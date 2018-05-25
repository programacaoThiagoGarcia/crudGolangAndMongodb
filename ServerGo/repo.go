package main

//Conexão com banco de dados

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

//=============================================
//Methods by DataBase
//=============================================

const (
	local      = "localhost:27017"
	collection = "user"
	db         = "users"
)

//CreateNewRegister save new register in DataBase
func CreateNewRegister(user User) User {
	session, err := mgo.Dial(local)
	if err != nil {
		log.Println("Deu ruim na conexao")
		return User{}
	}
	defer session.Close()

	c := session.DB(db).C(collection)
	err = c.Insert(user)
	if err != nil {
		log.Println("Nao gravou")
	}
	return user
}

//GetAllUsers get all users in dataBase
func GetAllUsers() (Users, error) {
	session, err := mgo.Dial(local)
	if err != nil {
		log.Println("Could not connect to mongo ", err.Error())
	}
	defer session.Close()
	c := session.DB(db).C(collection)
	var listUser Users
	err = c.Find(bson.M{}).All(&listUser)
	return listUser, err

}

//GetOnlyUser return only one User
func GetOnlyUser(id string) Users {
	session, err := mgo.Dial(local)
	if err != nil {
		fmt.Println("Problema conexao")
	}
	defer session.Close()
	var list Users
	c := session.DB(db).C(collection)
	err = c.Find(bson.M{"nome": id}).All(&list)
	return list
}

//DeleteUserDB delete user with name method delete
func DeleteUserDB(name string) error {
	session, err := mgo.Dial(local)
	if err != nil {
		fmt.Println("Problema conexao")
		return err
	}
	defer session.Close()
	c := session.DB(db).C(collection)
	err = c.Remove(bson.M{"nome": name})
	if err != nil {
		return err
	}
	return nil
}
