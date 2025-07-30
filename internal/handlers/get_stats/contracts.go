package getstats

/*
	{
	    "authorId": 2
	}
*/
type Payload struct {
	AuthorId int `json:"authorId"`
}

type Response struct {
	Count int `json:"count"`
}

type Logger interface {
	Error(message string, err error)
}
