package route

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

//测试配置文件解析
func TestHello(t *testing.T) {
	fmt.Println("hello")
	route := new(Route)
	route.Id = 4
	fmt.Println(route.Id)
}

var route Route // route for test

func TestInit(t *testing.T) {
	r := `{
            "id": 2000,
            "source_path": [
                "/ticket/req.do",
                "/newticket/req.do"
            ],
            "source_params": [
                "processcode=11002",
                "processcode=11003",
                "DeviceID=12345609888"
            ],
            "target_server": "http://10.150.150.247",
            "target_path": "/times/list",
            "target_param_name_swap": {
                "processcode": "ProcessCode",
                "DeviceID": "DeviceId",
                "city_id": "cityId"
            }
        }`
	route.Init(r)
}

func TestMatch(t *testing.T) {
	req_url := `http://localhost:10003/ticket/req.do?processcode=11002&mobile=10`
	req, _ := http.NewRequest("GET", req_url, nil)
	fmt.Println(req.URL.Path, strings.Split(req.URL.RawQuery, "&"))
	if res := route.IsMatch(req); !res {
		t.Error("#isMatch (request) match error !!")
	}
}

func TestMatchPost(t *testing.T) {
	req_url := `http://localhost:10003/ticket/req.do`
	data := url.Values{}
	data.Set("processcode", "11002")
	data.Set("mobile", "15001028927")
	req, _ := http.NewRequest("POST", req_url, bytes.NewBufferString(data.Encode()))
	fmt.Println(req.URL.Path, req.Body)
	if res := route.IsMatch(req); !res {
		t.Error("#isMatch Post(request) match error !!")
	}
}

func TestMatchMaxinPost(t *testing.T) {
	req_url := `http://localhost:10003/ticket/req.do?processcode=11002`
	data := url.Values{}
	data.Set("mobile", "15001028927")
	req, _ := http.NewRequest("POST", req_url, bytes.NewBufferString(data.Encode()))
	fmt.Println(req.URL.Path, req.Body)
	if res := route.IsMatch(req); !res {
		t.Error("#isMatch maxin Post(request) match error !!")
	}
}

func TestMatchNotMatchPathCase(t *testing.T) {
	req_url := `http://localhost:10003/ticket/req1.do?processcode=11002`
	req, _ := http.NewRequest("GET", req_url, nil)
	if res := route.IsMatch(req); res {
		t.Error("#isMatch (request) match error2 !!")
	}
}

func TestMatchNotMatchParamsCase(t *testing.T) {
	req_url := `http://localhost:10003/ticket/req.do?processcode=11005`
	req, _ := http.NewRequest("GET", req_url, nil)
	if res := route.IsMatch(req); res {
		t.Error("#isMatch (request) match error3 !!")
	}
}

func TestIsMatchSourcePaths(t *testing.T) {
	source_path := `/ticket/req.do`
	if res := route.isMatchSourcePaths(source_path); !res {
		t.Error("#IsMatchSourcePaths match error !!")
	}
}

func TestIsMatchSourceParams(t *testing.T) {
	source_params := []string{"processcode=11002", "DeviceID=12345609888"}
	if res := route.isMatchSourceParams(source_params); !res {
		t.Error("#isMatchSourceParams match error !!")
	}
}

// END
