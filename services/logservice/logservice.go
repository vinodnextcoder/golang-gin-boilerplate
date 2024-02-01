
// logservice/logservice.go
package logservice

import (
// "fmt"
    "github.com/sirupsen/logrus"
)

func logWithFields(logLevel logrus.Level, message string, fields ...logrus.Fields) {
    logEntry := Logger.WithFields(mergeFields(fields...))
    logEntry.Log(logLevel, message)
}

func Trace(message string, fields ...logrus.Fields) {
    logWithFields(logrus.TraceLevel, message, fields...)
}

func Info(message string, fields ...logrus.Fields) {
    logWithFields(logrus.InfoLevel, message, fields...)
}

func Warn(message string, fields ...logrus.Fields) {
    logWithFields(logrus.WarnLevel, message, fields...)
}

func Error(message string, fields ...logrus.Fields) {
    logWithFields(logrus.ErrorLevel, message, fields...)
}

// mergeFields merges variadic logrus.Fields into a single logrus.Fields
func mergeFields(fields ...logrus.Fields) logrus.Fields {
    result := make(logrus.Fields)
    for _, f := range fields {
        for k, v := range f {
            result[k] = v
        }
    }
    return result
}