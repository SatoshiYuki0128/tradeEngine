package trade

import (
	"tradeEngine/model"
	"tradeEngine/service"
)

//findBuyer functionCode:S1
func findBuyer(flowData *model.FlowData, controllerCode, serviceCode string) (id int, isOK bool) {
	req := flowData.Request.(model.Sell)

	db := service.DbConnect()
	defer db.Close()

	sql := "select id from Trade where type = 'buy' and quantity = ? and price = ? and close = 0 order by id asc limit 1;"
	row := db.QueryRow(sql, req.Quantity, req.Price)
	row.Scan(&id)

	isOK = true
	return
}

//pendingOrder functionCode:S2
func pendingOrder(flowData *model.FlowData, controllerCode, serviceCode string) {
	req := flowData.Request.(model.Sell)

	db := service.DbConnect()
	defer db.Close()

	sql := "insert into Trade(type, quantity, price, close, updateTime) values ('sell', ?, ?, 0, now());"
	_, err := db.Exec(sql, req.Quantity, req.Price)
	if err != nil {
		service.SetError(flowData, controllerCode, serviceCode, "S2", "DB寫入失敗", err)
		return
	}

	return
}

func findBuyerAndSell(flowData *model.FlowData, controllerCode, serviceCode string) (status string, isOK bool) {

	req := flowData.Request.(model.Sell)
	remain := req.Quantity
	finish := false

	db := service.DbConnect()
	defer db.Close()

	for !finish {

		sql := "select id, quantity from Trade where type = 'buy' and price = ? and close = 0  limit 1;"
		row := db.QueryRow(sql, req.Price)
		id, quantity := -1, -1
		row.Scan(&id, &quantity)

		if quantity >= remain {
			break
		}

		if id == -1 {
			break
		}

	}

	return
}
