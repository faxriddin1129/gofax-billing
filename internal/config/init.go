package config

import "microservice/pkg/env"

func init() {
	// LOAD ENVIRONMENTS
	env.LoadEnv()
}
