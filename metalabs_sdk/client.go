package metalabs_sdk

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Authorization string;
	Client *resty.Client;
}

func New(key string) Client {
	return Client {
		Authorization: key,
		Client: resty.New(),
	}
}

func (c Client) UpdateKey(license string, data map[string]interface{}) (License, error) {
	b, _ := json.Marshal(data)
	resp, err := c.Client.R().
						SetHeader("Content-Type", "application/json").
						SetHeader("Authorization", fmt.Sprintf("Bearer %s", c.Authorization)).
						SetBody([]byte(fmt.Sprintf(`{"metadata": %s}`, b))).
						Patch(c.BuildURL(fmt.Sprintf("/licenses/%s", license)))
	
	if err != nil {
		return License{}, err
	}
	
	if string(resp.Body()) == "Not Found" {
		return License{
			message: "License Not Found",
		}, nil
	}

	res, err := c.decodeResponse(resp.Body())

	if err != nil {
		return License{}, err
	}

	return res, nil
}

func (c Client) GetKey(license string) (License, error) {
	resp, err := c.Client.R().
						SetHeader("Authorization", fmt.Sprintf("Bearer %s", c.Authorization)).
						Get(c.BuildURL(fmt.Sprintf("/licenses/%s", license)))

	if err != nil {
		return License{}, err
	}

	if string(resp.Body()) == "Not Found" {
		return License{
			message: "License Not Found",
		}, nil
	}

	res, err := c.decodeResponse(resp.Body())

	if err != nil {
		return License{}, err
	}

	return res, nil
}

func (c Client) ValidateLicense(license string) bool {
	return false
}

func(c Client) BuildURL(endpoint string) string {
	return fmt.Sprintf("https://api.metalabs.io/v4/%s", endpoint)
}

func(c Client) FindLicense(license string) string {
	return license
}

func (c Client) decodeResponse(b []byte) (License, error) {
	var r License
	err := json.Unmarshal(b, &r)

	if err != nil {
		return r, nil
	}

	return r, nil
}