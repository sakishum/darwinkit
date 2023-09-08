// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"github.com/progrium/macdriver/objc"
)

// The interface a net service uses to inform its delegate about the state of the service it offers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate?language=objc
type PNetServiceDelegate interface {
	// optional
	NetServiceDidNotResolve(sender NetService, errorDict map[string]Number)
	HasNetServiceDidNotResolve() bool

	// optional
	NetServiceDidResolveAddress(sender NetService)
	HasNetServiceDidResolveAddress() bool

	// optional
	NetServiceDidStop(sender NetService)
	HasNetServiceDidStop() bool

	// optional
	NetServiceDidAcceptConnectionWithInputStreamOutputStream(sender NetService, inputStream InputStream, outputStream OutputStream)
	HasNetServiceDidAcceptConnectionWithInputStreamOutputStream() bool

	// optional
	NetServiceDidPublish(sender NetService)
	HasNetServiceDidPublish() bool

	// optional
	NetServiceDidUpdateTXTRecordData(sender NetService, data []byte)
	HasNetServiceDidUpdateTXTRecordData() bool

	// optional
	NetServiceWillPublish(sender NetService)
	HasNetServiceWillPublish() bool

	// optional
	NetServiceDidNotPublish(sender NetService, errorDict map[string]Number)
	HasNetServiceDidNotPublish() bool

	// optional
	NetServiceWillResolve(sender NetService)
	HasNetServiceWillResolve() bool
}

// A delegate implementation builder for the [PNetServiceDelegate] protocol.
type NetServiceDelegate struct {
	_NetServiceDidNotResolve                                  func(sender NetService, errorDict map[string]Number)
	_NetServiceDidResolveAddress                              func(sender NetService)
	_NetServiceDidStop                                        func(sender NetService)
	_NetServiceDidAcceptConnectionWithInputStreamOutputStream func(sender NetService, inputStream InputStream, outputStream OutputStream)
	_NetServiceDidPublish                                     func(sender NetService)
	_NetServiceDidUpdateTXTRecordData                         func(sender NetService, data []byte)
	_NetServiceWillPublish                                    func(sender NetService)
	_NetServiceDidNotPublish                                  func(sender NetService, errorDict map[string]Number)
	_NetServiceWillResolve                                    func(sender NetService)
}

func (di *NetServiceDelegate) HasNetServiceDidNotResolve() bool {
	return di._NetServiceDidNotResolve != nil
}

// Informs the delegate that an error occurred during resolution of a given service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1414161-netservice?language=objc
func (di *NetServiceDelegate) SetNetServiceDidNotResolve(f func(sender NetService, errorDict map[string]Number)) {
	di._NetServiceDidNotResolve = f
}

// Informs the delegate that an error occurred during resolution of a given service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1414161-netservice?language=objc
func (di *NetServiceDelegate) NetServiceDidNotResolve(sender NetService, errorDict map[string]Number) {
	di._NetServiceDidNotResolve(sender, errorDict)
}
func (di *NetServiceDelegate) HasNetServiceDidResolveAddress() bool {
	return di._NetServiceDidResolveAddress != nil
}

// Informs the delegate that the address for a given service was resolved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1408457-netservicedidresolveaddress?language=objc
func (di *NetServiceDelegate) SetNetServiceDidResolveAddress(f func(sender NetService)) {
	di._NetServiceDidResolveAddress = f
}

// Informs the delegate that the address for a given service was resolved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1408457-netservicedidresolveaddress?language=objc
func (di *NetServiceDelegate) NetServiceDidResolveAddress(sender NetService) {
	di._NetServiceDidResolveAddress(sender)
}
func (di *NetServiceDelegate) HasNetServiceDidStop() bool {
	return di._NetServiceDidStop != nil
}

// Informs the delegate that a [foundation/nsnetservice/publish] or [foundation/nsnetservice/resolvewithtimeout] request was stopped. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1409726-netservicedidstop?language=objc
func (di *NetServiceDelegate) SetNetServiceDidStop(f func(sender NetService)) {
	di._NetServiceDidStop = f
}

// Informs the delegate that a [foundation/nsnetservice/publish] or [foundation/nsnetservice/resolvewithtimeout] request was stopped. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1409726-netservicedidstop?language=objc
func (di *NetServiceDelegate) NetServiceDidStop(sender NetService) {
	di._NetServiceDidStop(sender)
}
func (di *NetServiceDelegate) HasNetServiceDidAcceptConnectionWithInputStreamOutputStream() bool {
	return di._NetServiceDidAcceptConnectionWithInputStreamOutputStream != nil
}

