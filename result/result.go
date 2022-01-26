package result

type ResultTag int

const (
	Ok ResultTag = iota
	Err
)

type Result struct {
	Tag   ResultTag
	Value interface{}
}

func NewResult(t ResultTag, v interface{}) *Result {
	return &Result{
		Tag:   t,
		Value: v,
	}
}

func NewOk(v interface{}) *Result {
	return NewResult(Ok, v)
}

func NewErr(v interface{}) *Result {
	return NewResult(Err, v)
}
