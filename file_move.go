package pcs
import (
  "net/url"
  "encoding/json"
)
type ResponseFileMove struct{
   From string `json:"from"`
   To string  `json:"to"`
}

func (rt *ResponseFileMove) String() string {
	bf, _ := json.Marshal(rt)
	return string(bf)
}

func (pcs *Pcs)FileMove(from string,to string)(*ResponseFileMove,error){
	var info ResponseFileMove
	url_values:=url.Values{}
   url_values.Add("from",from)
   url_values.Add("to",to)
	_, _, err := pcs.QuickRequest(pcs.BuildRequest(POST, "file?method=move&"+url_values.Encode(), nil), &info)
	return &info,err
}