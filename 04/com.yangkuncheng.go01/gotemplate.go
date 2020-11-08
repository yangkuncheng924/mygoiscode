package main

import(
	"fmt"
	"html/template"
	"net/http"

)
type User struct{
	Name    string
	Age     int
	Genader string
}

func sayHello(w http.ResponseWriter,r *http.Request){
	//解析模板
	t,err := template.ParseFiles("./gotemp.tmpl")
	if err!=nil {
		fmt.Println("error template")
		return
	}
	//渲染模ban


	u1 := User{
		Name : "xiaoyang",
		Age  : 19,
		Genader : "吃吃吃",
	}
	
	m1 := map[string]interface{}{ 
		"Name":"yangge",
		"Age":"20",
		"Genader":"好",
	}

	err = t.Execute(w,map[string]interface{}{
		"u1":u1,
		"m1":m1,
	}) //map、struct、string 
	if err!=nil{
		fmt.Println("渲染模板发生意外~")
		return
	}


}

func main(){
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090",nil)
	if err!=nil {
		fmt.Println("http error",err)
		return
	}
}