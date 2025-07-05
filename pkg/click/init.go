package click

import (
	"microservice/pkg/env"
)

var (
	SERVICE_ID       string
	MERCHANT_ID      string
	MERCHANT_USER_ID string
	SECRET_KEY       string
	SERVICE_URL      string
)

func LoadEnv() {
	SERVICE_ID = env.GetEnv("CLICK_SERVICE_ID")
	MERCHANT_ID = env.GetEnv("CLICK_MERCHANT_ID")
	MERCHANT_USER_ID = env.GetEnv("CLICK_MERCHANT_USER_ID")
	SECRET_KEY = env.GetEnv("CLICK_SECRET_KEY")
	SERVICE_URL = env.GetEnv("CLICK_SERVICE_URL")
}

func init() {
	env.LoadEnv()
	LoadEnv()
}
