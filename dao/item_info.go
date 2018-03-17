package dao

import (
	"fmt"
	"strconv"

	"optimusprime/common"
	"optimusprime/persistence/mysql"
	"bumblebee/model"
)

func GetItemInfo(item_id string) (*model.ItemInfo, error) {
	// 获取数据库连接
	dbConnStr, err := common.GetConfigByKey("dbconfig.database")
	if err != nil {
		return nil, err
	}
	dbConn, err := mysql.GetMysqlInstance(dbConnStr.(string))
	if err != nil {
		return nil, err
	}

	sqlStr := "SELECT a.account_id, a.item_id, a.item_name, b.url FROM t_item_info a, t_item_detail b WHERE a.item_id=b.item_id AND a.item_id=?"
	rows, err := dbConn.Select(sqlStr, item_id)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, fmt.Errorf("item_info not exist")
	}

	row := rows[0]

	account_id, _ := strconv.Atoi(row["account_id"])
	itemInfo := &model.ItemInfo{
		"", row["item_id"], account_id, row["item_name"], row["url"],
	}

	return itemInfo, nil
}

func ListItemInfo(account_id int) ([]*model.ItemInfo, error) {
	// 获取数据库连接
	dbConnStr, err := common.GetConfigByKey("dbconfig.database")
	if err != nil {
		return nil, err
	}
	dbConn, err := mysql.GetMysqlInstance(dbConnStr.(string))
	if err != nil {
		return nil, err
	}

	sqlStr := "SELECT item_id, item_name FROM t_item_info WHERE account_id=?"
	rows, err := dbConn.Select(sqlStr, account_id)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, fmt.Errorf("item_info not exist")
	}

	itemInfos := make([]*model.ItemInfo, 0)

	for _, row := range rows {
		itemInfos = append(itemInfos, &model.ItemInfo{
			ItemId:   row["item_id"],
			ItemName: row["item_name"],
		})
	}
	return itemInfos, nil
}
