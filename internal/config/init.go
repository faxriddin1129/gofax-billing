package config

import "gofax-billing/pkg/env"

func init() {
	// LOAD ENVIRONMENTS
	env.LoadEnv()
}
