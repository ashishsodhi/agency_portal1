package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
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

JSON Sample
-------------------------------------
{    "collection": "flHiCSnjMvfRRVftkBMOWKOTw",    "metadata": "TvZhctowLFCwRqOzQvWshIWkV",    "uuid": "TGbLjvElkjDXOqulNZDWjIbJE",    "namespace": "LENsNeCZAbMgQhlybjjULQYRL"}



*/

// BOOKING struct is a row record of the BOOKING table in the agency_portal database
type BOOKING struct {
	//[ 0] UUID                                           varchar(16)          null: false  primary: true   auto: false  col: varchar         len: 16      default: []
	UUID string `db:"UUID" protobuf:"string,0,opt,name=uuid"`
	//[ 1] NAMESPACE                                      char(4)              null: false  primary: false  auto: false  col: char            len: 4       default: []
	NAMESPACE string `db:"NAMESPACE" protobuf:"string,1,opt,name=namespace"`
	//[ 2] COLLECTION                                     char(19)             null: false  primary: false  auto: false  col: char            len: 19      default: []
	COLLECTION string `db:"COLLECTION" protobuf:"string,2,opt,name=collection"`
	//[ 3] METADATA                                       json                 null: true   primary: false  auto: false  col: json            len: -1      default: []
	METADATA sql.NullString `db:"METADATA" protobuf:"string,3,opt,name=metadata"`
}

// TableName sets the insert table name for this struct type
func (b *BOOKING) TableName() string {
	return "BOOKING"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *BOOKING) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *BOOKING) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *BOOKING) Validate(action Action) error {
	return nil
}
