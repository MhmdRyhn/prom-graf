package app

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)
/*
(rate(monitoring_promgraf_request_count{job="golang_app"}[1m]) * 100) / rate(monitoring_promgraf_request_count{status_code="200"}[1m])
*/

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	sleepFor := rand.Intn(300) + 550
	time.Sleep(time.Duration(sleepFor) * time.Millisecond)
	number := rand.Intn(10) + 1
	var statusCode int
	if number <= 5 {
		statusCode = http.StatusOK
		data := map[string]interface{}{
			"data":       fmt.Sprintf("After waiting %dms, Hello World :)", sleepFor),
			"statusCode": statusCode,
		}
		serveSuccess(w, data, statusCode)
	} else {
		statusCode = http.StatusInternalServerError
		err := errors.New(fmt.Sprintf("After waiting %dms, Hello Error :)", sleepFor))
		serveError(w, err, statusCode)
	}
	logrus.Info("statusCode:", statusCode)
	return
}
