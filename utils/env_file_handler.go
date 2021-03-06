package utils

import (
	"encoding/json"
	"github.com/neel1996/gitconvex/global"
	"go/types"
	"io/ioutil"
	"os"
)

type EnvConfig struct {
	DataBaseFile string `json:"databaseFile"`
	Port         string `json:"port"`
}

var logger global.Logger

// EnvConfigValidator checks if the env_config json file is present and accessible
// If the file is missing or unable to access, then an error will be thrown
func EnvConfigValidator() error {
	cwd, cwdErr := DefaultDirSetup()
	if cwdErr != nil {
		return cwdErr
	}

	fileString := cwd + "/" + global.EnvFileName
	_, openErr := os.Open(fileString)

	if openErr != nil {
		logger.Log(openErr.Error(), global.StatusError)
		return openErr
	} else {
		if envContent := EnvConfigFileReader(); envContent == nil {
			logger.Log("Unable to read env file", global.StatusError)
			return types.Error{Msg: "Invalid content in env_config file"}
		}
	}
	return nil
}

// EnvConfigFileReader reads the env_config json file and returns the config data as a struct
func EnvConfigFileReader() *EnvConfig {
	cwd, cwdErr := DefaultDirSetup()
	if cwdErr != nil {
		logger.Log(cwdErr.Error(), global.StatusError)
		return nil
	}

	fileString := cwd + "/" + global.EnvFileName
	envFile, err := os.Open(fileString)

	var envConfig *EnvConfig

	if err != nil {
		logger.Log(err.Error(), global.StatusError)
		return nil
	} else {
		if fileContent, openErr := ioutil.ReadAll(envFile); openErr == nil {
			unMarshallErr := json.Unmarshal(fileContent, &envConfig)
			if unMarshallErr == nil {
				return envConfig
			} else {
				logger.Log(unMarshallErr.Error(), global.StatusError)
				return nil
			}
		}
	}
	return nil
}

// EnvConfigFileGenerator will be invoked when the EnvConfigValidator returns an error or if EnvConfigFileReader returns no data
// The function generates a new env_config.json file and populates it with the default config data
func EnvConfigFileGenerator() error {
	cwd, cwdErr := DefaultDirSetup()
	if cwdErr != nil {
		return cwdErr
	}

	fileString := cwd + "/" + global.EnvFileName

	envContent, _ := json.MarshalIndent(&EnvConfig{
		DataBaseFile: cwd + "/" + global.DatabaseFilePath,
		Port:         global.DefaultPort,
	}, "", " ")

	return ioutil.WriteFile(fileString, envContent, 0755)
}
