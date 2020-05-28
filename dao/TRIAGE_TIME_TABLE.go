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
CREATE TABLE `TRIAGE_TIME_TABLE` (
  `UID` varchar(45) NOT NULL,
  `TRIAGE_UID` varchar(45) NOT NULL,
  `ORDER` int(11) DEFAULT NULL,
  `STATUS` enum('ACTIVE','INACTIVE') NOT NULL,
  `STATUS_TLM` datetime NOT NULL,
  `ASSIGNMENT_TIMER` int(11) DEFAULT NULL,
  `WHEN_CREATED` datetime NOT NULL,
  `TLM` datetime NOT NULL,
  PRIMARY KEY (`UID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8


PrimaryKeyNamesList: [UID]
PrimaryKeyNames    : UID
delSql             : DELETE FROM TRIAGE_TIME_TABLE where UID = $1
updateSql          : UPDATE TRIAGE_TIME_TABLE set TRIAGE_UID = $1, ORDER = $2, STATUS = $3, STATUS_TLM = $4, ASSIGNMENT_TIMER = $5, WHEN_CREATED = $6, TLM = $7 WHERE UID = $8
insertSql          : INSERT INTO TRIAGE_TIME_TABLE ( UID,  TRIAGE_UID,  ORDER,  STATUS,  STATUS_TLM,  ASSIGNMENT_TIMER,  WHEN_CREATED,  TLM) values ( $1, $2, $3, $4, $5, $6, $7, $8 )
selectOneSql       : SELECT * FROM TRIAGE_TIME_TABLE WHERE UID=?
selectMultiSql     : SELECT * FROM TRIAGE_TIME_TABLE

*/

// GetAllTRIAGETIMETABLES is a function to get a slice of record(s) from TRIAGE_TIME_TABLE table in the agency_portal database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllTRIAGETIMETABLES(ctx context.Context, page, pagesize int64, order string) (triagetimetables []*model.TRIAGETIMETABLE, totalRows int, err error) {
	sql := "SELECT * FROM TRIAGE_TIME_TABLE"

	if order != "" {
		sql = sql + " order by " + order
	}
	sql = fmt.Sprintf("%s LIMIT %d, %d", sql, page, pagesize)
	fmt.Printf("%s\n", sql)
	err = DB.SelectContext(ctx, &triagetimetables, sql, page, pagesize)
	return triagetimetables, len(triagetimetables), err
}

// GetTRIAGETIMETABLE is a function to get a single record to TRIAGE_TIME_TABLE table in the agency_portal database
// error - ErrNotFound, db Find error
func GetTRIAGETIMETABLE(ctx context.Context, argUID string) (record *model.TRIAGETIMETABLE, err error) {
	sql := "SELECT * FROM TRIAGE_TIME_TABLE WHERE UID=?"
	record = &model.TRIAGETIMETABLE{}
	err = DB.GetContext(ctx, record, sql, argUID)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// AddTRIAGETIMETABLE is a function to add a single record to TRIAGE_TIME_TABLE table in the agency_portal database
// error - ErrInsertFailed, db save call failed
func AddTRIAGETIMETABLE(ctx context.Context, record *model.TRIAGETIMETABLE) (result *model.TRIAGETIMETABLE, RowsAffected int64, err error) {
	sql := "INSERT INTO TRIAGE_TIME_TABLE ( UID,  TRIAGE_UID,  ORDER,  STATUS,  STATUS_TLM,  ASSIGNMENT_TIMER,  WHEN_CREATED,  TLM) values ( $1, $2, $3, $4, $5, $6, $7, $8 )"
	dbResult := DB.MustExecContext(ctx, sql, record.TRIAGEUID, record.ORDER, record.STATUS, record.STATUSTLM, record.ASSIGNMENTTIMER, record.WHENCREATED, record.TLM)
	id, err := dbResult.LastInsertId()
	fmt.Printf("LastInsertId: %d\n", id)
	rows, err := dbResult.RowsAffected()
	fmt.Printf("RowsAffected: %d\n", rows)
	record.UID = string(id)

	return record, rows, err
}

// UpdateTRIAGETIMETABLE is a function to update a single record from TRIAGE_TIME_TABLE table in the agency_portal database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateTRIAGETIMETABLE(ctx context.Context, argUID string, updated *model.TRIAGETIMETABLE) (result *model.TRIAGETIMETABLE, RowsAffected int64, err error) {
	sql := "UPDATE TRIAGE_TIME_TABLE set TRIAGE_UID = $1, ORDER = $2, STATUS = $3, STATUS_TLM = $4, ASSIGNMENT_TIMER = $5, WHEN_CREATED = $6, TLM = $7 WHERE UID = $8"
	fmt.Printf("sql: %s\n", sql)
	dbResult := DB.MustExecContext(ctx, sql, updated.TRIAGEUID, updated.ORDER, updated.STATUS, updated.STATUSTLM, updated.ASSIGNMENTTIMER, updated.WHENCREATED, updated.TLM, argUID)
	id, err := dbResult.LastInsertId()
	fmt.Printf("LastInsertId: %d\n", id)
	rows, err := dbResult.RowsAffected()
	fmt.Printf("RowsAffected: %d\n", rows)

	updated.UID = argUID
	return updated, rows, err
}

// DeleteTRIAGETIMETABLE is a function to delete a single record from TRIAGE_TIME_TABLE table in the agency_portal database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteTRIAGETIMETABLE(ctx context.Context, argUID string) (rowsAffected int64, err error) {
	sql := "DELETE FROM TRIAGE_TIME_TABLE where UID = $1"
	result := DB.MustExecContext(ctx, sql, argUID)
	return result.RowsAffected()
}
