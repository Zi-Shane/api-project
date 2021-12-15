# api-porject

## How to Run?

### MariaDB

I use `docker-compose` to quickly start database when developing.

### Environment variables

```
DB_CONN=username:password@tcp(127.0.0.1:3306)/nation
```

---

## VSCode setting

You can store environment variables in `.env` file, and load them when running the program in vscode.

```json
{
    "configurations": [
        {
            "name": "homepc",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/main.go"
            "envFile": "${workspaceFolder}/secret/homepc.shanesnotes.tw/.env"
        }
    ]
}
```

---

## Architecture diagram

This Repository is the implementation of the backend in the diagram.

![](https://i.imgur.com/lNeXLJZ.png)

---

## API Document

### All Path
```
GET    /api/language/:id              --> api/controllers.GetLanguage 
GET    /api/languageRange/:start/:end --> api/controllers.GetLanguages 
GET    /api/countryUselanguage        --> api/controllers.GetCountryUesdLanguages 
POST   /api/language                  --> api/controllers.AddLanguage 
DELETE /api/language/:language        --> api/controllers.RemoveLanguage 
PUT    /api/language                  --> api/controllers.UpdateLanguage
```


### `GET    /api/hostname`
- query: `api/hostname`
- response: 
```json
{
  "hostname": "shane-ubuntu"
}
```

### `GET    /api/GetLanguage/:id`
- query: `api/GetLanguage/1`
- response: 
```json
{
  "message": "success",
  "result": {
    "Language_id": 1,
    "Language": "Dutch"
  }
}
```

### `GET    /api/GetLanguageRange/:start/:end`
- query: `api/GetLanguageRange/10/20`
- response: 
```json
{
  "message": "success",
  "result": [
    {
      "Language_id": 10,
      "Language": "Ambo"
    },
    {
      "Language_id": 11,
      "Language": "Chokwe"
    },
    {
      "Language_id": 12,
      "Language": "Kongo"
    },
    {
      "Language_id": 13,
      "Language": "Luchazi"
    }
  ]
}
```
### `GET    /api/GetCountryUselanguage`
- query: `api/GetCountryUselanguage/?country=Canada`
- response: 
```json
{
  "message": "success",
  "result": [
    {
      "CountryName": "Canada",
      "Language": "Dutch"
    },
    {
      "CountryName": "Canada",
      "Language": "English"
    },
    ...
    {
      "CountryName": "Canada",
      "Language": "Chinese"
    },
    {
      "CountryName": "Canada",
      "Language": "Eskimo Languages"
    },
    {
      "CountryName": "Canada",
      "Language": "Punjabi"
    }
  ]
}
``` 

### `POST   /api/AddLanguage`
- query: `api/AddLanguage`
    - body: 
    ```json
    [
        {
            "Language_id": 500,
            "Language": "Test500"
        },
        {
            "Language_id": 501,
            "Language": "Test501"
        }
    ]
    ```
- response: 
```json
{
  "message": "success",
  "rowsAffected": 2
}
```

### `/api/RemoveLanguage/:language`

- query: `api/RemoveLanguage/Test501`
- response: 
```json
{
  "message": "success",
  "rowsAffected": 1
}
```

### `/api/UpdateLanguage`
- query: `api/UpdateLanguage`
    - body:
    ```json
    {
        "Language_id": 500,
        "Language": "Updated500"
    }
    ```
- response: 
```json
{
  "message": "success",
  "rowsAffected": 1
}
```

---

## Deploy to k8s

To deploy this application to k8s, you have to prepare Secret before you deploy to k8s.

Then you can deploy `manifests/deployment.yml` and `manifests/service.yml`

### Secret
```yaml
apiVersion: v1
data:
  dbconn: base64-encoded-connection-info
kind: Secret
metadata:
  name: mariadb-conn
  namespace: api-backend
type: Opaque
```

### How to generate secret?
- Use `kustomization`: 
    - [Official Tutorial](https://kubernetes.io/docs/tasks/manage-kubernetes-objects/kustomization/)
