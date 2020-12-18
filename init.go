
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