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


CREATE TABLE `TRIAGE_TIME_TABLE_PROVIDER` (
  `TRIAGE_TIME_TABLE_UUID` varchar(16) NOT NULL,
  `PROVIDER_UUID` varchar(16) NOT NULL,
  `STATUS` enum('ACTIVE','INACTIVE') NOT NULL,
  `STATUS_TLM` datetime NOT NULL,
  `WHEN_CREATED` datetime NOT NULL,
  `TLM` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`TRIAGE_TIME_TABLE_UUID`),
  KEY `IX_TLM` (`TLM`),
  KEY `IX_TRIAGE_TIME_TABLE_UUID_STATUS` (`TRIAGE_TIME_TABLE_UUID`,`STATUS`),
  KEY `FK_TRIAGE_TIME_TABLE_PROVIDER_UUID` (`PROVIDER_UUID`),
  CONSTRAINT `FK_TRIAGE_TIME_TABLE_PROVIDER_TIME_TABLE_UUID` FOREIGN KEY (`TRIAGE_TIME_TABLE_UUID`) REFERENCES `TRIAGE_TIME_TABLE` (`UUID`),
  CONSTRAINT `FK_TRIAGE_TIME_TABLE_PROVIDER_UUID` FOREIGN KEY (`PROVIDER_UUID`) REFERENCES `PROVIDER` (`UUID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "status_tlm": "2035-02-19T00:46:42.315556406-05:00",    "when_created": "2189-01-25T08:11:15.71116927-05:00",    "tlm": "2154-11-30T17:28:09.816641451-05:00",    "triage_time_table_uuid": "cIeByaPBbJlqSYWSdfeCzkPmB",    "provider_uuid": "wPPuHrRGwGlWzCUSbWZesPYhY",    "status": "cViYLWcWEiTTCDgbMWVYOaDZu"}



*/

// TRIAGETIMETABLEPROVIDER struct is a row record of the TRIAGE_TIME_TABLE_PROVIDER table in the agency_portal database
type TRIAGETIMETABLEPROVIDER struct {
	//[ 0] TRIAGE_TIME_TABLE_UUID                         varchar(16)          null: false  primary: true   auto: false  col: varchar         len: 16      default: []
	TRIAGETIMETABLEUUID string `db:"TRIAGE_TIME_TABLE_UUID" protobuf:"string,0,opt,name=triage_time_table_uuid"`
	//[ 1] PROVIDER_UUID                                  varchar(16)          null: false  primary: false  auto: false  col: varchar         len: 16      default: []
	PROVIDERUUID string `db:"PROVIDER_UUID" protobuf:"string,1,opt,name=provider_uuid"`
	//[ 2] STATUS                                         char(8)              null: false  primary: false  auto: false  col: char            len: 8       default: []
	STATUS string `db:"STATUS" protobuf:"string,2,opt,name=status"`
	//[ 3] STATUS_TLM                                     datetime             null: false  primary: false  auto: false  col: datetime        len: -1      default: []
	STATUSTLM time.Time `db:"STATUS_TLM" protobuf:"uint64,3,opt,name=status_tlm"`
	//[ 4] WHEN_CREATED                                   datetime             null: false  primary: false  auto: false  col: datetime        len: -1      default: []
	WHENCREATED time.Time `db:"WHEN_CREATED" protobuf:"uint64,4,opt,name=when_created"`
	//[ 5] TLM                                            timestamp            null: false  primary: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	TLM time.Time `db:"TLM" protobuf:"uint64,5,opt,name=tlm"`
}

// TableName sets the insert table name for this struct type
func (t *TRIAGETIMETABLEPROVIDER) TableName() string {
	return "TRIAGE_TIME_TABLE_PROVIDER"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *TRIAGETIMETABLEPROVIDER) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *TRIAGETIMETABLEPROVIDER) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *TRIAGETIMETABLEPROVIDER) Validate(action Action) error {
	return nil
}
