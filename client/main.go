package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

func main() {
	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	a := Item{"first", "first item"}
	b := Item{"second", "second item"}
	c := Item{"third", "third item"}

	client.Call("API.AddItem", a, &reply)
	client.Call("API.AddItem", b, &reply)
	client.Call("API.AddItem", c, &reply)
	client.Call("API.GetDB", "", &db)

	fmt.Println("Database: ", db)

	client.Call("API.EditItem", Item{"newItem", "new Item"}, &reply)

	client.Call("API.DeleteItem", c, &reply)
	client.Call("API.GetDB", "", &db)
	fmt.Println("Database: ", db)

	client.Call("API.GetByName", "first", &reply)
	fmt.Println("reply", reply)

}
