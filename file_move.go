package pcs
import (
  "net/url"
)

func (pcs *Pcs)FileMove(from string,to string)(info *ResponseFileCopy,pcs_err *PcsError){
	url_values:=url.Values{}
   url_values.Add("from",from)
   url_values.Add("to",to)
	_, _, pcs_err = pcs.QuickRequest(pcs.BuildRequest(POST, "file?method=move&"+url_values.Encode(), nil), &info)
	return info,pcs_err
}