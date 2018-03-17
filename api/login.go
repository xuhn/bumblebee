package api

import (
	"optimusprime/controller"
	"bumblebee/logic"
	"bumblebee/model"
	"bumblebee/util/errormsg"
	"bumblebee/util/request"
	"optimusprime/log"
	"optimusprime/common"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"time"
)


type Passport struct {
	*controller.Controller
}

func (c Passport) Login() controller.Result {
	_v := c.Params.Values
	code := _v.Get("code")
	encryptedData := _v.Get("encryptedData")
	iv := _v.Get("iv")

	_res := new(model.IJsCode2SessionResponse)
	err := request.LoginRequest(code, _res)
	if err != nil {
		log.ERRORF("LoginRequest: %s", err)
		return c.RenderError(err)
	}
	log.DEBUGF("LoginResponse: %v", _res)

	userInfo := new(model.UserInfo)
	err = common.AESDecrypt2Obj(encryptedData, _res.Session_key, iv, &userInfo)
	if err != nil {
		log.ERRORF("AESDecrypt: %s", err)
	}
	_userInfo, err := logic.GetUser(userInfo.ResourceId)
	var userId int64
	if err == errormsg.USER_NOT_EXIST {
		userId, err = logic.AddUser(userInfo)
		if err != nil {
			return c.RenderError(err)
		}
	}

	//生成token
	//初始化私钥
	signBytes, err := ioutil.ReadFile(controller.PrivKeyPath)
	if err != nil {
		log.ERRORF("No http.privatekey provided: %s", err)
		return c.RenderError(err)
	}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.ERRORF("No http.signkey provided: %s", err)
		return c.RenderError(err)
	}

	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	log.DEBUGF("token: %v", token)
	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return c.RenderError(err)
	}
	data := make(map[string]interface{})
	if _userInfo != nil {
		data["UserId"] = userId
	}
	data["Token"] = tokenString
	log.DEBUGF("Login: %v", data)
	return c.RenderJSON(data)
}