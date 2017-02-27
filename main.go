package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/justinas/alice"
	"github.com/rs/cors"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 5000, "Specify the port to listen to.")
	flag.Parse()

	appC := appContext{}
	commonHandlers := alice.New(context.ClearHandler, loggingHandler, contentTypeHandler)

	router := NewRouter()
	router.Get("/", commonHandlers.ThenFunc(appC.mainHandler))
	router.Post("/answer", commonHandlers.Append(bodyHandler).ThenFunc(appC.answerHandler))

	router.Get("/animals", commonHandlers.ThenFunc(appC.animalsHandler))

	corsHandler := cors.Default().Handler(router)
	log.Printf("Listening on port %d", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), corsHandler); err != nil {
		log.Fatal(err)
	}
	// log.Printf("Listening on port %d", port)
	// if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
	// 	log.Fatal(err)
	// }
}
