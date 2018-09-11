package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"wen/hook-api/upstream"
	"wen/svc-d/check"

	"github.com/chinglinwen/log"
	"github.com/labstack/echo"
)

/* var (
	prePod, prePhase string
	preTime          time.Time
) */

// should skip useless namespace too?
// did we should consider check before notify nginx add ip?
func hookHandler(c echo.Context) error {
	podname := c.FormValue("podname")
	appname := c.FormValue("appname")
	namespace := c.FormValue("namespace")
	ip := c.FormValue("ip")
	port := c.FormValue("port")
	phase := c.FormValue("state")

	if podname != "" {
		var err error
		appname, err = cutName(podname)
		if err != nil {
			e := fmt.Errorf("cutName err: %v", err)
			log.Println(e)
			return e
		}
		// it will cause error, since the name currently is not reversible
		//appname = strings.Replace(appname, "-", "_", -1)
	}

	out := fmt.Sprintf("name: %v, ns: %v,ip: %v, port: %v, phase: %v", appname, namespace, ip, port, phase)
	log.Printf("hook: %v\n", out)

	if *testproject != "" && appname != *testproject {
		log.Println("skip non-testing project")
		return nil
	}

	/* 	curPod := fmt.Sprintf("%v %v:%v", appname, ip, port)
	   	if curPod == prePod && phase == "ADD" && prePhase == "DEL" && time.Now().Sub(preTime) < 5*time.Second {
	   		log.Printf("hook: pre: %v, prephase: %v, cur: %v, curphase: %v\n", prePod, prePhase, curPod, phase)
	   		log.Println("ignore useless del and add(in 5 seconds), consider it as empty change")
	   		c.String(http.StatusOK, "empty change")
	   		return nil
	   	}
	   	prePod = fmt.Sprintf("%v %v:%v", appname, ip, port)
	   	prePhase = phase
	   	preTime = time.Now() */

	// if phase=add, do a extra check first
	// say wait 10 minutes

	// for rolling update issue
	// consider roll one by one, roll one, and check one, then roll another
	// did a check, and then roll another
	//
	// TODO: what if failed after retry? useless check?
	if phase == "ADD" {
		//for the later fetch config part, it will convert to project name
		err := check.CheckLonger(appname, ip, port, 5*time.Minute)
		if err != nil {
			e := fmt.Errorf("simple tcp check for %v after 5 minutes err:%v", appname, err)
			log.Println(e)
			return e
		}
		log.Printf("simple tcp check for %v ok\n", appname)
	}

	u := upstream.Upstream{
		Name:      appname,
		Namespace: namespace,
		State:     "1", // make nginx change
		Env:       *env,
		IP:        ip,
		Port:      port,
		IsDocker:  "1",
		NginxGrp:  *nginxGrp,
	}

	var err error
	if phase == "ADD" {
		err = u.Add()
	} else if phase == "DEL" {
		err = u.Del()
	}

	if err != nil {
		e := fmt.Errorf("call upstream for %v, ns:%v %v:%v, phase: %v, err: %v", appname, namespace, ip, port, phase, err)
		log.Println(e)
		return e
	}
	result := fmt.Sprintf("call upstream for %v, ns:%v %v:%v, phase: %v ok", appname, namespace, ip, port, phase)
	log.Println(result)

	return c.String(http.StatusOK, result)
}

// remove pod name's suffix
func cutName(podname string) (name string, err error) {
	s := strings.Split(podname, "-")
	n := len(s)
	if n < 3 {
		err = fmt.Errorf("podname invalid")
		return
	}
	suffix := fmt.Sprintf("-%v-%v", s[n-2], s[n-1])
	name = strings.TrimSuffix(podname, suffix)
	return
}
