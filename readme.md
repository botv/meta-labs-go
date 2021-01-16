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
  "User": {
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
