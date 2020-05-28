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
  `UID` varchar(45) NOT NULL,
  `TRIAGE_UID` varchar(45) NOT NULL,
  `ORDER` int(11) DEFAULT NULL,
  `STATUS` enum('ACTIVE','INACTIVE') NOT NULL,
  `STATUS_TLM` datetime NOT NULL,
  `ASSIGNMENT_TIMER` int(11) DEFAULT NULL,
  `WHEN_CREATED` datetime NOT NULL,
  `TLM` datetime NOT NULL,
  PRIMARY KEY (`UID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "triage_uid": "mcwHmfPUzlRtYfEnJJqBmBtGR",    "order": 74,    "status": "uZYBjlgVkalmBqeuKcjrYoIkH",    "status_tlm": "2275-12-03T17:24:42.73024445-05:00",    "assignment_timer": 52,    "when_created": "2143-03-19T19:18:57.778794625-05:00",    "tlm": "2302-11-18T06:18:24.092263288-05:00",    "uid": "KwgekpOaHRmCPoEAcEfHToBBN"}



*/

// TRIAGETIMETABLE struct is a row record of the TRIAGE_TIME_TABLE table in the agency_portal database
type TRIAGETIMETABLE struct {
	//[ 0] UID                                            varchar(45)          null: false  primary: true   auto: false  col: varchar         len: 45      default: []
	UID string `db:"UID" protobuf:"string,0,opt,name=uid"`
	//[ 1] TRIAGE_UID                                     varchar(45)          null: false  primary: false  auto: false  col: varchar         len: 45      default: []
	TRIAGEUID string `db:"TRIAGE_UID" protobuf:"string,1,opt,name=triage_uid"`
	//[ 2] ORDER                                          int                  null: true   primary: false  auto: false  col: int             len: -1      default: []
	ORDER sql.NullInt64 `db:"ORDER" protobuf:"int32,2,opt,name=order"`
	//[ 3] STATUS                                         char(8)              null: false  primary: false  auto: false  col: char            len: 8       default: []
	STATUS string `db:"STATUS" protobuf:"string,3,opt,name=status"`
	//[ 4] STATUS_TLM                                     datetime             null: false  primary: false  auto: false  col: datetime        len: -1      default: []
	STATUSTLM time.Time `db:"STATUS_TLM" protobuf:"uint64,4,opt,name=status_tlm"`
	//[ 5] ASSIGNMENT_TIMER                               int                  null: true   primary: false  auto: false  col: int             len: -1      default: []
	ASSIGNMENTTIMER sql.NullInt64 `db:"ASSIGNMENT_TIMER" protobuf:"int32,5,opt,name=assignment_timer"`
	//[ 6] WHEN_CREATED                                   datetime             null: false  primary: false  auto: false  col: datetime        len: -1      default: []
	WHENCREATED time.Time `db:"WHEN_CREATED" protobuf:"uint64,6,opt,name=when_created"`
	//[ 7] TLM                                            datetime             null: false  primary: false  auto: false  col: datetime        len: -1      default: []
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
