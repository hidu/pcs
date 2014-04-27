package pcs

import (
   "net/url"
	"io"
	"time"
//	"fmt"
	"net/http"
)

type FileDownloadInfo struct{
    ContentLength int64
    Content_MD5 string
    Content_Type string
    Mtime int64
}
//文件下载
func (pcs *Pcs)FileDownload(path string,writer io.Writer)(*FileDownloadInfo,error){
  url_values:=url.Values{}
  url_values.Add("path",path)
  req:=pcs.BuildRequest(GET, "file?method=download&"+url_values.Encode(), nil)
  res,err:=pcs.http_client.Do(req)
  if err!=nil{
   return nil,err
  }
  info:=new(FileDownloadInfo)

  info.ContentLength=res.ContentLength
  info.Content_Type=res.Header.Get("Content-Type")
  info.Content_MD5=res.Header.Get("Content-MD5")
  io.Copy(writer,res.Body)
  
  Last_Modified:=res.Header.Get("Last-Modified")
  t,_:=time.Parse(http.TimeFormat,Last_Modified)
  info.Mtime=t.Local().Unix()
//  fmt.Println(t.Local().Format("2006-01-02 15:04:05"))
  return info,err
}
