// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A set of optional methods implemented by delegates of NSControl subclasses to respond to editing actions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate?language=objc
type PControlTextEditingDelegate interface {
	// optional
	ControlTextDidBeginEditing(obj foundation.Notification)
	HasControlTextDidBeginEditing() bool

	// optional
	ControlIsValidObject(control Control, obj objc.Object) bool
	HasControlIsValidObject() bool

	// optional
	ControlTextViewDoCommandBySelector(control Control, textView TextView, commandSelector objc.Selector) bool
	HasControlTextViewDoCommandBySelector() bool

	// optional
	ControlTextShouldBeginEditing(control Control, fieldEditor Text) bool
	HasControlTextShouldBeginEditing() bool

	// optional
	ControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem(control Control, textView TextView, words []string, charRange foundation.Range, index *int) []string
	HasControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem() bool

	// optional
	ControlDidFailToValidatePartialStringErrorDescription(control Control, string_ string, error string)
	HasControlDidFailToValidatePartialStringErrorDescription() bool

	// optional
	ControlTextDidChange(obj foundation.Notification)
	HasControlTextDidChange() bool

	// optional
	ControlDidFailToFormatStringErrorDescription(control Control, string_ string, error string) bool
	HasControlDidFailToFormatStringErrorDescription() bool

	// optional
	ControlTextDidEndEditing(obj foundation.Notification)
	HasControlTextDidEndEditing() bool

	// optional
	ControlTextShouldEndEditing(control Control, fieldEditor Text) bool
	HasControlTextShouldEndEditing() bool
}

// A delegate implementation builder for the [PControlTextEditingDelegate] protocol.
type ControlTextEditingDelegate struct {
	_ControlTextDidBeginEditing                                       func(obj foundation.Notification)
	_ControlIsValidObject                                             func(control Control, obj objc.Object) bool
	_ControlTextViewDoCommandBySelector                               func(control Control, textView TextView, commandSelector objc.Selector) bool
	_ControlTextShouldBeginEditing                                    func(control Control, fieldEditor Text) bool
	_ControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem func(control Control, textView TextView, words []string, charRange foundation.Range, index *int) []string
	_ControlDidFailToValidatePartialStringErrorDescription            func(control Control, string_ string, error string)
	_ControlTextDidChange                                             func(obj foundation.Notification)
	_ControlDidFailToFormatStringErrorDescription                     func(control Control, string_ string, error string) bool
	_ControlTextDidEndEditing                                         func(obj foundation.Notification)
	_ControlTextShouldEndEditing                                      func(control Control, fieldEditor Text) bool
}

func (di *ControlTextEditingDelegate) HasControlTextDidBeginEditing() bool {
	return di._ControlTextDidBeginEditing != nil
}

// Tells the delegate that the control started editing its text content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/3005176-controltextdidbeginediting?language=objc
func (di *ControlTextEditingDelegate) SetControlTextDidBeginEditing(f func(obj foundation.Notification)) {
	di._ControlTextDidBeginEditing = f
}

// Tells the delegate that the control started editing its text content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/3005176-controltextdidbeginediting?language=objc
func (di *ControlTextEditingDelegate) ControlTextDidBeginEditing(obj foundation.Notification) {
	di._ControlTextDidBeginEditing(obj)
}
func (di *ControlTextEditingDelegate) HasControlIsValidObject() bool {
	return di._ControlIsValidObject != nil
}

// Invoked when the insertion point leaves a cell belonging to the specified control, but before the value of the cell’s object is displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/1428873-control?language=objc
func (di *ControlTextEditingDelegate) SetControlIsValidObject(f func(control Control, obj objc.Object) bool) {
	di._ControlIsValidObject = f
}

// Invoked when the insertion point leaves a cell belonging to the specified control, but before the value of the cell’s object is displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/1428873-control?language=objc
func (di *ControlTextEditingDelegate) ControlIsValidObject(control Control, obj objc.Object) bool {
	return di._ControlIsValidObject(control, obj)
}
func (di *ControlTextEditingDelegate) HasControlTextViewDoCommandBySelector() bool {
	return di._ControlTextViewDoCommandBySelector != nil
}

