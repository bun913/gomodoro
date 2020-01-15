package logger

import (
	"bytes"
	"testing"
)

func TestErrorf(t *testing.T) {
	cases := []struct {
		name     string
		logLevel Level
		format   string
		args     []interface{}
		want     string
	}{
		{
			name:     "success",
			logLevel: ErrorLevel,
			format:   "test %s",
			args:     []interface{}{"test"},
			want:     "[ERROR]test test\n",
		},
		{
			name:     "logLevel is InfoLevel",
			logLevel: InfoLevel,
			format:   "test",
			want:     "[ERROR]test\n",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			buf := new(bytes.Buffer)
			SetOutput(buf)
			SetLogLevel(c.logLevel)
			Errorf(c.format, c.args...)

			s := buf.String()

			if s != c.want {
				tt.Errorf("want: %s, got: %s", c.want, s)
			}
		})
	}
}

func TestWarnf(t *testing.T) {
	cases := []struct {
		name     string
		logLevel Level
		format   string
		args     []interface{}
		want     string
	}{
		{
			name:     "success",
			logLevel: WarnLevel,
			format:   "test %s",
			args:     []interface{}{"test"},
			want:     "[WARN]test test\n",
		},
		{
			name:     "logLevel is ErrorLevel",
			logLevel: ErrorLevel,
			format:   "test",
			want:     "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			buf := new(bytes.Buffer)
			SetOutput(buf)
			SetLogLevel(c.logLevel)
			Warnf(c.format, c.args...)

			s := buf.String()

			if s != c.want {
				tt.Errorf("want: %s, got: %s", c.want, s)
			}
		})
	}
}

func TestInfof(t *testing.T) {
	cases := []struct {
		name     string
		logLevel Level
		format   string
		args     []interface{}
		want     string
	}{
		{
			name:     "success",
			logLevel: InfoLevel,
			format:   "test %s",
			args:     []interface{}{"test"},
			want:     "[INFO]test test\n",
		},
		{
			name:     "logLevel is WarnLevel",
			logLevel: WarnLevel,
			format:   "test",
			want:     "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			buf := new(bytes.Buffer)
			SetOutput(buf)
			SetLogLevel(c.logLevel)
			Infof(c.format, c.args...)

			s := buf.String()

			if s != c.want {
				tt.Errorf("want: %s, got: %s", c.want, s)
			}
		})
	}
}

func TestDebugf(t *testing.T) {
	cases := []struct {
		name     string
		logLevel Level
		format   string
		args     []interface{}
		want     string
	}{
		{
			name:     "success",
			logLevel: DebugLevel,
			format:   "test %s",
			args:     []interface{}{"test"},
			want:     "[DEBUG]test test\n",
		},
		{
			name:     "logLevel is InfoLevel",
			logLevel: InfoLevel,
			format:   "test",
			want:     "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			buf := new(bytes.Buffer)
			SetOutput(buf)
			SetLogLevel(c.logLevel)
			Debugf(c.format, c.args...)

			s := buf.String()

			if s != c.want {
				tt.Errorf("want: %s, got: %s", c.want, s)
			}
		})
	}
}

// Logger instance test
func TestLoggerErrorf(t *testing.T) {
	cases := []struct {
		name     string
		logLevel Level
		format   string
		args     []interface{}
		want     string
	}{
		{
			name:     "success",
			logLevel: ErrorLevel,
			format:   "test %s",
			args:     []interface{}{"test"},
			want:     "[ERROR]test test\n",
		},
		{
			name:     "logLevel is InfoLevel",
			logLevel: InfoLevel,
			format:   "test",
			want:     "[ERROR]test\n",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			buf := new(bytes.Buffer)

			logger := NewLogger(c.logLevel)
			logger.out.SetOutput(buf)

			logger.Errorf(c.format, c.args...)

			s := buf.String()

			if s != c.want {
				tt.Errorf("want: %s, got: %s", c.want, s)
			}
		})
	}
}

func TestLoggerWarnf(t *testing.T) {
	cases := []struct {
		name     string
		logLevel Level
		format   string
		args     []interface{}
		want     string
	}{
		{
			name:     "success",
			logLevel: WarnLevel,
			format:   "test %s",
			args:     []interface{}{"test"},
			want:     "[WARN]test test\n",
		},
		{
			name:     "logLevel is ErrorLevel",
			logLevel: ErrorLevel,
			format:   "test",
			want:     "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			buf := new(bytes.Buffer)
			logger := NewLogger(c.logLevel)
			logger.out.SetOutput(buf)

			logger.Warnf(c.format, c.args...)

			s := buf.String()

			if s != c.want {
				tt.Errorf("want: %s, got: %s", c.want, s)
			}
		})
	}
}

func TestLoggerInfof(t *testing.T) {
	cases := []struct {
		name     string
		logLevel Level
		format   string
		args     []interface{}
		want     string
	}{
		{
			name:     "success",
			logLevel: InfoLevel,
			format:   "test %s",
			args:     []interface{}{"test"},
			want:     "[INFO]test test\n",
		},
		{
			name:     "logLevel is WarnLevel",
			logLevel: WarnLevel,
			format:   "test",
			want:     "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			buf := new(bytes.Buffer)
			logger := NewLogger(c.logLevel)
			logger.out.SetOutput(buf)
			logger.Infof(c.format, c.args...)

			s := buf.String()

			if s != c.want {
				tt.Errorf("want: %s, got: %s", c.want, s)
			}
		})
	}
}

func TestLoggerDebugf(t *testing.T) {
	cases := []struct {
		name     string
		logLevel Level
		format   string
		args     []interface{}
		want     string
	}{
		{
			name:     "success",
			logLevel: DebugLevel,
			format:   "test %s",
			args:     []interface{}{"test"},
			want:     "[DEBUG]test test\n",
		},
		{
			name:     "logLevel is InfoLevel",
			logLevel: InfoLevel,
			format:   "test",
			want:     "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			buf := new(bytes.Buffer)
			logger := NewLogger(c.logLevel)
			logger.out.SetOutput(buf)
			logger.Debugf(c.format, c.args...)

			s := buf.String()

			if s != c.want {
				tt.Errorf("want: %s, got: %s", c.want, s)
			}
		})
	}
}