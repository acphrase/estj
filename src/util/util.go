package util

import (
	"encoding/json"
	"reflect"
	"strings"
)

func ConvertStrToList(str string, sep string) *[]string {
	parsingStr := strings.TrimLeft(str, "[")
	parsingStr = strings.TrimRight(parsingStr, "]")
	parsingResult := strings.Split(parsingStr, sep)
	for i, v := range parsingResult {
		v1 := strings.Trim(v, " ")
		parsingResult[i] = strings.Trim(v1, "\"")
	}
	return &parsingResult
}

func ConvertStrToMap(str string) (*map[string]interface{}, error) {
	parsingResult := make(map[string]interface{})
	if err := json.Unmarshal([]byte(str), &parsingResult); err != nil {
		return nil, err
	} else {
		for key, value := range parsingResult {
			checkInterface(parsingResult, key, value)
		}
		return &parsingResult, nil
	}
}

func checkInterface(msi map[string]interface{}, key string, value interface{}) {
	valueType := reflect.ValueOf(value)
	switch valueType.Kind() {
	case reflect.Map:
		m := value.(map[string]interface{})
		msi[key] = m
		for k, v := range m {
			checkInterface(m, k, v)
		}
	case reflect.String:
		msi[key] = value.(string)
	}
}

func ConvertInterfaceVariableToMap(key string, variables *map[string]interface{}) map[string]interface{} {
	if _, ok := (*variables)[key]; ok {
		return (*variables)[key].(map[string]interface{})
	} else {
		return nil
	}
}

func ConvertInterfaceVariableToList(key string, variables *map[string]interface{}) []string {
	if _, ok := (*variables)[key]; ok {
		return (*variables)[key].([]string)
	} else {
		return nil
	}
}
