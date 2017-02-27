package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/context"
	"github.com/samuelmasuy/protobuff/repo"
)

// Middlewares

func contentTypeHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") == "application/json" {
			context.Set(r, "content-type", "json")
			defer func() {
				if err := recover(); err != nil {
					log.Printf("panic: %+v", err)
					WriteJsonError(w, ErrInternalServer)
				}
			}()
		} else if r.Header.Get("Content-Type") == "application/octet-stream" {
			context.Set(r, "content-type", "proto")
			defer func() {
				if err := recover(); err != nil {
					log.Printf("panic: %+v", err)
					WriteProtoError(w, ErrInternalServer)
				}
			}()
		} else {
			w.WriteHeader(415)
			w.Write([]byte("Content-Type header must be set to: 'application/json' or 'application/octet-stream'"))
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func bodyHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		contentType := context.Get(r, "content-type")
		if contentType == "proto" {
			request := repo.ProtoAnimalAnswer{}

			data, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				log.Println(err)
				WriteProtoError(w, ErrBadRequest)
				return
			}

			if err := proto.Unmarshal(data, &request); err != nil {
				log.Println("unmarshalling error", err)
				WriteProtoError(w, ErrBadRequest)
				return
			}

			if next != nil {
				context.Set(r, "body", request)
				next.ServeHTTP(w, r)
			}
		} else if contentType == "json" {
			request := &repo.JsonAnimalAnswer{}
			err := json.NewDecoder(r.Body).Decode(request)
			defer r.Body.Close()
			if err != nil {
				WriteJsonError(w, ErrBadRequest)
				return
			}

			if next != nil {
				context.Set(r, "body", request)
				next.ServeHTTP(w, r)
			}
		}
	}
	return http.HandlerFunc(fn)
}

func loggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}

	return http.HandlerFunc(fn)
}
