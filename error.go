package pcs
import(
	"encoding/json"
)

type PcsError struct{
 Error_code int `json:"error_code"`
 Error_msg string `json:"error_msg"`
 Request_id int `json:"request_id"`
 Extra map[string]interface{} `json:"extra"`
}

func (e *PcsError)Error() string {
	bf, _ := json.Marshal(e)
	return string(bf)
}

func (e *PcsError)JsonDecodeError(err error){
   e.Error_code=123
   e.Error_msg=err.Error()
}
func NewPcsError(err_code int,msg string) *PcsError{
 	return &PcsError{Error_code:err_code,Error_msg:msg,Request_id:-1}
}

const (
   ERROR_CUSTOM		=8000			 //一般用户操作错误
   ERROR_CUSTOM_JSON =8001			 //json解析错误
)