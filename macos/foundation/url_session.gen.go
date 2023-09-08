// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [URLSession] class.
var URLSessionClass = _URLSessionClass{objc.GetClass("NSURLSession")}

type _URLSessionClass struct {
	objc.Class
}

// An interface definition for the [URLSession] class.
type IURLSession interface {
	objc.IObject
	InvalidateAndCancel()
	WebSocketTaskWithURLProtocols(url IURL, protocols []string) URLSessionWebSocketTask
	DownloadTaskWithURL(url IURL) URLSessionDownloadTask
	DownloadTaskWithURLCompletionHandler(url IURL, completionHandler func(location URL, response URLResponse, error Error)) URLSessionDownloadTask
	GetAllTasksWithCompletionHandler(completionHandler func(tasks []URLSessionTask))
	DownloadTaskWithRequestCompletionHandler(request IURLRequest, completionHandler func(location URL, response URLResponse, error Error)) URLSessionDownloadTask
	DataTaskWithRequestCompletionHandler(request IURLRequest, completionHandler func(data []byte, response URLResponse, error Error)) URLSessionDataTask
	FinishTasksAndInvalidate()
	FlushWithCompletionHandler(completionHandler func())
	DataTaskWithURL(url IURL) URLSessionDataTask
	GetTasksWithCompletionHandler(completionHandler func(dataTasks []URLSessionDataTask, uploadTasks []URLSessionUploadTask, downloadTasks []URLSessionDownloadTask))
	UploadTaskWithRequestFromData(request IURLRequest, bodyData []byte) URLSessionUploadTask
	DownloadTaskWithResumeData(resumeData []byte) URLSessionDownloadTask
	DownloadTaskWithRequest(request IURLRequest) URLSessionDownloadTask
	UploadTaskWithRequestFromDataCompletionHandler(request IURLRequest, bodyData []byte, completionHandler func(data []byte, response URLResponse, error Error)) URLSessionUploadTask
	WebSocketTaskWithURL(url IURL) URLSessionWebSocketTask
	DataTaskWithURLCompletionHandler(url IURL, completionHandler func(data []byte, response URLResponse, error Error)) URLSessionDataTask
	ResetWithCompletionHandler(completionHandler func())
	UploadTaskWithRequestFromFile(request IURLRequest, fileURL IURL) URLSessionUploadTask
	DownloadTaskWithResumeDataCompletionHandler(resumeData []byte, completionHandler func(location URL, response URLResponse, error Error)) URLSessionDownloadTask
	StreamTaskWithHostNamePort(hostname string, port int) URLSessionStreamTask
	WebSocketTaskWithRequest(request IURLRequest) URLSessionWebSocketTask
	UploadTaskWithStreamedRequest(request IURLRequest) URLSessionUploadTask
	DataTaskWithRequest(request IURLRequest) URLSessionDataTask
	UploadTaskWithRequestFromFileCompletionHandler(request IURLRequest, fileURL IURL, completionHandler func(data []byte, response URLResponse, error Error)) URLSessionUploadTask
	Delegate() URLSessionDelegateObject
	DelegateQueue() OperationQueue
	SessionDescription() string
	SetSessionDescription(value string)
	Configuration() URLSessionConfiguration
}

// An object that coordinates a group of related, network data transfer tasks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession?language=objc
type URLSession struct {
	objc.Object
}

func URLSessionFrom(ptr unsafe.Pointer) URLSession {
	return URLSession{
		Object: objc.ObjectFrom(ptr),
	}
}

func (uc _URLSessionClass) Alloc() URLSession {
	rv := objc.Call[URLSession](uc, objc.Sel("alloc"))
	return rv
}

