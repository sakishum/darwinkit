// Code generated by DarwinKit. DO NOT EDIT.

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/coreimage"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/corevideo"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/imageio"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TargetedImageRequest] class.
var TargetedImageRequestClass = _TargetedImageRequestClass{objc.GetClass("VNTargetedImageRequest")}

type _TargetedImageRequestClass struct {
	objc.Class
}

// An interface definition for the [TargetedImageRequest] class.
type ITargetedImageRequest interface {
	IImageBasedRequest
}

// The abstract superclass for image analysis requests that operate on both the processed image and a secondary image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest?language=objc
type TargetedImageRequest struct {
	ImageBasedRequest
}

func TargetedImageRequestFrom(ptr unsafe.Pointer) TargetedImageRequest {
	return TargetedImageRequest{
		ImageBasedRequest: ImageBasedRequestFrom(ptr),
	}
}

func (t_ TargetedImageRequest) InitWithTargetedCGImageOrientationOptions(cgImage coregraphics.ImageRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedCGImage:orientation:options:"), cgImage, orientation, options)
	return rv
}

// Creates a new request targeting a Core Graphics image of known orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923444-initwithtargetedcgimage?language=objc
func NewTargetedImageRequestWithTargetedCGImageOrientationOptions(cgImage coregraphics.ImageRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedCGImageOrientationOptions(cgImage, orientation, options)
	instance.Autorelease()
	return instance
}

func (t_ TargetedImageRequest) InitWithTargetedCGImageOptions(cgImage coregraphics.ImageRef, options map[ImageOption]objc.IObject) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedCGImage:options:"), cgImage, options)
	return rv
}

// Creates a new request targeting a Core Graphics image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923452-initwithtargetedcgimage?language=objc
func NewTargetedImageRequestWithTargetedCGImageOptions(cgImage coregraphics.ImageRef, options map[ImageOption]objc.IObject) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedCGImageOptions(cgImage, options)
	instance.Autorelease()
	return instance
}

func (t_ TargetedImageRequest) InitWithTargetedImageURLOptionsCompletionHandler(imageURL foundation.IURL, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedImageURL:options:completionHandler:"), objc.Ptr(imageURL), options, completionHandler)
	return rv
}

// Creates a new request targeting an image at the specified URL, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923453-initwithtargetedimageurl?language=objc
func NewTargetedImageRequestWithTargetedImageURLOptionsCompletionHandler(imageURL foundation.IURL, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedImageURLOptionsCompletionHandler(imageURL, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (t_ TargetedImageRequest) InitWithTargetedImageDataOptions(imageData []byte, options map[ImageOption]objc.IObject) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedImageData:options:"), imageData, options)
	return rv
}

// Creates a new request targeting an image as raw data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923460-initwithtargetedimagedata?language=objc
func NewTargetedImageRequestWithTargetedImageDataOptions(imageData []byte, options map[ImageOption]objc.IObject) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedImageDataOptions(imageData, options)
	instance.Autorelease()
	return instance
}

func (t_ TargetedImageRequest) InitWithTargetedCVPixelBufferOptionsCompletionHandler(pixelBuffer corevideo.PixelBufferRef, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedCVPixelBuffer:options:completionHandler:"), pixelBuffer, options, completionHandler)
	return rv
}

