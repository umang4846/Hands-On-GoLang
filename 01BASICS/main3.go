package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

func PlaceOrderRequestConcurrency(c *gin.Context) {

	var data model.WebhookData
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	url := "https://demo.tradovateapi.com/v1/order/placeorder"
	method := "POST"

	var wg sync.WaitGroup
	var mu sync.Mutex
	var errors []error

	// Manually provide client names
	clientNames := []string{"benny", "benrafa"}

	fmt.Println(clientNames)

	for _, clientName := range clientNames {
		wg.Add(1)
		go func(clientName string) {
			defer wg.Done()

			requestBody, err := GenerateRequestBody(clientName, &data)
			if err != nil {
				// Handle error
				mu.Lock()
				errors = append(errors, err)
				mu.Unlock()
				return
			}

			accessToken := FetchAccessToken(clientName)
			if err != nil {
				// Handle error
				mu.Lock()
				errors = append(errors, err)
				mu.Unlock()
				return
			}

			requestBodyBytes, err := json.Marshal(requestBody)
			if err != nil {
				// Handle error
				mu.Lock()
				errors = append(errors, err)
				mu.Unlock()
				return
			}

			req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBodyBytes))
			if err != nil {
				// Handle error
				mu.Lock()
				errors = append(errors, err)
				mu.Unlock()
				return
			}

			req.Header.Set("Authorization", "Bearer "+accessToken)
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("Accept", "application/json")

			httpclient := &http.Client{}
			res, err := httpclient.Do(req)
			if err != nil {
				// Handle error
				mu.Lock()
				errors = append(errors, err)
				mu.Unlock()
				return
			}
			defer res.Body.Close()

			// Process response
			// responseBody, err := ioutil.ReadAll(res.Body)

		}(clientName)
	}

	// Wait for all Goroutines to finish
	wg.Wait()

	// Handle errors if needed
	// Example: if len(errors) > 0 { handleErrors(errors) }
	c.JSON(http.StatusOK, gin.H{"message": "Orders placed successfully"})
}
