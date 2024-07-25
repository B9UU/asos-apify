### Output
The output is a JSON array of objects, each representing a product with the following fields:
### `id`
- **Type**: integer
- **Description**: Unique identifier for the product.
- **Example**: 206382728
### `name`
- **Type**: string
- **Description**: The name of the product.
- **Example**: "Topshop bardot wrap top in red"
### `price`
- **Type**: object
- **Description**: Contains pricing details for the product.
	**`current`:**
	- **Type**: object
	- **Description**: Current price of the product.
	- **Properties**:
		**`value`:** number
		- **Description**: Numeric value of the current price.
		- **Example**: 21
		**`text`: `string`**
		- **Description**: Textual representation of the current price.
		- **Example**: "£21.00"
	**`previous`:**
	- **Type**: object
	- **Description**: Previous price of the product.
	- **Properties**:
		- **value**: number
		- **Description**: Numeric value of the previous price.
		- **Example**: 28
		- **text**: string
		- **Description**: Textual representation of the previous price.
		- **Example**: "£28.00"
	**`rrp`:**
	- **Type**: object
	- **Description**: Recommended retail price, if available.
	- **Properties**:
		- **value**: null or number
		- **Description**: Numeric value of the RRP, or null if not available.
		- **Example**: null
		- **text**: string
		- **Description**: Textual representation of the RRP.
		- **Example**: "" (empty string)
	**`isMarkedDown`**:
	- **Type**: boolean
	- **Description**: Indicates whether the product is currently marked down.
	- **Example**: true
	**`isOutletPrice`**:
	- **Type**: boolean
	- **Description**: Indicates if the price is an outlet price.
	- **Example**: false
	**`currency`**:
	- **Type**: string
	- **Description**: Currency in which the price is listed.
	- **Example**: "GBP"
