package api

import (
	"bumblebee/model"
	"fmt"
)

var responseCode = map[string]map[string]interface{}{
	"RETCODE_NOT_REGIST": {"RetCode": -1, "Message": "RetCode is not exists"},
	"SERVICE_ERROR":      {"RetCode": 130, "Message": "Service errormsg and break"},
	"MISSING_ACTION":     {"RetCode": 160, "Message": "MISSING_ACTION"},
	"MISSING_SIGNATURE":  {"RetCode": 170, "Message": "Missing signature"},
	"MISSING_PARAMS":     {"RetCode": 220, "Message": "Missing params [%s]"},
	"PARAMS_ERROR":       {"RetCode": 230, "Message": "Params [%s] not available"},

	"CREATE_USER_ERROR": {"RetCode": 4001, "Message": "check share bandwidth fail"},

	"CREATE_ERROR":   {"RetCode": 5016, "Message": "create errormsg [%s] "},
	"DESCRIBE_ERROR": {"RetCode": 5019, "Message": "describe errormsg [%s] "},

	"ERROR_MESSAGE": {"RetCode": 57000, "Message": "errormsg: %s"},
}

func generateErrorCode(name string, params ...interface{}) *model.APIResponse {
	res := new(model.APIResponse)
	res.RetCode = -1
	res.Message = "RetCode is not exists"
	errorCode, ok := responseCode[name]
	if !ok {
		return res
	}
	message, ok := errorCode["Message"].(string)
	if !ok {
		return res
	}
	res.RetCode = errorCode["RetCode"].(int)
	res.Message = fmt.Sprintf(message, params...)
	return res
}
