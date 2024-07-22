package config

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

var env map[string]string
var once sync.Once

func ReadEnv() map[string]string {
	once.Do(func() {

		path := os.Getenv("ENV_PATH")

		if path == "" {
			path = ".env"
		}

		rawEnvFile, err := os.ReadFile(path)

		if err != nil {
			log.Fatal(err.Error())
			panic(0)
		}

		fmt.Println("Env File Read Successfully")

		envFileLines := strings.Split(string(rawEnvFile), "\n")
		envVars := make(map[string]string)

		for _, line := range envFileLines {
			// Ignore Comments
			if line[0] == '#' {
				continue
			}
			lineValues := strings.Split(line, "=")
			key, value := lineValues[0], lineValues[1]
			envVars[key] = value
		}

		env = envVars
	})

	return env
}

var config map[string]string = ReadEnv()

var Port string = config["PORT"]
var DBConnectionString string = config["DB_CONNECTION_STRING"]