func (uc _URLSessionClass) New() URLSession {
	rv := objc.Call[URLSession](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewURLSession() URLSession {
	return URLSessionClass.New()
}

func (u_ URLSession) Init() URLSession {
	rv := objc.Call[URLSession](u_, objc.Sel("init"))
	return rv
}

// Cancels all outstanding tasks and then invalidates the session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411538-invalidateandcancel?language=objc
func (u_ URLSession) InvalidateAndCancel() {
	objc.Call[objc.Void](u_, objc.Sel("invalidateAndCancel"))
}

// Creates a session with the specified session configuration, delegate, and operation queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411597-sessionwithconfiguration?language=objc
func (uc _URLSessionClass) SessionWithConfigurationDelegateDelegateQueue(configuration IURLSessionConfiguration, delegate PURLSessionDelegate, queue IOperationQueue) URLSession {
	po1 := objc.WrapAsProtocol("NSURLSessionDelegate", delegate)
	rv := objc.Call[URLSession](uc, objc.Sel("sessionWithConfiguration:delegate:delegateQueue:"), objc.Ptr(configuration), po1, objc.Ptr(queue))
	return rv
}

// Creates a session with the specified session configuration, delegate, and operation queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411597-sessionwithconfiguration?language=objc
func URLSession_SessionWithConfigurationDelegateDelegateQueue(configuration IURLSessionConfiguration, delegate PURLSessionDelegate, queue IOperationQueue) URLSession {
	return URLSessionClass.SessionWithConfigurationDelegateDelegateQueue(configuration, delegate, queue)
}

// Creates a session with the specified session configuration, delegate, and operation queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411597-sessionwithconfiguration?language=objc
func (uc _URLSessionClass) SessionWithConfigurationDelegateObjectDelegateQueue(configuration IURLSessionConfiguration, delegateObject objc.IObject, queue IOperationQueue) URLSession {
	rv := objc.Call[URLSession](uc, objc.Sel("sessionWithConfiguration:delegate:delegateQueue:"), objc.Ptr(configuration), objc.Ptr(delegateObject), objc.Ptr(queue))
	return rv
}

// Creates a session with the specified session configuration, delegate, and operation queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411597-sessionwithconfiguration?language=objc
func URLSession_SessionWithConfigurationDelegateObjectDelegateQueue(configuration IURLSessionConfiguration, delegateObject objc.IObject, queue IOperationQueue) URLSession {
	return URLSessionClass.SessionWithConfigurationDelegateObjectDelegateQueue(configuration, delegateObject, queue)
}

// Creates a WebSocket task given a URL and an array of protocols. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/3181172-websockettaskwithurl?language=objc
func (u_ URLSession) WebSocketTaskWithURLProtocols(url IURL, protocols []string) URLSessionWebSocketTask {
	rv := objc.Call[URLSessionWebSocketTask](u_, objc.Sel("webSocketTaskWithURL:protocols:"), objc.Ptr(url), protocols)
	return rv
}

// Creates a download task that retrieves the contents of the specified URL and saves the results to a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411482-downloadtaskwithurl?language=objc
func (u_ URLSession) DownloadTaskWithURL(url IURL) URLSessionDownloadTask {
	rv := objc.Call[URLSessionDownloadTask](u_, objc.Sel("downloadTaskWithURL:"), objc.Ptr(url))
	return rv
}

// Creates a download task that retrieves the contents of the specified URL, saves the results to a file, and calls a handler upon completion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411608-downloadtaskwithurl?language=objc
func (u_ URLSession) DownloadTaskWithURLCompletionHandler(url IURL, completionHandler func(location URL, response URLResponse, error Error)) URLSessionDownloadTask {
	rv := objc.Call[URLSessionDownloadTask](u_, objc.Sel("downloadTaskWithURL:completionHandler:"), objc.Ptr(url), completionHandler)
	return rv
}

// Asynchronously calls a completion callback with all tasks in a session [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411618-getalltaskswithcompletionhandler?language=objc
func (u_ URLSession) GetAllTasksWithCompletionHandler(completionHandler func(tasks []URLSessionTask)) {
	objc.Call[objc.Void](u_, objc.Sel("getAllTasksWithCompletionHandler:"), completionHandler)
}

// Creates a download task that retrieves the contents of a URL based on the specified URL request object, saves the results to a file, and calls a handler upon completion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411511-downloadtaskwithrequest?language=objc
func (u_ URLSession) DownloadTaskWithRequestCompletionHandler(request IURLRequest, completionHandler func(location URL, response URLResponse, error Error)) URLSessionDownloadTask {
	rv := objc.Call[URLSessionDownloadTask](u_, objc.Sel("downloadTaskWithRequest:completionHandler:"), objc.Ptr(request), completionHandler)
	return rv
}

// Creates a task that retrieves the contents of a URL based on the specified URL request object, and calls a handler upon completion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1407613-datataskwithrequest?language=objc
func (u_ URLSession) DataTaskWithRequestCompletionHandler(request IURLRequest, completionHandler func(data []byte, response URLResponse, error Error)) URLSessionDataTask {
	rv := objc.Call[URLSessionDataTask](u_, objc.Sel("dataTaskWithRequest:completionHandler:"), objc.Ptr(request), completionHandler)
	return rv
}

// Invalidates the session, allowing any outstanding tasks to finish. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1407428-finishtasksandinvalidate?language=objc
func (u_ URLSession) FinishTasksAndInvalidate() {
	objc.Call[objc.Void](u_, objc.Sel("finishTasksAndInvalidate"))
}

// Flushes cookies and credentials to disk, clears transient caches, and ensures that future requests occur on a new TCP connection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411622-flushwithcompletionhandler?language=objc
func (u_ URLSession) FlushWithCompletionHandler(completionHandler func()) {
	objc.Call[objc.Void](u_, objc.Sel("flushWithCompletionHandler:"), completionHandler)
}

// Creates a task that retrieves the contents of the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411554-datataskwithurl?language=objc
func (u_ URLSession) DataTaskWithURL(url IURL) URLSessionDataTask {
	rv := objc.Call[URLSessionDataTask](u_, objc.Sel("dataTaskWithURL:"), objc.Ptr(url))
	return rv
}

// Asynchronously calls a completion callback with all data, upload, and download tasks in a session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411578-gettaskswithcompletionhandler?language=objc
func (u_ URLSession) GetTasksWithCompletionHandler(completionHandler func(dataTasks []URLSessionDataTask, uploadTasks []URLSessionUploadTask, downloadTasks []URLSessionDownloadTask)) {
	objc.Call[objc.Void](u_, objc.Sel("getTasksWithCompletionHandler:"), completionHandler)
}

// Creates a task that performs an HTTP request for the specified URL request object and uploads the provided data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1409763-uploadtaskwithrequest?language=objc
func (u_ URLSession) UploadTaskWithRequestFromData(request IURLRequest, bodyData []byte) URLSessionUploadTask {
	rv := objc.Call[URLSessionUploadTask](u_, objc.Sel("uploadTaskWithRequest:fromData:"), objc.Ptr(request), bodyData)
	return rv
}

// Creates a download task to resume a previously canceled or failed download. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1409226-downloadtaskwithresumedata?language=objc
func (u_ URLSession) DownloadTaskWithResumeData(resumeData []byte) URLSessionDownloadTask {
	rv := objc.Call[URLSessionDownloadTask](u_, objc.Sel("downloadTaskWithResumeData:"), resumeData)
	return rv
}

// Creates a download task that retrieves the contents of a URL based on the specified URL request object and saves the results to a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411481-downloadtaskwithrequest?language=objc
func (u_ URLSession) DownloadTaskWithRequest(request IURLRequest) URLSessionDownloadTask {
	rv := objc.Call[URLSessionDownloadTask](u_, objc.Sel("downloadTaskWithRequest:"), objc.Ptr(request))
	return rv
}

// Creates a task that performs an HTTP request for the specified URL request object, uploads the provided data, and calls a handler upon completion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411518-uploadtaskwithrequest?language=objc
func (u_ URLSession) UploadTaskWithRequestFromDataCompletionHandler(request IURLRequest, bodyData []byte, completionHandler func(data []byte, response URLResponse, error Error)) URLSessionUploadTask {
	rv := objc.Call[URLSessionUploadTask](u_, objc.Sel("uploadTaskWithRequest:fromData:completionHandler:"), objc.Ptr(request), bodyData, completionHandler)
	return rv
}

// Creates a WebSocket task for the provided URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/3181171-websockettaskwithurl?language=objc
func (u_ URLSession) WebSocketTaskWithURL(url IURL) URLSessionWebSocketTask {
	rv := objc.Call[URLSessionWebSocketTask](u_, objc.Sel("webSocketTaskWithURL:"), objc.Ptr(url))
	return rv
}

// Creates a session with the specified session configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411474-sessionwithconfiguration?language=objc
func (uc _URLSessionClass) SessionWithConfiguration(configuration IURLSessionConfiguration) URLSession {
	rv := objc.Call[URLSession](uc, objc.Sel("sessionWithConfiguration:"), objc.Ptr(configuration))
	return rv
}

// Creates a session with the specified session configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411474-sessionwithconfiguration?language=objc
func URLSession_SessionWithConfiguration(configuration IURLSessionConfiguration) URLSession {
	return URLSessionClass.SessionWithConfiguration(configuration)
}

// Creates a task that retrieves the contents of the specified URL, then calls a handler upon completion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1410330-datataskwithurl?language=objc
func (u_ URLSession) DataTaskWithURLCompletionHandler(url IURL, completionHandler func(data []byte, response URLResponse, error Error)) URLSessionDataTask {
	rv := objc.Call[URLSessionDataTask](u_, objc.Sel("dataTaskWithURL:completionHandler:"), objc.Ptr(url), completionHandler)
	return rv
}

// Empties all cookies, caches and credential stores, removes disk files, flushes in-progress downloads to disk, and ensures that future requests occur on a new socket. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411479-resetwithcompletionhandler?language=objc
func (u_ URLSession) ResetWithCompletionHandler(completionHandler func()) {
	objc.Call[objc.Void](u_, objc.Sel("resetWithCompletionHandler:"), completionHandler)
}

// Creates a task that performs an HTTP request for uploading the specified file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411550-uploadtaskwithrequest?language=objc
func (u_ URLSession) UploadTaskWithRequestFromFile(request IURLRequest, fileURL IURL) URLSessionUploadTask {
	rv := objc.Call[URLSessionUploadTask](u_, objc.Sel("uploadTaskWithRequest:fromFile:"), objc.Ptr(request), objc.Ptr(fileURL))
	return rv
}

// Creates a download task to resume a previously canceled or failed download and calls a handler upon completion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411598-downloadtaskwithresumedata?language=objc
func (u_ URLSession) DownloadTaskWithResumeDataCompletionHandler(resumeData []byte, completionHandler func(location URL, response URLResponse, error Error)) URLSessionDownloadTask {
	rv := objc.Call[URLSessionDownloadTask](u_, objc.Sel("downloadTaskWithResumeData:completionHandler:"), resumeData, completionHandler)
	return rv
}

// Creates a task that establishes a bidirectional TCP/IP connection to a specified hostname and port. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411587-streamtaskwithhostname?language=objc
func (u_ URLSession) StreamTaskWithHostNamePort(hostname string, port int) URLSessionStreamTask {
	rv := objc.Call[URLSessionStreamTask](u_, objc.Sel("streamTaskWithHostName:port:"), hostname, port)
	return rv
}

// Creates a WebSocket task for the provided URL request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/3235750-websockettaskwithrequest?language=objc
func (u_ URLSession) WebSocketTaskWithRequest(request IURLRequest) URLSessionWebSocketTask {
	rv := objc.Call[URLSessionWebSocketTask](u_, objc.Sel("webSocketTaskWithRequest:"), objc.Ptr(request))
	return rv
}

// Creates a task that performs an HTTP request for uploading data based on the specified URL request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1410934-uploadtaskwithstreamedrequest?language=objc
func (u_ URLSession) UploadTaskWithStreamedRequest(request IURLRequest) URLSessionUploadTask {
	rv := objc.Call[URLSessionUploadTask](u_, objc.Sel("uploadTaskWithStreamedRequest:"), objc.Ptr(request))
	return rv
}

// Creates a task that retrieves the contents of a URL based on the specified URL request object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1410592-datataskwithrequest?language=objc
func (u_ URLSession) DataTaskWithRequest(request IURLRequest) URLSessionDataTask {
	rv := objc.Call[URLSessionDataTask](u_, objc.Sel("dataTaskWithRequest:"), objc.Ptr(request))
	return rv
}

// Creates a task that performs an HTTP request for uploading the specified file, then calls a handler upon completion. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411638-uploadtaskwithrequest?language=objc
func (u_ URLSession) UploadTaskWithRequestFromFileCompletionHandler(request IURLRequest, fileURL IURL, completionHandler func(data []byte, response URLResponse, error Error)) URLSessionUploadTask {
	rv := objc.Call[URLSessionUploadTask](u_, objc.Sel("uploadTaskWithRequest:fromFile:completionHandler:"), objc.Ptr(request), objc.Ptr(fileURL), completionHandler)
	return rv
}

// The shared singleton session object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1409000-sharedsession?language=objc
func (uc _URLSessionClass) SharedSession() URLSession {
	rv := objc.Call[URLSession](uc, objc.Sel("sharedSession"))
	return rv
}

// The shared singleton session object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1409000-sharedsession?language=objc
func URLSession_SharedSession() URLSession {
	return URLSessionClass.SharedSession()
}

// The delegate assigned when this object was created. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411530-delegate?language=objc
func (u_ URLSession) Delegate() URLSessionDelegateObject {
	rv := objc.Call[URLSessionDelegateObject](u_, objc.Sel("delegate"))
	return rv
}

// The operation queue provided when this object was created. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411571-delegatequeue?language=objc
func (u_ URLSession) DelegateQueue() OperationQueue {
	rv := objc.Call[OperationQueue](u_, objc.Sel("delegateQueue"))
	return rv
}

// An app-defined descriptive label for the session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1408277-sessiondescription?language=objc
func (u_ URLSession) SessionDescription() string {
	rv := objc.Call[string](u_, objc.Sel("sessionDescription"))
	return rv
}

// An app-defined descriptive label for the session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1408277-sessiondescription?language=objc
func (u_ URLSession) SetSessionDescription(value string) {
	objc.Call[objc.Void](u_, objc.Sel("setSessionDescription:"), value)
}

// A copy of the configuration object for this session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsession/1411477-configuration?language=objc
func (u_ URLSession) Configuration() URLSessionConfiguration {
	rv := objc.Call[URLSessionConfiguration](u_, objc.Sel("configuration"))
	return rv
}
