package utils

import (
	"database/sql"
	"net/http"
	"text/template"
	"time"
)

var (
	Tp *template.Template
	Db *sql.DB
)

type ErrorPage struct {
	Code         int
	ErrorMessage string
}
type Users struct {
	Username string
	Email    string
	Password string
}
type Categories struct {
	Name   string
	Id     int
	PostID int
}
type Posts struct {
	Id            int
	Username      string
	Title         string
	Description   string
	Time          time.Time
	TimeFormatted string
	TotalLikes    int
	TotalDislikes int
	Comments      []Comments
	Categories    []Categories
}
type Catgs struct {
	Catgs []string
}
type Comments struct {
	PostID   int
	Id       int
	Username string
	Comment  string
	Time     time.Time
	TimeFormattedComment string
	TotalLikes    int
	TotalDislikes int
}

var (
	ErrorBadReq = ErrorPage{
		Code:         http.StatusBadRequest,
		ErrorMessage: "Oops! It looks like there was an issue with your request. Please check your input and try again.",
	}

	ErrorNotFound = ErrorPage{
		Code:         http.StatusNotFound,
		ErrorMessage: "Uh-oh! The page you're looking for doesn't exist. It might have been moved or deleted.",
	}

	ErrorMethodnotAll = ErrorPage{
		Code:         http.StatusMethodNotAllowed,
		ErrorMessage: "The request method is not supported for this resource. Please check and try again with a valid method.",
	}

	ErrorInternalServerErr = ErrorPage{
		Code:         http.StatusInternalServerError,
		ErrorMessage: "Something went wrong on our end. We're working on fixing it—please try again later!",
	}
)
