package mw

import (
	"fmt"
	"hackernews-service/helpers/response"
	"net/http"
)


type NotFoundError struct {
    Message string
}

type InternalServerError struct {
    Message string
}

type BadRequestError struct {
    Message string
}

func (e BadRequestError) Error() string {
    return e.Message
}

func (e NotFoundError) Error() string {
    return e.Message
}

func (e InternalServerError) Error() string {
    return e.Message
}

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				switch e := err.(type) {
				case *BadRequestError:
					response.Error(w, http.StatusBadRequest, e.Message)
				case *NotFoundError:
					response.Error(w, http.StatusNotFound, e.Message)
				case *InternalServerError:
					fmt.Println(e.Message)
					response.Error(w, http.StatusInternalServerError,"internal server error")
				default:
					response.Error(w, http.StatusInternalServerError,"internal server error")
				}
			}
		}()

		next.ServeHTTP(w, r)
	})
}