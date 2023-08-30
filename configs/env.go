package configs

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

// helper function to load the environment

func EnvMongoURI() string {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    return os.Getenv("MONGOURI")
}
