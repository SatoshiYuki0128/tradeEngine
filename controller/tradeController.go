package controller

import (
	"net/http"
	"tradeEngine/model"
	"tradeEngine/service"
	"tradeEngine/service/trade"
)

// Sell controllerCode:S1
func Sell(w http.ResponseWriter, r *http.Request) {
	flowData := model.FlowData{}
	if service.GetAndValidateRequest[model.Sell](r, &flowData, "S1") {
		trade.Sell(&flowData, "S1")
	}
	service.ServeResponse(w, &flowData)
}
