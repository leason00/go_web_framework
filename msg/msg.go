package msg

func SuccessRes(msg string, data map[string]interface{})(map[string]interface{}){
	if data == nil{
		return map[string]interface{}{"code": 0, "msg": msg}
	}
	return map[string]interface{}{"code": 0, "msg": msg, "data": data}
}

func ErrorRes(msg string, code int)(map[string]interface{}){
	return map[string]interface{}{"code": code, "msg": msg}
}

func ArrayRes(msg string, total int64, data []interface{})(map[string]interface{}){
	if data == nil{
		return map[string]interface{}{"code": 0, "msg": msg}
	}
	return map[string]interface{}{"code": 0, "msg": msg, "total":total, "data": data}
}