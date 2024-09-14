package main

import (
	"github.com/common-nighthawk/go-figure"
	"log/slog"
	"main/pkg"
	"os"
	"time"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	pkg.Init()
	pkg.Log(pkg.JSONFlag,pkg.DebugFlag,pkg.ColorFlag)

	dbUser, present := os.LookupEnv("DBUSER")
	if !present {
		dbUser = pkg.DBUser
		if dbUser == "" {
			slog.Error("Need to set Database User!")
			os.Exit(1)
		}
	}

	dbPassword, present := os.LookupEnv("DBPASS")
	if !present {
		dbPassword = pkg.DBPass
		if dbPassword == "" {
			slog.Error("Need to set Database Password!")
			os.Exit(1)
		}
	} 

	dbHost, present := os.LookupEnv("DBHOST")
	if !present {
		dbHost = pkg.DBHost
		if dbHost == "" {
			slog.Error("Need to set Database Host!")
			os.Exit(1)
		}
	}

	dbName, present := os.LookupEnv("DBNAME")
	if !present {
		dbName = pkg.DBName
		if dbName == "" {
			slog.Error("Need to set Database Name!")
			os.Exit(1)
		}
	}

	connection := dbUser + ":" + dbPassword + "@tcp" + "(" + dbHost + ")/" + dbName

	tables := pkg.DBTable
	if len(tables) == 0 {
		slog.Error("Need Tables!")
		os.Exit(1)
	}

	figure.NewColorFigure("Dragon", "graffiti","reset", true).Print()
	slog.Info("CPU:" + pkg.CPU() + " "+"Memory: " + pkg.Memory()+"MB")
	pkg.DatabaseConnect(connection)

	done := make(chan struct{})
	pkg.GetVersion()
	slog.Info("Checking Tables...")
	startcheck := time.Now()
	for _,table := range tables {
		go func(table string) {
			pkg.CheckTable(table)
			done <- struct{}{}
		}(table)
	}
	for range tables {
		<-done
	}
	
	elapsed := time.Since(startcheck)
	slog.Info("Checking Succesfully! in:"+elapsed.String())
	startcheck = time.Now()
	for _,table := range tables {
		go func(table string) {
			pkg.DeleteTable(table)
			done <- struct{}{}
		}(table)
	}

	for range tables {
		<-done
	}
	elapsed = time.Since(startcheck)
	slog.Info("Deleted Succesfully! in:"+elapsed.String())
}
