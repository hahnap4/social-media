package main

import (
	"errors"
	"testing"
)

func TestUserIsEligible(t *testing.T) {
	// Test Data
	var tests = []struct {
		email       string
		password    string
		age         int
		expectedErr error
	}{
		{
			email:       "test@example.com",
			password:    "testpassword",
			age:         21,
			expectedErr: nil,
		},
		{
			email:       "",
			password:    "passwordishere",
			age:         18,
			expectedErr: errors.New("email can't be empty"),
		},
		{
			email:       "test@testing123.com",
			password:    "",
			age:         50,
			expectedErr: errors.New("password can't be empty"),
		},
		{
			email:       "age@isunder18.com",
			password:    "under18",
			age:         8,
			expectedErr: errors.New("age must be at least 18 years old"),
		},
	}

	// Actual Tests
	for _, tc := range tests {
		err := userIsEligible(tc.email, tc.password, tc.age)
		errString := ""
		expectedErrString := ""

		if err != nil {
			errString = err.Error()
		}
		if tc.expectedErr != nil {
			expectedErrString = tc.expectedErr.Error()
		}
		if errString != expectedErrString {
			t.Errorf("got %v, want %v", errString, expectedErrString)
		}
	}
}
