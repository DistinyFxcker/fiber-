package zapLog

import (
	"fmt"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

var (
	Zaplog *zap.SugaredLogger
)

type ZapLog struct {
	Path string
	Day  time.Duration
	Hour time.Duration
}

func (this *ZapLog) InitLogger() error {
	if !Exists(this.Path) {
		file, err := os.Create(this.Path)
		defer file.Close()
		if err != nil {
			fmt.Println("mkdir logPath err!")
			return err
		}
	}
	encoder := initEncoder()

	// 想要将日常文件区分开来，可以实现多个日志等级接口
	/*infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel
	})*/
	debugLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})

	// 获取 info、warn日志文件的io.Writer
	//infoIoWriter := getWriter("D:/bian/logs/dspcollect.log")
	warnIoWriter := getWriter(this.Path, this.Day, this.Hour)

	// 创建Logger
	core := zapcore.NewTee(
		//zapcore.NewCore(encoder, zapcore.AddSync(infoIoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(warnIoWriter), debugLevel),
	)
	logger := zap.New(core, zap.AddCaller()) // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数
	Zaplog = logger.Sugar()
	return nil
}

//初始化Encoder
func initEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		TimeKey:     "time",
		CallerKey:   "file",
		EncodeLevel: zapcore.CapitalLevelEncoder, //基本zapcore.LowercaseLevelEncoder。将日志级别字符串转化为小写
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeCaller: zapcore.ShortCallerEncoder, //一般zapcore.ShortCallerEncoder，以包/文件:行号 格式化调用堆栈
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) { //一般zapcore.SecondsDurationEncoder,执行消耗的时间转化成浮点型的秒
			enc.AppendInt64(int64(d) / 1000000)
		},
	})
}

//日志文件切割
func getWriter(filename string, day, hour time.Duration) io.Writer {
	var (
		Day, Hour time.Duration
	)
	Day = time.Hour * 24 * 90
	if day != 0 {
		Day = time.Hour * 24 * day
	}
	Hour = time.Hour * 24
	if day != 0 {
		Hour = time.Hour * hour
	}
	// 保存30天内的日志，每24小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		filename+".%Y%m%d",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(Day),
		rotatelogs.WithRotationTime(Hour),
	)

	if err != nil {
		panic(err)
	}
	return hook
}

//查看文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}