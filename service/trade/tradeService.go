package trade

import (
	"tradeEngine/model"
)

// Sell serviceCode:S1
func Sell(flowData *model.FlowData, controllerCode string) {
	_, isOK := findBuyerAndSell(flowData, controllerCode, "S1")
	if !isOK {
		return
	}
}

//1 find buyer
//2 if
