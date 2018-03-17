package model

//=============================================
// 公用结构
//=============================================
type (
	/*
		UserInfo struct {
			Id       string `json:"_id"`
			UserName string
			Password string
			Phone    string
			//ResourceId string
		}
	*/
	UserInfo struct {
		Id         int64  `json:"_id"`
		UserName   string `json:"nickName"`
		Password   string `json:"password"`
		Phone      string `json:"phone"`
		ResourceId string `json:"openId"`
		Gender     int    `json:"gender"`
		AvatarUrl  string `json:"avatarUrl"`
	}
)

//=============================================
// 消息结构
//=============================================
type (
	// 添加用户
	AddUserRequest struct {
		UserName string `key:"UserName" required:"true"`
		Password string `key:"Password"`
		Phone    string `key:"Phone"`
	}
	AddUserResponse struct {
		APIResponse
		UserId int64
	}
	// 获取用户
	GetUserRequest struct {
		UserId string `key:"UserId"`
	}
	GetUserResponse struct {
		APIResponse
		UserName string
		Password string
		Phone    string
	}
)
