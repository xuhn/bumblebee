package request

import (
	"optimusprime/common"
	"optimusprime/log"
	"optimusprime/net"
	"bytes"
	"encoding/json"
)

func HttpPostRequest(url string, data map[string]interface{}, result interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	res, err := net.SendHTTPPostRequest(url, "application/json", bytes.NewBuffer(body), 10)
	if err != nil {
		return err
	}
	err = json.Unmarshal(res, result)
	if err != nil {
		return err
	}
	return nil
}

func HttpGetRequest(url string, data map[string]interface{}, result interface{}) error {
	log.DEBUGF("HttpGetRequest | %s \n    %+v", url, data)
	res, err := net.SendHTTPRequest(url, data, 10)
	if err != nil {
		return err
	}
	err = json.Unmarshal(res, result)
	if err != nil {
		return err
	}
	return nil
}

func PrivateAPIRequest(data map[string]interface{}, result interface{}) error {
	apiURL, err := common.GetConfigByKey("private_api")
	if err != nil {
		return err
	}

	if _, ok := data["request_uuid"]; !ok {
		data["request_uuid"] = common.NewUUIDV4().String()
	}

	log.INFOF(" PRIVATE_API_REQUEST (%s) | (%s.%s) \n    %+v", data["request_uuid"], data["Backend"], data["Action"], data)

	err = HttpPostRequest(apiURL.(string), data, result)

	log.INFOF(" PRIVATE_API_RESPONSE (%s) | (%s.%s) \n    %+v", data["request_uuid"], data["Backend"], data["Action"], result)

	return err
}

func LoginRequest(code string, result interface{}) error {
	url, err := common.GetConfigByKey("wxconfig.url")
	if err != nil {
		return err
	}
	appid, err := common.GetConfigByKey("wxconfig.appid")
	if err != nil {
		return err
	}
	secret, err := common.GetConfigByKey("wxconfig.secret")
	if err != nil {
		return err
	}
	grant_type, err := common.GetConfigByKey("wxconfig.grant_type")
	if err != nil {
		return err
	}
	params := map[string]interface{}{
		"appid":      appid,
		"secret":     secret,
		"js_code":    code,
		"grant_type": grant_type,
	}
	err = HttpGetRequest(url.(string), params, result)
	if err != nil {
		log.ERRORF("HttpGetRequest: %s", err)
	}
	return nil
}
