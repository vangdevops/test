package pkg 

import (
	"os"
	"strings"
	"log/slog"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
	err error
	version string
)

func DatabaseConnect(connection string) {
	DB, err = sql.Open("mysql", connection)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	err = DB.Ping()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	slog.Info("Successfully connected to the database!")
}

func GetVersion() {
	err = DB.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	slog.Info("Database version: " + version)
}

func CheckTable(table string) {
	res, err := DB.Query("SHOW TABLES")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	var tablesmap []string

	for res.Next() {
		var scanedtable string
		err := res.Scan(&scanedtable) 
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)      
		}
		tablesmap = append(tablesmap, scanedtable)
	}

	err = res.Err() 
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}


	if len(tablesmap) == 0 {
		CreateTable(table)
	} else {
		tables := strings.Join(tablesmap, ",")
		slog.Debug("Tables: ["+tables+"]")
		needtable := strings.Contains(tables, table)
		slog.Debug("Checking Table: "+table)
		if !needtable {	
			CreateTable(table)
		}
	}
}

func CreateTable(table string,) {
	_,err := DB.Query("CREATE TABLE " + table + "(test INT)")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	slog.Debug("Created table: "+table)
	CheckTable(table)
}

func DeleteTable(table string) {
	_,err := DB.Query("DROP TABLE " + table)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	slog.Debug("Deleted Table: "+table)
}
