package pkg

import (
	"flag"
	"strings"
)

var (
	JSONFlag  bool
	DebugFlag bool
	ColorFlag bool

	DBUser  string
	DBPass  string
	DBHost  string
	DBName  string
	DBTable []string
)

func stringSliceVar(value *[]string, flagName string, defaultValue string, usage string) {
	flag.Func(flagName, usage, func(val string) error {
		*value = strings.Split(val, ",")
		return nil
	})
}

func Init() {
	flag.BoolVar(&DebugFlag, "debug", false, "Enable Debug Mode (Default: false)")
	flag.BoolVar(&JSONFlag, "json", false, "Enable JSON View (Default: false)")
	flag.BoolVar(&ColorFlag, "color", true, "Enable Color View\n(example: -color=true/false)\n(WARNING: IF json - color disabled)")

	flag.StringVar(&DBUser, "dbuser", "", "Database User\n(example: -dbuser=test)\n(ENV: DBUSER)")
	flag.StringVar(&DBPass, "dbpass", "", "Database Password\n(example: -dbpass=test)\n(ENV: DBPASS)")
	flag.StringVar(&DBHost, "dbhost", "", "Database Host\n(example: -dbhost=localhost:3306)\n(ENV: DBHOST)")
	flag.StringVar(&DBName, "dbname", "", "Database Name\n(example: -dbname=test)\n(ENV: DBNAME)")
	stringSliceVar(&DBTable, "dbtable", "", "Database Tables\n(example: -dbtable=example,test)")

	flag.Parse()
}
