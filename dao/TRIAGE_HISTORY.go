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
CREATE TABLE `TRIAGE_HISTORY` (
  `UID` varchar(45) NOT NULL,
  `TRIAGE_UID` varchar(45) DEFAULT NULL,
  `NAME` varchar(255) DEFAULT NULL,
  `ZIP` json DEFAULT NULL,
  `OPEN_JOB_TIMER` varchar(255) DEFAULT NULL,
  `TIME_TABLE` varchar(255) DEFAULT NULL,
  `CSR_ID` int(11) DEFAULT NULL,
  `WHEN_CREATED` datetime DEFAULT NULL,
  `TLM` datetime DEFAULT NULL,
  PRIMARY KEY (`UID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8


PrimaryKeyNamesList: [UID]
PrimaryKeyNames    : UID
delSql             : DELETE FROM TRIAGE_HISTORY where UID = $1
updateSql          : UPDATE TRIAGE_HISTORY set TRIAGE_UID = $1, NAME = $2, ZIP = $3, OPEN_JOB_TIMER = $4, TIME_TABLE = $5, CSR_ID = $6, WHEN_CREATED = $7, TLM = $8 WHERE UID = $9
insertSql          : INSERT INTO TRIAGE_HISTORY ( UID,  TRIAGE_UID,  NAME,  ZIP,  OPEN_JOB_TIMER,  TIME_TABLE,  CSR_ID,  WHEN_CREATED,  TLM) values ( $1, $2, $3, $4, $5, $6, $7, $8, $9 )
selectOneSql       : SELECT * FROM TRIAGE_HISTORY WHERE UID=?
selectMultiSql     : SELECT * FROM TRIAGE_HISTORY

*/

// GetAllTRIAGEHISTORIES is a function to get a slice of record(s) from TRIAGE_HISTORY table in the agency_portal database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllTRIAGEHISTORIES(ctx context.Context, page, pagesize int64, order string) (triagehistories []*model.TRIAGEHISTORY, totalRows int, err error) {
	sql := "SELECT * FROM TRIAGE_HISTORY"

	if order != "" {
		sql = sql + " order by " + order
	}
	sql = fmt.Sprintf("%s LIMIT %d, %d", sql, page, pagesize)
	fmt.Printf("%s\n", sql)
	err = DB.SelectContext(ctx, &triagehistories, sql, page, pagesize)
	return triagehistories, len(triagehistories), err
}

// GetTRIAGEHISTORY is a function to get a single record to TRIAGE_HISTORY table in the agency_portal database
// error - ErrNotFound, db Find error
func GetTRIAGEHISTORY(ctx context.Context, argUID string) (record *model.TRIAGEHISTORY, err error) {
	sql := "SELECT * FROM TRIAGE_HISTORY WHERE UID=?"
	record = &model.TRIAGEHISTORY{}
	err = DB.GetContext(ctx, record, sql, argUID)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// AddTRIAGEHISTORY is a function to add a single record to TRIAGE_HISTORY table in the agency_portal database
// error - ErrInsertFailed, db save call failed
func AddTRIAGEHISTORY(ctx context.Context, record *model.TRIAGEHISTORY) (result *model.TRIAGEHISTORY, RowsAffected int64, err error) {
	sql := "INSERT INTO TRIAGE_HISTORY ( UID,  TRIAGE_UID,  NAME,  ZIP,  OPEN_JOB_TIMER,  TIME_TABLE,  CSR_ID,  WHEN_CREATED,  TLM) values ( $1, $2, $3, $4, $5, $6, $7, $8, $9 )"
	dbResult := DB.MustExecContext(ctx, sql, record.TRIAGEUID, record.NAME, record.ZIP, record.OPENJOBTIMER, record.TIMETABLE, record.CSRID, record.WHENCREATED, record.TLM)
	id, err := dbResult.LastInsertId()
	fmt.Printf("LastInsertId: %d\n", id)
	rows, err := dbResult.RowsAffected()
	fmt.Printf("RowsAffected: %d\n", rows)
	record.UID = string(id)

	return record, rows, err
}

// UpdateTRIAGEHISTORY is a function to update a single record from TRIAGE_HISTORY table in the agency_portal database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateTRIAGEHISTORY(ctx context.Context, argUID string, updated *model.TRIAGEHISTORY) (result *model.TRIAGEHISTORY, RowsAffected int64, err error) {
	sql := "UPDATE TRIAGE_HISTORY set TRIAGE_UID = $1, NAME = $2, ZIP = $3, OPEN_JOB_TIMER = $4, TIME_TABLE = $5, CSR_ID = $6, WHEN_CREATED = $7, TLM = $8 WHERE UID = $9"
	fmt.Printf("sql: %s\n", sql)
	dbResult := DB.MustExecContext(ctx, sql, updated.TRIAGEUID, updated.NAME, updated.ZIP, updated.OPENJOBTIMER, updated.TIMETABLE, updated.CSRID, updated.WHENCREATED, updated.TLM, argUID)
	id, err := dbResult.LastInsertId()
	fmt.Printf("LastInsertId: %d\n", id)
	rows, err := dbResult.RowsAffected()
	fmt.Printf("RowsAffected: %d\n", rows)

	updated.UID = argUID
	return updated, rows, err
}

// DeleteTRIAGEHISTORY is a function to delete a single record from TRIAGE_HISTORY table in the agency_portal database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteTRIAGEHISTORY(ctx context.Context, argUID string) (rowsAffected int64, err error) {
	sql := "DELETE FROM TRIAGE_HISTORY where UID = $1"
	result := DB.MustExecContext(ctx, sql, argUID)
	return result.RowsAffected()
}
