package service

import (
	"net/http"
	"tradeEngine/model"
)

//GetAndValidateRequest serviceCode:C1
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

//SetError FunctionCode:N/A
func SetError(flowData *model.FlowData, controllerCode, serviceCode, functionCode, errorMessage string, baseError error) {
	servErr := model.ServError{ServCode: serviceCode, FuncCode: functionCode, Msg: errorMessage, Err: baseError}
	flowData.CtrlError = model.CtrlError{CtrlCode: controllerCode, ServError: servErr}
}

// Test serviceCode:T1
//func Test(flowData *model.FlowData, controllerCode string) {
//	req := flowData.Request.(model.Test)
//	result := ""
//	for i := 0; i < req.Num; i++ {
//		result += viper.GetString("DB.host")
//	}
//
//	addr := viper.GetString("DB.connectString")
//	db, err := gorm.Open(mysql.Open(addr), &gorm.Config{})
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	flowData.Response = db
//}
