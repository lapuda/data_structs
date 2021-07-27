package queue

type QueueError struct {
	content string
}

func (qe QueueError) Error() string  {
	return qe.content
}
