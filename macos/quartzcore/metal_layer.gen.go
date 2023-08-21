// AUTO-GENERATED CODE, DO NOT MODIFY

package quartzcore

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MetalLayer] class.
var MetalLayerClass = _MetalLayerClass{objc.GetClass("CAMetalLayer")}

type _MetalLayerClass struct {
	objc.Class
}

// An interface definition for the [MetalLayer] class.
type IMetalLayer interface {
	ILayer
	NextDrawable() MetalDrawableWrapper
	WantsExtendedDynamicRangeContent() bool
	SetWantsExtendedDynamicRangeContent(value bool)
	FramebufferOnly() bool
	SetFramebufferOnly(value bool)
	PresentsWithTransaction() bool
	SetPresentsWithTransaction(value bool)
	Device() metal.DeviceWrapper
	SetDevice(value metal.PDevice)
	SetDeviceObject(valueObject objc.IObject)
	EDRMetadata() EDRMetadata
	SetEDRMetadata(value IEDRMetadata)
	Colorspace() coregraphics.ColorSpaceRef
	SetColorspace(value coregraphics.ColorSpaceRef)
	PreferredDevice() metal.DeviceWrapper
	AllowsNextDrawableTimeout() bool
	SetAllowsNextDrawableTimeout(value bool)
	DisplaySyncEnabled() bool
	SetDisplaySyncEnabled(value bool)
	PixelFormat() metal.PixelFormat
	SetPixelFormat(value metal.PixelFormat)
	MaximumDrawableCount() uint
	SetMaximumDrawableCount(value uint)
	DrawableSize() coregraphics.Size
	SetDrawableSize(value coregraphics.Size)
}

// A Core Animation layer that Metal can render into, typically displayed onscreen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer?language=objc
type MetalLayer struct {
	Layer
}

func MetalLayerFrom(ptr unsafe.Pointer) MetalLayer {
	return MetalLayer{
		Layer: LayerFrom(ptr),
	}
}

func (mc _MetalLayerClass) Alloc() MetalLayer {
	rv := objc.Call[MetalLayer](mc, objc.Sel("alloc"))
	return rv
}

func MetalLayer_Alloc() MetalLayer {
	return MetalLayerClass.Alloc()
}

