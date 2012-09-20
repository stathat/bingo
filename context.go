package bingo

import (
	"net/http"
)

type Context interface {
	Request() *http.Request
	Writer() http.ResponseWriter
	Session() *Session

	Before() (bool, error)
	After()

	ResultData(title string) TemplateData
}
