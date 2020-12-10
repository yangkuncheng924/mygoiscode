package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct{
	ID int
	Name string
	Age  string
}
 //初始化数据库连接
 func InitDatabase() (*sql.DB, error) {
     //将数据转换成数据库url作为返回值
    url := strings.Join([]string{"root", ":", "root","@tcp(", "127.0.0.1", ":", "3306", ")/", "starkgo"}, "")
     db, err := sql.Open("mysql", url)
     if err != nil {
         log.Printf("open database error:%v", err)
         return nil, err
	 }
	 fmt.Printf("调用数据库初始化")
     return db, nil
 }

 func Execute(db *sql.DB, sql string, params ...interface{}) error {
	     stmt, _ := db.Prepare(sql) //预编译
        defer stmt.Close()
        _, err := stmt.Exec(params...)
        if err != nil {
            log.Printf("execute sql error:%v\n", err)
            return err
        }
		log.Println("execute sql success")
		
		fmt.Printf("调用EXcute")
        return nil
  }

func main11(){
	r := gin.Default()

	r.POST("/form",func(c *gin.Context){
	//	tpe1 := c.DefaultPostForm("type","alert")

		id := c.PostForm("no")
		name := c.PostForm("name")
		age := c.PostForm("age")

		 
          //DATABASE BASH
          db, err := InitDatabase()
          defer db.Close()
          if err != nil {
              log.Println(err)
              return
          }
          ////增
          insertSql := "insert into student(id, name, age) values(?, ?, ?)"
          err = Execute(db, insertSql, id, name, age)
          if err != nil {
              log.Printf("insert data error : %v\n", err)
              return
          }
	})
	r.Run(":9080")
}