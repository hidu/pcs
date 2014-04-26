package pcs

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

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
func (pcs *Pcs) QuickRequest(req *http.Request, v interface{}) (resp *http.Response, resp_body []byte, err error) {
	resp, err = pcs.http_client.Do(req)
	if err != nil {
		return
	}
	resp_body, err = ioutil.ReadAll(resp.Body)
	log.Println("resp_body", string(resp_body))
	if err != nil {
		return
	}
	if err != nil {
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

func (pcs *Pcs) GetQuota() (*ResponseQuote, error) {
	var quote ResponseQuote
	_, _, err := pcs.QuickRequest(pcs.BuildRequest(GET, "quota?method=info", nil), &quote)
	return &quote, err
}
