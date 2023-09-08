// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageCopyToMatrix] class.
var ImageCopyToMatrixClass = _ImageCopyToMatrixClass{objc.GetClass("MPSImageCopyToMatrix")}

type _ImageCopyToMatrixClass struct {
	objc.Class
}

// An interface definition for the [ImageCopyToMatrix] class.
type IImageCopyToMatrix interface {
	IKernel
	EncodeToCommandBufferSourceImageDestinationMatrix(commandBuffer metal.PCommandBuffer, sourceImage IImage, destinationMatrix IMatrix)
	EncodeToCommandBufferObjectSourceImageDestinationMatrix(commandBufferObject objc.IObject, sourceImage IImage, destinationMatrix IMatrix)
	EncodeBatchToCommandBufferSourceImagesDestinationMatrix(commandBuffer metal.PCommandBuffer, sourceImages *foundation.Array, destinationMatrix IMatrix)
	EncodeBatchToCommandBufferObjectSourceImagesDestinationMatrix(commandBufferObject objc.IObject, sourceImages *foundation.Array, destinationMatrix IMatrix)
	DataLayout() DataLayout
	DestinationMatrixOrigin() metal.Origin
	SetDestinationMatrixOrigin(value metal.Origin)
	DestinationMatrixBatchIndex() uint
	SetDestinationMatrixBatchIndex(value uint)
}

// A class that copies image data to a matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix?language=objc
type ImageCopyToMatrix struct {
	Kernel
}

func ImageCopyToMatrixFrom(ptr unsafe.Pointer) ImageCopyToMatrix {
	return ImageCopyToMatrix{
		Kernel: KernelFrom(ptr),
	}
}

func (i_ ImageCopyToMatrix) InitWithDeviceDataLayout(device metal.PDevice, dataLayout DataLayout) ImageCopyToMatrix {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageCopyToMatrix](i_, objc.Sel("initWithDevice:dataLayout:"), po0, dataLayout)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/2873210-initwithdevice?language=objc
func NewImageCopyToMatrixWithDeviceDataLayout(device metal.PDevice, dataLayout DataLayout) ImageCopyToMatrix {
	instance := ImageCopyToMatrixClass.Alloc().InitWithDeviceDataLayout(device, dataLayout)
	instance.Autorelease()
	return instance
}

func (ic _ImageCopyToMatrixClass) Alloc() ImageCopyToMatrix {
	rv := objc.Call[ImageCopyToMatrix](ic, objc.Sel("alloc"))
	return rv
}

func (ic _ImageCopyToMatrixClass) New() ImageCopyToMatrix {
	rv := objc.Call[ImageCopyToMatrix](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageCopyToMatrix() ImageCopyToMatrix {
	return ImageCopyToMatrixClass.New()
}

func (i_ ImageCopyToMatrix) Init() ImageCopyToMatrix {
	rv := objc.Call[ImageCopyToMatrix](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageCopyToMatrix) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageCopyToMatrix {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageCopyToMatrix](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageCopyToMatrix_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageCopyToMatrix {
	instance := ImageCopyToMatrixClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

func (i_ ImageCopyToMatrix) InitWithDevice(device metal.PDevice) ImageCopyToMatrix {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageCopyToMatrix](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewImageCopyToMatrixWithDevice(device metal.PDevice) ImageCopyToMatrix {
	instance := ImageCopyToMatrixClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/2873212-encodetocommandbuffer?language=objc
func (i_ ImageCopyToMatrix) EncodeToCommandBufferSourceImageDestinationMatrix(commandBuffer metal.PCommandBuffer, sourceImage IImage, destinationMatrix IMatrix) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](i_, objc.Sel("encodeToCommandBuffer:sourceImage:destinationMatrix:"), po0, objc.Ptr(sourceImage), objc.Ptr(destinationMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/2873212-encodetocommandbuffer?language=objc
func (i_ ImageCopyToMatrix) EncodeToCommandBufferObjectSourceImageDestinationMatrix(commandBufferObject objc.IObject, sourceImage IImage, destinationMatrix IMatrix) {
	objc.Call[objc.Void](i_, objc.Sel("encodeToCommandBuffer:sourceImage:destinationMatrix:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceImage), objc.Ptr(destinationMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/3013769-encodebatchtocommandbuffer?language=objc
func (i_ ImageCopyToMatrix) EncodeBatchToCommandBufferSourceImagesDestinationMatrix(commandBuffer metal.PCommandBuffer, sourceImages *foundation.Array, destinationMatrix IMatrix) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](i_, objc.Sel("encodeBatchToCommandBuffer:sourceImages:destinationMatrix:"), po0, sourceImages, objc.Ptr(destinationMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/3013769-encodebatchtocommandbuffer?language=objc
func (i_ ImageCopyToMatrix) EncodeBatchToCommandBufferObjectSourceImagesDestinationMatrix(commandBufferObject objc.IObject, sourceImages *foundation.Array, destinationMatrix IMatrix) {
	objc.Call[objc.Void](i_, objc.Sel("encodeBatchToCommandBuffer:sourceImages:destinationMatrix:"), objc.Ptr(commandBufferObject), sourceImages, objc.Ptr(destinationMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/2873216-datalayout?language=objc
func (i_ ImageCopyToMatrix) DataLayout() DataLayout {
	rv := objc.Call[DataLayout](i_, objc.Sel("dataLayout"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/2873215-destinationmatrixorigin?language=objc
func (i_ ImageCopyToMatrix) DestinationMatrixOrigin() metal.Origin {
	rv := objc.Call[metal.Origin](i_, objc.Sel("destinationMatrixOrigin"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/2873215-destinationmatrixorigin?language=objc
func (i_ ImageCopyToMatrix) SetDestinationMatrixOrigin(value metal.Origin) {
	objc.Call[objc.Void](i_, objc.Sel("setDestinationMatrixOrigin:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/2873211-destinationmatrixbatchindex?language=objc
func (i_ ImageCopyToMatrix) DestinationMatrixBatchIndex() uint {
	rv := objc.Call[uint](i_, objc.Sel("destinationMatrixBatchIndex"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagecopytomatrix/2873211-destinationmatrixbatchindex?language=objc
func (i_ ImageCopyToMatrix) SetDestinationMatrixBatchIndex(value uint) {
	objc.Call[objc.Void](i_, objc.Sel("setDestinationMatrixBatchIndex:"), value)
}
