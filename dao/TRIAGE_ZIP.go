package dao

import (
	"context"
	"fmt"
	"time"

	"app/server/model"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = null.Bool{}
)

/*


DB Table Details
-------------------------------------
CREATE TABLE `TRIAGE_ZIP` (
  `ZIP` varchar(45) NOT NULL,
  `TRIAGE_UID` varchar(45) DEFAULT NULL,
  `WHEN_CREATED` datetime DEFAULT NULL,
  `TLM` datetime DEFAULT NULL,
  PRIMARY KEY (`ZIP`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8


PrimaryKeyNamesList: [ZIP]
PrimaryKeyNames    : ZIP
delSql             : DELETE FROM TRIAGE_ZIP where ZIP = $1
updateSql          : UPDATE TRIAGE_ZIP set TRIAGE_UID = $1, WHEN_CREATED = $2, TLM = $3 WHERE ZIP = $4
insertSql          : INSERT INTO TRIAGE_ZIP ( ZIP,  TRIAGE_UID,  WHEN_CREATED,  TLM) values ( $1, $2, $3, $4 )
selectOneSql       : SELECT * FROM TRIAGE_ZIP WHERE ZIP=?
selectMultiSql     : SELECT * FROM TRIAGE_ZIP

*/

// GetAllTRIAGEZIPS is a function to get a slice of record(s) from TRIAGE_ZIP table in the agency_portal database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllTRIAGEZIPS(ctx context.Context, page, pagesize int64, order string) (triagezips []*model.TRIAGEZIP, totalRows int, err error) {
	sql := "SELECT * FROM TRIAGE_ZIP"

	if order != "" {
		sql = sql + " order by " + order
	}
	sql = fmt.Sprintf("%s LIMIT %d, %d", sql, page, pagesize)
	fmt.Printf("%s\n", sql)
	err = DB.SelectContext(ctx, &triagezips, sql, page, pagesize)
	return triagezips, len(triagezips), err
}

// GetTRIAGEZIP is a function to get a single record to TRIAGE_ZIP table in the agency_portal database
// error - ErrNotFound, db Find error
func GetTRIAGEZIP(ctx context.Context, argZIP string) (record *model.TRIAGEZIP, err error) {
	sql := "SELECT * FROM TRIAGE_ZIP WHERE ZIP=?"
	record = &model.TRIAGEZIP{}
	err = DB.GetContext(ctx, record, sql, argZIP)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// AddTRIAGEZIP is a function to add a single record to TRIAGE_ZIP table in the agency_portal database
// error - ErrInsertFailed, db save call failed
func AddTRIAGEZIP(ctx context.Context, record *model.TRIAGEZIP) (result *model.TRIAGEZIP, RowsAffected int64, err error) {
	sql := "INSERT INTO TRIAGE_ZIP ( ZIP,  TRIAGE_UID,  WHEN_CREATED,  TLM) values ( $1, $2, $3, $4 )"
	dbResult := DB.MustExecContext(ctx, sql, record.TRIAGEUID, record.WHENCREATED, record.TLM)
	id, err := dbResult.LastInsertId()
	fmt.Printf("LastInsertId: %d\n", id)
	rows, err := dbResult.RowsAffected()
	fmt.Printf("RowsAffected: %d\n", rows)
	record.ZIP = string(id)

	return record, rows, err
}

// UpdateTRIAGEZIP is a function to update a single record from TRIAGE_ZIP table in the agency_portal database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateTRIAGEZIP(ctx context.Context, argZIP string, updated *model.TRIAGEZIP) (result *model.TRIAGEZIP, RowsAffected int64, err error) {
	sql := "UPDATE TRIAGE_ZIP set TRIAGE_UID = $1, WHEN_CREATED = $2, TLM = $3 WHERE ZIP = $4"
	fmt.Printf("sql: %s\n", sql)
	dbResult := DB.MustExecContext(ctx, sql, updated.TRIAGEUID, updated.WHENCREATED, updated.TLM, argZIP)
	id, err := dbResult.LastInsertId()
	fmt.Printf("LastInsertId: %d\n", id)
	rows, err := dbResult.RowsAffected()
	fmt.Printf("RowsAffected: %d\n", rows)

	updated.ZIP = argZIP
	return updated, rows, err
}

// DeleteTRIAGEZIP is a function to delete a single record from TRIAGE_ZIP table in the agency_portal database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteTRIAGEZIP(ctx context.Context, argZIP string) (rowsAffected int64, err error) {
	sql := "DELETE FROM TRIAGE_ZIP where ZIP = $1"
	result := DB.MustExecContext(ctx, sql, argZIP)
	return result.RowsAffected()
}
