package accounts

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"project/accounts/dto"
	"project/accounts/models"
	"sync"
)

func New() *Handler {
	return &Handler{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
}

type Handler struct {
	accounts map[string]*models.Account
	guard    *sync.RWMutex
}

func (h *Handler) CreateAccount(c echo.Context) error {
	var request dto.CreateAccountRequest // {"name": "alice", "amount": 50}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	if request.Amount < 0 {
		return c.String(http.StatusBadRequest, "negative amount")
	}

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; ok {
		h.guard.Unlock()

		return c.String(http.StatusForbidden, "account already exists")
	}

	h.accounts[request.Name] = &models.Account{
		Name:   request.Name,
		Amount: request.Amount,
	}

	h.guard.Unlock()

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) GetAccount(c echo.Context) error {
	name := c.QueryParams().Get("name")

	h.guard.RLock()

	account, ok := h.accounts[name]

	h.guard.RUnlock()

	if !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	response := dto.GetAccountResponse{
		Name:   account.Name,
		Amount: account.Amount,
	}

	return c.JSON(http.StatusOK, response)
}

// Удаляет аккаунт
func (h *Handler) DeleteAccount(c echo.Context) error {
	var request dto.DeleteAccountRequest // {"name": "alice", "amount": 20}

	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	h.guard.Lock()

	_, ok := h.accounts[request.Name]

	if !ok {
		h.guard.Unlock()

		return c.String(http.StatusNotFound, "account not found")
	}

	delete(h.accounts, request.Name)

	h.guard.Unlock()

	return c.NoContent(http.StatusOK)
}

// Меняет баланс
func (h *Handler) PatchAccount(c echo.Context) error {
	var request dto.PatchAccountRequest // {"name": "alice", "amount": 20}

	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	if request.Amount < 0 {
		return c.String(http.StatusBadRequest, "negative amount")
	}

	h.guard.Lock()

	account, ok := h.accounts[request.Name]

	if !ok {
		h.guard.Unlock()

		return c.String(http.StatusNotFound, "account not found")
	}

	account.Amount = request.Amount

	h.guard.Unlock()

	return c.NoContent(http.StatusOK)
}

// Меняет имя
func (h *Handler) ChangeAccount(c echo.Context) error {
	var request dto.ChangeAccountRequest // {"name": "alice", "new_name": "bob"}

	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}
	if len(request.NewName) == 0 {
		return c.String(http.StatusBadRequest, "empty new name")
	}

	h.guard.Lock()

	account, ok := h.accounts[request.Name]

	if !ok {
		h.guard.Unlock()

		return c.String(http.StatusNotFound, "account not found")
	}

	account.Name = request.NewName

	h.guard.Unlock()

	return c.NoContent(http.StatusOK)
}

// Написать клиент консольный, который делает запросы
