package main

import (
	"fmt"
	"os"

	// "time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"log"
)

var GOOGLE_CREDENTIALS = os.Getenv("GOOGLE_CLOUD_AUTH")
var PORT = os.Getenv("PORT")

func init() {
	var err error
	// time.Local, err = time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		panic(err)
	}
}

func main() {
	path, _ := os.Getwd()
	LoadGoogleCredentials(path)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":" + PORT))
}

func LoadGoogleCredentials(rootPath string) {
	filePath := fmt.Sprintf("%s/google_cloud_auth.json", rootPath)

	err := os.WriteFile(filePath, []byte(GOOGLE_CREDENTIALS), 0644)
	if err != nil {
		log.Fatalf("Could not write file having google credentials, error: %v", err)
	}

	err = os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", filePath)

	if err != nil {
		log.Fatalf("Could not set a environment variable having the file path, error: %v", err)
	}
}
