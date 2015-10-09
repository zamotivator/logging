package logging

import (
	"fmt"
)

type Level int
type ErrorLevelInvalidValue int
type ErrorLevelInvalidString string

func (self ErrorLevelInvalidValue) Error() string {
	return fmt.Sprintf("Invalid Level value: %v", self)
}

func (self ErrorLevelInvalidString) Error() string {
	return fmt.Sprintf("Invalid Level string: %v", self)
}

const (
	MIN     Level = -3
	TRACE   Level = -2
	DEBUG   Level = -1
	INFO    Level = 0
	WARN    Level = 1
	ERROR   Level = 2
	CRIT    Level = 3
	MAX     Level = 4
	DEFAULT Level = INFO
)

var levelString = []string{
	"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "CRIT",
}

func (self Level) String() string {
	if MIN < self && self < MAX {
		return levelString[self+2]
	} else {
		panic(ErrorLevelInvalidValue(self))
	}
}

func LevelParse(value string) (Level, error) {
	for index, str := range levelString {
		if str == value {
			return Level(int(MIN) + 1 + index), nil
		}
	}
	return MAX, ErrorLevelInvalidString(value)
}

func (self *Level) UnmarshalJSON(data []byte) error {
	data = data[1 : len(data)-1]
	if result, err := LevelParse(string(data)); err == nil {
		*self = result
		return nil
	} else {
		return err
	}
}

func (self *Level) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var data string
	unmarshal(&data)
	if result, err := LevelParse(string(data)); err == nil {
		*self = result
		return nil
	} else {
		return err
	}
}

func (self Level) MarshalJSON() ([]byte, error) {
	if self < MAX {
		return []byte(`"` + self.String() + `"`), nil
	} else {
		return nil, ErrorLevelInvalidValue(self)
	}
}

func (self Level) MarshalYAML() (interface{}, error) {
	if self < MAX {
		return self.String(), nil
	} else {
		return nil, ErrorLevelInvalidValue(self)
	}
}

/*type LevelAtomic atomic.Value

func (self *LevelAtomic) Get() Level {
	return Level((atomic.Value(*self)).Load())
}

func (self *LevelAtomic) Set(level Level) {
	(atomic.Value(*self)).Store(level)
}*/
