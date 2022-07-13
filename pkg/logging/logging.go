package logging

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func Init() *Logger {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.JSONFormatter{
		// CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
		// 	filename := path.Base(f.File)
		// 	return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		// },
		TimestampFormat: "2006-01-02 15:04:05",
	}

	// err := os.MkdirAll("logs", 0755)
	// if err != nil {
	// 	panic(err)
	// }

	// file, err := os.OpenFile("logs/log.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	// if err != nil {
	// 	panic(err)

	// }

	// wrt := io.MultiWriter(os.Stdout, file)
	// l.SetOutput(wrt)
	l.SetLevel(logrus.TraceLevel)
	// defer file.Close()
	return &Logger{l}
}
