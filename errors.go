package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/samuelmasuy/protobuff/repo"
)

type Error struct {
	Status  int    `json:"status"`
	Content string `json:"content"`
}

func WriteJsonError(w http.ResponseWriter, err *Error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Status)
	jsonError := &repo.JsonError{
		Content: err.Content,
	}
	json.NewEncoder(w).Encode(jsonError)
}

func WriteProtoError(w http.ResponseWriter, err *Error) {
	log.Printf("Got an error: %+v\n", err)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(err.Status)

	protoError := &repo.ProtoError{
		Content: err.Content,
	}

	data, erro := proto.Marshal(protoError)
	if erro != nil {
		log.Printf("panic: %+v", erro)
	}

	_, erro = w.Write(data)
	if erro != nil {
		log.Printf("panic: %+v", erro)
	}
}

var (
	ErrBadRequest           = &Error{400, "Request body is not well-formed."}
	ErrNotAcceptable        = &Error{406, "Accept header must be set to 'application/json' or 'application/octet-stream'."}
	ErrUnsupportedMediaType = &Error{415, "Content-Type header must be set to: 'application/json' or 'application/octet-stream'."}
	ErrInternalServer       = &Error{500, "Something went wrong."}
)
