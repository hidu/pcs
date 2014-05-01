package pcs
import (
  "net/url"
)

func (pcs *Pcs)FileDelete(path string)(*ResponseOk,error){
	var info ResponseOk
	url_values:=url.Values{}
   url_values.Add("path",path)
	_, _, err := pcs.QuickRequest(pcs.BuildRequest(POST, "file?method=delete&"+url_values.Encode(), nil), &info)
	return &info,err
}
//批量删除接口
func (pcs *Pcs)FileDeleteBatch(paths []string)(*ResponseOk,error){
	var info ResponseOk
   param_str,err:=paths_param_build(paths)
   if(err!=nil){
     return nil,err
    }
   url_values:=url.Values{}
   url_values.Add("param",param_str)
	_, _, err = pcs.QuickRequest(pcs.BuildRequest(POST, "file?method=delete&"+url_values.Encode(), nil), &info)
	return &info,err
}

