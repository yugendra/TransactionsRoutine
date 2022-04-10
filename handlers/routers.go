package handlers

//NewRoutes provides API routes to the handler
func (h *Handler) NewRoutes() {
	h.Router.HandleFunc("/account", h.createAccount).Methods("POST")
	h.Router.HandleFunc("/account/{account_id}", h.getAccount).Methods("GET")
	h.Router.HandleFunc("/transaction", h.createTransaction).Methods("POST")
}
