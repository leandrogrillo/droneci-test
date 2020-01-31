package main

import (
	"log"
	"os"
)

func main() {
	log.Println(os.Getenv("DRONE_SYSTEM_HOST"))
	log.Println(os.Getenv("DRONE_REPO_NAMESPACE"))
	log.Println(os.Getenv("DRONE_REPO_NAME"))
	log.Println(os.Getenv("GOCLIENT_OAUTH_TOKEN"))
}
