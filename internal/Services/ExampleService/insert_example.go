package ExampleService

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	//"github.com/google/uuid"
)

var InsertExampleSql = `
INSERT INTO examples(id,label)
VALUES ( UNHEX( REPLACE( ?,'-','' ), ?));
`

func (h *ExampleSvc) InsertExample(label string) (string, error) {
	//for {
	//	uid := uuid.New()

	statement, err := h.AppCtx.DB.Prepare(InsertExampleSql)
	if err != nil {
		h.AppCtx.Logger.Error(err)
		return "", err
	}

	if _, err := statement.Exec("77cc3dab-25dc-40bb-aca9-74bccb64767c", label); err != nil {
		h.AppCtx.Logger.Error(err)
		if sqlerr, ok := err.(*mysql.MySQLError); ok {
			fmt.Printf("full err: %+v\n", sqlerr)
			//if sqlerr.Number == sqlerr. {
			//	Handle the permission-denied error
			//}
		}
		return "", err
	}

	//return uid.String(), nil
	//}
	return "", nil
}
