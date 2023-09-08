// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNConvolutionTransposeGradientState] class.
var CNNConvolutionTransposeGradientStateClass = _CNNConvolutionTransposeGradientStateClass{objc.GetClass("MPSCNNConvolutionTransposeGradientState")}

type _CNNConvolutionTransposeGradientStateClass struct {
	objc.Class
}

// An interface definition for the [CNNConvolutionTransposeGradientState] class.
type ICNNConvolutionTransposeGradientState interface {
	ICNNConvolutionGradientState
	ConvolutionTranspose() CNNConvolutionTranspose
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradientstate?language=objc
type CNNConvolutionTransposeGradientState struct {
	CNNConvolutionGradientState
}

func CNNConvolutionTransposeGradientStateFrom(ptr unsafe.Pointer) CNNConvolutionTransposeGradientState {
	return CNNConvolutionTransposeGradientState{
		CNNConvolutionGradientState: CNNConvolutionGradientStateFrom(ptr),
	}
}

func (cc _CNNConvolutionTransposeGradientStateClass) Alloc() CNNConvolutionTransposeGradientState {
	rv := objc.Call[CNNConvolutionTransposeGradientState](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CNNConvolutionTransposeGradientStateClass) New() CNNConvolutionTransposeGradientState {
	rv := objc.Call[CNNConvolutionTransposeGradientState](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNConvolutionTransposeGradientState() CNNConvolutionTransposeGradientState {
	return CNNConvolutionTransposeGradientStateClass.New()
}

func (c_ CNNConvolutionTransposeGradientState) Init() CNNConvolutionTransposeGradientState {
	rv := objc.Call[CNNConvolutionTransposeGradientState](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNConvolutionTransposeGradientStateClass) TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) CNNConvolutionTransposeGradientState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[CNNConvolutionTransposeGradientState](cc, objc.Sel("temporaryStateWithCommandBuffer:resourceList:"), po0, objc.Ptr(resourceList))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947915-temporarystatewithcommandbuffer?language=objc
func CNNConvolutionTransposeGradientState_TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) CNNConvolutionTransposeGradientState {
	return CNNConvolutionTransposeGradientStateClass.TemporaryStateWithCommandBufferResourceList(commandBuffer, resourceList)
}

func (c_ CNNConvolutionTransposeGradientState) InitWithDeviceTextureDescriptor(device metal.PDevice, descriptor metal.ITextureDescriptor) CNNConvolutionTransposeGradientState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNConvolutionTransposeGradientState](c_, objc.Sel("initWithDevice:textureDescriptor:"), po0, objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942400-initwithdevice?language=objc
func NewCNNConvolutionTransposeGradientStateWithDeviceTextureDescriptor(device metal.PDevice, descriptor metal.ITextureDescriptor) CNNConvolutionTransposeGradientState {
	instance := CNNConvolutionTransposeGradientStateClass.Alloc().InitWithDeviceTextureDescriptor(device, descriptor)
	instance.Autorelease()
	return instance
}

func (c_ CNNConvolutionTransposeGradientState) InitWithResource(resource metal.PResource) CNNConvolutionTransposeGradientState {
	po0 := objc.WrapAsProtocol("MTLResource", resource)
	rv := objc.Call[CNNConvolutionTransposeGradientState](c_, objc.Sel("initWithResource:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942390-initwithresource?language=objc
func NewCNNConvolutionTransposeGradientStateWithResource(resource metal.PResource) CNNConvolutionTransposeGradientState {
	instance := CNNConvolutionTransposeGradientStateClass.Alloc().InitWithResource(resource)
	instance.Autorelease()
	return instance
}

func (c_ CNNConvolutionTransposeGradientState) InitWithDeviceBufferSize(device metal.PDevice, bufferSize uint) CNNConvolutionTransposeGradientState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNConvolutionTransposeGradientState](c_, objc.Sel("initWithDevice:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942392-initwithdevice?language=objc
func NewCNNConvolutionTransposeGradientStateWithDeviceBufferSize(device metal.PDevice, bufferSize uint) CNNConvolutionTransposeGradientState {
	instance := CNNConvolutionTransposeGradientStateClass.Alloc().InitWithDeviceBufferSize(device, bufferSize)
	instance.Autorelease()
	return instance
}

func (c_ CNNConvolutionTransposeGradientState) InitWithResources(resources []metal.PResource) CNNConvolutionTransposeGradientState {
	rv := objc.Call[CNNConvolutionTransposeGradientState](c_, objc.Sel("initWithResources:"), resources)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947895-initwithresources?language=objc
func NewCNNConvolutionTransposeGradientStateWithResources(resources []metal.PResource) CNNConvolutionTransposeGradientState {
	instance := CNNConvolutionTransposeGradientStateClass.Alloc().InitWithResources(resources)
	instance.Autorelease()
	return instance
}

func (cc _CNNConvolutionTransposeGradientStateClass) TemporaryStateWithCommandBufferBufferSize(cmdBuf metal.PCommandBuffer, bufferSize uint) CNNConvolutionTransposeGradientState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	rv := objc.Call[CNNConvolutionTransposeGradientState](cc, objc.Sel("temporaryStateWithCommandBuffer:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942391-temporarystatewithcommandbuffer?language=objc
func CNNConvolutionTransposeGradientState_TemporaryStateWithCommandBufferBufferSize(cmdBuf metal.PCommandBuffer, bufferSize uint) CNNConvolutionTransposeGradientState {
	return CNNConvolutionTransposeGradientStateClass.TemporaryStateWithCommandBufferBufferSize(cmdBuf, bufferSize)
}

func (cc _CNNConvolutionTransposeGradientStateClass) TemporaryStateWithCommandBuffer(cmdBuf metal.PCommandBuffer) CNNConvolutionTransposeGradientState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	rv := objc.Call[CNNConvolutionTransposeGradientState](cc, objc.Sel("temporaryStateWithCommandBuffer:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942393-temporarystatewithcommandbuffer?language=objc
func CNNConvolutionTransposeGradientState_TemporaryStateWithCommandBuffer(cmdBuf metal.PCommandBuffer) CNNConvolutionTransposeGradientState {
	return CNNConvolutionTransposeGradientStateClass.TemporaryStateWithCommandBuffer(cmdBuf)
}

func (cc _CNNConvolutionTransposeGradientStateClass) TemporaryStateWithCommandBufferTextureDescriptor(cmdBuf metal.PCommandBuffer, descriptor metal.ITextureDescriptor) CNNConvolutionTransposeGradientState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	rv := objc.Call[CNNConvolutionTransposeGradientState](cc, objc.Sel("temporaryStateWithCommandBuffer:textureDescriptor:"), po0, objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942395-temporarystatewithcommandbuffer?language=objc
func CNNConvolutionTransposeGradientState_TemporaryStateWithCommandBufferTextureDescriptor(cmdBuf metal.PCommandBuffer, descriptor metal.ITextureDescriptor) CNNConvolutionTransposeGradientState {
	return CNNConvolutionTransposeGradientStateClass.TemporaryStateWithCommandBufferTextureDescriptor(cmdBuf, descriptor)
}

func (c_ CNNConvolutionTransposeGradientState) InitWithDeviceResourceList(device metal.PDevice, resourceList IStateResourceList) CNNConvolutionTransposeGradientState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNConvolutionTransposeGradientState](c_, objc.Sel("initWithDevice:resourceList:"), po0, objc.Ptr(resourceList))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947908-initwithdevice?language=objc
func NewCNNConvolutionTransposeGradientStateWithDeviceResourceList(device metal.PDevice, resourceList IStateResourceList) CNNConvolutionTransposeGradientState {
	instance := CNNConvolutionTransposeGradientStateClass.Alloc().InitWithDeviceResourceList(device, resourceList)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradientstate/3131790-convolutiontranspose?language=objc
func (c_ CNNConvolutionTransposeGradientState) ConvolutionTranspose() CNNConvolutionTranspose {
	rv := objc.Call[CNNConvolutionTranspose](c_, objc.Sel("convolutionTranspose"))
	return rv
}
