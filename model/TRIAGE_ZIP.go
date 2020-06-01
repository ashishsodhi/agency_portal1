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
  `ZIP` char(5) NOT NULL,
  `TRIAGE_UUID` varbinary(16) NOT NULL,
  `WHEN_CREATED` datetime NOT NULL,
  `TLM` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`ZIP`),
  KEY `IX_TLM` (`TLM`),
  KEY `FK_TRIAGE_ZIP_TRIAGE_UUID` (`TRIAGE_UUID`),
  CONSTRAINT `FK_TRIAGE_ZIP_TRIAGE_UUID` FOREIGN KEY (`TRIAGE_UUID`) REFERENCES `TRIAGE` (`UUID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "when_created": "2179-06-17T13:35:29.469736263-05:00",    "tlm": "2274-05-03T10:08:28.531619631-05:00",    "zip": "bIDOGCnnXYRiUbzmRJlmiCUAO",    "triage_uuid": "KhUSQig9OTNWX0EeE0oFS0wCWRcdSBU="}



*/

// TRIAGEZIP struct is a row record of the TRIAGE_ZIP table in the agency_portal database
type TRIAGEZIP struct {
	//[ 0] ZIP                                            char(5)              null: false  primary: true   auto: false  col: char            len: 5       default: []
	ZIP string `db:"ZIP" protobuf:"string,0,opt,name=zip"`
	//[ 1] TRIAGE_UUID                                    varbinary            null: false  primary: false  auto: false  col: varbinary       len: -1      default: []
	TRIAGEUUID []byte `db:"TRIAGE_UUID" protobuf:"bytes,1,opt,name=triage_uuid"`
	//[ 2] WHEN_CREATED                                   datetime             null: false  primary: false  auto: false  col: datetime        len: -1      default: []
	WHENCREATED time.Time `db:"WHEN_CREATED" protobuf:"uint64,2,opt,name=when_created"`
	//[ 3] TLM                                            timestamp            null: false  primary: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
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
