package main

import (
	"fmt"
	"html/template"
	//_ "github.com/go-sql-driver/mysql"
	"net/http"
)

//type User struct {
//	Name string `json:"name"`
//	Age  uint16 `json:"age"`
//}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "create", nil)
}

func handlerFunc() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handlerFunc()
	//db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/goFirstLabDB")
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()

	// Установка данных
	//insert, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES('Denis', 26)")
	//if err != nil {
	//	panic(err)
	//}
	//defer insert.Close()

	// Выборка данных
	//res, err := db.Query("SELECT `name`,`age` FROM `users`")
	//if err != nil {
	//	panic(err)
	//}
	//for res.Next() {
	//	var user User
	//	err = res.Scan(&user.Name, &user.Age)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(fmt.Sprintf("User: %s with age %d", user.Name, user.Age))
	//}
}
