package main

import (
	"fmt"
	"gopodcast/graphql"
	"gopodcast/graphql/generated"
	"gopodcast/itunes"
	"gopodcast/utils"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/websocket"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/cors"
)

var (
	cacheTTL, _ = strconv.ParseInt(os.Getenv("CACHE_TTL_IN_HOURS"), 10, 64)
	port        = os.Getenv("APP_PORT")
	url         = os.Getenv("APP_URL")
)

func main() {
	router := chi.NewRouter()
	cache, err := utils.NewCache(os.Getenv("REDIS_ADDRESS"), time.Duration(cacheTTL)*time.Hour)
	if err != nil {
		log.Fatalf("Cannot create APQ redis cache: %v", err)
	}
	resolver := &graphql.Resolver{
		Api: &itunes.API{},
	}
	c := generated.Config{Resolvers: resolver}
	queryHandler := handler.GraphQL(
		generated.NewExecutableSchema(c),
		handler.EnablePersistedQueryCache(cache),
		handler.ComplexityLimit(200),
		handler.WebsocketUpgrader(
			websocket.Upgrader{
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			}),
	)
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{fmt.Sprintf("%s:%s", url, port)},
		AllowCredentials: true,
	}).Handler)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Handle("/", handler.Playground("GraphQL playground", "/query"))
	router.Handle("/query", queryHandler)
	log.Printf("connect to %s:%s/ for GraphQL playground", url, port)

	http.ListenAndServe(":"+port, router)
}
