package resource

type List struct {
	// The total count of items in the list.
	Count int `json:"count"`

	// URL to the next page of items (if available).
	Next string `json:"next"`

	// URL to the previous page of items (if available).
	Previous interface{} `json:"previous"`

	// An array of items, where each item has a name and a URL.
	Results []struct {
		// The name or identifier of the item.
		Name string `json:"name"`

		// The URL to access detailed information about the item.
		Url string `json:"url"`
	} `json:"results"`
}
