package main

import (
	"fmt"
	"net/http"
)

const (
	TASK_TYPE_DB_MYSQL_QUERY = 1
	TASK_TYPE_DB_MYSQL_EXEC = 2
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Print("URL: ")
	fmt.Println(r.URL)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	str := fmt.Sprintf(`
	{
		"type": %d,
		"config": {
			"type": "mysql",
			"dsn": "root:root@tcp(localhost:3306)/"
		},
		"payload": "SELECT * FROM test"
	}
	`, TASK_TYPE_DB_MYSQL_QUERY)
	fmt.Fprintf(w, str)
}

func main() {

	fmt.Println("Starting task server...")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8109", nil)
}
