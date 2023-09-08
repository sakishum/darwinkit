// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/macdriver/macos/quartzcore"
	"github.com/progrium/macdriver/objc"
)

// An optional layer delegate method for handling resolution changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewlayercontentscaledelegate?language=objc
type PViewLayerContentScaleDelegate interface {
	// optional
	LayerShouldInheritContentsScaleFromWindow(layer quartzcore.Layer, newScale float64, window Window) bool
	HasLayerShouldInheritContentsScaleFromWindow() bool
}

// A delegate implementation builder for the [PViewLayerContentScaleDelegate] protocol.
type ViewLayerContentScaleDelegate struct {
	_LayerShouldInheritContentsScaleFromWindow func(layer quartzcore.Layer, newScale float64, window Window) bool
}

func (di *ViewLayerContentScaleDelegate) HasLayerShouldInheritContentsScaleFromWindow() bool {
	return di._LayerShouldInheritContentsScaleFromWindow != nil
}

// Notifies you when a resolution changes occurs for the window that hosts the layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewlayercontentscaledelegate/3005294-layer?language=objc
func (di *ViewLayerContentScaleDelegate) SetLayerShouldInheritContentsScaleFromWindow(f func(layer quartzcore.Layer, newScale float64, window Window) bool) {
	di._LayerShouldInheritContentsScaleFromWindow = f
}

// Notifies you when a resolution changes occurs for the window that hosts the layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewlayercontentscaledelegate/3005294-layer?language=objc
func (di *ViewLayerContentScaleDelegate) LayerShouldInheritContentsScaleFromWindow(layer quartzcore.Layer, newScale float64, window Window) bool {
	return di._LayerShouldInheritContentsScaleFromWindow(layer, newScale, window)
}

// ensure impl type implements protocol interface
var _ PViewLayerContentScaleDelegate = (*ViewLayerContentScaleDelegateObject)(nil)

// A concrete type for the [PViewLayerContentScaleDelegate] protocol.
type ViewLayerContentScaleDelegateObject struct {
	objc.Object
}

func (v_ ViewLayerContentScaleDelegateObject) HasLayerShouldInheritContentsScaleFromWindow() bool {
	return v_.RespondsToSelector(objc.Sel("layer:shouldInheritContentsScale:fromWindow:"))
}

// Notifies you when a resolution changes occurs for the window that hosts the layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewlayercontentscaledelegate/3005294-layer?language=objc
func (v_ ViewLayerContentScaleDelegateObject) LayerShouldInheritContentsScaleFromWindow(layer quartzcore.Layer, newScale float64, window Window) bool {
	rv := objc.Call[bool](v_, objc.Sel("layer:shouldInheritContentsScale:fromWindow:"), objc.Ptr(layer), newScale, objc.Ptr(window))
	return rv
}
