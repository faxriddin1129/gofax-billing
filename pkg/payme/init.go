package payme

import (
	"gofax-billing/pkg/env"
)

var (
	SERVICE_URL        string
	HIKMAT_PROD_KEY    string
	HIKMAT_MERCHANT_ID string
	ASIA_PROD_KEY      string
	ASIA_MERCHANT_ID   string
)

func LoadEnv() {
	SERVICE_URL = env.GetEnv("PAYME_SERVICE_URL")
	HIKMAT_PROD_KEY = env.GetEnv("PAYME_HIKMAT_PROD_KEY")
	HIKMAT_MERCHANT_ID = env.GetEnv("PAYME_HIKMAT_MERCHANT_ID")
	ASIA_PROD_KEY = env.GetEnv("PAYME_ASIA_PROD_KEY")
	ASIA_MERCHANT_ID = env.GetEnv("PAYME_ASIA_MERCHANT_ID")
}

func init() {
	env.LoadEnv()
	LoadEnv()
}
