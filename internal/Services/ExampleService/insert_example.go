package ExampleService

import (
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
)

var InsertExampleSql = `
INSERT INTO example(id,label)
VALUES ( UNHEX( REPLACE( ?,'-','' )), ?);
`

func (h *ExampleSvc) InsertExample(label string) (string, error) {
	for {
		// loop in case uid collision
		uid := uuid.New()

		statement, err := h.AppCtx.DB.Prepare(InsertExampleSql)
		if err != nil {
			return "", tracerr.Wrap(err)
		}

		if _, err := statement.Exec(uid, label); err != nil {
			if sqlerr, ok := err.(*mysql.MySQLError); ok {
				if sqlerr.Number == 1062 { // 1062 is 'Duplicate entry for key'
					continue
				}
			}
			return "", tracerr.Wrap(err)
		}

		return uid.String(), nil
	}
}
