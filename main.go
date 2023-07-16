package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func response(w http.ResponseWriter, statusCode int, read io.Reader) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	_, err := io.Copy(w, read)
	if err != nil {
		logrus.Error("io.Copy error: ", err.Error())
	}
}

func makeBuffer(data interface{}) (io.Reader, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	err := encoder.Encode(data)
	return buffer, err
}

func serveSuccess(w http.ResponseWriter, data interface{}, statusCode int) {
	buffer, err := makeBuffer(data)
	if err != nil {
		logrus.Error("http.ServeData Encode: ", err)
		return
	}
	response(w, statusCode, buffer)
}

func serveError(w http.ResponseWriter, err error, statusCode int) {
	data := map[string]interface{}{
		"error":      err.Error(),
		"statusCode": statusCode,
	}
	reader, _ := makeBuffer(data)
	response(w, http.StatusInternalServerError, reader)
}

func PromPullHandler(w http.ResponseWriter, r *http.Request) {
	number := rand.Intn(10) + 1
	var statusCode int
	if number <= 5 {
		statusCode = http.StatusOK
		data := map[string]interface{}{
			"data":       fmt.Sprintf("Desired number: %d (which is <= 5)", number),
			"statusCode": statusCode,
		}
		serveSuccess(w, data, statusCode)
	} else {
		statusCode = http.StatusInternalServerError
		err := errors.New(fmt.Sprintf("Unwanted number: %d  (which is > 5)", number))
		serveError(w, err, statusCode)
	}
	logrus.Info("statusCode:", statusCode)
	return
}

func main() {
	r := chi.NewRouter()
	r.Get("/greetings", PromPullHandler)
	err := http.ListenAndServe(":2021", r)
	if err != nil {
		logrus.Errorf("Error while http.ListenAndServe: %s", err.Error())
	}
}