// Creates a new request targeting an image in a CVPixelBufferRef, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923446-initwithtargetedcvpixelbuffer?language=objc
func NewTargetedImageRequestWithTargetedCVPixelBufferOptionsCompletionHandler(pixelBuffer corevideo.PixelBufferRef, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedCVPixelBufferOptionsCompletionHandler(pixelBuffer, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (t_ TargetedImageRequest) InitWithTargetedCMSampleBufferOptionsCompletionHandler(sampleBuffer coremedia.SampleBufferRef, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedCMSampleBuffer:options:completionHandler:"), sampleBuffer, options, completionHandler)
	return rv
}

// Creates a new request with a completion handler that targets an image in a sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/3571275-initwithtargetedcmsamplebuffer?language=objc
func NewTargetedImageRequestWithTargetedCMSampleBufferOptionsCompletionHandler(sampleBuffer coremedia.SampleBufferRef, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedCMSampleBufferOptionsCompletionHandler(sampleBuffer, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (t_ TargetedImageRequest) InitWithTargetedImageDataOrientationOptionsCompletionHandler(imageData []byte, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedImageData:orientation:options:completionHandler:"), imageData, orientation, options, completionHandler)
	return rv
}

// Creates a new request targeting a raw data image of known orientation, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923443-initwithtargetedimagedata?language=objc
func NewTargetedImageRequestWithTargetedImageDataOrientationOptionsCompletionHandler(imageData []byte, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedImageDataOrientationOptionsCompletionHandler(imageData, orientation, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (t_ TargetedImageRequest) InitWithTargetedCMSampleBufferOrientationOptions(sampleBuffer coremedia.SampleBufferRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedCMSampleBuffer:orientation:options:"), sampleBuffer, orientation, options)
	return rv
}

// Creates a new request that targets an image of a known orientation in a sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/3571276-initwithtargetedcmsamplebuffer?language=objc
func NewTargetedImageRequestWithTargetedCMSampleBufferOrientationOptions(sampleBuffer coremedia.SampleBufferRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedCMSampleBufferOrientationOptions(sampleBuffer, orientation, options)
	instance.Autorelease()
	return instance
}

func (t_ TargetedImageRequest) InitWithTargetedCVPixelBufferOptions(pixelBuffer corevideo.PixelBufferRef, options map[ImageOption]objc.IObject) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedCVPixelBuffer:options:"), pixelBuffer, options)
	return rv
}

// Creates a new request targeting an image in a CVPixelBufferRef. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923445-initwithtargetedcvpixelbuffer?language=objc
func NewTargetedImageRequestWithTargetedCVPixelBufferOptions(pixelBuffer corevideo.PixelBufferRef, options map[ImageOption]objc.IObject) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedCVPixelBufferOptions(pixelBuffer, options)
	instance.Autorelease()
	return instance
}

func (t_ TargetedImageRequest) InitWithTargetedImageDataOptionsCompletionHandler(imageData []byte, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedImageData:options:completionHandler:"), imageData, options, completionHandler)
	return rv
}

// Creates a new request targeting an image as raw data, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923455-initwithtargetedimagedata?language=objc
func NewTargetedImageRequestWithTargetedImageDataOptionsCompletionHandler(imageData []byte, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedImageDataOptionsCompletionHandler(imageData, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (t_ TargetedImageRequest) InitWithTargetedCIImageOptionsCompletionHandler(ciImage coreimage.IImage, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedCIImage:options:completionHandler:"), objc.Ptr(ciImage), options, completionHandler)
	return rv
}

// Creates a new request targeting a CIImage, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923454-initwithtargetedciimage?language=objc
func NewTargetedImageRequestWithTargetedCIImageOptionsCompletionHandler(ciImage coreimage.IImage, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedCIImageOptionsCompletionHandler(ciImage, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (t_ TargetedImageRequest) InitWithTargetedCIImageOrientationOptions(ciImage coreimage.IImage, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedCIImage:orientation:options:"), objc.Ptr(ciImage), orientation, options)
	return rv
}

// Creates a new request targeting a CIImage of known orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923459-initwithtargetedciimage?language=objc
func NewTargetedImageRequestWithTargetedCIImageOrientationOptions(ciImage coreimage.IImage, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedCIImageOrientationOptions(ciImage, orientation, options)
	instance.Autorelease()
	return instance
}

func (t_ TargetedImageRequest) InitWithTargetedImageURLOrientationOptions(imageURL foundation.IURL, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedImageURL:orientation:options:"), objc.Ptr(imageURL), orientation, options)
	return rv
}

// Creates a new request targeting an image of known orientation, at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923456-initwithtargetedimageurl?language=objc
func NewTargetedImageRequestWithTargetedImageURLOrientationOptions(imageURL foundation.IURL, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedImageURLOrientationOptions(imageURL, orientation, options)
	instance.Autorelease()
	return instance
}

func (t_ TargetedImageRequest) InitWithTargetedCIImageOrientationOptionsCompletionHandler(ciImage coreimage.IImage, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedCIImage:orientation:options:completionHandler:"), objc.Ptr(ciImage), orientation, options, completionHandler)
	return rv
}

// Creates a new request targeting a CIImage of known orientation, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923451-initwithtargetedciimage?language=objc
func NewTargetedImageRequestWithTargetedCIImageOrientationOptionsCompletionHandler(ciImage coreimage.IImage, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedCIImageOrientationOptionsCompletionHandler(ciImage, orientation, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (t_ TargetedImageRequest) InitWithTargetedCGImageOptionsCompletionHandler(cgImage coregraphics.ImageRef, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedCGImage:options:completionHandler:"), cgImage, options, completionHandler)
	return rv
}

// Creates a new request targeting a Core Graphics image, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923448-initwithtargetedcgimage?language=objc
func NewTargetedImageRequestWithTargetedCGImageOptionsCompletionHandler(cgImage coregraphics.ImageRef, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedCGImageOptionsCompletionHandler(cgImage, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (t_ TargetedImageRequest) InitWithTargetedImageDataOrientationOptions(imageData []byte, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedImageData:orientation:options:"), imageData, orientation, options)
	return rv
}

// Creates a new request targeting a raw data image of known orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923441-initwithtargetedimagedata?language=objc
func NewTargetedImageRequestWithTargetedImageDataOrientationOptions(imageData []byte, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedImageDataOrientationOptions(imageData, orientation, options)
	instance.Autorelease()
	return instance
}

func (t_ TargetedImageRequest) InitWithTargetedCMSampleBufferOptions(sampleBuffer coremedia.SampleBufferRef, options map[ImageOption]objc.IObject) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedCMSampleBuffer:options:"), sampleBuffer, options)
	return rv
}

// Creates a new request that targets an image in a sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/3571274-initwithtargetedcmsamplebuffer?language=objc
func NewTargetedImageRequestWithTargetedCMSampleBufferOptions(sampleBuffer coremedia.SampleBufferRef, options map[ImageOption]objc.IObject) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedCMSampleBufferOptions(sampleBuffer, options)
	instance.Autorelease()
	return instance
}

func (t_ TargetedImageRequest) InitWithTargetedImageURLOrientationOptionsCompletionHandler(imageURL foundation.IURL, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedImageURL:orientation:options:completionHandler:"), objc.Ptr(imageURL), orientation, options, completionHandler)
	return rv
}

// Creates a new request targeting an image of known orientation, at the specified URL, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923457-initwithtargetedimageurl?language=objc
func NewTargetedImageRequestWithTargetedImageURLOrientationOptionsCompletionHandler(imageURL foundation.IURL, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedImageURLOrientationOptionsCompletionHandler(imageURL, orientation, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (t_ TargetedImageRequest) InitWithTargetedCMSampleBufferOrientationOptionsCompletionHandler(sampleBuffer coremedia.SampleBufferRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedCMSampleBuffer:orientation:options:completionHandler:"), sampleBuffer, orientation, options, completionHandler)
	return rv
}

// Creates a new request with a completion handler that targets an image of a known orientation in a sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/3571277-initwithtargetedcmsamplebuffer?language=objc
func NewTargetedImageRequestWithTargetedCMSampleBufferOrientationOptionsCompletionHandler(sampleBuffer coremedia.SampleBufferRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedCMSampleBufferOrientationOptionsCompletionHandler(sampleBuffer, orientation, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (t_ TargetedImageRequest) InitWithTargetedCIImageOptions(ciImage coreimage.IImage, options map[ImageOption]objc.IObject) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedCIImage:options:"), objc.Ptr(ciImage), options)
	return rv
}

// Creates a new request targeting a CIImage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923447-initwithtargetedciimage?language=objc
func NewTargetedImageRequestWithTargetedCIImageOptions(ciImage coreimage.IImage, options map[ImageOption]objc.IObject) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedCIImageOptions(ciImage, options)
	instance.Autorelease()
	return instance
}

func (t_ TargetedImageRequest) InitWithTargetedCGImageOrientationOptionsCompletionHandler(cgImage coregraphics.ImageRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedCGImage:orientation:options:completionHandler:"), cgImage, orientation, options, completionHandler)
	return rv
}

// Creates a new request targeting a Core Graphics image of known orientation, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923450-initwithtargetedcgimage?language=objc
func NewTargetedImageRequestWithTargetedCGImageOrientationOptionsCompletionHandler(cgImage coregraphics.ImageRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedCGImageOrientationOptionsCompletionHandler(cgImage, orientation, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (t_ TargetedImageRequest) InitWithTargetedCVPixelBufferOrientationOptions(pixelBuffer corevideo.PixelBufferRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedCVPixelBuffer:orientation:options:"), pixelBuffer, orientation, options)
	return rv
}

// Creates a new request targeting an image in a CVPixelBufferRef of known orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923442-initwithtargetedcvpixelbuffer?language=objc
func NewTargetedImageRequestWithTargetedCVPixelBufferOrientationOptions(pixelBuffer corevideo.PixelBufferRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedCVPixelBufferOrientationOptions(pixelBuffer, orientation, options)
	instance.Autorelease()
	return instance
}

func (t_ TargetedImageRequest) InitWithTargetedCVPixelBufferOrientationOptionsCompletionHandler(pixelBuffer corevideo.PixelBufferRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedCVPixelBuffer:orientation:options:completionHandler:"), pixelBuffer, orientation, options, completionHandler)
	return rv
}

// Creates a new request targeting an image in a CVPixelBufferRef of known orientation, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923449-initwithtargetedcvpixelbuffer?language=objc
func NewTargetedImageRequestWithTargetedCVPixelBufferOrientationOptionsCompletionHandler(pixelBuffer corevideo.PixelBufferRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedCVPixelBufferOrientationOptionsCompletionHandler(pixelBuffer, orientation, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (t_ TargetedImageRequest) InitWithTargetedImageURLOptions(imageURL foundation.IURL, options map[ImageOption]objc.IObject) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithTargetedImageURL:options:"), objc.Ptr(imageURL), options)
	return rv
}

// Creates a new request targeting an image at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923458-initwithtargetedimageurl?language=objc
func NewTargetedImageRequestWithTargetedImageURLOptions(imageURL foundation.IURL, options map[ImageOption]objc.IObject) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithTargetedImageURLOptions(imageURL, options)
	instance.Autorelease()
	return instance
}

func (tc _TargetedImageRequestClass) Alloc() TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](tc, objc.Sel("alloc"))
	return rv
}

func (tc _TargetedImageRequestClass) New() TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTargetedImageRequest() TargetedImageRequest {
	return TargetedImageRequestClass.New()
}

func (t_ TargetedImageRequest) Init() TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("init"))
	return rv
}

func (t_ TargetedImageRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) TargetedImageRequest {
	rv := objc.Call[TargetedImageRequest](t_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewTargetedImageRequestWithCompletionHandler(completionHandler RequestCompletionHandler) TargetedImageRequest {
	instance := TargetedImageRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}
