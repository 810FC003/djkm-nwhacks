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
	caps.AddChrome(chrome.Capabilities{
		Path: chromeBinary,
		Prefs: map[string]interface{}{
			"profile.default_content_settings.popups": 0,
			"download.default_directory":              tempDir,
			"safebrowsing.enabled":                    true,
			"download.prompt_for_download":            false,
		},
		Args: []string{
			"safebrowsing-disable-download-protection",
			"window-size=1280,720",
		},
	})

	defer os.RemoveAll(tempDir)
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", driverPort))
	if err != nil