// Invoked when users press keys with predefined bindings in a cell of the specified control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/1428898-control?language=objc
func (di *ControlTextEditingDelegate) SetControlTextViewDoCommandBySelector(f func(control Control, textView TextView, commandSelector objc.Selector) bool) {
	di._ControlTextViewDoCommandBySelector = f
}

// Invoked when users press keys with predefined bindings in a cell of the specified control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/1428898-control?language=objc
func (di *ControlTextEditingDelegate) ControlTextViewDoCommandBySelector(control Control, textView TextView, commandSelector objc.Selector) bool {
	return di._ControlTextViewDoCommandBySelector(control, textView, commandSelector)
}
func (di *ControlTextEditingDelegate) HasControlTextShouldBeginEditing() bool {
	return di._ControlTextShouldBeginEditing != nil
}

// Invoked when the user tries to enter a character in a cell of a control that allows editing of text (such as a text field or form field). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/1428865-control?language=objc
func (di *ControlTextEditingDelegate) SetControlTextShouldBeginEditing(f func(control Control, fieldEditor Text) bool) {
	di._ControlTextShouldBeginEditing = f
}

// Invoked when the user tries to enter a character in a cell of a control that allows editing of text (such as a text field or form field). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/1428865-control?language=objc
func (di *ControlTextEditingDelegate) ControlTextShouldBeginEditing(control Control, fieldEditor Text) bool {
	return di._ControlTextShouldBeginEditing(control, fieldEditor)
}
func (di *ControlTextEditingDelegate) HasControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem() bool {
	return di._ControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem != nil
}

// Invoked to allow you to control the list of proposed text completions generated by text fields and other controls. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/1428925-control?language=objc
func (di *ControlTextEditingDelegate) SetControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem(f func(control Control, textView TextView, words []string, charRange foundation.Range, index *int) []string) {
	di._ControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem = f
}

// Invoked to allow you to control the list of proposed text completions generated by text fields and other controls. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/1428925-control?language=objc
func (di *ControlTextEditingDelegate) ControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem(control Control, textView TextView, words []string, charRange foundation.Range, index *int) []string {
	return di._ControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem(control, textView, words, charRange, index)
}
func (di *ControlTextEditingDelegate) HasControlDidFailToValidatePartialStringErrorDescription() bool {
	return di._ControlDidFailToValidatePartialStringErrorDescription != nil
}

// Invoked when the formatter for the cell belonging to control (or selected cell) rejects a partial string a user is typing into the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/1428941-control?language=objc
func (di *ControlTextEditingDelegate) SetControlDidFailToValidatePartialStringErrorDescription(f func(control Control, string_ string, error string)) {
	di._ControlDidFailToValidatePartialStringErrorDescription = f
}

// Invoked when the formatter for the cell belonging to control (or selected cell) rejects a partial string a user is typing into the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/1428941-control?language=objc
func (di *ControlTextEditingDelegate) ControlDidFailToValidatePartialStringErrorDescription(control Control, string_ string, error string) {
	di._ControlDidFailToValidatePartialStringErrorDescription(control, string_, error)
}
func (di *ControlTextEditingDelegate) HasControlTextDidChange() bool {
	return di._ControlTextDidChange != nil
}

// Tells the delegate that the control made changes to its text content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/3005177-controltextdidchange?language=objc
func (di *ControlTextEditingDelegate) SetControlTextDidChange(f func(obj foundation.Notification)) {
	di._ControlTextDidChange = f
}

// Tells the delegate that the control made changes to its text content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/3005177-controltextdidchange?language=objc
func (di *ControlTextEditingDelegate) ControlTextDidChange(obj foundation.Notification) {
	di._ControlTextDidChange(obj)
}
func (di *ControlTextEditingDelegate) HasControlDidFailToFormatStringErrorDescription() bool {
	return di._ControlDidFailToFormatStringErrorDescription != nil
}

// Invoked when the formatter for the cell belonging to the specified control cannot convert a string to an underlying object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/1428883-control?language=objc
func (di *ControlTextEditingDelegate) SetControlDidFailToFormatStringErrorDescription(f func(control Control, string_ string, error string) bool) {
	di._ControlDidFailToFormatStringErrorDescription = f
}

