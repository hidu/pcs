package pcs

import (
  "net/url"
  "encoding/json"
  "fmt"
)

type ResponseFileList struct{
  List []ResponseFileListDetail `json:"list"`
  Request_id uint64 `json:"request_id"`
}
type ResponseFileListDetail struct{
   Fs_id uint64 `json:"fs_id"`
   Path string  `json:"path"`
   Ctime int64 `json:"ctime"`
	Mtime int64 `json:"mtime"`
   Md5 string  `json:"md5"`
   Size  uint64 `json:"size"`   //文件大小（byte）。 
   IsDir  uint64 `json:"isdir"`   //0为文件  1为目录 
}

func (rt *ResponseFileList) String() string {
	bf, _ := json.Marshal(rt)
	return string(bf)
}

const (
 FileList_OrderBy_MTime="time"  //修改时间
 FileList_OrderBy_Name="name"  //文件名
 FileList_OrderBy_Size="size"  //大小，注意目录无大小
 
 FileList_Order_Asc="asc"  //升序
 FileList_Order_Desc="desc"  //降序 缺省采用降序排序
)

/**
*@param path string
*@param orderby string 排序字段
*@param order string 排序方式
*@param offset int 结果集偏移量 从0开始
*@param limit int 结果集条数
*/
func (pcs *Pcs)FileList(path string,orderby string,order string,offset int,limit int)(*ResponseFileList,error){
	var info ResponseFileList
	url_values:=url.Values{}
   url_values.Add("path",path)
   if orderby!=""{
     if (orderby==FileList_OrderBy_MTime||orderby==FileList_OrderBy_Name||orderby==FileList_OrderBy_Size){
       url_values.Add("by",orderby)
     }else{
        panic(fmt.Sprintf("unknow order field [%s],only sup:[time|name|size]",orderby))
       }
   }
   if order!=""{
	  if(order==FileList_Order_Asc ||order==FileList_Order_Desc){
	     url_values.Add("order",order)
	   }else{
	     panic(fmt.Sprintf("unknow order method [%s],only sup:[asc|desc]",order))
	   }
   }
   if(limit>0 && offset>=0){
     url_values.Add("limit",fmt.Sprintf("%d-%d",offset,offset+limit))
   }
   
	_, _, err := pcs.QuickRequest(pcs.BuildRequest(GET, "file?method=list&"+url_values.Encode(), nil), &info)
	return &info,err
}

func (pcs *Pcs)FileListEasy(path string) (*ResponseFileList,error){
  return pcs.FileList(path,FileList_OrderBy_Name,FileList_Order_Desc,0,0)
}