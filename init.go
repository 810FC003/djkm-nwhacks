
package main

import (
	"flag"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

const defaultChromedriver = "./chromedriver"
const defaultDownDir = "."
const driverPort = 54678
const defaultChromeBinary = `C:\Program Files (x86)\Google\Chrome\Application\chrome.exe`
const defaultStart = 1
const defaultEnd = 0
const defaultStartIter = 1

var argDriver string
var argDownloadDir string
var chromeBinary string
var argStart int
var argEnd int
var argStartIter int

var tempDir string

func init() {
	flag.StringVar(&argDriver, "driver", defaultChromedriver, "Path to Chromedriver binary")
	flag.StringVar(&argDownloadDir, "o", defaultDownDir, "Where to save the reports")
	flag.StringVar(&chromeBinary, "chrome", defaultChromeBinary, "Path to Chrome(ium) binary")