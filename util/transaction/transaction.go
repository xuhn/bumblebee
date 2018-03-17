package transaction

import (
	"github.com/xuhn/optimusprime/common"
	"github.com/xuhn/optimusprime/persistence/mysql"
)

func Begin() (string, error) {
	// 获取数据库连接
	dbConnStr, err := common.GetConfigByKey("dbconfig.database")
	if err != nil {
		return "", err
	}
	dbConn, err := mysql.GetMysqlInstance(dbConnStr.(string))
	if err != nil {
		return "", err
	}
	txId, err := dbConn.BeginTx()
	if err != nil {
		return "", err
	}
	return txId, nil
}

func Commit(txId string) error {
	dbConnStr, err := common.GetConfigByKey("dbconfig.database")
	if err != nil {
		return err
	}
	dbConn, err := mysql.GetMysqlInstance(dbConnStr.(string))
	if err != nil {
		return err
	}
	err = dbConn.Commit(txId)
	if err != nil {
		return err
	}
	return nil
}

func Rollback(txId string) error {
	dbConnStr, err := common.GetConfigByKey("dbconfig.database")
	if err != nil {
		return err
	}
	dbConn, err := mysql.GetMysqlInstance(dbConnStr.(string))
	if err != nil {
		return err
	}
	err = dbConn.Rollback(txId)
	if err != nil {
		return err
	}
	return nil
}
