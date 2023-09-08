// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageLaplacianPyramid] class.
var ImageLaplacianPyramidClass = _ImageLaplacianPyramidClass{objc.GetClass("MPSImageLaplacianPyramid")}

type _ImageLaplacianPyramidClass struct {
	objc.Class
}

// An interface definition for the [ImageLaplacianPyramid] class.
type IImageLaplacianPyramid interface {
	IImagePyramid
	GetLaplacianBias() float64
	SetLaplacianBias(value float64)
	GetLaplacianScale() float64
	SetLaplacianScale(value float64)
}

// A filter that convolves an image with a Laplacian filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagelaplacianpyramid?language=objc
type ImageLaplacianPyramid struct {
	ImagePyramid
}

func ImageLaplacianPyramidFrom(ptr unsafe.Pointer) ImageLaplacianPyramid {
	return ImageLaplacianPyramid{
		ImagePyramid: ImagePyramidFrom(ptr),
	}
}

func (ic _ImageLaplacianPyramidClass) Alloc() ImageLaplacianPyramid {
	rv := objc.Call[ImageLaplacianPyramid](ic, objc.Sel("alloc"))
	return rv
}

func (ic _ImageLaplacianPyramidClass) New() ImageLaplacianPyramid {
	rv := objc.Call[ImageLaplacianPyramid](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageLaplacianPyramid() ImageLaplacianPyramid {
	return ImageLaplacianPyramidClass.New()
}

func (i_ ImageLaplacianPyramid) Init() ImageLaplacianPyramid {
	rv := objc.Call[ImageLaplacianPyramid](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageLaplacianPyramid) InitWithDevice(device metal.PDevice) ImageLaplacianPyramid {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageLaplacianPyramid](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a downwards 5-tap image pyramid with the default filter kernel and device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagepyramid/1648935-initwithdevice?language=objc
func NewImageLaplacianPyramidWithDevice(device metal.PDevice) ImageLaplacianPyramid {
	instance := ImageLaplacianPyramidClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageLaplacianPyramid) InitWithDeviceCenterWeight(device metal.PDevice, centerWeight float64) ImageLaplacianPyramid {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageLaplacianPyramid](i_, objc.Sel("initWithDevice:centerWeight:"), po0, centerWeight)
	return rv
}

// Initialize a downwards 5-tap image pyramid with a central weight parameter and device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagepyramid/1648889-initwithdevice?language=objc
func NewImageLaplacianPyramidWithDeviceCenterWeight(device metal.PDevice, centerWeight float64) ImageLaplacianPyramid {
	instance := ImageLaplacianPyramidClass.Alloc().InitWithDeviceCenterWeight(device, centerWeight)
	instance.Autorelease()
	return instance
}

func (i_ ImageLaplacianPyramid) InitWithDeviceKernelWidthKernelHeightWeights(device metal.PDevice, kernelWidth uint, kernelHeight uint, kernelWeights *float64) ImageLaplacianPyramid {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageLaplacianPyramid](i_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:weights:"), po0, kernelWidth, kernelHeight, kernelWeights)
	return rv
}

// Initialize a downwards n-tap image pyramid with a custom filter kernel and device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagepyramid/1648821-initwithdevice?language=objc
func NewImageLaplacianPyramidWithDeviceKernelWidthKernelHeightWeights(device metal.PDevice, kernelWidth uint, kernelHeight uint, kernelWeights *float64) ImageLaplacianPyramid {
	instance := ImageLaplacianPyramidClass.Alloc().InitWithDeviceKernelWidthKernelHeightWeights(device, kernelWidth, kernelHeight, kernelWeights)
	instance.Autorelease()
	return instance
}

func (i_ ImageLaplacianPyramid) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageLaplacianPyramid {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageLaplacianPyramid](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageLaplacianPyramid_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageLaplacianPyramid {
	instance := ImageLaplacianPyramidClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagelaplacianpyramid/2966645-laplacianbias?language=objc
func (i_ ImageLaplacianPyramid) GetLaplacianBias() float64 {
	rv := objc.Call[float64](i_, objc.Sel("getLaplacianBias"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagelaplacianpyramid/2966645-laplacianbias?language=objc
func (i_ ImageLaplacianPyramid) SetLaplacianBias(value float64) {
	objc.Call[objc.Void](i_, objc.Sel("setLaplacianBias:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagelaplacianpyramid/2966646-laplacianscale?language=objc
func (i_ ImageLaplacianPyramid) GetLaplacianScale() float64 {
	rv := objc.Call[float64](i_, objc.Sel("getLaplacianScale"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagelaplacianpyramid/2966646-laplacianscale?language=objc
func (i_ ImageLaplacianPyramid) SetLaplacianScale(value float64) {
	objc.Call[objc.Void](i_, objc.Sel("setLaplacianScale:"), value)
}
