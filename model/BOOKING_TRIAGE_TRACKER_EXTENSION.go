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
  `UID` varchar(45) NOT NULL,
  `PROVIDER_URI` varchar(45) DEFAULT NULL,
  `BOOKING_URI` varchar(45) DEFAULT NULL,
  `PERMANENTLY_DECLINED` bit(1) DEFAULT NULL,
  `NOTES` varchar(10000) DEFAULT NULL,
  `SOURCE` enum('TRIAGE','CSR') DEFAULT NULL,
  `CSR_ID` varchar(45) DEFAULT NULL,
  `WHEN_CREATED` datetime DEFAULT NULL,
  `TLM` datetime DEFAULT NULL,
  PRIMARY KEY (`UID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "uid": "nyCpxZgAzhMBuULVkDhdrlDzF",    "permanently_declined": false,    "csr_id": "eJroetjMYZfSXJsQHHdLNPfJF",    "when_created": "2155-02-05T00:34:35.624961863-05:00",    "tlm": "2242-03-15T07:09:09.778548991-05:00",    "provider_uri": "FIcQcntWYIHbOGCtZPlGsCxWY",    "booking_uri": "iRnmlqzJwhDtcGscUGkpIdyuV",    "notes": "sLYIyaXHpsnXBVlSKfStGCxBV",    "source": "AWSOJPZAkBAmZJUPCvtWQXgAZ"}



*/

// BOOKINGTRIAGETRACKEREXTENSION struct is a row record of the BOOKING_TRIAGE_TRACKER_EXTENSION table in the agency_portal database
type BOOKINGTRIAGETRACKEREXTENSION struct {
	//[ 0] UID                                            varchar(45)          null: false  primary: true   auto: false  col: varchar         len: 45      default: []
	UID string `db:"UID" protobuf:"string,0,opt,name=uid"`
	//[ 1] PROVIDER_URI                                   varchar(45)          null: true   primary: false  auto: false  col: varchar         len: 45      default: []
	PROVIDERURI sql.NullString `db:"PROVIDER_URI" protobuf:"string,1,opt,name=provider_uri"`
	//[ 2] BOOKING_URI                                    varchar(45)          null: true   primary: false  auto: false  col: varchar         len: 45      default: []
	BOOKINGURI sql.NullString `db:"BOOKING_URI" protobuf:"string,2,opt,name=booking_uri"`
	//[ 3] PERMANENTLY_DECLINED                           bit                  null: true   primary: false  auto: false  col: bit             len: -1      default: []
	PERMANENTLYDECLINED sql.NullBool `db:"PERMANENTLY_DECLINED" protobuf:"bool,3,opt,name=permanently_declined"`
	//[ 4] NOTES                                          varchar(10000)       null: true   primary: false  auto: false  col: varchar         len: 10000   default: []
	NOTES sql.NullString `db:"NOTES" protobuf:"string,4,opt,name=notes"`
	//[ 5] SOURCE                                         char(6)              null: true   primary: false  auto: false  col: char            len: 6       default: []
	SOURCE sql.NullString `db:"SOURCE" protobuf:"string,5,opt,name=source"`
	//[ 6] CSR_ID                                         varchar(45)          null: true   primary: false  auto: false  col: varchar         len: 45      default: []
	CSRID sql.NullString `db:"CSR_ID" protobuf:"string,6,opt,name=csr_id"`
	//[ 7] WHEN_CREATED                                   datetime             null: true   primary: false  auto: false  col: datetime        len: -1      default: []
	WHENCREATED time.Time `db:"WHEN_CREATED" protobuf:"uint64,7,opt,name=when_created"`
	//[ 8] TLM                                            datetime             null: true   primary: false  auto: false  col: datetime        len: -1      default: []
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
