package pcs
import (
  "net/url"
  "encoding/json"
)
type ResponseFileCopy struct{
   Extra ResponseFileCopyList `json:"extra"`
   Request_id uint64 `json:"request_id"`
}
type ResponseFileCopyList struct{
  List []ResponseFileCopyDetail `json:"list"`
}
type ResponseFileCopyDetail struct{
   From string `json:"from"`
   To string  `json:"to"`
}

func (rt *ResponseFileCopy) String() string {
	bf, _ := json.Marshal(rt)
	return string(bf)
}

func (pcs *Pcs)FileCopy(from string,to string)(info *ResponseFileCopy,pcs_err *PcsError){
	url_values:=url.Values{}
   url_values.Add("from",from)
   url_values.Add("to",to)
	_, _, pcs_err = pcs.QuickRequest(pcs.BuildRequest(POST, "file?method=copy&"+url_values.Encode(), nil), &info)
	return info,pcs_err
}