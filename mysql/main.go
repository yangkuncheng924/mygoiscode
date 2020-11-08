package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)
type Person struct{
	Id int
	Name string
	Age	int
}

func main(){
	consts := "root:root@tcp(127.0.0.1:3306)/starkgo"
	db,err := sql.Open("mysql",consts)
	if err!= nil{
		log.Fatal(err.Error())
		return
	}
	// _,err = db.Exec("create table person(" +
	// "id int auto_increment primary key," +
	// "name varchar(12) not null," +
	// "age int default 1" +
	// ");")
	// if err != nil{
	// 	log.Fatal(err.Error())
	// 	return
	// }
	_,err = db.Exec("insert into person(name,age)"+
	"values(?,?);","小杨",18)

	if err != nil{
		log.Fatal(err.Error())
	}else{
		fmt.Println("数据已插入")
	}
	Ro,err := db.Query("select id, name,age from person")
	if err != nil {
		log.Fatal(err.Error())
	}
	Scane:
	if Ro.Next(){
		person := new(Person)
		err = Ro.Scan(&person.Id,&person.Name,&person.Age)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println(person.Id,person.Name,person.Age)
goto Scane
	}
}
}