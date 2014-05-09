package pcs

import (
  "net/url"
  "encoding/json"
)

//单个文件或目录的元信息
type ResponseFileMeta struct{
   List  []ResponseFileMeta_SubInfo `json:"list"`
	Request_id  uint64 `json:"request_id"` //是否含有子目录的标识符 “0”表示没有子目录  “1”表示有子目录 
}

type ResponseFileMeta_SubInfo struct{
	Fs_id   uint64 `json:"fs_id"`   //文件或目录在PCS的临时唯一标识ID。 
	Path   string `json:"path"`
   Ctime int64 `json:"ctime"`
	Mtime int64 `json:"mtime"`
	Block_list    string `json:"block_list"`  //文件所有分片的md5数组JSON字符串。 
	Size  uint64 `json:"size"`
	Isdir  int `json:"isdir"`  //“0”为文件 “1”为目录 
	Filenum  int `json:"filenum"` 
	Ifhassubdir  int `json:"ifhassubdir"` //是否含有子目录的标识符 “0”表示没有子目录  “1”表示有子目录 
}


func (rt *ResponseFileMeta) String() string {
	bf, _ := json.Marshal(rt)
	return string(bf)
}

//批量获取文件或目录的元信息
func (pcs *Pcs)FileMeta(path string)(*ResponseFileMeta,*PcsError){
  var meta ResponseFileMeta
  url_values:=url.Values{}
  url_values.Add("path",path)
  _, _, err := pcs.QuickRequest(pcs.BuildRequest(GET, "file?method=meta&"+url_values.Encode(), nil), &meta)
  return &meta,err
}

func (pcs *Pcs)FileMetaBatch(paths []string)(meta *ResponseFileMeta,pcs_err *PcsError){
  param_str,err:=paths_param_build(paths)
  if(err!=nil){
  	 pcs_err.Error_msg=err.Error()
    return nil,pcs_err
  }
  url_values:=url.Values{}
  url_values.Add("param",param_str)
  _, _, pcs_err = pcs.QuickRequest(pcs.BuildRequest(GET, "file?method=meta&"+url_values.Encode(), nil), &meta)
  return meta,pcs_err
}

func aaa(){
}
