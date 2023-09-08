// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// An interface that enables the delegate of a metadata query to provide substitute results or attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquerydelegate?language=objc
type PMetadataQueryDelegate interface {
	// optional
	MetadataQueryReplacementObjectForResultObject(query MetadataQuery, result MetadataItem) objc.Object
	HasMetadataQueryReplacementObjectForResultObject() bool

	// optional
	MetadataQueryReplacementValueForAttributeValue(query MetadataQuery, attrName string, attrValue objc.Object) objc.Object
	HasMetadataQueryReplacementValueForAttributeValue() bool
}

// A delegate implementation builder for the [PMetadataQueryDelegate] protocol.
type MetadataQueryDelegate struct {
	_MetadataQueryReplacementObjectForResultObject  func(query MetadataQuery, result MetadataItem) objc.Object
	_MetadataQueryReplacementValueForAttributeValue func(query MetadataQuery, attrName string, attrValue objc.Object) objc.Object
}

func (di *MetadataQueryDelegate) HasMetadataQueryReplacementObjectForResultObject() bool {
	return di._MetadataQueryReplacementObjectForResultObject != nil
}

// Returns a different object for a given query result object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquerydelegate/1407317-metadataquery?language=objc
func (di *MetadataQueryDelegate) SetMetadataQueryReplacementObjectForResultObject(f func(query MetadataQuery, result MetadataItem) objc.Object) {
	di._MetadataQueryReplacementObjectForResultObject = f
}

// Returns a different object for a given query result object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquerydelegate/1407317-metadataquery?language=objc
func (di *MetadataQueryDelegate) MetadataQueryReplacementObjectForResultObject(query MetadataQuery, result MetadataItem) objc.Object {
	return di._MetadataQueryReplacementObjectForResultObject(query, result)
}
func (di *MetadataQueryDelegate) HasMetadataQueryReplacementValueForAttributeValue() bool {
	return di._MetadataQueryReplacementValueForAttributeValue != nil
}

// Returns a different value for a given attribute and value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquerydelegate/1414215-metadataquery?language=objc
func (di *MetadataQueryDelegate) SetMetadataQueryReplacementValueForAttributeValue(f func(query MetadataQuery, attrName string, attrValue objc.Object) objc.Object) {
	di._MetadataQueryReplacementValueForAttributeValue = f
}

// Returns a different value for a given attribute and value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquerydelegate/1414215-metadataquery?language=objc
func (di *MetadataQueryDelegate) MetadataQueryReplacementValueForAttributeValue(query MetadataQuery, attrName string, attrValue objc.Object) objc.Object {
	return di._MetadataQueryReplacementValueForAttributeValue(query, attrName, attrValue)
}

// ensure impl type implements protocol interface
var _ PMetadataQueryDelegate = (*MetadataQueryDelegateObject)(nil)

// A concrete type for the [PMetadataQueryDelegate] protocol.
type MetadataQueryDelegateObject struct {
	objc.Object
}

func (m_ MetadataQueryDelegateObject) HasMetadataQueryReplacementObjectForResultObject() bool {
	return m_.RespondsToSelector(objc.Sel("metadataQuery:replacementObjectForResultObject:"))
}

// Returns a different object for a given query result object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquerydelegate/1407317-metadataquery?language=objc
func (m_ MetadataQueryDelegateObject) MetadataQueryReplacementObjectForResultObject(query MetadataQuery, result MetadataItem) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("metadataQuery:replacementObjectForResultObject:"), objc.Ptr(query), objc.Ptr(result))
	return rv
}

func (m_ MetadataQueryDelegateObject) HasMetadataQueryReplacementValueForAttributeValue() bool {
	return m_.RespondsToSelector(objc.Sel("metadataQuery:replacementValueForAttribute:value:"))
}

// Returns a different value for a given attribute and value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataquerydelegate/1414215-metadataquery?language=objc
func (m_ MetadataQueryDelegateObject) MetadataQueryReplacementValueForAttributeValue(query MetadataQuery, attrName string, attrValue objc.Object) objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("metadataQuery:replacementValueForAttribute:value:"), objc.Ptr(query), attrName, attrValue)
	return rv
}
