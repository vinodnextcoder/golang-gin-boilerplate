// logservice/logger.go
package logservice

import (
    "fmt"
    "io"
    "os"
    "github.com/sirupsen/logrus"
    "github.com/t-tomalak/logrus-easy-formatter"
)

var Logger *logrus.Logger

func InitLogger() {
    f, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        fmt.Println("Failed to create logfile" + "log.txt")
        panic(err)
    }

    Logger = &logrus.Logger{
        // Log into f file handler and on os.Stdout
        Out:   io.MultiWriter(f, os.Stdout),
        Level: logrus.DebugLevel,
        Formatter: &easy.Formatter{
            TimestampFormat: "2006-01-02 15:04:05",
            LogFormat:       "[%lvl%]: %time% - %msg%\n",
        },
    }
}
