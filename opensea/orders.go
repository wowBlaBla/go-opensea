package opensea

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetOrders() {
	var osResp GetOrdersResponse
	u, err := url.Parse(fmt.Sprintf("%s/api/v1/assets", c.baseURL))
	if err != nil {
		c.Log.Errorf("Error parsing url: %s", err)
		return osResp, err
	}

	// Set query params
	q := u.Query()
	q.Set("owner", owner)
	q.Set("limit", fmt.Sprint(c.limitAssets))
	q.Set("offset", fmt.Sprint(offset))
	u.RawQuery = q.Encode()

	resp, err := c.Get(u)
	if err != nil {
		c.Log.Errorf("Error getting assets: %s", err)
		return osResp, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&osResp)
	if err != nil {
		c.Log.Errorf("Error decoding response: %s", err)
		return osResp, err
	}

	// TODO: Filter out assets with hidden collections
	// var filtered = []Asset{}
	// for _, asset := range osResp.Assets {
	// 	if !asset.Collection.Hidden {
	// 		filtered = append(filtered, asset)
	// 	}
	// }

	return osResp, nil

	url := "https://api.opensea.io/wyvern/v1/orders?asset_contract_address=0x1a92f7381b9f03921564a437210bb9396471050c&bundled=false&include_bundled=false&token_id=6069&side=1&limit=20&offset=0&order_by=created_date&order_direction=desc"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-API-KEY", "5de1db9cee694d7d91d7f80669c57659")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
