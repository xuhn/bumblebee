package model

//=============================================
// 公用结构
//=============================================
type (
	ItemInfo struct {
		Id        string `json:"_id"`
		ItemId    string
		AccountId int
		ItemName  string
		Url       string
	}
)

//=============================================
// 消息结构
//=============================================
type (
	// Get Item
	GetItemInfoRequest struct {
		ItemId string `key:"ItemId"`
	}

	// List Item
	ListItemInfoRequest struct {
		AccountId int `key:"AccountId"`
	}

	ListItemInfoResponse struct {
		APIResponse
	}
)
