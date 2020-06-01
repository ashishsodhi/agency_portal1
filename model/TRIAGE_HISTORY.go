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
  `UUID` varchar(16) NOT NULL,
  `TRIAGE_UUID` varchar(16) NOT NULL,
  `NAME` varchar(255) NOT NULL,
  `ZIP` json NOT NULL,
  `OPEN_JOB_TIMER` json NOT NULL,
  `TIME_TABLE` json NOT NULL,
  `CSR_UUID` varbinary(16) NOT NULL,
  `WHEN_CREATED` datetime NOT NULL,
  `TLM` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`UUID`),
  KEY `IX_TLM` (`TLM`),
  KEY `FK_TRIAGE_HISTORY_TRIAGE_UUID` (`TRIAGE_UUID`),
  CONSTRAINT `FK_TRIAGE_HISTORY_TRIAGE_UUID` FOREIGN KEY (`TRIAGE_UUID`) REFERENCES `TRIAGE` (`UUID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "open_job_timer": "chHksWkoaPzkRXlSkPdIQeQBy",    "time_table": "IMdUtRojtDGFVapxrTnOQPvjP",    "when_created": "2156-02-01T09:15:06.084351346-05:00",    "uuid": "vJCAeDdBdqwSbFkeaqQPYXtSd",    "name": "VlfhvcCwORbMvZYEaVXeruVzT",    "zip": "qPCwvlycdoxcFgyBiKCRAWNWs",    "triage_uuid": "nAAywmLwKZNfgcJQEJfVbjhti",    "csr_uuid": "URAPMloYXzM1BgYPKiYhDjUhRhgnWA0PQSpNVEMDGD1WNUckUg4F",    "tlm": "2279-08-25T09:30:43.05450629-05:00"}



*/

// TRIAGEHISTORY struct is a row record of the TRIAGE_HISTORY table in the agency_portal database
type TRIAGEHISTORY struct {
	//[ 0] UUID                                           varchar(16)          null: false  primary: true   auto: false  col: varchar         len: 16      default: []
	UUID string `db:"UUID" protobuf:"string,0,opt,name=uuid"`
	//[ 1] TRIAGE_UUID                                    varchar(16)          null: false  primary: false  auto: false  col: varchar         len: 16      default: []
	TRIAGEUUID string `db:"TRIAGE_UUID" protobuf:"string,1,opt,name=triage_uuid"`
	//[ 2] NAME                                           varchar(255)         null: false  primary: false  auto: false  col: varchar         len: 255     default: []
	NAME string `db:"NAME" protobuf:"string,2,opt,name=name"`
	//[ 3] ZIP                                            json                 null: false  primary: false  auto: false  col: json            len: -1      default: []
	ZIP string `db:"ZIP" protobuf:"string,3,opt,name=zip"`
	//[ 4] OPEN_JOB_TIMER                                 json                 null: false  primary: false  auto: false  col: json            len: -1      default: []
	OPENJOBTIMER string `db:"OPEN_JOB_TIMER" protobuf:"string,4,opt,name=open_job_timer"`
	//[ 5] TIME_TABLE                                     json                 null: false  primary: false  auto: false  col: json            len: -1      default: []
	TIMETABLE string `db:"TIME_TABLE" protobuf:"string,5,opt,name=time_table"`
	//[ 6] CSR_UUID                                       varbinary            null: false  primary: false  auto: false  col: varbinary       len: -1      default: []
	CSRUUID []byte `db:"CSR_UUID" protobuf:"bytes,6,opt,name=csr_uuid"`
	//[ 7] WHEN_CREATED                                   datetime             null: false  primary: false  auto: false  col: datetime        len: -1      default: []
	WHENCREATED time.Time `db:"WHEN_CREATED" protobuf:"uint64,7,opt,name=when_created"`
	//[ 8] TLM                                            timestamp            null: false  primary: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
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
