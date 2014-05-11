package pcs

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
)

type ResponseQuote struct {
	QuotaSize int64 `json:"quota"`
	UsedSize  uint64 `json:"used"`
	RequestId uint64 `json:"request_id"`
}
func (rt *ResponseQuote) String() string {
	bf, _ := json.Marshal(rt)
	return string(bf)
}


type Pcs struct {
	accessToken string
	http_client *http.Client
	Debug bool
}

func NewPcs(accessToken string) *Pcs {
	return &Pcs{accessToken:accessToken, http_client:&http.Client{}}
}

func (pcs *Pcs) BuildRequest(method string, params string, body io.Reader) *http.Request {
	url := PCS_HOST + params + "&access_token=" + pcs.accessToken
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Println(err)
		return nil
	}
	return req
}

func (pcs *Pcs)DoRequest(req *http.Request)(resp *http.Response,err error){
	if(pcs.Debug){
		PrintRequestInfo(req)
	}
   resp,err=pcs.http_client.Do(req)
   if(pcs.Debug){
   	PrintResponseInfo(resp,err)
    }
   return
}

func (pcs *Pcs) QuickRequest(req *http.Request, v interface{}) (resp *http.Response, resp_body []byte, pcs_err *PcsError) {
 	var err error
	resp, err = pcs.DoRequest(req)
	resp_body,pcs_err=parseResponse(resp,err)
	if v != nil {
		err = json.Unmarshal(resp_body, &v)
		if err != nil {
			pcs_err=NewPcsError(ERROR_OTHER,fmt.Sprintf("json decode response failed.status [%d],response [%s],err:[%s]",resp.StatusCode,string(resp_body),err.Error()))
			return
		}
	}
	return
}

func parseResponse(resp *http.Response,respErr error)(resp_body []byte,pcs_err *PcsError) {
	if respErr != nil {
		pcs_err=NewPcsError(ERROR_OTHER,respErr.Error())
		return
	}
 	var err error
	resp_body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
	   pcs_err=NewPcsError(ERROR_OTHER,"read response body failed:"+err.Error())
		return
	}
	if(resp.StatusCode!=200){
		err=json.Unmarshal(resp_body, &pcs_err)
		if(err!=nil){
			pcs_err=NewPcsError(ERROR_OTHER,fmt.Sprintf("json decode response failed.status [%d],response [%s],err:[%s]",resp.StatusCode,string(resp_body),err.Error()))
		}
		return
	}
	return
}



type ResponseOk struct {
	RequestId uint64 `json:"request_id"`
}
func (rt *ResponseOk) String() string {
	bf, _ := json.Marshal(rt)
	return string(bf)
}

