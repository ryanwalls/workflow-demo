// Package logger should be used for all logging.  Other code should reference this package instead of referencing log15 directly
package logger

import (
	"github.com/3dsim/workflow/config"
	log "github.com/inconshreveable/log15"
	logext "github.com/inconshreveable/log15/ext"
	"os"
)

// Log is the root log for the rest of the application.  It is an instance of a
// github.com/inconshreveable/log15 Logger.  It has been setup with
// the following handlers:
//
// FatalHandler makes a "Crit" log stop the main execution
// LvlFilterHandler filters out messages below the "LogLevel" config property (see config-*.yml)
// CallerFileHandler prints the file and line number from where the line was logged
// (Non prod) StdoutHandler prints the logs in a nice format for reading from standard out
// (Prod only) JsonFormat prints the logs as json
var Log log.Logger

func init() {
	Log = log.New()
	var handler log.Handler
	env := config.Viper.GetString("env")
	lvl, err := log.LvlFromString(config.Viper.GetString("LogLevel"))
	if err != nil {
		panic("Could not read configuration for LogLevel, check the 'config-" + env +
			".json' file: " + err.Error())
	}
	if env == "dev" || env == "qa" {
		handler = logext.FatalHandler(log.LvlFilterHandler(lvl, log.CallerFileHandler(log.StdoutHandler)))
	} else {
		handler = logext.FatalHandler(log.LvlFilterHandler(lvl, log.CallerFileHandler(log.StreamHandler(os.Stdout, log.JsonFormat()))))
	}
	Log.SetHandler(handler)
}
