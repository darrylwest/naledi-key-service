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
}

var (
	log            *logger.Logger
	currentContext Context
)

func (c *Context) GetShutdownPort() int {
	return c.shutdownPort
}

func IsProduction(env string) bool {
	return env == "production"
}

func (c *Context) ToMap() map[string]interface{} {
	hash := make(map[string]interface{})

	hash["env"] = c.env
	hash["logpath"] = c.logpath
	hash["logname"] = c.logname

	hash["baseport"] = c.baseport
	hash["shutdownPort"] = c.shutdownPort
	hash["serverCount"] = c.serverCount

	return hash
}

func NewDefaultContext() *Context {
	ctx := new(Context)

	ctx.env = "production"
	ctx.logpath = path.Join(os.Getenv("HOME"), "logs")
	ctx.logname = "webserver"

	ctx.baseport = 9001
	ctx.shutdownPort = 9009
	ctx.serverCount = 2

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

	return ctx
}

func CreateLogger(ctx *Context) *logger.Logger {
	if log == nil {
		filename := path.Join(ctx.logpath, ctx.logname)
		handler, err := logger.NewRotatingDayHandler(filename)

		if err != nil {
			panic("logger could not be created")
		}

		fmt.Printf("created logger at %s\n", filename)

		log = logger.NewLogger(handler)
	}

	return log
}

func Version() string {
	return version
}
