package presenter

type Inventory struct {
	ID              string `json:"id,omitempty"`
	UserID          string `json:"user_id"`
	ProductID       string `json:"product_id"`
	Quantity        int    `json:"quantity"`
	UnitOfMeasureID string `json:"unit_of_measure_id"`
}
