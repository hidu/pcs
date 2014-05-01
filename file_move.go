package pcs
import (
  "net/url"
)

func (pcs *Pcs)FileMove(from string,to string)(*ResponseFileCopy,error){
	var info ResponseFileCopy
	url_values:=url.Values{}
   url_values.Add("from",from)
   url_values.Add("to",to)
	_, _, err := pcs.QuickRequest(pcs.BuildRequest(POST, "file?method=move&"+url_values.Encode(), nil), &info)
	return &info,err
}