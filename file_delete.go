package pcs
import (
  "net/url"
)

func (pcs *Pcs)FileDelete(path string)(info *ResponseOk,pcs_err *PcsError){
	url_values:=url.Values{}
   url_values.Add("path",path)
	_, _, pcs_err = pcs.QuickRequest(pcs.BuildRequest(POST, "file?method=delete&"+url_values.Encode(), nil), &info)
	return info,pcs_err
}
//批量删除接口
func (pcs *Pcs)FileDeleteBatch(paths []string)(info *ResponseOk,pcs_err *PcsError){
   param_str,err:=paths_param_build(paths)
   if(err!=nil){
  	  pcs_err=NewPcsError(ERROR_OTHER,"FileDeleteBatch params build error:"+err.Error())
     return nil,pcs_err
    }
   url_values:=url.Values{}
   url_values.Add("param",param_str)
	_, _, pcs_err = pcs.QuickRequest(pcs.BuildRequest(POST, "file?method=delete&"+url_values.Encode(), nil), &info)
	return info,pcs_err
}

