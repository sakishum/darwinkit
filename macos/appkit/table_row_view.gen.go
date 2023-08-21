// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TableRowView] class.
var TableRowViewClass = _TableRowViewClass{objc.GetClass("NSTableRowView")}

type _TableRowViewClass struct {
	objc.Class
}

// An interface definition for the [TableRowView] class.
type ITableRowView interface {
	IView
	ViewAtColumn(column int) objc.Object
	DrawBackgroundInRect(dirtyRect foundation.Rect)
	DrawSeparatorInRect(dirtyRect foundation.Rect)
	DrawSelectionInRect(dirtyRect foundation.Rect)
	DrawDraggingDestinationFeedbackInRect(dirtyRect foundation.Rect)
	IsPreviousRowSelected() bool
	SetPreviousRowSelected(value bool)
	IsTargetForDropOperation() bool
	SetTargetForDropOperation(value bool)
	IsEmphasized() bool
	SetEmphasized(value bool)
	IsSelected() bool
	SetSelected(value bool)
	IsFloating() bool
	SetFloating(value bool)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	IsGroupRowStyle() bool
	SetGroupRowStyle(value bool)
	InteriorBackgroundStyle() BackgroundStyle
	SelectionHighlightStyle() TableViewSelectionHighlightStyle
	SetSelectionHighlightStyle(value TableViewSelectionHighlightStyle)
	IndentationForDropOperation() float64
	SetIndentationForDropOperation(value float64)
	NumberOfColumns() int
	IsNextRowSelected() bool
	SetNextRowSelected(value bool)
	DraggingDestinationFeedbackStyle() TableViewDraggingDestinationFeedbackStyle
	SetDraggingDestinationFeedbackStyle(value TableViewDraggingDestinationFeedbackStyle)
}

// The view shown for a row in a table view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview?language=objc
type TableRowView struct {
	View
}

func TableRowViewFrom(ptr unsafe.Pointer) TableRowView {
	return TableRowView{
		View: ViewFrom(ptr),
	}
}

func (tc _TableRowViewClass) Alloc() TableRowView {
	rv := objc.Call[TableRowView](tc, objc.Sel("alloc"))
	return rv
}

func TableRowView_Alloc() TableRowView {
	return TableRowViewClass.Alloc()
}

