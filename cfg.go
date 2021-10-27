package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func MustGetPropAsStringFromYaml(key string) string {
	if val := viper.GetString(key); val != "" {
		return val
	}

	fmt.Printf(`The key '%s' is required`, key)
	os.Exit(-1)
	return ""
}
