package env

import (
	"github.com/joho/godotenv"
)

const (
	envKey              = "ENV"
	envFileDevelopment  = ".env"
	envFileGithubDeploy = ".env.github-actions"
	envDevelopment      = "development"
)

var Env map[string]string

func GetEnv(key, def string) string {
	if val, ok := Env[key]; ok {
		return val
	}
	return def
}

func SetupEnvFile() {
	envFile := envFileDevelopment
	if IsProduction() {
		envFile = envFileGithubDeploy
	}

	var err error
	Env, err = godotenv.Read(envFile)
	if err != nil {
		panic(err)
	}
}

func IsProduction() bool {
	value, exist := Env[envKey]
	if !exist || value == envDevelopment {
		return false
	}

	return true
}
