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


CREATE TABLE `TRIAGE_ZIP` (
  `ZIP` varchar(45) NOT NULL,
  `TRIAGE_UID` varchar(45) DEFAULT NULL,
  `WHEN_CREATED` datetime DEFAULT NULL,
  `TLM` datetime DEFAULT NULL,
  PRIMARY KEY (`ZIP`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "zip": "UjGpEFIeFRxPqimfVtcOYpvki",    "triage_uid": "mVfdGfMHseGMwVGSgriLBmnzO",    "when_created": "2102-02-23T06:54:25.723288882-05:00",    "tlm": "2146-09-30T18:04:55.349302523-05:00"}



*/

// TRIAGEZIP struct is a row record of the TRIAGE_ZIP table in the agency_portal database
type TRIAGEZIP struct {
	//[ 0] ZIP                                            varchar(45)          null: false  primary: true   auto: false  col: varchar         len: 45      default: []
	ZIP string `db:"ZIP" protobuf:"string,0,opt,name=zip"`
	//[ 1] TRIAGE_UID                                     varchar(45)          null: true   primary: false  auto: false  col: varchar         len: 45      default: []
	TRIAGEUID sql.NullString `db:"TRIAGE_UID" protobuf:"string,1,opt,name=triage_uid"`
	//[ 2] WHEN_CREATED                                   datetime             null: true   primary: false  auto: false  col: datetime        len: -1      default: []
	WHENCREATED time.Time `db:"WHEN_CREATED" protobuf:"uint64,2,opt,name=when_created"`
	//[ 3] TLM                                            datetime             null: true   primary: false  auto: false  col: datetime        len: -1      default: []
	TLM time.Time `db:"TLM" protobuf:"uint64,3,opt,name=tlm"`
}

// TableName sets the insert table name for this struct type
func (t *TRIAGEZIP) TableName() string {
	return "TRIAGE_ZIP"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *TRIAGEZIP) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *TRIAGEZIP) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *TRIAGEZIP) Validate(action Action) error {
	return nil
}
