package msg

type ControllerError struct {
	Code     int    `json:"code"`
	Message  string `json:"msg"`
}

var (
		Err404 			= &ControllerError{404, "page not found"}
		ErrNoUser 		= &ControllerError{400, "用户不存在！"}
		ErrPass 		= &ControllerError{400, "密码有误！"}
		ErrPath 		= &ControllerError{500, "文件保存失败！"}
)




func SuccessRes(msg string, data map[string]interface{})(map[string]interface{}){
	if data == nil{
		return map[string]interface{}{"code": 200, "msg": msg}
	}
	return map[string]interface{}{"code": 200, "msg": msg, "data": data}
}


func ArrayRes(msg string, total int64, data []interface{})(map[string]interface{}){
	if data == nil{
		return map[string]interface{}{"code": 200, "msg": msg}
	}
	return map[string]interface{}{"code": 200, "msg": msg, "total":total, "data": data}
}