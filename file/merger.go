package main

import (
	"os"
	"io"
	"io/ioutil"
	"strconv"
	"fmt"
)


func main(){

	mergerFile("file.txt" ,crateFile(3))
}


// 合并文件
func mergerFile(fileName string ,file_num int){
	//fl, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE, 0644)
	fl, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer fl.Close()

	for i:= 1;i<=file_num ;i++  {
		fileNmae := strconv.Itoa(i)+".txt"
		buff,err:=	readFile(fileNmae)
		if err != nil {
			panic(err)
		}
		n, err := fl.Write(buff)
		if err == nil && n < len(buff) {
			err = io.ErrShortWrite
		}
		fmt.Println("read file",i)
		err = os.Remove(fileNmae)
		if  err != nil{
			fmt.Println(err)
		}
	}
}

// 读取文件
func readFile(filePth string)([]byte, error){
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ioutil.ReadAll(f)
}

//创建三个文件 并写入  1 2 3
// 读取文件
func crateFile(file_num int) (num int) {
	for i:= 1;i<=file_num ;i++  {
		fileNmae:= strconv.Itoa(i)+".txt"
		ioutil.WriteFile(fileNmae,[]byte(fileNmae),0666)
		num = i
	}
	return
}

