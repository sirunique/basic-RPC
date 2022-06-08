package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	title string
	body  string
}

type API int

var database []Item

func (a *API) GetByName(title string, reply *Item) error {
	var getItem Item

	for _, val := range database {
		if val.title == title {
			getItem = val
		}
	}

	*reply = getItem

	return nil
}

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item

	return nil
}

func (a *API) EditItem(edit Item, reply *Item) error {
	var changed Item

	for idx, val := range database {
		if val.title == edit.title {
			database[idx] = Item{edit.title, edit.body}
			changed = database[idx]
		}
	}

	*reply = changed

	return nil
}

func (a *API) DeleteItem(item Item, reply *Item) error {
	var del Item
	for idx, val := range database {
		if val.title == item.title && val.body == item.body {
			database = append(database[:idx], database[idx+1:]...)
			del = item
			break
		}
	}

	*reply = del

	return nil
}

func main() {
	api := new(API)
	err := rpc.Register(api)

	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal("Listener error", err)
	}

	log.Printf("serving rpc on port %d", 4040)
	http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error serving: ", err)
	}

	// fmt.Println("Initial database: ", database)
	// a := Item{"first", "first item"}
	// b := Item{"second", "second item"}
	// c := Item{"third", "third item"}

	// AddItem(a)
	// AddItem(b)
	// AddItem(c)
	// fmt.Println("Second database: ", database)

	// DeleteItem(b)
	// fmt.Println("Third database: ", database)

	// EditItem("third", Item{"fourth", "fourth item"})
	// fmt.Println("Fourth database: ", database)

	// x := GetByName("fourth")
	// y := GetByName("first")
	// fmt.Println(x, y)

}
