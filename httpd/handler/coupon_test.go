package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

const (
	amount = 350000.00
	amountInsuficient = 3500.00
)

func TestItemsCoupon(t *testing.T) {
	os.Setenv("api_mercado_libre", "https://api.mercadolibre.com/items/")
	ts := httptest.NewServer(Routes())
	defer ts.Close()

	items := fmt.Sprintf(`
    {
		"item_ids": [
			"MCO492245833",
        	"MCO492245834",
			"MCO492245835",
			"MCO556943801"
		],
		"amount": %f
	}`, amount)

	resp, err := http.Post(fmt.Sprintf("%s/coupon", ts.URL), "application/json", strings.NewReader(items))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}
}

func TestItemsCouponAmountInsuficient(t *testing.T) {
	os.Setenv("api_mercado_libre", "https://api.mercadolibre.com/items/")
	ts := httptest.NewServer(Routes())
	defer ts.Close()

	items := fmt.Sprintf(`
    {
		"item_ids": [
			"MCO492245833",
        	"MCO492245834",
			"MCO492245835",
			"MCO556943801"
		],
		"amount": %f
	}`, amountInsuficient)

	resp, err := http.Post(fmt.Sprintf("%s/coupon", ts.URL), "application/json", strings.NewReader(items))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 404 {
		t.Fatalf("Expected status code 404, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]

	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}

	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}
}
