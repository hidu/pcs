package pcs
import (
	"net/http"
	"fmt"
	"io"
	"io/ioutil"
   "bytes"
)
func PrintRequestInfo(req *http.Request){
  fmt.Println(`==request start=============================================`)
  format:="%10s : %s\n"
  fmt.Printf(format,"url",req.URL)
  fmt.Printf(format,"method",req.Method)
  req.ParseForm()
  fmt.Printf("%10s\n","params")
  for k,v:=range req.Form{
  	fmt.Printf("%10s %5s : %s\n","",k,v)
  }
  fmt.Println(`==request end===============================================`)
  fmt.Println()
}

func PrintResponseInfo(res *http.Response,err error){
   fmt.Println(`==response start===============================================`)
  format:="%10s : %v\n"
  if(err!=nil){
	  fmt.Printf(format,"err",err)
  }
  if(res!=nil){
   fmt.Printf(format,"status",res.Status)
   fmt.Printf(format,"req_url",res.Request.URL)
   fmt.Printf(format,"req_method",res.Request.Method)
   content_type:=res.Header.Get("Content-Type")
   if(content_type[:4]=="text"){
     buf:=bytes.NewBuffer([]byte{})
     io.Copy(buf,res.Body)
     fmt.Printf(format,"body",buf.String())
     res.Body=ioutil.NopCloser(buf).(io.ReadCloser)
    }
   fmt.Printf("%10s\n","header")
   for k,v:=range res.Header{
      fmt.Printf("%2s %30s : %s\n","",k,v)
   }
  }
   fmt.Println(`==response end===============================================`)
  fmt.Println()
}