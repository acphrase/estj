package core

import (
	"estj/src/exception"
	log "estj/src/logger"
	"estj/src/util"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var environment *Environment

type Environment struct {
	profile      string
	fileName     string
	envVariables map[string]interface{}
}

func init() {
	environment = new(Environment)
	environment.envVariables = make(map[string]interface{})
}

func InitEnvironment(level string) {
	if environment == nil {
		environment = new(Environment)
		environment.envVariables = make(map[string]interface{})
	}
	environment.profile = level
	setEnvironment()
}

func setEnvironment() {
	// Check profile.
	switch environment.profile {
	case "prod":
		environment.fileName = ".prod"
	case "stage":
		environment.fileName = ".stage"
	case "dev":
		environment.fileName = ".dev"
	default:
		environment.fileName = ".dev"
	}

	// Loads different environment variables depending on the profile.
	path, _ := os.Getwd()
	err := godotenv.Load(path + "/src/" + environment.fileName)
	if err != nil {
		createProfileErrors := exception.CreateProfileErrors(reflect.TypeOf(environment).String(), "")
		log.Fatal(fmt.Sprintf("%+v", errors.Wrap(createProfileErrors, createProfileErrors.GetMessage())))
	}

	// Save to memory.
	setVariable()
}

func setVariable() {
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
					createProfileErrors := exception.CreateProfileErrors(reflect.TypeOf(environment).String(), "An error occurred while converting the variable. Let's check the environment variable "+pairs[0])
					log.Fatal(fmt.Sprintf("%+v", errors.Wrap(createProfileErrors, createProfileErrors.GetMessage())))
				} else {
					environment.envVariables[pairs[0]] = convertedBool
				}
			} else {
				switch (pairs[1])[0:1] {
				case "{": // Set map value.
					convertedMap, err := util.ConvertStrToMap(pairs[1])
					if err != nil {
						createProfileErrors := exception.CreateProfileErrors(reflect.TypeOf(environment).String(), "An error occurred while converting the variable. Let's check the environment variable "+pairs[0])
						log.Fatal(fmt.Sprintf("%+v", errors.Wrap(createProfileErrors, createProfileErrors.GetMessage())))
					} else {
						environment.envVariables[pairs[0]] = *convertedMap
					}
				case "[": // Set list value.
					environment.envVariables[pairs[0]] = *(util.ConvertStrToList(pairs[1], ","))
				default: // Set string value.
					environment.envVariables[pairs[0]] = pairs[1]
				}
			}
		} else {
			// Set number value.
			environment.envVariables[pairs[0]] = num
		}
	}
}

func GetVariable(key string) interface{} {
	if _, ok := environment.envVariables[key]; ok {
		return environment.envVariables[key]
	} else {
		return nil
	}
}

func GetMapVariable(key string) map[string]interface{} {
	return ConvertVariableToMap(key)
}

func ConvertVariableToMap(key string) map[string]interface{} {
	if _, ok := environment.envVariables[key]; ok {
		return environment.envVariables[key].(map[string]interface{})
	} else {
		return nil
	}
}

func GetListVariable(key string) []string {
	return ConvertVariableToList(key)
}

func ConvertVariableToList(key string) []string {
	if _, ok := environment.envVariables[key]; ok {
		return environment.envVariables[key].([]string)
	} else {
		return nil
	}
}
