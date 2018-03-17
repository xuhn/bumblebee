package model

//=============================================
// 消息结构
//=============================================
type (
	// API请求通用结构
	APIRequest struct {
		Backend     string `json:"Backend"`      // 服务后端
		Action      string `json:"Action"`       // 请求名字
		RequestUUID string `json:"request_uuid"` // 请求UUID
	}
	APIResponse struct {
		RetCode int    `json:"RetCode"` // 状态码 0:成功
		Message string `json:"Message"` // 错误消息
	}
)
