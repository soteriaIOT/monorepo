package main

import (
	"context"
	"sync"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/arora-aditya/monorepo/application-server/data"
	"github.com/arora-aditya/monorepo/application-server/graph"
	"github.com/arora-aditya/monorepo/application-server/graph/generated"
	"github.com/arora-aditya/monorepo/application-server/auth"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"

)

const defaultPort = "8081"

func main() {
	ctx, kafka_cancel := context.WithCancel(context.Background())

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	var wg sync.WaitGroup
	wg.Add(1)

	d := data.NewDemoRepository()
	go d.ReadMessage(ctx, &wg)

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	router.Use(auth.JwtMiddleware())

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := &graph.Resolver{
		Repository: d,
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: r}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	s := &http.Server{
		Addr: ":" + defaultPort,
		Handler: router,
	}
	
	// Start the server
	go func() {
		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
		err := s.ListenAndServe()
		if err == http.ErrServerClosed {
			log.Println("server closed")
		} else {
			log.Fatalf("server closed: %v", err)
		}
	}()


	// Wait for interrupt signal to gracefully shutdown the server
	<-done
	log.Println("Server Stopped")

	// Shutdown kafka gracefully
	kafka_cancel()
	log.Println("Waiting for kafka reader to finish")

	wg.Wait()
	log.Println("Done waiting")


	ctx_timeout, timeout := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		timeout()
	}()


	if err := s.Shutdown(ctx_timeout); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	log.Println("Shutdown")
}
