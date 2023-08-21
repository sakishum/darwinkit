// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ClipView] class.
var ClipViewClass = _ClipViewClass{objc.GetClass("NSClipView")}

type _ClipViewClass struct {
	objc.Class
}

// An interface definition for the [ClipView] class.
type IClipView interface {
	IView
	ViewBoundsChanged(notification foundation.INotification)
	ScrollToPoint(newOrigin foundation.Point)
	ConstrainBoundsRect(proposedBounds foundation.Rect) foundation.Rect
	ViewFrameChanged(notification foundation.INotification)
	DocumentCursor() Cursor
	SetDocumentCursor(value ICursor)
	ContentInsets() foundation.EdgeInsets
	SetContentInsets(value foundation.EdgeInsets)
	BackgroundColor() Color
	SetBackgroundColor(value IColor)
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	DocumentVisibleRect() foundation.Rect
	DocumentRect() foundation.Rect
	DocumentView() View
	SetDocumentView(value IView)
	AutomaticallyAdjustsContentInsets() bool
	SetAutomaticallyAdjustsContentInsets(value bool)
}

// An object that clips a document view to a scroll view's frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclipview?language=objc
type ClipView struct {
	View
}

func ClipViewFrom(ptr unsafe.Pointer) ClipView {
	return ClipView{
		View: ViewFrom(ptr),
	}
}

func (cc _ClipViewClass) Alloc() ClipView {
	rv := objc.Call[ClipView](cc, objc.Sel("alloc"))
	return rv
}

func ClipView_Alloc() ClipView {
	return ClipViewClass.Alloc()
}

func (cc _ClipViewClass) New() ClipView {
	rv := objc.Call[ClipView](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewClipView() ClipView {
	return ClipViewClass.New()
}

func (c_ ClipView) Init() ClipView {
	rv := objc.Call[ClipView](c_, objc.Sel("init"))
	return rv
}

func (c_ ClipView) InitWithFrame(frameRect foundation.Rect) ClipView {
	rv := objc.Call[ClipView](c_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func NewClipViewWithFrame(frameRect foundation.Rect) ClipView {
	instance := ClipViewClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

// Handles an NSViewBoundsDidChangeNotification, passed in the aNotification argument, by updating a containing NSScrollView based on the new bounds. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclipview/1531354-viewboundschanged?language=objc
func (c_ ClipView) ViewBoundsChanged(notification foundation.INotification) {
	objc.Call[objc.Void](c_, objc.Sel("viewBoundsChanged:"), objc.Ptr(notification))
}

// Changes the origin of the clip view’s bounds rectangle to newOrigin. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclipview/1528826-scrolltopoint?language=objc
func (c_ ClipView) ScrollToPoint(newOrigin foundation.Point) {
	objc.Call[objc.Void](c_, objc.Sel("scrollToPoint:"), newOrigin)
}

// Constrains the bounds of the clip view while the user is magnifying and scrolling. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclipview/1534160-constrainboundsrect?language=objc
func (c_ ClipView) ConstrainBoundsRect(proposedBounds foundation.Rect) foundation.Rect {
	rv := objc.Call[foundation.Rect](c_, objc.Sel("constrainBoundsRect:"), proposedBounds)
	return rv
}

// Handles an NSViewFrameDidChangeNotification, passed in the aNotification argument, by updating a containing NSScrollView based on the new frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclipview/1526364-viewframechanged?language=objc
func (c_ ClipView) ViewFrameChanged(notification foundation.INotification) {
	objc.Call[objc.Void](c_, objc.Sel("viewFrameChanged:"), objc.Ptr(notification))
}

// The cursor object used when the pointer lies over the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclipview/1535377-documentcursor?language=objc
func (c_ ClipView) DocumentCursor() Cursor {
	rv := objc.Call[Cursor](c_, objc.Sel("documentCursor"))
	return rv
}

// The cursor object used when the pointer lies over the view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclipview/1535377-documentcursor?language=objc
func (c_ ClipView) SetDocumentCursor(value ICursor) {
	objc.Call[objc.Void](c_, objc.Sel("setDocumentCursor:"), objc.Ptr(value))
}

// The distance that the content view is inset from the enclosing scroll view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclipview/1524329-contentinsets?language=objc
func (c_ ClipView) ContentInsets() foundation.EdgeInsets {
	rv := objc.Call[foundation.EdgeInsets](c_, objc.Sel("contentInsets"))
	return rv
}

// The distance that the content view is inset from the enclosing scroll view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclipview/1524329-contentinsets?language=objc
func (c_ ClipView) SetContentInsets(value foundation.EdgeInsets) {
	objc.Call[objc.Void](c_, objc.Sel("setContentInsets:"), value)
}

// The color of the clip view’s background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclipview/1525469-backgroundcolor?language=objc
func (c_ ClipView) BackgroundColor() Color {
	rv := objc.Call[Color](c_, objc.Sel("backgroundColor"))
	return rv
}

// The color of the clip view’s background. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclipview/1525469-backgroundcolor?language=objc
func (c_ ClipView) SetBackgroundColor(value IColor) {
	objc.Call[objc.Void](c_, objc.Sel("setBackgroundColor:"), objc.Ptr(value))
}

// A Boolean value that indicates if the clip view draws its background color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclipview/1534684-drawsbackground?language=objc
func (c_ ClipView) DrawsBackground() bool {
	rv := objc.Call[bool](c_, objc.Sel("drawsBackground"))
	return rv
}

// A Boolean value that indicates if the clip view draws its background color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclipview/1534684-drawsbackground?language=objc
func (c_ ClipView) SetDrawsBackground(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setDrawsBackground:"), value)
}

// The exposed rectangle of the clip view’s document view, in the document view’s own coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclipview/1527958-documentvisiblerect?language=objc
func (c_ ClipView) DocumentVisibleRect() foundation.Rect {
	rv := objc.Call[foundation.Rect](c_, objc.Sel("documentVisibleRect"))
	return rv
}

// The rectangle defining the document view’s frame, adjusted to the size of the clip view if the document view is smaller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclipview/1533338-documentrect?language=objc
func (c_ ClipView) DocumentRect() foundation.Rect {
	rv := objc.Call[foundation.Rect](c_, objc.Sel("documentRect"))
	return rv
}

// The clip view’s document view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclipview/1524587-documentview?language=objc
func (c_ ClipView) DocumentView() View {
	rv := objc.Call[View](c_, objc.Sel("documentView"))
	return rv
}

// The clip view’s document view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclipview/1524587-documentview?language=objc
func (c_ ClipView) SetDocumentView(value IView) {
	objc.Call[objc.Void](c_, objc.Sel("setDocumentView:"), objc.Ptr(value))
}

// A Boolean value that indicates if the clip view automatically accounts for other scroll view subviews. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclipview/1527540-automaticallyadjustscontentinset?language=objc
func (c_ ClipView) AutomaticallyAdjustsContentInsets() bool {
	rv := objc.Call[bool](c_, objc.Sel("automaticallyAdjustsContentInsets"))
	return rv
}

// A Boolean value that indicates if the clip view automatically accounts for other scroll view subviews. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsclipview/1527540-automaticallyadjustscontentinset?language=objc
func (c_ ClipView) SetAutomaticallyAdjustsContentInsets(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setAutomaticallyAdjustsContentInsets:"), value)
}