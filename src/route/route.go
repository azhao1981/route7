/*
  route.go by weizhao
  route for transit
  1. sort by id
  2.
*/

package route

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type StringSlice []string
type Route struct {
	Id                  int               `json:"id"`
	SourcePaths         StringSlice       `json:"source_path"`
	SourceParams        StringSlice       `json:"source_params"`
	TargetServer        string            `json:"target_server"`
	TargetPath          string            `json:"target_path"`
	TargetParamNameSwap map[string]string `json:"target_param_name_swap"`
	ConnectionTimeout   int               `json:"connection_timeout"`
	ResponseTimeout     int               `json:"response_timeout"`
	Redirect            bool              `json:"redirect"`
	need_check_spaths   bool
	need_check_sparams  bool
}

func (r *Route) Init(b string) {
	if json.Unmarshal([]byte(b), &r) != nil {
		panic("Parse json failed.")
	}
	if len(r.SourcePaths) > 0 {
		r.need_check_spaths = true
	}
	if len(r.SourceParams) > 0 {
		r.need_check_sparams = true
	}
	fmt.Println("route ready") // delete after
}

// 是否匹配本条路由的规则
func (r *Route) IsMatch(request *http.Request) (ok bool) {
	source_path := request.URL.Path
	source_params := strings.Split(request.URL.RawQuery, "&")

	var param_ok, path_ok bool
	param_ok = !r.need_check_sparams || r.isMatchSourceParams(source_params)
	path_ok = !r.need_check_spaths || r.isMatchSourcePaths(source_path)

	ok = param_ok && path_ok
	return
}

// Matcher of SourcePaths,one or more params match will return true
func (r *Route) isMatchSourcePaths(source_path string) (ok bool) {
	for _, path := range r.SourcePaths {
		if path == source_path {
			ok = true
			return
		}
	}
	return
}

// Matcher of SourceParam, one or more params match will return true
func (r *Route) isMatchSourceParams(params []string) (ok bool) {
	for _, param := range params { // for each params
		for _, sParam := range r.SourceParams { // for each SourceParams
			if sParam == param {
				ok = true
				return
			}
		}
	}
	return
}

// END
