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

	logrus.Infof("Found %d results", resultCountInt)
	logrus.Infof("Staring export")

	resultsLeft := resultCountInt
	if argEnd == 0 {
		resultsLeft -= (argStart - 1)
	} else {
		resultsLeft = argEnd - argStart + 1
	}

	iterationN := argStartIter
	lastDownloaded := argStart - 1

	logrus.Infof("Will download %d from %d to %d \n", resultsLeft, argStart, argEnd)

	for resultsLeft > 0 {

		logrus.Infof("Iteration %d, need to download %d", iterationN, resultsLeft)

		toDownload := 500
		if resultsLeft < 500 {
			toDownload = resultsLeft
		}
		firstTD := lastDownloaded + 1
		lastTD := lastDownloaded + toDownload

		logrus.Infof("Downloading %d from %d to %d", toDownload, firstTD, lastTD)

		_, err = wd.FindElement(selenium.ByXPATH, `//*[@id="exportTypeName"]`)
		if err == nil {
			mustClickByXPATH(wd, `//*[@id="exportTypeName"]`)
		} else {
			mustClickByXPATH(wd, `//*[@id="page"]/div[1]/div[26]/div[2]/div/div/div/div[2]/div[3]/div[3]/div[2]/div[1]/button`)
		}

		time.Sleep(500 * time.Millisecond)

		_, err = wd.FindElement(selenium.ByXPATH, `/html/body/div[1]/div[26]/div[2]/div/div/div/div[2]/div[3]/div[3]/div[2]/div[1]/ul/li/span/ul/li[3]/a`)
		if err == nil {
			mustClickByXPATH(wd, `/html/body/div[1]/div[26]/div[2]/div/div/div/div[2]/div[3]/div[3]/div[2]/div[1]/ul/li/span/ul/li[3]/a`)
		}

		time.Sleep(1 * time.Second)

		mustClickByXPATH(wd, `//*[@id="numberOfRecordsRange"]`)
		time.Sleep(30 * time.Millisecond)

		setCntByXPATH(wd, `//*[@id="markFrom"]`, fmt.Sprintf("%d", firstTD))
		time.Sleep(30 * time.Millisecond)

		setCntByXPATH(wd, `//*[@id="markTo"]`, fmt.Sprintf("%d", lastTD))
		time.Sleep(30 * time.Millisecond)

		mustClickByXPATH(wd, `/html/body/div[11]/div[2]/form/div[2]/div[2]/div/span/span[1]/span/span[2]`)
		time.Sleep(30 * time.Millisecond)
		mustClickByXPATH(wd, `//*[@id="bib_fields:fullrec_fields_option"]`)
		time.Sleep(30 * time.Millisecond)
		mustClickByXPATH(wd, `/html/body/div[11]/div[2]/form/div[2]/div[2]/div/span/span[1]/span/span[2]`)
		time.Sleep(30 * time.Millisecond)

		mustClickByXPATH(wd, `/html/body/div[11]/div[2]/form/div[3]/div/div/div/span/span[1]/span/span[2]`)
		time.Sleep(30 * time.Millisecond)
		mustClickByXPATH(wd, `//*[@id="saveOptions"]/option[4]`)
		time.Sleep(30 * time.Millisecond)
		mustClickByXPATH(wd, `/html/body/div[11]/div[2]/form/div[3]/div/div/div/span/span[1]/span/span[2]`)
		time.Sleep(30 * time.Millisecond)

		mustClickByXPATH(wd, `//*[@id="exportButton"]`)

		fname, err := waitForDl(20 * time.Minute)
		if err != nil {
			logrus.Panic(err)
		}

		err = portableMoveFile(tempDir+"/"+fname, argDownloadDir+"/"+fmt.Sprintf("%s_%d", "WOS_EXP", iterationN)+"_"+fname)
		if err != nil {
			logrus.Panic(err)
		}

		time.Sleep(700 * time.Millisecond)
		mustClickByXPATH(wd, `//*[@id="page"]/div[11]/div[2]/form/div[2]/a`)

		logrus.Infof("File saved to %s\n", argDownloadDir+"/"+fmt.Sprintf("%s_%d", "WOS_EXP", iterationN)+"_"+fname)

		it