package keyservice

/**
 * all constants and global variables are defined here...
 */

import (
	"flag"
	"fmt"
	"github.com/darrylwest/cassava-logger/logger"
	"os"
	"path"
)

const (
	version = "0.90.100"
)

type Context struct {
	env          string
	logpath      string
	logname      string
	baseport     int
	shutdownPort int
	serverCount  int
	workFolder   string
	configFile   string
}

var (
	log            *logger.Logger
	currentContext Context
	config         *Config
)

func Version() string {
	return version
}

func IsProduction(env string) bool {
	return env == "production"
}

func NewDefaultContext() *Context {
	ctx := new(Context)

	home := os.Getenv("HOME")

	ctx.env = "production"
	ctx.logpath = path.Join(home, "logs")
	ctx.logname = "keyservice"

	ctx.baseport = 9001
	ctx.shutdownPort = 9009
	ctx.serverCount = 2

	ctx.workFolder = path.Join(home, ".keyservice")
	ctx.configFile = path.Join(ctx.workFolder, "config.json")

	return ctx
}

func NewContextForEnvironment(env string) *Context {
	ctx := NewDefaultContext()

	ctx.env = env

	if !IsProduction(env) {
		ctx.logname = env + "-keyservice"
	}

	return ctx
}

func ParseArgs() *Context {
	dflt := NewDefaultContext()

	vers := flag.Bool("version", false, "show the version and exit")

	env := flag.String("env", dflt.env, "set the environment, defaults to "+dflt.env)

	baseport := flag.Int("baseport", dflt.baseport, "set the server's base port number (e.g., 3001)...")
	serverCount := flag.Int("serverCount", dflt.serverCount, "set the number of server/listeners")
	shutdownPort := flag.Int("shutdownPort", dflt.shutdownPort, "set the service shutdown port")

	logpath := flag.String("logpath", dflt.logpath, "set the log directory")
	logname := flag.String("logname", dflt.logname, "set the name of the rolling log file")

	workFolder := flag.String("workFolder", dflt.workFolder, "set the application's working folder")
	configFile := flag.String("configFile", dflt.configFile, "set the configuration file")

	flag.Parse()

	fmt.Printf("%s Version: %s\n", path.Base(os.Args[0]), Version())

	if *vers == true {
		os.Exit(0)
	}

	ctx := new(Context)

	ctx.env = *env

	ctx.logpath = *logpath
	ctx.logname = *logname

	ctx.baseport = *baseport
	ctx.shutdownPort = *shutdownPort
	ctx.serverCount = *serverCount

	ctx.workFolder = *workFolder
	ctx.configFile = *configFile

	return ctx
}

func (c *Context) CreateLogger() *logger.Logger {
	if log == nil {
		filename := path.Join(c.logpath, c.logname)
		handler, err := logger.NewRotatingDayHandler(filename)

		if err != nil {
			panic("logger could not be created")
		}

		fmt.Printf("created logger at %s\n", filename)

		log = logger.NewLogger(handler)
	}

	return log
}

func (c *Context) GetShutdownPort() int {
	return c.shutdownPort
}

func (c *Context) ToMap() map[string]interface{} {
	hash := make(map[string]interface{})

	hash["env"] = c.env
	hash["logpath"] = c.logpath
	hash["logname"] = c.logname

	hash["baseport"] = c.baseport
	hash["shutdownPort"] = c.shutdownPort
	hash["serverCount"] = c.serverCount

	hash["workFolder"] = c.workFolder
	hash["configFile"] = c.configFile

	return hash
}

func (c Context) StartService() error {
	if log == nil {
		log = c.CreateLogger()
	}

	log.Info("StartService, version: %s, env: %s", version, c.env)

	if config == nil {
		log.Info("read configuration from: %s", c.configFile)
		conf, err := ReadConfig(c.configFile)

		if err != nil {
			panic(err)
		}

		log.Info("config parsed, name: %s", conf.name)
		config = conf
	}

	log.Info("start the servers with context: %v", c.ToMap())

	return nil
}