// Called when a client connects to a service managed by Bonjour. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1407489-netservice?language=objc
func (di *NetServiceDelegate) SetNetServiceDidAcceptConnectionWithInputStreamOutputStream(f func(sender NetService, inputStream InputStream, outputStream OutputStream)) {
	di._NetServiceDidAcceptConnectionWithInputStreamOutputStream = f
}

// Called when a client connects to a service managed by Bonjour. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1407489-netservice?language=objc
func (di *NetServiceDelegate) NetServiceDidAcceptConnectionWithInputStreamOutputStream(sender NetService, inputStream InputStream, outputStream OutputStream) {
	di._NetServiceDidAcceptConnectionWithInputStreamOutputStream(sender, inputStream, outputStream)
}
func (di *NetServiceDelegate) HasNetServiceDidPublish() bool {
	return di._NetServiceDidPublish != nil
}

// Notifies the delegate that a service was successfully published. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1416802-netservicedidpublish?language=objc
func (di *NetServiceDelegate) SetNetServiceDidPublish(f func(sender NetService)) {
	di._NetServiceDidPublish = f
}

// Notifies the delegate that a service was successfully published. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1416802-netservicedidpublish?language=objc
func (di *NetServiceDelegate) NetServiceDidPublish(sender NetService) {
	di._NetServiceDidPublish(sender)
}
func (di *NetServiceDelegate) HasNetServiceDidUpdateTXTRecordData() bool {
	return di._NetServiceDidUpdateTXTRecordData != nil
}

// Notifies the delegate that the TXT record for a given service has been updated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1413199-netservice?language=objc
func (di *NetServiceDelegate) SetNetServiceDidUpdateTXTRecordData(f func(sender NetService, data []byte)) {
	di._NetServiceDidUpdateTXTRecordData = f
}

// Notifies the delegate that the TXT record for a given service has been updated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1413199-netservice?language=objc
func (di *NetServiceDelegate) NetServiceDidUpdateTXTRecordData(sender NetService, data []byte) {
	di._NetServiceDidUpdateTXTRecordData(sender, data)
}
func (di *NetServiceDelegate) HasNetServiceWillPublish() bool {
	return di._NetServiceWillPublish != nil
}

// Notifies the delegate that the network is ready to publish the service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1414277-netservicewillpublish?language=objc
func (di *NetServiceDelegate) SetNetServiceWillPublish(f func(sender NetService)) {
	di._NetServiceWillPublish = f
}

// Notifies the delegate that the network is ready to publish the service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1414277-netservicewillpublish?language=objc
func (di *NetServiceDelegate) NetServiceWillPublish(sender NetService) {
	di._NetServiceWillPublish(sender)
}
func (di *NetServiceDelegate) HasNetServiceDidNotPublish() bool {
	return di._NetServiceDidNotPublish != nil
}

// Notifies the delegate that a service could not be published. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1417101-netservice?language=objc
func (di *NetServiceDelegate) SetNetServiceDidNotPublish(f func(sender NetService, errorDict map[string]Number)) {
	di._NetServiceDidNotPublish = f
}

// Notifies the delegate that a service could not be published. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1417101-netservice?language=objc
func (di *NetServiceDelegate) NetServiceDidNotPublish(sender NetService, errorDict map[string]Number) {
	di._NetServiceDidNotPublish(sender, errorDict)
}
func (di *NetServiceDelegate) HasNetServiceWillResolve() bool {
	return di._NetServiceWillResolve != nil
}

// Notifies the delegate that the network is ready to resolve the service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1416022-netservicewillresolve?language=objc
func (di *NetServiceDelegate) SetNetServiceWillResolve(f func(sender NetService)) {
	di._NetServiceWillResolve = f
}

// Notifies the delegate that the network is ready to resolve the service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1416022-netservicewillresolve?language=objc
func (di *NetServiceDelegate) NetServiceWillResolve(sender NetService) {
	di._NetServiceWillResolve(sender)
}

// ensure impl type implements protocol interface
var _ PNetServiceDelegate = (*NetServiceDelegateObject)(nil)

// A concrete type for the [PNetServiceDelegate] protocol.
type NetServiceDelegateObject struct {
	objc.Object
}

func (n_ NetServiceDelegateObject) HasNetServiceDidNotResolve() bool {
	return n_.RespondsToSelector(objc.Sel("netService:didNotResolve:"))
}

// Informs the delegate that an error occurred during resolution of a given service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1414161-netservice?language=objc
func (n_ NetServiceDelegateObject) NetServiceDidNotResolve(sender NetService, errorDict map[string]Number) {
	objc.Call[objc.Void](n_, objc.Sel("netService:didNotResolve:"), objc.Ptr(sender), errorDict)
}

