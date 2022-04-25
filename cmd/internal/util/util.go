package util

import (
	"reflect"
	"strings"
)

func FindValue(findingKey string, yamlMap map[interface{}]interface{}) (map[interface{}]interface{}, string) {
	splitKey := strings.Split(findingKey, ".")
	currentFindingKey := splitKey[0]
	splitKey = splitKey[1:]
	for key, value := range yamlMap {
		if key == currentFindingKey && len(splitKey) == 0 {
			return yamlMap, currentFindingKey
		} else if key == currentFindingKey && len(splitKey) > 0 && reflect.TypeOf(value).Kind() == reflect.Map {
			return FindValue(strings.Join(splitKey, "."), value.(map[interface{}]interface{}))
		}
	}
	return nil, ""
}
