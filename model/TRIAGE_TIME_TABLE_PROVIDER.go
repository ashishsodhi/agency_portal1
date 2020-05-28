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
  `TRIAGE_TIME_TABLE_ID` varchar(45) NOT NULL,
  `PROVIDER_URI` varchar(45) DEFAULT NULL,
  `STATUS` enum('ACTIVE','INACTIVE') NOT NULL,
  `STATUS_TLM` datetime DEFAULT NULL,
  `WHEN_CREATED` datetime NOT NULL,
  `TLM` datetime NOT NULL,
  PRIMARY KEY (`TRIAGE_TIME_TABLE_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "when_created": "2280-03-05T12:06:19.999753995-05:00",    "tlm": "2247-08-28T13:23:48.220637041-05:00",    "triage_time_table_id": "yloiKLAjUSTWyqzyIsCjqfcrt",    "provider_uri": "XJrMvUkrXuSflWCyTZJgXnbGR",    "status": "reBRfIUoqkrlGUDDSrPkrIgmn",    "status_tlm": "2135-11-26T02:12:00.559921893-05:00"}



*/

// TRIAGETIMETABLEPROVIDER struct is a row record of the TRIAGE_TIME_TABLE_PROVIDER table in the agency_portal database
type TRIAGETIMETABLEPROVIDER struct {
	//[ 0] TRIAGE_TIME_TABLE_ID                           varchar(45)          null: false  primary: true   auto: false  col: varchar         len: 45      default: []
	TRIAGETIMETABLEID string `db:"TRIAGE_TIME_TABLE_ID" protobuf:"string,0,opt,name=triage_time_table_id"`
	//[ 1] PROVIDER_URI                                   varchar(45)          null: true   primary: false  auto: false  col: varchar         len: 45      default: []
	PROVIDERURI sql.NullString `db:"PROVIDER_URI" protobuf:"string,1,opt,name=provider_uri"`
	//[ 2] STATUS                                         char(8)              null: false  primary: false  auto: false  col: char            len: 8       default: []
	STATUS string `db:"STATUS" protobuf:"string,2,opt,name=status"`
	//[ 3] STATUS_TLM                                     datetime             null: true   primary: false  auto: false  col: datetime        len: -1      default: []
	STATUSTLM time.Time `db:"STATUS_TLM" protobuf:"uint64,3,opt,name=status_tlm"`
	//[ 4] WHEN_CREATED                                   datetime             null: false  primary: false  auto: false  col: datetime        len: -1      default: []
	WHENCREATED time.Time `db:"WHEN_CREATED" protobuf:"uint64,4,opt,name=when_created"`
	//[ 5] TLM                                            datetime             null: false  primary: false  auto: false  col: datetime        len: -1      default: []
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
