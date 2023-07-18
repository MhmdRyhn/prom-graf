package app

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

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
