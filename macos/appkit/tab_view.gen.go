// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TabView] class.
var TabViewClass = _TabViewClass{objc.GetClass("NSTabView")}

type _TabViewClass struct {
	objc.Class
}

// An interface definition for the [TabView] class.
type ITabView interface {
	IView
	TabViewItemAtIndex(index int) TabViewItem
	TabViewItemAtPoint(point foundation.Point) TabViewItem
	RemoveTabViewItem(tabViewItem ITabViewItem)
	TakeSelectedTabViewItemFromSender(sender objc.IObject)
	SelectFirstTabViewItem(sender objc.IObject)
	SelectPreviousTabViewItem(sender objc.IObject)
	IndexOfTabViewItemWithIdentifier(identifier objc.IObject) int
	SelectNextTabViewItem(sender objc.IObject)
	SelectTabViewItemAtIndex(index int)
	SelectLastTabViewItem(sender objc.IObject)
	SelectTabViewItem(tabViewItem ITabViewItem)
	InsertTabViewItemAtIndex(tabViewItem ITabViewItem, index int)
	IndexOfTabViewItem(tabViewItem ITabViewItem) int
	AddTabViewItem(tabViewItem ITabViewItem)
	SelectTabViewItemWithIdentifier(identifier objc.IObject)
	MinimumSize() foundation.Size
	ControlSize() ControlSize
	SetControlSize(value ControlSize)
	TabViewItems() []TabViewItem
	SetTabViewItems(value []ITabViewItem)
	ContentRect() foundation.Rect
	NumberOfTabViewItems() int
	AllowsTruncatedLabels() bool
	SetAllowsTruncatedLabels(value bool)
	SelectedTabViewItem() TabViewItem
	Delegate() TabViewDelegateWrapper
	SetDelegate(value PTabViewDelegate)
	SetDelegateObject(valueObject objc.IObject)
	TabPosition() TabPosition
	SetTabPosition(value TabPosition)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	TabViewBorderType() TabViewBorderType
	SetTabViewBorderType(value TabViewBorderType)
	Font() Font
	SetFont(value IFont)
	TabViewType() TabViewType
	SetTabViewType(value TabViewType)
}

// A multipage interface that displays one page at a time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview?language=objc
type TabView struct {
	View
}

func TabViewFrom(ptr unsafe.Pointer) TabView {
	return TabView{
		View: ViewFrom(ptr),
	}
}

func (tc _TabViewClass) Alloc() TabView {
	rv := objc.Call[TabView](tc, objc.Sel("alloc"))
	return rv
}

func TabView_Alloc() TabView {
	return TabViewClass.Alloc()
}

func (tc _TabViewClass) New() TabView {
	rv := objc.Call[TabView](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTabView() TabView {
	return TabViewClass.New()
}

func (t_ TabView) Init() TabView {
	rv := objc.Call[TabView](t_, objc.Sel("init"))
	return rv
}

func (t_ TabView) InitWithFrame(frameRect foundation.Rect) TabView {
	rv := objc.Call[TabView](t_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func NewTabViewWithFrame(frameRect foundation.Rect) TabView {
	instance := TabViewClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

// Returns the tab view item at index in the tab view’s array of items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391637-tabviewitematindex?language=objc
func (t_ TabView) TabViewItemAtIndex(index int) TabViewItem {
	rv := objc.Call[TabViewItem](t_, objc.Sel("tabViewItemAtIndex:"), index)
	return rv
}

// Returns the tab view item at the specified point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391606-tabviewitematpoint?language=objc
func (t_ TabView) TabViewItemAtPoint(point foundation.Point) TabViewItem {
	rv := objc.Call[TabViewItem](t_, objc.Sel("tabViewItemAtPoint:"), point)
	return rv
}

// Removes the specified item from the tab view’s array of tab view items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391641-removetabviewitem?language=objc
func (t_ TabView) RemoveTabViewItem(tabViewItem ITabViewItem) {
	objc.Call[objc.Void](t_, objc.Sel("removeTabViewItem:"), objc.Ptr(tabViewItem))
}

// Sets the selected tab view item to the selected item obtained from the sender. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391653-takeselectedtabviewitemfromsende?language=objc
func (t_ TabView) TakeSelectedTabViewItemFromSender(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("takeSelectedTabViewItemFromSender:"), sender)
}

// This action method selects the first tab view item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391627-selectfirsttabviewitem?language=objc
func (t_ TabView) SelectFirstTabViewItem(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("selectFirstTabViewItem:"), sender)
}

// This action method selects the previous tab view item in the sequence. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391594-selectprevioustabviewitem?language=objc
func (t_ TabView) SelectPreviousTabViewItem(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("selectPreviousTabViewItem:"), sender)
}

// Returns the index of the item that matches the specified identifier or NSNotFound if the item is not found. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391586-indexoftabviewitemwithidentifier?language=objc
func (t_ TabView) IndexOfTabViewItemWithIdentifier(identifier objc.IObject) int {
	rv := objc.Call[int](t_, objc.Sel("indexOfTabViewItemWithIdentifier:"), identifier)
	return rv
}

// This action method selects the next tab view item in the sequence. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391639-selectnexttabviewitem?language=objc
func (t_ TabView) SelectNextTabViewItem(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("selectNextTabViewItem:"), sender)
}

// Selects the tab view item specified by index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391590-selecttabviewitematindex?language=objc
func (t_ TabView) SelectTabViewItemAtIndex(index int) {
	objc.Call[objc.Void](t_, objc.Sel("selectTabViewItemAtIndex:"), index)
}

// This action method selects the last tab view item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391635-selectlasttabviewitem?language=objc
func (t_ TabView) SelectLastTabViewItem(sender objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("selectLastTabViewItem:"), sender)
}

