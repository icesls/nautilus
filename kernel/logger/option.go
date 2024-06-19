// +----------------------------------------------------------------------
// | nautilus [ logger option]
// +----------------------------------------------------------------------
// | Copyright (c) 2013~2024 https://www.secdos.com All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: ︶ㄣ逍遥楓 <admin@secdos.com>
// +----------------------------------------------------------------------

package logger

func WithFileName(filename string) Option {
	return func(l *Options) {
		l.FileName = filename
	}
}

func WithMaxSize(maxSize int) Option {
	return func(l *Options) {
		l.MaxSize = maxSize
	}
}

func WithMaxAge(maxAge int) Option {
	return func(l *Options) {
		l.MaxAge = maxAge
	}
}

func WithMaxBackup(maxBackup int) Option {
	return func(l *Options) {
		l.MaxBackup = maxBackup
	}
}

func WithCompress(compress bool) Option {
	return func(l *Options) {
		l.Compress = compress
	}
}

func WithOptionsType(logType string) Option {
	return func(l *Options) {
		l.LogType = logType
	}
}
