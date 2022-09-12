package service

import (
	"net/http"
	"tradeEngine/model"
)

func GetAndValidateRequest[T any](httpRequest *http.Request, flowData *model.FlowData, controllerCode string) (result bool) {
	initFlowData(flowData, controllerCode, "C1")
	data, isOK := getRequestBody(httpRequest, flowData, controllerCode, "C1")
	if isOK {
		checkModel[T](data, httpRequest, flowData, controllerCode, "C1")

		if flowData.Err == nil {
			result = true
		}
	}

	return
}

func ServeResponse(w http.ResponseWriter, flowData *model.FlowData) {
	JsonResponse(w, flowData)
}

// Test serviceCode:T1
func Test(flowData *model.FlowData, controllerCode string) {
	req := flowData.Request.(model.Test)
	result := ""
	for i := 0; i < req.Num; i++ {
		result += "hello"
	}

	flowData.Response = result
}