// Selects the specified tab view item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391649-selecttabviewitem?language=objc
func (t_ TabView) SelectTabViewItem(tabViewItem ITabViewItem) {
	objc.Call[objc.Void](t_, objc.Sel("selectTabViewItem:"), objc.Ptr(tabViewItem))
}

// Inserts the specified item into the tab view’s array of tab view items at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391584-inserttabviewitem?language=objc
func (t_ TabView) InsertTabViewItemAtIndex(tabViewItem ITabViewItem, index int) {
	objc.Call[objc.Void](t_, objc.Sel("insertTabViewItem:atIndex:"), objc.Ptr(tabViewItem), index)
}

// Returns the index of the specified item in the tab view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391604-indexoftabviewitem?language=objc
func (t_ TabView) IndexOfTabViewItem(tabViewItem ITabViewItem) int {
	rv := objc.Call[int](t_, objc.Sel("indexOfTabViewItem:"), objc.Ptr(tabViewItem))
	return rv
}

// Adds the specified tab item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391607-addtabviewitem?language=objc
func (t_ TabView) AddTabViewItem(tabViewItem ITabViewItem) {
	objc.Call[objc.Void](t_, objc.Sel("addTabViewItem:"), objc.Ptr(tabViewItem))
}

// Selects the tab view item specified by identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391629-selecttabviewitemwithidentifier?language=objc
func (t_ TabView) SelectTabViewItemWithIdentifier(identifier objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("selectTabViewItemWithIdentifier:"), identifier)
}

// The minimum size necessary for the tab view to display tabs in a useful way. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391598-minimumsize?language=objc
func (t_ TabView) MinimumSize() foundation.Size {
	rv := objc.Call[foundation.Size](t_, objc.Sel("minimumSize"))
	return rv
}

// The size of the tab view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391633-controlsize?language=objc
func (t_ TabView) ControlSize() ControlSize {
	rv := objc.Call[ControlSize](t_, objc.Sel("controlSize"))
	return rv
}

// The size of the tab view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391633-controlsize?language=objc
func (t_ TabView) SetControlSize(value ControlSize) {
	objc.Call[objc.Void](t_, objc.Sel("setControlSize:"), value)
}

// The tab view’s array of tab view items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391613-tabviewitems?language=objc
func (t_ TabView) TabViewItems() []TabViewItem {
	rv := objc.Call[[]TabViewItem](t_, objc.Sel("tabViewItems"))
	return rv
}

// The tab view’s array of tab view items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391613-tabviewitems?language=objc
func (t_ TabView) SetTabViewItems(value []ITabViewItem) {
	objc.Call[objc.Void](t_, objc.Sel("setTabViewItems:"), value)
}

