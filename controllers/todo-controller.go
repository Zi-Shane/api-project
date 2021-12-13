package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var TodoList []Todo

type Todo struct {
	Id   int64
	Item string
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024)) //io.LimitReader限制大小
	if err != nil {
		fmt.Println(err)
	}
	var addTodo Todo
	_ = json.Unmarshal(body, &addTodo) //轉為json
	TodoList = append(TodoList, addTodo)
	defer r.Body.Close()
	response := ApiResponse{"200", TodoList}

	ResponseWithJson(w, http.StatusOK, response) //回傳

}

func GetTodoById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queryId := vars["id"] //獲取url參數
	var targetTodo Todo
	for _, Todo := range TodoList { //比對TodoList內是否有符合的Todo
		intQueryId, _ := strconv.ParseInt(queryId, 10, 64) //將string轉為int64
		if Todo.Id == intQueryId {
			targetTodo = Todo
		}
	}
	response := ApiResponse{"200", targetTodo}
	ResponseWithJson(w, http.StatusOK, response)

}
