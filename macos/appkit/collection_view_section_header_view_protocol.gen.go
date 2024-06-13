// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that defines a button to control the collapse of a collection view’s section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewsectionheaderview?language=objc
type PCollectionViewSectionHeaderView interface {
	// optional
	SetSectionCollapseButton(value Button)
	HasSetSectionCollapseButton() bool

	// optional
	SectionCollapseButton() Button
	HasSectionCollapseButton() bool
}

// ensure impl type implements protocol interface
var _ PCollectionViewSectionHeaderView = (*CollectionViewSectionHeaderViewObject)(nil)

// A concrete type for the [PCollectionViewSectionHeaderView] protocol.
type CollectionViewSectionHeaderViewObject struct {
	objc.Object
}

func (c_ CollectionViewSectionHeaderViewObject) HasSetSectionCollapseButton() bool {
	return c_.RespondsToSelector(objc.Sel("setSectionCollapseButton:"))
}

// A control that lets users collapse and open a collection view section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewsectionheaderview/1644266-sectioncollapsebutton?language=objc
func (c_ CollectionViewSectionHeaderViewObject) SetSectionCollapseButton(value Button) {
	objc.Call[objc.Void](c_, objc.Sel("setSectionCollapseButton:"), objc.Ptr(value))
}

func (c_ CollectionViewSectionHeaderViewObject) HasSectionCollapseButton() bool {
	return c_.RespondsToSelector(objc.Sel("sectionCollapseButton"))
}

// A control that lets users collapse and open a collection view section. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewsectionheaderview/1644266-sectioncollapsebutton?language=objc
func (c_ CollectionViewSectionHeaderViewObject) SectionCollapseButton() Button {
	rv := objc.Call[Button](c_, objc.Sel("sectionCollapseButton"))
	return rv
}