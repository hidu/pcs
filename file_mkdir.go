package pcs

import (
	"encoding/json"
	"net/url"
)

type ResponseFileMakeDir struct {
	Fs_id uint64 `json:"fs_id"` //文件或目录在PCS的临时唯一标识ID。
	Path  string `json:"path"`
	Ctime int64  `json:"ctime"`
	Mtime int64  `json:"mtime"`
}

func (rt *ResponseFileMakeDir) String() string {
	bf, _ := json.Marshal(rt)
	return string(bf)
}

func (pcs *Pcs) FileMakeDir(path string) (info *ResponseFileMakeDir, pcs_err *PcsError) {
	url_values := url.Values{}
	url_values.Add("path", path)
	_, _, pcs_err = pcs.QuickRequest(pcs.BuildRequest(POST, "file?method=mkdir&"+url_values.Encode(), nil), &info)
	return info, pcs_err
}
