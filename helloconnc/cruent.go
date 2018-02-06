package main

import (
	"fmt"
	"sort"
	"io"
	"encoding/binary"
	"math/rand"
	"os"
	"bufio"
	"net"
	"strconv"
)

const fileNameIn = "same.in"
const fileNameOut = "same.out"
const countSum = 50
const chunCount = 4
const fileSize  = 512

func main() {
	SouertNeteDemo()
}

func SouertNeteDemo(){
	p:= CreateNetPiple(fileNameIn,fileSize,chunCount)
	WriterToFile(p,fileNameOut)
	PrintFile(fileNameOut)
}


func SouertFileDemo(){
	p:= CreatePiple(fileNameIn,fileSize,chunCount)
	WriterToFile(p,fileNameOut)
	PrintFile(fileNameOut)
}

func NetworkSink(addr string,in <- chan int){
	listener,err:= net.Listen("tcp",addr)
	if err != nil{
		panic(err)
	}

	go func() {
		defer listener.Close()
		conn,err:=  listener.Accept()
		if err !=nil{
			panic(err)
		}
		defer conn.Close()

		writer:= bufio.NewWriter(conn)
		defer writer.Flush()
		WriterSink(writer,in)
	}()
}

func NetSourece(addr string) <- chan int{
	out := make(chan int)
	go func() {
		conn,err:=  net.Dial("tcp",addr)
		if err !=nil{
			panic(err)
		}
		r:= ReadSource(conn,-1)
		for v:= range r{
			out<- v
		}
		close(out)
	}()
	return out
}


func CreateNetPiple(fileName string,fileSize,chunkCount int) <-chan int{
	chunkSize := fileSize/chunkCount
	sourtAddr:= []string {}
	for i:=0;i<chunkCount ;i++  {
		file,err  := os.Open(fileName)
		if err !=nil{
			panic(err)
		}
		file.Seek( int64(i*chunkSize),0)

		addr:= "0.0.0.0:"+ strconv.Itoa(8000+i)
		source:= ReadSource(bufio.NewReader(file),chunkSize)
		sourtsource := InMerSort(source)
		NetworkSink(addr ,sourtsource)
		sourtAddr = append(sourtAddr,addr)
	}

	resultSoure := []<-chan int{}
	for _,v:=  range sourtAddr{
		resultSoure = append(resultSoure,NetSourece(v))
	}

	return  MergerN(resultSoure...)
}

func PrintFile(fileName string){
	file ,err:= os.Open(fileName)
	if err !=nil{
		panic(err)
	}
	defer file.Close()

	pile:= ReadSource(file,fileSize)
	for v :=range pile  {
		fmt.Println(v)
	}
}

 func WriterToFile(ins <-chan int,fileName string){
 	 file ,err:= os.Create(fileName)
	 if err !=nil{
		 panic(err)
	 }
	 defer file.Close()

	  writer:= bufio.NewWriter(file)
	  defer writer.Flush()

	  WriterSink(writer,ins)
 }

func CreatePiple(fileName string,fileSize,chunkCount int) <-chan int{
	chunSize := fileSize/chunkCount
	resultSoure:= []<-chan int{}
	for i:=0;i<chunkCount ;i++  {
		file,err  := os.Open(fileName)
		if err !=nil{
			panic(err)
		}
		file.Seek( int64(i*chunSize),0)
		source:= ReadSource(bufio.NewReader(file),chunSize)
		resultSoure = append(resultSoure,  InMerSort(source))
	}

	return  MergerN(resultSoure...)
}

func CreateRandomIo(){
	const filename = "samlle.in"
	const n  = 64
	file,err:=  os.Create(filename)
	if err !=nil{
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	p:= RandomSource(n)
	WriterSink(writer,p )
	writer.Flush()


	file,err = os.Open(filename)
	if err !=nil{
		panic(err)
	}
	defer file.Close()

	reader:= bufio.NewReader(file)
	p =  ReadSource(reader,-1)
	for v := range p{
		fmt.Println(v)
	}

}


func MergerDemo(){
	arr1 := InMerSort(arraySort(1, 4, 8, 9))
	arr2 := InMerSort(arraySort(2, 7, 3, 5))
	p := Merge(arr1, arr2)

	for val := range p {
		fmt.Println(val)
	}
}

func arraySort(arr ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, v := range arr {
			ch <- v
		}
		close(ch)
	}()

	return ch
}

func InMerSort(in <-chan int) <-chan int {
	out := make(chan int)

	// 接收数据
	go func() {
		arr := []int{}
		for val := range in {
			arr = append(arr, val)
		}

		// 排序数据
		sort.Ints(arr)

		//重新发送排序后的数据
		for _, v := range arr {
			out <- v
			fmt.Println(v)
		}
		close(out)
	}()

	return out
}

func Merge(in1, in2 <-chan int) <-chan int {
	ch := make(chan int)
	// 接收来自两个节点已经排序后的数据
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				ch <- v1
				v1, ok1 = <-in1
			} else {
				ch <- v2
				v2, ok2 = <-in2
			}
		}
		close(ch)
	}()

	return ch
}

func ReadSource( reader io.Reader,chunkSize int)<-chan int{
	out := make(chan int)

	go func() {
		buffer := make([]byte ,8)
		byteRead := 0
		for{
			n,err := reader.Read(buffer)
			if n > 0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out<-v
			}
			if err != nil|| ( chunkSize != -1 && byteRead >= chunkSize){
				break
			}
			byteRead++
		}
		close(out)
	}()

	return out
}

func WriterSink( writer io.Writer, in <-chan int){
	for v := range in {
		buffer:= make([]byte,8)
		fmt.Println(v)
		binary.BigEndian.PutUint64(buffer,uint64(v))
		writer.Write(buffer)
	}
}

func RandomSource(count int) <-chan int {
	out := make(chan int)
	go func() {
		for i:= 1; i<= count ; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}


func MergerN( ins ...<- chan int)<- chan int{
	if len(ins) == 1{
		return  ins[0]
	}
	mid:= len(ins)>> 1

	return Merge(
		MergerN(ins[:mid]...) ,
		MergerN(ins[:mid]...))
}