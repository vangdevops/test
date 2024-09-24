package main

import (
	"github.com/common-nighthawk/go-figure"
	"syscall"
	"log/slog"
	"main/pkg"
	"os"
	"time"
	"runtime"
<<<<<<< HEAD
	"gitlab.com/vangdevops/mylibrary/database"
	"gitlab.com/vangdevops/mylibrary/info"
=======
	"github.com/vangdevops/library/database"
	"github.com/vangdevops/library/info"
>>>>>>> dev
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	pkg.Init()
	info.Log(pkg.JSONFlag,pkg.DebugFlag,pkg.ColorFlag)

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

<<<<<<< HEAD
	memory,err := info.Memory()
=======
	memory,err := info.Memory(syscall.Sysinfo)
>>>>>>> dev
	if err != nil {
		slog.Error("Error get memory: "+err.Error())
		os.Exit(1)
	}

	cpu := info.CPU()

	figure.NewColorFigure("Dragon", "graffiti","reset", true).Print()
	slog.Info("CPU:" + cpu + " "+"Memory: " + memory +"MB")
<<<<<<< HEAD
	err = database.DatabaseConnect(connection)
=======
	db,err := database.DatabaseConnect(connection)
>>>>>>> dev
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	done := make(chan struct{})
<<<<<<< HEAD
	version,err := database.GetVersion()
=======
	version,err := database.GetVersion(db)
>>>>>>> dev
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	slog.Info("Database version: " + version)

	slog.Info("Checking Tables...")
	startcheck := time.Now()
	for _,table := range tables {
		go func(table string) {
<<<<<<< HEAD
			found,err := database.CheckTable(table)
			if err != nil {
				slog.Error(err.Error())
				os.Exit(1)
			}
			if !found {
				err := database.CreateTable(table)
=======
			err := database.CheckTable(db,table)
			if err != nil {
				err := database.CreateTable(db,table)
>>>>>>> dev
				if err != nil {
					slog.Error(err.Error())
					os.Exit(1)
				}
				slog.Debug("Created table: "+table)
			}
<<<<<<< HEAD
=======
			slog.Debug("Table found: "+table)
>>>>>>> dev
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
<<<<<<< HEAD
			err := database.DeleteTable(table)
=======
			err := database.DeleteTable(db,table)
>>>>>>> dev
			if err != nil {
				slog.Error("Error delete tables: "+table)
				os.Exit(1)
			}
			slog.Debug("Deleted table: "+table)
			done <- struct{}{}
		}(table)
	}

	for range tables {
		<-done
	}
	elapsed = time.Since(startcheck)
	slog.Info("Deleted Succesfully! in:"+elapsed.String())
}
