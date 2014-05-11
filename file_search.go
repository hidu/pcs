package pcs

import (
   "net/url"
)

/**
*@param path string
*@param query string  关键词
*@param recursive string 是否递归
*/
func (pcs *Pcs)FileSearch(path string,query string,recursive bool)(info *ResponseFileMeta,pcs_err *PcsError){
	url_values:=url.Values{}
	url_values.Add("path",path)
	if(query==""){
		pcs_err=NewPcsError(ERROR_CUSTOM,"file search with empty query")
	   return nil,pcs_err
	}
	url_values.Add("wd",query)
	if(recursive){
		url_values.Add("re","1")
	}else{
		url_values.Add("re","0")
	}
	_, _, pcs_err = pcs.QuickRequest(pcs.BuildRequest(GET, "file?method=search&"+url_values.Encode(), nil), &info)
	return info,pcs_err
}