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
CREATE TABLE `BOOKING_TRIAGE_HISTORY` (
  `UUID` varchar(16) NOT NULL,
  `PROVIDER_UUID` varchar(16) NOT NULL,
  `BOOKING_UUID` varchar(16) NOT NULL,
  `ACTION` enum('PENDING_ACCEPTANCE','ACCEPTED','DECLINED_TRIAGE','TRIAGE_TIMED_OUT','ASSIGNMENT_TIMED_OUT','DECLINED_PERMANENTLY','FULFILLED') NOT NULL,
  `USER_UUID` varbinary(16) NOT NULL,
  `CSR_UUID` varbinary(16) NOT NULL,
  `WHEN_CREATED` datetime NOT NULL,
  `TLM` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`UUID`),
  KEY `IX_TLM` (`TLM`),
  KEY `BOOKING_TRIAGE_HISTORY_PROVIDER_UUID_fk` (`PROVIDER_UUID`),
  KEY `FK_BOOKING_TRIAGE_HISTORY_BOOKING_TRIAGE_TRACKER` (`BOOKING_UUID`),
  CONSTRAINT `BOOKING_TRIAGE_HISTORY_PROVIDER_UUID_fk` FOREIGN KEY (`PROVIDER_UUID`) REFERENCES `PROVIDER` (`UUID`),
  CONSTRAINT `FK_BOOKING_TRIAGE_HISTORY_BOOKING_TRIAGE_TRACKER` FOREIGN KEY (`BOOKING_UUID`) REFERENCES `BOOKING` (`UUID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8


PrimaryKeyNamesList: [UUID]
PrimaryKeyNames    : UUID
delSql             : DELETE FROM BOOKING_TRIAGE_HISTORY where UUID = $1
updateSql          : UPDATE BOOKING_TRIAGE_HISTORY set PROVIDER_UUID = $1, BOOKING_UUID = $2, ACTION = $3, USER_UUID = $4, CSR_UUID = $5, WHEN_CREATED = $6, TLM = $7 WHERE UUID = $8
insertSql          : INSERT INTO BOOKING_TRIAGE_HISTORY ( UUID,  PROVIDER_UUID,  BOOKING_UUID,  ACTION,  USER_UUID,  CSR_UUID,  WHEN_CREATED,  TLM) values ( $1, $2, $3, $4, $5, $6, $7, $8 )
selectOneSql       : SELECT * FROM BOOKING_TRIAGE_HISTORY WHERE UUID=?
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
func GetBOOKINGTRIAGEHISTORY(ctx context.Context, argUUID string) (record *model.BOOKINGTRIAGEHISTORY, err error) {
	sql := "SELECT * FROM BOOKING_TRIAGE_HISTORY WHERE UUID=?"
	record = &model.BOOKINGTRIAGEHISTORY{}
	err = DB.GetContext(ctx, record, sql, argUUID)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// AddBOOKINGTRIAGEHISTORY is a function to add a single record to BOOKING_TRIAGE_HISTORY table in the agency_portal database
// error - ErrInsertFailed, db save call failed
func AddBOOKINGTRIAGEHISTORY(ctx context.Context, record *model.BOOKINGTRIAGEHISTORY) (result *model.BOOKINGTRIAGEHISTORY, RowsAffected int64, err error) {
	sql := "INSERT INTO BOOKING_TRIAGE_HISTORY ( UUID,  PROVIDER_UUID,  BOOKING_UUID,  ACTION,  USER_UUID,  CSR_UUID,  WHEN_CREATED,  TLM) values ( $1, $2, $3, $4, $5, $6, $7, $8 )"
	dbResult := DB.MustExecContext(ctx, sql, record.PROVIDERUUID, record.BOOKINGUUID, record.ACTION, record.USERUUID, record.CSRUUID, record.WHENCREATED, record.TLM)
	id, err := dbResult.LastInsertId()
	fmt.Printf("LastInsertId: %d\n", id)
	rows, err := dbResult.RowsAffected()
	fmt.Printf("RowsAffected: %d\n", rows)
	record.UUID = string(id)

	return record, rows, err
}

// UpdateBOOKINGTRIAGEHISTORY is a function to update a single record from BOOKING_TRIAGE_HISTORY table in the agency_portal database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateBOOKINGTRIAGEHISTORY(ctx context.Context, argUUID string, updated *model.BOOKINGTRIAGEHISTORY) (result *model.BOOKINGTRIAGEHISTORY, RowsAffected int64, err error) {
	sql := "UPDATE BOOKING_TRIAGE_HISTORY set PROVIDER_UUID = $1, BOOKING_UUID = $2, ACTION = $3, USER_UUID = $4, CSR_UUID = $5, WHEN_CREATED = $6, TLM = $7 WHERE UUID = $8"
	fmt.Printf("sql: %s\n", sql)
	dbResult := DB.MustExecContext(ctx, sql, updated.PROVIDERUUID, updated.BOOKINGUUID, updated.ACTION, updated.USERUUID, updated.CSRUUID, updated.WHENCREATED, updated.TLM, argUUID)
	id, err := dbResult.LastInsertId()
	fmt.Printf("LastInsertId: %d\n", id)
	rows, err := dbResult.RowsAffected()
	fmt.Printf("RowsAffected: %d\n", rows)

	updated.UUID = argUUID
	return updated, rows, err
}

// DeleteBOOKINGTRIAGEHISTORY is a function to delete a single record from BOOKING_TRIAGE_HISTORY table in the agency_portal database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteBOOKINGTRIAGEHISTORY(ctx context.Context, argUUID string) (rowsAffected int64, err error) {
	sql := "DELETE FROM BOOKING_TRIAGE_HISTORY where UUID = $1"
	result := DB.MustExecContext(ctx, sql, argUUID)
	return result.RowsAffected()
}
