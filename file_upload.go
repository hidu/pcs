package pcs

import (
	"bytes"
	"net/url"
	"fmt"
	"mime/multipart"
	"encoding/json"
	"io"
)

/**
*单个文件上传的
 */
type ResponseFileUpload struct {
	Path  string `json:"path"`
	Size  uint64 `json:"size"`
	Ctime int64 `json:"ctime"`
	Mtime int64 `json:"mtime"`
	Md5   string `json:"md5"`
	Fs_id string `json:"fs_id "`
}
/**
*文件分片
 */
type ResponseFileUploadSlice struct {
	Md5   string `json:"md5"`
	Fs_id string `json:"fs_id "`
}

func (rt *ResponseFileUpload) String() string {
	bf, _ := json.Marshal(rt)
	return string(bf)
}

func (pcs *Pcs) FileUploadSingle(data io.Reader, server_path string, ondup_overwrite bool) (resSingle *ResponseFileUpload, err error) {
	ondup := "overwrite"
	if !ondup_overwrite {
		ondup = "newcopy"
	}
	values := url.Values{
		"method": {"upload"},
		"path":   {server_path},
		"ondup":  {ondup},
	}
	buf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(buf)
	fileWriter, err := bodyWriter.CreateFormFile("file", "hidu")
	io.Copy(fileWriter,data)
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	
	req := pcs.BuildRequest(POST, "file?"+values.Encode(), buf)
	req.Header.Add("Content-Type", contentType)
	
   resSingle=new(ResponseFileUpload)
	_, _, err = pcs.QuickRequest(req, resSingle)
	return resSingle, err
}

//分片上传，2G以内
func (pcs *Pcs)FileUploadSlice(data io.Reader)(resSlice *ResponseFileUploadSlice,err error){
	buf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(buf)
	fileWriter, err := bodyWriter.CreateFormFile("file", "hidu")
	io.Copy(fileWriter,data)
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	
	req := pcs.BuildRequest(POST, "file?method=upload&type=tmpfile", buf)
	req.Header.Add("Content-Type", contentType)
	resSlice=new(ResponseFileUploadSlice)
	_, _, err = pcs.QuickRequest(req, resSlice)
	return resSlice,err
}

// 分片上传—合并分片文件 
func (pcs *Pcs)FileUploadSliceMerge(block_list []string,server_path string,ondup_overwrite bool)(resInfo *ResponseFileUpload,err error){
  block_size:=len(block_list)
  if(block_size>1024){
     err=fmt.Errorf("slice size out of range [%d > 1024]",block_size)
     return
   }else if (block_size<2){
     err=fmt.Errorf("min slice is 2.now is [%d]",block_size)
     return
   }
  for i,md5_str:=range block_list{
    if(len(md5_str)!=32){
       err=fmt.Errorf("the %d's md5 str [%s] length is not 32",i,md5_str)
       return
     }
   }
   param_map:=make(map[string][]string)
   param_map["param"]=block_list
   var param_byte []byte
   param_byte,err=json.Marshal(param_map)
	if err!=nil{
	   return
	}
   
   ondup := "overwrite"
	if !ondup_overwrite {
		ondup = "newcopy"
	}
	values := url.Values{
		"method": {"createsuperfile"},
		"path":   {server_path},
		"ondup":  {ondup},
	}
	
	post_values:=url.Values{"param":{string(param_byte)}}
	req := pcs.BuildRequest(POST, "file?"+values.Encode(), bytes.NewBufferString(post_values.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencode")
	resInfo=new(ResponseFileUpload)
	_, _, err = pcs.QuickRequest(req, resInfo)
	return
}
