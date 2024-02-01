
// logservice/logservice.go
package logservice

import (
// "fmt"
    "github.com/sirupsen/logrus"
)

func logWithFields(logLevel logrus.Level, message string, fields logrus.Fields) {
    logEntry := Logger.WithFields(fields)
    logEntry.Log(logLevel, message)
}

func Trace(message string, fields logrus.Fields) {
    logWithFields(logrus.TraceLevel, message, fields)
}

func InfoData(message string, fields logrus.Fields) {
    logWithFields(logrus.InfoLevel, message, fields)
}

func Warn(message string, fields logrus.Fields) {
    logWithFields(logrus.WarnLevel, message, fields)
}

func Error(message string, fields logrus.Fields) {
    logWithFields(logrus.ErrorLevel, message, fields)
}