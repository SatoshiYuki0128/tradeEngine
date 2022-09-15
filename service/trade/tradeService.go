package trade

import (
	"tradeEngine/model"
)

// Sell serviceCode:S1
func Sell(flowData *model.FlowData, controllerCode string) {
	findBuyerAndSell(flowData, controllerCode, "S1")
}

// Buy serviceCode:S2
func Buy(flowData *model.FlowData, controllerCode string) {
	findSellerAndBuy(flowData, controllerCode, "S2")
}

// Search serviceCode:S3
func Search(flowData *model.FlowData, controllerCode string) {
	searchTrade(flowData, controllerCode, "S3")
}