// The rectangle describing the content area of the tab view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391659-contentrect?language=objc
func (t_ TabView) ContentRect() foundation.Rect {
	rv := objc.Call[foundation.Rect](t_, objc.Sel("contentRect"))
	return rv
}

// The number of items in the tab view’s array of tab view items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391596-numberoftabviewitems?language=objc
func (t_ TabView) NumberOfTabViewItems() int {
	rv := objc.Call[int](t_, objc.Sel("numberOfTabViewItems"))
	return rv
}

// A Boolean value that indicates if the tab view allows truncating for labels that don’t fit on a tab. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391645-allowstruncatedlabels?language=objc
func (t_ TabView) AllowsTruncatedLabels() bool {
	rv := objc.Call[bool](t_, objc.Sel("allowsTruncatedLabels"))
	return rv
}

// A Boolean value that indicates if the tab view allows truncating for labels that don’t fit on a tab. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391645-allowstruncatedlabels?language=objc
func (t_ TabView) SetAllowsTruncatedLabels(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAllowsTruncatedLabels:"), value)
}

// The tab view item for the currently selected tab. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391625-selectedtabviewitem?language=objc
func (t_ TabView) SelectedTabViewItem() TabViewItem {
	rv := objc.Call[TabViewItem](t_, objc.Sel("selectedTabViewItem"))
	return rv
}

// The tab view’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391615-delegate?language=objc
func (t_ TabView) Delegate() TabViewDelegateWrapper {
	rv := objc.Call[TabViewDelegateWrapper](t_, objc.Sel("delegate"))
	return rv
}

// The tab view’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391615-delegate?language=objc
func (t_ TabView) SetDelegate(value PTabViewDelegate) {
	po0 := objc.WrapAsProtocol("NSTabViewDelegate", value)
	objc.SetAssociatedObject(t_, objc.AssociationKey("setDelegate"), po0, objc.ASSOCIATION_RETAIN)
	objc.Call[objc.Void](t_, objc.Sel("setDelegate:"), po0)
}

// The tab view’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391615-delegate?language=objc
func (t_ TabView) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/2097105-tabposition?language=objc
func (t_ TabView) TabPosition() TabPosition {
	rv := objc.Call[TabPosition](t_, objc.Sel("tabPosition"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/2097105-tabposition?language=objc
func (t_ TabView) SetTabPosition(value TabPosition) {
	objc.Call[objc.Void](t_, objc.Sel("setTabPosition:"), value)
}

// A Boolean value that indicates if the tab view draws a background color when its type is NSNoTabsNoBorder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391588-drawsbackground?language=objc
func (t_ TabView) DrawsBackground() bool {
	rv := objc.Call[bool](t_, objc.Sel("drawsBackground"))
	return rv
}

// A Boolean value that indicates if the tab view draws a background color when its type is NSNoTabsNoBorder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391588-drawsbackground?language=objc
func (t_ TabView) SetDrawsBackground(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setDrawsBackground:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/2097112-tabviewbordertype?language=objc
func (t_ TabView) TabViewBorderType() TabViewBorderType {
	rv := objc.Call[TabViewBorderType](t_, objc.Sel("tabViewBorderType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/2097112-tabviewbordertype?language=objc
func (t_ TabView) SetTabViewBorderType(value TabViewBorderType) {
	objc.Call[objc.Void](t_, objc.Sel("setTabViewBorderType:"), value)
}

// The font used for the tab view’s label text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391617-font?language=objc
func (t_ TabView) Font() Font {
	rv := objc.Call[Font](t_, objc.Sel("font"))
	return rv
}

// The font used for the tab view’s label text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391617-font?language=objc
func (t_ TabView) SetFont(value IFont) {
	objc.Call[objc.Void](t_, objc.Sel("setFont:"), objc.Ptr(value))
}

// The tab type to display the tabs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391602-tabviewtype?language=objc
func (t_ TabView) TabViewType() TabViewType {
	rv := objc.Call[TabViewType](t_, objc.Sel("tabViewType"))
	return rv
}

// The tab type to display the tabs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstabview/1391602-tabviewtype?language=objc
func (t_ TabView) SetTabViewType(value TabViewType) {
	objc.Call[objc.Void](t_, objc.Sel("setTabViewType:"), value)
}