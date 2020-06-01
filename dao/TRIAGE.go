package dao

import (
	"context"
	"fmt"
	"time"

	"example.com/example/model"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = null.Bool{}
)

/*


DB Table Details
-------------------------------------
CREATE TABLE `TRIAGE` (
  `UUID` varchar(16) NOT NULL,
  `NAME` varchar(255) NOT NULL,
  `ACCEPTANCE_TIMER` json NOT NULL,
  `WHEN_CREATED` datetime NOT NULL,
  `TLM` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`UUID`),
  UNIQUE KEY `UNQ_NAME` (`NAME`),
  KEY `IX_TLM` (`TLM`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8


PrimaryKeyNamesList: [UUID]
PrimaryKeyNames    : UUID
delSql             : DELETE FROM TRIAGE where UUID = $1
updateSql          : UPDATE TRIAGE set NAME = $1, ACCEPTANCE_TIMER = $2, WHEN_CREATED = $3, TLM = $4 WHERE UUID = $5
insertSql          : INSERT INTO TRIAGE ( UUID,  NAME,  ACCEPTANCE_TIMER,  WHEN_CREATED,  TLM) values ( $1, $2, $3, $4, $5 )
selectOneSql       : SELECT * FROM TRIAGE WHERE UUID=?
selectMultiSql     : SELECT * FROM TRIAGE

*/

// GetAllTRIAGES is a function to get a slice of record(s) from TRIAGE table in the agency_portal database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllTRIAGES(ctx context.Context, page, pagesize int64, order string) (triages []*model.TRIAGE, totalRows int, err error) {
	sql := "SELECT * FROM TRIAGE"

	if order != "" {
		sql = sql + " order by " + order
	}
	sql = fmt.Sprintf("%s LIMIT %d, %d", sql, page, pagesize)
	fmt.Printf("%s\n", sql)
	err = DB.SelectContext(ctx, &triages, sql, page, pagesize)
	return triages, len(triages), err
}

// GetTRIAGE is a function to get a single record to TRIAGE table in the agency_portal database
// error - ErrNotFound, db Find error
func GetTRIAGE(ctx context.Context, argUUID string) (record *model.TRIAGE, err error) {
	sql := "SELECT * FROM TRIAGE WHERE UUID=?"
	record = &model.TRIAGE{}
	err = DB.GetContext(ctx, record, sql, argUUID)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// AddTRIAGE is a function to add a single record to TRIAGE table in the agency_portal database
// error - ErrInsertFailed, db save call failed
func AddTRIAGE(ctx context.Context, record *model.TRIAGE) (result *model.TRIAGE, RowsAffected int64, err error) {
	sql := "INSERT INTO TRIAGE ( UUID,  NAME,  ACCEPTANCE_TIMER,  WHEN_CREATED,  TLM) values ( $1, $2, $3, $4, $5 )"
	dbResult := DB.MustExecContext(ctx, sql, record.NAME, record.ACCEPTANCETIMER, record.WHENCREATED, record.TLM)
	id, err := dbResult.LastInsertId()
	fmt.Printf("LastInsertId: %d\n", id)
	rows, err := dbResult.RowsAffected()
	fmt.Printf("RowsAffected: %d\n", rows)
	record.UUID = string(id)

	return record, rows, err
}

// UpdateTRIAGE is a function to update a single record from TRIAGE table in the agency_portal database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateTRIAGE(ctx context.Context, argUUID string, updated *model.TRIAGE) (result *model.TRIAGE, RowsAffected int64, err error) {
	sql := "UPDATE TRIAGE set NAME = $1, ACCEPTANCE_TIMER = $2, WHEN_CREATED = $3, TLM = $4 WHERE UUID = $5"
	fmt.Printf("sql: %s\n", sql)
	dbResult := DB.MustExecContext(ctx, sql, updated.NAME, updated.ACCEPTANCETIMER, updated.WHENCREATED, updated.TLM, argUUID)
	id, err := dbResult.LastInsertId()
	fmt.Printf("LastInsertId: %d\n", id)
	rows, err := dbResult.RowsAffected()
	fmt.Printf("RowsAffected: %d\n", rows)

	updated.UUID = argUUID
	return updated, rows, err
}

// DeleteTRIAGE is a function to delete a single record from TRIAGE table in the agency_portal database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteTRIAGE(ctx context.Context, argUUID string) (rowsAffected int64, err error) {
	sql := "DELETE FROM TRIAGE where UUID = $1"
	result := DB.MustExecContext(ctx, sql, argUUID)
	return result.RowsAffected()
}
