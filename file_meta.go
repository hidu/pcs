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
   Ctime uint64 `json:"ctime"`
	Mtime uint64 `json:"mtime"`
	Block_list    string `json:"block_list"`  //文件所有分片的md5数组JSON字符串。 
	Size  uint64 `json:"size"`
	Isdir  uint64 `json:"isdir"`  //“0”为文件 “1”为目录 
	Filenum  uint64 `json:"filenum"` 
	Ifhassubdir  uint64 `json:"ifhassubdir"` //是否含有子目录的标识符 “0”表示没有子目录  “1”表示有子目录 
}

func (rt *ResponseFileMeta) String() string {
	bf, _ := json.Marshal(rt)
	return string(bf)
}

//批量获取文件或目录的元信息
func (pcs *Pcs)FileMeta(path string)(*ResponseFileMeta,error){
  var meta ResponseFileMeta
  url_values:=url.Values{}
  url_values.Add("path",path)
  _, _, err := pcs.QuickRequest(pcs.BuildRequest(GET, "file?method=meta&"+url_values.Encode(), nil), &meta)
  return &meta,err
}

func (pcs *Pcs)FileMetaBatch(paths []string)(*ResponseFileMeta,error){
  var meta ResponseFileMeta
  param_map:=make(map[string][]map[string]string)
  param_map["list"]=[]map[string]string{}
  for _,path:=range paths{
     m:=make(map[string]string)
     m["path"]=path
    param_map["list"]=append(param_map["list"],m)
  }
  url_values:=url.Values{}
  param_byte,err:=json.Marshal(param_map)
  if err!=nil{
     return nil,err
  }
  url_values.Add("param",string(param_byte))
  _, _, err = pcs.QuickRequest(pcs.BuildRequest(GET, "file?method=meta&"+url_values.Encode(), nil), &meta)
  return &meta,err
}

