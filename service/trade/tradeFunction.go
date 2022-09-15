package trade

import (
	"database/sql"
	"fmt"
	"tradeEngine/model"
	"tradeEngine/service"
)

//findBuyerAndSell functionCode:S1
func findBuyerAndSell(flowData *model.FlowData, controllerCode, serviceCode string) {
	req := flowData.Request.(model.Sell)
	res := model.SellResponse{}
	remain := req.Quantity
	var buyer []int

	db := service.DbConnect()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	for {
		query := "select id, quantity from Trade where type = 'buy' and price = ? and close = 0 order by id limit 1;"
		row := db.QueryRow(query, req.Price)
		id, quantity := -1, -1
		err := row.Scan(&id, &quantity)
		if err != nil {
			service.SetError(flowData, controllerCode, serviceCode, "S1", "DB搜尋失敗", err)
			return
		}

		if id == -1 {
			query = "insert into Trade(type, quantity, price, close, updateTime) values ('sell', ?, ?, 0, now());"
			_, err = db.Exec(query, remain, req.Price)
			if err != nil {
				service.SetError(flowData, controllerCode, serviceCode, "S1", "DB寫入失敗", err)
				return
			}
			res.Message = fmt.Sprintf("完成 %d 個，剩餘 %d 個掛單中", req.Quantity-remain, remain)
			res.BuyerId = buyer
			break
		} else if quantity-remain > 0 {
			buyer = append(buyer, id)
			query = "update Trade set quantity = ?, updateTime = now() where id = ?;"
			_, err = db.Exec(query, quantity-remain, id)
			if err != nil {
				service.SetError(flowData, controllerCode, serviceCode, "S1", "DB更新失敗", err)
				return
			}
			res.Message = "交易完成"
			res.BuyerId = buyer
			break
		} else if quantity-remain == 0 {
			buyer = append(buyer, id)
			query = "update Trade set quantity = 0, close = 1, updateTime = now() where id = ?;"
			_, err = db.Exec(query, id)
			if err != nil {
				service.SetError(flowData, controllerCode, serviceCode, "S1", "DB更新失敗", err)
				return
			}
			res.Message = "交易完成"
			res.BuyerId = buyer
			break
		} else {
			buyer = append(buyer, id)
			query = "update Trade set quantity = 0, close = 1, updateTime = now() where id = ?;"
			_, err = db.Exec(query, id)
			if err != nil {
				service.SetError(flowData, controllerCode, serviceCode, "S1", "DB更新失敗", err)
				return
			}
			remain -= quantity
		}

	}

	flowData.Response = res
	return
}

//findSellerAndBuy functionCode:S2
func findSellerAndBuy(flowData *model.FlowData, controllerCode, serviceCode string) {
	req := flowData.Request.(model.Buy)
	res := model.BuyResponse{}
	remain := req.Quantity
	var seller []int

	db := service.DbConnect()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			service.SetError(flowData, controllerCode, serviceCode, "S2", "DB關閉失敗", err)
			return
		}
	}(db)

	for {
		query := "select id, quantity from Trade where type = 'sell' and price = ? and close = 0 order by id limit 1;"
		row := db.QueryRow(query, req.Price)
		id, quantity := -1, -1
		err := row.Scan(&id, &quantity)
		if err != nil {
			service.SetError(flowData, controllerCode, serviceCode, "S2", "DB遭料存取失敗", err)
			return
		}

		if id == -1 {
			query = "insert into Trade(type, quantity, price, close, updateTime) values ('buy', ?, ?, 0, now());"
			_, err := db.Exec(query, remain, req.Price)
			if err != nil {
				service.SetError(flowData, controllerCode, serviceCode, "S2", "DB寫入失敗", err)
				return
			}
			res.Message = fmt.Sprintf("完成 %d 個，剩餘 %d 個掛單中", req.Quantity-remain, remain)
			res.SellerId = seller
			break
		} else if quantity-remain > 0 {
			seller = append(seller, id)
			query = "update Trade set quantity = ?, updateTime = now() where id = ?;"
			_, err := db.Exec(query, quantity-remain, id)
			if err != nil {
				service.SetError(flowData, controllerCode, serviceCode, "S2", "DB更新失敗", err)
				return
			}
			res.Message = "交易完成"
			res.SellerId = seller
			break
		} else if quantity-remain == 0 {
			seller = append(seller, id)
			query = "update Trade set quantity = 0, close = 1, updateTime = now() where id = ?;"
			_, err := db.Exec(query, id)
			if err != nil {
				service.SetError(flowData, controllerCode, serviceCode, "S2", "DB更新失敗", err)
				return
			}
			res.Message = "交易完成"
			res.SellerId = seller
			break
		} else {
			seller = append(seller, id)
			query = "update Trade set quantity = 0, close = 1, updateTime = now() where id = ?;"
			_, err := db.Exec(query, id)
			if err != nil {
				service.SetError(flowData, controllerCode, serviceCode, "S2", "DB更新失敗", err)
				return
			}
			remain -= quantity
		}
	}

	flowData.Response = res
	return
}

//searchTrade functionCode:S3
func searchTrade(flowData *model.FlowData, controllerCode, serviceCode string) {
	req := flowData.Request.(model.Search)
	var arguments []interface{}

	where := ""
	if req.Id > 0 {
		where += " and id = ?"
		arguments = append(arguments, req.Id)
	}
	if req.Type != "" {
		where += " and type = ?"
		arguments = append(arguments, req.Type)
	}
	if req.Quantity != -1 {
		where += " and quantity = ?"
		arguments = append(arguments, req.Quantity)
	}
	if req.Price != -1 {
		where += " and price = ?"
		arguments = append(arguments, req.Price)
	}
	if req.Close != -1 {
		where += " and close = ?"
		arguments = append(arguments, req.Close)
	}
	if req.UpdateTimeStart != "" {
		where += " and updateTime >= ?"
		arguments = append(arguments, req.UpdateTimeStart)
	}
	if req.UpdateTimeEnd != "" {
		where += " and updateTime <= ?"
		arguments = append(arguments, req.UpdateTimeEnd)
	}
	if len(where) != 0 {
		where = " where " + where[5:]
	}

	query := "select id, type, quantity, price, close, updateTime from Trade" + where + ";"

	flowData.Response = query
}
