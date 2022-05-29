package log

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
	"time"

	"io"
	"os"
)

func init() {

	//设置输出样式
	log.SetFormatter(&nested.Formatter{
		NoColors:        true,
		ShowFullLevel:   true,
		HideKeys:        true,
		TimestampFormat: time.RFC3339,
	})
	log.SetOutput(colorable.NewColorableStdout())
	//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	file, err := os.OpenFile("Collette.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	writers := []io.Writer{
		file,
		os.Stdout}
	//同时写文件和屏幕
	fileAndStdoutWriter := io.MultiWriter(writers...)
	if err == nil {
		log.SetOutput(fileAndStdoutWriter)
	} else {
		log.Info("failed to log to file.")
	}
	//设置最低loglevel
	log.SetLevel(log.InfoLevel)
}
