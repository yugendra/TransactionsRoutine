package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/yugendra/TransactionsRoutine/database"
	"github.com/yugendra/TransactionsRoutine/entities"
	"github.com/yugendra/TransactionsRoutine/transactionsroutine"
	"github.com/yugendra/TransactionsRoutine/transactionsroutine/interactor"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"strconv"
)

//Handler ...
type Handler struct {
	Router   *mux.Router
	db       interactor.DatabaseInteractor
	validate *validator.Validate
}

/*NewHandler initializer for handler, router and database
 */
func NewHandler() *Handler {
	router := mux.NewRouter()
	db := database.NewDatabase()
	validate := validator.New()
	return &Handler{Router: router, db: db, validate: validate}
}

/*createAccount handler function for /account POST request
 * Validate the inputs.
 * Create a new user account for given document number.
 * Document number should be unique for each user.
 * Returns new user id along with document number if user account is created successfully else return error.
 */
func (h *Handler) createAccount(w http.ResponseWriter, r *http.Request) {
	account := &Account{}

	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error: invalid request", http.StatusBadRequest)
		return
	}

	err = h.validate.Struct(account)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error: invalid request", http.StatusBadRequest)
		return
	}

	entityAccount := accountModulesToEntity(account)
	err = transactionsroutine.CreateAccount(h.db, entityAccount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	account = accountEntityToModule(entityAccount)
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(account)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

/*getAccount handler function for /account GET request
 *Validate the inputs.
 *If account exists in database then this API will return its id and document number else return error.
 */
func (h *Handler) getAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountID := vars["account_id"]
	intAccountID, err := strconv.Atoi(accountID)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if intAccountID <= 0 {
		errMsg := fmt.Sprintf("invalid account id %s", accountID)
		log.Println(errMsg)
		http.Error(w, errMsg, http.StatusBadRequest)
	}

	entityAccount, err := transactionsroutine.GetAccount(h.db, intAccountID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(accountEntityToModule(entityAccount))
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

/*createTransaction handler function for /transaction POST request
 *Validate the inputs.
 *Append new transaction in database for existing user.
 *If transaction is created in database successfully then return transaction id along with other information else return error.
 */
func (h *Handler) createTransaction(w http.ResponseWriter, r *http.Request) {
	transaction := &Transaction{}

	err := json.NewDecoder(r.Body).Decode(transaction)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error: invalid request", http.StatusBadRequest)
		return
	}

	err = h.validate.Struct(transaction)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error: invalid request", http.StatusBadRequest)
		return
	}

	entityTransaction := transactionModuleToEntity(transaction)
	err = transactionsroutine.CreateTransaction(h.db, entityTransaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	transaction = transactionEntityToModule(entityTransaction)
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(transaction)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func accountModulesToEntity(account *Account) *entities.Account {
	return &entities.Account{
		DocumentNumber: account.DocumentNumber,
	}
}

func accountEntityToModule(accountEntity *entities.Account) *Account {
	return &Account{
		AccountID:      accountEntity.AccountID,
		DocumentNumber: accountEntity.DocumentNumber,
	}
}

func transactionModuleToEntity(transaction *Transaction) *entities.Transaction {
	return &entities.Transaction{
		AccountID:     transaction.AccountID,
		OperationType: entities.GetOperationsType(transaction.OperationType),
		Amount:        transaction.Amount,
	}
}

func transactionEntityToModule(entityTransaction *entities.Transaction) *Transaction {
	return &Transaction{
		TransactionID: entityTransaction.TransactionID,
		AccountID:     entityTransaction.AccountID,
		OperationType: entityTransaction.OperationType.String(),
		Amount:        entityTransaction.Amount,
		EventDate:     entityTransaction.EventDate,
	}
}
