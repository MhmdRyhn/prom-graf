package main

import (
	"github.com/mhmdryhn/prom-graf/app"
	"github.com/sirupsen/logrus"
)

func main() {
	port := 2020
	logrus.Infof("Serving app on port: %d", port)
	err := app.Start(port)
	if err != nil {
		logrus.Errorf("Error while http.ListenAndServe: %s", err.Error())
	}
}
