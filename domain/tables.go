package domain

import "github.com/vingarcia/ksql"

var UsersTable = ksql.NewTable("users", "id")

// ============================================================================
// Meu

var ProductsTable = ksql.NewTable("products", "id")
