package octo

import (
	"gofax-billing/pkg/env"
	"strconv"
)

var (
	OCTO_SHOP      string
	OCOT_SECRET    string
	OCTO_TEST_MODE bool
	OCTO_URL       string
)

func LoadEnv() {
	OCTO_SHOP = env.GetEnv("OCTO_SHOP")
	OCOT_SECRET = env.GetEnv("OCOT_SECRET")
	OCTO_TEST_MODE, _ = strconv.ParseBool(env.GetEnv("OCTO_TEST_MODE"))
	OCTO_URL = env.GetEnv("OCTO_URL")
}

func init() {
	env.LoadEnv()
	LoadEnv()
}
