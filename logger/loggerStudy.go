package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var ll = *logrus.New()

func init() {
	ll.SetOutput(os.Stdout)
	ll.SetFormatter(&logrus.TextFormatter{})
}

func Infof(format string, args ...interface{}) {
	ll.Infof("one: %12s, two: %v", "one", time.Now())
}
