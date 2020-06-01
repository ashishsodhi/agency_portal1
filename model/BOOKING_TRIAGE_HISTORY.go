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


CREATE TABLE `BOOKING_TRIAGE_HISTORY` (
  `UUID` varchar(16) NOT NULL,
  `PROVIDER_UUID` varchar(16) NOT NULL,
  `BOOKING_UUID` varchar(16) NOT NULL,
  `ACTION` enum('PENDING_ACCEPTANCE','ACCEPTED','DECLINED_TRIAGE','TRIAGE_TIMED_OUT','ASSIGNMENT_TIMED_OUT','DECLINED_PERMANENTLY','FULFILLED') NOT NULL,
  `USER_UUID` varbinary(16) NOT NULL,
  `CSR_UUID` varbinary(16) NOT NULL,
  `WHEN_CREATED` datetime NOT NULL,
  `TLM` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`UUID`),
  KEY `IX_TLM` (`TLM`),
  KEY `BOOKING_TRIAGE_HISTORY_PROVIDER_UUID_fk` (`PROVIDER_UUID`),
  KEY `FK_BOOKING_TRIAGE_HISTORY_BOOKING_TRIAGE_TRACKER` (`BOOKING_UUID`),
  CONSTRAINT `BOOKING_TRIAGE_HISTORY_PROVIDER_UUID_fk` FOREIGN KEY (`PROVIDER_UUID`) REFERENCES `PROVIDER` (`UUID`),
  CONSTRAINT `FK_BOOKING_TRIAGE_HISTORY_BOOKING_TRIAGE_TRACKER` FOREIGN KEY (`BOOKING_UUID`) REFERENCES `BOOKING` (`UUID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "tlm": "2184-12-17T20:36:57.523861643-05:00",    "uuid": "rMZgRKlOegCayctpneEoXMNyi",    "provider_uuid": "PWybZuKzNjjAKaWjBtqLTOINl",    "booking_uuid": "dATxuZqVDJAjXwHzQzDgXDbuW",    "action": "vWrwoKsnPudRAfzLsDuOvIivb",    "user_uuid": "DjonS1YKNhQxPGEGClcSKFNVJw8FUwZYUxYTJSBLLw==",    "csr_uuid": "ElUEAx0TFiBHCwYfYU0KIyxZ",    "when_created": "2312-06-30T10:16:17.690474855-05:00"}



*/

// BOOKINGTRIAGEHISTORY struct is a row record of the BOOKING_TRIAGE_HISTORY table in the agency_portal database
type BOOKINGTRIAGEHISTORY struct {
	//[ 0] UUID                                           varchar(16)          null: false  primary: true   auto: false  col: varchar         len: 16      default: []
	UUID string `db:"UUID" protobuf:"string,0,opt,name=uuid"`
	//[ 1] PROVIDER_UUID                                  varchar(16)          null: false  primary: false  auto: false  col: varchar         len: 16      default: []
	PROVIDERUUID string `db:"PROVIDER_UUID" protobuf:"string,1,opt,name=provider_uuid"`
	//[ 2] BOOKING_UUID                                   varchar(16)          null: false  primary: false  auto: false  col: varchar         len: 16      default: []
	BOOKINGUUID string `db:"BOOKING_UUID" protobuf:"string,2,opt,name=booking_uuid"`
	//[ 3] ACTION                                         char(20)             null: false  primary: false  auto: false  col: char            len: 20      default: []
	ACTION string `db:"ACTION" protobuf:"string,3,opt,name=action"`
	//[ 4] USER_UUID                                      varbinary            null: false  primary: false  auto: false  col: varbinary       len: -1      default: []
	USERUUID []byte `db:"USER_UUID" protobuf:"bytes,4,opt,name=user_uuid"`
	//[ 5] CSR_UUID                                       varbinary            null: false  primary: false  auto: false  col: varbinary       len: -1      default: []
	CSRUUID []byte `db:"CSR_UUID" protobuf:"bytes,5,opt,name=csr_uuid"`
	//[ 6] WHEN_CREATED                                   datetime             null: false  primary: false  auto: false  col: datetime        len: -1      default: []
	WHENCREATED time.Time `db:"WHEN_CREATED" protobuf:"uint64,6,opt,name=when_created"`
	//[ 7] TLM                                            timestamp            null: false  primary: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	TLM time.Time `db:"TLM" protobuf:"uint64,7,opt,name=tlm"`
}

// TableName sets the insert table name for this struct type
func (b *BOOKINGTRIAGEHISTORY) TableName() string {
	return "BOOKING_TRIAGE_HISTORY"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *BOOKINGTRIAGEHISTORY) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *BOOKINGTRIAGEHISTORY) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *BOOKINGTRIAGEHISTORY) Validate(action Action) error {
	return nil
}
