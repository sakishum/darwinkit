// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MutableOrderedSet] class.
var MutableOrderedSetClass = _MutableOrderedSetClass{objc.GetClass("NSMutableOrderedSet")}

type _MutableOrderedSetClass struct {
	objc.Class
}

// An interface definition for the [MutableOrderedSet] class.
type IMutableOrderedSet interface {
	IOrderedSet
	AddObject(object objc.IObject)
	RemoveObject(object objc.IObject)
	MinusOrderedSet(other IOrderedSet)
	UnionOrderedSet(other IOrderedSet)
	ReplaceObjectsInRangeWithObjectsCount(range_ Range, objects objc.IObject, count uint)
	ReplaceObjectsAtIndexesWithObjects(indexes IIndexSet, objects []objc.IObject)
	ExchangeObjectAtIndexWithObjectAtIndex(idx1 uint, idx2 uint)
	ApplyDifference(difference IOrderedCollectionDifference)
	SetObjectAtIndexedSubscript(obj objc.IObject, idx uint)
	IntersectSet(other ISet)
	IntersectOrderedSet(other IOrderedSet)
	RemoveObjectAtIndex(idx uint)
	SortUsingComparator(cmptr Comparator)
	RemoveObjectsAtIndexes(indexes IIndexSet)
	SortRangeOptionsUsingComparator(range_ Range, opts SortOptions, cmptr Comparator)
	RemoveAllObjects()
	InsertObjectsAtIndexes(objects []objc.IObject, indexes IIndexSet)
	SetObjectAtIndex(obj objc.IObject, idx uint)
	SortWithOptionsUsingComparator(opts SortOptions, cmptr Comparator)
	SortUsingDescriptors(sortDescriptors []ISortDescriptor)
	UnionSet(other ISet)
	MoveObjectsAtIndexesToIndex(indexes IIndexSet, idx uint)
	ReplaceObjectAtIndexWithObject(idx uint, object objc.IObject)
	FilterUsingPredicate(p IPredicate)
	AddObjectsFromArray(array []objc.IObject)
	RemoveObjectsInArray(array []objc.IObject)
	RemoveObjectsInRange(range_ Range)
	MinusSet(other ISet)
	AddObjectsCount(objects objc.IObject, count uint)
	InsertObjectAtIndex(object objc.IObject, idx uint)
}

// A dynamic, ordered collection of unique objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset?language=objc
type MutableOrderedSet struct {
	OrderedSet
}

func MutableOrderedSetFrom(ptr unsafe.Pointer) MutableOrderedSet {
	return MutableOrderedSet{
		OrderedSet: OrderedSetFrom(ptr),
	}
}

func (m_ MutableOrderedSet) InitWithCapacity(numItems uint) MutableOrderedSet {
	rv := objc.Call[MutableOrderedSet](m_, objc.Sel("initWithCapacity:"), numItems)
	return rv
}

// Returns an initialized mutable ordered set with a given initial capacity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1411583-initwithcapacity?language=objc
func NewMutableOrderedSetWithCapacity(numItems uint) MutableOrderedSet {
	instance := MutableOrderedSetClass.Alloc().InitWithCapacity(numItems)
	instance.Autorelease()
	return instance
}

func (mc _MutableOrderedSetClass) OrderedSetWithCapacity(numItems uint) MutableOrderedSet {
	rv := objc.Call[MutableOrderedSet](mc, objc.Sel("orderedSetWithCapacity:"), numItems)
	return rv
}

// Creates and returns an mutable ordered set with a given initial capacity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1543283-orderedsetwithcapacity?language=objc
func MutableOrderedSet_OrderedSetWithCapacity(numItems uint) MutableOrderedSet {
	return MutableOrderedSetClass.OrderedSetWithCapacity(numItems)
}

func (m_ MutableOrderedSet) Init() MutableOrderedSet {
	rv := objc.Call[MutableOrderedSet](m_, objc.Sel("init"))
	return rv
}

func (mc _MutableOrderedSetClass) Alloc() MutableOrderedSet {
	rv := objc.Call[MutableOrderedSet](mc, objc.Sel("alloc"))
	return rv
}

