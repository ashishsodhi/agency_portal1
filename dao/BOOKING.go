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
CREATE TABLE `BOOKING` (
  `UUID` varchar(16) NOT NULL,
  `NAMESPACE` enum('CZEN') NOT NULL,
  `COLLECTION` enum('BACKUP_CARE_BOOKING') NOT NULL,
  `METADATA` json DEFAULT NULL,
  PRIMARY KEY (`UUID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8


PrimaryKeyNamesList: [UUID]
PrimaryKeyNames    : UUID
delSql             : DELETE FROM BOOKING where UUID = $1
updateSql          : UPDATE BOOKING set NAMESPACE = $1, COLLECTION = $2, METADATA = $3 WHERE UUID = $4
insertSql          : INSERT INTO BOOKING ( UUID,  NAMESPACE,  COLLECTION,  METADATA) values ( $1, $2, $3, $4 )
selectOneSql       : SELECT * FROM BOOKING WHERE UUID=?
selectMultiSql     : SELECT * FROM BOOKING

*/

// GetAllBOOKINGS is a function to get a slice of record(s) from BOOKING table in the agency_portal database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllBOOKINGS(ctx context.Context, page, pagesize int64, order string) (bookings []*model.BOOKING, totalRows int, err error) {
	sql := "SELECT * FROM BOOKING"

	if order != "" {
		sql = sql + " order by " + order
	}
	sql = fmt.Sprintf("%s LIMIT %d, %d", sql, page, pagesize)
	fmt.Printf("%s\n", sql)
	err = DB.SelectContext(ctx, &bookings, sql, page, pagesize)
	return bookings, len(bookings), err
}

// GetBOOKING is a function to get a single record to BOOKING table in the agency_portal database
// error - ErrNotFound, db Find error
func GetBOOKING(ctx context.Context, argUUID string) (record *model.BOOKING, err error) {
	sql := "SELECT * FROM BOOKING WHERE UUID=?"
	record = &model.BOOKING{}
	err = DB.GetContext(ctx, record, sql, argUUID)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// AddBOOKING is a function to add a single record to BOOKING table in the agency_portal database
// error - ErrInsertFailed, db save call failed
func AddBOOKING(ctx context.Context, record *model.BOOKING) (result *model.BOOKING, RowsAffected int64, err error) {
	sql := "INSERT INTO BOOKING ( UUID,  NAMESPACE,  COLLECTION,  METADATA) values ( $1, $2, $3, $4 )"
	dbResult := DB.MustExecContext(ctx, sql, record.NAMESPACE, record.COLLECTION, record.METADATA)
	id, err := dbResult.LastInsertId()
	fmt.Printf("LastInsertId: %d\n", id)
	rows, err := dbResult.RowsAffected()
	fmt.Printf("RowsAffected: %d\n", rows)
	record.UUID = string(id)

	return record, rows, err
}

// UpdateBOOKING is a function to update a single record from BOOKING table in the agency_portal database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateBOOKING(ctx context.Context, argUUID string, updated *model.BOOKING) (result *model.BOOKING, RowsAffected int64, err error) {
	sql := "UPDATE BOOKING set NAMESPACE = $1, COLLECTION = $2, METADATA = $3 WHERE UUID = $4"
	fmt.Printf("sql: %s\n", sql)
	dbResult := DB.MustExecContext(ctx, sql, updated.NAMESPACE, updated.COLLECTION, updated.METADATA, argUUID)
	id, err := dbResult.LastInsertId()
	fmt.Printf("LastInsertId: %d\n", id)
	rows, err := dbResult.RowsAffected()
	fmt.Printf("RowsAffected: %d\n", rows)

	updated.UUID = argUUID
	return updated, rows, err
}

// DeleteBOOKING is a function to delete a single record from BOOKING table in the agency_portal database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteBOOKING(ctx context.Context, argUUID string) (rowsAffected int64, err error) {
	sql := "DELETE FROM BOOKING where UUID = $1"
	result := DB.MustExecContext(ctx, sql, argUUID)
	return result.RowsAffected()
}
