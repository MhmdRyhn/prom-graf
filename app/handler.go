package app

import (
	"errors"
	"math/rand"
	"net/http"

	"github.com/sirupsen/logrus"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	number := rand.Intn(10) + 1
	var statusCode int
	if number <= 5 {
		statusCode = http.StatusOK
		data := map[string]interface{}{
			"data":       "Hello World :)",
			"statusCode": statusCode,
		}
		serveSuccess(w, data, statusCode)
	} else {
		statusCode = http.StatusInternalServerError
		err := errors.New("Hello Error :(")
		serveError(w, err, statusCode)
	}
	logrus.Info("statusCode:", statusCode)
	return
}