func (mc _MutableOrderedSetClass) New() MutableOrderedSet {
	rv := objc.Call[MutableOrderedSet](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableOrderedSet() MutableOrderedSet {
	return MutableOrderedSetClass.New()
}

func (mc _MutableOrderedSetClass) OrderedSetWithObjectsCount(objects objc.IObject, cnt uint) MutableOrderedSet {
	rv := objc.Call[MutableOrderedSet](mc, objc.Sel("orderedSetWithObjects:count:"), objc.Ptr(objects), cnt)
	return rv
}

// Creates and returns a set containing a specified number of objects from a given C array of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1543334-orderedsetwithobjects?language=objc
func MutableOrderedSet_OrderedSetWithObjectsCount(objects objc.IObject, cnt uint) MutableOrderedSet {
	return MutableOrderedSetClass.OrderedSetWithObjectsCount(objects, cnt)
}

func (m_ MutableOrderedSet) InitWithSetCopyItems(set ISet, flag bool) MutableOrderedSet {
	rv := objc.Call[MutableOrderedSet](m_, objc.Sel("initWithSet:copyItems:"), objc.Ptr(set), flag)
	return rv
}

// Initializes a new ordered set with the contents of a set, optionally copying the objects in the set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1411246-initwithset?language=objc
func NewMutableOrderedSetWithSetCopyItems(set ISet, flag bool) MutableOrderedSet {
	instance := MutableOrderedSetClass.Alloc().InitWithSetCopyItems(set, flag)
	instance.Autorelease()
	return instance
}

func (mc _MutableOrderedSetClass) OrderedSetWithSetCopyItems(set ISet, flag bool) MutableOrderedSet {
	rv := objc.Call[MutableOrderedSet](mc, objc.Sel("orderedSetWithSet:copyItems:"), objc.Ptr(set), flag)
	return rv
}

// Creates and returns an ordered set with the contents of a set, optionally copying the items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1543331-orderedsetwithset?language=objc
func MutableOrderedSet_OrderedSetWithSetCopyItems(set ISet, flag bool) MutableOrderedSet {
	return MutableOrderedSetClass.OrderedSetWithSetCopyItems(set, flag)
}

func (m_ MutableOrderedSet) InitWithObject(object objc.IObject) MutableOrderedSet {
	rv := objc.Call[MutableOrderedSet](m_, objc.Sel("initWithObject:"), objc.Ptr(object))
	return rv
}

// Initializes a new ordered set with the object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1413883-initwithobject?language=objc
func NewMutableOrderedSetWithObject(object objc.IObject) MutableOrderedSet {
	instance := MutableOrderedSetClass.Alloc().InitWithObject(object)
	instance.Autorelease()
	return instance
}

func (m_ MutableOrderedSet) InitWithArrayCopyItems(set []objc.IObject, flag bool) MutableOrderedSet {
	rv := objc.Call[MutableOrderedSet](m_, objc.Sel("initWithArray:copyItems:"), set, flag)
	return rv
}

// Initializes a newly allocated set with the objects that are contained in a given array, optionally copying the items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1418006-initwitharray?language=objc
func NewMutableOrderedSetWithArrayCopyItems(set []objc.IObject, flag bool) MutableOrderedSet {
	instance := MutableOrderedSetClass.Alloc().InitWithArrayCopyItems(set, flag)
	instance.Autorelease()
	return instance
}

func (mc _MutableOrderedSetClass) OrderedSetWithArray(array []objc.IObject) MutableOrderedSet {
	rv := objc.Call[MutableOrderedSet](mc, objc.Sel("orderedSetWithArray:"), array)
	return rv
}

// Creates and returns a set containing a uniqued collection of the objects contained in a given array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1543310-orderedsetwitharray?language=objc
func MutableOrderedSet_OrderedSetWithArray(array []objc.IObject) MutableOrderedSet {
	return MutableOrderedSetClass.OrderedSetWithArray(array)
}

func (m_ MutableOrderedSet) InitWithSet(set ISet) MutableOrderedSet {
	rv := objc.Call[MutableOrderedSet](m_, objc.Sel("initWithSet:"), objc.Ptr(set))
	return rv
}

// Initializes a new ordered set with the contents of a set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1416344-initwithset?language=objc
func NewMutableOrderedSetWithSet(set ISet) MutableOrderedSet {
	instance := MutableOrderedSetClass.Alloc().InitWithSet(set)
	instance.Autorelease()
	return instance
}

func (mc _MutableOrderedSetClass) OrderedSetWithOrderedSet(set IOrderedSet) MutableOrderedSet {
	rv := objc.Call[MutableOrderedSet](mc, objc.Sel("orderedSetWithOrderedSet:"), objc.Ptr(set))
	return rv
}

// Creates and returns an ordered set containing the objects from another ordered set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1543280-orderedsetwithorderedset?language=objc
func MutableOrderedSet_OrderedSetWithOrderedSet(set IOrderedSet) MutableOrderedSet {
	return MutableOrderedSetClass.OrderedSetWithOrderedSet(set)
}

func (mc _MutableOrderedSetClass) OrderedSetWithObject(object objc.IObject) MutableOrderedSet {
	rv := objc.Call[MutableOrderedSet](mc, objc.Sel("orderedSetWithObject:"), objc.Ptr(object))
	return rv
}

// Creates and returns a ordered set that contains a single given object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1543339-orderedsetwithobject?language=objc
func MutableOrderedSet_OrderedSetWithObject(object objc.IObject) MutableOrderedSet {
	return MutableOrderedSetClass.OrderedSetWithObject(object)
}

func (m_ MutableOrderedSet) InitWithOrderedSetRangeCopyItems(set IOrderedSet, range_ Range, flag bool) MutableOrderedSet {
	rv := objc.Call[MutableOrderedSet](m_, objc.Sel("initWithOrderedSet:range:copyItems:"), objc.Ptr(set), range_, flag)
	return rv
}

// Initializes a new ordered set with the contents of an ordered set, optionally copying the items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1417751-initwithorderedset?language=objc
func NewMutableOrderedSetWithOrderedSetRangeCopyItems(set IOrderedSet, range_ Range, flag bool) MutableOrderedSet {
	instance := MutableOrderedSetClass.Alloc().InitWithOrderedSetRangeCopyItems(set, range_, flag)
	instance.Autorelease()
	return instance
}

func (mc _MutableOrderedSetClass) OrderedSet() MutableOrderedSet {
	rv := objc.Call[MutableOrderedSet](mc, objc.Sel("orderedSet"))
	return rv
}

// Creates and returns an empty ordered set [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1543313-orderedset?language=objc
func MutableOrderedSet_OrderedSet() MutableOrderedSet {
	return MutableOrderedSetClass.OrderedSet()
}

func (mc _MutableOrderedSetClass) OrderedSetWithOrderedSetRangeCopyItems(set IOrderedSet, range_ Range, flag bool) MutableOrderedSet {
	rv := objc.Call[MutableOrderedSet](mc, objc.Sel("orderedSetWithOrderedSet:range:copyItems:"), objc.Ptr(set), range_, flag)
	return rv
}

// Creates and returns a new ordered set for a specified range of objects in an ordered set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1543292-orderedsetwithorderedset?language=objc
func MutableOrderedSet_OrderedSetWithOrderedSetRangeCopyItems(set IOrderedSet, range_ Range, flag bool) MutableOrderedSet {
	return MutableOrderedSetClass.OrderedSetWithOrderedSetRangeCopyItems(set, range_, flag)
}

func (mc _MutableOrderedSetClass) OrderedSetWithArrayRangeCopyItems(array []objc.IObject, range_ Range, flag bool) MutableOrderedSet {
	rv := objc.Call[MutableOrderedSet](mc, objc.Sel("orderedSetWithArray:range:copyItems:"), array, range_, flag)
	return rv
}

// Creates and returns a new ordered set for a specified range of objects in an array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1543321-orderedsetwitharray?language=objc
func MutableOrderedSet_OrderedSetWithArrayRangeCopyItems(array []objc.IObject, range_ Range, flag bool) MutableOrderedSet {
	return MutableOrderedSetClass.OrderedSetWithArrayRangeCopyItems(array, range_, flag)
}

func (m_ MutableOrderedSet) InitWithOrderedSet(set IOrderedSet) MutableOrderedSet {
	rv := objc.Call[MutableOrderedSet](m_, objc.Sel("initWithOrderedSet:"), objc.Ptr(set))
	return rv
}

// Initializes a new ordered set with the contents of a set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1412402-initwithorderedset?language=objc
func NewMutableOrderedSetWithOrderedSet(set IOrderedSet) MutableOrderedSet {
	instance := MutableOrderedSetClass.Alloc().InitWithOrderedSet(set)
	instance.Autorelease()
	return instance
}

func (m_ MutableOrderedSet) InitWithOrderedSetCopyItems(set IOrderedSet, flag bool) MutableOrderedSet {
	rv := objc.Call[MutableOrderedSet](m_, objc.Sel("initWithOrderedSet:copyItems:"), objc.Ptr(set), flag)
	return rv
}

// Initializes a new ordered set with the contents of a set, optionally copying the items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1411658-initwithorderedset?language=objc
func NewMutableOrderedSetWithOrderedSetCopyItems(set IOrderedSet, flag bool) MutableOrderedSet {
	instance := MutableOrderedSetClass.Alloc().InitWithOrderedSetCopyItems(set, flag)
	instance.Autorelease()
	return instance
}

func (m_ MutableOrderedSet) InitWithArray(array []objc.IObject) MutableOrderedSet {
	rv := objc.Call[MutableOrderedSet](m_, objc.Sel("initWithArray:"), array)
	return rv
}

// Initializes a newly allocated set with the objects that are contained in a given array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1408623-initwitharray?language=objc
func NewMutableOrderedSetWithArray(array []objc.IObject) MutableOrderedSet {
	instance := MutableOrderedSetClass.Alloc().InitWithArray(array)
	instance.Autorelease()
	return instance
}

func (m_ MutableOrderedSet) InitWithArrayRangeCopyItems(set []objc.IObject, range_ Range, flag bool) MutableOrderedSet {
	rv := objc.Call[MutableOrderedSet](m_, objc.Sel("initWithArray:range:copyItems:"), set, range_, flag)
	return rv
}

// Initializes a newly allocated set with the objects that are contained in the specified range of an array, optionally copying the items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1409272-initwitharray?language=objc
func NewMutableOrderedSetWithArrayRangeCopyItems(set []objc.IObject, range_ Range, flag bool) MutableOrderedSet {
	instance := MutableOrderedSetClass.Alloc().InitWithArrayRangeCopyItems(set, range_, flag)
	instance.Autorelease()
	return instance
}

func (m_ MutableOrderedSet) InitWithObjects(firstObj objc.IObject, args ...any) MutableOrderedSet {
	rv := objc.Call[MutableOrderedSet](m_, objc.Sel("initWithObjects:"), append([]any{objc.Ptr(firstObj)}, args...)...)
	return rv
}

// Initializes a newly allocated set with members taken from the specified list of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1543287-initwithobjects?language=objc
func NewMutableOrderedSetWithObjects(firstObj objc.IObject, args ...any) MutableOrderedSet {
	instance := MutableOrderedSetClass.Alloc().InitWithObjects(firstObj, args...)
	instance.Autorelease()
	return instance
}

func (mc _MutableOrderedSetClass) OrderedSetWithObjects(firstObj objc.IObject, args ...any) MutableOrderedSet {
	rv := objc.Call[MutableOrderedSet](mc, objc.Sel("orderedSetWithObjects:"), append([]any{objc.Ptr(firstObj)}, args...)...)
	return rv
}

// Creates and returns a ordered set containing the objects in a given argument list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1543312-orderedsetwithobjects?language=objc
func MutableOrderedSet_OrderedSetWithObjects(firstObj objc.IObject, args ...any) MutableOrderedSet {
	return MutableOrderedSetClass.OrderedSetWithObjects(firstObj, args...)
}

func (mc _MutableOrderedSetClass) OrderedSetWithSet(set ISet) MutableOrderedSet {
	rv := objc.Call[MutableOrderedSet](mc, objc.Sel("orderedSetWithSet:"), objc.Ptr(set))
	return rv
}

// Creates and returns an ordered set with the contents of a set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1543298-orderedsetwithset?language=objc
func MutableOrderedSet_OrderedSetWithSet(set ISet) MutableOrderedSet {
	return MutableOrderedSetClass.OrderedSetWithSet(set)
}

func (m_ MutableOrderedSet) InitWithObjectsCount(objects objc.IObject, cnt uint) MutableOrderedSet {
	rv := objc.Call[MutableOrderedSet](m_, objc.Sel("initWithObjects:count:"), objc.Ptr(objects), cnt)
	return rv
}

// Initializes a newly allocated set with a specified number of objects from a given C array of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsorderedset/1411910-initwithobjects?language=objc
func NewMutableOrderedSetWithObjectsCount(objects objc.IObject, cnt uint) MutableOrderedSet {
	instance := MutableOrderedSetClass.Alloc().InitWithObjectsCount(objects, cnt)
	instance.Autorelease()
	return instance
}

// Appends a given object to the end of the mutable ordered set, if it is not already a member. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1408009-addobject?language=objc
func (m_ MutableOrderedSet) AddObject(object objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("addObject:"), objc.Ptr(object))
}

// Removes a given object from the mutable ordered set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1416776-removeobject?language=objc
func (m_ MutableOrderedSet) RemoveObject(object objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("removeObject:"), objc.Ptr(object))
}

