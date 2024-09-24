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

	flag.StringVar(&DBUser, "dbuser", "test", "Database User\n(ENV: DBUSER)")
	flag.StringVar(&DBPass, "dbpass", "test", "Database Password\n(ENV: DBPASS)")
	flag.StringVar(&DBHost, "dbhost", "127.0.0.1:3306", "Database Host\n(ENV: DBHOST)")
	flag.StringVar(&DBName, "dbname", "test", "Database Name\n(ENV: DBNAME)")
	stringSliceVar(&DBTable, "dbtable", "", "Database Tables")

	flag.Parse()
}
