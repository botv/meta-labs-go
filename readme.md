# Methods
* MetaSDK.New("pk_...") -> Client
* Client.GetKey("...") -> (License, Error)
* Client.UpdateKey("...", map[string]interface{}) -> (License, Error)

# Example
```
import (
    "github.com/meta-labs/meta-labs-go/metalabs_sdk"
    "log"
)

func main(){
    Client := New("pk...")
    
    /* Fetching a license */
    license, err := Client.GetKey("0000-0000-0000-0000")
    
    if err != nil {
        panic(err)
    }
    
    log.Println(license)  
    
    /* ----------- ----------- ----------- */
    
    /* Updating a license */
    
    /* Strings, Structs, Slices, Ints Supported */
    data := make(map[string]interface{})
    data["..."] = "..."
    data["..."] = {}
    data["..."] = 1
    
    updatedLicense, err := Client.GetKey("0000-0000-0000-0000")
    
    if err != nil {
        panic(err)
    }
    
    log.Println(updatedLicense)
    /* ----------- ----------- ----------- */
}
```

# License
```
type License struct {
	message       string
	ID            string `json:"id"`
	Email         string `json:"email"`
	Key           string `json:"key"`
	Customer      string `json:"customer"`
	Status        string `json:"status"`
	Created       int `json:"created"`
	CancelAt      string `json:"cancel_at"`
	Subscription  string `json:"subscription"`
	PaymentMethod string `json:"payment_method"`
	Plan struct {
		Account string `json:"account"`
		Active  bool `json:"active"`
		Product string `json:"product"`
		Price   string `json:"price"`
		Name    string `json:"name"`
		AllowUnbinding bool `json:"allow_unbinding"`
		AllowReselling bool `json:"allow_reselling"`
		Amount         int `json:"amount"`
		Created        int `json:"created"`
		Currency       string `json:"currency"`
		Roles          []string `json:"roles"`
		Recurring struct {
			Interval string `json:"interval"`
			IntervalCount int `json:"interval_count"`
		} `json:"recurring"`   
	} `json:"plan"`
	User struct {
		Id     string `json:"id"`
		Avatar string `json:"avatar"`
		Username string `json:"username"`
		Discriminator string `json:"discriminator"`
		AccessToken   string `json:"access_token"`
		RefreshToken  string `json:"refresh_token"`
	}
	Release string `json:"release"`
	Metadata map[string]interface {

	} `json:"metadata"`
}
```

# Response
```
{
  "id": "g2cx1Fp1F",
  "email": "ben@metalabs.io",
  "key": "1HFX-FCMH-ISYQ-K1RX",
  "customer": "cus_IYn34ncv62CI",
  "status": "active",
  "created": 1607795734795,
  "cancel_at": "",
  "subscription": "sub_HEgnhG85a0rw4d",
  "payment_method": "pm_1GgDPyLTrosAq6ebT1dCGDmD",
  "plan": {
    "account": "1HRQtzJDD",
    "active": true,
    "product": "prod_xs0PqrBwa7gs",
    "price": "price_1HobMtHQ1LAgRMRUXM2oI8ii",
    "name": "Monthly Subscription",
    "allow_unbinding": true,
    "allow_reselling": false,
    "amount": 40000,
    "created": 1605647139350,
    "currency": "usd",
    "roles": [
      "785537975491493980"
    ],
    "recurring": {
        "interval": "month",
        "interval_count": 1
    }
  },
  "user": {
    "id": "409443509163130880",
    "username": "Stray Dog",
    "discriminator": "0001",
    "avatar": "",
    "access_token": "",
    "refresh_token": ""
  },
  "release": "",
  "metadata": {}
}
```
