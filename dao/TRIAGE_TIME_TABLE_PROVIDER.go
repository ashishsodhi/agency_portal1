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
CREATE TABLE `TRIAGE_TIME_TABLE_PROVIDER` (
  `TRIAGE_TIME_TABLE_ID` varchar(45) NOT NULL,
  `PROVIDER_URI` varchar(45) DEFAULT NULL,
  `STATUS` enum('ACTIVE','INACTIVE') NOT NULL,
  `STATUS_TLM` datetime DEFAULT NULL,
  `WHEN_CREATED` datetime NOT NULL,
  `TLM` datetime NOT NULL,
  PRIMARY KEY (`TRIAGE_TIME_TABLE_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8


PrimaryKeyNamesList: [TRIAGE_TIME_TABLE_ID]
PrimaryKeyNames    : TRIAGE_TIME_TABLE_ID
delSql             : DELETE FROM TRIAGE_TIME_TABLE_PROVIDER where TRIAGE_TIME_TABLE_ID = $1
updateSql          : UPDATE TRIAGE_TIME_TABLE_PROVIDER set PROVIDER_URI = $1, STATUS = $2, STATUS_TLM = $3, WHEN_CREATED = $4, TLM = $5 WHERE TRIAGE_TIME_TABLE_ID = $6
insertSql          : INSERT INTO TRIAGE_TIME_TABLE_PROVIDER ( TRIAGE_TIME_TABLE_ID,  PROVIDER_URI,  STATUS,  STATUS_TLM,  WHEN_CREATED,  TLM) values ( $1, $2, $3, $4, $5, $6 )
selectOneSql       : SELECT * FROM TRIAGE_TIME_TABLE_PROVIDER WHERE TRIAGE_TIME_TABLE_ID=?
selectMultiSql     : SELECT * FROM TRIAGE_TIME_TABLE_PROVIDER

*/

// GetAllTRIAGETIMETABLEPROVIDERS is a function to get a slice of record(s) from TRIAGE_TIME_TABLE_PROVIDER table in the agency_portal database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllTRIAGETIMETABLEPROVIDERS(ctx context.Context, page, pagesize int64, order string) (triagetimetableproviders []*model.TRIAGETIMETABLEPROVIDER, totalRows int, err error) {
	sql := "SELECT * FROM TRIAGE_TIME_TABLE_PROVIDER"

	if order != "" {
		sql = sql + " order by " + order
	}
	sql = fmt.Sprintf("%s LIMIT %d, %d", sql, page, pagesize)
	fmt.Printf("%s\n", sql)
	err = DB.SelectContext(ctx, &triagetimetableproviders, sql, page, pagesize)
	return triagetimetableproviders, len(triagetimetableproviders), err
}

// GetTRIAGETIMETABLEPROVIDER is a function to get a single record to TRIAGE_TIME_TABLE_PROVIDER table in the agency_portal database
// error - ErrNotFound, db Find error
func GetTRIAGETIMETABLEPROVIDER(ctx context.Context, argTRIAGETIMETABLEID string) (record *model.TRIAGETIMETABLEPROVIDER, err error) {
	sql := "SELECT * FROM TRIAGE_TIME_TABLE_PROVIDER WHERE TRIAGE_TIME_TABLE_ID=?"
	record = &model.TRIAGETIMETABLEPROVIDER{}
	err = DB.GetContext(ctx, record, sql, argTRIAGETIMETABLEID)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// AddTRIAGETIMETABLEPROVIDER is a function to add a single record to TRIAGE_TIME_TABLE_PROVIDER table in the agency_portal database
// error - ErrInsertFailed, db save call failed
func AddTRIAGETIMETABLEPROVIDER(ctx context.Context, record *model.TRIAGETIMETABLEPROVIDER) (result *model.TRIAGETIMETABLEPROVIDER, RowsAffected int64, err error) {
	sql := "INSERT INTO TRIAGE_TIME_TABLE_PROVIDER ( TRIAGE_TIME_TABLE_ID,  PROVIDER_URI,  STATUS,  STATUS_TLM,  WHEN_CREATED,  TLM) values ( $1, $2, $3, $4, $5, $6 )"
	dbResult := DB.MustExecContext(ctx, sql, record.PROVIDERURI, record.STATUS, record.STATUSTLM, record.WHENCREATED, record.TLM)
	id, err := dbResult.LastInsertId()
	fmt.Printf("LastInsertId: %d\n", id)
	rows, err := dbResult.RowsAffected()
	fmt.Printf("RowsAffected: %d\n", rows)
	record.TRIAGETIMETABLEID = string(id)

	return record, rows, err
}

// UpdateTRIAGETIMETABLEPROVIDER is a function to update a single record from TRIAGE_TIME_TABLE_PROVIDER table in the agency_portal database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateTRIAGETIMETABLEPROVIDER(ctx context.Context, argTRIAGETIMETABLEID string, updated *model.TRIAGETIMETABLEPROVIDER) (result *model.TRIAGETIMETABLEPROVIDER, RowsAffected int64, err error) {
	sql := "UPDATE TRIAGE_TIME_TABLE_PROVIDER set PROVIDER_URI = $1, STATUS = $2, STATUS_TLM = $3, WHEN_CREATED = $4, TLM = $5 WHERE TRIAGE_TIME_TABLE_ID = $6"
	fmt.Printf("sql: %s\n", sql)
	dbResult := DB.MustExecContext(ctx, sql, updated.PROVIDERURI, updated.STATUS, updated.STATUSTLM, updated.WHENCREATED, updated.TLM, argTRIAGETIMETABLEID)
	id, err := dbResult.LastInsertId()
	fmt.Printf("LastInsertId: %d\n", id)
	rows, err := dbResult.RowsAffected()
	fmt.Printf("RowsAffected: %d\n", rows)

	updated.TRIAGETIMETABLEID = argTRIAGETIMETABLEID
	return updated, rows, err
}

// DeleteTRIAGETIMETABLEPROVIDER is a function to delete a single record from TRIAGE_TIME_TABLE_PROVIDER table in the agency_portal database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteTRIAGETIMETABLEPROVIDER(ctx context.Context, argTRIAGETIMETABLEID string) (rowsAffected int64, err error) {
	sql := "DELETE FROM TRIAGE_TIME_TABLE_PROVIDER where TRIAGE_TIME_TABLE_ID = $1"
	result := DB.MustExecContext(ctx, sql, argTRIAGETIMETABLEID)
	return result.RowsAffected()
}
