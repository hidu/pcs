package pcs
import (
  "net/url"
  "encoding/json"
)
type ResponseFileCopy struct{
   From string `json:"from"`
   To string  `json:"to"`
}

func (rt *ResponseFileCopy) String() string {
	bf, _ := json.Marshal(rt)
	return string(bf)
}

func (pcs *Pcs)FileCopy(from string,to string)(*ResponseFileCopy,error){
	var info ResponseFileCopy
	url_values:=url.Values{}
   url_values.Add("from",from)
   url_values.Add("to",to)
	_, _, err := pcs.QuickRequest(pcs.BuildRequest(POST, "file?method=copy&"+url_values.Encode(), nil), &info)
	return &info,err
}