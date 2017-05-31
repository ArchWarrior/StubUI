package main

import (
"os/exec"
	"fmt"
	"strings"
)

var state  bool

func getServerState() bool{
//	command:=exec.Command("cmd", "cd")
//	fmt.Println(command.Output())
	command:=exec.Command("tasklist.exe", "/fo", "csv", "/nh")
//command:= exec.Command("tasklist", "/fi", "\"imagename eq notepad.exe\"")
if cmdoutput, err :=  command.Output(); err != nil { 
        fmt.Println("Error: ", err)
       // state=false
        //ch<-state
    } else{      //fmt.Println("Mocky state: ", string(cmdoutput))
                 processlist:=string(cmdoutput)
                      if strings.Contains(processlist,"Virtualizer.exe") {
                         state=true
                       }else{state=false}
          }
    return state
}

func toggleServerState(ch chan<- bool){
//	pwd,_:=os.Getwd()
//	fmt.Println(pwd)
state=getServerState()
    
if state==false{
	state=true
	ch <- state
	command:= exec.Command("Virtualizer.exe")
	if cmdoutput, err1 :=  command.Output(); err1 != nil { 
        fmt.Println("Error: ", err1)
       // state=false
        //ch<-state
    } else{fmt.Println("Mocky log: ", string(cmdoutput))}
	state =false
	ch <- state
}else{
	command:= exec.Command("taskkill", "/IM", "Virtualizer.exe", "/F")
    if cmdoutput, err2 := command.Output(); err2 != nil { 
        fmt.Println("Error: ", err2)
    }else{fmt.Println("Mocky log: ", string(cmdoutput))}
    state =false
	ch <- state
}

}

	
	 

    