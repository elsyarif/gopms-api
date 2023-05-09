package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var (
	log *logrus.Logger
)

func init() {
	timeFormat := new(logrus.TextFormatter)
	log = logrus.New()
	timeFormat.TimestampFormat = "2006-01-02 15:04:05"
	log.SetFormatter(timeFormat)
	timeFormat.FullTimestamp = true

	path := "log/"
	logFileName := "app.log"

	file, err := os.OpenFile(path+logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	log.Out = file

	fileInfo, err := os.Stat(path + logFileName)
	if err != nil {
		panic(err)
	}
	fileModTime := fileInfo.ModTime()
	now := time.Now()

	if fileModTime.Day() != now.Day() {
		file.Close()
		newLogFileName := "app_log_" + fileModTime.Format("02012006") + ".log"

		err := os.Rename(path+logFileName, path+newLogFileName)
		if err != nil {
			panic(err)
		}
		file, err = os.OpenFile(path+logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		log.Out = file
	}
}

func Info(message string, fields map[string]interface{}) {
	log.WithFields(fields).Info(message)
}

func Debug(message string, fields map[string]interface{}) {
	log.WithFields(fields).Debug(message)
}

func Warn(message string, fields map[string]interface{}) {
	log.WithFields(fields).Warn(message)
}

func Error(message string, fields map[string]interface{}) {
	log.WithFields(fields).Error(message)
}

func Fatal(message string, fields map[string]interface{}) {
	log.WithFields(fields).Fatal(message)
}

func Close() {

}
