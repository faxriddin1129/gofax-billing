package click

import (
	"gofax-billing/pkg/env"
)

var (
	SERVICE_URL             string
	ASIA_SERVICE_ID         string
	ASIA_MERCHANT_ID        string
	ASIA_SECRET_KEY         string
	ASIA_MERCHANT_USER_ID   string
	HIKMAT_SERVICE_ID       string
	HIKMAT_MERCHANT_ID      string
	HIKMAT_SECRET_KEY       string
	HIKMAT_MERCHANT_USER_ID string
)

func LoadEnv() {
	SERVICE_URL = env.GetEnv("CLICK_SERVICE_URL")
	ASIA_SERVICE_ID = env.GetEnv("CLICK_ASIA_SERVICE_ID")
	ASIA_MERCHANT_ID = env.GetEnv("CLICK_ASIA_MERCHANT_ID")
	ASIA_SECRET_KEY = env.GetEnv("CLICK_ASIA_SECRET_KEY")
	ASIA_MERCHANT_USER_ID = env.GetEnv("CLICK_ASIA_MERCHANT_USER_ID")
	HIKMAT_SERVICE_ID = env.GetEnv("CLICK_HIKMAT_SERVICE_ID")
	HIKMAT_MERCHANT_ID = env.GetEnv("CLICK_HIKMAT_MERCHANT_ID")
	HIKMAT_SECRET_KEY = env.GetEnv("CLICK_HIKMAT_SECRET_KEY")
	HIKMAT_MERCHANT_USER_ID = env.GetEnv("CLICK_HIKMAT_MERCHANT_USER_ID")
}

func init() {
	env.LoadEnv()
	LoadEnv()
}
