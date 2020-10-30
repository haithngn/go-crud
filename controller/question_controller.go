package controller

import (
	"encoding/json"
	"fmt"
	"github.com/haithngn/go-crud/db"
	"github.com/haithngn/go-crud/model"
	"github.com/haithngn/go-crud/value"
	"io/ioutil"
	"net/http"
	"strconv"
)

type QuestionError struct {
	Reason string
}

func (error QuestionError) Error() string {
	return error.Reason
}

type QuestionController struct {
	Store db.Storage
}

func (controller *QuestionController) GetQuestion(writer http.ResponseWriter, request *http.Request) {
	//question?=ids={question_id}
	ids, err := request.URL.Query()["id"]
	if !err || len(ids) < 1 {
		controller.abortRequest(writer)
		return
	}
	questionId, _ := strconv.Atoi(ids[0])
	fmt.Println("client retrieving the detail of the question that has id is ", uint32(questionId))

	question, error2 := controller.findQuestion(questionId)
	if error2 != nil {
		writer.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(writer).Encode(value.EntityError{
			Code:   600,
			Phrase: "The question could not be found",
		})
		return
	}

	_ = json.NewEncoder(writer).Encode(question)
}

func (controller *QuestionController) CreateQuestion(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		controller.abortRequest(writer)
		return
	}

	defer request.Body.Close()
	if err != nil {
		controller.abortRequest(writer)
		return
	}

	fmt.Println("client is attempting to submit a question that has body ", string(body))

	input := struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}{}

	err = json.Unmarshal(body, &input)
	if err != nil {
		controller.abortRequest(writer)
		return
	}

	question, err := controller.createNewQuestion(input)
	if err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	_ = json.NewEncoder(writer).Encode(question)
}

func (controller *QuestionController) createNewQuestion(params struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}) (*model.Question, error) {
	question, err := controller.Store.Create(params)

	if err != nil {
		return nil, err
	}

	return question, nil
}

func (controller *QuestionController) findQuestion(id int) (*model.Question, error) {
	return controller.Store.Find(id)
}

func (controller *QuestionController) abortRequest(writer http.ResponseWriter) {
	http.Error(writer, QuestionError{Reason: "Invalid request"}.Error(), http.StatusUnprocessableEntity)
	return
}
