
package main

import "github.com/gin-gonic/gin"

var router *gin.Engine

func main() {

  // Set the router as the default one provided by Gin
  router = gin.Default()

  // Process the templates at the start so that they don't have to be loaded
  // from the disk again. This makes serving HTML pages very fast.
//  router.LoadHTMLGlob("src/main/templates/*")
   router.LoadHTMLGlob("D:/gagan/workspace/EasyMock/templates/*")
 router.StaticFile("favicon.ico", "D:/gagan/workspace/EasyMock/favicon.ico")
  // Initialize the routes
  initializeRoutes()

  // Start serving the application
  router.Run(":4242")

}
//func getSession() *mgo.Session {
//	// Connect to our local mongo
//	s, err := mgo.Dial("mongodb://localhost")
//
//	// Check if connection error, is mongo running?
//	if err != nil {
//		panic(err)
//	}
//
//	// Deliver session
//	return s
//}