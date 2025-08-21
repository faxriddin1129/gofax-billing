package main_server

import (
	"gofax-billing/pkg/env"
)

var (
	MAIN_SERVER_TOKEN string
	MAIN_SERVER_URL   string
)

func LoadEnv() {
	MAIN_SERVER_TOKEN = env.GetEnv("MAIN_SERVER_TOKEN")
	MAIN_SERVER_URL = env.GetEnv("MAIN_SERVER_URL")
}

func init() {
	env.LoadEnv()
	LoadEnv()
}
