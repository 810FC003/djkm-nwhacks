
package main

import (
	"io/ioutil"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/tebeka/selenium"
)

func waitForDl(timeout time.Duration) (string, error) {
	startTime := time.Now()
	for {
		fileList, err := ioutil.ReadDir(tempDir)
		if err != nil {
			return "", err
		}

		for _, file := range fileList {
			fname := file.Name()
			if strings.HasSuffix(fname, "txt") {
				time.Sleep(1500 * time.Millisecond)
				return fname, nil
			}
			if time.Since(startTime) > timeout {
				time.Sleep(1500 * time.Millisecond)
				return "", errors.Errorf("Timeout exceeded")
			}
		}
		time.Sleep(1000 * time.Millisecond)
	}
}

func waitForResults(wd selenium.WebDriver) {
	wd.WaitWithTimeout(func(iwd selenium.WebDriver) (bool, error) {
		modal, err := iwd.FindElement(selenium.ByXPATH, "//*[@id=\"w_loader\"]")
		if err != nil {
			return true, err
		}
		disp, err := modal.IsDisplayed()
		return !disp, err
	}, 300*time.Second)
}

func setCntByXPATH(wd selenium.WebDriver, elemXPATH, content string) error {
	requestBox, err := wd.FindElement(selenium.ByXPATH, elemXPATH)
	if err != nil {
		return err
	}

	err = requestBox.Clear()
	if err != nil {
		return err