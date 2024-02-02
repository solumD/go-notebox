package main

import (
	"log"
	"net/http"
)

func main() {
	// Используется функция http.NewServeMux() для инициализации нового рутера, затем
	// функция "home" регистрируется как обработчик для URL-шаблона "/",
	// аналогично "home" регистрируются showNote и createNote
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/note", showNote)
	mux.HandleFunc("/note/create", createNote)

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

//
