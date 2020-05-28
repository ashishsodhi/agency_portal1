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
  `UID` varchar(45) NOT NULL,
  `PROVIDER_URI` varchar(45) DEFAULT NULL,
  `BOOKING_URI` varchar(45) DEFAULT NULL,
  `ACTION` enum('PENDING_ACCEPTANCE','ACCEPTED','DECLINE_TRIAGE','TRIAGE_TIMED_OUT','DECLINED_PERMANENTLY','FULFILLED') DEFAULT NULL,
  `AGENCY_USER_ID` int(11) DEFAULT NULL,
  `CSR_ID` int(11) DEFAULT NULL,
  `WHEN_CREATED` datetime DEFAULT NULL,
  `TLM` datetime DEFAULT NULL,
  PRIMARY KEY (`UID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "agency_user_id": 48,    "csr_id": 8,    "when_created": "2254-05-17T11:57:56.814906396-05:00",    "tlm": "2193-12-30T01:13:10.362551521-05:00",    "uid": "ASyicMdtEhkBbSOiwpvCIDqpn",    "provider_uri": "UMIqYcLrrWUtmOiLRetWUqxZv",    "booking_uri": "PhuQpWjavzclyzscDCfsqMeSA",    "action": "minJIVmJAchWbEiOqNcqPEPwH"}



*/

// BOOKINGTRIAGEHISTORY struct is a row record of the BOOKING_TRIAGE_HISTORY table in the agency_portal database
type BOOKINGTRIAGEHISTORY struct {
	//[ 0] UID                                            varchar(45)          null: false  primary: true   auto: false  col: varchar         len: 45      default: []
	UID string `db:"UID" protobuf:"string,0,opt,name=uid"`
	//[ 1] PROVIDER_URI                                   varchar(45)          null: true   primary: false  auto: false  col: varchar         len: 45      default: []
	PROVIDERURI sql.NullString `db:"PROVIDER_URI" protobuf:"string,1,opt,name=provider_uri"`
	//[ 2] BOOKING_URI                                    varchar(45)          null: true   primary: false  auto: false  col: varchar         len: 45      default: []
	BOOKINGURI sql.NullString `db:"BOOKING_URI" protobuf:"string,2,opt,name=booking_uri"`
	//[ 3] ACTION                                         char(20)             null: true   primary: false  auto: false  col: char            len: 20      default: []
	ACTION sql.NullString `db:"ACTION" protobuf:"string,3,opt,name=action"`
	//[ 4] AGENCY_USER_ID                                 int                  null: true   primary: false  auto: false  col: int             len: -1      default: []
	AGENCYUSERID sql.NullInt64 `db:"AGENCY_USER_ID" protobuf:"int32,4,opt,name=agency_user_id"`
	//[ 5] CSR_ID                                         int                  null: true   primary: false  auto: false  col: int             len: -1      default: []
	CSRID sql.NullInt64 `db:"CSR_ID" protobuf:"int32,5,opt,name=csr_id"`
	//[ 6] WHEN_CREATED                                   datetime             null: true   primary: false  auto: false  col: datetime        len: -1      default: []
	WHENCREATED time.Time `db:"WHEN_CREATED" protobuf:"uint64,6,opt,name=when_created"`
	//[ 7] TLM                                            datetime             null: true   primary: false  auto: false  col: datetime        len: -1      default: []
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
