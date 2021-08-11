package main

import (
	"github.com/sirupsen/logrus"
	"twitter.go/log"
	"twitter.go/server"
	"twitter.go/src/migrate"
)

func main() {
	err := log.SetupLog()
	if err != nil {
		logrus.Fatal(err)
	}
	err = migrate.InitConfig()
	if err != nil {
		logrus.Fatal(err)
	}
	server.BuildServer().Run()
}
