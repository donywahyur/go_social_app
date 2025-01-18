package env

import (
	"fmt"
	"os"
	"strconv"
)

func Get(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		fmt.Println("Warning: Using default value for", key, ":", fallback)
		return fallback
	}

	return val
}
func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		fmt.Println("Warning: Using default value for", key, ":", fallback)
		return fallback
	}
	intVal, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return intVal
}