func (tc _TableRowViewClass) New() TableRowView {
	rv := objc.Call[TableRowView](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTableRowView() TableRowView {
	return TableRowViewClass.New()
}

func (t_ TableRowView) Init() TableRowView {
	rv := objc.Call[TableRowView](t_, objc.Sel("init"))
	return rv
}

func (t_ TableRowView) InitWithFrame(frameRect foundation.Rect) TableRowView {
	rv := objc.Call[TableRowView](t_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func NewTableRowViewWithFrame(frameRect foundation.Rect) TableRowView {
	instance := TableRowViewClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

// Provides access to the given view at a particular column. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1534440-viewatcolumn?language=objc
func (t_ TableRowView) ViewAtColumn(column int) objc.Object {
	rv := objc.Call[objc.Object](t_, objc.Sel("viewAtColumn:"), column)
	return rv
}

// Draws the background of the row in the rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1531936-drawbackgroundinrect?language=objc
func (t_ TableRowView) DrawBackgroundInRect(dirtyRect foundation.Rect) {
	objc.Call[objc.Void](t_, objc.Sel("drawBackgroundInRect:"), dirtyRect)
}

// Draws the horizontal separator between table rows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1525167-drawseparatorinrect?language=objc
func (t_ TableRowView) DrawSeparatorInRect(dirtyRect foundation.Rect) {
	objc.Call[objc.Void](t_, objc.Sel("drawSeparatorInRect:"), dirtyRect)
}

// Draws the selected row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1526425-drawselectioninrect?language=objc
func (t_ TableRowView) DrawSelectionInRect(dirtyRect foundation.Rect) {
	objc.Call[objc.Void](t_, objc.Sel("drawSelectionInRect:"), dirtyRect)
}

// Draws the row’s dragging destination feedback when the entire row is a drop target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1528434-drawdraggingdestinationfeedbacki?language=objc
func (t_ TableRowView) DrawDraggingDestinationFeedbackInRect(dirtyRect foundation.Rect) {
	objc.Call[objc.Void](t_, objc.Sel("drawDraggingDestinationFeedbackInRect:"), dirtyRect)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1535313-previousrowselected?language=objc
func (t_ TableRowView) IsPreviousRowSelected() bool {
	rv := objc.Call[bool](t_, objc.Sel("isPreviousRowSelected"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1535313-previousrowselected?language=objc
func (t_ TableRowView) SetPreviousRowSelected(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setPreviousRowSelected:"), value)
}

// Specifies whether this row will draw a drop indicator based on the current dragging feedback style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1533914-targetfordropoperation?language=objc
func (t_ TableRowView) IsTargetForDropOperation() bool {
	rv := objc.Call[bool](t_, objc.Sel("isTargetForDropOperation"))
	return rv
}

// Specifies whether this row will draw a drop indicator based on the current dragging feedback style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1533914-targetfordropoperation?language=objc
func (t_ TableRowView) SetTargetForDropOperation(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setTargetForDropOperation:"), value)
}

// Determines whether the row will draw with the alternate or secondary color (unless overridden). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1526258-emphasized?language=objc
func (t_ TableRowView) IsEmphasized() bool {
	rv := objc.Call[bool](t_, objc.Sel("isEmphasized"))
	return rv
}

// Determines whether the row will draw with the alternate or secondary color (unless overridden). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1526258-emphasized?language=objc
func (t_ TableRowView) SetEmphasized(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setEmphasized:"), value)
}

// Determines whether the row is selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1529508-selected?language=objc
func (t_ TableRowView) IsSelected() bool {
	rv := objc.Call[bool](t_, objc.Sel("isSelected"))
	return rv
}

// Determines whether the row is selected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1529508-selected?language=objc
func (t_ TableRowView) SetSelected(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setSelected:"), value)
}

// Specifies whether the row is drawn using the floating style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1534291-floating?language=objc
func (t_ TableRowView) IsFloating() bool {
	rv := objc.Call[bool](t_, objc.Sel("isFloating"))
	return rv
}

// Specifies whether the row is drawn using the floating style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1534291-floating?language=objc
func (t_ TableRowView) SetFloating(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setFloating:"), value)
}

// The background color of the row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1534057-backgroundcolor?language=objc
func (t_ TableRowView) BackgroundColor() Color {
	rv := objc.Call[Color](t_, objc.Sel("backgroundColor"))
	return rv
}

// The background color of the row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1534057-backgroundcolor?language=objc
func (t_ TableRowView) SetBackgroundColor(value IColor) {
	objc.Call[objc.Void](t_, objc.Sel("setBackgroundColor:"), objc.Ptr(value))
}

// Specifies whether this row view is a group row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1530499-grouprowstyle?language=objc
func (t_ TableRowView) IsGroupRowStyle() bool {
	rv := objc.Call[bool](t_, objc.Sel("isGroupRowStyle"))
	return rv
}

// Specifies whether this row view is a group row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1530499-grouprowstyle?language=objc
func (t_ TableRowView) SetGroupRowStyle(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setGroupRowStyle:"), value)
}

// Specifies how the subviews should draw. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1535905-interiorbackgroundstyle?language=objc
func (t_ TableRowView) InteriorBackgroundStyle() BackgroundStyle {
	rv := objc.Call[BackgroundStyle](t_, objc.Sel("interiorBackgroundStyle"))
	return rv
}

// Specifies the selection highlight style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1531083-selectionhighlightstyle?language=objc
func (t_ TableRowView) SelectionHighlightStyle() TableViewSelectionHighlightStyle {
	rv := objc.Call[TableViewSelectionHighlightStyle](t_, objc.Sel("selectionHighlightStyle"))
	return rv
}

// Specifies the selection highlight style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1531083-selectionhighlightstyle?language=objc
func (t_ TableRowView) SetSelectionHighlightStyle(value TableViewSelectionHighlightStyle) {
	objc.Call[objc.Void](t_, objc.Sel("setSelectionHighlightStyle:"), value)
}

// Defines the amount the drag target for a row should be indented. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1535836-indentationfordropoperation?language=objc
func (t_ TableRowView) IndentationForDropOperation() float64 {
	rv := objc.Call[float64](t_, objc.Sel("indentationForDropOperation"))
	return rv
}

// Defines the amount the drag target for a row should be indented. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1535836-indentationfordropoperation?language=objc
func (t_ TableRowView) SetIndentationForDropOperation(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setIndentationForDropOperation:"), value)
}

// Returns the number of columns represented by views in the table row view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1525610-numberofcolumns?language=objc
func (t_ TableRowView) NumberOfColumns() int {
	rv := objc.Call[int](t_, objc.Sel("numberOfColumns"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1529083-nextrowselected?language=objc
func (t_ TableRowView) IsNextRowSelected() bool {
	rv := objc.Call[bool](t_, objc.Sel("isNextRowSelected"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1529083-nextrowselected?language=objc
func (t_ TableRowView) SetNextRowSelected(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setNextRowSelected:"), value)
}

// Specifies the dragging destination feedback style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1525012-draggingdestinationfeedbackstyle?language=objc
func (t_ TableRowView) DraggingDestinationFeedbackStyle() TableViewDraggingDestinationFeedbackStyle {
	rv := objc.Call[TableViewDraggingDestinationFeedbackStyle](t_, objc.Sel("draggingDestinationFeedbackStyle"))
	return rv
}

// Specifies the dragging destination feedback style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstablerowview/1525012-draggingdestinationfeedbackstyle?language=objc
func (t_ TableRowView) SetDraggingDestinationFeedbackStyle(value TableViewDraggingDestinationFeedbackStyle) {
	objc.Call[objc.Void](t_, objc.Sel("setDraggingDestinationFeedbackStyle:"), value)
}