package pcs

import (
  "net/url"
  "encoding/json"
)

type ResponseFileMakeDir struct{
  	Fs_id   uint64 `json:"fs_id"`   //文件或目录在PCS的临时唯一标识ID。
  	Path   string `json:"path"`
  	Ctime uint64 `json:"ctime"`
	Mtime uint64 `json:"mtime"`
}
func (rt *ResponseFileMakeDir) String() string {
	bf, _ := json.Marshal(rt)
	return string(bf)
}

func (pcs *Pcs)FileMakeDir(path string)(*ResponseFileMakeDir,error){
	var info ResponseFileMakeDir
	url_values:=url.Values{}
   url_values.Add("path",path)
	_, _, err := pcs.QuickRequest(pcs.BuildRequest(POST, "file?method=mkdir&"+url_values.Encode(), nil), &info)
	return &info,err
}