// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a fourfold reflected tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldreflectedtile?language=objc
type PFourfoldReflectedTile interface {
	// optional
	SetWidth(value float64)
	HasSetWidth() bool

	// optional
	Width() float64
	HasWidth() bool

	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetAngle(value float64)
	HasSetAngle() bool

	// optional
	Angle() float64
	HasAngle() bool

	// optional
	SetAcuteAngle(value float64)
	HasSetAcuteAngle() bool

	// optional
	AcuteAngle() float64
	HasAcuteAngle() bool

	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool
}

// A concrete type wrapper for the [PFourfoldReflectedTile] protocol.
type FourfoldReflectedTileWrapper struct {
	objc.Object
}

func (f_ FourfoldReflectedTileWrapper) HasSetWidth() bool {
	return f_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The width of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldreflectedtile/3228448-width?language=objc
func (f_ FourfoldReflectedTileWrapper) SetWidth(value float64) {
	objc.Call[objc.Void](f_, objc.Sel("setWidth:"), value)
}

func (f_ FourfoldReflectedTileWrapper) HasWidth() bool {
	return f_.RespondsToSelector(objc.Sel("width"))
}

// The width of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldreflectedtile/3228448-width?language=objc
func (f_ FourfoldReflectedTileWrapper) Width() float64 {
	rv := objc.Call[float64](f_, objc.Sel("width"))
	return rv
}

func (f_ FourfoldReflectedTileWrapper) HasSetInputImage() bool {
	return f_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldreflectedtile/3228447-inputimage?language=objc
func (f_ FourfoldReflectedTileWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](f_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (f_ FourfoldReflectedTileWrapper) HasInputImage() bool {
	return f_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldreflectedtile/3228447-inputimage?language=objc
func (f_ FourfoldReflectedTileWrapper) InputImage() Image {
	rv := objc.Call[Image](f_, objc.Sel("inputImage"))
	return rv
}

func (f_ FourfoldReflectedTileWrapper) HasSetAngle() bool {
	return f_.RespondsToSelector(objc.Sel("setAngle:"))
}

// The angle, in radians, of the tiled pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldreflectedtile/3228445-angle?language=objc
func (f_ FourfoldReflectedTileWrapper) SetAngle(value float64) {
	objc.Call[objc.Void](f_, objc.Sel("setAngle:"), value)
}

func (f_ FourfoldReflectedTileWrapper) HasAngle() bool {
	return f_.RespondsToSelector(objc.Sel("angle"))
}

// The angle, in radians, of the tiled pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldreflectedtile/3228445-angle?language=objc
func (f_ FourfoldReflectedTileWrapper) Angle() float64 {
	rv := objc.Call[float64](f_, objc.Sel("angle"))
	return rv
}

func (f_ FourfoldReflectedTileWrapper) HasSetAcuteAngle() bool {
	return f_.RespondsToSelector(objc.Sel("setAcuteAngle:"))
}

// The primary angle for the repeating reflected tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldreflectedtile/3228444-acuteangle?language=objc
func (f_ FourfoldReflectedTileWrapper) SetAcuteAngle(value float64) {
	objc.Call[objc.Void](f_, objc.Sel("setAcuteAngle:"), value)
}

func (f_ FourfoldReflectedTileWrapper) HasAcuteAngle() bool {
	return f_.RespondsToSelector(objc.Sel("acuteAngle"))
}

// The primary angle for the repeating reflected tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldreflectedtile/3228444-acuteangle?language=objc
func (f_ FourfoldReflectedTileWrapper) AcuteAngle() float64 {
	rv := objc.Call[float64](f_, objc.Sel("acuteAngle"))
	return rv
}

func (f_ FourfoldReflectedTileWrapper) HasSetCenter() bool {
	return f_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldreflectedtile/3228446-center?language=objc
func (f_ FourfoldReflectedTileWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](f_, objc.Sel("setCenter:"), value)
}

func (f_ FourfoldReflectedTileWrapper) HasCenter() bool {
	return f_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourfoldreflectedtile/3228446-center?language=objc
func (f_ FourfoldReflectedTileWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](f_, objc.Sel("center"))
	return rv
}