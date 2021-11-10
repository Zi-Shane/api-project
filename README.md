# api-porject

## preparation

1. Generate the Self-Signed certificate and private key for your domain name.

```sh
openssl genrsa -out tls.key 2048

# .csr for domain
openssl req -nodes -new -key tls.key -subj "/CN=<your.domain.com>/" -out tls.csr

openssl x509 -req -sha256 -days 3650 -in tls.csr -signkey tls.key -out tls.crt
```

put `tls.crt` and `tls.key` into the folder `./secret/<DOMAINNAME>` 

2. Environment variables

```
DOMAINNAME=<your.domain.com>
```

---

## VSCode setting

You can store environment variables in `.env` file, and load them when running the program in vscode.

```
{
    "configurations": [
        {
            "name": "homepc",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${fileDirname}",
            "envFile": "${workspaceFolder}/secret/homepc.shanesnotes.tw/.env"
        }
    ]
}
```

---

## Deploy to k8s

To deploy this application to k8s, you have to prepare resources before you deploy to k8s.

### Secret
```
kubectl create secret tls my-tls-secret \
--key ./tls.key \
--cert ./tls.crt

kubectl describe secrets/my-tls-secret
```

### ConfigMap
```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: myconfigyaml
data:
  domainname: <your.domain.com>
  bottoken: xxxxx:aaaaaaaaaaaaaaaaaaaa
```

### Service
- LoadBalancer
```yaml
apiVersion: v1
kind: Service
metadata:
  name: test-server-service
spec:
  type: LoadBalancer
  ports:
  - port: 8443
  selector:
    app: test-server
```
- NodePort
```yaml
apiVersion: v1
kind: Service
metadata:
  name: test-server-service
spec:
  type: NodePort
  selector:
    app: test-server
  ports:
    - protocol: TCP
      port: 8443
```

After you apply all resources, you can correctly deploy your application to k8s.

### Deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: test-server
spec:
  replicas: 1
  selector:
    matchLabels:
      app: test-server
  template:
    metadata:
      labels:
        app: test-server
    spec:
      nodeSelector:
        "beta.kubernetes.io/os": linux
      containers:
      - name: test-server
        image: tgbotreg.azurecr.io/api-example:v1.0
        env:
        - name: DOMAINNAME
          valueFrom:
            configMapKeyRef:
              name: myconfigyaml
              key: domainname
        - name: TELEGRAM_APITOKEN
          valueFrom:
            configMapKeyRef:
              name: myconfigyaml
              key: bottoken
        volumeMounts:
        - name: certificate-volume
          mountPath: /app/secret/<your.domain.com>
          readOnly: true
        ports:
        - containerPort: 8443
      volumes:
      - name: certificate-volume
        secret:
          secretName: my-tls-secret
```
