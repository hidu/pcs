package pcs

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
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
}

func NewPcs(accessToken string) *Pcs {
	return &Pcs{accessToken, &http.Client{}}
}

func (pcs *Pcs) BuildRequest(method string, params string, body io.Reader) *http.Request {
	url := PCS_HOST + params + "&access_token=" + pcs.accessToken
	log.Println("url:", url, "\n")
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Println(err)
		return nil
	}
	return req
}
func (pcs *Pcs) QuickRequest(req *http.Request, v interface{}) (resp *http.Response, resp_body []byte, pcs_err *PcsError) {
 	var err error
	resp, err = pcs.http_client.Do(req)
	if err != nil {
		pcs_err.Error_msg=err.Error()
		return
	}
	resp_body, err = ioutil.ReadAll(resp.Body)
	log.Println("resp_code", resp.StatusCode)
	log.Println("resp_body", string(resp_body))
	if err != nil {
		return
	}
	if(resp.StatusCode!=200){
		err=json.Unmarshal(resp_body, &pcs_err)
		return
	}
	if v != nil {
		err = json.Unmarshal(resp_body, &v)
		if err != nil {
			log.Println("json unmarshar failed,response:", string(resp_body), "err:", err)
			return
		}
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

