package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	
   //"github.com/Sirupsen/logrus"
   "os"
   "errors"
 
//	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	
	"strings"
)

func connect() (session *mgo.Session) {
		
	connectURL := "localhost"
	session, err := mgo.Dial(connectURL)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}
	session.SetSafe(&mgo.Safe{})
	return session
}

func getServicesFromDB() (Services []Service) {
	session := connect()
	defer session.Close()

	collection := session.DB("MockyServices").C("StubsNew")
	fmt.Println("Query Executing\n")

	// Retrieve the list of service.
	//var Services []Service
	err := collection.Find(nil).All(&Services)
	if err != nil {
		fmt.Println("RunQuery : ERROR : %s\n", err)
		return
	}

	fmt.Println("RunQuery:no of services :\n", len(Services))
	
	
	return Services
}

func saveServiceinDB( s *Service) error{
	session := connect()
	defer session.Close()

	c := session.DB("MockyServices").C("StubsNew")
	err := c.Insert(s)

	if err == nil {
		return nil
	}else{
	return errors.New("Service not added")
	}
}
func getAllServices() []Service {
	return getServicesFromDB()
}
func getServiceByID(id string) (*Service, error) {
	session := connect()
	defer session.Close()
	z:=  strings.Split(id, "\"")
	
	id=z[1]
	c := session.DB("MockyServices").C("StubsNew")
     result := Service{}
         c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
//        if err != nil {
//                panic(err)
//        }
  //fmt.Println(result)
      return &result, nil
    
  
  return nil, errors.New("Service not found")
}