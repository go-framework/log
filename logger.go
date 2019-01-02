package log

type Logger interface {
	Print(v ...interface{})
	Println(v ...interface{})
	Printf(format string, v ...interface{})
	Fatal(v ...interface{})
	Fatalln(v ...interface{})
	Fatalf(format string, v ...interface{})
	Panic(v ...interface{})
	Panicln(v ...interface{})
	Panicf(format string, v ...interface{})
}

type AdvancedLogger interface {
	Logger

	Debug(v ...interface{})
	Debugln(v ...interface{})
	Debugf(format string, v ...interface{})
	Info(v ...interface{})
	Infoln(v ...interface{})
	Infof(format string, v ...interface{})
	Warn(v ...interface{})
	Warnln(v ...interface{})
	Warnf(format string, v ...interface{})
	Error(v ...interface{})
	Errorn(v ...interface{})
	Errorf(format string, v ...interface{})
}
