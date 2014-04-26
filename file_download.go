package pcs

import (
  "bytes"
   "net/url"
   "fmt"
    "time"
    "net/http"
)

type FileDownloadInfo struct{
    Buffer bytes.Buffer
    ContentLength int64
    Content_MD5 string
    Content_Type string
    Mtime uint64
}

func (pcs *Pcs)FileDownload(path string)(*FileDownloadInfo,error){
  url_values:=url.Values{}
  url_values.Add("path",path)
  req:=pcs.BuildRequest(GET, "file?method=download&"+url_values.Encode(), nil)
  res,err:=pcs.http_client.Do(req)
  if err!=nil{
   return nil,err
  }
  info:=new(FileDownloadInfo)
  info.Buffer=bytes.Buffer{}
  info.Buffer.ReadFrom(res.Body)

  info.ContentLength=res.ContentLength
  info.Content_Type=res.Header.Get("Content-Type")
  info.Content_MD5=res.Header.Get("Content-MD5")
  //@todo
//  Last_Modified:=res.Header.Get("Last-Modified")
//  t,_:=time.Parse(http.TimeFormat,Last_Modified)
//  fmt.Print(t)
  return info,err
}
