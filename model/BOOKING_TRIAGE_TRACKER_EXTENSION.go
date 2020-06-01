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


CREATE TABLE `BOOKING_TRIAGE_TRACKER_EXTENSION` (
  `UUID` varchar(16) NOT NULL,
  `PROVIDER_UUID` varchar(16) NOT NULL,
  `BOOKING_UUID` varchar(16) NOT NULL,
  `PERMANENTLY_DECLINED` bit(1) NOT NULL DEFAULT b'0',
  `NOTES` varchar(10000) DEFAULT NULL,
  `SOURCE` enum('TRIAGE','CSR') NOT NULL DEFAULT 'TRIAGE',
  `CSR_UUID` varbinary(16) NOT NULL,
  `WHEN_CREATED` datetime NOT NULL,
  `TLM` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`UUID`),
  KEY `IX_TLM` (`TLM`),
  KEY `IX_PROVIDER_UUID_PERMANENTLY_DECLINED` (`PROVIDER_UUID`,`PERMANENTLY_DECLINED`),
  KEY `BOOKING_TRIAGE_TRACKER_EXTENSION_BOOKING_UUID_fk` (`BOOKING_UUID`),
  CONSTRAINT `BOOKING_TRIAGE_TRACKER_EXTENSION_BOOKING_UUID_fk` FOREIGN KEY (`BOOKING_UUID`) REFERENCES `BOOKING` (`UUID`),
  CONSTRAINT `BOOKING_TRIAGE_TRACKER_EXTENSION_PROVIDER_UUID` FOREIGN KEY (`PROVIDER_UUID`) REFERENCES `PROVIDER` (`UUID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "csr_uuid": "YBJVXyliBlFjUTJMNQliKCgRDEgLSRkDLkJCTV4BTWI6RlJdJQRjCy4pRw==",    "provider_uuid": "NaIkMQfRSVPXzTkaBFVHPtxXT",    "booking_uuid": "hFBjMauMFooHeUxtCLSfKDLiQ",    "permanently_declined": true,    "source": "HrGYKCOCCWYEVkkTjhDwqspeP",    "uuid": "iWFtiBYJtQspAlpFjqeGGdHHg",    "notes": "tNYChtJTvjWydQVCJuvOsCeZx",    "when_created": "2141-12-27T16:15:46.529511296-05:00",    "tlm": "2089-06-19T03:40:51.497109564-05:00"}



*/

// BOOKINGTRIAGETRACKEREXTENSION struct is a row record of the BOOKING_TRIAGE_TRACKER_EXTENSION table in the agency_portal database
type BOOKINGTRIAGETRACKEREXTENSION struct {
	//[ 0] UUID                                           varchar(16)          null: false  primary: true   auto: false  col: varchar         len: 16      default: []
	UUID string `db:"UUID" protobuf:"string,0,opt,name=uuid"`
	//[ 1] PROVIDER_UUID                                  varchar(16)          null: false  primary: false  auto: false  col: varchar         len: 16      default: []
	PROVIDERUUID string `db:"PROVIDER_UUID" protobuf:"string,1,opt,name=provider_uuid"`
	//[ 2] BOOKING_UUID                                   varchar(16)          null: false  primary: false  auto: false  col: varchar         len: 16      default: []
	BOOKINGUUID string `db:"BOOKING_UUID" protobuf:"string,2,opt,name=booking_uuid"`
	//[ 3] PERMANENTLY_DECLINED                           bit                  null: false  primary: false  auto: false  col: bit             len: -1      default: [b'0']
	PERMANENTLYDECLINED bool `db:"PERMANENTLY_DECLINED" protobuf:"bool,3,opt,name=permanently_declined"`
	//[ 4] NOTES                                          varchar(10000)       null: true   primary: false  auto: false  col: varchar         len: 10000   default: []
	NOTES sql.NullString `db:"NOTES" protobuf:"string,4,opt,name=notes"`
	//[ 5] SOURCE                                         char(6)              null: false  primary: false  auto: false  col: char            len: 6       default: [TRIAGE]
	SOURCE string `db:"SOURCE" protobuf:"string,5,opt,name=source"`
	//[ 6] CSR_UUID                                       varbinary            null: false  primary: false  auto: false  col: varbinary       len: -1      default: []
	CSRUUID []byte `db:"CSR_UUID" protobuf:"bytes,6,opt,name=csr_uuid"`
	//[ 7] WHEN_CREATED                                   datetime             null: false  primary: false  auto: false  col: datetime        len: -1      default: []
	WHENCREATED time.Time `db:"WHEN_CREATED" protobuf:"uint64,7,opt,name=when_created"`
	//[ 8] TLM                                            timestamp            null: false  primary: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	TLM time.Time `db:"TLM" protobuf:"uint64,8,opt,name=tlm"`
}

// TableName sets the insert table name for this struct type
func (b *BOOKINGTRIAGETRACKEREXTENSION) TableName() string {
	return "BOOKING_TRIAGE_TRACKER_EXTENSION"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *BOOKINGTRIAGETRACKEREXTENSION) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *BOOKINGTRIAGETRACKEREXTENSION) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *BOOKINGTRIAGETRACKEREXTENSION) Validate(action Action) error {
	return nil
}
