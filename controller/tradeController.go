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

// Buy controllerCode:S2
func Buy(w http.ResponseWriter, r *http.Request) {
	flowData := model.FlowData{}
	if service.GetAndValidateRequest[model.Buy](r, &flowData, "S2") {
		trade.Buy(&flowData, "S2")
	}
	service.ServeResponse(w, &flowData)
}

// Search controllerCode:S3
func Search(w http.ResponseWriter, r *http.Request) {
	flowData := model.FlowData{}
	if service.GetAndValidateRequest[model.Search](r, &flowData, "S3") {
		trade.Search(&flowData, "S3")
	}
	service.ServeResponse(w, &flowData)
}
