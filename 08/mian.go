
 package main
 
 import (
     "database/sql"
     "fmt"
     _ "github.com/go-sql-driver/mysql"
 )
 
 // 定义一个全局对象db
 var db *sql.DB
 //定义结构体
 type User struct {
     Id   int
     Name string
     Age  string
 }
 
 //初始化数据库
 func InitDB() (err error) {
     //连接数据库
     dsn := "root:root@tcp(localhost)/starkgo"
     db, err = sql.Open("mysql", dsn)
     //错误处理
     if err != nil {
         return err
     }
     //检验dsn是否正确
     err = db.Ping()
     if err != nil {
         return err
     }
     return nil
 }
 func main(){
	var i int = 10

	for i=0;i<10;i++{ 
	InitDB() 
	//InsertDemo() 
	DeleteDemo()
	}
	}
	 func InsertDemo() {
		     //sql语句
 sqlStr := "insert into user(id,name,age) value(?,?,?)"
		     //Exec执行一次命令（包括查询、删除、更新、插入等），返回的Result是对已执行的SQL命令的总结。参数args表示query中的占位参数
 ret, err := db.Exec(sqlStr, "赵爸", "31")
if err != nil {
		        fmt.Println(err)
		        return
 }
 fmt.Println(ret)
		     // 新插入数据的id是theID
		     theID, err := ret.LastInsertId()
		     if err != nil {
		         fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		         return
		     }
		     fmt.Printf("insert success, the id is %d.\n", theID)
}
 //删除数据
 func DeleteDemo() {
	     sqlStr := "delete from user where id=?"
	     ret, err := db.Exec(sqlStr, 4)
	     if err != nil {
	         fmt.Println(err)
	        return
	     }
	     n, err := ret.RowsAffected()
	     if err != nil {
	         fmt.Printf("get RowsAffected failed, err:%v\n", err)
	         return
	     }
	     fmt.Printf("delete success, affected rows:%d\n", n)
	 
	 }