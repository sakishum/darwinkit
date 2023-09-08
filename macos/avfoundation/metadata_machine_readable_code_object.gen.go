// Code generated by DarwinKit. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coreimage"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MetadataMachineReadableCodeObject] class.
var MetadataMachineReadableCodeObjectClass = _MetadataMachineReadableCodeObjectClass{objc.GetClass("AVMetadataMachineReadableCodeObject")}

type _MetadataMachineReadableCodeObjectClass struct {
	objc.Class
}

// An interface definition for the [MetadataMachineReadableCodeObject] class.
type IMetadataMachineReadableCodeObject interface {
	IMetadataObject
	Descriptor() coreimage.BarcodeDescriptor
	Corners() []foundation.Dictionary
	StringValue() string
}

// Barcode information detected by a metadata capture output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatamachinereadablecodeobject?language=objc
type MetadataMachineReadableCodeObject struct {
	MetadataObject
}

func MetadataMachineReadableCodeObjectFrom(ptr unsafe.Pointer) MetadataMachineReadableCodeObject {
	return MetadataMachineReadableCodeObject{
		MetadataObject: MetadataObjectFrom(ptr),
	}
}

func (mc _MetadataMachineReadableCodeObjectClass) Alloc() MetadataMachineReadableCodeObject {
	rv := objc.Call[MetadataMachineReadableCodeObject](mc, objc.Sel("alloc"))
	return rv
}

func (mc _MetadataMachineReadableCodeObjectClass) New() MetadataMachineReadableCodeObject {
	rv := objc.Call[MetadataMachineReadableCodeObject](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMetadataMachineReadableCodeObject() MetadataMachineReadableCodeObject {
	return MetadataMachineReadableCodeObjectClass.New()
}

func (m_ MetadataMachineReadableCodeObject) Init() MetadataMachineReadableCodeObject {
	rv := objc.Call[MetadataMachineReadableCodeObject](m_, objc.Sel("init"))
	return rv
}

// A barcode description for use in Core Image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatamachinereadablecodeobject/2875944-descriptor?language=objc
func (m_ MetadataMachineReadableCodeObject) Descriptor() coreimage.BarcodeDescriptor {
	rv := objc.Call[coreimage.BarcodeDescriptor](m_, objc.Sel("descriptor"))
	return rv
}

// The points defining the (x, 	y) locations of the corners. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatamachinereadablecodeobject/1618815-corners?language=objc
func (m_ MetadataMachineReadableCodeObject) Corners() []foundation.Dictionary {
	rv := objc.Call[[]foundation.Dictionary](m_, objc.Sel("corners"))
	return rv
}

// Returns the error-corrected data decoded into a human-readable string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatamachinereadablecodeobject/1618800-stringvalue?language=objc
func (m_ MetadataMachineReadableCodeObject) StringValue() string {
	rv := objc.Call[string](m_, objc.Sel("stringValue"))
	return rv
}