// Invoked when the formatter for the cell belonging to the specified control cannot convert a string to an underlying object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/1428883-control?language=objc
func (di *ControlTextEditingDelegate) ControlDidFailToFormatStringErrorDescription(control Control, string_ string, error string) bool {
	return di._ControlDidFailToFormatStringErrorDescription(control, string_, error)
}
func (di *ControlTextEditingDelegate) HasControlTextDidEndEditing() bool {
	return di._ControlTextDidEndEditing != nil
}

// Tells the delegate that the control finished editing its text content and committed the changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/3005178-controltextdidendediting?language=objc
func (di *ControlTextEditingDelegate) SetControlTextDidEndEditing(f func(obj foundation.Notification)) {
	di._ControlTextDidEndEditing = f
}

// Tells the delegate that the control finished editing its text content and committed the changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/3005178-controltextdidendediting?language=objc
func (di *ControlTextEditingDelegate) ControlTextDidEndEditing(obj foundation.Notification) {
	di._ControlTextDidEndEditing(obj)
}
func (di *ControlTextEditingDelegate) HasControlTextShouldEndEditing() bool {
	return di._ControlTextShouldEndEditing != nil
}

// Invoked when the insertion point tries to leave a cell of the control that has been edited. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/1428984-control?language=objc
func (di *ControlTextEditingDelegate) SetControlTextShouldEndEditing(f func(control Control, fieldEditor Text) bool) {
	di._ControlTextShouldEndEditing = f
}

// Invoked when the insertion point tries to leave a cell of the control that has been edited. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/1428984-control?language=objc
func (di *ControlTextEditingDelegate) ControlTextShouldEndEditing(control Control, fieldEditor Text) bool {
	return di._ControlTextShouldEndEditing(control, fieldEditor)
}

// ensure impl type implements protocol interface
var _ PControlTextEditingDelegate = (*ControlTextEditingDelegateObject)(nil)

// A concrete type for the [PControlTextEditingDelegate] protocol.
type ControlTextEditingDelegateObject struct {
	objc.Object
}

func (c_ ControlTextEditingDelegateObject) HasControlTextDidBeginEditing() bool {
	return c_.RespondsToSelector(objc.Sel("controlTextDidBeginEditing:"))
}

// Tells the delegate that the control started editing its text content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/3005176-controltextdidbeginediting?language=objc
func (c_ ControlTextEditingDelegateObject) ControlTextDidBeginEditing(obj foundation.Notification) {
	objc.Call[objc.Void](c_, objc.Sel("controlTextDidBeginEditing:"), objc.Ptr(obj))
}

func (c_ ControlTextEditingDelegateObject) HasControlIsValidObject() bool {
	return c_.RespondsToSelector(objc.Sel("control:isValidObject:"))
}

// Invoked when the insertion point leaves a cell belonging to the specified control, but before the value of the cell’s object is displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/1428873-control?language=objc
func (c_ ControlTextEditingDelegateObject) ControlIsValidObject(control Control, obj objc.Object) bool {
	rv := objc.Call[bool](c_, objc.Sel("control:isValidObject:"), objc.Ptr(control), obj)
	return rv
}

func (c_ ControlTextEditingDelegateObject) HasControlTextViewDoCommandBySelector() bool {
	return c_.RespondsToSelector(objc.Sel("control:textView:doCommandBySelector:"))
}

// Invoked when users press keys with predefined bindings in a cell of the specified control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/1428898-control?language=objc
func (c_ ControlTextEditingDelegateObject) ControlTextViewDoCommandBySelector(control Control, textView TextView, commandSelector objc.Selector) bool {
	rv := objc.Call[bool](c_, objc.Sel("control:textView:doCommandBySelector:"), objc.Ptr(control), objc.Ptr(textView), commandSelector)
	return rv
}

func (c_ ControlTextEditingDelegateObject) HasControlTextShouldBeginEditing() bool {
	return c_.RespondsToSelector(objc.Sel("control:textShouldBeginEditing:"))
}

