package main

import (
	"fmt"
	"os"
//	"bytes"
	"bufio"
	"io/ioutil"
	"strings"
	"flag"
	"github.com/hidu/pcs"
)

var method=flag.String("m","quote","run which metod")

func GetPcs() *pcs.Pcs{
   data,err:=ioutil.ReadFile("./token.txt")
   if err!=nil{
      fmt.Println("read ./token.txt failed!",err)
      os.Exit(1)
    }
	_pcs:= pcs.NewPcs(strings.TrimSpace(string(data)))
	_pcs.Debug=true;
	return _pcs;
}

var base_dir string="/apps/pcstest_oauth/"

func main(){
	flag.Parse()
   func_name:=fmt.Sprintf("%s",*method)
   funcs:=make(map[string]func())
   funcs["quote"]=run_quote
   funcs["upload"]=run_upload
   funcs["meta"]=run_meta
   funcs["meta_batch"]=run_metabatch
   funcs["makedir"]=run_makedir
   funcs["download"]=run_download
   funcs["filelist"]=run_filelist
   funcs["filecopy"]=run_filecopy
   funcs["filemove"]=run_filemv
   funcs["filedelete"]=run_filedelete
   funcs["filedeletebatch"]=run_filedeletebatch
   funcs["filesearch"]=run_filesearch
   if fun,has:=funcs[func_name];has{
     fun()
   }else if(func_name=="all"){
	   for name,fun:=range funcs{
	   fmt.Println("============",name,"===================")
	   fun()
	   fmt.Println("============",name,"===================\n")
	    }
   }else{
      fmt.Println("unknow method:",func_name)
   }
}

func run_quote(){
	pcs := GetPcs()
	quote, err := pcs.GetQuota()
	fmt.Println(quote,err)
}

func run_upload(){
	pcs := GetPcs()
	f,_:=os.Open("../pcs.go")
	obj, err := pcs.FileUploadSingle(f, base_dir+"pcs.go", true)
	fmt.Println(obj, err)
}
func run_meta(){
	pcs := GetPcs()
	obj, err := pcs.FileMeta(base_dir+"pcs.go")
	fmt.Println(obj, err)
}
func run_metabatch(){
	pcs := GetPcs()
	obj, err := pcs.FileMetaBatch([]string{base_dir+"pcs.go",base_dir+"yun.jpg"})
	fmt.Println(obj, err)
}
func run_makedir(){
	pcs := GetPcs()
	obj, err := pcs.FileMakeDir(base_dir+"test_dir")
	fmt.Println(obj, err)
}
func run_download(){
	pcs := GetPcs()
	w:=bufio.NewWriter(os.Stdout)
	obj,err := pcs.FileDownload(base_dir+"pcs.go",w)
//	w.Flush()
	fmt.Println(err)
	if(obj!=nil){
	fmt.Println(obj.ContentLength, )
	}
}

func run_filelist(){
	pcs := GetPcs()
	info,err:=pcs.FileListEasy(base_dir)
	fmt.Println(info, err)
}
func run_filecopy(){
	pcs := GetPcs()
	info,err:=pcs.FileCopy(base_dir+"pcs.go",base_dir+"pcs.go.copy1")
	fmt.Println(info, err)
}
func run_filemv(){
	pcs := GetPcs()
	info,err:=pcs.FileMove(base_dir+"pcs.go.copy1",base_dir+"pcs.go.copy2")
	fmt.Println(info, err)
}
func run_filedelete(){
	pcs := GetPcs()
	info,err:=pcs.FileDelete(base_dir+"pcs.go.copy2")
	fmt.Println(info, err)
}
func run_filedeletebatch(){
	pcs := GetPcs()
	info,err:=pcs.FileDeleteBatch([]string{base_dir+"pcs.go.copy2"})
	fmt.Println(info, err)
}
func run_filesearch(){
	pcs := GetPcs()
	info,err:=pcs.FileSearch(base_dir,"pcs",false)
	fmt.Println(info, err)
}