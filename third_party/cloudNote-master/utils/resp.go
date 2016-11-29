package utils

//Response 客户端应答
type Response struct {
	//Code 错误码
	Code    int         `json:"Code"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
}

const (
	SYSTEM_CODE                  = 1000
	CATEGORYID_WRONG_CODE        = 1001
	CATEGORYID_WRONG_MESSAGE     = "该笔记分组不存在"
	WITHOUT_PARAMETER_CODE       = 1002
	WITHOUT_PARAMETER_MESSAGE    = "参数不全"
	NOTE_NOT_EXIST_CODE          = 1003
	NOTE_NOT_EXIST               = "笔记不存在"
	NOTE_ALREADY_EXIST_CODE      = 1004
	NOTE_ALREADY_EXIST           = "该笔记或者附件已经存在，不能重复创建"
	NOTE_ALREADY_DELETE_CODE     = 1005
	NOTE_ALREADY_DELETE          = "该笔记已经被删除"
	CATEGORY_ALREADY_DELETE_CODE = 1006
	CATEGORY_ALREADY_DELETE      = "该笔记分组已经被删除"

	Base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	UserId      = "test"
)

//NewResponse 初始化response
func NewResponse(Code int, Msg string, data interface{}) *Response {
	if Code != 0 {
		data = nil
	}
	return &Response{
		Code:    Code,
		Message: Msg,
		Data:    data,
	}
}

func InvalidParameter() *Response {
	return NewResponse(WITHOUT_PARAMETER_CODE, WITHOUT_PARAMETER_MESSAGE, nil)
}
