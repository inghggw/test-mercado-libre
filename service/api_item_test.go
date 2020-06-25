package service

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetItemMercadoLibre(t *testing.T) {
	os.Setenv("api_mercado_libre", "https://api.mercadolibre.com/items/")
	item, _ := GetItemMercadoLibre("MCO492245833")

	if item.Status != "active" {
		t.Errorf("GetItemMercadoLibre() = %s; want status active", item.Status)
	}
}

func TestGetItemMercadoLibreErr(t *testing.T) {
	os.Setenv("api_mercado_libre", "https://")
	_, err := GetItemMercadoLibre("MCO492245833")

	if err == nil {
		t.Errorf("GetItemMercadoLibre() = %s; want err", err)
	}
}

func TestGetItemMercadoLibreNotFound(t *testing.T) {
	os.Setenv("api_mercado_libre", "https://api.mercadolibre.com/items/")
	assert := assert.New(t)
	item, _ := GetItemMercadoLibre("MLA599260060")
	status := item.Status

	if status != "404" {
		t.Errorf("Status: %s want 404 " + item.Status)
	}

	assert.Equal(status, "404", "El item no fue encontrado")
}

func TestGetItemMercadoLibreErrUnmarshal(t *testing.T) {
	os.Setenv("api_mercado_libre", "https://api.mercadolibre.com/items")
	assert := assert.New(t)
	_, err := GetItemMercadoLibre("?ids=MLA1,MLA2,MLA3")

	if err == nil {
		t.Errorf("GetItemMercadoLibre() = %s; want err", err)
	}

	assert.EqualError(err, "json: cannot unmarshal array into Go value of type service.ResponseItem")
}
