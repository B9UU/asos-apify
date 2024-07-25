package main

type AsosResp struct {
	ItemCount int `json:"itemCount"`
	Products  []struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Price struct {
			Current struct {
				Value float64 `json:"value"`
				Text  string  `json:"text"`
			} `json:"current"`
			Previous struct {
				Value any    `json:"value"`
				Text  string `json:"text"`
			} `json:"previous"`
			Rrp struct {
				Value any    `json:"value"`
				Text  string `json:"text"`
			} `json:"rrp"`
			IsMarkedDown  bool   `json:"isMarkedDown"`
			IsOutletPrice bool   `json:"isOutletPrice"`
			Currency      string `json:"currency"`
		} `json:"price"`
		Colour              string   `json:"colour"`
		ColourWayID         int      `json:"colourWayId"`
		BrandName           string   `json:"brandName"`
		HasVariantColours   bool     `json:"hasVariantColours"`
		HasMultiplePrices   bool     `json:"hasMultiplePrices"`
		GroupID             any      `json:"groupId"`
		ProductCode         int      `json:"productCode"`
		ProductType         string   `json:"productType"`
		URL                 string   `json:"url"`
		ImageURL            string   `json:"imageUrl"`
		AdditionalImageUrls []string `json:"additionalImageUrls"`
		VideoURL            any      `json:"videoUrl"`
		ShowVideo           bool     `json:"showVideo"`
		IsSellingFast       bool     `json:"isSellingFast"`
		IsRestockingSoon    bool     `json:"isRestockingSoon"`
		IsPromotion         bool     `json:"isPromotion"`
		SponsoredCampaignID any      `json:"sponsoredCampaignId"`
		FacetGroupings      []any    `json:"facetGroupings"`
		Advertisement       any      `json:"advertisement"`
	} `json:"products"`
}
