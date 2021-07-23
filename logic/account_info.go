package logic

import (
	"github.com/xuhn/bumblebee/dao"
	"github.com/xuhn/bumblebee/model"
	"github.com/xuhn/bumblebee/util/errormsg"
	"github.com/xuhn/bumblebee/util/transaction"

	"github.com/xuhn/optimusprime/log"
)

func AddUser(userInfo *model.UserInfo) (int64, error) {
	txId, err := transaction.Begin()
	if err != nil {
		return 0, errormsg.CREATE_USER_ERROR
	}
	defer transaction.Rollback(txId)
	userId, err := dao.AddUserInfo(txId, userInfo)
	if err != nil {
		log.ERRORF("AddUserInfo_Error: %s", err)
		return 0, errormsg.CREATE_USER_ERROR
	}
	log.DEBUGF("AddUserInfo_user_id: %d, tx_id: %v", userId, txId)
	transaction.Commit(txId)
	return userId, nil
}

func GetUser(userName string) (*model.UserInfo, error) {
	userInfo, err := dao.GetUserInfo(userName)
	if err == errormsg.USER_NOT_EXIST {
		log.ERRORF("GetUser not exist: %s", err)
		return nil, err
	}
	if err != nil {
		log.ERRORF("GetUser_Error: %s", err)
		return nil, errormsg.SERVICE_ERROR
	}
	return userInfo, nil
}
