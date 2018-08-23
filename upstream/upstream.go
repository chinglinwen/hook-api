package upstream

import (
	"encoding/json"
	"flag"
	"fmt"
	"strings"

	"github.com/chinglinwen/log"
	resty "gopkg.in/resty.v1"
)

var (
	UpstreamBase       = flag.String("upstream", "http://upstream-test.sched.qianbao-inc.com:8010", "upstream base api url")
	UpstreamAddAPI     = *UpstreamBase + "/add_nginx_upstream/"
	UpstreamDelAPI     = *UpstreamBase + "/wk_deleted_upstream/"
	UpstreamnChangeAPI = *UpstreamBase + "/up_nginx_state/"

	UpstreamAllAPI    = *UpstreamBase + "/get_upstream_all_instance/"
	UpstreamSingleAPI = *UpstreamBase + "/get_nginx_all/"
)

type Upstream struct {
	Name      string
	Namespace string
	State     string // "0 disable, 1 enable"
	Env       string // "pre|pro|qa"
	IP        string // single ip only for now
	Port      string
	IsDocker  string // "0 is vm ,1 is docker"
	NginxGrp  string // nginx group, "BJ-SH"

	// "fail_timeout": "int default 30"
	// "weight": "int default 1000"
}

func (u *Upstream) Add() error {
	resp, err := resty.SetRetryCount(3).
		//SetDebug(true).
		R().SetFormData(map[string]string{
		"wk_name":   u.Name,
		"namespace": u.Namespace,
		"state":     u.State,
		"ip_list":   u.IP,
		"port":      u.Port,
		"env":       u.Env,
		"is_docker": u.IsDocker,
		"nginx":     u.NginxGrp,
	}).
		Post(UpstreamAddAPI)

	if err != nil {
		return err
	}
	log.Println("resp: ", limit(resp.Body()))

	state, err := parseState(resp.Body())
	if err != nil {
		return fmt.Errorf("upstream add err resp: %v", limit(resp.Body()))
	}
	if state != true {
		return fmt.Errorf("upstream add failed resp: %v", limit(resp.Body()))
	}

	return nil
}

func limit(body []byte) string {
	n := len(body)
	if n >= 100 {
		return string(body[n-100 : n])
	}
	return string(body)
}

func (u *Upstream) Del() error {
	resp, err := resty.SetRetryCount(3).
		//SetDebug(true).
		R().SetFormData(map[string]string{
		"wk_name":   u.Name,
		"namespace": u.Namespace,
		"ip_list":   u.IP,
		"port":      u.Port,
		"env":       u.Env,
		"nginx":     u.NginxGrp,
	}).
		Post(UpstreamDelAPI)

	if err != nil {
		return err
	}
	log.Println("resp: ", strings.Replace(limit(resp.Body()), "\n", "", -1))

	state, err := parseState(resp.Body())
	if err != nil {
		return fmt.Errorf("upstream add err resp: %v", limit(resp.Body()))
	}
	if state != true {
		return fmt.Errorf("upstream add failed resp: %v", limit(resp.Body()))
	}

	return nil
}

// ChangeState change project specific ip state, remove item from nginx
// The logic may need to distinguish VM and docker
// We currently check based on docker first.
//
// Upstream will make it disabled ( need rethink?)
// Service recovery need human manual operation.
func ChangeState(endpoint, title, state string) (bool, error) {
	ip, port := endpoint2ip(endpoint)
	resp, err := resty.SetRetryCount(3).
		//SetDebug(true).
		R().SetFormData(map[string]string{
		"appname": title,
		"ip":      ip,
		"port":    port,
		"state":   state, // int 1:up or 0:down
	}).
		Post(UpstreamnChangeAPI)

	if err != nil {
		return false, err
	}

	//log.Println("resp: ", string(resp.Body()))
	result, err := parseState(resp.Body())
	if result != true {
		log.Println("ChangeState resp: ", string(resp.Body()))
	}
	return result, err
}

func parseState(body []byte) (state bool, err error) {
	var result []interface{}
	err = json.Unmarshal(body, &result)
	if err != nil || len(result) == 0 {
		return
	}
	if state, _ = result[0].(bool); state != true {
		return
	}
	return true, err
}

func endpoint2ip(e string) (ip, port string) {
	str := strings.Split(e, "/")
	var s string
	if len(str) > 2 {
		s = str[2]
	} else if len(str) == 1 {
		s = str[0]
	}
	ipport := strings.Split(s, ":")
	if len(ipport) == 2 {
		ip, port = ipport[0], ipport[1]
	}
	return
}