// Removes each object in another given ordered set from the receiving mutable ordered set, if present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1407987-minusorderedset?language=objc
func (m_ MutableOrderedSet) MinusOrderedSet(other IOrderedSet) {
	objc.Call[objc.Void](m_, objc.Sel("minusOrderedSet:"), objc.Ptr(other))
}

// Adds each object in another given ordered set to the receiving mutable ordered set, if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1410973-unionorderedset?language=objc
func (m_ MutableOrderedSet) UnionOrderedSet(other IOrderedSet) {
	objc.Call[objc.Void](m_, objc.Sel("unionOrderedSet:"), objc.Ptr(other))
}

// Replaces the objects in the receiving mutable ordered set at the range with the specified number of objects from a given C array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1415340-replaceobjectsinrange?language=objc
func (m_ MutableOrderedSet) ReplaceObjectsInRangeWithObjectsCount(range_ Range, objects objc.IObject, count uint) {
	objc.Call[objc.Void](m_, objc.Sel("replaceObjectsInRange:withObjects:count:"), range_, objc.Ptr(objects), count)
}

// Replaces the objects at the specified indexes with the new objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1416127-replaceobjectsatindexes?language=objc
func (m_ MutableOrderedSet) ReplaceObjectsAtIndexesWithObjects(indexes IIndexSet, objects []objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("replaceObjectsAtIndexes:withObjects:"), objc.Ptr(indexes), objects)
}

