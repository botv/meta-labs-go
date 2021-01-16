package metalabs_sdk

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