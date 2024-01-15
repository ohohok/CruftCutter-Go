package logger

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

type BaseTime struct {
	NowTime time.Time
}

func (b BaseTime) Now() time.Time {
	// init
	now := time.Now()
	nowUnix := now.Unix()
	baseTime := nowUnix - int64(now.Hour()) - int64(now.Second())
	b.NowTime = time.Unix(baseTime, 0)
	//
	return b.NowTime
}

func newLfsHook() logrus.Hook {
	writer, err := rotatelogs.New(
		LogPath+".%Y%m%d%H",
		// WithLinkName为最新的日志建立软连接，以方便随着找到当前日志文件
		rotatelogs.WithLinkName(LogPath),

		// WithRotationTime设置日志分割的时间，这里设置为12小时分割一次
		rotatelogs.WithRotationTime(24*time.Hour),

		rotatelogs.WithClock(BaseTime{}),

		// WithMaxAge和WithRotationCount二者只能设置一个，
		// WithMaxAge设置文件清理前的最长保存时间，
		// WithRotationCount设置文件清理前最多保存的个数。
		rotatelogs.WithMaxAge(time.Hour*24*10),
		//rotatelogs.WithRotationCount(maxRemainCnt),

		rotatelogs.WithRotationSize(100*1024*1024),
	)

	if err != nil {
		logrus.Errorf("config local file system for logger error: %v", err)
	}

	StdLogger.Out = writer

	return nil
}
