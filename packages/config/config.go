package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type FileType int

const (
	FileType_NONE FileType = iota + 1
	FileType_JSON
	FileType_YAML
)

func MustLoad(fileType FileType, configPath string, dest interface{}) (err error) {
	if configPath == "" {
		err = errors.New("error when config path is empty")
		return
	}
	file, err := os.ReadFile(configPath)
	if err != nil {
		err = errors.New(fmt.Sprintf("error when read config file: %v", err.Error()))
		return
	}
	bytesReader := bytes.NewReader(file)
	switch fileType {
	case FileType_JSON:
		jsonDecoder := json.NewDecoder(bytesReader)
		err = jsonDecoder.Decode(dest)
	case FileType_YAML:
		yamlDecoder := yaml.NewDecoder(bytesReader)
		err = yamlDecoder.Decode(dest)
	default:
		err = errors.New(fmt.Sprintf("error when file type is invalid"))
	}
	if err != nil {
		err = errors.New(fmt.Sprintf("error when decode file: %v", err.Error()))
	}
	return
}
