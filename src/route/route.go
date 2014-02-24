/*
  route.go by weizhao
  route for transit
  1. sort by id
  2.
*/

package route

import (
	"encoding/json"
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
}

func (r *Route) Init(b string) {
	if json.Unmarshal([]byte(b), &r) != nil {
		panic("Parse json failed.")
	}
}

// 是否匹配本条路由的规则
func (r *Route) IsMatch(request string) (ok bool) {
	source_path := request
	ok = r.IsMatchSourcePaths(source_path)
	return
}

// 是否符本路由的的SourcePaths规则
func (r *Route) IsMatchSourcePaths(source_path string) (ok bool) {
	for _, path := range r.SourcePaths {
		if path == source_path {
			ok = true
			return
		}
	}
	return
}
