package main

import (
	"log"
	"os"

	"github.com/AntonioMartinezFernandez/hexagonal-go/config"
	http_server "github.com/AntonioMartinezFernandez/hexagonal-go/internal/http-server"
	"github.com/AntonioMartinezFernandez/hexagonal-go/pkg/logger"
	"github.com/AntonioMartinezFernandez/hexagonal-go/pkg/utils"
)

// @title Hexagonal Go
// @version 1.0
// @description Hexagonal Go
// @contact.name Antonio Martínez
// @contact.url https://github.com/AntonioMartinezFernandez
// @contact.email antonio@weffective.com
// @BasePath /api/v1
func main() {
	// Env variable that determine the config to use
	configPath := utils.GetConfigPath(os.Getenv("config"))
	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode)

	s := http_server.NewServer(cfg, appLogger)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
