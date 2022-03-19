package opensea

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type Collection struct {
	Editors                     []string                          `json:"editors"`
	PaymentTokens               []CollectionPaymentTokens         `json:"payment_tokens"`
	PrimaryAssetContracts       []CollectionPrimaryAssetContracts `json:"primary_asset_contracts"`
	Stats                       CollectionStats                   `json:"stats"`
	BannerImageURL              string                            `json:"banner_image_url"`
	ChatURL                     string                            `json:"chat_url"`
	CreatedDate                 string                            `json:"created_date"`
	DefaultToFiat               bool                              `json:"default_to_fiat"`
	Description                 string                            `json:"description"`
	DevBuyerFeeBasisPoints      string                            `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     string                            `json:"dev_seller_fee_basis_points"`
	DiscordURL                  string                            `json:"discord_url"`
	DisplayData                 CollectionDisplayData             `json:"display_data"`
	ExternalURL                 string                            `json:"external_url"`
	Featured                    bool                              `json:"featured"`
	FeaturedImageURL            string                            `json:"featured_image_url"`
	Hidden                      bool                              `json:"hidden"`
	SafelistRequestStatus       string                            `json:"safelist_request_status"`
	ImageURL                    string                            `json:"image_url"`
	IsSubjectToWhitelist        bool                              `json:"is_subject_to_whitelist"`
	LargeImageURL               string                            `json:"large_image_url"`
	MediumUsername              string                            `json:"medium_username"`
	Name                        string                            `json:"name"`
	OnlyProxiedTransfers        bool                              `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  string                            `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints string                            `json:"opensea_seller_fee_basis_points"`
	PayoutAddress               string                            `json:"payout_address"`
	RequireEmail                bool                              `json:"require_email"`
	ShortDescription            string                            `json:"short_description"`
	Slug                        string                            `json:"slug"`
	TelegramURL                 string                            `json:"telegram_url"`
	TwitterUsername             string                            `json:"twitter_username"`
	InstagramUsername           string                            `json:"instagram_username"`
	WikiURL                     string                            `json:"wiki_url"`
}

type CollectionPaymentTokens struct {
	ID       int     `json:"id"`
	Symbol   string  `json:"symbol"`
	Address  string  `json:"address"`
	ImageURL string  `json:"image_url"`
	Name     string  `json:"name"`
	Decimals int     `json:"decimals"`
	EthPrice float64 `json:"eth_price"`
	UsdPrice float64 `json:"usd_price"`
}

type CollectionPrimaryAssetContracts struct {
	Address                     string `json:"address"`
	AssetContractType           string `json:"asset_contract_type"`
	CreatedDate                 string `json:"created_date"`
	Name                        string `json:"name"`
	NftVersion                  string `json:"nft_version"`
	OpenseaVersion              string `json:"opensea_version"`
	Owner                       int    `json:"owner"`
	SchemaName                  string `json:"schema_name"`
	Symbol                      string `json:"symbol"`
	TotalSupply                 string `json:"total_supply"`
	Description                 string `json:"description"`
	ExternalLink                string `json:"external_link"`
	ImageURL                    string `json:"image_url"`
	DefaultToFiat               bool   `json:"default_to_fiat"`
	DevBuyerFeeBasisPoints      int    `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     int    `json:"dev_seller_fee_basis_points"`
	OnlyProxiedTransfers        bool   `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  int    `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints int    `json:"opensea_seller_fee_basis_points"`
	BuyerFeeBasisPoints         int    `json:"buyer_fee_basis_points"`
	SellerFeeBasisPoints        int    `json:"seller_fee_basis_points"`
	PayoutAddress               string `json:"payout_address"`
}

type CollectionStats struct {
	OneDayVolume          float64 `json:"one_day_volume"`
	OneDayChange          float64 `json:"one_day_change"`
	OneDaySales           float64 `json:"one_day_sales"`
	OneDayAveragePrice    float64 `json:"one_day_average_price"`
	SevenDayVolume        float64 `json:"seven_day_volume"`
	SevenDayChange        float64 `json:"seven_day_change"`
	SevenDaySales         float64 `json:"seven_day_sales"`
	SevenDayAveragePrice  float64 `json:"seven_day_average_price"`
	ThirtyDayVolume       float64 `json:"thirty_day_volume"`
	ThirtyDayChange       float64 `json:"thirty_day_change"`
	ThirtyDaySales        float64 `json:"thirty_day_sales"`
	ThirtyDayAveragePrice float64 `json:"thirty_day_average_price"`
	TotalVolume           float64 `json:"total_volume"`
	TotalSales            float64 `json:"total_sales"`
	TotalSupply           float64 `json:"total_supply"`
	Count                 float64 `json:"count"`
	NumOwners             int     `json:"num_owners"`
	AveragePrice          float64 `json:"average_price"`
	NumReports            int     `json:"num_reports"`
	MarketCap             float64 `json:"market_cap"`
	FloorPrice            float64 `json:"floor_price"`
}

type CollectionDisplayData struct {
	CardDisplayStyle string `json:"card_display_style"`
}

type GetCollectionResponse struct {
	Collection Collection `json:"collection"`
}

func (c *OpenSeaClient) GetCollection(slug string) (Collection, error) {
	var collection Collection

	u, err := url.Parse(fmt.Sprintf("%s/api/v1/collection/%s", c.baseURL, slug))
	if err != nil {
		c.Log.Errorf("Error parsing url: %s", err)
		return collection, err
	}

	resp, err := c.Get(u)
	if err != nil {
		c.Log.Errorf("Error getting collection: %s", err)
		return collection, err
	}

	defer resp.Body.Close()

	var osResp GetCollectionResponse
	err = json.NewDecoder(resp.Body).Decode(&osResp)
	if err != nil {
		c.Log.Errorf("Error decoding response: %s", err)
		return collection, err
	}

	return osResp.Collection, nil
}
