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


CREATE TABLE `TRIAGE_TIME_TABLE` (
  `UUID` varchar(16) NOT NULL,
  `TRIAGE_UUID` varchar(16) NOT NULL,
  `ORDER` smallint(5) unsigned NOT NULL,
  `STATUS` enum('ACTIVE','INACTIVE') NOT NULL,
  `STATUS_TLM` datetime NOT NULL,
  `ASSIGNMENT_TIMER` json NOT NULL,
  `WHEN_CREATED` datetime NOT NULL,
  `TLM` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`UUID`),
  UNIQUE KEY `UNQ_TRIAGE_UUID_ORDER` (`TRIAGE_UUID`,`ORDER`),
  KEY `IX_TLM` (`TLM`),
  KEY `IX_TRIAGE_UUID_STATUS` (`TRIAGE_UUID`,`STATUS`),
  CONSTRAINT `FK_TRIAGE_TIME_TABLE_TRAIGE_UUID` FOREIGN KEY (`TRIAGE_UUID`) REFERENCES `TRIAGE` (`UUID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "triage_uuid": "PlJDYnoaxtfVxDvCMpigJtgNW",    "order": 25,    "status": "gMRLgnlGalMHcHHdFpEbGtZSD",    "status_tlm": "2134-01-20T20:05:30.797147447-05:00",    "assignment_timer": "ZAdHvIrcrVzhJAvvbjWfBXoSP",    "when_created": "2030-06-22T13:21:18.667452179-04:00",    "tlm": "2139-10-06T09:13:05.442641661-05:00",    "uuid": "rgjfTTlyOMsBuVujTRQUhOWDZ"}



*/

// TRIAGETIMETABLE struct is a row record of the TRIAGE_TIME_TABLE table in the agency_portal database
type TRIAGETIMETABLE struct {
	//[ 0] UUID                                           varchar(16)          null: false  primary: true   auto: false  col: varchar         len: 16      default: []
	UUID string `db:"UUID" protobuf:"string,0,opt,name=uuid"`
	//[ 1] TRIAGE_UUID                                    varchar(16)          null: false  primary: false  auto: false  col: varchar         len: 16      default: []
	TRIAGEUUID string `db:"TRIAGE_UUID" protobuf:"string,1,opt,name=triage_uuid"`
	//[ 2] ORDER                                          smallint             null: false  primary: false  auto: false  col: smallint        len: -1      default: []
	ORDER int `db:"ORDER" protobuf:"int16,2,opt,name=order"`
	//[ 3] STATUS                                         char(8)              null: false  primary: false  auto: false  col: char            len: 8       default: []
	STATUS string `db:"STATUS" protobuf:"string,3,opt,name=status"`
	//[ 4] STATUS_TLM                                     datetime             null: false  primary: false  auto: false  col: datetime        len: -1      default: []
	STATUSTLM time.Time `db:"STATUS_TLM" protobuf:"uint64,4,opt,name=status_tlm"`
	//[ 5] ASSIGNMENT_TIMER                               json                 null: false  primary: false  auto: false  col: json            len: -1      default: []
	ASSIGNMENTTIMER string `db:"ASSIGNMENT_TIMER" protobuf:"string,5,opt,name=assignment_timer"`
	//[ 6] WHEN_CREATED                                   datetime             null: false  primary: false  auto: false  col: datetime        len: -1      default: []
	WHENCREATED time.Time `db:"WHEN_CREATED" protobuf:"uint64,6,opt,name=when_created"`
	//[ 7] TLM                                            timestamp            null: false  primary: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	TLM time.Time `db:"TLM" protobuf:"uint64,7,opt,name=tlm"`
}

// TableName sets the insert table name for this struct type
func (t *TRIAGETIMETABLE) TableName() string {
	return "TRIAGE_TIME_TABLE"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *TRIAGETIMETABLE) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *TRIAGETIMETABLE) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *TRIAGETIMETABLE) Validate(action Action) error {
	return nil
}
