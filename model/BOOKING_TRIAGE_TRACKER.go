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


CREATE TABLE `BOOKING_TRIAGE_TRACKER` (
  `UUID` varchar(16) NOT NULL,
  `PROVIDER_UUID` varchar(16) DEFAULT NULL,
  `TRIAGE_ORDER` smallint(5) unsigned DEFAULT NULL,
  `BOOKING_UUID` varchar(16) NOT NULL,
  `TIME_OFFERED_IN_MINS` smallint(5) unsigned DEFAULT NULL,
  `TIMER_EXECUTION_ID` varchar(255) DEFAULT NULL,
  `TIMER_EXPIRES_AT` datetime DEFAULT NULL,
  `TRACKING_TYPE` enum('ACCEPTANCE_TIMER','ASSIGNMENT_TIMER','NO_TIMER') DEFAULT NULL,
  `STATUS` enum('PENDING','FULFILLED','CANCELLED') NOT NULL,
  `STATUS_TLM` datetime NOT NULL,
  `TRIAGE_COMPLETED` bit(1) NOT NULL DEFAULT b'0',
  `WHEN_CREATED` datetime DEFAULT NULL,
  `TLM` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`UUID`),
  UNIQUE KEY `UNQ_BOOKING_UUID` (`BOOKING_UUID`),
  KEY `IX_TLM` (`TLM`),
  KEY `IX_TIMER_EXECUTION_ID` (`TIMER_EXECUTION_ID`),
  KEY `IX_PROVIDER_UUID_STATUS` (`PROVIDER_UUID`,`STATUS`),
  CONSTRAINT `FK_BOOKING_TRIAGE_TRACKER_BOOKING_BELONGS_TO` FOREIGN KEY (`BOOKING_UUID`) REFERENCES `BOOKING` (`UUID`),
  CONSTRAINT `FK_BOOKING_TRIAGE_TRACKER_PROVIDER_BELONGS_TO` FOREIGN KEY (`PROVIDER_UUID`) REFERENCES `PROVIDER` (`UUID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "triage_completed": true,    "tlm": "2104-05-08T23:45:40.717608522-05:00",    "uuid": "IAjbbTROumqPZCEsHeKNGAkjo",    "provider_uuid": "yIosxdiJNlBrWdVicumwnxwgN",    "booking_uuid": "QqVUWgCqCZLkcPNSCgCTLAVcq",    "time_offered_in_mins": 46,    "status": "TpXdUKeHgfatFbRrnPOSCRKFI",    "when_created": "2257-08-09T18:29:52.347571386-05:00",    "triage_order": 57,    "timer_execution_id": "ZbIFYwRKYzzIUhifFxKQqyZYB",    "timer_expires_at": "2082-12-16T03:30:36.87672429-05:00",    "tracking_type": "FoFKTSigQDxGMZVbBeKGSmKgb",    "status_tlm": "2114-08-25T06:56:58.229255186-05:00"}



*/

// BOOKINGTRIAGETRACKER struct is a row record of the BOOKING_TRIAGE_TRACKER table in the agency_portal database
type BOOKINGTRIAGETRACKER struct {
	//[ 0] UUID                                           varchar(16)          null: false  primary: true   auto: false  col: varchar         len: 16      default: []
	UUID string `db:"UUID" protobuf:"string,0,opt,name=uuid"`
	//[ 1] PROVIDER_UUID                                  varchar(16)          null: true   primary: false  auto: false  col: varchar         len: 16      default: []
	PROVIDERUUID sql.NullString `db:"PROVIDER_UUID" protobuf:"string,1,opt,name=provider_uuid"`
	//[ 2] TRIAGE_ORDER                                   smallint             null: true   primary: false  auto: false  col: smallint        len: -1      default: []
	TRIAGEORDER sql.NullInt64 `db:"TRIAGE_ORDER" protobuf:"int16,2,opt,name=triage_order"`
	//[ 3] BOOKING_UUID                                   varchar(16)          null: false  primary: false  auto: false  col: varchar         len: 16      default: []
	BOOKINGUUID string `db:"BOOKING_UUID" protobuf:"string,3,opt,name=booking_uuid"`
	//[ 4] TIME_OFFERED_IN_MINS                           smallint             null: true   primary: false  auto: false  col: smallint        len: -1      default: []
	TIMEOFFEREDINMINS sql.NullInt64 `db:"TIME_OFFERED_IN_MINS" protobuf:"int16,4,opt,name=time_offered_in_mins"`
	//[ 5] TIMER_EXECUTION_ID                             varchar(255)         null: true   primary: false  auto: false  col: varchar         len: 255     default: []
	TIMEREXECUTIONID sql.NullString `db:"TIMER_EXECUTION_ID" protobuf:"string,5,opt,name=timer_execution_id"`
	//[ 6] TIMER_EXPIRES_AT                               datetime             null: true   primary: false  auto: false  col: datetime        len: -1      default: []
	TIMEREXPIRESAT time.Time `db:"TIMER_EXPIRES_AT" protobuf:"uint64,6,opt,name=timer_expires_at"`
	//[ 7] TRACKING_TYPE                                  char(16)             null: true   primary: false  auto: false  col: char            len: 16      default: []
	TRACKINGTYPE sql.NullString `db:"TRACKING_TYPE" protobuf:"string,7,opt,name=tracking_type"`
	//[ 8] STATUS                                         char(9)              null: false  primary: false  auto: false  col: char            len: 9       default: []
	STATUS string `db:"STATUS" protobuf:"string,8,opt,name=status"`
	//[ 9] STATUS_TLM                                     datetime             null: false  primary: false  auto: false  col: datetime        len: -1      default: []
	STATUSTLM time.Time `db:"STATUS_TLM" protobuf:"uint64,9,opt,name=status_tlm"`
	//[10] TRIAGE_COMPLETED                               bit                  null: false  primary: false  auto: false  col: bit             len: -1      default: [b'0']
	TRIAGECOMPLETED bool `db:"TRIAGE_COMPLETED" protobuf:"bool,10,opt,name=triage_completed"`
	//[11] WHEN_CREATED                                   datetime             null: true   primary: false  auto: false  col: datetime        len: -1      default: []
	WHENCREATED time.Time `db:"WHEN_CREATED" protobuf:"uint64,11,opt,name=when_created"`
	//[12] TLM                                            timestamp            null: false  primary: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	TLM time.Time `db:"TLM" protobuf:"uint64,12,opt,name=tlm"`
}

// TableName sets the insert table name for this struct type
func (b *BOOKINGTRIAGETRACKER) TableName() string {
	return "BOOKING_TRIAGE_TRACKER"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *BOOKINGTRIAGETRACKER) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *BOOKINGTRIAGETRACKER) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *BOOKINGTRIAGETRACKER) Validate(action Action) error {
	return nil
}
