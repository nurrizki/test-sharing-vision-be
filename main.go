package main

import (
	"flag"
	"log"
	"test-sharing-vision/internal/config"
	"test-sharing-vision/internal/database"
	"test-sharing-vision/router"
	"test-sharing-vision/service/controller"
)

func main() {
	log.Println("[START]", config.Cfg.App.Name)
	defer log.Println("[END]", config.Cfg.App.Name)

	err := database.SetupDB()
	if err != nil {
		log.Fatalf("error when setupDB, Error: %v", err)
	}

	// initialize flag for setup data dummy
	isSetupData := flag.Bool("isSetupData", false, "Flag untuk setup awal, insert data dummy")

	// parse flag
	flag.Parse()

	if *isSetupData {
		log.Println("[DATABASE] Insert data dummy")
		err := controller.InsertDataDummy()
		if err != nil {
			log.Fatalf("[DATABASE] Insert data dummy, Error: %v", err)
		}
	}

	err = router.SetupRouter()
	if err != nil {
		log.Fatalf("[SERVER] Error: %v", err)
	}
}
