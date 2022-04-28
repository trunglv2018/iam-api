package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"live-coding/models"
	"net/http"
)

const (
	customerEndpoint = "https://challenge.cyberlitmus.com/api/customer/"
)

func buildEndpoint(customerID string, queryAll bool) string {
	var reqQuery = ""
	if queryAll {
		reqQuery = "?all=true"
	}
	return fmt.Sprintf("%s/%s%s", customerEndpoint, customerID, reqQuery)
}

func GetCustomerByID(customerID string, queryAll bool, token string) (*models.Customer, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", buildEndpoint(customerID, queryAll), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var customer = new(models.Customer)
	err = json.Unmarshal(resBody, &customer)
	if err != nil {
		return nil, err
	}
	if customer != nil {
		customer.RefactorEmail()
	}
	return customer, nil
}
