package main

// load required packages
import (
	"github.com/nethsaraPrabash/rabc-in-go/database"
	"github.com/nethsaraPrabash/rabc-in-go/model"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load environment file
	loadEnv()
	// load database configuration and connection
	loadDatabase()
	// start the server
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(".env file loaded successfully")
}
func loadDatabase() {
    database.InitDb()
    database.Db.AutoMigrate(&model.Role{})
    database.Db.AutoMigrate(&model.User{})
    seedData()
}

func seedData() {
    var roles = []model.Role{{Name: "admin", Description: "Administrator role"}, {Name: "customer", Description: "Authenticated customer role"}, {Name: "anonymous", Description: "Unauthenticated customer role"}}
    var user = []model.User{{Username: os.Getenv("ADMIN_USERNAME"), Email: os.Getenv("ADMIN_EMAIL"), Password: os.Getenv("ADMIN_PASSWORD"), RoleID: 1}}
    database.Db.Save(&roles)
    database.Db.Save(&user)
}


func serveApplication() {
	router := gin.Default()

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}