package logic

import (
	"optimusprime/log"
	"bumblebee/dao"
	"bumblebee/model"
)

func GetItem(itemId string) (*model.ItemInfo, string) {
	itemInfo, err := dao.GetItemInfo(itemId)
	if err != nil {
		log.ERRORF("GetItem_Error: %s", err)
		return nil, "DESCRIBE_ERROR"
	}
	return itemInfo, ""
}

func ListItem(accountId int) ([]*model.ItemInfo, string) {
	itemInfos, err := dao.ListItemInfo(accountId)
	if err != nil {
		log.ERRORF("ListItem: %s", err)
		return nil, "DESCRIBE_ERROR"
	}
	return itemInfos, ""
}
