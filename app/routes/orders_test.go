package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/l-leniac-l/money-ordering-api/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestOrdersRouteInvalidBody(t *testing.T) {
	router := SetupRouter()

	reqBody := ""

	body, _ := json.Marshal(reqBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/orders", bytes.NewReader(body))

	router.ServeHTTP(w, req)

	var m map[string]interface{}

	json.NewDecoder(w.Body).Decode(&m)

	expected := map[string]interface{}{
		"message": "invalid request body",
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, expected, m)
}

func TestOrdersRoute(t *testing.T) {
	router := SetupRouter()

	reqBody := &entities.Order{
		BankId:         "f0b48033-dbb2-4bd5-b24d-8f8763e9461f",
		Amount:         100000.56,
		AdditionalInfo: "credit cards",
	}

	body, _ := json.Marshal(reqBody)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/orders", bytes.NewReader(body))

	router.ServeHTTP(w, req)

	var order entities.Order

	json.NewDecoder(w.Body).Decode(&order)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.IsType(t, entities.Order{}, order)
	assert.Equal(t, order.BankId, reqBody.BankId)
	assert.Equal(t, reqBody.Amount, order.Amount)
}
