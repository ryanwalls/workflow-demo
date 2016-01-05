// Package logger should be used for all logging.  Other code should reference this package instead of referencing log15 directly
package logger

import (
	"github.com/3dsim/workflow/config"
	log "github.com/inconshreveable/log15"
	"os"
)

// Log is the root log for the rest of the application
var Log log.Logger

func init() {
	Log = log.New()
	var handler log.Handler
	// see https://godoc.org/gopkg.in/inconshreveable/log15.v2
	env := config.Viper.GetString("env")
	lvl, err := log.LvlFromString(config.Viper.GetString("LogLevel"))
	if err != nil {
		panic("Could not read configuration for LogLevel, check the 'config-" + env +
			".json' file: " + err.Error())
	}
	if env == "dev" || env == "qa" {
		handler = log.LvlFilterHandler(lvl, log.CallerFileHandler(log.StdoutHandler))
	} else {
		handler = log.LvlFilterHandler(lvl, log.CallerFileHandler(log.StreamHandler(os.Stdout, log.JsonFormat())))
	}
	Log.SetHandler(handler)
}
