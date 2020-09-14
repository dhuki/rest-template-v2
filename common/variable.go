package common

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

// create type of contextKey for key context
type (
	contextKey uint
)

// cons db
var (
	DbUsername string
	DbPassword string
	DbName     string
	DbPort     string
	DbHost     string
)

// cons redis
var (
	RedisHost string
	RedisPort string
)

// cons redis
var (
	ElasticsHost string
	ElasticsPort string
)

// cons url
var (
	Host    string
	Port    string
	BaseUrl string
)

// cons email
var (
	EmailUsername string
	EmailPassword string
	EmailSmtpHost string
	EmailSmtpPort string
)

// cons error
var (
	ErrDataNotFound = errors.New("Data Not Found")
	ErrAssertion    = errors.New("Error Assertion")
	ErrCancelled    = errors.New("Request Cancelled")
	ErrLimitExceed  = errors.New("Request Limit Exceeded")
)

func LoadCons(path string) error {
	// by default you can use env directly by using "go env" to see list available value
	// but if you want use another key you should load another env file using
	// library "github.com/joho/godotenv".
	// if file env is in current dir just use .Load() only w/o parameter
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	Host = os.Getenv("base.host")
	Port = os.Getenv("base.port")
	BaseUrl = os.Getenv("base.url")

	DbUsername = os.Getenv("db.username")
	DbPassword = os.Getenv("db.password")
	DbName = os.Getenv("db.name")
	DbPort = os.Getenv("db.port")
	DbHost = os.Getenv("db.host")

	RedisHost = os.Getenv("redis.host")
	RedisPort = os.Getenv("redis.port")

	ElasticsPort = os.Getenv("elastics.port")
	ElasticsPort = os.Getenv("elastics.port")

	return nil
}
