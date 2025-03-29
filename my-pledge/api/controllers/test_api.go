package controllers

import (
	"net/http"
)

func PledgeHandler(w http.ResponseWriter, r *http.Request) {
	// 处理pledge请求的逻辑
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pledge successful"))
}
