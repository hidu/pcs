package pcs

import (
	"bytes"
	"net/url"
	//	"fmt"
	"mime/multipart"
	"encoding/json"
	"io"
)

/**
*单个文件上传的
 */
type ResponseFileUploadSingle struct {
	Path  string `json:"path"`
	Size  uint64 `json:"size"`
	Ctime int64 `json:"ctime"`
	Mtime int64 `json:"mtime"`
	Md5   string `json:"md5"`
	Fs_id string `json:"fs_id "`
}

func (rt *ResponseFileUploadSingle) String() string {
	bf, _ := json.Marshal(rt)
	return string(bf)
}

func (pcs *Pcs) FileUploadSingle(data io.Reader, server_path string, ondup_overwrite bool) (resSingle *ResponseFileUploadSingle, err error) {
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
	
   resSingle=new(ResponseFileUploadSingle)
	_, _, err = pcs.QuickRequest(req, resSingle)
	return resSingle, err
}
