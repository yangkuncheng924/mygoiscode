 package main
 
 import (
     "database/sql"
     "fmt"
     "github.com/gin-gonic/gin"
     _ "github.com/go-sql-driver/mysql"
     "log"
     "strings"
 )
 
 //定义一个和表对应的结构体
 type Student struct {
     id    int
     no    string //学号
     name  string //姓名
     score uint   //成绩
 } 
 func main() {

    //1. 创建路由
     r := gin.Default()
 
     //2.
   r.POST("/form", func(c *gin.Context) {
 
         //表单参数设置默认值
         type1 := c.DefaultPostForm("type", "alert")
 
         //接收username,password
        no := c.PostForm("no")
         name := c.PostForm("name")
         score := c.PostForm("score")
 
        //DATABASE BASH
         db, err := InitDatabase()
         defer db.Close()
         if err != nil {
             log.Println(err)
             return
         }
         ////增
         insertSql := "insert into student(no, name, score) values(?, ?, ?)"
         err = Execute(db, insertSql, no, name, score)
         if err != nil {
             log.Printf("insert data error : %v\n", err)
            return
         }
 
         //查
         querySql := "select id, no, name, score from student where name = ?"
         rows, err := QueryData(db, querySql, "Jack")
         defer rows.Close()
         if err != nil {
            log.Printf("query data error:%v\n", err)
             return
         }
         s := new(Student)
 
         for rows.Next() {
             rows.Scan(&s.id, &s.no, &s.name, &s.score)
             log.Println(*s)
         }
 
         ////改
         updateSql := "update student set name = ? where no = ?"
         Execute(db, updateSql, "Rose", "123456")
         //
        ////删
        deleteSql := "delete from student where no = ? "
         Execute(db, deleteSql,"123456")
 
         c.String(200,
             fmt.Sprintf(type1, no, name, score))
 
     })
     //3. 监听
     r.Run()
 
 }
 
 //初始化数据库连接
 func InitDatabase() (*sql.DB, error) {
     //将数据转换成数据库url作为返回值
     url := strings.Join([]string{"root", ":", "root", "@tcp(", "127.0.0.1", ":", "3306", ")/", "starkgo"}, "")
     db, err := sql.Open("mysql", url)
     if err != nil {
         log.Printf("open database error:%v", err)
         return nil, err
     }
     return db, nil
 }
 
 //执行增、改、删任务
 func Execute(db *sql.DB, sql string, params ...interface{}) error {
     stmt, _ := db.Prepare(sql) //预编译
     defer stmt.Close()
     _, err := stmt.Exec(params...)
     if err != nil {
         log.Printf("execute sql error:%v\n", err)
         return err
     }
     log.Println("execute sql success")
     return nil
 }
 
 //查询数据库数据
 func QueryData(db *sql.DB, sql string, params ...interface{}) (*sql.Rows, error) {
     stmt, _ := db.Prepare(sql)
     defer stmt.Close()
     rows, err := stmt.Query(params...)
     if err != nil {
         log.Printf("query data error:%v", err)
         return nil, err
     }
     log.Println("query data success")
     return rows, nil
 }