// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PointerArray] class.
var PointerArrayClass = _PointerArrayClass{objc.GetClass("NSPointerArray")}

type _PointerArrayClass struct {
	objc.Class
}

// An interface definition for the [PointerArray] class.
type IPointerArray interface {
	objc.IObject
	RemovePointerAtIndex(index uint)
	PointerAtIndex(index uint) unsafe.Pointer
	Compact()
	InsertPointerAtIndex(item unsafe.Pointer, index uint)
	ReplacePointerAtIndexWithPointer(index uint, item unsafe.Pointer)
	AddPointer(pointer unsafe.Pointer)
	PointerFunctions() PointerFunctions
	Count() uint
	SetCount(value uint)
	AllObjects() []objc.Object
}

// A collection similar to an array, but with a broader range of available memory semantics. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerarray?language=objc
type PointerArray struct {
	objc.Object
}

func PointerArrayFrom(ptr unsafe.Pointer) PointerArray {
	return PointerArray{
		Object: objc.ObjectFrom(ptr),
	}
}

func (p_ PointerArray) InitWithPointerFunctions(functions IPointerFunctions) PointerArray {
	rv := objc.Call[PointerArray](p_, objc.Sel("initWithPointerFunctions:"), objc.Ptr(functions))
	return rv
}

// Initializes the receiver to use the given functions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerarray/1416727-initwithpointerfunctions?language=objc
func NewPointerArrayWithPointerFunctions(functions IPointerFunctions) PointerArray {
	instance := PointerArrayClass.Alloc().InitWithPointerFunctions(functions)
	instance.Autorelease()
	return instance
}

func (p_ PointerArray) InitWithOptions(options PointerFunctionsOptions) PointerArray {
	rv := objc.Call[PointerArray](p_, objc.Sel("initWithOptions:"), options)
	return rv
}

// Initializes the receiver to use the given options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerarray/1408229-initwithoptions?language=objc
func NewPointerArrayWithOptions(options PointerFunctionsOptions) PointerArray {
	instance := PointerArrayClass.Alloc().InitWithOptions(options)
	instance.Autorelease()
	return instance
}

func (pc _PointerArrayClass) Alloc() PointerArray {
	rv := objc.Call[PointerArray](pc, objc.Sel("alloc"))
	return rv
}

func PointerArray_Alloc() PointerArray {
	return PointerArrayClass.Alloc()
}

func (pc _PointerArrayClass) New() PointerArray {
	rv := objc.Call[PointerArray](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPointerArray() PointerArray {
	return PointerArrayClass.New()
}

func (p_ PointerArray) Init() PointerArray {
	rv := objc.Call[PointerArray](p_, objc.Sel("init"))
	return rv
}

// Removes the pointer at a given index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerarray/1407403-removepointeratindex?language=objc
func (p_ PointerArray) RemovePointerAtIndex(index uint) {
	objc.Call[objc.Void](p_, objc.Sel("removePointerAtIndex:"), index)
}

// Returns a new pointer array initialized to use the given options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerarray/1564845-pointerarraywithoptions?language=objc
func (pc _PointerArrayClass) PointerArrayWithOptions(options PointerFunctionsOptions) PointerArray {
	rv := objc.Call[PointerArray](pc, objc.Sel("pointerArrayWithOptions:"), options)
	return rv
}

// Returns a new pointer array initialized to use the given options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerarray/1564845-pointerarraywithoptions?language=objc
func PointerArray_PointerArrayWithOptions(options PointerFunctionsOptions) PointerArray {
	return PointerArrayClass.PointerArrayWithOptions(options)
}

// Returns the pointer at a given index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerarray/1417655-pointeratindex?language=objc
func (p_ PointerArray) PointerAtIndex(index uint) unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](p_, objc.Sel("pointerAtIndex:"), index)
	return rv
}

// Removes NULL values from the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerarray/1413659-compact?language=objc
func (p_ PointerArray) Compact() {
	objc.Call[objc.Void](p_, objc.Sel("compact"))
}

// A new pointer array initialized to use the given functions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerarray/1564844-pointerarraywithpointerfunctions?language=objc
func (pc _PointerArrayClass) PointerArrayWithPointerFunctions(functions IPointerFunctions) PointerArray {
	rv := objc.Call[PointerArray](pc, objc.Sel("pointerArrayWithPointerFunctions:"), objc.Ptr(functions))
	return rv
}

// A new pointer array initialized to use the given functions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerarray/1564844-pointerarraywithpointerfunctions?language=objc
func PointerArray_PointerArrayWithPointerFunctions(functions IPointerFunctions) PointerArray {
	return PointerArrayClass.PointerArrayWithPointerFunctions(functions)
}

// Inserts a pointer at a given index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerarray/1408589-insertpointer?language=objc
func (p_ PointerArray) InsertPointerAtIndex(item unsafe.Pointer, index uint) {
	objc.Call[objc.Void](p_, objc.Sel("insertPointer:atIndex:"), item, index)
}

// Replaces the pointer at a given index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerarray/1412654-replacepointeratindex?language=objc
func (p_ PointerArray) ReplacePointerAtIndexWithPointer(index uint, item unsafe.Pointer) {
	objc.Call[objc.Void](p_, objc.Sel("replacePointerAtIndex:withPointer:"), index, item)
}

// Returns a new pointer array that maintains weak references to its elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerarray/1412795-weakobjectspointerarray?language=objc
func (pc _PointerArrayClass) WeakObjectsPointerArray() PointerArray {
	rv := objc.Call[PointerArray](pc, objc.Sel("weakObjectsPointerArray"))
	return rv
}

// Returns a new pointer array that maintains weak references to its elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerarray/1412795-weakobjectspointerarray?language=objc
func PointerArray_WeakObjectsPointerArray() PointerArray {
	return PointerArrayClass.WeakObjectsPointerArray()
}

// Returns a new pointer array that maintains strong references to its elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerarray/1413102-strongobjectspointerarray?language=objc
func (pc _PointerArrayClass) StrongObjectsPointerArray() PointerArray {
	rv := objc.Call[PointerArray](pc, objc.Sel("strongObjectsPointerArray"))
	return rv
}

// Returns a new pointer array that maintains strong references to its elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerarray/1413102-strongobjectspointerarray?language=objc
func PointerArray_StrongObjectsPointerArray() PointerArray {
	return PointerArrayClass.StrongObjectsPointerArray()
}

// Adds a given pointer to the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerarray/1411636-addpointer?language=objc
func (p_ PointerArray) AddPointer(pointer unsafe.Pointer) {
	objc.Call[objc.Void](p_, objc.Sel("addPointer:"), pointer)
}

// The functions in use by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerarray/1413669-pointerfunctions?language=objc
func (p_ PointerArray) PointerFunctions() PointerFunctions {
	rv := objc.Call[PointerFunctions](p_, objc.Sel("pointerFunctions"))
	return rv
}

// The number of elements in the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerarray/1418453-count?language=objc
func (p_ PointerArray) Count() uint {
	rv := objc.Call[uint](p_, objc.Sel("count"))
	return rv
}

// The number of elements in the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerarray/1418453-count?language=objc
func (p_ PointerArray) SetCount(value uint) {
	objc.Call[objc.Void](p_, objc.Sel("setCount:"), value)
}

// All the objects in the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerarray/1408081-allobjects?language=objc
func (p_ PointerArray) AllObjects() []objc.Object {
	rv := objc.Call[[]objc.Object](p_, objc.Sel("allObjects"))
	return rv
}