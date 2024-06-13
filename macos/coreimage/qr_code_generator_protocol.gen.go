// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a QR code generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciqrcodegenerator?language=objc
type PQRCodeGenerator interface {
	// optional
	SetCorrectionLevel(value string)
	HasSetCorrectionLevel() bool

	// optional
	CorrectionLevel() string
	HasCorrectionLevel() bool

	// optional
	SetMessage(value []byte)
	HasSetMessage() bool

	// optional
	Message() []byte
	HasMessage() bool
}

// ensure impl type implements protocol interface
var _ PQRCodeGenerator = (*QRCodeGeneratorObject)(nil)

// A concrete type for the [PQRCodeGenerator] protocol.
type QRCodeGeneratorObject struct {
	objc.Object
}

func (q_ QRCodeGeneratorObject) HasSetCorrectionLevel() bool {
	return q_.RespondsToSelector(objc.Sel("setCorrectionLevel:"))
}

// The QR code correction level: L, M, Q, or H. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciqrcodegenerator/3228682-correctionlevel?language=objc
func (q_ QRCodeGeneratorObject) SetCorrectionLevel(value string) {
	objc.Call[objc.Void](q_, objc.Sel("setCorrectionLevel:"), value)
}

func (q_ QRCodeGeneratorObject) HasCorrectionLevel() bool {
	return q_.RespondsToSelector(objc.Sel("correctionLevel"))
}

// The QR code correction level: L, M, Q, or H. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciqrcodegenerator/3228682-correctionlevel?language=objc
func (q_ QRCodeGeneratorObject) CorrectionLevel() string {
	rv := objc.Call[string](q_, objc.Sel("correctionLevel"))
	return rv
}

func (q_ QRCodeGeneratorObject) HasSetMessage() bool {
	return q_.RespondsToSelector(objc.Sel("setMessage:"))
}

// The message to encode in the QR code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciqrcodegenerator/3228683-message?language=objc
func (q_ QRCodeGeneratorObject) SetMessage(value []byte) {
	objc.Call[objc.Void](q_, objc.Sel("setMessage:"), value)
}

func (q_ QRCodeGeneratorObject) HasMessage() bool {
	return q_.RespondsToSelector(objc.Sel("message"))
}

// The message to encode in the QR code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciqrcodegenerator/3228683-message?language=objc
func (q_ QRCodeGeneratorObject) Message() []byte {
	rv := objc.Call[[]byte](q_, objc.Sel("message"))
	return rv
}