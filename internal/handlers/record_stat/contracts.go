package recordstat

type Payload struct {
	UserId   int `json:"userId"`
	AuthorId int `json:"authorId"`
}

type Logger interface {
	Error(message string, err error)
}
