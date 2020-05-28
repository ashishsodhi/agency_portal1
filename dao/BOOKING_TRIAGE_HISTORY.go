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
CREATE TABLE `BOOKING_TRIAGE_HISTORY` (
  `UID` varchar(45) NOT NULL,
  `PROVIDER_URI` varchar(45) DEFAULT NULL,
  `BOOKING_URI` varchar(45) DEFAULT NULL,
  `ACTION` enum('PENDING_ACCEPTANCE','ACCEPTED','DECLINE_TRIAGE','TRIAGE_TIMED_OUT','DECLINED_PERMANENTLY','FULFILLED') DEFAULT NULL,
  `AGENCY_USER_ID` int(11) DEFAULT NULL,
  `CSR_ID` int(11) DEFAULT NULL,
  `WHEN_CREATED` datetime DEFAULT NULL,
  `TLM` datetime DEFAULT NULL,
  PRIMARY KEY (`UID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8


PrimaryKeyNamesList: [UID]
PrimaryKeyNames    : UID
delSql             : DELETE FROM BOOKING_TRIAGE_HISTORY where UID = $1
updateSql          : UPDATE BOOKING_TRIAGE_HISTORY set PROVIDER_URI = $1, BOOKING_URI = $2, ACTION = $3, AGENCY_USER_ID = $4, CSR_ID = $5, WHEN_CREATED = $6, TLM = $7 WHERE UID = $8
insertSql          : INSERT INTO BOOKING_TRIAGE_HISTORY ( UID,  PROVIDER_URI,  BOOKING_URI,  ACTION,  AGENCY_USER_ID,  CSR_ID,  WHEN_CREATED,  TLM) values ( $1, $2, $3, $4, $5, $6, $7, $8 )
selectOneSql       : SELECT * FROM BOOKING_TRIAGE_HISTORY WHERE UID=?
selectMultiSql     : SELECT * FROM BOOKING_TRIAGE_HISTORY

*/

// GetAllBOOKINGTRIAGEHISTORIES is a function to get a slice of record(s) from BOOKING_TRIAGE_HISTORY table in the agency_portal database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllBOOKINGTRIAGEHISTORIES(ctx context.Context, page, pagesize int64, order string) (bookingtriagehistories []*model.BOOKINGTRIAGEHISTORY, totalRows int, err error) {
	sql := "SELECT * FROM BOOKING_TRIAGE_HISTORY"

	if order != "" {
		sql = sql + " order by " + order
	}
	sql = fmt.Sprintf("%s LIMIT %d, %d", sql, page, pagesize)
	fmt.Printf("%s\n", sql)
	err = DB.SelectContext(ctx, &bookingtriagehistories, sql, page, pagesize)
	return bookingtriagehistories, len(bookingtriagehistories), err
}

// GetBOOKINGTRIAGEHISTORY is a function to get a single record to BOOKING_TRIAGE_HISTORY table in the agency_portal database
// error - ErrNotFound, db Find error
func GetBOOKINGTRIAGEHISTORY(ctx context.Context, argUID string) (record *model.BOOKINGTRIAGEHISTORY, err error) {
	sql := "SELECT * FROM BOOKING_TRIAGE_HISTORY WHERE UID=?"
	record = &model.BOOKINGTRIAGEHISTORY{}
	err = DB.GetContext(ctx, record, sql, argUID)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// AddBOOKINGTRIAGEHISTORY is a function to add a single record to BOOKING_TRIAGE_HISTORY table in the agency_portal database
// error - ErrInsertFailed, db save call failed
func AddBOOKINGTRIAGEHISTORY(ctx context.Context, record *model.BOOKINGTRIAGEHISTORY) (result *model.BOOKINGTRIAGEHISTORY, RowsAffected int64, err error) {
	sql := "INSERT INTO BOOKING_TRIAGE_HISTORY ( UID,  PROVIDER_URI,  BOOKING_URI,  ACTION,  AGENCY_USER_ID,  CSR_ID,  WHEN_CREATED,  TLM) values ( $1, $2, $3, $4, $5, $6, $7, $8 )"
	dbResult := DB.MustExecContext(ctx, sql, record.PROVIDERURI, record.BOOKINGURI, record.ACTION, record.AGENCYUSERID, record.CSRID, record.WHENCREATED, record.TLM)
	id, err := dbResult.LastInsertId()
	fmt.Printf("LastInsertId: %d\n", id)
	rows, err := dbResult.RowsAffected()
	fmt.Printf("RowsAffected: %d\n", rows)
	record.UID = string(id)

	return record, rows, err
}

// UpdateBOOKINGTRIAGEHISTORY is a function to update a single record from BOOKING_TRIAGE_HISTORY table in the agency_portal database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateBOOKINGTRIAGEHISTORY(ctx context.Context, argUID string, updated *model.BOOKINGTRIAGEHISTORY) (result *model.BOOKINGTRIAGEHISTORY, RowsAffected int64, err error) {
	sql := "UPDATE BOOKING_TRIAGE_HISTORY set PROVIDER_URI = $1, BOOKING_URI = $2, ACTION = $3, AGENCY_USER_ID = $4, CSR_ID = $5, WHEN_CREATED = $6, TLM = $7 WHERE UID = $8"
	fmt.Printf("sql: %s\n", sql)
	dbResult := DB.MustExecContext(ctx, sql, updated.PROVIDERURI, updated.BOOKINGURI, updated.ACTION, updated.AGENCYUSERID, updated.CSRID, updated.WHENCREATED, updated.TLM, argUID)
	id, err := dbResult.LastInsertId()
	fmt.Printf("LastInsertId: %d\n", id)
	rows, err := dbResult.RowsAffected()
	fmt.Printf("RowsAffected: %d\n", rows)

	updated.UID = argUID
	return updated, rows, err
}

// DeleteBOOKINGTRIAGEHISTORY is a function to delete a single record from BOOKING_TRIAGE_HISTORY table in the agency_portal database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteBOOKINGTRIAGEHISTORY(ctx context.Context, argUID string) (rowsAffected int64, err error) {
	sql := "DELETE FROM BOOKING_TRIAGE_HISTORY where UID = $1"
	result := DB.MustExecContext(ctx, sql, argUID)
	return result.RowsAffected()
}
