package presenter

type Product struct {
	ID            string `json:"id,omitempty"`
	Name          string `json:"name"`
	BaseProductID string `json:"base_product_id"`
}
