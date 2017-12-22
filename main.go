package logspout_amqp_rancher

import (
	"github.com/gliderlabs/logspout/router"
	"log"
	"os"
)

func init() {
	router.AdapterFactories.Register(NewAmqpAdapter, "amqp")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func getEnv(key string, defaultValue string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return
}
