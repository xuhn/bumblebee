package api

import (
	"github.com/xuhn/optimusprime/log"
	"bumblebee/logic"
	"bumblebee/model"
	"net/http"
)

func GetItem(w http.ResponseWriter, r *http.Request) {
	req := new(model.GetItemInfoRequest)
	if err := InputRequest(r, req); err != nil {
		OutputResponse(w, generateErrorCode("MISSING_PARAMS", err))
		return
	}

	res := new(model.ListItemInfoResponse)
	res.RetCode = 0

	_, errKey := logic.GetItem(req.ItemId)
	if errKey != "" {
		log.ERRORF("", errKey)
		OutputResponse(w, generateErrorCode(errKey))
		return
	}

	/*


		res.DataSet = append(res.DataSet, map[string]interface{}{
			"ItemId":    info.ItemId,
			"ItemName":  info.ItemName,
			"AccountId": info.AccountId,
			"Url":       info.Url,
		})

	*/
	OutputResponse(w, res)
	return
}
