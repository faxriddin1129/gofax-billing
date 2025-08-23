package ipak

import (
	"gofax-billing/pkg/env"
)

var (
	IPAK_URL   string
	ASIA_TOKEN string
)

func LoadEnv() {
	IPAK_URL = env.GetEnv("IPAK_URL")
	ASIA_TOKEN = env.GetEnv("IPAK_ASIA_TOKEN")
}

func init() {
	env.LoadEnv()
	LoadEnv()
}
