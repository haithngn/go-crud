package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Storage struct {
	Questions []Question
}

type RequestHandler struct {
	Store Storage
}

func (handler *RequestHandler) Home(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Requested At: ", time.Now().String())
	fmt.Println("======== REQUEST LINE ========")
	fmt.Println("Method: ", request.Method)
	fmt.Println("Host: ", request.Host)
	fmt.Println("Path: ", request.URL.Path)
	fmt.Println("Query: ", request.URL.RawQuery)
	fmt.Println("Fragment: ", request.URL.RawFragment)
	fmt.Println("======== REQUEST HEADER ========")
	fmt.Println("User-Agent: ", request.Header.Get("User-Agent"))
	fmt.Println("Content-type: ", request.Header.Get("Content-type"))
	fmt.Println("Connection: ", request.Header.Get("Connection"))
	fmt.Println("Accept: ", request.Header.Get("Accept"))
	fmt.Println("Accept-Encoding: ", request.Header.Get("Accept-Encoding"))

	fmt.Println("======== MESSAGE BODY ========")

	message, _ := ioutil.ReadAll(request.Body)
	fmt.Printf("Message: %s\n", message)

	writer.Header().Set("Content-type", "text/html; charset=utf-8")
	writer.WriteHeader(http.StatusOK)

	_, _ = fmt.Fprintf(writer, "<h1 align=\"center\">Welcome to Golang Viet Nam Workshop 2 - FAQ app</h1>")

	return
}

func (handler *RequestHandler) Question(writer http.ResponseWriter, request *http.Request) {
	//Re-direct the quest base on the http method
	switch request.Method {
	case http.MethodGet:
		handler.getQuestion(writer, request)
		break
	case http.MethodPost:
		handler.createQuestion(writer, request)
		break
	default:
		handler.abortRequest(writer)
		break
	}
}

func (handler *RequestHandler) getQuestion(writer http.ResponseWriter, request *http.Request) {
	//question?=ids={question_id}
	ids, err := request.URL.Query()["id"]
	if err != true || len(ids) < 1 {
		handler.abortRequest(writer)
		return
	}
	questionId, _ := strconv.Atoi(ids[0])
	fmt.Println("client retrieving the detail of the question that has id is ", uint32(questionId))

	question, qErr := handler.findQuestion(questionId)
	if qErr != nil {
		writer.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(writer).Encode(EntityError{
			Code:   600,
			Phrase: "The question could not be found",
		})
		return
	}

	_ = json.NewEncoder(writer).Encode(question)
}

func (handler *RequestHandler) createQuestion(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		handler.abortRequest(writer)
	}

	defer request.Body.Close()
	if err != nil {
		handler.abortRequest(writer)
		return
	}

	fmt.Println("client is attempting to submit a question that has body ", string(body))

	input := struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}{}

	err = json.Unmarshal(body, &input)
	if err != nil {
		handler.abortRequest(writer)
		return
	}

	question, err := handler.createNewQuestion(input)
	if err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	_ = json.NewEncoder(writer).Encode(question)
}

//Handle Unexpected request
func (handler *RequestHandler) abortRequest(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusBadRequest)
	return
}

func (handler *RequestHandler) createNewQuestion(params struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}) (Question, error) {
	question := Question{
		ID:      handler.AllocQuestionID(),
		Title:   params.Title,
		Content: params.Content,
	}

	handler.Store.Questions = append(handler.Store.Questions, question)

	fmt.Println(handler.Store.Questions)

	return question, nil
}

func (handler *RequestHandler) findQuestion(id int) (*Question, error) {
	for _, question := range handler.Store.Questions {
		if question.ID == id {
			return &question, nil
		}
	}

	return nil, EntityError{Phrase: "The Question could not be found."}
}

func (handler *RequestHandler) AllocQuestionID() int {
	result := len(handler.Store.Questions) + 1
	for {
		randID := rand.Intn(1000)
		_, err := handler.findQuestion(randID)
		if err != nil {
			result = randID
			break
		}
	}

	return result
}
