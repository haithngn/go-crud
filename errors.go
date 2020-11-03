package main

type EntityError struct {
	Code   int32
	Phrase string
}

func (e EntityError) Error() string {
	return e.Phrase
}

type QuestionError struct {
	Reason string
}

func (error QuestionError) Error() string {
	return error.Reason
}
