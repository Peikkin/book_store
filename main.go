package main

import (
	"net/http"
	"os"

	"github.com/Peikkin/book_store/routes"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)

	log.Info().Msg("Запуск сервера на порту :8080")
	log.Fatal().Err(http.ListenAndServe("localhost:8080", router)).Msg("Ошибка запуска сервера")
}
