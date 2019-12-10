package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pborman/uuid"

	supermart "github.com/alka/supermartchannel"
	"github.com/alka/supermartchannel/api"
	"github.com/alka/supermartchannel/http/utils"
	"github.com/alka/supermartchannel/store"
	"github.com/alka/supermartchannel/store/model"
)

type MartHandler struct {
	manager   supermart.MartManager
	martStore store.MartStore
}

func (m *MartHandler) InstallRoutes(mux *mux.Router) {
	// It will Create a supermart
	mux.Methods(http.MethodPost).Path("/supermarts").HandlerFunc(m.CreateNewMart)
	// List all items of a supermart
	mux.Methods(http.MethodGet).Path("/supermarts/{supermart_name}/items").HandlerFunc(m.GetItems)
	// Create a new item of super mart
	mux.Methods(http.MethodPost).Path("/supermarts/{supermart_name}/items").HandlerFunc(m.CreateItem)
	// Update an existing item of a mart
	mux.Methods(http.MethodPut).Path("/supermarts/{supermart_name}/items/{itemID}").HandlerFunc(m.UpdateItem)
	// delete any item of a mart
	mux.Methods(http.MethodDelete).Path("/supermarts/{supermart_name}/items/{itemID}").HandlerFunc(m.DeleteItem)
}

func (m *MartHandler) CreateNewMart(w http.ResponseWriter, r *http.Request) {
	createReq := &api.SuperMart{}
	if err := json.NewDecoder(r.Body).Decode(createReq); err != nil {
		utils.WriteErrorResponse(http.StatusBadRequest, err, w)
		return
	}
	createReq.ID = uuid.New()
	if err := m.martStore.CreateMart(r.Context(), &model.SuperMart{
		Name: createReq.Name,
		ID:   createReq.ID,
	}); err != nil {
		utils.WriteErrorResponse(http.StatusInternalServerError, err, w)
		return
	}
	utils.WriteResponse(http.StatusCreated, createReq, w)
}

func (m *MartHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	martName := mux.Vars(r)["supermart_name"]
	items, err := m.manager.GetItems(r.Context(), martName)
	if err != nil {
		utils.WriteErrorResponse(http.StatusInternalServerError, err, w)
		return
	}
	utils.WriteResponse(http.StatusOK, items, w)
}

func (m *MartHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	martName := mux.Vars(r)["supermart_name"]
	itemCreateReq := &api.ItemRequest{}
	if err := json.NewDecoder(r.Body).Decode(itemCreateReq); err != nil {
		utils.WriteErrorResponse(http.StatusBadRequest, err, w)
		return
	}
	resp, err := m.manager.CreateItem(r.Context(), martName, itemCreateReq)
	if err != nil {
		utils.WriteErrorResponse(http.StatusInternalServerError, err, w)
		return
	}
	utils.WriteResponse(http.StatusCreated, resp, w)
}

func (m *MartHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	martName := mux.Vars(r)["supermart_name"]
	itemID := mux.Vars(r)["itemID"]
	itemCreateReq := &api.ItemRequest{}
	if err := json.NewDecoder(r.Body).Decode(itemCreateReq); err != nil {
		utils.WriteErrorResponse(http.StatusBadRequest, err, w)
		return
	}
	resp, err := m.manager.UpdateItem(r.Context(), martName, itemID, itemCreateReq)
	if err != nil {
		utils.WriteErrorResponse(http.StatusInternalServerError, err, w)
		return
	}
	utils.WriteResponse(http.StatusCreated, resp, w)
}

func (m *MartHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	martName := mux.Vars(r)["supermart_name"]
	itemID := mux.Vars(r)["itemID"]

	if err := m.manager.DeleteItem(r.Context(), martName, itemID); err != nil {
		utils.WriteErrorResponse(http.StatusInternalServerError, err, w)
		return
	}
	utils.WriteResponse(http.StatusNoContent, nil, w)
	return
}

func NewMartHandler(manager supermart.MartManager, martStore store.MartStore) *MartHandler {
	return &MartHandler{
		manager:   manager,
		martStore: martStore,
	}
}
