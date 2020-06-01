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
CREATE TABLE `TRIAGE_HISTORY` (
  `UUID` varchar(16) NOT NULL,
  `TRIAGE_UUID` varchar(16) NOT NULL,
  `NAME` varchar(255) NOT NULL,
  `ZIP` json NOT NULL,
  `OPEN_JOB_TIMER` json NOT NULL,
  `TIME_TABLE` json NOT NULL,
  `CSR_UUID` varbinary(16) NOT NULL,
  `WHEN_CREATED` datetime NOT NULL,
  `TLM` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`UUID`),
  KEY `IX_TLM` (`TLM`),
  KEY `FK_TRIAGE_HISTORY_TRIAGE_UUID` (`TRIAGE_UUID`),
  CONSTRAINT `FK_TRIAGE_HISTORY_TRIAGE_UUID` FOREIGN KEY (`TRIAGE_UUID`) REFERENCES `TRIAGE` (`UUID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8


PrimaryKeyNamesList: [UUID]
PrimaryKeyNames    : UUID
delSql             : DELETE FROM TRIAGE_HISTORY where UUID = $1
updateSql          : UPDATE TRIAGE_HISTORY set TRIAGE_UUID = $1, NAME = $2, ZIP = $3, OPEN_JOB_TIMER = $4, TIME_TABLE = $5, CSR_UUID = $6, WHEN_CREATED = $7, TLM = $8 WHERE UUID = $9
insertSql          : INSERT INTO TRIAGE_HISTORY ( UUID,  TRIAGE_UUID,  NAME,  ZIP,  OPEN_JOB_TIMER,  TIME_TABLE,  CSR_UUID,  WHEN_CREATED,  TLM) values ( $1, $2, $3, $4, $5, $6, $7, $8, $9 )
selectOneSql       : SELECT * FROM TRIAGE_HISTORY WHERE UUID=?
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
func GetTRIAGEHISTORY(ctx context.Context, argUUID string) (record *model.TRIAGEHISTORY, err error) {
	sql := "SELECT * FROM TRIAGE_HISTORY WHERE UUID=?"
	record = &model.TRIAGEHISTORY{}
	err = DB.GetContext(ctx, record, sql, argUUID)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// AddTRIAGEHISTORY is a function to add a single record to TRIAGE_HISTORY table in the agency_portal database
// error - ErrInsertFailed, db save call failed
func AddTRIAGEHISTORY(ctx context.Context, record *model.TRIAGEHISTORY) (result *model.TRIAGEHISTORY, RowsAffected int64, err error) {
	sql := "INSERT INTO TRIAGE_HISTORY ( UUID,  TRIAGE_UUID,  NAME,  ZIP,  OPEN_JOB_TIMER,  TIME_TABLE,  CSR_UUID,  WHEN_CREATED,  TLM) values ( $1, $2, $3, $4, $5, $6, $7, $8, $9 )"
	dbResult := DB.MustExecContext(ctx, sql, record.TRIAGEUUID, record.NAME, record.ZIP, record.OPENJOBTIMER, record.TIMETABLE, record.CSRUUID, record.WHENCREATED, record.TLM)
	id, err := dbResult.LastInsertId()
	fmt.Printf("LastInsertId: %d\n", id)
	rows, err := dbResult.RowsAffected()
	fmt.Printf("RowsAffected: %d\n", rows)
	record.UUID = string(id)

	return record, rows, err
}

// UpdateTRIAGEHISTORY is a function to update a single record from TRIAGE_HISTORY table in the agency_portal database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateTRIAGEHISTORY(ctx context.Context, argUUID string, updated *model.TRIAGEHISTORY) (result *model.TRIAGEHISTORY, RowsAffected int64, err error) {
	sql := "UPDATE TRIAGE_HISTORY set TRIAGE_UUID = $1, NAME = $2, ZIP = $3, OPEN_JOB_TIMER = $4, TIME_TABLE = $5, CSR_UUID = $6, WHEN_CREATED = $7, TLM = $8 WHERE UUID = $9"
	fmt.Printf("sql: %s\n", sql)
	dbResult := DB.MustExecContext(ctx, sql, updated.TRIAGEUUID, updated.NAME, updated.ZIP, updated.OPENJOBTIMER, updated.TIMETABLE, updated.CSRUUID, updated.WHENCREATED, updated.TLM, argUUID)
	id, err := dbResult.LastInsertId()
	fmt.Printf("LastInsertId: %d\n", id)
	rows, err := dbResult.RowsAffected()
	fmt.Printf("RowsAffected: %d\n", rows)

	updated.UUID = argUUID
	return updated, rows, err
}

// DeleteTRIAGEHISTORY is a function to delete a single record from TRIAGE_HISTORY table in the agency_portal database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteTRIAGEHISTORY(ctx context.Context, argUUID string) (rowsAffected int64, err error) {
	sql := "DELETE FROM TRIAGE_HISTORY where UUID = $1"
	result := DB.MustExecContext(ctx, sql, argUUID)
	return result.RowsAffected()
}
