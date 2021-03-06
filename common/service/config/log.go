// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package config

import (
	"github.com/Sirupsen/logrus"
	"github.com/uber-common/bark"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const fileMode = os.FileMode(0644)

// NewBarkLogger builds and returns a new bark
// logger for this logging configuration
func (cfg *Logger) NewBarkLogger() bark.Logger {

	logger := logrus.New()
	logger.Out = ioutil.Discard
	logger.Level = parseLogrusLevel(cfg.Level)
	logger.Formatter = getFormatter()

	if cfg.Stdout {
		logger.Out = os.Stdout
	}

	if len(cfg.OutputFile) > 0 {
		outFile := createLogFile(cfg.OutputFile)
		logger.Out = outFile
		if cfg.Stdout {
			logger.Out = io.MultiWriter(os.Stdout, outFile)
		}
	}

	return bark.NewLoggerFromLogrus(logger)
}

func getFormatter() logrus.Formatter {
	formatter := &logrus.TextFormatter{}
	formatter.FullTimestamp = true
	return formatter
}

func createLogFile(path string) *os.File {
	dir := filepath.Dir(path)
	if len(dir) > 0 && dir != "." {
		if err := os.MkdirAll(dir, fileMode); err != nil {
			log.Fatalf("error creating log directory %v, err=%v", dir, err)
		}
	}
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, fileMode)
	if err != nil {
		log.Fatalf("error creating log file %v, err=%v", path, err)
	}
	return file
}

// parseLogrusLevel converts the string log
// level into a logrus level
func parseLogrusLevel(level string) logrus.Level {
	switch strings.ToLower(level) {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	default:
		return logrus.InfoLevel
	}
}
