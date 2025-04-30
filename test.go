package main

import (
	"fmt"
	"time"
)

func generateToken() (string, error) {
	secretKeyID := "cvv1duqi10vs73ahjcc0"
	secretKey := "zZnS4ZrmyOPvIj31cLFl3CNpVeWO85oR"

	exp := time.Now().Add(24 * time.Hour).Unix()

	claims := jwt.MapClaims{
		"key_id":   secretKeyID,
		"exp":      exp,
		"event_id": "30c1d59e-49eb-42cf-bb6a-6684aa92d4c8",
		"ip":       "1.2.3.4",
		"sub":      "abcdef123456",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	fmt.Println(token)
	return token, nil
}
