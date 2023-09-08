// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NibConnector] class.
var NibConnectorClass = _NibConnectorClass{objc.GetClass("NSNibConnector")}

type _NibConnectorClass struct {
	objc.Class
}

// An interface definition for the [NibConnector] class.
type INibConnector interface {
	objc.IObject
	ReplaceObjectWithObject(oldObject objc.IObject, newObject objc.IObject)
	EstablishConnection()
	Label() string
	SetLabel(value string)
	Destination() objc.Object
	SetDestination(value objc.IObject)
	Source() objc.Object
	SetSource(value objc.IObject)
}

// A connection between two nibs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsnibconnector?language=objc
type NibConnector struct {
	objc.Object
}

func NibConnectorFrom(ptr unsafe.Pointer) NibConnector {
	return NibConnector{
		Object: objc.ObjectFrom(ptr),
	}
}

func (nc _NibConnectorClass) Alloc() NibConnector {
	rv := objc.Call[NibConnector](nc, objc.Sel("alloc"))
	return rv
}

func (nc _NibConnectorClass) New() NibConnector {
	rv := objc.Call[NibConnector](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNibConnector() NibConnector {
	return NibConnectorClass.New()
}

func (n_ NibConnector) Init() NibConnector {
	rv := objc.Call[NibConnector](n_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsnibconnector/1387489-replaceobject?language=objc
func (n_ NibConnector) ReplaceObjectWithObject(oldObject objc.IObject, newObject objc.IObject) {
	objc.Call[objc.Void](n_, objc.Sel("replaceObject:withObject:"), oldObject, newObject)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsnibconnector/1387491-establishconnection?language=objc
func (n_ NibConnector) EstablishConnection() {
	objc.Call[objc.Void](n_, objc.Sel("establishConnection"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsnibconnector/1387493-label?language=objc
func (n_ NibConnector) Label() string {
	rv := objc.Call[string](n_, objc.Sel("label"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsnibconnector/1387493-label?language=objc
func (n_ NibConnector) SetLabel(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setLabel:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsnibconnector/1387482-destination?language=objc
func (n_ NibConnector) Destination() objc.Object {
	rv := objc.Call[objc.Object](n_, objc.Sel("destination"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsnibconnector/1387482-destination?language=objc
func (n_ NibConnector) SetDestination(value objc.IObject) {
	objc.Call[objc.Void](n_, objc.Sel("setDestination:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsnibconnector/1387484-source?language=objc
func (n_ NibConnector) Source() objc.Object {
	rv := objc.Call[objc.Object](n_, objc.Sel("source"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsnibconnector/1387484-source?language=objc
func (n_ NibConnector) SetSource(value objc.IObject) {
	objc.Call[objc.Void](n_, objc.Sel("setSource:"), value)
}
