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
CREATE TABLE `BOOKING_TRIAGE_TRACKER` (
  `UUID` varchar(16) NOT NULL,
  `PROVIDER_UUID` varchar(16) DEFAULT NULL,
  `TRIAGE_ORDER` smallint(5) unsigned DEFAULT NULL,
  `BOOKING_UUID` varchar(16) NOT NULL,
  `TIME_OFFERED_IN_MINS` smallint(5) unsigned DEFAULT NULL,
  `TIMER_EXECUTION_ID` varchar(255) DEFAULT NULL,
  `TIMER_EXPIRES_AT` datetime DEFAULT NULL,
  `TRACKING_TYPE` enum('ACCEPTANCE_TIMER','ASSIGNMENT_TIMER','NO_TIMER') DEFAULT NULL,
  `STATUS` enum('PENDING','FULFILLED','CANCELLED') NOT NULL,
  `STATUS_TLM` datetime NOT NULL,
  `TRIAGE_COMPLETED` bit(1) NOT NULL DEFAULT b'0',
  `WHEN_CREATED` datetime DEFAULT NULL,
  `TLM` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`UUID`),
  UNIQUE KEY `UNQ_BOOKING_UUID` (`BOOKING_UUID`),
  KEY `IX_TLM` (`TLM`),
  KEY `IX_TIMER_EXECUTION_ID` (`TIMER_EXECUTION_ID`),
  KEY `IX_PROVIDER_UUID_STATUS` (`PROVIDER_UUID`,`STATUS`),
  CONSTRAINT `FK_BOOKING_TRIAGE_TRACKER_BOOKING_BELONGS_TO` FOREIGN KEY (`BOOKING_UUID`) REFERENCES `BOOKING` (`UUID`),
  CONSTRAINT `FK_BOOKING_TRIAGE_TRACKER_PROVIDER_BELONGS_TO` FOREIGN KEY (`PROVIDER_UUID`) REFERENCES `PROVIDER` (`UUID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8


PrimaryKeyNamesList: [UUID]
PrimaryKeyNames    : UUID
delSql             : DELETE FROM BOOKING_TRIAGE_TRACKER where UUID = $1
updateSql          : UPDATE BOOKING_TRIAGE_TRACKER set PROVIDER_UUID = $1, TRIAGE_ORDER = $2, BOOKING_UUID = $3, TIME_OFFERED_IN_MINS = $4, TIMER_EXECUTION_ID = $5, TIMER_EXPIRES_AT = $6, TRACKING_TYPE = $7, STATUS = $8, STATUS_TLM = $9, TRIAGE_COMPLETED = $10, WHEN_CREATED = $11, TLM = $12 WHERE UUID = $13
insertSql          : INSERT INTO BOOKING_TRIAGE_TRACKER ( UUID,  PROVIDER_UUID,  TRIAGE_ORDER,  BOOKING_UUID,  TIME_OFFERED_IN_MINS,  TIMER_EXECUTION_ID,  TIMER_EXPIRES_AT,  TRACKING_TYPE,  STATUS,  STATUS_TLM,  TRIAGE_COMPLETED,  WHEN_CREATED,  TLM) values ( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13 )
selectOneSql       : SELECT * FROM BOOKING_TRIAGE_TRACKER WHERE UUID=?
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
func GetBOOKINGTRIAGETRACKER(ctx context.Context, argUUID string) (record *model.BOOKINGTRIAGETRACKER, err error) {
	sql := "SELECT * FROM BOOKING_TRIAGE_TRACKER WHERE UUID=?"
	record = &model.BOOKINGTRIAGETRACKER{}
	err = DB.GetContext(ctx, record, sql, argUUID)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// AddBOOKINGTRIAGETRACKER is a function to add a single record to BOOKING_TRIAGE_TRACKER table in the agency_portal database
// error - ErrInsertFailed, db save call failed
func AddBOOKINGTRIAGETRACKER(ctx context.Context, record *model.BOOKINGTRIAGETRACKER) (result *model.BOOKINGTRIAGETRACKER, RowsAffected int64, err error) {
	sql := "INSERT INTO BOOKING_TRIAGE_TRACKER ( UUID,  PROVIDER_UUID,  TRIAGE_ORDER,  BOOKING_UUID,  TIME_OFFERED_IN_MINS,  TIMER_EXECUTION_ID,  TIMER_EXPIRES_AT,  TRACKING_TYPE,  STATUS,  STATUS_TLM,  TRIAGE_COMPLETED,  WHEN_CREATED,  TLM) values ( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13 )"
	dbResult := DB.MustExecContext(ctx, sql, record.PROVIDERUUID, record.TRIAGEORDER, record.BOOKINGUUID, record.TIMEOFFEREDINMINS, record.TIMEREXECUTIONID, record.TIMEREXPIRESAT, record.TRACKINGTYPE, record.STATUS, record.STATUSTLM, record.TRIAGECOMPLETED, record.WHENCREATED, record.TLM)
	id, err := dbResult.LastInsertId()
	fmt.Printf("LastInsertId: %d\n", id)
	rows, err := dbResult.RowsAffected()
	fmt.Printf("RowsAffected: %d\n", rows)
	record.UUID = string(id)

	return record, rows, err
}

// UpdateBOOKINGTRIAGETRACKER is a function to update a single record from BOOKING_TRIAGE_TRACKER table in the agency_portal database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateBOOKINGTRIAGETRACKER(ctx context.Context, argUUID string, updated *model.BOOKINGTRIAGETRACKER) (result *model.BOOKINGTRIAGETRACKER, RowsAffected int64, err error) {
	sql := "UPDATE BOOKING_TRIAGE_TRACKER set PROVIDER_UUID = $1, TRIAGE_ORDER = $2, BOOKING_UUID = $3, TIME_OFFERED_IN_MINS = $4, TIMER_EXECUTION_ID = $5, TIMER_EXPIRES_AT = $6, TRACKING_TYPE = $7, STATUS = $8, STATUS_TLM = $9, TRIAGE_COMPLETED = $10, WHEN_CREATED = $11, TLM = $12 WHERE UUID = $13"
	fmt.Printf("sql: %s\n", sql)
	dbResult := DB.MustExecContext(ctx, sql, updated.PROVIDERUUID, updated.TRIAGEORDER, updated.BOOKINGUUID, updated.TIMEOFFEREDINMINS, updated.TIMEREXECUTIONID, updated.TIMEREXPIRESAT, updated.TRACKINGTYPE, updated.STATUS, updated.STATUSTLM, updated.TRIAGECOMPLETED, updated.WHENCREATED, updated.TLM, argUUID)
	id, err := dbResult.LastInsertId()
	fmt.Printf("LastInsertId: %d\n", id)
	rows, err := dbResult.RowsAffected()
	fmt.Printf("RowsAffected: %d\n", rows)

	updated.UUID = argUUID
	return updated, rows, err
}

// DeleteBOOKINGTRIAGETRACKER is a function to delete a single record from BOOKING_TRIAGE_TRACKER table in the agency_portal database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteBOOKINGTRIAGETRACKER(ctx context.Context, argUUID string) (rowsAffected int64, err error) {
	sql := "DELETE FROM BOOKING_TRIAGE_TRACKER where UUID = $1"
	result := DB.MustExecContext(ctx, sql, argUUID)
	return result.RowsAffected()
}
