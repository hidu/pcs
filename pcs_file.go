package pcs

import (
	"bytes"
	"io/ioutil"
	"net/url"
	//	"fmt"
	"mime/multipart"
)

func (pcs *Pcs) FileUploadSingle(local_path string, file_data []byte, server_path string, ondup_overwrite bool) (resSingle *ResponseFileUploadSingle, err error) {
	if local_path != "" {
		file_data, err = ioutil.ReadFile(local_path)
		if err != nil {
			return
		}
	}
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
	fileWriter.Write(file_data)

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	req := pcs.BuildRequest(POST, "file?"+values.Encode(), buf)
	req.Header.Add("Content-Type", contentType)
	
   resSingle=new(ResponseFileUploadSingle)
	_, _, err = pcs.QuickRequest(req, resSingle)
	return resSingle, err

}
