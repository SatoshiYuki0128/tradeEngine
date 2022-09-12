package controller

import (
	"net/http"
	"tradeEngine/model"
	"tradeEngine/service"
)

// Test controllerCode:T1
func Test(w http.ResponseWriter, r *http.Request) {
	flowData := model.FlowData{}
	if service.GetAndValidateRequest[model.Test](r, &flowData, "T1") {
		service.Test(&flowData, "T1")
	}
	service.ServeResponse(w, &flowData)
}
