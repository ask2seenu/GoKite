package main

import (
	"fmt"

	kiteconnect "github.com/zerodhatech/gokiteconnect"
)

const (
	apiKey    string = "fcfb1hnzj0w37m6w"
	apiSecret string = "1ecuyv86kfswoe0mi1dlpxprrsayw5cz"
)

type OrderParams struct {
    Exchange        string `url:"NSE"`
    Tradingsymbol   string `url:"SBIN-EQ"`
    Validity        string `url:"DAY"`
    Product         string `url:"REGULAR"`
    OrderType       string `url:"MIS"`
    TransactionType string `url:"LIMIT"`

    Quantity          int     `url:"1"`
    DisclosedQuantity int     `url:""`
    Price             float64 `url:"200"`
    TriggerPrice      float64 `url:""`

    Squareoff        float64 `url:""`
    Stoploss         float64 `url:""`
    TrailingStoploss float64 `url:""`

    Tag string `json:"tag" url:"tag,omitempty"`
}

func main() {
	// Create a new Kite connect instance
	kc := kiteconnect.New(apiKey)

	// Login URL from which request token can be obtained
	fmt.Println(kc.GetLoginURL())

	// Obtained request token after Kite Connect login flow
	// simulated here by scanning from stdin
	var requestToken string
	fmt.Scanf("%s\n", &requestToken)

	// Get user details and access token
	data, err := kc.GenerateSession(requestToken, apiSecret)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	// Set access token
	kc.SetAccessToken(data.AccessToken)

	// Get margins
	margins, err := kc.GetUserMargins()
	if err != nil {
		fmt.Printf("Error getting margins: %v", err)
	}
	fmt.Println("margins: ", margins)

	//get mf holdings

	//(func(c *Client) GetMFHoldings)()(MFHoldings, error)
	MFHoldings, err := kc.GetMFHoldings()
	if err != nil {
		fmt.Printf("Error getting mf holdings: %v", err)

	}
	fmt.Println("mf holdings:", MFHoldings)

	//get Holdings

	GetHoldings, err := kc.GetHoldings()
	if err != nil {
		fmt.Printf("Error getting mf holdings: %v", err)

	}
	fmt.Println("mf holdings:", GetHoldings)

	//get  orders

	GetOrders, err := kc.GetOrders()
	if err != nil {
		fmt.Printf("Error getting orders: %v", err)
	}
	fmt.Println("orders:", GetOrders)

	//get  positions

	GetPositions, err := kc.GetPositions()
	if err != nil {
		fmt.Printf("Error getting positions: %v", err)
	}
	fmt.Println("orders:", GetPositions)

	//place order

	PlaceOrder, err := kc.PlaceOrder(variety string, OrderParams)
	if err != nil {
		fmt.Printf("Error in placing order: %v", err)
	}
}
