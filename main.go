package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

var storages []Question

func main() {

	//Basic Auth Accounts
	accounts := make([]string, 1)
	accounts = append(accounts, "hai:pwd")

	//Home
	http.HandleFunc("/", Log(Home))

	//Handle request on question endpoint
	http.HandleFunc("/v1/question", Group(Log(QuestionAPI), EnsureAuthorize(accounts)))

	log.Fatal(http.ListenAndServe(":1313", nil))
}

func Home(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-type", "text/html; charset=utf-8")
	writer.WriteHeader(http.StatusOK)

	_, _ = fmt.Fprintf(writer, "<h2 align=\"center\">Welcome to Golang Viet Nam Workshop 2 - FAQ app</h2>")

	return
}

func QuestionAPI(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		GetQuestion(writer, request)
		break
	case http.MethodPost:
		CreateQuestion(writer, request)
		break
	default:
		AbortRequest(writer)
		break
	}
}

func GetQuestion(writer http.ResponseWriter, request *http.Request) {
	//question?=ids={question_id}
	ids, err := request.URL.Query()["id"]
	if !err || len(ids) < 1 {
		AbortRequest(writer)
		return
	}
	questionId, _ := strconv.Atoi(ids[0])
	fmt.Println("client retrieving the detail of the question that has id is ", uint32(questionId))

	question, error2 := findQuestion(questionId)
	if error2 != nil {
		writer.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(writer).Encode(EntityError{
			Code:   600,
			Phrase: "The question could not be found",
		})
		return
	}

	_ = json.NewEncoder(writer).Encode(question)
}

func CreateQuestion(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		AbortRequest(writer)
		return
	}

	defer request.Body.Close()
	if err != nil {
		AbortRequest(writer)
		return
	}

	fmt.Println("client is attempting to submit a question that has body ", string(body))

	input := struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}{}

	err = json.Unmarshal(body, &input)
	if err != nil {
		AbortRequest(writer)
		return
	}

	question, err := CreateNewQuestion(input)
	if err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	_ = json.NewEncoder(writer).Encode(question)
}

func CreateNewQuestion(params struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}) (Question, error) {
	question := Question{
		ID:      allocQuestionID(),
		Title:   params.Title,
		Content: params.Content,
	}

	storages = append(storages, question)

	fmt.Println(storages)

	return question, nil
}

func AbortRequest(writer http.ResponseWriter) {
	http.Error(writer, QuestionError{Reason: "Invalid request"}.Error(), http.StatusUnprocessableEntity)
	return
}

func findQuestion(id int) (*Question, error) {
	for _, question := range storages {
		if question.ID == id {
			return &question, nil
		}
	}

	return nil, EntityError{Phrase: "The Question could not be found."}
}

func allocQuestionID() int {
	result := len(storages) + 1
	for {
		randID := rand.Intn(1000)
		_, err := findQuestion(randID)
		if err != nil {
			result = randID
			break
		}
	}

	return result
}
