package main

import (
	"fmt"

	supermart "github.com/alka/supermarttask1"
	"github.com/alka/supermarttask1/api"
	"github.com/alka/supermarttask1/http"
	"github.com/alka/supermarttask1/store"
)

func main() {
	martStore := store.NewMartStore()
	martService := supermart.NewSuperMartService(store.NewItemStore(), martStore)
	martHandler := http.NewMartHandler(martService, martStore)
	//to create a supermart
	if err := martHandler.CreateNewMart("FoodWorld"); err != nil {
		fmt.Printf("error while creating new mart %v", err)
	}

	//to create a items under supermart
	itemCreation := []api.ItemRequest{
		{
			Name:  "Item1",
			Price: 35.50,
		},
		{
			Name:  "Item2",
			Price: 50.50,
		},
		{
			Name:  "Item3",
			Price: 95.57,
		},
	}
	idupdate := ""
	for _, item := range itemCreation {
		res, err := martHandler.CreateItem("FoodWorld", &item)
		if err != nil {
			fmt.Printf("error while creating items %v", err)
		}
		idupdate = res.ID
		fmt.Println("Item has been created successfully ", res.ID)
	}

	//to get the items under a mart
	items, err := martHandler.GetItems("FoodWorld")
	if err != nil {
		fmt.Printf("error while getting items from mart %v", err)
	}
	fmt.Println("below are the items of given mart")
	for _, item := range items {
		fmt.Println(item.ItemRequest.Name, item.ItemRequest.Price, item.ID)
	}
	//to update the item
	createReq := &api.ItemRequest{Name: "ItemNew", Price: 40.56}
	res, err := martHandler.UpdateItem("FoodWorld", idupdate, createReq)
	if err != nil {
		fmt.Printf("error while getting items from mart %v", err)
	}
	fmt.Println("Item has been updated, please find updated item below\n", res.ItemRequest.Name, res.ItemRequest.Price)

	//to delete the items
	err = martHandler.DeleteItem("FoodWorld", idupdate)
	if err != nil {
		fmt.Println("Id does not exist or issues in deleting the item")
	}

	//Items post deletion
	items, err = martHandler.GetItems("FoodWorld")
	if err != nil {
		fmt.Printf("error while getting items from mart %v", err)
	}
	fmt.Println("below are the items of given mart post deletion")
	for _, item := range items {
		fmt.Println(item.ItemRequest.Name, item.ItemRequest.Price, item.ID)
	}

}
