package pcs

import (
   "net/url"
)

/**
*@param path string
*@param query string  关键词
*@param recursive string 是否递归
*/
func (pcs *Pcs)FileSearch(path string,query string,recursive bool)(*ResponseFileMeta,error){
   var info ResponseFileMeta
	url_values:=url.Values{}
	url_values.Add("path",path)
	if(query==""){
	  panic("file search with empty query")
	}
	url_values.Add("wd",query)
	if(recursive){
		url_values.Add("re","1")
	}else{
		url_values.Add("re","0")
	}
	_, _, err := pcs.QuickRequest(pcs.BuildRequest(GET, "file?method=search&"+url_values.Encode(), nil), &info)
	return &info,err
}