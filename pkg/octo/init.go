package octo

import (
	"gofax-billing/pkg/env"
	"strconv"
)

var (
	OCTO_URL string

	OCTO_HIKMAT_SHOP      string
	OCOT_HIKMAT_SECRET    string
	OCTO_HIKMAT_TEST_MODE bool

	OCTO_ASIA_SHOP      string
	OCOT_ASIA_SECRET    string
	OCTO_ASIA_TEST_MODE bool
)

func LoadEnv() {
	OCTO_URL = env.GetEnv("OCTO_URL")
	OCTO_HIKMAT_SHOP = env.GetEnv("OCTO_HIKMAT_SHOP")
	OCOT_HIKMAT_SECRET = env.GetEnv("OCOT_HIKMAT_SECRET")
	OCTO_HIKMAT_TEST_MODE, _ = strconv.ParseBool(env.GetEnv("OCTO_HIKMAT_TEST_MODE"))
	OCTO_ASIA_SHOP = env.GetEnv("OCTO_ASIA_SHOP")
	OCOT_ASIA_SECRET = env.GetEnv("OCOT_ASIA_SECRET")
	OCTO_ASIA_TEST_MODE, _ = strconv.ParseBool(env.GetEnv("OCTO_ASIA_TEST_MODE"))
}

func init() {
	env.LoadEnv()
	LoadEnv()
}
