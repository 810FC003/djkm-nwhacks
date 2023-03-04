
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