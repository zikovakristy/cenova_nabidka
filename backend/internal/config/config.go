package config

import (
    "log"

    "cenova_nabidka/backend/internal/database"
)

func SetupDatabase() {
    database.Init()
    log.Println("Database setup completed")
}
