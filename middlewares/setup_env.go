package middlewares

import (
	"os"
	"web/config"

	log "github.com/sirupsen/logrus"
)

func SetupEnv() *config.EnvConfig {
	env, err := config.ReadEnvFile(os.Getenv("ENV"))
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Fatal("Could not read env file")
	}

	log.WithFields(log.Fields{
		"env": env,
	}).Info("Read env file")

	return env
}
