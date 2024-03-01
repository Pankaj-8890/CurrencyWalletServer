package Currencyconverter

import (
	"context"
	pb "currencyServer/proto"
	"testing"
)

func TestConvertToINR(t *testing.T) {

	type Money struct {
		currency string
		value    float32
	}
	tests := []struct {
		name           string
		money          Money
		expectedResult float32
	}{
		{
			"TestWhenConvert1USDTOINR",
			Money{"USD", 1},
			83.10,
		},
		{
			"TestWhenConvert1INRTOINR",
			Money{"INR", 1},
			1.0,
		},
		{
			"TestWhenConvert1EURTOINR",
			Money{"EUR", 1},
			89.04,
		},
	}

	server := &Server{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &pb.Request{
				Currency: tt.money.currency,
				Value:    tt.money.value,
			}

			ctx := context.Background()
			res, err := server.ConvertToINR(ctx, req)
			if err != nil {
				t.Fatalf("ConvertToINR failed: %v", err)
			}
			if res.Value != tt.expectedResult {
				t.Errorf("Unexpected response value. Expected %f, got %f", tt.expectedResult, res.Value)
			}
		})
	}
}

func TestConvertFromINR(t *testing.T) {

	type Money struct {
		currency string
		value    float32
	}
	tests := []struct {
		name           string
		money          Money
		convertTo      string
		expectedResult float32
	}{
		{
			"TestWhenConvert1INRTOUSD",
			Money{"INR", 100.0},
			"USD",
			1.20,
		},
		{
			"TestWhenConvert1INRTOINR",
			Money{"INR", 1.0},
			"INR",
			1.0,
		},
		{
			"TestWhenConvert1INRTOEUR",
			Money{"INR", 89.04},
			"EUR",
			1.0,
		},
	}

	server := &Server{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &pb.Request{
				Currency: tt.money.currency,
				TargetCurrency: tt.convertTo,
				Value:    tt.money.value,
			}

			ctx := context.Background()
			res, err := server.ConvertFromINR(ctx, req)
			if err != nil {
				t.Fatalf("ConvertToINR failed: %v", err)
			}
			if res.Value != tt.expectedResult {
				t.Errorf("Unexpected response value. Expected %f, got %f", tt.expectedResult, res.Value)
			}
		})
	}
}
