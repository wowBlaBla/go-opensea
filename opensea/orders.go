package opensea

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func (c *OpenSeaClient) GetCheapestOrders(contract_addr string, token_id string, side string) (map[string]interface{}, error) {
	var osResp map[string]interface{}
	u, err := url.Parse(fmt.Sprintf("%s/wyvern/v1/orders", c.baseURL))
	if err != nil {
		c.Log.Errorf("Error parsing url: %s", err)
		return osResp, err
	}

	// Set query params
	q := u.Query()
	q.Set("asset_contract_address", contract_addr)
	q.Set("token_id", token_id)
	q.Set("side", side)
	q.Set("bundled", "false")
	q.Set("include_bundled", "false")
	q.Set("order_by", "eth_price")
	q.Set("order_direction", "asc")
	u.RawQuery = q.Encode()

	resp, err := c.Get(u)
	if err != nil {
		c.Log.Errorf("Error getting Orders: %s", err)
		return osResp, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&osResp)
	if err != nil {
		c.Log.Errorf("Error decoding response: %s", err)
		return osResp, err
	}

	// TODO: Filter out Orders with hidden collections
	// var filtered = []Order{}
	// for _, Order := range osResp.Orders {
	// 	if !Order.Collection.Hidden {
	// 		filtered = append(filtered, Order)
	// 	}
	// }

	return osResp, nil
}
