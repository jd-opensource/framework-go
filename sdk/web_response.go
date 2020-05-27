package sdk

/*
 * Author: imuge
 * Date: 2020/5/27 下午6:49
 */

type WebResponse struct {
	Success bool         `json:"success"`
	Data    interface{}  `json:"data"`
	Error   ErrorMessage `json:"error"`
}

type ErrorMessage struct {
	ErrorCode    int32  `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}
