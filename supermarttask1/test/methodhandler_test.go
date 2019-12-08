package main

import (
	"testing"

	supermart "github.com/alka/supermarttask1"
	"github.com/alka/supermarttask1/api"
	"github.com/alka/supermarttask1/http"
	"github.com/alka/supermarttask1/store"
)

func TestCreateNewMart(t *testing.T) {

	martStore := store.NewMartStore()
	martService := supermart.NewSuperMartService(store.NewItemStore(), martStore)
	martHandler := http.NewMartHandler(martService, martStore)

	got := martHandler.CreateNewMart("FoodWorld")
	if got != nil {
		t.Errorf("Failed to create mart %v", got)
	}
}

func TestCreateItem(t *testing.T) {
	martStore := store.NewMartStore()
	martService := supermart.NewSuperMartService(store.NewItemStore(), martStore)
	martHandler := http.NewMartHandler(martService, martStore)

	got, err := martHandler.CreateItem("FoodWorld", &api.ItemRequest{Name: "Item1", Price: 50.89})

	if got == nil && err != nil {
		t.Errorf("Failed in creation of items %v", err)
	}
}

func TestGetItems(t *testing.T) {
	martStore := store.NewMartStore()
	martService := supermart.NewSuperMartService(store.NewItemStore(), martStore)
	martHandler := http.NewMartHandler(martService, martStore)

	got, err := martHandler.GetItems("FoodWorld")

	if got == nil && err != nil {
		t.Errorf("Failed in getting items %v", err)
	}
}

func TestUpdateItem(t *testing.T) {

	martStore := store.NewMartStore()
	martService := supermart.NewSuperMartService(store.NewItemStore(), martStore)
	martHandler := http.NewMartHandler(martService, martStore)

	createReq := &api.ItemRequest{Name: "ItemNew", Price: 40.56}
	got, err := martHandler.UpdateItem("FoodWorld", "", createReq)
	if got == nil && err != nil {
		t.Errorf("Failed in getting items %v", err)
	}

}

func Test(t *testing.T) {
	martStore := store.NewMartStore()
	martService := supermart.NewSuperMartService(store.NewItemStore(), martStore)
	martHandler := http.NewMartHandler(martService, martStore)

	err := martHandler.DeleteItem("FoodWorld", "")
	if err != nil {
		t.Errorf("error in deleting item %v", err)
	}

}
