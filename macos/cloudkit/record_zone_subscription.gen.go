// AUTO-GENERATED CODE, DO NOT MODIFY

package cloudkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RecordZoneSubscription] class.
var RecordZoneSubscriptionClass = _RecordZoneSubscriptionClass{objc.GetClass("CKRecordZoneSubscription")}

type _RecordZoneSubscriptionClass struct {
	objc.Class
}

// An interface definition for the [RecordZoneSubscription] class.
type IRecordZoneSubscription interface {
	ISubscription
	ZoneID() RecordZoneID
	RecordType() RecordType
	SetRecordType(value RecordType)
}

// A subscription that generates push notifications when CloudKit modifies records in a specific record zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordzonesubscription?language=objc
type RecordZoneSubscription struct {
	Subscription
}

func RecordZoneSubscriptionFrom(ptr unsafe.Pointer) RecordZoneSubscription {
	return RecordZoneSubscription{
		Subscription: SubscriptionFrom(ptr),
	}
}

func (rc _RecordZoneSubscriptionClass) Alloc() RecordZoneSubscription {
	rv := objc.Call[RecordZoneSubscription](rc, objc.Sel("alloc"))
	return rv
}

func RecordZoneSubscription_Alloc() RecordZoneSubscription {
	return RecordZoneSubscriptionClass.Alloc()
}

func (rc _RecordZoneSubscriptionClass) New() RecordZoneSubscription {
	rv := objc.Call[RecordZoneSubscription](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRecordZoneSubscription() RecordZoneSubscription {
	return RecordZoneSubscriptionClass.New()
}

func (r_ RecordZoneSubscription) Init() RecordZoneSubscription {
	rv := objc.Call[RecordZoneSubscription](r_, objc.Sel("init"))
	return rv
}

// The ID of the record zone that the subscription queries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordzonesubscription/1640367-zoneid?language=objc
func (r_ RecordZoneSubscription) ZoneID() RecordZoneID {
	rv := objc.Call[RecordZoneID](r_, objc.Sel("zoneID"))
	return rv
}

// The type of record that the subscription queries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordzonesubscription/1640479-recordtype?language=objc
func (r_ RecordZoneSubscription) RecordType() RecordType {
	rv := objc.Call[RecordType](r_, objc.Sel("recordType"))
	return rv
}

// The type of record that the subscription queries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckrecordzonesubscription/1640479-recordtype?language=objc
func (r_ RecordZoneSubscription) SetRecordType(value RecordType) {
	objc.Call[objc.Void](r_, objc.Sel("setRecordType:"), value)
}