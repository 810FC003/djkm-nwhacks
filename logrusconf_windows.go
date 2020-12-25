
// +build windows

package main

import (
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
	"golang.org/x/sys/windows"
)

func logrusConf() {
	//if we're running on windows, enable VT sequence support
	//and then force logrus to emit those sequences
	if runtime.GOOS == "windows" {
		var originalMode uint32