package main

import (
	bytes2 "bytes"
	"fmt"
	"os"
)

func main(){
	var bytes bytes2.Buffer
	bytes.Write([]byte("hello"))

	fmt.Fprintf(&bytes,"world")
	bytes.WriteTo(os.Stdout)
}
