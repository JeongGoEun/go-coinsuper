// Package types of both request and response for API
package types

// ResponseData structure
type ResponseData struct {
	Timestamp int64       `json:"timestamp"`
	Result    interface{} `json:"result"`
}

// Response structure
type Response struct {
	Code string        `json:"code"`
	Msg  string        `json:"msg"`
	Data *ResponseData `json:"data"`
}

// RequestCommon structure
type RequestCommon struct {
	AccessKey string `json:"accesskey"`
	Timestamp string `json:"timestamp,omitempty"`
	Sign      string `json:"sign"`
}

// Request structure
type Request struct {
	Common *RequestCommon `json:"common"`
	Data   *Options       `json:"data"`
}

// Asset structure
type Asset struct {
	Available string `json:"available"`
	Total     string `json:"total"`
}

// UserInfo structure
type UserInfo struct {
	UserNo int64             `json:"userNo"`
	Email  string            `json:"email"`
	Assets map[string]*Asset `json:"asset"`
}
