package model

//=============================================
// 消息结构
//=============================================
type (
	LoginRequest struct {
		Code          string `key:"code" required:"true"`
		EncryptedData string `key:"encryptedData" required:"true"`
		Iv            string `key:"iv" required:"true"`
	}

	LoginResponse struct {
		APIResponse
		ThirdSessionId string
		UserId         int64
	}

	IJsCode2SessionResponse struct {
		Session_key string
		Expires_in  int
		Openid      string
	}
)
