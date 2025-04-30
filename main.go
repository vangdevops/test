package main

import (
	"log/slog"
	"main/pkg"
	"os"
	"runtime"
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

	dbUser, present := os.LookupEnv("DBUSER")
	if !present {
		dbUser = pkg.DBUser
		if dbUser == "" {
			slog.Error("Need to set Database User!")
			os.Exit(0)
		}
	}

	dbPassword, present := os.LookupEnv("DBPASS")
	if !present {
		dbPassword = pkg.DBPass
		if dbPassword == "" {
			slog.Error("Need to set Database Password!")
			os.Exit(0)
		}
	}

	dbHost, present := os.LookupEnv("DBHOST")
	if !present {
		dbHost = pkg.DBHost
		if dbHost == "" {
			slog.Error("Need to set Database Host!")
			os.Exit(0)
		}
	}

	dbName, present := os.LookupEnv("DBNAME")
	if !present {
		dbName = pkg.DBName
		if dbName == "" {
			slog.Error("Need to set Database Name!")
			os.Exit(0)
		}
	}

	connection := "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + "/" + dbName + "?sslmode=disable"

	tables := pkg.DBTable
	if len(tables) == 0 {
		slog.Error("Need Tables!")
		os.Exit(0)
	}

	memory, err := info.Memory(syscall.Sysinfo)
	if err != nil {
		slog.Error("Error get memory: " + err.Error())
		os.Exit(1)
	}

	cpu := info.CPU()

	figure.NewColorFigure("Dragon", "graffiti", "reset", true).Print()
	slog.Info("CPU:" + cpu + " " + "Memory: " + memory + "MB")
	db, err := database.DatabaseConnect(connection)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	done := make(chan struct{})
	version, err := database.GetVersion(db)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	slog.Info("Database version: " + version)

	slog.Info("Checking Tables...")
	startcheck := time.Now()
	for _, table := range tables {
		go func(table string) {
			err := database.CheckTable(db, table)
			if err != nil {
				err := database.CreateTable(db, table)
				if err != nil {
					slog.Error(err.Error())
					os.Exit(1)
				}
				slog.Debug("Created table: " + table)
			}
			slog.Debug("Table found: " + table)
			done <- struct{}{}
		}(table)
	}
	for range tables {
		<-done
	}

	elapsed := time.Since(startcheck)
	slog.Info("Checking Succesfully! in:" + elapsed.String())
	startcheck = time.Now()
	for _, table := range tables {
		go func(table string) {
			err := database.DeleteTable(db, table)
			if err != nil {
				slog.Error("Error delete tables: " + table)
				os.Exit(1)
			}
			slog.Debug("Deleted table: " + table)
			done <- struct{}{}
		}(table)
	}

	for range tables {
		<-done
	}
	elapsed = time.Since(startcheck)
	slog.Info("Deleted Succesfully! in:" + elapsed.String())
}
