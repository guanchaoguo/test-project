package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	FileName   string   = "我有一只小毛驴.txt" //这是我们需要打开的文件，当然你也可以把它定义到从某个配置文件来获取变量。
	InputFile  *os.File                 //变量 InputFile 是 *os.File 类型的。该类型是一个结构，表示一个打开文件的描述符（文件句柄）。
	InputError error                    //我们使用 os 包里的 Open 函数来打开一个文件。如果文件不存在或者程序没有足够的权限打开这个文件，Open函数会返回一个错误，InputError变量就是用来接收这个错误的。
	Count      int                      //这个变量是我们用来统计行号的，默认值为0.
)

func main() {

}

func readfile() {
	//InputFile,InputError = os.OpenFile(FileName,os.O_CREATE|os.O_RDWR,0644) //打开FileName文件，如果不存在就创建新文件，打开的权限是可读可写，权限是644。这种打开方式相对下面的打开方式权限会更大一些。
	InputFile, InputError = os.Open(FileName) //使用 os 包里的 Open 函数来打开一个文件。该函数的参数是文件名，类型为 string 。我们以只读模式打开"FileName"文件。
	if InputError != nil {                    //如果打开文件出错，那么我们可以给用户一些提示，然后在推出函数。
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer InputFile.Close()                   //defer关键字是用在程序即将结束时执行的代码,确保在程序退出前关闭该文件。
	inputReader := bufio.NewReader(InputFile) //我们使用 bufio.NewReader()函数来获得一个读取器变量（读取器）。我们可以很方便的操作相对高层的 string 对象，而避免了去操作比较底层的字节。
	for {
		Count += 1
		inputString, readerError := inputReader.ReadString('\n') //我们将inputReader里面的字符串按行进行读取。
		if readerError == io.EOF {
			return //如果遇到错误就终止循环。
		}
		fmt.Printf("The %d line is: %s", Count, inputString) //将文件的内容逐行（行结束符'\n'）读取出来。
	}
}

func file_seek(){
	f, err := os.OpenFile("a.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("yinzhengjie\n")
	//表示文件的其实位置，从第二个字符往后写入。
	f.Seek(1, 0)

	f.WriteString("$$$")
	f.Close()
	os.Remove("a.txt")
}

func romveDir(){

}