func (mc _MetalLayerClass) New() MetalLayer {
	rv := objc.Call[MetalLayer](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMetalLayer() MetalLayer {
	return MetalLayerClass.New()
}

func (m_ MetalLayer) Init() MetalLayer {
	rv := objc.Call[MetalLayer](m_, objc.Sel("init"))
	return rv
}

func (mc _MetalLayerClass) Layer() MetalLayer {
	rv := objc.Call[MetalLayer](mc, objc.Sel("layer"))
	return rv
}

// Creates and returns an instance of the layer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410793-layer?language=objc
func MetalLayer_Layer() MetalLayer {
	return MetalLayerClass.Layer()
}

func (m_ MetalLayer) InitWithLayer(layer objc.IObject) MetalLayer {
	rv := objc.Call[MetalLayer](m_, objc.Sel("initWithLayer:"), layer)
	return rv
}

// Override to copy or initialize custom fields of the specified layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410842-initwithlayer?language=objc
func NewMetalLayerWithLayer(layer objc.IObject) MetalLayer {
	instance := MetalLayerClass.Alloc().InitWithLayer(layer)
	instance.Autorelease()
	return instance
}

func (m_ MetalLayer) ModelLayer() MetalLayer {
	rv := objc.Call[MetalLayer](m_, objc.Sel("modelLayer"))
	return rv
}

// Returns the model layer object associated with the receiver, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410853-modellayer?language=objc
func MetalLayer_ModelLayer() MetalLayer {
	instance := MetalLayerClass.Alloc().ModelLayer()
	instance.Autorelease()
	return instance
}

func (m_ MetalLayer) PresentationLayer() MetalLayer {
	rv := objc.Call[MetalLayer](m_, objc.Sel("presentationLayer"))
	return rv
}

// Returns a copy of the presentation layer object that represents the state of the layer as it currently appears onscreen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410744-presentationlayer?language=objc
func MetalLayer_PresentationLayer() MetalLayer {
	instance := MetalLayerClass.Alloc().PresentationLayer()
	instance.Autorelease()
	return instance
}

// Waits until a Metal drawable is available, and then returns it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/1478172-nextdrawable?language=objc
func (m_ MetalLayer) NextDrawable() MetalDrawableWrapper {
	rv := objc.Call[MetalDrawableWrapper](m_, objc.Sel("nextDrawable"))
	return rv
}

// Enables extended dynamic range values onscreen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/1478161-wantsextendeddynamicrangecontent?language=objc
func (m_ MetalLayer) WantsExtendedDynamicRangeContent() bool {
	rv := objc.Call[bool](m_, objc.Sel("wantsExtendedDynamicRangeContent"))
	return rv
}

// Enables extended dynamic range values onscreen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/1478161-wantsextendeddynamicrangecontent?language=objc
func (m_ MetalLayer) SetWantsExtendedDynamicRangeContent(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setWantsExtendedDynamicRangeContent:"), value)
}

// A Boolean value that determines whether the layer’s textures are used only for rendering. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/1478168-framebufferonly?language=objc
func (m_ MetalLayer) FramebufferOnly() bool {
	rv := objc.Call[bool](m_, objc.Sel("framebufferOnly"))
	return rv
}

// A Boolean value that determines whether the layer’s textures are used only for rendering. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/1478168-framebufferonly?language=objc
func (m_ MetalLayer) SetFramebufferOnly(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setFramebufferOnly:"), value)
}

// A Boolean value that determines whether the layer presents its content using a Core Animation transaction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/1478157-presentswithtransaction?language=objc
func (m_ MetalLayer) PresentsWithTransaction() bool {
	rv := objc.Call[bool](m_, objc.Sel("presentsWithTransaction"))
	return rv
}

// A Boolean value that determines whether the layer presents its content using a Core Animation transaction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/1478157-presentswithtransaction?language=objc
func (m_ MetalLayer) SetPresentsWithTransaction(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setPresentsWithTransaction:"), value)
}

// The Metal device responsible for the layer’s drawable resources. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/1478163-device?language=objc
func (m_ MetalLayer) Device() metal.DeviceWrapper {
	rv := objc.Call[metal.DeviceWrapper](m_, objc.Sel("device"))
	return rv
}

// The Metal device responsible for the layer’s drawable resources. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/1478163-device?language=objc
func (m_ MetalLayer) SetDevice(value metal.PDevice) {
	po0 := objc.WrapAsProtocol("MTLDevice", value)
	objc.Call[objc.Void](m_, objc.Sel("setDevice:"), po0)
}

// The Metal device responsible for the layer’s drawable resources. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/1478163-device?language=objc
func (m_ MetalLayer) SetDeviceObject(valueObject objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("setDevice:"), objc.Ptr(valueObject))
}

// Metadata describing the tone mapping to apply to the extended dynamic range (EDR) values in the layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/3182052-edrmetadata?language=objc
func (m_ MetalLayer) EDRMetadata() EDRMetadata {
	rv := objc.Call[EDRMetadata](m_, objc.Sel("EDRMetadata"))
	return rv
}

// Metadata describing the tone mapping to apply to the extended dynamic range (EDR) values in the layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/3182052-edrmetadata?language=objc
func (m_ MetalLayer) SetEDRMetadata(value IEDRMetadata) {
	objc.Call[objc.Void](m_, objc.Sel("setEDRMetadata:"), objc.Ptr(value))
}

// The color space of the rendered content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/1478170-colorspace?language=objc
func (m_ MetalLayer) Colorspace() coregraphics.ColorSpaceRef {
	rv := objc.Call[coregraphics.ColorSpaceRef](m_, objc.Sel("colorspace"))
	return rv
}

// The color space of the rendered content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/1478170-colorspace?language=objc
func (m_ MetalLayer) SetColorspace(value coregraphics.ColorSpaceRef) {
	objc.Call[objc.Void](m_, objc.Sel("setColorspace:"), value)
}

// The device object that the system recommends using for this layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/3175021-preferreddevice?language=objc
func (m_ MetalLayer) PreferredDevice() metal.DeviceWrapper {
	rv := objc.Call[metal.DeviceWrapper](m_, objc.Sel("preferredDevice"))
	return rv
}

// A Boolean value that determines whether requests for a new buffer expire if the system can’t satisfy them. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/2887086-allowsnextdrawabletimeout?language=objc
func (m_ MetalLayer) AllowsNextDrawableTimeout() bool {
	rv := objc.Call[bool](m_, objc.Sel("allowsNextDrawableTimeout"))
	return rv
}

// A Boolean value that determines whether requests for a new buffer expire if the system can’t satisfy them. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/2887086-allowsnextdrawabletimeout?language=objc
func (m_ MetalLayer) SetAllowsNextDrawableTimeout(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setAllowsNextDrawableTimeout:"), value)
}

// A Boolean value that determines whether the layer synchronizes its updates to the display’s refresh rate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/2887087-displaysyncenabled?language=objc
func (m_ MetalLayer) DisplaySyncEnabled() bool {
	rv := objc.Call[bool](m_, objc.Sel("displaySyncEnabled"))
	return rv
}

// A Boolean value that determines whether the layer synchronizes its updates to the display’s refresh rate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/2887087-displaysyncenabled?language=objc
func (m_ MetalLayer) SetDisplaySyncEnabled(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setDisplaySyncEnabled:"), value)
}

// The pixel format of the layer’s textures. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/1478155-pixelformat?language=objc
func (m_ MetalLayer) PixelFormat() metal.PixelFormat {
	rv := objc.Call[metal.PixelFormat](m_, objc.Sel("pixelFormat"))
	return rv
}

// The pixel format of the layer’s textures. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/1478155-pixelformat?language=objc
func (m_ MetalLayer) SetPixelFormat(value metal.PixelFormat) {
	objc.Call[objc.Void](m_, objc.Sel("setPixelFormat:"), value)
}

// The number of Metal drawables in the resource pool managed by Core Animation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/2938720-maximumdrawablecount?language=objc
func (m_ MetalLayer) MaximumDrawableCount() uint {
	rv := objc.Call[uint](m_, objc.Sel("maximumDrawableCount"))
	return rv
}

// The number of Metal drawables in the resource pool managed by Core Animation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/2938720-maximumdrawablecount?language=objc
func (m_ MetalLayer) SetMaximumDrawableCount(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setMaximumDrawableCount:"), value)
}

// The size, in pixels, of textures for rendering layer content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/1478174-drawablesize?language=objc
func (m_ MetalLayer) DrawableSize() coregraphics.Size {
	rv := objc.Call[coregraphics.Size](m_, objc.Sel("drawableSize"))
	return rv
}

// The size, in pixels, of textures for rendering layer content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cametallayer/1478174-drawablesize?language=objc
func (m_ MetalLayer) SetDrawableSize(value coregraphics.Size) {
	objc.Call[objc.Void](m_, objc.Sel("setDrawableSize:"), value)
}