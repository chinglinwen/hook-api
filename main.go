package main

import (
	"flag"
	"strings"

	"wen/hook-api/upstream"

	"github.com/chinglinwen/log"
	"github.com/labstack/echo"
)

var (
	env          = flag.String("env", "qa", "env includes (qa,pre,pro)")
	nginxGrp     = flag.String("nginx-grp", "BJ-SH", "nginx group for upstream includes (BJ-SH,)")
	port         = flag.String("p", "8081", "listening port")
	runingNS     = flag.String("n", "", "specify namespace, use comma to seperate many")
	testproject  = flag.String("test", "", "test project(wk name)")
	upstreamBase = flag.String("upstream", "http://upstream-test.sched.qianbao-inc.com:8010", "upstream base api url")

	namespaces []string
)

func getNS(namespaces string) []string {
	return strings.Split(namespaces, ",")
}

//  define a global variable
//  add new check, update it, and store the config as file(update config)

func init() {
	log.Println("starting...")
	log.Debug.Println("debug is on")

	flag.Parse()
	if *testproject != "" {
		log.Println("test project only, project: ", *testproject)
	}

	upstream.Init(*upstreamBase)
	log.Println("using upstream", *upstreamBase)

	namespaces = getNS(*runingNS)
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
