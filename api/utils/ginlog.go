package utils

//gin框架日志重写
import (
	"os"
	"path"
	"time"
)

//gin log
type Log struct {
	files map[string]*os.File
	path  string
}

func Loger() *Log {
	logpath := Cfg.MustValue("log", "path")
	if logpath == "" {
		logpath = "./log/"
	}
	l := new(Log)
	l.path = logpath
	l.files = make(map[string]*os.File)
	filename := time.Now().Format("2006-01-02")
	file, err := os.Create(l.path + filename + ".log")
	if err != nil {
		os.MkdirAll(path.Dir(logpath), os.ModePerm)
		file, err = os.Create(l.path + filename + ".log")
		if err != nil {
			panic(err.Error())
		}
	}
	l.files[filename] = file
	return l
}
func (l *Log) Write(p []byte) (n int, err error) {
	filename := time.Now().Format("2006-01-02")
	file, has := l.files[filename]
	if !has {
		for k, f := range l.files {
			f.Close()
			delete(l.files, k)
		}
		var err error
		file, err = os.Create(l.path + filename + ".log")
		if err != nil {
			return 0, err
			//panic(err.Error())
		}
		l.files[filename] = file
	}
	return file.Write(p)
}
