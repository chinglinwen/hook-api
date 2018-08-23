package main

import (
	"flag"

	"github.com/chinglinwen/log"
	"github.com/labstack/echo"
)

var (
	env         = flag.String("env", "qa", "env includes (qa,pre,pro)")
	nginxGrp    = flag.String("nginx-grp", "BJ-SH", "nginx group for upstream includes (BJ-SH,)")
	port        = flag.String("p", "8081", "listening port")
	testproject = flag.String("test", "", "test project(wk name)")
)

//  define a global variable
//  add new check, update it, and store the config as file(update config)

func init() {
	log.Println("starting...")
	log.Debug.Println("debug is on")

	if *testproject != "" {
		log.Println("test project only, project: ", *testproject)
	}

	flag.Parse()
}

func main() {

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
