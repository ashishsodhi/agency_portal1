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
CREATE TABLE `BOOKING_TRIAGE_TRACKER` (
  `PROVIDER_URI` varchar(45) NOT NULL,
  `TRIAGE_ORDER` int(11) DEFAULT NULL,
  `BOOKING_URI` varchar(45) DEFAULT NULL,
  `TIME_OFFERED_IN_MINS` int(11) DEFAULT NULL,
  `TRACKER_TYPE` enum('ACCEPTANCE_TIMER','ASSIGNMENT_TIMER','NO_TIMER') NOT NULL,
  `ACTION_TAKEN` enum('PENDING','FULFILLED','CANCELLED') DEFAULT NULL,
  `ACTION_TAKEN_TLM` datetime DEFAULT NULL,
  `STATUS` enum('ACTIVE','INACTIVE') DEFAULT NULL,
  `STATUS_TLM` datetime DEFAULT NULL,
  `TRIAGE_COMPLETED` bit(1) DEFAULT NULL,
  `WHEN_CREATED` datetime DEFAULT NULL,
  `TLM` datetime DEFAULT NULL,
  PRIMARY KEY (`PROVIDER_URI`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8


PrimaryKeyNamesList: [PROVIDER_URI]
PrimaryKeyNames    : PROVIDER_URI
delSql             : DELETE FROM BOOKING_TRIAGE_TRACKER where PROVIDER_URI = $1
updateSql          : UPDATE BOOKING_TRIAGE_TRACKER set TRIAGE_ORDER = $1, BOOKING_URI = $2, TIME_OFFERED_IN_MINS = $3, TRACKER_TYPE = $4, ACTION_TAKEN = $5, ACTION_TAKEN_TLM = $6, STATUS = $7, STATUS_TLM = $8, TRIAGE_COMPLETED = $9, WHEN_CREATED = $10, TLM = $11 WHERE PROVIDER_URI = $12
insertSql          : INSERT INTO BOOKING_TRIAGE_TRACKER ( PROVIDER_URI,  TRIAGE_ORDER,  BOOKING_URI,  TIME_OFFERED_IN_MINS,  TRACKER_TYPE,  ACTION_TAKEN,  ACTION_TAKEN_TLM,  STATUS,  STATUS_TLM,  TRIAGE_COMPLETED,  WHEN_CREATED,  TLM) values ( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12 )
selectOneSql       : SELECT * FROM BOOKING_TRIAGE_TRACKER WHERE PROVIDER_URI=?
selectMultiSql     : SELECT * FROM BOOKING_TRIAGE_TRACKER

*/

// GetAllBOOKINGTRIAGETRACKERS is a function to get a slice of record(s) from BOOKING_TRIAGE_TRACKER table in the agency_portal database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllBOOKINGTRIAGETRACKERS(ctx context.Context, page, pagesize int64, order string) (bookingtriagetrackers []*model.BOOKINGTRIAGETRACKER, totalRows int, err error) {
	sql := "SELECT * FROM BOOKING_TRIAGE_TRACKER"

	if order != "" {
		sql = sql + " order by " + order
	}
	sql = fmt.Sprintf("%s LIMIT %d, %d", sql, page, pagesize)
	fmt.Printf("%s\n", sql)
	err = DB.SelectContext(ctx, &bookingtriagetrackers, sql, page, pagesize)
	return bookingtriagetrackers, len(bookingtriagetrackers), err
}

// GetBOOKINGTRIAGETRACKER is a function to get a single record to BOOKING_TRIAGE_TRACKER table in the agency_portal database
// error - ErrNotFound, db Find error
func GetBOOKINGTRIAGETRACKER(ctx context.Context, argPROVIDERURI string) (record *model.BOOKINGTRIAGETRACKER, err error) {
	sql := "SELECT * FROM BOOKING_TRIAGE_TRACKER WHERE PROVIDER_URI=?"
	record = &model.BOOKINGTRIAGETRACKER{}
	err = DB.GetContext(ctx, record, sql, argPROVIDERURI)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// AddBOOKINGTRIAGETRACKER is a function to add a single record to BOOKING_TRIAGE_TRACKER table in the agency_portal database
// error - ErrInsertFailed, db save call failed
func AddBOOKINGTRIAGETRACKER(ctx context.Context, record *model.BOOKINGTRIAGETRACKER) (result *model.BOOKINGTRIAGETRACKER, RowsAffected int64, err error) {
	sql := "INSERT INTO BOOKING_TRIAGE_TRACKER ( PROVIDER_URI,  TRIAGE_ORDER,  BOOKING_URI,  TIME_OFFERED_IN_MINS,  TRACKER_TYPE,  ACTION_TAKEN,  ACTION_TAKEN_TLM,  STATUS,  STATUS_TLM,  TRIAGE_COMPLETED,  WHEN_CREATED,  TLM) values ( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12 )"
	dbResult := DB.MustExecContext(ctx, sql, record.TRIAGEORDER, record.BOOKINGURI, record.TIMEOFFEREDINMINS, record.TRACKERTYPE, record.ACTIONTAKEN, record.ACTIONTAKENTLM, record.STATUS, record.STATUSTLM, record.TRIAGECOMPLETED, record.WHENCREATED, record.TLM)
	id, err := dbResult.LastInsertId()
	fmt.Printf("LastInsertId: %d\n", id)
	rows, err := dbResult.RowsAffected()
	fmt.Printf("RowsAffected: %d\n", rows)
	record.PROVIDERURI = string(id)

	return record, rows, err
}

// UpdateBOOKINGTRIAGETRACKER is a function to update a single record from BOOKING_TRIAGE_TRACKER table in the agency_portal database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateBOOKINGTRIAGETRACKER(ctx context.Context, argPROVIDERURI string, updated *model.BOOKINGTRIAGETRACKER) (result *model.BOOKINGTRIAGETRACKER, RowsAffected int64, err error) {
	sql := "UPDATE BOOKING_TRIAGE_TRACKER set TRIAGE_ORDER = $1, BOOKING_URI = $2, TIME_OFFERED_IN_MINS = $3, TRACKER_TYPE = $4, ACTION_TAKEN = $5, ACTION_TAKEN_TLM = $6, STATUS = $7, STATUS_TLM = $8, TRIAGE_COMPLETED = $9, WHEN_CREATED = $10, TLM = $11 WHERE PROVIDER_URI = $12"
	fmt.Printf("sql: %s\n", sql)
	dbResult := DB.MustExecContext(ctx, sql, updated.TRIAGEORDER, updated.BOOKINGURI, updated.TIMEOFFEREDINMINS, updated.TRACKERTYPE, updated.ACTIONTAKEN, updated.ACTIONTAKENTLM, updated.STATUS, updated.STATUSTLM, updated.TRIAGECOMPLETED, updated.WHENCREATED, updated.TLM, argPROVIDERURI)
	id, err := dbResult.LastInsertId()
	fmt.Printf("LastInsertId: %d\n", id)
	rows, err := dbResult.RowsAffected()
	fmt.Printf("RowsAffected: %d\n", rows)

	updated.PROVIDERURI = argPROVIDERURI
	return updated, rows, err
}

// DeleteBOOKINGTRIAGETRACKER is a function to delete a single record from BOOKING_TRIAGE_TRACKER table in the agency_portal database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteBOOKINGTRIAGETRACKER(ctx context.Context, argPROVIDERURI string) (rowsAffected int64, err error) {
	sql := "DELETE FROM BOOKING_TRIAGE_TRACKER where PROVIDER_URI = $1"
	result := DB.MustExecContext(ctx, sql, argPROVIDERURI)
	return result.RowsAffected()
}
