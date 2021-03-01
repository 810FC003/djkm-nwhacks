package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/tebeka/selenium/chrome"

	"github.com/sirupsen/logrus"
	"github.com/tebeka/selenium"
)

func main() {
	logrus.Infoln("Starting Chromedriver")
	s, err := selenium.NewChromeDriverService(argDriver, driverPort)
	if err != nil {
		logrus.Panicln(err)
	}
	defer s.Stop()

	logrus.Infof("Temporary download directory %s", tempDir)

	logrus.Infoln("Starting Chrome")
	caps := selenium.Capabilities{"browserName": "chrome"}
	caps.AddChrome(chrome.Capabilit