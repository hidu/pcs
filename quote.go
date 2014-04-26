package pcs
import (
//   "fmt"
   "encoding/json"
)

type Quote struct{
   QuotaSize int64 `json:"quota"`
   UsedSize int64 `json:"used"`
   RequestId int64 `json:"request_id"`
}

func (quote *Quote)String() string{
   bf, _ := json.Marshal(quote)
   return string(bf)
}

func (pcs *Pcs) GetQuota() (*ResponseQuote, error) {
	var quote ResponseQuote
	_, _, err := pcs.QuickRequest(pcs.BuildRequest(GET, "quota?method=info", nil), &quote)
	return &quote, err
}