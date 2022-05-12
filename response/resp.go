package response

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

//Response struct
type Response struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
	Body    interface{} `json:"body"`
}

// ErrorFromError prepares and returns data to user
func ErrorFromError(err error) *Response {
	return &Response{Status: -1, Message: "error", Body: ListOfErrors(err)}
}

// ErrorFromString returns error representation
func ErrorFromString(message string) *Response {
	return &Response{Status: -1, Message: "error", Body: message}
}

// ErrorFromStringAndLog returns error representation
func ErrorFrom(message, str string, err error) *Response {
	log.Error(str, err)
	return &Response{Status: -1, Message: "error", Body: message}
}

// Forbidden response
func Forbidden() *Response {
	return &Response{Status: -1, Message: "forbidden", Body: "insufficient privelegies"}
}

//// NotFound response
func NotFound() *Response {
	return &Response{Status: -1,
		Message: http.StatusText(http.StatusNotFound),
		Body:    "Not Found Route 4 0 4 ",
	}
}

// MenyRequest response
func TooMenyRequest() *Response {
	return &Response{Status: -1, Message: "Meny Request", Body: "Too many requests"}
}

// CorrectWithData prepares and returns data to user
func CorrectWithData(data interface{}) *Response {
	return &Response{Status: 0, Message: "success", Body: data}
}

// Correct prepares and returns data to user
func Correct() *Response {
	return &Response{Status: 0, Message: "success"}
}

//ServerError ....
func ServerError() *Response {
	return &Response{Status: -1, Message: http.StatusText(500)}
}

// Created prepares and returns data to user
func Createdd() *Response {
	return &Response{Status: 0, Message: "Created"}
}

// ListOfErrors returns formmated error
func ListOfErrors(e error) []map[string]string {
	InvalidFields := make([]map[string]string, 0)

	switch e.(type) {
	case validator.ValidationErrors:
		ve := e.(validator.ValidationErrors)

		for _, e := range ve {
			errors := map[string]string{}
			errors[e.Tag()] = e.Field()
			InvalidFields = append(InvalidFields, errors)
		}
	default:
		InvalidFields = append(InvalidFields, map[string]string{"error": e.Error()})
	}

	return InvalidFields
}
