package stack

type StackError struct {
	content string
}

func (se StackError) Error() string  {
	return se.content
}
