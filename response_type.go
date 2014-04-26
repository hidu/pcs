package pcs

import (
	//   "fmt"
	"encoding/json"
)


type ResponseQuote struct {
	QuotaSize int64 `json:"quota"`
	UsedSize  int64 `json:"used"`
	RequestId int64 `json:"request_id"`
}
func (rt *ResponseQuote) String() string {
	bf, _ := json.Marshal(rt)
	return string(bf)
}


/**
*单个文件上传的
 */
type ResponseFileUploadSingle struct {
	Path  string `json:"path"`
	Size  uint64 `json:"size"`
	Ctime uint64 `json:"ctime"`
	Mtime uint64 `json:"mtime"`
	Md5   string `json:"md5"`
	Fs_id string `json:"fs_id "`
}
func (rt *ResponseFileUploadSingle) String() string {
	bf, _ := json.Marshal(rt)
	return string(bf)
}