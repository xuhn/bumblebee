package logic

import (
	"github.com/xuhn/bumblebee/dao"
	"github.com/xuhn/bumblebee/model"

	"github.com/xuhn/optimusprime/log"
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
