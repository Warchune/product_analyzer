package main

import (
	"errors"
	"testing"
)

func TestProcessingCSV(t *testing.T) {
	tests := []struct {
		name                  string
		inputFile             string
		expectedMostExpensive string
		expectedHighestRating string
		expectedErr           error
	}{
		{
			name:                  "valid file csv",
			inputFile:             "../data/db_test.csv",
			expectedMostExpensive: "Печенье",
			expectedHighestRating: "Печенье",
			expectedErr:           nil,
		},
	}
	for _, test := range tests {
		var mostExpensive, highestRating item
		gotErr := processingCSV(test.inputFile, &mostExpensive, &highestRating)
		if !errors.Is(gotErr, test.expectedErr) || mostExpensive.Product != test.expectedMostExpensive || highestRating.Product != test.expectedHighestRating {
			t.Errorf("%s:\n got err: %v, mostExpensive: %s, highestRating:%s\nwant err: %v, mostExpensive: %s, highestRating:%s\n",
				test.inputFile, gotErr, mostExpensive.Product, highestRating.Product, test.expectedErr, test.expectedMostExpensive, test.expectedHighestRating)
		}
	}
}

func TestProcessingJSON(t *testing.T) {
	tests := []struct {
		name                  string
		inputFile             string
		expectedMostExpensive string
		expectedHighestRating string
		expectedErr           error
	}{
		{
			name:                  "valid file json",
			inputFile:             "../data/db_test.json",
			expectedMostExpensive: "Варенье",
			expectedHighestRating: "Варенье",
			expectedErr:           nil,
		},
	}
	for _, test := range tests {
		var mostExpensive, highestRating item
		gotErr := processingJSON(test.inputFile, &mostExpensive, &highestRating)
		if !errors.Is(gotErr, test.expectedErr) || mostExpensive.Product != test.expectedMostExpensive || highestRating.Product != test.expectedHighestRating {
			t.Errorf("%s:\n got err: %v, mostExpensive: %s, highestRating:%s\nwant err: %v, mostExpensive: %s, highestRating:%s\n",
				test.inputFile, gotErr, mostExpensive.Product, highestRating.Product, test.expectedErr, test.expectedMostExpensive, test.expectedHighestRating)
		}
	}
}
