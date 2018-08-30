package upstream

import (
	"fmt"
	"testing"
)

var (
	testUstreamapi         = "http://upstream-test.sched.qianbao-inc.com:8010/get_upstream_all_instance/"
	testUpstreamnChangeAPI = "http://upstream-test.sched.qianbao-inc.com:8010/add_nginx_upstream/"
)

func TestAdd(t *testing.T) {
	u := Upstream{
		Name:      "ops-fs",
		Namespace: "qb-qa-10",
		State:     "1",
		Env:       "qa",
		IP:        "172.28.136.100",
		Port:      "8000",
		IsDocker:  "1",
		NginxGrp:  "BJ-SH",
	}
	if err := u.Add(); err != nil {
		t.Error("add error: ", err)
	}
}

func TestDel(t *testing.T) {
	u := Upstream{
		Name:      "ops-fs",
		Namespace: "qb-qa-10",
		Env:       "qa",
		IP:        "172.28.136.100",
		Port:      "8000",
		NginxGrp:  "BJ-SH",
	}
	if err := u.Del(); err != nil {
		t.Error("del error: ", err)
	}
}

func TestParseState(t *testing.T) {
	test := `[
  true, 
  "bb 104.16.25.88:80 \u6ca1\u6709\u6ce8\u518c\u5230\u8c03\u5ea6\u4e2d\u5fc3,"
]`
	fmt.Println(parseState([]byte(test)))
}

func TestChangeState(t *testing.T) {
	tests := []struct {
		endpoint, name, ns, state string
		result                    bool
	}{
		{"http://172.28.137.22:8000", "ops_fs", "qb-qa-10", "0", true}, //must be real
		//{"http://172.28.137.22:8000", "ops_fs", "qb-qa-10", "1", true}, //must be real
		//{"tcp://104.16.25.88:80", "bb","qb-qa-10","1", true},
		//{"104.16.25.88:80", "cc", "80"},
	}
	for _, v := range tests {
		ok, err := ChangeState(v.endpoint, v.name, v.ns, v.state)
		if err != nil || ok != v.result {
			t.Error("err", v, "got", ok, "want", v.result)
			continue
		}
	}
}

func BenchmarkChangeState(b *testing.B) {
	endpoint, name, ns := "http://172.28.137.144:8080", "ismsgproject_ismsgapiweb_v1", "qb-qa-10"
	var n int
	for ; n < b.N; n++ {
		ChangeState(endpoint, name, ns, "1")
	}
	fmt.Println("runed: ", n)
}

func TestEndpoint2ip(t *testing.T) {
	tests := []struct {
		endpoint, ip, port string
	}{
		{"http://104.16.25.88:80", "104.16.25.88", "80"},
		{"tcp://104.16.25.88:80", "104.16.25.88", "80"},
		{"104.16.25.88:80", "104.16.25.88", "80"},
	}
	for _, v := range tests {
		i, p := endpoint2ip(v.endpoint)
		if i != v.ip || p != v.port {
			t.Error(v.endpoint, "got", i, p, "want", v.ip, v.port)
		}
	}
}
