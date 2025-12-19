package logging

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() Logger {
	return Logger{e}
}

type MyHook struct {
	Writers []io.Writer
	LogLevels []logrus.Level
}


func (h *MyHook) Fire(entry *logrus.Entry) error {
	line , err := entry.String()
	if err != nil {
		log.Fatal("error")
	}

	for _ , w := range h.Writers {
		w.Write([]byte(line))
	}

	return nil
}

func (h *MyHook) Levels() []logrus.Level {
	return h.LogLevels
}

func init() {
	logger := logrus.New()

	logger.SetReportCaller(true)

	logger.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			filename := filepath.Base(f.File)
			return f.Function, fmt.Sprintf("%s:%d", filename, f.Line)
		},

		FullTimestamp: true,
		DisableColors: false,
	}

	err := os.MkdirAll("logs" , 0755)
	if err != nil {
		fmt.Println(err)
	}

	file , err := os.OpenFile("logs/all.log" , os.O_CREATE|os.O_WRONLY|os.O_APPEND , 0644)
	if err != nil {
		fmt.Println(err)
	}
	

	logger.AddHook(&MyHook{
		Writers: []io.Writer{file , os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	logger.SetOutput(io.Discard)
	logger.SetLevel(logrus.TraceLevel)
	
	
	
	e = logrus.NewEntry(logger)
}
