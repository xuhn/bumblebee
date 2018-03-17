package api

import (
	"optimusprime/log"
	"bumblebee/logic"
	"bumblebee/model"
	"net/http"
)

func ListItem(w http.ResponseWriter, r *http.Request) {
	req := new(model.ListItemInfoRequest)
	if err := InputRequest(r, req); err != nil {
		OutputResponse(w, generateErrorCode("MISSING_PARAMS", err))
		return
	}

	res := new(model.ListItemInfoResponse)
	res.RetCode = 0

	_, errKey := logic.ListItem(req.AccountId)

	if errKey != "" {
		log.ERRORF("", errKey)
		OutputResponse(w, generateErrorCode(errKey))
		return
	}
	/*
		// 处理返回消息
		for _, info := range infos {
			res.DataSet = append(res.DataSet, map[string]interface{}{
				"ItemId":   info,
				"ItemName": info.ItemName,
			})
		}
	*/

	OutputResponse(w, res)
	return
}
