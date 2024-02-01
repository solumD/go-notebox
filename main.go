package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Создается функция-обработчик "home", которая записывает байтовый слайс, содержащий
// текст "Привет из Snippetbox" как тело ответа.
func home(w http.ResponseWriter, r *http.Request) {

	// Проверяется, если текущий путь URL запроса точно совпадает с шаблоном "/". Если нет, вызывается
	// функция http.NotFound() для возвращения клиенту ошибки 404.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return //Завершение работы клиента, чтобы функция не вывела "Привет из NoteBox"
	}

	w.Write([]byte("Привет из NoteBox"))
}

// Обработчик для отображения содержимого заметки.
func showNote(w http.ResponseWriter, r *http.Request) {
	// Извлекаем значение параметра id из URL и попытаемся
	// конвертировать строку в integer используя функцию strconv.Atoi(). Если его нельзя
	// конвертировать в integer, или значение меньше 1, возвращаем ответ
	// 404 - страница не найдена!

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Отображение выбранной записи с ID %d...", id)
}

// Обработчик для создания новой заметки.
func createNote(w http.ResponseWriter, r *http.Request) {
	// Используем r.Method для проверки, использует ли запрос метод POST или нет.
	// http.MethodPost является строкой "POST"
	if r.Method != http.MethodPost {
		// Используем метод Header().Set() для добавления заголовка 'Allow: POST' в
		// карту HTTP-заголовков. Первый параметр - название заголовка, а
		// второй параметр - значение заголовка.
		w.Header().Set("Allow", http.MethodPost)

		// Используем функцию http.Error() для отправки кода состояния 405 с соответствующим сообщением.
		http.Error(w, "Метод запрещен!", 405)
		return
	}

	w.Write([]byte("Создания новой записи..."))
}

func main() {
	// Используется функция http.NewServeMux() для инициализации нового рутера, затем
	// функцию "home" регистрируется как обработчик для URL-шаблона "/",
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
