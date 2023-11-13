package env

import (
	"github.com/joho/godotenv"
	"io"
	"net/http"
)

const (
	envKey         = "ENV"
	envFile        = ".env"
	envProduction  = "production"
	envDevelopment = "development"
)

var Env map[string]string

func GetEnv(key, def string) string {
	if val, ok := Env[key]; ok {
		return val
	}
	return def
}

func SetupEnvFile(embedFS http.FileSystem) {
	file, err := embedFS.Open(envFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	Env, err = godotenv.Unmarshal(string(data))
	if err != nil {
		panic(err)
	}
}

func IsProduction() bool {
	value, exist := Env[envKey]
	if !exist || value == envDevelopment {
		return false
	}

	if value == envProduction {
		return true
	}

	return false
}
