# ASOS Search Scraper Actor

---

## Asos Search Scraper
This Apify Actor scrapes search results from ASOS.com based on the search query you provide. You can specify where to start and where to stop in the search results.

## Input Parameters
The Actor accepts the following input parameters:

- search_query (required): The search query string. This is the term you want to search for on ASOS.
- start (optional): The index of the item to start scraping from. Defaults to 0.
- last_items (optional): The index of the item to stop scraping at. If not provided, scraping will continue until the last item.
### Example Input
```json
{
  "search_query": "shoes",
  "start": 0,
  "last_items": 10
}
```
In this example, the Actor will start scraping from the first item and stop at the 10th item.

## How It Works
1. Initialization: The Actor initializes with the search query q and the optional start and last_items parameters.
2. Pagination: The Actor navigates through the search results pages to gather items.
3. Data Extraction: For each item in the specified range, the Actor extracts relevant information such as item title, price, URL, and image.
4. Result Output: The extracted data is output to the dataset
### Output
The output is a JSON array of objects, each representing a product with the following fields:
- **`id`**
	- **Type**: `integer`
	- **Description**: Unique identifier for the product.
	- **Example**: `206382728`
- `name`
	- **Type**: `string`
	- **Description**: The name of the product.
	- **Example**: `"Topshop bardot wrap top in red"`
- `price`
	- **Type**: `object`
	- **Description**: Contains pricing details for the product.
		- **`current`:**
			- **Type**: `object`
			- **Description**: Current price of the product.
			- **Properties**:
				- **`value`:** `number`
					* **Description**: Numeric value of the current price.
					* **Example**: `21`
				- **`text`: `string`**
					- **Description**: Textual representation of the current price.
					- **Example**: `"£21.00"`
		- **`previous`:**
			* **Type**: `object`
			- **Description**: Previous price of the product.
			- **Properties**:
				- **value**: `number`
					- **Description**: Numeric value of the previous price.
					- **Example**: `28`
				- **text**: `string`
					- **Description**: Textual representation of the previous price.
					- **Example**: `"£28.00"`
		- **`rrp`:**
			- **Type**: `object`
			- **Description**: Recommended retail price, if available.
			- **Properties**:
				- **value**: `null` or `number`
					- **Description**: Numeric value of the RRP, or null if not available.
					- **Example**: null
				- **text**: `string`
					- **Description**: Textual representation of the RRP.
					- **Example**: `""` (empty string)
	- **`isMarkedDown`**:
		- **Type**: `boolean`
		- **Description**: Indicates whether the product is currently marked down.
		- **Example**: `true`
	- **`isOutletPrice`**:
		- **Type**: `boolean`
		- **Description**: Indicates if the price is an outlet price.
		- **Example**: `false`
	- **`currency`:**
		- **Type**: `string`
		- **Description**: Currency in which the price is listed.
		- **Example**: `"GBP"`
-  **`url`**
	- **Type**: `string`
	- **Description**: URL for the product's page on the retailer's website.
	- **Example**: "`topshop/topshop-bardot-wrap-top-in-red/prd/206382728#colourWayId-206382729"`
- **`imageUrl`**
	- **Type**: `string`
	- **Description**: URL of the main image for the product.
	- **Example**: "`images.asos-media.com/products/topshop-bardot-wrap-top-in-red/206382728-1-red"`
- **`additionalImageUrls`**
	- **Type**: `array` of `string`
	- **Description**: URLs of additional images of the product.
	- **Example**:
	``` json
	[
	  "images.asos-media.com/products/topshop-bardot-wrap-top-in-red/206382728-2",
	  "images.asos-media.com/products/topshop-bardot-wrap-top-in-red/206382728-3",
	  "images.asos-media.com/products/topshop-bardot-wrap-top-in-red/206382728-4"
	]
	```
### Example Output
```json 
[
{
  "id": 206382728,
  "name": "Topshop bardot wrap top in red",
  "price": {
    "current": {
      "value": 21,
      "text": "£21.00"
    },
    "previous": {
      "value": 28,
      "text": "£28.00"
    },
    "rrp": {
      "value": null,
      "text": ""
    },
    "isMarkedDown": true,
    "isOutletPrice": false,
    "currency": "GBP"
  },
  "colour": "",
  "colourWayId": 206382729,
  "brandName": "Topshop",
  "hasVariantColours": false,
  "hasMultiplePrices": false,
  "groupId": null,
  "productCode": 135048803,
  "productType": "Product",
  "url": "topshop/topshop-bardot-wrap-top-in-red/prd/206382728#colourWayId-206382729",
  "imageUrl": "images.asos-media.com/products/topshop-bardot-wrap-top-in-red/206382728-1-red",
  "additionalImageUrls": [
    "images.asos-media.com/products/topshop-bardot-wrap-top-in-red/206382728-2",
    "images.asos-media.com/products/topshop-bardot-wrap-top-in-red/206382728-3",
    "images.asos-media.com/products/topshop-bardot-wrap-top-in-red/206382728-4"
  ],
  "videoUrl": null,
  "showVideo": false,
  "isSellingFast": false,
  "isRestockingSoon": false,
  "isPromotion": false,
  "sponsoredCampaignId": null,
  "facetGroupings": [],
  "advertisement": null
},
]
```
### Usage
To use this Actor, follow these steps:

1. Create an Apify account: If you don't already have one, sign up at Apify.
2. Create a new Actor: Go to the Actors section and create a new Actor.
3. Configure the Actor: Use the provided code and configure the input parameters.
4. Run the Actor: Start the Actor and review the results in your Apify account.

For more information on how to set up and use Apify Actors, refer to the [Apify documentation.](https://docs.apify.com/)
