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
CREATE TABLE `BOOKING_TRIAGE_TRACKER_EXTENSION` (
  `UUID` varchar(16) NOT NULL,
  `PROVIDER_UUID` varchar(16) NOT NULL,
  `BOOKING_UUID` varchar(16) NOT NULL,
  `PERMANENTLY_DECLINED` bit(1) NOT NULL DEFAULT b'0',
  `NOTES` varchar(10000) DEFAULT NULL,
  `SOURCE` enum('TRIAGE','CSR') NOT NULL DEFAULT 'TRIAGE',
  `CSR_UUID` varbinary(16) NOT NULL,
  `WHEN_CREATED` datetime NOT NULL,
  `TLM` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`UUID`),
  KEY `IX_TLM` (`TLM`),
  KEY `IX_PROVIDER_UUID_PERMANENTLY_DECLINED` (`PROVIDER_UUID`,`PERMANENTLY_DECLINED`),
  KEY `BOOKING_TRIAGE_TRACKER_EXTENSION_BOOKING_UUID_fk` (`BOOKING_UUID`),
  CONSTRAINT `BOOKING_TRIAGE_TRACKER_EXTENSION_BOOKING_UUID_fk` FOREIGN KEY (`BOOKING_UUID`) REFERENCES `BOOKING` (`UUID`),
  CONSTRAINT `BOOKING_TRIAGE_TRACKER_EXTENSION_PROVIDER_UUID` FOREIGN KEY (`PROVIDER_UUID`) REFERENCES `PROVIDER` (`UUID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8


PrimaryKeyNamesList: [UUID]
PrimaryKeyNames    : UUID
delSql             : DELETE FROM BOOKING_TRIAGE_TRACKER_EXTENSION where UUID = $1
updateSql          : UPDATE BOOKING_TRIAGE_TRACKER_EXTENSION set PROVIDER_UUID = $1, BOOKING_UUID = $2, PERMANENTLY_DECLINED = $3, NOTES = $4, SOURCE = $5, CSR_UUID = $6, WHEN_CREATED = $7, TLM = $8 WHERE UUID = $9
insertSql          : INSERT INTO BOOKING_TRIAGE_TRACKER_EXTENSION ( UUID,  PROVIDER_UUID,  BOOKING_UUID,  PERMANENTLY_DECLINED,  NOTES,  SOURCE,  CSR_UUID,  WHEN_CREATED,  TLM) values ( $1, $2, $3, $4, $5, $6, $7, $8, $9 )
selectOneSql       : SELECT * FROM BOOKING_TRIAGE_TRACKER_EXTENSION WHERE UUID=?
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
func GetBOOKINGTRIAGETRACKEREXTENSION(ctx context.Context, argUUID string) (record *model.BOOKINGTRIAGETRACKEREXTENSION, err error) {
	sql := "SELECT * FROM BOOKING_TRIAGE_TRACKER_EXTENSION WHERE UUID=?"
	record = &model.BOOKINGTRIAGETRACKEREXTENSION{}
	err = DB.GetContext(ctx, record, sql, argUUID)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// AddBOOKINGTRIAGETRACKEREXTENSION is a function to add a single record to BOOKING_TRIAGE_TRACKER_EXTENSION table in the agency_portal database
// error - ErrInsertFailed, db save call failed
func AddBOOKINGTRIAGETRACKEREXTENSION(ctx context.Context, record *model.BOOKINGTRIAGETRACKEREXTENSION) (result *model.BOOKINGTRIAGETRACKEREXTENSION, RowsAffected int64, err error) {
	sql := "INSERT INTO BOOKING_TRIAGE_TRACKER_EXTENSION ( UUID,  PROVIDER_UUID,  BOOKING_UUID,  PERMANENTLY_DECLINED,  NOTES,  SOURCE,  CSR_UUID,  WHEN_CREATED,  TLM) values ( $1, $2, $3, $4, $5, $6, $7, $8, $9 )"
	dbResult := DB.MustExecContext(ctx, sql, record.PROVIDERUUID, record.BOOKINGUUID, record.PERMANENTLYDECLINED, record.NOTES, record.SOURCE, record.CSRUUID, record.WHENCREATED, record.TLM)
	id, err := dbResult.LastInsertId()
	fmt.Printf("LastInsertId: %d\n", id)
	rows, err := dbResult.RowsAffected()
	fmt.Printf("RowsAffected: %d\n", rows)
	record.UUID = string(id)

	return record, rows, err
}

// UpdateBOOKINGTRIAGETRACKEREXTENSION is a function to update a single record from BOOKING_TRIAGE_TRACKER_EXTENSION table in the agency_portal database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateBOOKINGTRIAGETRACKEREXTENSION(ctx context.Context, argUUID string, updated *model.BOOKINGTRIAGETRACKEREXTENSION) (result *model.BOOKINGTRIAGETRACKEREXTENSION, RowsAffected int64, err error) {
	sql := "UPDATE BOOKING_TRIAGE_TRACKER_EXTENSION set PROVIDER_UUID = $1, BOOKING_UUID = $2, PERMANENTLY_DECLINED = $3, NOTES = $4, SOURCE = $5, CSR_UUID = $6, WHEN_CREATED = $7, TLM = $8 WHERE UUID = $9"
	fmt.Printf("sql: %s\n", sql)
	dbResult := DB.MustExecContext(ctx, sql, updated.PROVIDERUUID, updated.BOOKINGUUID, updated.PERMANENTLYDECLINED, updated.NOTES, updated.SOURCE, updated.CSRUUID, updated.WHENCREATED, updated.TLM, argUUID)
	id, err := dbResult.LastInsertId()
	fmt.Printf("LastInsertId: %d\n", id)
	rows, err := dbResult.RowsAffected()
	fmt.Printf("RowsAffected: %d\n", rows)

	updated.UUID = argUUID
	return updated, rows, err
}

// DeleteBOOKINGTRIAGETRACKEREXTENSION is a function to delete a single record from BOOKING_TRIAGE_TRACKER_EXTENSION table in the agency_portal database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteBOOKINGTRIAGETRACKEREXTENSION(ctx context.Context, argUUID string) (rowsAffected int64, err error) {
	sql := "DELETE FROM BOOKING_TRIAGE_TRACKER_EXTENSION where UUID = $1"
	result := DB.MustExecContext(ctx, sql, argUUID)
	return result.RowsAffected()
}
