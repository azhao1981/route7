package route

import (
	"fmt"
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
	route.Init([]byte(r))
}

func TestMatch(t *testing.T) {
	fmt.Printf("%s\n", route)
	request := `http://localhost:10003/callback`
	if res := route.IsMatch(request); res {
		fmt.Printf("request:%v ,isMatch is %v\n", request, res)
	}
}

// END
