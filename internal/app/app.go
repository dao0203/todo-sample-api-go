package app

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dao0203/todo-sample/config"
	v1 "github.com/dao0203/todo-sample/internal/controller/http/v1"
	mysql "github.com/dao0203/todo-sample/internal/database"
	"github.com/dao0203/todo-sample/internal/repository"
	"github.com/dao0203/todo-sample/internal/usecase"
	"github.com/dao0203/todo-sample/pkg/httpserver"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Go API with Docker!")
}

func Run(cfg *config.Config) {
	// MySQL connection
	mySql, err := mysql.New()
	if err != nil {
		panic(err)
	}
	defer mySql.Close()

	todoUseCase := usecase.New(
		repository.New(mySql),
	)

	httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port))
	v1.NewTodoRoutes(todoUseCase, httpServer.App)

	// start http server
	httpServer.Start()
	// wait signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		fmt.Println("Got signal:", s)
	case err = <-httpServer.Notify():
		fmt.Println("Got error:", err)
	}

	err = httpServer.Shutdown()
	if err != nil {
		fmt.Println("Got error:", err)
	}
}
