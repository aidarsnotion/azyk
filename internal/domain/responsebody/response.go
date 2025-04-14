package responsebody

type ResponseBody struct {
	Id      string      `json:"id"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
