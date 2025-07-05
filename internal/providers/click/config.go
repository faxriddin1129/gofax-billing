package click

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	ServiceID      string
	MerchantID     string
	MerchantUserID string
	SecretKey      string
	Logo           string
	ServiceURL     string
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ServiceID = os.Getenv("CLICK_SERVICE_ID")
	MerchantID = os.Getenv("CLICK_MERCHANT_ID")
	MerchantUserID = os.Getenv("CLICK_MERCHANT_USER_ID")
	SecretKey = os.Getenv("CLICK_SECRET_KEY")
	Logo = os.Getenv("CLICK_LOGO")
	ServiceURL = os.Getenv("CLICK_SERVICE_URL")
	SecretKey = os.Getenv("CLICK_SECRET_KEY")
}

func init() {
	LoadEnv()
}
