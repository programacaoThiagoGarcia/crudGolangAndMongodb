package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//=====================================================
// MÉTODOS PARA BANCO DE DADOS
//=====================================================

//CreateUser cria um usuário no banco
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(400)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}
	u := CreateNewRegister(user)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(u); err != nil {
		panic(err)
	}

}

//GetUsers busca a lista de usuários no banco
func GetUsers(w http.ResponseWriter, r *http.Request) {

	var users Users
	var err error
	users, err = GetAllUsers()
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		panic(err)
	}

}

//GetUserName filtra pelo nome do usuário
func GetUserName(w http.ResponseWriter, r *http.Request) {
	//name
	vars := mux.Vars(r)
	userName := vars["name"]
	u := GetOnlyUser(userName)
	if err := json.NewEncoder(w).Encode(u); err != nil {
		fmt.Println("Nao enviou")
	}

}

//DeleteUser deleta usuário enviado no body  method Delete
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user User
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		panic(err)
	}
	u := DeleteUserDB(user.Nome)
	if u != nil {
		e := fmt.Sprint("Erro: ", u)
		var warning = Warning{
			true,
			e,
		}
		_ = json.NewEncoder(w).Encode(warning)
		return
	}
	var warning = Warning{
		false,
		"Usuário deletado com sucesso",
	}
	_ = json.NewEncoder(w).Encode(warning)

}
