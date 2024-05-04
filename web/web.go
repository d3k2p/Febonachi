package web

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"main.go/logic"
)

func handleServe(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile("./web/index.html")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	w.Write(file)
}

func handleForm(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile("./web/result.html")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	fs := logic.FibonacciService{}

	number, err := strconv.Atoi(r.PostFormValue("numberValue"))
	if err != nil || number < 0 {
		w.Write([]byte("Введённое число некорректно!!!"))
		return
	}

	w.Write([]byte(fmt.Sprintf("Введённое число: %d\n", number)))

	if fs.IsFibonacci(number) {
		prev, next := fs.GetAdjacentFibonacci(number)
		w.Write([]byte(fmt.Sprintf("Предыдущее число: %d", prev)))
		w.Write([]byte(fmt.Sprintf("Следующее число: %d", next)))
	} else {
		closest := fs.GetNearestFibonacci(number)
		w.Write([]byte(fmt.Sprintf("Ближайшее число: %d", closest)))
	}

	w.Write([]byte("<a href=\"/\">Вернуться</a>"))
}

func StartServer() {
	http.HandleFunc("/", handleServe)
	http.HandleFunc("/result", handleForm)

	log.Fatal(http.ListenAndServe(":8181", nil))
}
