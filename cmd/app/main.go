package main

import (
	"fmt"
	"net/http"

	"github.com/dao0203/todo-sample/config"
	"github.com/dao0203/todo-sample/internal/app"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Go API with Docker!")
}

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	app.Run(cfg)
}
