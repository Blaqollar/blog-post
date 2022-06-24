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

	//Create a reference to the NewArticle Handler
	na := handlers.NewArticles(l)

	//Replace the serve mux with the GorrilaMux router package
	//Create a subrouter that is only applicable to a method then register your handler unto the route
	sm := mux.NewRouter()

	//This gives a subrouter that is filtered for the http verb "GET"
	//The .Subrouter converts this into a router so you can add your handle
	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", na.GetArticles)

	//You use the curly brackets to define a regexp in the URI
	putRouter := sm.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", na.UpdateArticles)

	//You can use different methods to call the http verbs
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", na.AddArticle)

	s := &http.Server{ //creating a server to tune the elements to timeout to avoid bad connections errors from server
		Addr:         ":9098",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	//Start the server
	go func() { //Handling the listen and serve in a go func so it deosn't block when integrating shutdown
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt) //(.Notify) This broadcasts a message to the sigChan whenever os.interrupt/kill is recieved
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig) //This sends a signal when the process is terminated

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc) // This waits until all the requests on the server has been completed before shutting down
}
