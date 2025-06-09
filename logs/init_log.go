package logs

import (
	"fmt"
	formatter "github.com/buhuang28/logrus-formatter"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
	"sync"
	"time"
)

var (
	o = new(sync.Once)
)

func LogInit() {
	o.Do(func() {
		fmt.Println("日志初始化")
		// 输出到命令行
		f := &formatter.Formatter{
			Caller: true,
			CallerFormat: func(f *runtime.Frame) string {
				_, l := f.Func.FileLine(f.PC)
				return fmt.Sprintf("%s:%d ", f.File, l)
			},
			Level: true,
		}
		log.SetFormatter(f)
		log.SetOutput(os.Stdout)
		log.SetReportCaller(true)

		// 输出到文件
		rotateLogs, err := rotatelogs.New(path.Join("logs", "%Y-%m-%d.log"),
			//rotatelogs.WithLinkName(path.Join("logs", "latest.log")), // 最新日志软链接
			rotatelogs.WithRotationTime(time.Hour*24), // 每天一个新文件
			rotatelogs.WithMaxAge(time.Hour*24*3),     // 日志保留3天
		)
		if err != nil {
			log.Info(err)
			return
		}
		log.AddHook(lfshook.NewHook(
			lfshook.WriterMap{
				log.InfoLevel:  rotateLogs,
				log.WarnLevel:  rotateLogs,
				log.ErrorLevel: rotateLogs,
				log.FatalLevel: rotateLogs,
				log.PanicLevel: rotateLogs,
			},
			f,
		))
	})
}
