package logging

import (
	"fmt"
)

type Logger struct {
	// path to logger, i.e. for logger "http.request" it would be ["http", "request"]
	path Name
	// children loggers
	children map[string]*Logger
	// files attached to logger
	file fileArray
	// logger level
	level Level
}

func (self *Logger) Name() Name {
	return self.path
}

func (self *Logger) Level() Level {
	global.mutex.Lock()
	defer global.mutex.Unlock()

	return self.level
}

func (self *Logger) GetLogger(name string) *Logger {
	global.mutex.Lock()
	defer global.mutex.Unlock()

	if result, ok := self.children[name]; ok {
		return result
	}

	result := &Logger{
		self.Name().GetChild(name),
		make(map[string]*Logger),
		self.file,
		INFO,
	}
	self.children[name] = result

	return result
}

func (self *Logger) String() string {
	global.mutex.Lock()
	defer global.mutex.Unlock()

	var children []string
	for name, _ := range self.children {
		children = append(children, name)
	}
	return fmt.Sprintf("Logger{path: '%v' children: %v file: %v level %v}",
		self.path, children, self.file, self.level)
}
