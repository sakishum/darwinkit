// Code generated by DarwinKit. DO NOT EDIT.

package mpsgraph

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Convolution2DOpDescriptor] class.
var Convolution2DOpDescriptorClass = _Convolution2DOpDescriptorClass{objc.GetClass("MPSGraphConvolution2DOpDescriptor")}

type _Convolution2DOpDescriptorClass struct {
	objc.Class
}

// An interface definition for the [Convolution2DOpDescriptor] class.
type IConvolution2DOpDescriptor interface {
	objc.IObject
	SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom(paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint)
	Groups() uint
	SetGroups(value uint)
	DilationRateInY() uint
	SetDilationRateInY(value uint)
	DataLayout() TensorNamedDataLayout
	SetDataLayout(value TensorNamedDataLayout)
	PaddingStyle() PaddingStyle
	SetPaddingStyle(value PaddingStyle)
	PaddingBottom() uint
	SetPaddingBottom(value uint)
	WeightsLayout() TensorNamedDataLayout
	SetWeightsLayout(value TensorNamedDataLayout)
	PaddingTop() uint
	SetPaddingTop(value uint)
	StrideInX() uint
	SetStrideInX(value uint)
	StrideInY() uint
	SetStrideInY(value uint)
	PaddingLeft() uint
	SetPaddingLeft(value uint)
	PaddingRight() uint
	SetPaddingRight(value uint)
	DilationRateInX() uint
	SetDilationRateInX(value uint)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor?language=objc
type Convolution2DOpDescriptor struct {
	objc.Object
}

func Convolution2DOpDescriptorFrom(ptr unsafe.Pointer) Convolution2DOpDescriptor {
	return Convolution2DOpDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _Convolution2DOpDescriptorClass) DescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYGroupsPaddingStyleDataLayoutWeightsLayout(strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, groups uint, paddingStyle PaddingStyle, dataLayout TensorNamedDataLayout, weightsLayout TensorNamedDataLayout) Convolution2DOpDescriptor {
	rv := objc.Call[Convolution2DOpDescriptor](cc, objc.Sel("descriptorWithStrideInX:strideInY:dilationRateInX:dilationRateInY:groups:paddingStyle:dataLayout:weightsLayout:"), strideInX, strideInY, dilationRateInX, dilationRateInY, groups, paddingStyle, dataLayout, weightsLayout)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/3564600-descriptorwithstrideinx?language=objc
func Convolution2DOpDescriptor_DescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYGroupsPaddingStyleDataLayoutWeightsLayout(strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, groups uint, paddingStyle PaddingStyle, dataLayout TensorNamedDataLayout, weightsLayout TensorNamedDataLayout) Convolution2DOpDescriptor {
	return Convolution2DOpDescriptorClass.DescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYGroupsPaddingStyleDataLayoutWeightsLayout(strideInX, strideInY, dilationRateInX, dilationRateInY, groups, paddingStyle, dataLayout, weightsLayout)
}

func (cc _Convolution2DOpDescriptorClass) DescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYGroupsPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayoutWeightsLayout(strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, groups uint, paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, paddingStyle PaddingStyle, dataLayout TensorNamedDataLayout, weightsLayout TensorNamedDataLayout) Convolution2DOpDescriptor {
	rv := objc.Call[Convolution2DOpDescriptor](cc, objc.Sel("descriptorWithStrideInX:strideInY:dilationRateInX:dilationRateInY:groups:paddingLeft:paddingRight:paddingTop:paddingBottom:paddingStyle:dataLayout:weightsLayout:"), strideInX, strideInY, dilationRateInX, dilationRateInY, groups, paddingLeft, paddingRight, paddingTop, paddingBottom, paddingStyle, dataLayout, weightsLayout)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/3564599-descriptorwithstrideinx?language=objc
func Convolution2DOpDescriptor_DescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYGroupsPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayoutWeightsLayout(strideInX uint, strideInY uint, dilationRateInX uint, dilationRateInY uint, groups uint, paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint, paddingStyle PaddingStyle, dataLayout TensorNamedDataLayout, weightsLayout TensorNamedDataLayout) Convolution2DOpDescriptor {
	return Convolution2DOpDescriptorClass.DescriptorWithStrideInXStrideInYDilationRateInXDilationRateInYGroupsPaddingLeftPaddingRightPaddingTopPaddingBottomPaddingStyleDataLayoutWeightsLayout(strideInX, strideInY, dilationRateInX, dilationRateInY, groups, paddingLeft, paddingRight, paddingTop, paddingBottom, paddingStyle, dataLayout, weightsLayout)
}

func (cc _Convolution2DOpDescriptorClass) Alloc() Convolution2DOpDescriptor {
	rv := objc.Call[Convolution2DOpDescriptor](cc, objc.Sel("alloc"))
	return rv
}

func (cc _Convolution2DOpDescriptorClass) New() Convolution2DOpDescriptor {
	rv := objc.Call[Convolution2DOpDescriptor](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewConvolution2DOpDescriptor() Convolution2DOpDescriptor {
	return Convolution2DOpDescriptorClass.New()
}

func (c_ Convolution2DOpDescriptor) Init() Convolution2DOpDescriptor {
	rv := objc.Call[Convolution2DOpDescriptor](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/3564609-setexplicitpaddingwithpaddinglef?language=objc
func (c_ Convolution2DOpDescriptor) SetExplicitPaddingWithPaddingLeftPaddingRightPaddingTopPaddingBottom(paddingLeft uint, paddingRight uint, paddingTop uint, paddingBottom uint) {
	objc.Call[objc.Void](c_, objc.Sel("setExplicitPaddingWithPaddingLeft:paddingRight:paddingTop:paddingBottom:"), paddingLeft, paddingRight, paddingTop, paddingBottom)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/3564603-groups?language=objc
func (c_ Convolution2DOpDescriptor) Groups() uint {
	rv := objc.Call[uint](c_, objc.Sel("groups"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/3564603-groups?language=objc
func (c_ Convolution2DOpDescriptor) SetGroups(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setGroups:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/3564602-dilationrateiny?language=objc
func (c_ Convolution2DOpDescriptor) DilationRateInY() uint {
	rv := objc.Call[uint](c_, objc.Sel("dilationRateInY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/3564602-dilationrateiny?language=objc
func (c_ Convolution2DOpDescriptor) SetDilationRateInY(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setDilationRateInY:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/3564598-datalayout?language=objc
func (c_ Convolution2DOpDescriptor) DataLayout() TensorNamedDataLayout {
	rv := objc.Call[TensorNamedDataLayout](c_, objc.Sel("dataLayout"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/3564598-datalayout?language=objc
func (c_ Convolution2DOpDescriptor) SetDataLayout(value TensorNamedDataLayout) {
	objc.Call[objc.Void](c_, objc.Sel("setDataLayout:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/3564607-paddingstyle?language=objc
func (c_ Convolution2DOpDescriptor) PaddingStyle() PaddingStyle {
	rv := objc.Call[PaddingStyle](c_, objc.Sel("paddingStyle"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/3564607-paddingstyle?language=objc
func (c_ Convolution2DOpDescriptor) SetPaddingStyle(value PaddingStyle) {
	objc.Call[objc.Void](c_, objc.Sel("setPaddingStyle:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/3564604-paddingbottom?language=objc
func (c_ Convolution2DOpDescriptor) PaddingBottom() uint {
	rv := objc.Call[uint](c_, objc.Sel("paddingBottom"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/3564604-paddingbottom?language=objc
func (c_ Convolution2DOpDescriptor) SetPaddingBottom(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setPaddingBottom:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/3564612-weightslayout?language=objc
func (c_ Convolution2DOpDescriptor) WeightsLayout() TensorNamedDataLayout {
	rv := objc.Call[TensorNamedDataLayout](c_, objc.Sel("weightsLayout"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/3564612-weightslayout?language=objc
func (c_ Convolution2DOpDescriptor) SetWeightsLayout(value TensorNamedDataLayout) {
	objc.Call[objc.Void](c_, objc.Sel("setWeightsLayout:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/3564608-paddingtop?language=objc
func (c_ Convolution2DOpDescriptor) PaddingTop() uint {
	rv := objc.Call[uint](c_, objc.Sel("paddingTop"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/3564608-paddingtop?language=objc
func (c_ Convolution2DOpDescriptor) SetPaddingTop(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setPaddingTop:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/3564610-strideinx?language=objc
func (c_ Convolution2DOpDescriptor) StrideInX() uint {
	rv := objc.Call[uint](c_, objc.Sel("strideInX"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/3564610-strideinx?language=objc
func (c_ Convolution2DOpDescriptor) SetStrideInX(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setStrideInX:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/3564611-strideiny?language=objc
func (c_ Convolution2DOpDescriptor) StrideInY() uint {
	rv := objc.Call[uint](c_, objc.Sel("strideInY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/3564611-strideiny?language=objc
func (c_ Convolution2DOpDescriptor) SetStrideInY(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setStrideInY:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/3564605-paddingleft?language=objc
func (c_ Convolution2DOpDescriptor) PaddingLeft() uint {
	rv := objc.Call[uint](c_, objc.Sel("paddingLeft"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/3564605-paddingleft?language=objc
func (c_ Convolution2DOpDescriptor) SetPaddingLeft(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setPaddingLeft:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/3564606-paddingright?language=objc
func (c_ Convolution2DOpDescriptor) PaddingRight() uint {
	rv := objc.Call[uint](c_, objc.Sel("paddingRight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/3564606-paddingright?language=objc
func (c_ Convolution2DOpDescriptor) SetPaddingRight(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setPaddingRight:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/3564601-dilationrateinx?language=objc
func (c_ Convolution2DOpDescriptor) DilationRateInX() uint {
	rv := objc.Call[uint](c_, objc.Sel("dilationRateInX"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphconvolution2dopdescriptor/3564601-dilationrateinx?language=objc
func (c_ Convolution2DOpDescriptor) SetDilationRateInX(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setDilationRateInX:"), value)
}
