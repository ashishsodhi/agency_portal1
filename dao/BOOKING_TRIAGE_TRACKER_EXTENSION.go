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
CREATE TABLE `BOOKING_TRIAGE_TRACKER_EXTENSION` (
  `UID` varchar(45) NOT NULL,
  `PROVIDER_URI` varchar(45) DEFAULT NULL,
  `BOOKING_URI` varchar(45) DEFAULT NULL,
  `PERMANENTLY_DECLINED` bit(1) DEFAULT NULL,
  `NOTES` varchar(10000) DEFAULT NULL,
  `SOURCE` enum('TRIAGE','CSR') DEFAULT NULL,
  `CSR_ID` varchar(45) DEFAULT NULL,
  `WHEN_CREATED` datetime DEFAULT NULL,
  `TLM` datetime DEFAULT NULL,
  PRIMARY KEY (`UID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8


PrimaryKeyNamesList: [UID]
PrimaryKeyNames    : UID
delSql             : DELETE FROM BOOKING_TRIAGE_TRACKER_EXTENSION where UID = $1
updateSql          : UPDATE BOOKING_TRIAGE_TRACKER_EXTENSION set PROVIDER_URI = $1, BOOKING_URI = $2, PERMANENTLY_DECLINED = $3, NOTES = $4, SOURCE = $5, CSR_ID = $6, WHEN_CREATED = $7, TLM = $8 WHERE UID = $9
insertSql          : INSERT INTO BOOKING_TRIAGE_TRACKER_EXTENSION ( UID,  PROVIDER_URI,  BOOKING_URI,  PERMANENTLY_DECLINED,  NOTES,  SOURCE,  CSR_ID,  WHEN_CREATED,  TLM) values ( $1, $2, $3, $4, $5, $6, $7, $8, $9 )
selectOneSql       : SELECT * FROM BOOKING_TRIAGE_TRACKER_EXTENSION WHERE UID=?
selectMultiSql     : SELECT * FROM BOOKING_TRIAGE_TRACKER_EXTENSION

*/

// GetAllBOOKINGTRIAGETRACKEREXTENSIONS is a function to get a slice of record(s) from BOOKING_TRIAGE_TRACKER_EXTENSION table in the agency_portal database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllBOOKINGTRIAGETRACKEREXTENSIONS(ctx context.Context, page, pagesize int64, order string) (bookingtriagetrackerextensions []*model.BOOKINGTRIAGETRACKEREXTENSION, totalRows int, err error) {
	sql := "SELECT * FROM BOOKING_TRIAGE_TRACKER_EXTENSION"

	if order != "" {
		sql = sql + " order by " + order
	}
	sql = fmt.Sprintf("%s LIMIT %d, %d", sql, page, pagesize)
	fmt.Printf("%s\n", sql)
	err = DB.SelectContext(ctx, &bookingtriagetrackerextensions, sql, page, pagesize)
	return bookingtriagetrackerextensions, len(bookingtriagetrackerextensions), err
}

// GetBOOKINGTRIAGETRACKEREXTENSION is a function to get a single record to BOOKING_TRIAGE_TRACKER_EXTENSION table in the agency_portal database
// error - ErrNotFound, db Find error
func GetBOOKINGTRIAGETRACKEREXTENSION(ctx context.Context, argUID string) (record *model.BOOKINGTRIAGETRACKEREXTENSION, err error) {
	sql := "SELECT * FROM BOOKING_TRIAGE_TRACKER_EXTENSION WHERE UID=?"
	record = &model.BOOKINGTRIAGETRACKEREXTENSION{}
	err = DB.GetContext(ctx, record, sql, argUID)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// AddBOOKINGTRIAGETRACKEREXTENSION is a function to add a single record to BOOKING_TRIAGE_TRACKER_EXTENSION table in the agency_portal database
// error - ErrInsertFailed, db save call failed
func AddBOOKINGTRIAGETRACKEREXTENSION(ctx context.Context, record *model.BOOKINGTRIAGETRACKEREXTENSION) (result *model.BOOKINGTRIAGETRACKEREXTENSION, RowsAffected int64, err error) {
	sql := "INSERT INTO BOOKING_TRIAGE_TRACKER_EXTENSION ( UID,  PROVIDER_URI,  BOOKING_URI,  PERMANENTLY_DECLINED,  NOTES,  SOURCE,  CSR_ID,  WHEN_CREATED,  TLM) values ( $1, $2, $3, $4, $5, $6, $7, $8, $9 )"
	dbResult := DB.MustExecContext(ctx, sql, record.PROVIDERURI, record.BOOKINGURI, record.PERMANENTLYDECLINED, record.NOTES, record.SOURCE, record.CSRID, record.WHENCREATED, record.TLM)
	id, err := dbResult.LastInsertId()
	fmt.Printf("LastInsertId: %d\n", id)
	rows, err := dbResult.RowsAffected()
	fmt.Printf("RowsAffected: %d\n", rows)
	record.UID = string(id)

	return record, rows, err
}

// UpdateBOOKINGTRIAGETRACKEREXTENSION is a function to update a single record from BOOKING_TRIAGE_TRACKER_EXTENSION table in the agency_portal database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateBOOKINGTRIAGETRACKEREXTENSION(ctx context.Context, argUID string, updated *model.BOOKINGTRIAGETRACKEREXTENSION) (result *model.BOOKINGTRIAGETRACKEREXTENSION, RowsAffected int64, err error) {
	sql := "UPDATE BOOKING_TRIAGE_TRACKER_EXTENSION set PROVIDER_URI = $1, BOOKING_URI = $2, PERMANENTLY_DECLINED = $3, NOTES = $4, SOURCE = $5, CSR_ID = $6, WHEN_CREATED = $7, TLM = $8 WHERE UID = $9"
	fmt.Printf("sql: %s\n", sql)
	dbResult := DB.MustExecContext(ctx, sql, updated.PROVIDERURI, updated.BOOKINGURI, updated.PERMANENTLYDECLINED, updated.NOTES, updated.SOURCE, updated.CSRID, updated.WHENCREATED, updated.TLM, argUID)
	id, err := dbResult.LastInsertId()
	fmt.Printf("LastInsertId: %d\n", id)
	rows, err := dbResult.RowsAffected()
	fmt.Printf("RowsAffected: %d\n", rows)

	updated.UID = argUID
	return updated, rows, err
}

// DeleteBOOKINGTRIAGETRACKEREXTENSION is a function to delete a single record from BOOKING_TRIAGE_TRACKER_EXTENSION table in the agency_portal database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteBOOKINGTRIAGETRACKEREXTENSION(ctx context.Context, argUID string) (rowsAffected int64, err error) {
	sql := "DELETE FROM BOOKING_TRIAGE_TRACKER_EXTENSION where UID = $1"
	result := DB.MustExecContext(ctx, sql, argUID)
	return result.RowsAffected()
}
