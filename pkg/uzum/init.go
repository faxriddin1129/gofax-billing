package uzum

import (
	"gofax-billing/pkg/env"
)

var (
	SERVICE_ID       string
	SECRET_KEY       string
	MERCHANT_ID      string
	MERCHANT_USER_ID string
)

func LoadEnv() {
	SERVICE_ID = env.GetEnv("UZUM_SERVICE_ID")
	SECRET_KEY = env.GetEnv("UZUM_SECRET_KEY")
	MERCHANT_ID = env.GetEnv("UZUM_MERCHANT_ID")
	MERCHANT_USER_ID = env.GetEnv("UZUM_MERCHANT_USER_ID")
}

func init() {
	env.LoadEnv()
	LoadEnv()
}
