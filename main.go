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
	if err != nil {
		logrus.Panicln(err)
	}
	defer wd.Quit()

	logrus.Infoln("Chrome is running!")
	logrus.Infoln("Go to the results page, wait till it fully loads, then press Enter to export results")
	fmt.Scanln()

	resultCountE, err := wd.FindElement(selenium.ByXPATH, `//*[@id="hitCount.top"]`)
	if err != nil {
		logrus.Panicf("Cannot get result count, error %v", err)
	}

	resultCountText, err := resultCountE.Text()
	if err != nil {
		logrus.Panicf("Cannot get result count, error %v", err)
	}
	resultCountTextGood := strings.ReplaceAll(resultCountText, ",", "")

	var resultCountInt int
	_, err = fmt.Sscanf(resultCountTextGood, "%d", &resultCountInt)
	if err != nil {
		logrus.Panicf("Cannot get result count, error %v", err)
	}

	