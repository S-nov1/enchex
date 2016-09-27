package main

import (
   "fmt"
    "io/ioutil"
    "strings"
    "os"
)

func main() {

    inpArgs := os.Args
    count:=len(inpArgs)
    if(count>1){
        fmt.Printf(inpArgs[1]+"\n")
    } else {
        fmt.Printf("No input file!\n")
        return 
    }
    content, err := ioutil.ReadFile(inpArgs[1])
    if err != nil {
        fmt.Printf("Can not open file "+inpArgs[1]+"\n")
        return 
    }
    lines := strings.Split(string(content), "\n")
    str := ""
    test:=""
    outstr:=""
    for lin :=0; lin < len(lines); lin++ {
        test=""
        str=lines[lin]
        outstr+=str+"   "
       // fmt.Printf(lines[lin]+"\n")
        for ix :=0; ix < len(lines[lin]); ix++ {
            test+=fmt.Sprintf("0x%x,",str[ix])
    }
    test+="0x00\n"
    outstr+=test
    }
    outname:="out_"
    if(count>2){
        outname=inpArgs[2]
    } else {
       outname+=inpArgs[1] 
    }
    f, err := os.OpenFile(outname, os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0)
    defer f.Close()
    f.WriteString(outstr)
}