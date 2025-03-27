package config

import (
	"log"
	"path/filepath"
	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

func InitGlobalEnv() {
	once.Do(func() {
		rootDir, _ := filepath.Abs("../../.env")

		err := godotenv.Load(rootDir)

		if err != nil {
			log.Fatal("Error loading .env file")
		}

	})
}
