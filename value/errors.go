package value

type EntityError struct {
	Code   int32
	Phrase string
}

func (e EntityError) Error() string {
	return e.Phrase
}
