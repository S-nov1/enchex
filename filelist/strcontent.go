package main

import (
   "fmt"
    "io/ioutil"
    "strings"
    "os"
)

func main() {

	files, err := ioutil.ReadDir(".")
	if err != nil {
		return
	}
	namfile:="";
	for _, file := range files {
		namfile=file.Name()
		if strings.Contains(namfile, ".txt") {
			fmt.Println(namfile)
			writehex(namfile)
		}
	}
 }
func writehex(namfile string) {

    content, err := ioutil.ReadFile(namfile)
    if err != nil {
        fmt.Printf("Can not open file "+namfile+"\n")
        return 
    }
    result :=strings.Replace(string(content), "\r\n", "\n", -1)
    lines := strings.Split(string(result), "\n")
       str := ""
    test:=""
    outstr:=""
    for lin :=0; lin < len(lines); lin++ {
        test=""
        str=lines[lin]
        outstr+=str+"   "
        for ix :=0; ix < len(lines[lin]); ix++ {
            test+=fmt.Sprintf("0x%x,",str[ix])
    }
    test+="0x00\n"
    outstr+=test
    }
   outname1:=namfile
   outname:=strings.Replace(outname1, ".txt", ".hxf", 1)
   f, err := os.OpenFile(outname, os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0)
   defer f.Close()
    f.WriteString(outstr)
 }