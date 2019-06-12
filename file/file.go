package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//文件读
func main() {

	inputfile,inputerr := os.Open("nql.dat")
	if inputerr!=nil{
		fmt.Println(inputerr)
		return
	}
	defer inputfile.Close()

	inputreder := bufio.NewReader(inputfile)

	for {
		inputstring,inputstringerr:= inputreder.ReadString('\n')
		fmt.Println(inputstring)
		if inputstringerr == io.EOF{
			break
		}
	}

	outputfile,outputerr := os.OpenFile("nql.dat",os.O_WRONLY|os.O_CREATE,0666)

	if outputerr!=nil {
		fmt.Println(outputerr)
		return
	}

	defer outputfile.Close()

	oupputwriter := bufio.NewWriter(outputfile)
	outstring := "Hello World \n"

	for i:=0;i<100 ;i++  {
		oupputwriter.WriteString(outstring)
	}

	oupputwriter.Flush()

}
