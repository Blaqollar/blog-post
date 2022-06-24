package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"github.com/blaqollar/blog-post/app/handlers"
)

func main() {
	l := log.New(os.Stdout, "blog", log.LstdFlags)

	na := handlers.NewArticles(l)

	sm := mux.NewRouter()

	
	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", na.GetArticles)

	
	putRouter := sm.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", na.UpdateArticles)

	
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", na.AddArticle)

	s := &http.Server{ 
		Addr:         ":9098",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	go func() { 
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt) 
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig) 

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc) 
}
