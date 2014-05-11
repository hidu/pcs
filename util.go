package pcs

import (
	"encoding/json"
)

func paths_param_build(paths []string) (string, error) {
	param_map := make(map[string][]map[string]string)
	param_map["list"] = []map[string]string{}
	for _, path := range paths {
		m := make(map[string]string)
		m["path"] = path
		param_map["list"] = append(param_map["list"], m)
	}

	param_byte, err := json.Marshal(param_map)
	if err != nil {
		return "", err
	}
	return string(param_byte), nil
}
