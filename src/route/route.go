/*
  route.go by weizhao
  route for transit
  1. sort by id
  2.
*/

package route

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	get_query := get_get_values(request)
	post_query := get_post_values(request)
	source_params := merge_params(get_query, post_query)

	var param_ok, path_ok bool
	param_ok = !r.need_check_sparams || r.isMatchSourceParams(source_params)
	path_ok = !r.need_check_spaths || r.isMatchSourcePaths(source_path)

	ok = param_ok && path_ok
	return
}

// merge two slice
func merge_params(params_a []string, params_b []string) (params_m []string) {
	params_m = append(params_a, params_b...)
	return
}

// get Get values from request.URL.RawQuery
func get_get_values(request *http.Request) (gms []string) {
	gms = strings.Split(request.URL.RawQuery, "&")
	return
}

// get post values from request.Body
func get_post_values(request *http.Request) (pms []string) {
	if request.Body != nil {
		pms_bytes, _ := ioutil.ReadAll(request.Body)
		pms = params_values(pms_bytes)
	}
	return
}

// Get params values from "X=x&B=b" to ["X=x","B=b"]
func params_values(raw_bytes []byte) (ps []string) {
	body_buffer := bytes.NewBuffer(raw_bytes)
	ps = strings.Split(body_buffer.String(), "&")
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
