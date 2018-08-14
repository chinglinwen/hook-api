package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/chinglinwen/log"
	"github.com/labstack/echo"
)

var prePod, prePhase string
var preTime time.Time

func hookHandler(c echo.Context) error {
	appname := c.FormValue("appname")
	ip := c.FormValue("ip")
	port := c.FormValue("port")
	phase := c.FormValue("state")

	out := fmt.Sprintf("name: %v,ip: %v, port: %v, phase: %v", appname, ip, port, phase)
	log.Printf("hook: %v\n", out)

	curPod := fmt.Sprintf("%v %v:%v", appname, ip, port)
	log.Printf("hook: pre: %v, prephase: %v, cur: %v, curphase: %v\n", prePod, prePhase, curPod, phase)

	if curPod == prePod && phase == "ADD" && prePhase == "DEL" && time.Now().Sub(preTime) < 5*time.Second {
		log.Println("ignore useless del and add(in 5 seconds), consider it as empty change")
		c.String(http.StatusOK, "empty change")
		return nil
	}
	prePod = fmt.Sprintf("%v %v:%v", appname, ip, port)
	prePhase = phase
	preTime = time.Now()

	// here call other api
	return c.String(http.StatusOK, out)
}
