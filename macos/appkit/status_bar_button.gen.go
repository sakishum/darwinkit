// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [StatusBarButton] class.
var StatusBarButtonClass = _StatusBarButtonClass{objc.GetClass("NSStatusBarButton")}

type _StatusBarButtonClass struct {
	objc.Class
}

// An interface definition for the [StatusBarButton] class.
type IStatusBarButton interface {
	IButton
	AppearsDisabled() bool
	SetAppearsDisabled(value bool)
}

// The appearance and behavior of an item in the systemwide menu bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusbarbutton?language=objc
type StatusBarButton struct {
	Button
}

func StatusBarButtonFrom(ptr unsafe.Pointer) StatusBarButton {
	return StatusBarButton{
		Button: ButtonFrom(ptr),
	}
}

func (sc _StatusBarButtonClass) Alloc() StatusBarButton {
	rv := objc.Call[StatusBarButton](sc, objc.Sel("alloc"))
	return rv
}

func StatusBarButton_Alloc() StatusBarButton {
	return StatusBarButtonClass.Alloc()
}

func (sc _StatusBarButtonClass) New() StatusBarButton {
	rv := objc.Call[StatusBarButton](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewStatusBarButton() StatusBarButton {
	return StatusBarButtonClass.New()
}

func (s_ StatusBarButton) Init() StatusBarButton {
	rv := objc.Call[StatusBarButton](s_, objc.Sel("init"))
	return rv
}

func (sc _StatusBarButtonClass) CheckboxWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) StatusBarButton {
	rv := objc.Call[StatusBarButton](sc, objc.Sel("checkboxWithTitle:target:action:"), title, target, action)
	return rv
}

// Creates a standard checkbox with the title you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1644525-checkboxwithtitle?language=objc
func StatusBarButton_CheckboxWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) StatusBarButton {
	return StatusBarButtonClass.CheckboxWithTitleTargetAction(title, target, action)
}

func (sc _StatusBarButtonClass) ButtonWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) StatusBarButton {
	rv := objc.Call[StatusBarButton](sc, objc.Sel("buttonWithTitle:target:action:"), title, target, action)
	return rv
}

// Creates a standard push button with the title you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1644256-buttonwithtitle?language=objc
func StatusBarButton_ButtonWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) StatusBarButton {
	return StatusBarButtonClass.ButtonWithTitleTargetAction(title, target, action)
}

func (sc _StatusBarButtonClass) ButtonWithImageTargetAction(image IImage, target objc.IObject, action objc.Selector) StatusBarButton {
	rv := objc.Call[StatusBarButton](sc, objc.Sel("buttonWithImage:target:action:"), objc.Ptr(image), target, action)
	return rv
}

// Creates a standard push button with the image you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1644659-buttonwithimage?language=objc
func StatusBarButton_ButtonWithImageTargetAction(image IImage, target objc.IObject, action objc.Selector) StatusBarButton {
	return StatusBarButtonClass.ButtonWithImageTargetAction(image, target, action)
}

func (sc _StatusBarButtonClass) RadioButtonWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) StatusBarButton {
	rv := objc.Call[StatusBarButton](sc, objc.Sel("radioButtonWithTitle:target:action:"), title, target, action)
	return rv
}

// Creates a standard radio button with the title you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/1644340-radiobuttonwithtitle?language=objc
func StatusBarButton_RadioButtonWithTitleTargetAction(title string, target objc.IObject, action objc.Selector) StatusBarButton {
	return StatusBarButtonClass.RadioButtonWithTitleTargetAction(title, target, action)
}

func (s_ StatusBarButton) InitWithFrame(frameRect foundation.Rect) StatusBarButton {
	rv := objc.Call[StatusBarButton](s_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes a control with the specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428900-initwithframe?language=objc
func NewStatusBarButtonWithFrame(frameRect foundation.Rect) StatusBarButton {
	instance := StatusBarButtonClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusbarbutton/1409292-appearsdisabled?language=objc
func (s_ StatusBarButton) AppearsDisabled() bool {
	rv := objc.Call[bool](s_, objc.Sel("appearsDisabled"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsstatusbarbutton/1409292-appearsdisabled?language=objc
func (s_ StatusBarButton) SetAppearsDisabled(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setAppearsDisabled:"), value)
}