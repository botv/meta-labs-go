package metalabs_sdk

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	access_key string;
	client *resty.Client;
}

type Response struct {
	message string
	ID string `json:"id"`
	Key string `json:"key"`
	Plan struct {
		ID string `json:"id"`
		Price string `json:"price"`
		Roles []string `json:"roles"`
		Recurring struct {
			Interval string `json:"interval"`
			IntervalCount int `json:"interval_count"`
		} `json:"recurring"`
		Name string `json:"name"`
		Amount int `json:"amount"`
		Type string `json:"type"`
		Currency string `json:"currency"`
	} `json:"plan"`
	Member struct {
		Email string `json:"email"`
		Discord struct {
			ID string `json:"id"`
			Tag string `json:"tag"`
			Username string `json:"username"`
			Discriminator string `json:"discriminator"`
			Avatar string `json:"avatar"`
		} `json:"discord"`
	} `json:"member"`
	Customer string `json:"customer"`
	Payment_Method string `json:"payment_method"`
	Subscription string `json:"subscription"`
	Metadata map[string]interface {

	} `json:"metadata"`
}

func New(key string) Client {
	return Client {
		access_key: key,
		client: resty.New(),
	}
}

func (c Client) UpdateKey(key string, meta map[string]string) (Response, error) {
	data, _ := json.Marshal(meta)
	resp, err := c.client.R().SetHeader("Content-Type", "application/json").SetHeader("Authorization", fmt.Sprintf("Basic %s", c.access_key)).SetBody([]byte(fmt.Sprintf(`{"metadata": %s}`, data))).Patch(BuildAPIUrl(key))
	
	if err != nil {
		return Response{}, err
	}

	if string(resp.Body()) == "Not Found" {
		return Response{ message: "License Not Found" }, nil
	} else {
		r, err := c.decodeToResponse(resp.Body())

		if err != nil {
			return Response{}, err;
		}
	
		return r, nil
	}
}

func(c Client) GetKey(key string) (Response, error) {
	resp, err := c.client.R().SetHeader("Authorization", fmt.Sprintf("Basic %s", c.access_key)).Get(BuildAPIUrl(key))
	if err != nil {
		return Response{}, err
	}

	if string(resp.Body()) == "Not Found" {
		return Response{ message: "License Not Found" }, nil
	} else {
		r, err := c.decodeToResponse(resp.Body())

		if err != nil {
			return Response{}, err;
		}
	
		return r, nil
	}
}

func (c Client) decodeToResponse(b []byte) (Response, error) {
	var r Response
	err := json.Unmarshal(b, &r)

	if err != nil {
		return r, nil
	}
	
	return r, nil
}

func BuildAPIUrl(key string) string { return fmt.Sprintf("https://api.metalabs.io/v4/licenses/%s", key) }