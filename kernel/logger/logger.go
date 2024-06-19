// +----------------------------------------------------------------------
// | nautilus [ logger ]
// +----------------------------------------------------------------------
// | Copyright (c) 2013~2024 https://www.secdos.com All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: ︶ㄣ逍遥楓 <admin@secdos.com>
// +----------------------------------------------------------------------

package logger

import (
	"github.com/icesls/nautilus/kernel/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"
)

var (
	Logger *zap.Logger
)

// Options 参数选项
type Options struct {
	FileName  string `json:"file_name"`
	MaxSize   int    `json:"max_size"`
	MaxAge    int    `json:"max_age"`
	MaxBackup int    `json:"max_backup"`
	Compress  bool   `json:"compress"`
	LogType   string `json:"log_type"`
	Debug     bool   `json:"debug"`
	Level     string `json:"level"`
}

type Option func(*Options)

// NewLogger 构造方法
func NewLogger(options ...Option) *Options {
	c := &Options{
		FileName:  "nautilus", // 日志文件
		MaxSize:   64,         // 每个日志文件保存的最大尺寸 单位：mb
		MaxAge:    30,         // 最多保存多少天
		MaxBackup: 5,          // 最多保存日志文件数
		Compress:  false,      // 是否压缩
		LogType:   "single",   // daily按照日期每日一个
		Debug:     false,      // 是否为调试模式，调试模式日志会终端输出
		Level:     "debug",    // 日志级别
	}
	for _, option := range options {
		option(c)
	}

	return c
}

// Init 初始化
func (l *Options) Init() {

	// 获取日志写入介质
	logWriterStore := l.logWriteStore()

	// 设置日志等级
	logLevel := new(zapcore.Level)
	utils.Throw(logLevel.UnmarshalText([]byte(l.Level)))
	core := zapcore.NewCore(l.encoder(), logWriterStore, logLevel) // 初始化

	Logger = zap.New(
		core,
		zap.AddCaller(),                   // 调用文件和行号，内部使用 runtime.Caller
		zap.AddCallerSkip(1),              // 封装了一层，调用文件去除一层(runtime.Caller(1))
		zap.AddStacktrace(zap.ErrorLevel), // Error 时才会显示 stacktrace
	)

	zap.ReplaceGlobals(Logger)
}

// 设置日志存储格式
func (l *Options) encoder() zapcore.Encoder {
	config := zapcore.EncoderConfig{
		MessageKey:          "message",
		LevelKey:            "level",
		TimeKey:             "time",
		NameKey:             "logger",
		CallerKey:           "caller",
		FunctionKey:         zapcore.OmitKey,
		StacktraceKey:       "stacktrace",
		SkipLineEnding:      false,
		LineEnding:          zapcore.DefaultLineEnding,      // 每行日志的结尾添加 "\n"
		EncodeLevel:         zapcore.CapitalLevelEncoder,    // 日志级别名称大写，如 ERROR、INFO
		EncodeTime:          l.customTimeEncoder,            // 时间格式，我们自定义为 2006-01-02 15:04:05
		EncodeDuration:      zapcore.SecondsDurationEncoder, // 执行时间，以秒为单位
		EncodeCaller:        zapcore.ShortCallerEncoder,     // Caller 短格式
		NewReflectedEncoder: nil,
	}

	// debug模式
	if l.Debug {
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder // 终端输出的关键词高亮
		return zapcore.NewConsoleEncoder(config)
	}

	// 线上环境使用 JSON 编码器
	return zapcore.NewConsoleEncoder(config)
}

// 自定义友好的时间格式
func (l *Options) customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// 存储介质
func (l *Options) logWriteStore() zapcore.WriteSyncer {
	var (
		logName  string
		fileName string
	)

	// 按日期记录日志
	if l.LogType == "daily" {
		logName = time.Now().Format("2006-01-02.log")
		fileName = strings.ReplaceAll(l.FileName, "nautilus.log", logName)
	}

	// 滚动日志
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    l.MaxSize,
		MaxAge:     l.MaxAge,
		MaxBackups: l.MaxBackup,
		Compress:   l.Compress,
	}

	if l.Debug {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}

	return zapcore.AddSync(lumberJackLogger)
}

func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

// LogIf 当 err != nil 时记录 error 等级的日志
func LogIf(err error) {
	if err != nil {
		Logger.Error("Error:", zap.Error(err))
	}
}

// LogWarnIf 当 err != nil 时记录 warning 等级的日志
func LogWarnIf(err error) {
	if err != nil {
		Logger.Warn("Error Occurred:", zap.Error(err))
	}
}

// LogInfoIf 当 err != nil 时记录 info 等级的日志
func LogInfoIf(err error) {
	if err != nil {
		Logger.Info("Error Occurred:", zap.Error(err))
	}
}
