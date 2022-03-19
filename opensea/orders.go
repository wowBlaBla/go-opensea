package opensea

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (c *OpenSeaClient) GetOrders(contract_addr string, token_id int, side int) (map[string]interface{}, error) {
	var osResp map[string]interface{}
	u, err := url.Parse(fmt.Sprintf("%s/api/v1/orders", c.baseURL))
	if err != nil {
		c.Log.Errorf("Error parsing url: %s", err)
		return osResp, err
	}

	// Set query params
	q := u.Query()
	q.Set("asset_contract_address", contract_addr)
	q.Set("token_id", fmt.Sprintf("%d", token_id))
	q.Set("side", fmt.Sprintf("%d", side))
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

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-API-KEY", "5de1db9cee694d7d91d7f80669c57659")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
