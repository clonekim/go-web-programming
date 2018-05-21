package netserver

import (
	echoLog "github.com/labstack/gommon/log"
	"github.com/neko-neko/echo-logrus/log"
	"github.com/sirupsen/logrus"
	"os"
)

var Log *logrus.Logger = log.Logger().Logger

func SetupLogger() {
	log.Logger().SetOutput(os.Stdout)
	if Conf.Debug {
		log.Logger().SetLevel(echoLog.DEBUG)
	} else {
		log.Logger().SetLevel(echoLog.INFO)
	}
}
