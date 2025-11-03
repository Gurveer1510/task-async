// package migrate

// import (
// 	"context"

// 	"github.com/jackc/pgx/v5/pgxpool"
// )

// type Migrate struct {
// 	path string
// 	db   *pgxpool.Pool
// 	txn  *pgxpool.Tx
// }

// func NewMigrate(db *pgxpool.Pool, dirPath string) Migrate {
// 	return Migrate{
// 		db: db,
// 		path: dirPath,
// 	}
// }

// func (m *Migrate) RunMigrations() error {

// 	m.txn, err := m.db.BeginTx(context.Background(), nil)
// 	if err != nil {
// 		return err
// 	}

// 	defer m.txn.Rollback(context.Background())



// }

package migrate

func Migrate(){
	
}