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


CREATE TABLE `TRIAGE` (
  `UUID` varchar(16) NOT NULL,
  `NAME` varchar(255) NOT NULL,
  `ACCEPTANCE_TIMER` json NOT NULL,
  `WHEN_CREATED` datetime NOT NULL,
  `TLM` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`UUID`),
  UNIQUE KEY `UNQ_NAME` (`NAME`),
  KEY `IX_TLM` (`TLM`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "when_created": "2150-04-27T09:23:02.941241833-05:00",    "tlm": "2193-08-19T03:55:00.595889911-05:00",    "uuid": "QkmXhsoigdYecDmRIncBADpJQ",    "name": "TEfQxouylNcYjxkXYYLAqFjsR",    "acceptance_timer": "KrDuhvklEWeYlLMpqKsemycLb"}



*/

// TRIAGE struct is a row record of the TRIAGE table in the agency_portal database
type TRIAGE struct {
	//[ 0] UUID                                           varchar(16)          null: false  primary: true   auto: false  col: varchar         len: 16      default: []
	UUID string `db:"UUID" protobuf:"string,0,opt,name=uuid"`
	//[ 1] NAME                                           varchar(255)         null: false  primary: false  auto: false  col: varchar         len: 255     default: []
	NAME string `db:"NAME" protobuf:"string,1,opt,name=name"`
	//[ 2] ACCEPTANCE_TIMER                               json                 null: false  primary: false  auto: false  col: json            len: -1      default: []
	ACCEPTANCETIMER string `db:"ACCEPTANCE_TIMER" protobuf:"string,2,opt,name=acceptance_timer"`
	//[ 3] WHEN_CREATED                                   datetime             null: false  primary: false  auto: false  col: datetime        len: -1      default: []
	WHENCREATED time.Time `db:"WHEN_CREATED" protobuf:"uint64,3,opt,name=when_created"`
	//[ 4] TLM                                            timestamp            null: false  primary: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	TLM time.Time `db:"TLM" protobuf:"uint64,4,opt,name=tlm"`
}

// TableName sets the insert table name for this struct type
func (t *TRIAGE) TableName() string {
	return "TRIAGE"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *TRIAGE) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *TRIAGE) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *TRIAGE) Validate(action Action) error {
	return nil
}
