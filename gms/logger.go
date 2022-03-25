package gms

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
)

type GLogger struct {
	logrus.Logger
}

func (me *GLogger) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *me
	//newlogger.Level = level
	return &newlogger
}

// Info print info
func (me *GLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	// if l.LogLevel >= Info {
	// 	l.Printf(l.infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	// }
}

// Warn print warn messages
func (me *GLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	// if l.LogLevel >= Warn {
	// 	l.Printf(l.warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	// }
}

// Error print error messages
func (me *GLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	// if l.LogLevel >= Error {
	// 	l.Printf(l.errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	// }
}

// Trace print sql message
func (me *GLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	// if l.LogLevel <= Silent {
	// 	return
	// }

	// elapsed := time.Since(begin)
	// switch {
	// case err != nil && l.LogLevel >= Error && (!errors.Is(err, ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
	// 	sql, rows := fc()
	// 	if rows == -1 {
	// 		l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
	// 	} else {
	// 		l.Printf(l.traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
	// 	}
	// case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= Warn:
	// 	sql, rows := fc()
	// 	slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
	// 	if rows == -1 {
	// 		l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
	// 	} else {
	// 		l.Printf(l.traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
	// 	}
	// case l.LogLevel == Info:
	// 	sql, rows := fc()
	// 	if rows == -1 {
	// 		l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
	// 	} else {
	// 		l.Printf(l.traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
	// 	}
	// }
}
