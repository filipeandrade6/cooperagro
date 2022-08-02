package domain

import "github.com/vingarcia/ksql"

var (
	UnitsOfMeasureTable = ksql.NewTable("units_of_measure", "id")
	BaseProductsTable   = ksql.NewTable("base_products", "id")
	ProductsTable       = ksql.NewTable("products", "id")
	RolesTable          = ksql.NewTable("roles", "id")
	UsersTable          = ksql.NewTable("users", "id")
)
