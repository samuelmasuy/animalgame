package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/context"
	"github.com/samuelmasuy/animalgame/repo"
)

// Main handlers
type appContext struct {
	tree *Tree
}

func (c *appContext) mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	_, err := w.Write([]byte(`{"name":"Alice","body":"Hello"}`))
	if err != nil {
		panic(err)
	}
}

func (c *appContext) animalsHandler(w http.ResponseWriter, r *http.Request) {
	c.tree = NewAnimalTree()
	animals := GetAnimals(c.tree)
	contentType := context.Get(r, "content-type")
	if contentType == "proto" {
		response := &repo.ProtoAnimalList{
			Content: animals,
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(200)

		data, err := proto.Marshal(response)
		if err != nil {
			panic(err)
		}

		_, err = w.Write(data)
		if err != nil {
			panic(err)
		}
	} else if contentType == "json" {
		response := &repo.JsonAnimalList{
			Content: animals,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response)
	}
}

func (c *appContext) answerHandler(w http.ResponseWriter, r *http.Request) {
	contentType := context.Get(r, "content-type")
	if contentType == "proto" {
		request := context.Get(r, "body").(repo.ProtoAnimalAnswer)
		id := request.Id
		tree, question := c.tree.Next(request.Content)
		if question == "" {
			id = -1
			if request.Content {
				question = fmt.Sprintf("You see, I knew you were thinking about a %s!", c.tree.Value)
			} else {
				question = fmt.Sprintf("Oops, I was sure you were thinking of a %s!", c.tree.Value)
			}
		} else {
			c.tree = tree
			if tree.IsLeaf() {
				question = fmt.Sprintf("Is it a %s?", question)
			}
		}
		response := &repo.ProtoAnimalQuestion{
			Id:      id,
			Content: question,
		}

		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(201)

		data, err := proto.Marshal(response)
		if err != nil {
			panic(err)
		}

		_, err = w.Write(data)
		if err != nil {
			panic(err)
		}
	} else if contentType == "json" {
		request := context.Get(r, "body").(*repo.JsonAnimalAnswer)
		id := request.Id
		tree, question := c.tree.Next(request.Content)
		if question == "" {
			id = -1
			if request.Content {
				question = fmt.Sprintf("You see, I knew you were thinking about a %s!", c.tree.Value)
			} else {
				question = fmt.Sprintf("Oops, I was sure you were thinking of a %s!", c.tree.Value)
			}
		} else {
			c.tree = tree
			if tree.IsLeaf() {
				question = fmt.Sprintf("Is it a %s?", question)
			}
		}
		response := &repo.JsonAnimalQuestion{
			Id:      id,
			Content: question,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(response)
	}
}
