// Code generated by DarwinKit. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [VideoCompositionInstruction] class.
var VideoCompositionInstructionClass = _VideoCompositionInstructionClass{objc.GetClass("AVVideoCompositionInstruction")}

type _VideoCompositionInstructionClass struct {
	objc.Class
}

// An interface definition for the [VideoCompositionInstruction] class.
type IVideoCompositionInstruction interface {
	objc.IObject
	EnablePostProcessing() bool
	TimeRange() coremedia.TimeRange
	ContainsTweening() bool
	PassthroughTrackID() objc.Object
	RequiredSourceTrackIDs() []foundation.Value
	RequiredSourceSampleDataTrackIDs() []foundation.Number
	LayerInstructions() []VideoCompositionLayerInstruction
	BackgroundColor() coregraphics.ColorRef
}

// An operation that a compositor performs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositioninstruction?language=objc
type VideoCompositionInstruction struct {
	objc.Object
}

func VideoCompositionInstructionFrom(ptr unsafe.Pointer) VideoCompositionInstruction {
	return VideoCompositionInstruction{
		Object: objc.ObjectFrom(ptr),
	}
}

func (vc _VideoCompositionInstructionClass) Alloc() VideoCompositionInstruction {
	rv := objc.Call[VideoCompositionInstruction](vc, objc.Sel("alloc"))
	return rv
}

func (vc _VideoCompositionInstructionClass) New() VideoCompositionInstruction {
	rv := objc.Call[VideoCompositionInstruction](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewVideoCompositionInstruction() VideoCompositionInstruction {
	return VideoCompositionInstructionClass.New()
}

func (v_ VideoCompositionInstruction) Init() VideoCompositionInstruction {
	rv := objc.Call[VideoCompositionInstruction](v_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether the composition enables post-processing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/1386654-avvideocompositioninstruction/1386216-enablepostprocessing?language=objc
func (v_ VideoCompositionInstruction) EnablePostProcessing() bool {
	rv := objc.Call[bool](v_, objc.Sel("enablePostProcessing"))
	return rv
}

// The time range during which the instruction is effective. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/1386654-avvideocompositioninstruction/1389873-timerange?language=objc
func (v_ VideoCompositionInstruction) TimeRange() coremedia.TimeRange {
	rv := objc.Call[coremedia.TimeRange](v_, objc.Sel("timeRange"))
	return rv
}

// A Boolean value that indicates whether the composition contains tweening. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/1386654-avvideocompositioninstruction/1389376-containstweening?language=objc
func (v_ VideoCompositionInstruction) ContainsTweening() bool {
	rv := objc.Call[bool](v_, objc.Sel("containsTweening"))
	return rv
}

// An identifier of a source track to pass through without compositing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/1386654-avvideocompositioninstruction/1389919-passthroughtrackid?language=objc
func (v_ VideoCompositionInstruction) PassthroughTrackID() objc.Object {
	rv := objc.Call[objc.Object](v_, objc.Sel("passthroughTrackID"))
	return rv
}

// The identifiers of the video tracks the instruction requires to compose frames. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/1386654-avvideocompositioninstruction/1388661-requiredsourcetrackids?language=objc
func (v_ VideoCompositionInstruction) RequiredSourceTrackIDs() []foundation.Value {
	rv := objc.Call[[]foundation.Value](v_, objc.Sel("requiredSourceTrackIDs"))
	return rv
}

// The identifiers of the sample data tracks the instruction requires to compose frames. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/1386654-avvideocompositioninstruction/3750315-requiredsourcesampledatatrackids?language=objc
func (v_ VideoCompositionInstruction) RequiredSourceSampleDataTrackIDs() []foundation.Number {
	rv := objc.Call[[]foundation.Number](v_, objc.Sel("requiredSourceSampleDataTrackIDs"))
	return rv
}

// Instructions that specify how to layer and compose video frames from source tracks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositioninstruction/1389689-layerinstructions?language=objc
func (v_ VideoCompositionInstruction) LayerInstructions() []VideoCompositionLayerInstruction {
	rv := objc.Call[[]VideoCompositionLayerInstruction](v_, objc.Sel("layerInstructions"))
	return rv
}

// The background color of the composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocompositioninstruction/1389384-backgroundcolor?language=objc
func (v_ VideoCompositionInstruction) BackgroundColor() coregraphics.ColorRef {
	rv := objc.Call[coregraphics.ColorRef](v_, objc.Sel("backgroundColor"))
	return rv
}
