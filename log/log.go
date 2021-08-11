package log

import (
	"github.com/ilya1st/rotatewriter"
	"github.com/sirupsen/logrus"
)

func SetupLog() error {
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "02-01-2006 15:04:05",
	})
	writer, err := rotatewriter.NewRotateWriter("./log/error.log", 1)
	if err != nil {
		return err
	}

	logrus.SetOutput(writer)
	return nil
}
