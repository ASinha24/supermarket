package api

type ItemRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateItemRespose struct {
	*ItemRequest
	ID string `json:"id"`
}

type UpdateItemRequest struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}
