package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"test-sms-2-pro/config"
	"test-sms-2-pro/internal/models"
	"test-sms-2-pro/internal/repositories/db"
	"test-sms-2-pro/internal/repositories/jsonFile"
	"test-sms-2-pro/internal/routers"
	"test-sms-2-pro/internal/services"
	"test-sms-2-pro/loggers"
	"test-sms-2-pro/utils"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Customers struct {
	ID   int
	Name string
	Age  int
}

func main() {
	cfg := config.InitConfig()
	loggers.InitLogger(cfg.App)
	utils.LoadJwtSecret(cfg.Secrets.JwtKey)
	DB := initSqlite(cfg.Sqlite)
	migrateDB(DB, models.UsersRepository{})
	pokemonData, err := readJSON(cfg.JsonDataFiles.PathFile, cfg.JsonDataFiles.NameFile)
	if err != nil {
		loggers.Fatal(fmt.Sprintf("cannot read json file error=%v", err.Error()), zap.Error(err))
	}
	jsonFile.LoadPokemonsData(pokemonData)
	// repository
	customerRepo := db.NewUsersRepository(DB)
	pokemonRepo := jsonFile.NewPokemonsRepository()
	// service
	usersSvc := services.NewUsersService(customerRepo)
	pokemonSvc := services.NewPokemonService(pokemonRepo)

	e := routers.InitRouter(usersSvc, pokemonSvc)
	go run(e, cfg.App)
	quit := make(chan os.Signal, 1)
	<-quit
	loggers.Info("receive signal: shutting down...\n")
	if err := e.Shutdown(context.Background()); err != nil {
		loggers.Fatal(err.Error())
	}

}

func run(e *echo.Echo, cfg config.App) {
	serverPort := fmt.Sprintf(":%v", cfg.Port)
	err := e.Start(serverPort)
	if err != nil {
		loggers.Fatal("router shutdown....")
	}

}

func initSqlite(configSqlite config.Sqlite) *gorm.DB {
	sqlitePath := configSqlite.Name
	if strings.TrimSpace(configSqlite.Path) != "" {
		sqlitePath = fmt.Sprintf("%v/%v", configSqlite.Path, configSqlite.Name)

	}
	db, err := gorm.Open(sqlite.Open(sqlitePath), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent), TranslateError: true})
	if err != nil {
		loggers.Fatal(fmt.Sprintf("cannot connect sqlite error=%v", err.Error()), zap.Error(err))
	}
	sql, _ := db.DB()
	sql.SetMaxIdleConns(configSqlite.MaxIdleConns)
	sql.SetMaxOpenConns(configSqlite.MaxOpenConns)
	sql.SetConnMaxLifetime(configSqlite.MaxLifeTimeMinutes)
	loggers.Info("connect DB successfully.")
	return db
}

func migrateDB(db *gorm.DB, customerTable models.UsersRepository) {

	err := db.AutoMigrate(customerTable)
	if err != nil {
		loggers.Fatal(fmt.Sprintf("AutoMigrate error:%v", err.Error()), zap.Error(err))
	}
	loggers.Info("migrate DB successfully.")
}

// ReadJSON reads a JSON file and returns map[string]interface{}
func readJSON(path, name string) (map[string]interface{}, error) {
	// Check if file exists
	filePath := fmt.Sprintf("%v/%v.json", path, name)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file does not exist: %s", filePath)
	}

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Read the file content
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	// Parse JSON to map
	var data map[string]interface{}
	if err := json.Unmarshal(content, &data); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}
	header := make(map[string]interface{})
	header[name] = data
	return header, nil
	//return data, nil
}
