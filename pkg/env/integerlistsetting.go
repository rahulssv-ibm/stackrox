package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IntegerListSetting struct {
	envVar       string
	defaultValue []int
}

// EnvVar returns the string name of the environment variable
func (s *IntegerListSetting) EnvVar() string {
	return s.envVar
}

// DefaultValue returns the default vaule for the setting
func (s *IntegerListSetting) DefaultValue() []int {
	return s.defaultValue
}

// Setting returns the string form of the integer environment variable
func (s *IntegerListSetting) Setting() string {
	return fmt.Sprintf("%d", s.IntegerListSetting())
}

// IntegerListSetting returns the integer object represented by the environment variable
func (s *IntegerListSetting) IntegerListSetting() []int {
	val := os.Getenv(s.envVar)
	if val == "" {
		return s.defaultValue
	}
	var result []int
	for _, v := range strings.Split(val, ",") {
		intV, err := strconv.Atoi(v)
		if err != nil {
			return s.defaultValue
		}
		result = append(result, intV)
	}
	return result
}

// RegisterIntegerListSetting globally registers and returns a new integer setting.
func RegisterIntegerListSetting(envVar string, defaultValue []int) *IntegerListSetting {
	s := &IntegerListSetting{
		envVar:       envVar,
		defaultValue: defaultValue,
	}

	Settings[s.EnvVar()] = s
	return s
}
