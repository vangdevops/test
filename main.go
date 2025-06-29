package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"main/pkg"
	"os"
	"runtime"
	"sync"
	"syscall"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/vangdevops/library/database"
	"github.com/vangdevops/library/info"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	pkg.Init()
	info.Log(pkg.JSONFlag, pkg.DebugFlag, pkg.ColorFlag)

	cfg, err := loadConfig()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	connection := createConnectionString(cfg)
	tables := pkg.DBTable
	if len(tables) == 0 {
		slog.Error("Need Tables!")
		os.Exit(1)
	}

	memory, err := info.Memory(syscall.Sysinfo)
	if err != nil {
		slog.Error("Error getting memory: " + err.Error())
		os.Exit(1)
	}
	cpu := info.CPU()

	figure.NewColorFigure("Dragon", "graffiti", "reset", true).Print()
	slog.Info("CPU:" + cpu + " Memory: " + memory + "MB")

	db, err := database.DatabaseConnect(connection)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	if err := checkAndCreateTables(db, tables); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	if err := deleteTables(db, tables); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

type dbConfig struct {
	User     string
	Password string
	Host     string
	Name     string
}

func loadConfig() (dbConfig, error) {
	cfg := dbConfig{
		User:     getEnvOrDefault("DBUSER", pkg.DBUser),
		Password: getEnvOrDefault("DBPASS", pkg.DBPass),
		Host:     getEnvOrDefault("DBHOST", pkg.DBHost),
		Name:     getEnvOrDefault("DBNAME", pkg.DBName),
	}
	if cfg.User == "" || cfg.Password == "" || cfg.Host == "" || cfg.Name == "" {
		return cfg, fmt.Errorf("database configuration incomplete")
	}
	return cfg, nil
}

func getEnvOrDefault(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

func createConnectionString(cfg dbConfig) string {
	return "postgres://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host + "/" + cfg.Name + "?sslmode=disable"
}

func checkAndCreateTables(db *sql.DB, tables []string) error {
	slog.Info("Checking Tables...")
	start := time.Now()
	var wg sync.WaitGroup
	errChan := make(chan error, len(tables))
	for _, table := range tables {
		wg.Add(1)
		go func(tbl string) {
			defer wg.Done()
			if err := database.CheckTable(db, tbl); err != nil {
				if err := database.CreateTable(db, tbl); err != nil {
					errChan <- fmt.Errorf("failed to create table %s: %w", tbl, err)
					return
				}
				slog.Debug("Created table: " + tbl)
			} else {
				slog.Debug("Table found: " + tbl)
			}
		}(table)
	}
	wg.Wait()
	close(errChan)
	if err := <-errChan; err != nil {
		return err
	}
	slog.Info("Checking Successfully! in: " + time.Since(start).String())
	return nil
}

func deleteTables(db *sql.DB, tables []string) error {
	slog.Info("Deleting Tables...")
	start := time.Now()
	var wg sync.WaitGroup
	errChan := make(chan error, len(tables))
	for _, table := range tables {
		wg.Add(1)
		go func(tbl string) {
			defer wg.Done()
			if err := database.DeleteTable(db, tbl); err != nil {
				errChan <- fmt.Errorf("error deleting table %s: %w", tbl, err)
				return
			}
			slog.Debug("Deleted table: " + tbl)
		}(table)
	}
	wg.Wait()
	close(errChan)
	if err := <-errChan; err != nil {
		return err
	}
	slog.Info("Deleted Successfully! in: " + time.Since(start).String())
	return nil
}
