package main

import (
  "net/http"
  "strconv"
  "github.com/gin-gonic/gin"
  "fmt"
//  "os/exec"
//  "os"
//  "gopkg.in/mgo.v2/bson"
//  "strings"
  
  
)

func welcome(c *gin.Context) {
	c.HTML(
    // Set the HTTP status to 200 (OK)
    http.StatusOK,
    // Use the index.html template
    "welcome.html",
    // Pass the data that the page uses
    gin.H{
      "title":   "EasyMock",
     },
  )
}
func showIndexPage(c *gin.Context) {
	
  services := getAllServices()

  // Call the HTML method of the Context to render a template
  c.HTML(
    // Set the HTTP status to 200 (OK)
    http.StatusOK,
    // Use the index.html template
    "index.html",
    // Pass the data that the page uses
    gin.H{
      "title":   "View Services",
      "payload": services,
    },
  )

}

func getService(c *gin.Context) {
  // Check if the service ID is valid
   serviceID:= c.Param("service_id")
   if len(serviceID)!=0{
    // Check if the service exists
    if service, err := getServiceByID(serviceID); err == nil {
      // Call the HTML method of the Context to render a template
      c.HTML(
        // Set the HTTP status to 200 (OK)
        http.StatusOK,
        // Use the index.html template
        "service.html",
        // Pass the data that the page uses
        gin.H{
          "title":   service.Sname,
          "payload":service,
        },
      )

    } else {
      // If the article is not found, abort with an error
      c.AbortWithError(http.StatusNotFound, err)
    }

  } else {
    // If an invalid article ID is specified in the URL, abort with an error
    c.AbortWithStatus(http.StatusNotFound)
  }
}
func addService(c *gin.Context){
//	var file, err = os.OpenFile("D:/Users/ggwala001c/workspace/Mockymodules/config.toml", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
//	if err!=nil{panic(err)}
//	defer file.Close()
//if err := toml.NewEncoder(file).Encode(config); err != nil {
//    panic(err)
//}
//fmt.Println(buf.String())
	 c.HTML(
    // Set the HTTP status to 200 (OK)
    http.StatusOK,
    // Use the index.html template
    "AddNewService.html",
    // Pass the data that the page uses
    gin.H{
      "title":   "Add Service",
      
    },
  )
}
func saveService(c *gin.Context){
	
	//message := c.PostForm("InputEndpoint")
	//fmt.Println("inside handeler")
	
	ServiceForm:=new(Service)
	var op Operation 
	var output Output 
   Path:=c.PostForm("InputEndpoint")
   Sname:= c.PostForm("InputServiceName")
   Type:=c.PostForm("responseType")
    //ServiceForm.Operations[0]=new(Operation)
    Opname:= c.PostForm("InputOperationName")
         Delay, _ := strconv.Atoi(c.PostForm("InputDelay"))
         Tagvalue:= c.PostForm("InputTag")
    Response:= c.PostForm("response")
    
//    ServiceForm.Operation
output.Tagvalue=Tagvalue
output.Response=Response
op.Opname=Opname
op.Delay=Delay
op.Outputs=[]Output{output}
ServiceForm.Path=Path
ServiceForm.Sname=Sname
ServiceForm.Type=Type
ServiceForm.Operations=[]Operation{op}


//
//s := &Service{bson.NewObjectId(),
//	Sname,
//	Path,
//	[1]Operation[
//	             {Operation{
//	             		      Opname,
//	             		      Delay,
//	             		      [1]Output[
//	             		                {Tagvalue,
//	             		                  Response
//	             		                }
//	             		               ]
//	                        }
//	             }
//	           ]}
////        c.Bind(&ServiceForm) 
       //fmt.Println(ServiceForm)
       err:=saveServiceinDB(ServiceForm)
       if err !=nil{
      panic(err)
       }
       c.HTML(
    // Set the HTTP status to 200 (OK)
    http.StatusOK,
    // Use the index.html template
    "welcome.html",
    // Pass the data that the page uses
    gin.H{
      "title":   "EasyMock",
     },
  )
}

func server(c *gin.Context) {
	var button string
	serverstate:=getServerState()
	if(serverstate==true){
		button="Stop Server"
	}else{button="Start Server"}
	c.HTML(
    // Set the HTTP status to 200 (OK)
    http.StatusOK,
    // Use the index.html template
    "server.html",
    // Pass the data that the page uses
    gin.H{
      "title":   "Server",
      "payload": button,
     },
  )
}

func StartStopServer(c *gin.Context) {
//	pwd,_:=os.Getwd()
//	fmt.Println(pwd)
var button string
     ch := make(chan bool)
         go toggleServerState(ch)
         <-ch 
         fmt.Println("toggled: ", state)
         serverstate:=getServerState()
	if(serverstate==true){
		button="Stop Server"
	    }else{button="Start Server"}
//	 command:= exec.Command("mocky.exe")
//
//    if cmdoutput, err1 :=  command.Output(); err1 != nil { 
//        fmt.Println("Error: ", err1)
//    } else{fmt.Println("Mocky log: ", string(cmdoutput))}
    
//    command= exec.Command("taskkill", "/IM", "mocky.exe", "/F")
//    if cmdoutput, err2 := command.Output(); err2 != nil { 
//        fmt.Println("Error: ", err2)
//    }else{fmt.Println("Mocky log: ", string(cmdoutput))}
//      
	
	
	c.HTML(
    // Set the HTTP status to 200 (OK)
    http.StatusOK,
    // Use the index.html template
    "server.html",
    // Pass the data that the page uses
    gin.H{
      "title":   "Server",
      "payload": button,
     },
  )
}

//func stopServer(c *gin.Context) {
////	pwd,_:=os.Getwd()
////	fmt.Println(pwd)
////	 command:= exec.Command("mocky.exe")
//
////    if cmdoutput, err1 := command.Output(); err1 != nil { 
////        fmt.Println("Error: ", err1)
////    } else{fmt.Println("Mocky log: ", string(cmdoutput))}
//    
//    command:= exec.Command("taskkill", "/IM", "mocky.exe", "/F")
//    if cmdoutput, err2 := command.Output(); err2 != nil { 
//        fmt.Println("Error: ", err2)
//    }else{fmt.Println("Mocky log: ", string(cmdoutput))}
//      
//	
//	
//	c.HTML(
//    // Set the HTTP status to 200 (OK)
//    http.StatusOK,
//    // Use the index.html template
//    "server.html",
//    // Pass the data that the page uses
//    gin.H{
//      "title":   "Server",
//     },
//  )
//}