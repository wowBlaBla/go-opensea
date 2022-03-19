package opensea

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// Asset represents an asset on OpenSea.
// https://docs.opensea.io/reference/asset-object
type Asset struct {
	AnimationOriginalURL    string             `json:"animation_original_url"`
	AnimationURL            string             `json:"animation_url"`
	AssetContract           AssetAssetContract `json:"asset_contract"`
	BackgroundColor         string             `json:"background_color"`
	Collection              AssetCollection    `json:"collection"`
	Creator                 AssetCreator       `json:"creator"`
	Decimals                int                `json:"decimals"`
	Description             string             `json:"description"`
	ExternalLink            string             `json:"external_link"`
	ID                      int                `json:"id"`
	ImageOriginalURL        string             `json:"image_original_url"`
	ImagePreviewURL         string             `json:"image_preview_url"`
	ImageThumbnailURL       string             `json:"image_thumbnail_url"`
	ImageURL                string             `json:"image_url"`
	IsNsfw                  bool               `json:"is_nsfw"`
	IsPresale               bool               `json:"is_presale"`
	LastSale                AssetLastSale      `json:"last_sale"`
	ListingDate             string             `json:"listing_date"`
	Name                    string             `json:"name"`
	NumSales                int                `json:"num_sales"`
	Owner                   AssetOwner         `json:"owner"`
	Permalink               string             `json:"permalink"`
	SellOrders              interface{}        `json:"sell_orders"`
	TokenID                 string             `json:"token_id"`
	TokenMetadata           string             `json:"token_metadata"`
	TopBid                  string             `json:"top_bid"`
	TransferFee             string             `json:"transfer_fee"`
	TransferFeePaymentToken string             `json:"transfer_fee_payment_token"`
	// TODO: Support traits
	// Traits                  []string           `json:"traits"`
}

type AssetAssetContract struct {
	Address                     string `json:"address"`
	AssetContractType           string `json:"asset_contract_type"`
	BuyerFeeBasisPoints         int    `json:"buyer_fee_basis_points"`
	CreatedDate                 string `json:"created_date"`
	DefaultToFiat               bool   `json:"default_to_fiat"`
	Description                 string `json:"description"`
	DevBuyerFeeBasisPoints      int    `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     int    `json:"dev_seller_fee_basis_points"`
	ExternalLink                string `json:"external_link"`
	ImageURL                    string `json:"image_url"`
	Name                        string `json:"name"`
	NftVersion                  string `json:"nft_version"`
	OnlyProxiedTransfers        bool   `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  int    `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints int    `json:"opensea_seller_fee_basis_points"`
	OpenseaVersion              string `json:"opensea_version"`
	Owner                       int    `json:"owner"`
	PayoutAddress               string `json:"payout_address"`
	SchemaName                  string `json:"schema_name"`
	SellerFeeBasisPoints        int    `json:"seller_fee_basis_points"`
	Symbol                      string `json:"symbol"`
	TotalSupply                 string `json:"total_supply"`
}

type AssetCollection struct {
	BannerImageURL              string           `json:"banner_image_url"`
	ChatURL                     string           `json:"chat_url"`
	CreatedDate                 string           `json:"created_date"`
	DefaultToFiat               bool             `json:"default_to_fiat"`
	Description                 string           `json:"description"`
	DevBuyerFeeBasisPoints      string           `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     string           `json:"dev_seller_fee_basis_points"`
	DiscordURL                  string           `json:"discord_url"`
	DisplayData                 AssetDisplayData `json:"display_data"`
	ExternalURL                 string           `json:"external_url"`
	Featured                    bool             `json:"featured"`
	FeaturedImageURL            string           `json:"featured_image_url"`
	Hidden                      bool             `json:"hidden"`
	ImageURL                    string           `json:"image_url"`
	InstagramUsername           string           `json:"instagram_username"`
	IsSubjectToWhitelist        bool             `json:"is_subject_to_whitelist"`
	LargeImageURL               string           `json:"large_image_url"`
	MediumUsername              string           `json:"medium_username"`
	Name                        string           `json:"name"`
	OnlyProxiedTransfers        bool             `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  string           `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints string           `json:"opensea_seller_fee_basis_points"`
	PayoutAddress               string           `json:"payout_address"`
	RequireEmail                bool             `json:"require_email"`
	SafelistRequestStatus       string           `json:"safelist_request_status"`
	ShortDescription            string           `json:"short_description"`
	Slug                        string           `json:"slug"`
	TelegramURL                 string           `json:"telegram_url"`
	TwitterUsername             string           `json:"twitter_username"`
	WikiURL                     string           `json:"wiki_url"`
}

type AssetOwner struct {
	Address       string    `json:"address"`
	Config        string    `json:"config"`
	ProfileImgURL string    `json:"profile_img_url"`
	User          AssetUser `json:"user"`
}

type AssetCreator struct {
	Address       string    `json:"address"`
	Config        string    `json:"config"`
	ProfileImgURL string    `json:"profile_img_url"`
	User          AssetUser `json:"user"`
}

type AssetUser struct {
	Username string `json:"username"`
}

type AssetDisplayData struct {
	CardDisplayStyle string `json:"card_display_style"`
}

type AssetLastSale struct {
	Asset AssetLastSaleAsset `json:"asset"`
}

type AssetLastSaleAsset struct {
	TokenID  string `json:"token_id"`
	Decimals int    `json:"decimals"`
}

type GetAssetsResponse struct {
	Assets []Asset `json:"assets"`
}

// GetAssetsWithOffset gets a list of assets with an offset
// https://docs.opensea.io/reference/getting-assets
func (c *OpenSeaClient) GetAssetsWithOffset(owner string, offset int) (GetAssetsResponse, error) {
	var osResp GetAssetsResponse
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
}

// GetAssets returns the assets for an address
func (c *OpenSeaClient) GetAssets(address string) ([]Asset, error) {
	var (
		allAssets []Asset
		offset    int
	)

	for {
		resp, err := c.GetAssetsWithOffset(address, offset)
		if err != nil {
			return allAssets, err
		}

		if len(resp.Assets) == 0 {
			break
		}

		allAssets = append(allAssets, resp.Assets...)
		offset += c.limitAssets
		time.Sleep(c.requestDelay)
	}

	return allAssets, nil
}
