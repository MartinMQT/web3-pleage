package common

import "my-pledge/config"

const (
	CN      = "CN"
	EN      = "EN"
	Success = 200
)

var Msg = map[int]map[string]string{

	Success: {
		CN: "成功",
		EN: "SUCC",
	},
}

func getMsg(code int) string {
	return Msg[code][config.Config.Env.Lang]
}
