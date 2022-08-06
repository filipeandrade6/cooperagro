package interfaces

import "github.com/filipeandrade6/cooperagro/src/usecases"

type Row interface {
	Scan(dest ...interface{})
	Next() bool
}

type DBHandler interface {
	Execute(statement string)
	Query(statement string) Row
}

type DBRepo struct {
	dbHandlers map[string]DBHandler
	dbHandler  DBHandler
}

type (
	DBUserRepo          DBRepo
	DBBaseProductRepo   DBRepo
	DBProductRepo       DBRepo
	DBUnitOfMeasureRepo DBRepo
	DBInventoryRepo     DBRepo
)

func NewDBUserRepo(dbHandlers map[string]DBHandler) *DBUserRepo {
	return &DBUserRepo{
		dbHandlers: dbHandlers,
		dbHandler:  dbHandlers["DBUserRepo"],
	}
}

func (r *DBUserRepo) FindByID(id int) usecases.User {
	row := r.dbHandler.Query(fmt.Sprintf("SELECT role, customer FROM users WHERE id = '%d'", id))

	var role

	row.Next()
	row.Scan()


}

func NewDBBaseProductRepo(dbHandler map[string]DBHandler) *DBBaseProductRepo     {}
func NewDBProductRepo(dbHandler map[string]DBHandler) *DBProductRepo             {}
func NewDBUnitOfMeasureRepo(dbHandler map[string]DBHandler) *DBUnitOfMeasureRepo {}
func NewDBInventoryRepo(dbHandler map[string]DBHandler) *DBInventoryRepo         {}