// Exchanges the object at the specified index with the object at the other index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1416821-exchangeobjectatindex?language=objc
func (m_ MutableOrderedSet) ExchangeObjectAtIndexWithObjectAtIndex(idx1 uint, idx2 uint) {
	objc.Call[objc.Void](m_, objc.Sel("exchangeObjectAtIndex:withObjectAtIndex:"), idx1, idx2)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/3152194-applydifference?language=objc
func (m_ MutableOrderedSet) ApplyDifference(difference IOrderedCollectionDifference) {
	objc.Call[objc.Void](m_, objc.Sel("applyDifference:"), objc.Ptr(difference))
}

// Replaces the given object at the specified index of the mutable ordered set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1543323-setobject?language=objc
func (m_ MutableOrderedSet) SetObjectAtIndexedSubscript(obj objc.IObject, idx uint) {
	objc.Call[objc.Void](m_, objc.Sel("setObject:atIndexedSubscript:"), objc.Ptr(obj), idx)
}

// Removes from the receiving ordered set each object that isn’t a member of another set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1415257-intersectset?language=objc
func (m_ MutableOrderedSet) IntersectSet(other ISet) {
	objc.Call[objc.Void](m_, objc.Sel("intersectSet:"), objc.Ptr(other))
}

// Removes from the receiving ordered set each object that isn’t a member of another ordered set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1408541-intersectorderedset?language=objc
func (m_ MutableOrderedSet) IntersectOrderedSet(other IOrderedSet) {
	objc.Call[objc.Void](m_, objc.Sel("intersectOrderedSet:"), objc.Ptr(other))
}

// Removes a the object at the specified index from the mutable ordered set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1415040-removeobjectatindex?language=objc
func (m_ MutableOrderedSet) RemoveObjectAtIndex(idx uint) {
	objc.Call[objc.Void](m_, objc.Sel("removeObjectAtIndex:"), idx)
}

// Sorts the mutable ordered set using the comparison method specified by the comparator block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1414566-sortusingcomparator?language=objc
func (m_ MutableOrderedSet) SortUsingComparator(cmptr Comparator) {
	objc.Call[objc.Void](m_, objc.Sel("sortUsingComparator:"), cmptr)
}

// Removes the objects at the specified indexes from the mutable ordered set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1418161-removeobjectsatindexes?language=objc
func (m_ MutableOrderedSet) RemoveObjectsAtIndexes(indexes IIndexSet) {
	objc.Call[objc.Void](m_, objc.Sel("removeObjectsAtIndexes:"), objc.Ptr(indexes))
}

// Sorts the specified range of the mutable ordered set using the specified options and the comparison method specified by a given comparator block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1407529-sortrange?language=objc
func (m_ MutableOrderedSet) SortRangeOptionsUsingComparator(range_ Range, opts SortOptions, cmptr Comparator) {
	objc.Call[objc.Void](m_, objc.Sel("sortRange:options:usingComparator:"), range_, opts, cmptr)
}

// Removes all the objects from the mutable ordered set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1414262-removeallobjects?language=objc
func (m_ MutableOrderedSet) RemoveAllObjects() {
	objc.Call[objc.Void](m_, objc.Sel("removeAllObjects"))
}

// Inserts the objects in the array at the specified indexes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1410287-insertobjects?language=objc
func (m_ MutableOrderedSet) InsertObjectsAtIndexes(objects []objc.IObject, indexes IIndexSet) {
	objc.Call[objc.Void](m_, objc.Sel("insertObjects:atIndexes:"), objects, objc.Ptr(indexes))
}

// Appends or replaces the object at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1411158-setobject?language=objc
func (m_ MutableOrderedSet) SetObjectAtIndex(obj objc.IObject, idx uint) {
	objc.Call[objc.Void](m_, objc.Sel("setObject:atIndex:"), objc.Ptr(obj), idx)
}

// Sorts the mutable ordered set using the specified options and the comparison method specified by a given comparator block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1411561-sortwithoptions?language=objc
func (m_ MutableOrderedSet) SortWithOptionsUsingComparator(opts SortOptions, cmptr Comparator) {
	objc.Call[objc.Void](m_, objc.Sel("sortWithOptions:usingComparator:"), opts, cmptr)
}

// Sorts the receiving ordered set using a given array of sort descriptors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1410023-sortusingdescriptors?language=objc
func (m_ MutableOrderedSet) SortUsingDescriptors(sortDescriptors []ISortDescriptor) {
	objc.Call[objc.Void](m_, objc.Sel("sortUsingDescriptors:"), sortDescriptors)
}

// Adds each object in another given set to the receiving mutable ordered set, if not present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1413853-unionset?language=objc
func (m_ MutableOrderedSet) UnionSet(other ISet) {
	objc.Call[objc.Void](m_, objc.Sel("unionSet:"), objc.Ptr(other))
}

// Moves the objects at the specified indexes to the new location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1417677-moveobjectsatindexes?language=objc
func (m_ MutableOrderedSet) MoveObjectsAtIndexesToIndex(indexes IIndexSet, idx uint) {
	objc.Call[objc.Void](m_, objc.Sel("moveObjectsAtIndexes:toIndex:"), objc.Ptr(indexes), idx)
}

// Replaces the object at the specified index with the new object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1412115-replaceobjectatindex?language=objc
func (m_ MutableOrderedSet) ReplaceObjectAtIndexWithObject(idx uint, object objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("replaceObjectAtIndex:withObject:"), idx, objc.Ptr(object))
}

