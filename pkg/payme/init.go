package payme

import (
	"microservice/pkg/env"
)

var (
	PROD_KEY    string
	TEST_KEY    string
	SERVICE_URL string
	MERCHANT_ID string
)

func LoadEnv() {
	PROD_KEY = env.GetEnv("PAYME_PROD_KEY")
	TEST_KEY = env.GetEnv("PAYME_TEST_KEY")
	SERVICE_URL = env.GetEnv("PAYME_SERVICE_URL")
	MERCHANT_ID = env.GetEnv("PAYME_MERCHANT_ID")
}

func init() {
	env.LoadEnv()
	LoadEnv()
}
