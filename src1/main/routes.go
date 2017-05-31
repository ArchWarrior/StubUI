package main

func initializeRoutes() {

  // Handle the index route
  router.GET("/", welcome)
  router.GET("/viewServices", showIndexPage)
   // Handle GET requests at /article/view/some_article_id
  router.GET("/service/view/:service_id", getService)
   router.GET("/service/add", addService)
   router.POST("/service/save", saveService)
   router.GET("/server", server)
   router.GET("/StartStopServer", StartStopServer)
//    router.GET("/stopServer", stopServer)
  
  
}
//<script>
//$(document).ready(function(){
//$('#download-btn').click(function() {
// $("#download-btn").text('Downloaded');
//})
//});    
//</script>