func (n_ NetServiceDelegateObject) HasNetServiceDidResolveAddress() bool {
	return n_.RespondsToSelector(objc.Sel("netServiceDidResolveAddress:"))
}

// Informs the delegate that the address for a given service was resolved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1408457-netservicedidresolveaddress?language=objc
func (n_ NetServiceDelegateObject) NetServiceDidResolveAddress(sender NetService) {
	objc.Call[objc.Void](n_, objc.Sel("netServiceDidResolveAddress:"), objc.Ptr(sender))
}

func (n_ NetServiceDelegateObject) HasNetServiceDidStop() bool {
	return n_.RespondsToSelector(objc.Sel("netServiceDidStop:"))
}

// Informs the delegate that a [foundation/nsnetservice/publish] or [foundation/nsnetservice/resolvewithtimeout] request was stopped. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1409726-netservicedidstop?language=objc
func (n_ NetServiceDelegateObject) NetServiceDidStop(sender NetService) {
	objc.Call[objc.Void](n_, objc.Sel("netServiceDidStop:"), objc.Ptr(sender))
}

func (n_ NetServiceDelegateObject) HasNetServiceDidAcceptConnectionWithInputStreamOutputStream() bool {
	return n_.RespondsToSelector(objc.Sel("netService:didAcceptConnectionWithInputStream:outputStream:"))
}

// Called when a client connects to a service managed by Bonjour. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1407489-netservice?language=objc
func (n_ NetServiceDelegateObject) NetServiceDidAcceptConnectionWithInputStreamOutputStream(sender NetService, inputStream InputStream, outputStream OutputStream) {
	objc.Call[objc.Void](n_, objc.Sel("netService:didAcceptConnectionWithInputStream:outputStream:"), objc.Ptr(sender), objc.Ptr(inputStream), objc.Ptr(outputStream))
}

func (n_ NetServiceDelegateObject) HasNetServiceDidPublish() bool {
	return n_.RespondsToSelector(objc.Sel("netServiceDidPublish:"))
}

// Notifies the delegate that a service was successfully published. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1416802-netservicedidpublish?language=objc
func (n_ NetServiceDelegateObject) NetServiceDidPublish(sender NetService) {
	objc.Call[objc.Void](n_, objc.Sel("netServiceDidPublish:"), objc.Ptr(sender))
}

func (n_ NetServiceDelegateObject) HasNetServiceDidUpdateTXTRecordData() bool {
	return n_.RespondsToSelector(objc.Sel("netService:didUpdateTXTRecordData:"))
}

// Notifies the delegate that the TXT record for a given service has been updated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1413199-netservice?language=objc
func (n_ NetServiceDelegateObject) NetServiceDidUpdateTXTRecordData(sender NetService, data []byte) {
	objc.Call[objc.Void](n_, objc.Sel("netService:didUpdateTXTRecordData:"), objc.Ptr(sender), data)
}

func (n_ NetServiceDelegateObject) HasNetServiceWillPublish() bool {
	return n_.RespondsToSelector(objc.Sel("netServiceWillPublish:"))
}

// Notifies the delegate that the network is ready to publish the service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1414277-netservicewillpublish?language=objc
func (n_ NetServiceDelegateObject) NetServiceWillPublish(sender NetService) {
	objc.Call[objc.Void](n_, objc.Sel("netServiceWillPublish:"), objc.Ptr(sender))
}

func (n_ NetServiceDelegateObject) HasNetServiceDidNotPublish() bool {
	return n_.RespondsToSelector(objc.Sel("netService:didNotPublish:"))
}

// Notifies the delegate that a service could not be published. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1417101-netservice?language=objc
func (n_ NetServiceDelegateObject) NetServiceDidNotPublish(sender NetService, errorDict map[string]Number) {
	objc.Call[objc.Void](n_, objc.Sel("netService:didNotPublish:"), objc.Ptr(sender), errorDict)
}

func (n_ NetServiceDelegateObject) HasNetServiceWillResolve() bool {
	return n_.RespondsToSelector(objc.Sel("netServiceWillResolve:"))
}

// Notifies the delegate that the network is ready to resolve the service. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnetservicedelegate/1416022-netservicewillresolve?language=objc
func (n_ NetServiceDelegateObject) NetServiceWillResolve(sender NetService) {
	objc.Call[objc.Void](n_, objc.Sel("netServiceWillResolve:"), objc.Ptr(sender))
}