// Invoked when the user tries to enter a character in a cell of a control that allows editing of text (such as a text field or form field). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/1428865-control?language=objc
func (c_ ControlTextEditingDelegateObject) ControlTextShouldBeginEditing(control Control, fieldEditor Text) bool {
	rv := objc.Call[bool](c_, objc.Sel("control:textShouldBeginEditing:"), objc.Ptr(control), objc.Ptr(fieldEditor))
	return rv
}

func (c_ ControlTextEditingDelegateObject) HasControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem() bool {
	return c_.RespondsToSelector(objc.Sel("control:textView:completions:forPartialWordRange:indexOfSelectedItem:"))
}

// Invoked to allow you to control the list of proposed text completions generated by text fields and other controls. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/1428925-control?language=objc
func (c_ ControlTextEditingDelegateObject) ControlTextViewCompletionsForPartialWordRangeIndexOfSelectedItem(control Control, textView TextView, words []string, charRange foundation.Range, index *int) []string {
	rv := objc.Call[[]string](c_, objc.Sel("control:textView:completions:forPartialWordRange:indexOfSelectedItem:"), objc.Ptr(control), objc.Ptr(textView), words, charRange, index)
	return rv
}

func (c_ ControlTextEditingDelegateObject) HasControlDidFailToValidatePartialStringErrorDescription() bool {
	return c_.RespondsToSelector(objc.Sel("control:didFailToValidatePartialString:errorDescription:"))
}

// Invoked when the formatter for the cell belonging to control (or selected cell) rejects a partial string a user is typing into the cell. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/1428941-control?language=objc
func (c_ ControlTextEditingDelegateObject) ControlDidFailToValidatePartialStringErrorDescription(control Control, string_ string, error string) {
	objc.Call[objc.Void](c_, objc.Sel("control:didFailToValidatePartialString:errorDescription:"), objc.Ptr(control), string_, error)
}

func (c_ ControlTextEditingDelegateObject) HasControlTextDidChange() bool {
	return c_.RespondsToSelector(objc.Sel("controlTextDidChange:"))
}

// Tells the delegate that the control made changes to its text content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/3005177-controltextdidchange?language=objc
func (c_ ControlTextEditingDelegateObject) ControlTextDidChange(obj foundation.Notification) {
	objc.Call[objc.Void](c_, objc.Sel("controlTextDidChange:"), objc.Ptr(obj))
}

func (c_ ControlTextEditingDelegateObject) HasControlDidFailToFormatStringErrorDescription() bool {
	return c_.RespondsToSelector(objc.Sel("control:didFailToFormatString:errorDescription:"))
}

// Invoked when the formatter for the cell belonging to the specified control cannot convert a string to an underlying object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/1428883-control?language=objc
func (c_ ControlTextEditingDelegateObject) ControlDidFailToFormatStringErrorDescription(control Control, string_ string, error string) bool {
	rv := objc.Call[bool](c_, objc.Sel("control:didFailToFormatString:errorDescription:"), objc.Ptr(control), string_, error)
	return rv
}

func (c_ ControlTextEditingDelegateObject) HasControlTextDidEndEditing() bool {
	return c_.RespondsToSelector(objc.Sel("controlTextDidEndEditing:"))
}

// Tells the delegate that the control finished editing its text content and committed the changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/3005178-controltextdidendediting?language=objc
func (c_ ControlTextEditingDelegateObject) ControlTextDidEndEditing(obj foundation.Notification) {
	objc.Call[objc.Void](c_, objc.Sel("controlTextDidEndEditing:"), objc.Ptr(obj))
}

func (c_ ControlTextEditingDelegateObject) HasControlTextShouldEndEditing() bool {
	return c_.RespondsToSelector(objc.Sel("control:textShouldEndEditing:"))
}

// Invoked when the insertion point tries to leave a cell of the control that has been edited. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontroltexteditingdelegate/1428984-control?language=objc
func (c_ ControlTextEditingDelegateObject) ControlTextShouldEndEditing(control Control, fieldEditor Text) bool {
	rv := objc.Call[bool](c_, objc.Sel("control:textShouldEndEditing:"), objc.Ptr(control), objc.Ptr(fieldEditor))
	return rv
}
