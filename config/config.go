package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		App      App
		Db       Db
		DbType   DbType
		Kafka    Kafka
		Grpc     Grpc
		Firebase Firebase
	}

	App struct {
		Name  string
		Url   string
		Stage string
	}

	Db struct {
		Url string
	}

	DbType struct {
		Type string
	}

	Kafka struct {
		Url    string
		ApiKey string
		Secret string
	}

	Grpc struct {
		AuthUrl string
		UserUrl string
	}

	Firebase struct {
		Url               string
		ApiKey            string
		ProjectId         string
		StorageBucket     string
		MessagingSenderId string
		AppId             string
		MeasurementId     string
	}
)

func LoadConfig(path string) Config {
	if err := godotenv.Load(path); err != nil {
		log.Fatal("Error loading .env file")
		log.Fatal(err.Error())
	}

	return Config{
		App: App{
			Name:  os.Getenv("APP_NAME"),
			Url:   os.Getenv("APP_URL"),
			Stage: os.Getenv("APP_STAGE"),
		},
		Db: Db{
			Url: os.Getenv("DB_URL"),
		},
		DbType: DbType{
			Type: os.Getenv("DB_TYPE"),
		},
		Kafka: Kafka{
			Url:    os.Getenv("KAFKA_URL"),
			ApiKey: os.Getenv("KAFKA_API_KEY"),
			Secret: os.Getenv("KAFKA_SECRET"),
		},
		Grpc: Grpc{
			AuthUrl: os.Getenv("GRPC_AUTH_URL"),
			UserUrl: os.Getenv("GRPC_USER_URL"),
		},
		Firebase: Firebase{
			Url:               os.Getenv("FB_URL"),
			ApiKey:            os.Getenv("FB_API_KEY"),
			ProjectId:         os.Getenv("FB_PROJECT_ID"),
			StorageBucket:     os.Getenv("FB_STORAGE_BUCKET"),
			MessagingSenderId: os.Getenv("FB_MESSAGING_SENDER_ID"),
			AppId:             os.Getenv("FB_APP_ID"),
			MeasurementId:     os.Getenv("FB_MEASUREMENT_ID"),
		},
	}
}
