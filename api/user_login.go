package api

import (
	"optimusprime/common"
	"optimusprime/log"
	"bumblebee/logic"
	"bumblebee/model"
	"bumblebee/util/request"
	"bumblebee/util/errormsg"
	"net/http"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {
	req := new(model.LoginRequest)
	if err := InputRequest(r, req); err != nil {
		OutputResponse(w, generateErrorCode("MISSING_PARAMS", err))
		return
	}
	_res := new(model.IJsCode2SessionResponse)
	err := request.LoginRequest(req.Code, _res)
	if err != nil {
		OutputResponse(w, generateErrorCode("MISSING_PARAMS", err))
		log.ERRORF("LoginRequest: %s", err)
	}
	log.DEBUGF("LoginResponse: %v", _res)
	res := new(model.LoginResponse)
	res.RetCode = 0
	userInfo := new(model.UserInfo)

	err = common.AESDecrypt2Obj(req.EncryptedData, _res.Session_key, req.Iv, &userInfo)
	if err != nil {
		log.ERRORF("AESDecrypt: %s", err)
	}
	_userInfo, err := logic.GetUser(userInfo.ResourceId)
	var userId int64
	if err == errormsg.USER_NOT_EXIST {
		userId, err = logic.AddUser(userInfo)
		if err != nil {
			OutputResponse(w, generateErrorCode("ERROR_MESSAGE", userId))
			return
		}
	}
	if _userInfo != nil {
		res.UserId = _userInfo.Id
		OutputResponse(w, res)
		return
	}
	res.UserId = userId
	OutputResponse(w, res)
	return
}
