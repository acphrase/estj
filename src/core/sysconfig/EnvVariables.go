package sysconfig

import (
	"estj/src/exception"
	"estj/src/util"
	"github.com/joho/godotenv"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var envVariables *EnvVariables

type EnvVariables struct {
	profile   string
	fileName  string
	variables map[string]interface{}
}

func init() {
	initInstance()
}

func InitEnvVariables(level string) {
	initInstance()
	envVariables.profile = level
	envVariables.setEnvVariables()
}

func initInstance() {
	if envVariables == nil {
		envVariables = new(EnvVariables)
		envVariables.variables = make(map[string]interface{})
	}
}

func GetEnvVariables() *EnvVariables {
	return envVariables
}

func (envVariables *EnvVariables) setEnvVariables() {
	// Check profile.
	switch envVariables.profile {
	case "prod":
		envVariables.fileName = ".prod"
	case "stage":
		envVariables.fileName = ".stage"
	case "dev":
		envVariables.fileName = ".dev"
	default:
		envVariables.fileName = ".dev"
	}

	// Loads different environment variables depending on the profile.
	path, _ := os.Getwd()
	err := godotenv.Load(path + "/src/" + envVariables.fileName)
	if err != nil {
		createProfileErrors := exception.CreateProfileErrors(reflect.TypeOf(envVariables).String(), "")
		panic(createProfileErrors)
	}

	// Save to memory.
	envVariables.saveEnvVariables()
}

func (envVariables *EnvVariables) saveEnvVariables() {
	// Get all env variables.
	for _, e := range os.Environ() {
		// Split env variable with "=".
		pairs := strings.SplitN(e, "=", 2)
		// Check if it is a number or not.
		num, err := strconv.Atoi(pairs[1])
		if err != nil {
			// Set boolean value.
			if pairs[1] == "True" || pairs[1] == "true" || pairs[1] == "Flase" || pairs[1] == "false" {
				convertedBool, err := strconv.ParseBool(pairs[1])
				if err != nil {
					createProfileErrors := exception.CreateProfileErrors(reflect.TypeOf(envVariables).String(), "An error occurred while converting the variable. Let's check the environment variable "+pairs[0])
					panic(createProfileErrors)
				} else {
					envVariables.variables[pairs[0]] = convertedBool
				}
			} else {
				switch (pairs[1])[0:1] {
				case "{": // Set map value.
					convertedMap, err := util.ConvertStrToMap(pairs[1])
					if err != nil {
						createProfileErrors := exception.CreateProfileErrors(reflect.TypeOf(envVariables).String(), "An error occurred while converting the variable. Let's check the environment variable "+pairs[0])
						panic(createProfileErrors)
					} else {
						envVariables.variables[pairs[0]] = *convertedMap
					}
				case "[": // Set list value.
					envVariables.variables[pairs[0]] = *(util.ConvertStrToList(pairs[1], ","))
				default: // Set string value.
					envVariables.variables[pairs[0]] = pairs[1]
				}
			}
		} else {
			// Set number value.
			envVariables.variables[pairs[0]] = num
		}
	}
}

func (envVariables *EnvVariables) GetVariable(key string) interface{} {
	if _, ok := envVariables.variables[key]; ok {
		return envVariables.variables[key]
	} else {
		return nil
	}
}

func (envVariables *EnvVariables) GetStringVariable(key string) string {
	if _, ok := envVariables.variables[key]; ok {
		return envVariables.variables[key].(string)
	} else {
		return ""
	}
}

func (envVariables *EnvVariables) GetMapVariable(key string) map[string]interface{} {
	return util.ConvertInterfaceVariableToMap(key, &envVariables.variables)
}

func (envVariables *EnvVariables) GetListVariable(key string) []string {
	return util.ConvertInterfaceVariableToList(key, &envVariables.variables)
}
