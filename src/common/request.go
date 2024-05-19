package common

const (
	RequestStatusRequested = "requested"
	RequestStatusCancel    = "cancel"
	RequestStatusInReview  = "in_review"
	RequestStatusApproved  = "approved"
	RequestStatusRejected  = "rejected"

	ApprovalStatusInReview = RequestStatusInReview
	ApprovalStatusApproved = RequestStatusApproved
	ApprovalStatusRejected = RequestStatusRejected

	RequestTypeLeaveSick                   = "sick-leave"
	RequestTypeLeaveOthers                 = "other-leave"
	RequestTypeLeaveMaternity              = "maternity-leave"
	RequestTypeLeavePersonal               = "personal-leave"
	RequestTypeLeaveMarriage               = "marriage-leave"
	RequestTypeLeaveChildMarriage          = "child_marriage-leave"
	RequestTypeLeaveChildCircumcision      = "child_circumcision-leave"
	RequestTypeLeaveChildBaptism           = "child_baptism-leave"
	RequestTypeLeavePaternity              = "paternity-leave"
	RequestTypeLeaveCompassionateFamily    = "compassionate_family-leave"
	RequestTypeLeaveCompassionateHousehold = "compassionate_household-leave"
	RequestTypeLeaveUnpaid                 = "unpaid-leave"
	RequestTypeBKO                         = "bko"
	RequestTypeAbsence                     = "absence"
	RequestTypeOvertime                    = "overtime"
)

var (
	ValidRequestType = map[string]bool{
		RequestTypeLeavePersonal:               true,
		RequestTypeLeaveSick:                   true,
		RequestTypeLeaveOthers:                 true,
		RequestTypeLeaveMaternity:              true,
		RequestTypeLeaveMarriage:               true,
		RequestTypeLeaveChildMarriage:          true,
		RequestTypeLeaveChildCircumcision:      true,
		RequestTypeLeaveChildBaptism:           true,
		RequestTypeLeavePaternity:              true,
		RequestTypeLeaveCompassionateFamily:    true,
		RequestTypeLeaveCompassionateHousehold: true,
		RequestTypeLeaveUnpaid:                 true,
		RequestTypeBKO:                         true,
		RequestTypeAbsence:                     true,
		RequestTypeOvertime:                    true,
	}

	RequestTypeLeaveOnly = map[string]bool{
		RequestTypeLeavePersonal:               true,
		RequestTypeLeaveSick:                   true,
		RequestTypeLeaveMaternity:              true,
		RequestTypeLeaveMarriage:               true,
		RequestTypeLeaveChildMarriage:          true,
		RequestTypeLeaveChildCircumcision:      true,
		RequestTypeLeaveChildBaptism:           true,
		RequestTypeLeavePaternity:              true,
		RequestTypeLeaveCompassionateFamily:    true,
		RequestTypeLeaveCompassionateHousehold: true,
		RequestTypeLeaveUnpaid:                 true,
	}

	ValidRequestStatus = map[string]bool{
		RequestStatusRequested: true,
		RequestStatusCancel:    true,
		RequestStatusInReview:  true,
		RequestStatusApproved:  true,
		RequestStatusRejected:  true,
	}
)
