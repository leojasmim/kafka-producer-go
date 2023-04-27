package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/leojasmim/kafka-producer-go/application/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Post("/message", handlers.SendMessage)

	server := &http.Server{
		Addr:    os.Getenv("KS_APPLICATION_PORT"),
		Handler: r,
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			interrupt <- syscall.SIGTERM
			fmt.Println("err: ", err)
		}
	}()

	fmt.Println(logo())
	fmt.Println("Application is up and running => ApplicationPort: " + os.Getenv("KS_APPLICATION_PORT"))
	shutdown(server, interrupt)
}

func shutdown(server *http.Server, interrupt chan os.Signal) {
	<-interrupt
	fmt.Println("Shutdown Down", "Start Shutdown server", nil)
	fmt.Println("--> starting shutdown server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("--> shutdown completed")
}

func logo() string {
	return `
	 _  __      __ _          _____                       _ 
	| |/ /     / _| |        / ____|                     | |     
	| ' / __ _| |_| | ____ _| (___   __ _ _ __ ___  _ __ | | ___ 
	|  < / _' |  _| |/ / _' |\___ \ / _' | '_ ' _ \| '_ \| |/ _ \
	| . \ (_| | | |   < (_| |____) | (_| | | | | | | |_) | |  __/
	|_|\_\__,_|_| |_|\_\__,_|_____/ \__,_|_| |_| |_| .__/|_|\___|
												   | |           			           
												   |_|				           																		 
	`
}
