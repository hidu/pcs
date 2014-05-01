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