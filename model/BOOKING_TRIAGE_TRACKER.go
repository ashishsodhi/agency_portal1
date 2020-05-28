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
  `PROVIDER_URI` varchar(45) NOT NULL,
  `TRIAGE_ORDER` int(11) DEFAULT NULL,
  `BOOKING_URI` varchar(45) DEFAULT NULL,
  `TIME_OFFERED_IN_MINS` int(11) DEFAULT NULL,
  `TRACKER_TYPE` enum('ACCEPTANCE_TIMER','ASSIGNMENT_TIMER','NO_TIMER') NOT NULL,
  `ACTION_TAKEN` enum('PENDING','FULFILLED','CANCELLED') DEFAULT NULL,
  `ACTION_TAKEN_TLM` datetime DEFAULT NULL,
  `STATUS` enum('ACTIVE','INACTIVE') DEFAULT NULL,
  `STATUS_TLM` datetime DEFAULT NULL,
  `TRIAGE_COMPLETED` bit(1) DEFAULT NULL,
  `WHEN_CREATED` datetime DEFAULT NULL,
  `TLM` datetime DEFAULT NULL,
  PRIMARY KEY (`PROVIDER_URI`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "action_taken_tlm": "2132-09-20T14:32:04.88902184-05:00",    "status": "FjXrRHVSYvvokzTOoAGqKACsl",    "triage_completed": true,    "when_created": "2148-10-22T22:47:55.280908071-05:00",    "triage_order": 34,    "booking_uri": "arZcIISQFXUieuygdRKIcuSXQ",    "time_offered_in_mins": 16,    "action_taken": "fRRdyUlzznVqkznnOcGPBBEDt",    "tlm": "2243-01-26T15:34:08.721960576-05:00",    "provider_uri": "CgFakbegxEgHWXBvyWjHFtypB",    "tracker_type": "dtfbJIPkZIBOzFznEuPVMwUPU",    "status_tlm": "2059-09-14T21:12:31.392773321-05:00"}



*/

// BOOKINGTRIAGETRACKER struct is a row record of the BOOKING_TRIAGE_TRACKER table in the agency_portal database
type BOOKINGTRIAGETRACKER struct {
	//[ 0] PROVIDER_URI                                   varchar(45)          null: false  primary: true   auto: false  col: varchar         len: 45      default: []
	PROVIDERURI string `db:"PROVIDER_URI" protobuf:"string,0,opt,name=provider_uri"`
	//[ 1] TRIAGE_ORDER                                   int                  null: true   primary: false  auto: false  col: int             len: -1      default: []
	TRIAGEORDER sql.NullInt64 `db:"TRIAGE_ORDER" protobuf:"int32,1,opt,name=triage_order"`
	//[ 2] BOOKING_URI                                    varchar(45)          null: true   primary: false  auto: false  col: varchar         len: 45      default: []
	BOOKINGURI sql.NullString `db:"BOOKING_URI" protobuf:"string,2,opt,name=booking_uri"`
	//[ 3] TIME_OFFERED_IN_MINS                           int                  null: true   primary: false  auto: false  col: int             len: -1      default: []
	TIMEOFFEREDINMINS sql.NullInt64 `db:"TIME_OFFERED_IN_MINS" protobuf:"int32,3,opt,name=time_offered_in_mins"`
	//[ 4] TRACKER_TYPE                                   char(16)             null: false  primary: false  auto: false  col: char            len: 16      default: []
	TRACKERTYPE string `db:"TRACKER_TYPE" protobuf:"string,4,opt,name=tracker_type"`
	//[ 5] ACTION_TAKEN                                   char(9)              null: true   primary: false  auto: false  col: char            len: 9       default: []
	ACTIONTAKEN sql.NullString `db:"ACTION_TAKEN" protobuf:"string,5,opt,name=action_taken"`
	//[ 6] ACTION_TAKEN_TLM                               datetime             null: true   primary: false  auto: false  col: datetime        len: -1      default: []
	ACTIONTAKENTLM time.Time `db:"ACTION_TAKEN_TLM" protobuf:"uint64,6,opt,name=action_taken_tlm"`
	//[ 7] STATUS                                         char(8)              null: true   primary: false  auto: false  col: char            len: 8       default: []
	STATUS sql.NullString `db:"STATUS" protobuf:"string,7,opt,name=status"`
	//[ 8] STATUS_TLM                                     datetime             null: true   primary: false  auto: false  col: datetime        len: -1      default: []
	STATUSTLM time.Time `db:"STATUS_TLM" protobuf:"uint64,8,opt,name=status_tlm"`
	//[ 9] TRIAGE_COMPLETED                               bit                  null: true   primary: false  auto: false  col: bit             len: -1      default: []
	TRIAGECOMPLETED sql.NullBool `db:"TRIAGE_COMPLETED" protobuf:"bool,9,opt,name=triage_completed"`
	//[10] WHEN_CREATED                                   datetime             null: true   primary: false  auto: false  col: datetime        len: -1      default: []
	WHENCREATED time.Time `db:"WHEN_CREATED" protobuf:"uint64,10,opt,name=when_created"`
	//[11] TLM                                            datetime             null: true   primary: false  auto: false  col: datetime        len: -1      default: []
	TLM time.Time `db:"TLM" protobuf:"uint64,11,opt,name=tlm"`
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