// Evaluates a given predicate against the mutable ordered set’s content and leaves only objects that match. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1408348-filterusingpredicate?language=objc
func (m_ MutableOrderedSet) FilterUsingPredicate(p IPredicate) {
	objc.Call[objc.Void](m_, objc.Sel("filterUsingPredicate:"), objc.Ptr(p))
}

// Appends to the end of the mutable ordered set each object contained in a given array that is not already a member. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1417200-addobjectsfromarray?language=objc
func (m_ MutableOrderedSet) AddObjectsFromArray(array []objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("addObjectsFromArray:"), array)
}

// Removes the objects in the array from the mutable ordered set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1411635-removeobjectsinarray?language=objc
func (m_ MutableOrderedSet) RemoveObjectsInArray(array []objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("removeObjectsInArray:"), array)
}

// Removes from the mutable ordered set each of the objects within a given range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1417539-removeobjectsinrange?language=objc
func (m_ MutableOrderedSet) RemoveObjectsInRange(range_ Range) {
	objc.Call[objc.Void](m_, objc.Sel("removeObjectsInRange:"), range_)
}

// Removes each object in another given set from the receiving mutable ordered set, if present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1411229-minusset?language=objc
func (m_ MutableOrderedSet) MinusSet(other ISet) {
	objc.Call[objc.Void](m_, objc.Sel("minusSet:"), objc.Ptr(other))
}

// Appends the given number of objects from a given C array to the end of the mutable ordered set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1413840-addobjects?language=objc
func (m_ MutableOrderedSet) AddObjectsCount(objects objc.IObject, count uint) {
	objc.Call[objc.Void](m_, objc.Sel("addObjects:count:"), objc.Ptr(objects), count)
}

// Inserts the given object at the specified index of the mutable ordered set, if it is not already a member. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableorderedset/1416634-insertobject?language=objc
func (m_ MutableOrderedSet) InsertObjectAtIndex(object objc.IObject, idx uint) {
	objc.Call[objc.Void](m_, objc.Sel("insertObject:atIndex:"), objc.Ptr(object), idx)
}
