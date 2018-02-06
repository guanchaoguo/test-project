package main

import (
	"os"
	"fmt"
	"log"
	"strings"
	"io/ioutil"
	"io"
	"path/filepath"
)



func main() {


	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	fmt.Println(slice2, slice1)

	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println(slice2, slice1)



}



func redderCloser(){
	s := strings.NewReader("hello world!")
	r := ioutil.NopCloser(s)
	r.Close()                                  //此处Close不起作用？！
	p := make([]byte, 10)
	r.Read(p)
	fmt.Println(string(p))   //hello worl

}



func devnull(){
	a := strings.NewReader("hello")
	p := make([]byte, 20)
	io.Copy(ioutil.Discard, a)
	ioutil.Discard.Write(p)
	fmt.Println(p)
}


func tempFile(){
	content := []byte("temporary file's content")

	tmpfile, err := ioutil.TempFile("", "example")

	if err != nil {

		log.Fatal(err)

	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {

		log.Fatal(err)

	}

	if err := tmpfile.Close(); err != nil {

		log.Fatal(err)

	}
}

func tempDir(){
	content := []byte("temporary file's content")

	dir, err := ioutil.TempDir("", "temp")
	if err != nil {

		log.Fatal(err)

	}

	defer os.RemoveAll(dir) // clean up


	tmpfn := filepath.Join(dir, "string.txt")

	fmt.Println(dir,tmpfn)

	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {

		log.Fatal(err)

	}
}

const GO_EOL = "\n"
func readDir(){
	rd,err:= ioutil.ReadDir(".")
	if err!= nil{
		fmt.Printf("%s read dir  fail\n")
	}
	for  k, fl := range rd{
		fmt.Println(k,GO_EOL,
			fl.Name(),GO_EOL,
			fl.IsDir(), GO_EOL,
			fl.ModTime(), GO_EOL,
			fl.Size(), GO_EOL,
			fl.Mode(), GO_EOL,
			fl.Sys(),GO_EOL)
	}
}

func fileReadWrite(){
	// 将字符串写入文件
	buf,err:= ioutil.ReadAll(strings.NewReader("123456"))
	if err!=nil{
		fmt.Printf("%s read buff  fail\n")
	}

	if ioutil.WriteFile("test.txt",buf,0666);err!=nil{
		fmt.Printf("%s write buff to file  fail\n")
	}

	//从文件读取数据再次写入
	buf,err = ioutil.ReadFile("test.txt")
	if err!=nil{
		fmt.Printf("%s read  file  fail\n")
	}
	buf = append(buf,[]byte("789")... )
	if ioutil.WriteFile("test.txt",buf,0666);err!=nil{
		fmt.Printf("%s write buff to file  fail\n")
	}

}

func fileExist ()  {


	path :="aa/bb"
	fi, err := os.Stat(path);
	// 目录是否存在
	if os.IsNotExist(err) {
		fmt.Printf("%s does not exist\n", path)

		// 目录不存在则创建
		if err:= os.MkdirAll(path,0666);err !=nil {
			fmt.Printf("%s create path  fail\n", path)
		}
	}else {
		// 判断是否为目录
		if fi.IsDir() {
			fmt.Printf("%s is a directory\n", path)
		}

		// 文件是否存在
		path :=path+"test.go"
		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Printf("%s does not exist\n", path)
		}
	}
}
