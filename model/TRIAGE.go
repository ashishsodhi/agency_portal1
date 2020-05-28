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
  `UID` varchar(45) NOT NULL,
  `NAME` varchar(45) NOT NULL,
  `ACCEPTANCE_TIMER` int(11) DEFAULT NULL,
  `WHEN_CREATED` datetime NOT NULL,
  `TLM` datetime NOT NULL,
  PRIMARY KEY (`UID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "when_created": "2147-08-13T12:13:45.760941899-05:00",    "tlm": "2250-09-17T00:44:06.954452463-05:00",    "uid": "vzFJtNSLnELbbqQpRycBIVzux",    "name": "DZjqVmSnPQgtotywcdDUuZnmP",    "acceptance_timer": 48}



*/

// TRIAGE struct is a row record of the TRIAGE table in the agency_portal database
type TRIAGE struct {
	//[ 0] UID                                            varchar(45)          null: false  primary: true   auto: false  col: varchar         len: 45      default: []
	UID string `db:"UID" protobuf:"string,0,opt,name=uid"`
	//[ 1] NAME                                           varchar(45)          null: false  primary: false  auto: false  col: varchar         len: 45      default: []
	NAME string `db:"NAME" protobuf:"string,1,opt,name=name"`
	//[ 2] ACCEPTANCE_TIMER                               int                  null: true   primary: false  auto: false  col: int             len: -1      default: []
	ACCEPTANCETIMER sql.NullInt64 `db:"ACCEPTANCE_TIMER" protobuf:"int32,2,opt,name=acceptance_timer"`
	//[ 3] WHEN_CREATED                                   datetime             null: false  primary: false  auto: false  col: datetime        len: -1      default: []
	WHENCREATED time.Time `db:"WHEN_CREATED" protobuf:"uint64,3,opt,name=when_created"`
	//[ 4] TLM                                            datetime             null: false  primary: false  auto: false  col: datetime        len: -1      default: []
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
