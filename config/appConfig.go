package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort            string
	Dsn                   string
	AppSecret             string
	TwilioAccountSid      string
	TwilioAuthToken       string
	TwilioFromPhoneNumber string
}

// function to read environment variables and return application struct

func SetupEnv() (cfg AppConfig, err error) {
	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load(".env")
	}

	httpPort := os.Getenv("HTTP_PORT")

	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("env variables not found")
	}

	// Dev
	// Dsn := os.Getenv("DSN")
	// if len(Dsn) < 1 {
	// 	return AppConfig{}, errors.New("env variables not found")
	// }
	// Production
	Dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),		
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	appSecret := os.Getenv("APP_SECRET")
	if len(appSecret) < 1 {
		return AppConfig{}, errors.New("env variables not found")
	}

	return AppConfig{ServerPort: httpPort, Dsn: Dsn, AppSecret: appSecret,
		TwilioAccountSid:      os.Getenv("TWILIO_ACCOUNT_SID"),
		TwilioAuthToken:       os.Getenv("TWILIO_AUTH_TOKEN"),
		TwilioFromPhoneNumber: os.Getenv("TWILIO_FROM_PHONE_NUMBER"),
	}, nil
}
