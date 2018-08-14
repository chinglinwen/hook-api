package main

import (
	"flag"

	"github.com/chinglinwen/log"
	"github.com/labstack/echo"
)

var (
	env                = flag.String("env", "test", "env includes (test,pre,pro)")
	port               = flag.String("p", "8080", "port")
	testproject        = flag.String("t", "ops_test", "test project")
	checkonetime       = flag.Bool("once", false, "check only once")
	concurrentChecks   = flag.Int("cc", 100, "number of concurrent checks")
	upstreamAPI        = flag.String("upstream", "http://upstream-pre.sched.qianbao-inc.com/get_upstream_all_instance/", "upstream fetch api url")
	upstreamnChangeAPI = flag.String("upstreamc", "http://upstream-pre.sched.qianbao-inc.com/up_nginx_state/", "upstream change api url")
)

//define a global variable
//add new check, update it, and store the config as file(update config)

func main() {
	log.Println("starting...")
	log.Debug.Println("debug is on")

	e := echo.New()
	//e.Use(middleware.Logger())

	//e.Use(middleware.Recover())
	//e.Use(middleware.Static("/static"))

	e.POST("/hook", hookHandler)

	err := e.Start(":" + *port)
	log.Println("fatal", err)
	//e.Logger.Fatal()

	log.Println("exit")
}
