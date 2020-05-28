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


CREATE TABLE `TRIAGE_HISTORY` (
  `UID` varchar(45) NOT NULL,
  `TRIAGE_UID` varchar(45) DEFAULT NULL,
  `NAME` varchar(255) DEFAULT NULL,
  `ZIP` json DEFAULT NULL,
  `OPEN_JOB_TIMER` varchar(255) DEFAULT NULL,
  `TIME_TABLE` varchar(255) DEFAULT NULL,
  `CSR_ID` int(11) DEFAULT NULL,
  `WHEN_CREATED` datetime DEFAULT NULL,
  `TLM` datetime DEFAULT NULL,
  PRIMARY KEY (`UID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "name": "sqFEKvpRVtqBWIztKbMypvItz",    "csr_id": 47,    "when_created": "2218-05-16T07:26:27.305796383-05:00",    "tlm": "2288-09-28T00:33:40.088701636-05:00",    "uid": "hhEjxVojSlIBJSRJUhkyfTsRm",    "triage_uid": "MjKzTkAZlgEXGmXNiLgJxNqZl",    "zip": "oEtbICtsqJRccVFtFpsAtKsWA",    "open_job_timer": "ltvRhNhwGDDmsfBdYebYMMhoL",    "time_table": "TrjgEKAFyYUpYbVCYeBthHsFn"}



*/

// TRIAGEHISTORY struct is a row record of the TRIAGE_HISTORY table in the agency_portal database
type TRIAGEHISTORY struct {
	//[ 0] UID                                            varchar(45)          null: false  primary: true   auto: false  col: varchar         len: 45      default: []
	UID string `db:"UID" protobuf:"string,0,opt,name=uid"`
	//[ 1] TRIAGE_UID                                     varchar(45)          null: true   primary: false  auto: false  col: varchar         len: 45      default: []
	TRIAGEUID sql.NullString `db:"TRIAGE_UID" protobuf:"string,1,opt,name=triage_uid"`
	//[ 2] NAME                                           varchar(255)         null: true   primary: false  auto: false  col: varchar         len: 255     default: []
	NAME sql.NullString `db:"NAME" protobuf:"string,2,opt,name=name"`
	//[ 3] ZIP                                            json                 null: true   primary: false  auto: false  col: json            len: -1      default: []
	ZIP sql.NullString `db:"ZIP" protobuf:"string,3,opt,name=zip"`
	//[ 4] OPEN_JOB_TIMER                                 varchar(255)         null: true   primary: false  auto: false  col: varchar         len: 255     default: []
	OPENJOBTIMER sql.NullString `db:"OPEN_JOB_TIMER" protobuf:"string,4,opt,name=open_job_timer"`
	//[ 5] TIME_TABLE                                     varchar(255)         null: true   primary: false  auto: false  col: varchar         len: 255     default: []
	TIMETABLE sql.NullString `db:"TIME_TABLE" protobuf:"string,5,opt,name=time_table"`
	//[ 6] CSR_ID                                         int                  null: true   primary: false  auto: false  col: int             len: -1      default: []
	CSRID sql.NullInt64 `db:"CSR_ID" protobuf:"int32,6,opt,name=csr_id"`
	//[ 7] WHEN_CREATED                                   datetime             null: true   primary: false  auto: false  col: datetime        len: -1      default: []
	WHENCREATED time.Time `db:"WHEN_CREATED" protobuf:"uint64,7,opt,name=when_created"`
	//[ 8] TLM                                            datetime             null: true   primary: false  auto: false  col: datetime        len: -1      default: []
	TLM time.Time `db:"TLM" protobuf:"uint64,8,opt,name=tlm"`
}

// TableName sets the insert table name for this struct type
func (t *TRIAGEHISTORY) TableName() string {
	return "TRIAGE_HISTORY"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *TRIAGEHISTORY) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *TRIAGEHISTORY) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *TRIAGEHISTORY) Validate(action Action) error {
	return nil
}
