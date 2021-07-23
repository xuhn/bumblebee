package dao

import (
	"strconv"
	_ "strconv"

	"github.com/xuhn/bumblebee/model"
	"github.com/xuhn/bumblebee/util/errormsg"

	"github.com/xuhn/optimusprime/common"
	"github.com/xuhn/optimusprime/persistence/mysql"
)

func GetUserInfo(user_id string) (*model.UserInfo, error) {
	// 获取数据库连接
	dbConnStr, err := common.GetConfigByKey("dbconfig.database")
	if err != nil {
		return nil, err
	}
	dbConn, err := mysql.GetMysqlInstance(dbConnStr.(string))
	if err != nil {
		return nil, err
	}
	sqlStr := "SELECT * FROM t_user_info WHERE resource_id=?"
	rows, err := dbConn.Select(sqlStr, user_id)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, errormsg.USER_NOT_EXIST
	}
	row := rows[0]
	_gender, _ := strconv.Atoi(row["gender"])
	_id, _ := strconv.ParseInt(row["id"], 10, 64)
	userInfo := &model.UserInfo{
		UserName:   row["name"],
		ResourceId: row["resource_id"],
		Gender:     _gender,
		AvatarUrl:  row["avatar_url"],
		Id:         _id,
	}
	return userInfo, nil
}

func AddUserInfo(txId string, userInfo *model.UserInfo) (int64, error) {
	// 获取数据库连接
	dbConnStr, err := common.GetConfigByKey("dbconfig.database")
	if err != nil {
		return 0, err
	}
	dbConn, err := mysql.GetMysqlInstance(dbConnStr.(string))
	if err != nil {
		return 0, err
	}

	sqlStr := "INSERT INTO t_user_info(resource_id, name, gender, avatar_url) VALUES(?,?,?,?)"
	sqlArgs := []interface{}{userInfo.ResourceId, userInfo.UserName, userInfo.Gender, userInfo.AvatarUrl}

	_id, err := dbConn.TxInsert(txId, sqlStr, sqlArgs...)
	if err != nil {
		return 0, err
	}
	return _id, nil
}

const (
	Host       = "23.91.98.209"
	Username   = "system"
	Password   = "1qaz@WSX#EDC"
	Database   = "test"
	Collection = "wechat"
)

/*
func AddUserInfoMongo(userInfo *model.UserInfo) (string, errormsg) {
	// 获取数据库连接

	dbConnStr, err := common.GetConfigByKey("dbconfig.mongo")
	if err != nil {
		return "", err
	}

	dbConn, err := mongodb.GetMgoInstance(dbConnStr.(string))
	if err != nil {
		return "", err
	}

	userInfo.Id = common.NewUUIDV4().String()
	data := make([]interface{}, 0)
	data = append(data, userInfo)
	err = dbConn.InsertDocs(Database, Collection, data)

	if err != nil {
		return "", err
	}

	return userInfo.Id, nil
}
*/
