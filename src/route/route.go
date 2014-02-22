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

func (r *Route) Init(b []byte) {
	if json.Unmarshal([]byte(b), &r) != nil {
		panic("Parse json failed.")
	}
}

func (r *Route) IsMatch(request string) (ok bool) {
	ok = true
	return
}
