// Code generated by DarwinKit. DO NOT EDIT.

package coregraphics

// Component information for a bitmap image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgbitmapinfo?language=objc
type BitmapInfo uint32

const (
	KBitmapAlphaInfoMask     BitmapInfo = 31
	KBitmapByteOrder16Big    BitmapInfo = 12288
	KBitmapByteOrder16Little BitmapInfo = 4096
	KBitmapByteOrder32Big    BitmapInfo = 16384
	KBitmapByteOrder32Little BitmapInfo = 8192
	KBitmapByteOrderDefault  BitmapInfo = 0
	KBitmapByteOrderMask     BitmapInfo = 28672
	KBitmapFloatComponents   BitmapInfo = 256
	KBitmapFloatInfoMask     BitmapInfo = 3840
)

// Compositing operations for images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgblendmode?language=objc
type BlendMode int32

const (
	KBlendModeClear           BlendMode = 16
	KBlendModeColor           BlendMode = 14
	KBlendModeColorBurn       BlendMode = 7
	KBlendModeColorDodge      BlendMode = 6
	KBlendModeCopy            BlendMode = 17
	KBlendModeDarken          BlendMode = 4
	KBlendModeDestinationAtop BlendMode = 24
	KBlendModeDestinationIn   BlendMode = 22
	KBlendModeDestinationOut  BlendMode = 23
	KBlendModeDestinationOver BlendMode = 21
	KBlendModeDifference      BlendMode = 10
	KBlendModeExclusion       BlendMode = 11
	KBlendModeHardLight       BlendMode = 9
	KBlendModeHue             BlendMode = 12
	KBlendModeLighten         BlendMode = 5
	KBlendModeLuminosity      BlendMode = 15
	KBlendModeMultiply        BlendMode = 1
	KBlendModeNormal          BlendMode = 0
	KBlendModeOverlay         BlendMode = 3
	KBlendModePlusDarker      BlendMode = 26
	KBlendModePlusLighter     BlendMode = 27
	KBlendModeSaturation      BlendMode = 13
	KBlendModeScreen          BlendMode = 2
	KBlendModeSoftLight       BlendMode = 8
	KBlendModeSourceAtop      BlendMode = 20
	KBlendModeSourceIn        BlendMode = 18
	KBlendModeSourceOut       BlendMode = 19
	KBlendModeXOR             BlendMode = 25
)

// Represents the number of buttons being set in a synthetic mouse event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgbuttoncount?language=objc
type ButtonCount uint32

// Configuration parameters that are used when capturing displays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcaptureoptions?language=objc
type CaptureOptions uint32

const (
	KCaptureNoFill    CaptureOptions = 1
	KCaptureNoOptions CaptureOptions = 0
)

// Represents a character generated by pressing one or more keys on a keyboard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcharcode?language=objc
type CharCode uint16

// Constants describing how a color conversion uses color spaces. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorconversioninfotransformtype?language=objc
type ColorConversionInfoTransformType uint32

const (
	KColorConversionTransformApplySpace ColorConversionInfoTransformType = 2
	KColorConversionTransformFromSpace  ColorConversionInfoTransformType = 0
	KColorConversionTransformToSpace    ColorConversionInfoTransformType = 1
)

// Handling options for colors that are not located within the destination color space of a graphics context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorrenderingintent?language=objc
type ColorRenderingIntent int32

const (
	KRenderingIntentAbsoluteColorimetric ColorRenderingIntent = 1
	KRenderingIntentDefault              ColorRenderingIntent = 0
	KRenderingIntentPerceptual           ColorRenderingIntent = 3
	KRenderingIntentRelativeColorimetric ColorRenderingIntent = 2
	KRenderingIntentSaturation           ColorRenderingIntent = 4
)

// Models for color spaces. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgcolorspacemodel?language=objc
type ColorSpaceModel int32

const (
	KColorSpaceModelCMYK       ColorSpaceModel = 2
	KColorSpaceModelDeviceN    ColorSpaceModel = 4
	KColorSpaceModelIndexed    ColorSpaceModel = 5
	KColorSpaceModelLab        ColorSpaceModel = 3
	KColorSpaceModelMonochrome ColorSpaceModel = 0
	KColorSpaceModelPattern    ColorSpaceModel = 6
	KColorSpaceModelRGB        ColorSpaceModel = 1
	KColorSpaceModelUnknown    ColorSpaceModel = -1
	KColorSpaceModelXYZ        ColorSpaceModel = 7
)

// The scope of the changes in a display configuration transaction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgconfigureoption?language=objc
type ConfigureOption uint32

const (
	KConfigureForAppOnly  ConfigureOption = 0
	KConfigureForSession  ConfigureOption = 1
	KConfigurePermanently ConfigureOption = 2
)

// A unique identifier for an attached display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdirectdisplayid?language=objc
type DirectDisplayID uint32

// The percentage of blend color used in a fade operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayblendfraction?language=objc
type DisplayBlendFraction float64

// The configuration parameters that are passed to a display reconfiguration callback function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaychangesummaryflags?language=objc
type DisplayChangeSummaryFlags uint32

const (
	KDisplayAddFlag                 DisplayChangeSummaryFlags = 16
	KDisplayBeginConfigurationFlag  DisplayChangeSummaryFlags = 1
	KDisplayDesktopShapeChangedFlag DisplayChangeSummaryFlags = 4096
	KDisplayDisabledFlag            DisplayChangeSummaryFlags = 512
	KDisplayEnabledFlag             DisplayChangeSummaryFlags = 256
	KDisplayMirrorFlag              DisplayChangeSummaryFlags = 1024
	KDisplayMovedFlag               DisplayChangeSummaryFlags = 2
	KDisplayRemoveFlag              DisplayChangeSummaryFlags = 32
	KDisplaySetMainFlag             DisplayChangeSummaryFlags = 4
	KDisplaySetModeFlag             DisplayChangeSummaryFlags = 8
	KDisplayUnMirrorFlag            DisplayChangeSummaryFlags = 2048
)

// The number of displays in various lists. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaycount?language=objc
type DisplayCount uint32

// The duration in seconds of a fade operation or a fade hardware reservation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayfadeinterval?language=objc
type DisplayFadeInterval float64

// A token issued by Quartz when reserving one or more displays for a fade operation during a specified interval. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayfadereservationtoken?language=objc
type DisplayFadeReservationToken uint32

// The time interval for a fade reservation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplayreservationinterval?language=objc
type DisplayReservationInterval float64

// Describes a frame update event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaystreamframestatus?language=objc
type DisplayStreamFrameStatus int32

const (
	KDisplayStreamFrameStatusFrameBlank    DisplayStreamFrameStatus = 2
	KDisplayStreamFrameStatusFrameComplete DisplayStreamFrameStatus = 0
	KDisplayStreamFrameStatusFrameIdle     DisplayStreamFrameStatus = 1
	KDisplayStreamFrameStatusStopped       DisplayStreamFrameStatus = 3
)

// Use these constants to determine which rectangles your app is interested in. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgdisplaystreamupdaterecttype?language=objc
type DisplayStreamUpdateRectType int32

const (
	KDisplayStreamUpdateDirtyRects        DisplayStreamUpdateRectType = 2
	KDisplayStreamUpdateMovedRects        DisplayStreamUpdateRectType = 1
	KDisplayStreamUpdateReducedDirtyRects DisplayStreamUpdateRectType = 3
	KDisplayStreamUpdateRefreshedRects    DisplayStreamUpdateRectType = 0
)

// A uniform type for result codes returned by functions in Core Graphics. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgerror?language=objc
type Error int32

const (
	KErrorCannotComplete    Error = 1004
	KErrorFailure           Error = 1000
	KErrorIllegalArgument   Error = 1001
	KErrorInvalidConnection Error = 1002
	KErrorInvalidContext    Error = 1003
	KErrorInvalidOperation  Error = 1010
	KErrorNoneAvailable     Error = 1011
	KErrorNotImplemented    Error = 1006
	KErrorRangeCheck        Error = 1007
	KErrorSuccess           Error = 0
	KErrorTypeCheck         Error = 1008
)

// Constants used as keys to access specialized fields in low-level events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventfield?language=objc
type EventField uint32

const (
	KEventSourceGroupID                                      EventField = 44
	KEventSourceStateID                                      EventField = 45
	KEventSourceUnixProcessID                                EventField = 41
	KEventSourceUserData                                     EventField = 42
	KEventSourceUserID                                       EventField = 43
	KEventTargetProcessSerialNumber                          EventField = 39
	KEventTargetUnixProcessID                                EventField = 40
	KEventUnacceleratedPointerMovementX                      EventField = 170
	KEventUnacceleratedPointerMovementY                      EventField = 171
	KKeyboardEventAutorepeat                                 EventField = 8
	KKeyboardEventKeyboardType                               EventField = 10
	KKeyboardEventKeycode                                    EventField = 9
	KMouseEventButtonNumber                                  EventField = 3
	KMouseEventClickState                                    EventField = 1
	KMouseEventDeltaX                                        EventField = 4
	KMouseEventDeltaY                                        EventField = 5
	KMouseEventInstantMouser                                 EventField = 6
	KMouseEventNumber                                        EventField = 0
	KMouseEventPressure                                      EventField = 2
	KMouseEventSubtype                                       EventField = 7
	KMouseEventWindowUnderMousePointer                       EventField = 91
	KMouseEventWindowUnderMousePointerThatCanHandleThisEvent EventField = 92
	KScrollWheelEventDeltaAxis1                              EventField = 11
	KScrollWheelEventDeltaAxis2                              EventField = 12
	KScrollWheelEventDeltaAxis3                              EventField = 13
	KScrollWheelEventFixedPtDeltaAxis1                       EventField = 93
	KScrollWheelEventFixedPtDeltaAxis2                       EventField = 94
	KScrollWheelEventFixedPtDeltaAxis3                       EventField = 95
	KScrollWheelEventInstantMouser                           EventField = 14
	KScrollWheelEventIsContinuous                            EventField = 88
	KScrollWheelEventMomentumPhase                           EventField = 123
	KScrollWheelEventPointDeltaAxis1                         EventField = 96
	KScrollWheelEventPointDeltaAxis2                         EventField = 97
	KScrollWheelEventPointDeltaAxis3                         EventField = 98
	KScrollWheelEventScrollCount                             EventField = 100
	KScrollWheelEventScrollPhase                             EventField = 99
	KTabletEventDeviceID                                     EventField = 24
	KTabletEventPointButtons                                 EventField = 18
	KTabletEventPointPressure                                EventField = 19
	KTabletEventPointX                                       EventField = 15
	KTabletEventPointY                                       EventField = 16
	KTabletEventPointZ                                       EventField = 17
	KTabletEventRotation                                     EventField = 22
	KTabletEventTangentialPressure                           EventField = 23
	KTabletEventTiltX                                        EventField = 20
	KTabletEventTiltY                                        EventField = 21
	KTabletEventVendor1                                      EventField = 25
	KTabletEventVendor2                                      EventField = 26
	KTabletEventVendor3                                      EventField = 27
	KTabletProximityEventCapabilityMask                      EventField = 36
	KTabletProximityEventDeviceID                            EventField = 31
	KTabletProximityEventEnterProximity                      EventField = 38
	KTabletProximityEventPointerID                           EventField = 30
	KTabletProximityEventPointerType                         EventField = 37
	KTabletProximityEventSystemTabletID                      EventField = 32
	KTabletProximityEventTabletID                            EventField = 29
	KTabletProximityEventVendorID                            EventField = 28
	KTabletProximityEventVendorPointerSerialNumber           EventField = 34
	KTabletProximityEventVendorPointerType                   EventField = 33
	KTabletProximityEventVendorUniqueID                      EventField = 35
)

// Specify masks for classes of low-level events that can be filtered during event suppression states. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventfiltermask?language=objc
type EventFilterMask uint32

const (
	KEventFilterMaskPermitLocalKeyboardEvents EventFilterMask = 2
	KEventFilterMaskPermitLocalMouseEvents    EventFilterMask = 1
	KEventFilterMaskPermitSystemDefinedEvents EventFilterMask = 4
)

// Constants that indicate the modifier key state at the time an event is created, as well as other event-related states. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventflags?language=objc
type EventFlags uint64

const (
	KEventFlagMaskAlphaShift   EventFlags = 65536
	KEventFlagMaskAlternate    EventFlags = 524288
	KEventFlagMaskCommand      EventFlags = 1048576
	KEventFlagMaskControl      EventFlags = 262144
	KEventFlagMaskHelp         EventFlags = 4194304
	KEventFlagMaskNonCoalesced EventFlags = 256
	KEventFlagMaskNumericPad   EventFlags = 2097152
	KEventFlagMaskSecondaryFn  EventFlags = 8388608
	KEventFlagMaskShift        EventFlags = 131072
)

// Defines a mask that identifies the set of Quartz events to be observed in an event tap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventmask?language=objc
type EventMask uint64

// Constants used with the kCGMouseEventSubtype event field. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventmousesubtype?language=objc
type EventMouseSubtype uint32

const (
	KEventMouseSubtypeDefault         EventMouseSubtype = 0
	KEventMouseSubtypeTabletPoint     EventMouseSubtype = 1
	KEventMouseSubtypeTabletProximity EventMouseSubtype = 2
)

// Defines a code that represents the type of keyboard used with a specified event source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventsourcekeyboardtype?language=objc
type EventSourceKeyboardType uint32

// Constants that specify the possible source states of an event source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventsourcestateid?language=objc
type EventSourceStateID int32

const (
	KEventSourceStateCombinedSessionState EventSourceStateID = 0
	KEventSourceStateHIDSystemState       EventSourceStateID = 1
	KEventSourceStatePrivate              EventSourceStateID = -1
)

// Specify the event suppression states that can occur after posting an event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventsuppressionstate?language=objc
type EventSuppressionState uint32

const (
	KEventSuppressionStateRemoteMouseDrag     EventSuppressionState = 1
	KEventSuppressionStateSuppressionInterval EventSuppressionState = 0
	KNumberOfEventSuppressionStates           EventSuppressionState = 2
)

// Constants that specify possible tapping points for events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventtaplocation?language=objc
type EventTapLocation uint32

const (
	KAnnotatedSessionEventTap EventTapLocation = 2
	KHIDEventTap              EventTapLocation = 0
	KSessionEventTap          EventTapLocation = 1
)

// Constants that specify whether a new event tap is an active filter or a passive listener. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventtapoptions?language=objc
type EventTapOptions uint32

const (
	KEventTapOptionDefault    EventTapOptions = 0
	KEventTapOptionListenOnly EventTapOptions = 1
)

// Constants that specify where a new event tap is inserted into the list of active event taps. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventtapplacement?language=objc
type EventTapPlacement uint32

const (
	KHeadInsertEventTap EventTapPlacement = 0
	KTailAppendEventTap EventTapPlacement = 1
)

// Defines the elapsed time in nanoseconds since startup that a Quartz event occurred. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventtimestamp?language=objc
type EventTimestamp uint64

// Constants that specify the different types of input events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgeventtype?language=objc
type EventType uint32

const (
	KEventFlagsChanged           EventType = 12
	KEventKeyDown                EventType = 10
	KEventKeyUp                  EventType = 11
	KEventLeftMouseDown          EventType = 1
	KEventLeftMouseDragged       EventType = 6
	KEventLeftMouseUp            EventType = 2
	KEventMouseMoved             EventType = 5
	KEventNull                   EventType = 0
	KEventOtherMouseDown         EventType = 25
	KEventOtherMouseDragged      EventType = 27
	KEventOtherMouseUp           EventType = 26
	KEventRightMouseDown         EventType = 3
	KEventRightMouseDragged      EventType = 7
	KEventRightMouseUp           EventType = 4
	KEventScrollWheel            EventType = 22
	KEventTabletPointer          EventType = 23
	KEventTabletProximity        EventType = 24
	KEventTapDisabledByTimeout   EventType = 4294967294
	KEventTapDisabledByUserInput EventType = 4294967295
)

// The basic type for all floating-point values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfloat?language=objc
type Float = float64

// An index into a font table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfontindex?language=objc
type FontIndex int

const (
	KFontIndexInvalid FontIndex = 65535
	KFontIndexMax     FontIndex = 65534
	KGlyphMax         FontIndex = 65534
)

// Possible formats for a PostScript font subset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgfontpostscriptformat?language=objc
type FontPostScriptFormat int32

const (
	KFontPostScriptFormatType1  FontPostScriptFormat = 1
	KFontPostScriptFormatType3  FontPostScriptFormat = 3
	KFontPostScriptFormatType42 FontPostScriptFormat = 42
)

// A value used to map a color generated in software to a color supported by the display hardware. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cggammavalue?language=objc
type GammaValue float64

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cggesturephase?language=objc
type GesturePhase uint32

const (
	KGesturePhaseBegan     GesturePhase = 1
	KGesturePhaseCancelled GesturePhase = 8
	KGesturePhaseChanged   GesturePhase = 2
	KGesturePhaseEnded     GesturePhase = 4
	KGesturePhaseMayBegin  GesturePhase = 128
	KGesturePhaseNone      GesturePhase = 0
)

// An index into the internal glyph table of a font. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgglyph?language=objc
type Glyph FontIndex

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgglyphdeprecatedenum?language=objc
type GlyphDeprecatedEnum int32

// Drawing locations for gradients. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cggradientdrawingoptions?language=objc
type GradientDrawingOptions uint32

const (
	KGradientDrawsAfterEndLocation    GradientDrawingOptions = 2
	KGradientDrawsBeforeStartLocation GradientDrawingOptions = 1
)

// Storage options for alpha component data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimagealphainfo?language=objc
type ImageAlphaInfo uint32

const (
	KImageAlphaFirst              ImageAlphaInfo = 4
	KImageAlphaLast               ImageAlphaInfo = 3
	KImageAlphaNone               ImageAlphaInfo = 0
	KImageAlphaNoneSkipFirst      ImageAlphaInfo = 6
	KImageAlphaNoneSkipLast       ImageAlphaInfo = 5
	KImageAlphaOnly               ImageAlphaInfo = 7
	KImageAlphaPremultipliedFirst ImageAlphaInfo = 2
	KImageAlphaPremultipliedLast  ImageAlphaInfo = 1
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimagebyteorderinfo?language=objc
type ImageByteOrderInfo uint32

const (
	KImageByteOrder16Big    ImageByteOrderInfo = 12288
	KImageByteOrder16Little ImageByteOrderInfo = 4096
	KImageByteOrder32Big    ImageByteOrderInfo = 16384
	KImageByteOrder32Little ImageByteOrderInfo = 8192
	KImageByteOrderDefault  ImageByteOrderInfo = 0
	KImageByteOrderMask     ImageByteOrderInfo = 28672
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgimagepixelformatinfo?language=objc
type ImagePixelFormatInfo uint32

const (
	KImagePixelFormatMask      ImagePixelFormatInfo = 983040
	KImagePixelFormatPacked    ImagePixelFormatInfo = 0
	KImagePixelFormatRGB101010 ImagePixelFormatInfo = 196608
	KImagePixelFormatRGB555    ImagePixelFormatInfo = 65536
	KImagePixelFormatRGB565    ImagePixelFormatInfo = 131072
	KImagePixelFormatRGBCIF10  ImagePixelFormatInfo = 262144
)

// Levels of interpolation quality for rendering an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cginterpolationquality?language=objc
type InterpolationQuality int32

const (
	KInterpolationDefault InterpolationQuality = 0
	KInterpolationHigh    InterpolationQuality = 3
	KInterpolationLow     InterpolationQuality = 2
	KInterpolationMedium  InterpolationQuality = 4
	KInterpolationNone    InterpolationQuality = 1
)

// Represents the virtual key codes used in keyboard events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgkeycode?language=objc
type KeyCode uint16

// Styles for rendering the endpoint of a stroked line. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cglinecap?language=objc
type LineCap int32

const (
	KLineCapButt   LineCap = 0
	KLineCapRound  LineCap = 1
	KLineCapSquare LineCap = 2
)

// Junction types for stroked lines. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cglinejoin?language=objc
type LineJoin int32

const (
	KLineJoinBevel LineJoin = 2
	KLineJoinMiter LineJoin = 0
	KLineJoinRound LineJoin = 1
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgmomentumscrollphase?language=objc
type MomentumScrollPhase uint32

const (
	KMomentumScrollPhaseBegin    MomentumScrollPhase = 1
	KMomentumScrollPhaseContinue MomentumScrollPhase = 2
	KMomentumScrollPhaseEnd      MomentumScrollPhase = 3
	KMomentumScrollPhaseNone     MomentumScrollPhase = 0
)

// Constants that specify buttons on a one, two, or three-button mouse. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgmousebutton?language=objc
type MouseButton uint32

const (
	KMouseButtonCenter MouseButton = 2
	KMouseButtonLeft   MouseButton = 0
	KMouseButtonRight  MouseButton = 1
)

// A bitmask used in OpenGL to specify a set of attached displays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgopengldisplaymask?language=objc
type OpenGLDisplayMask uint32

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfaccesspermissions?language=objc
type PDFAccessPermissions uint32

const (
	KPDFAllowsCommenting           PDFAccessPermissions = 64
	KPDFAllowsContentAccessibility PDFAccessPermissions = 32
	KPDFAllowsContentCopying       PDFAccessPermissions = 16
	KPDFAllowsDocumentAssembly     PDFAccessPermissions = 8
	KPDFAllowsDocumentChanges      PDFAccessPermissions = 4
	KPDFAllowsFormFieldEntry       PDFAccessPermissions = 128
	KPDFAllowsHighQualityPrinting  PDFAccessPermissions = 2
	KPDFAllowsLowQualityPrinting   PDFAccessPermissions = 1
)

// A PDF Boolean value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfboolean?language=objc
type PDFBoolean uint8

// Box types for a PDF page. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfbox?language=objc
type PDFBox int32

const (
	KPDFArtBox   PDFBox = 4
	KPDFBleedBox PDFBox = 2
	KPDFCropBox  PDFBox = 1
	KPDFMediaBox PDFBox = 0
	KPDFTrimBox  PDFBox = 3
)

// The encoding format of PDF data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfdataformat?language=objc
type PDFDataFormat int32

const (
	PDFDataFormatJPEG2000    PDFDataFormat = 2
	PDFDataFormatJPEGEncoded PDFDataFormat = 1
	PDFDataFormatRaw         PDFDataFormat = 0
)

// A PDF integer value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfinteger?language=objc
type PDFInteger int32

// Types of PDF object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfobjecttype?language=objc
type PDFObjectType int32

const (
	KPDFObjectTypeArray      PDFObjectType = 7
	KPDFObjectTypeBoolean    PDFObjectType = 2
	KPDFObjectTypeDictionary PDFObjectType = 8
	KPDFObjectTypeInteger    PDFObjectType = 3
	KPDFObjectTypeName       PDFObjectType = 5
	KPDFObjectTypeNull       PDFObjectType = 1
	KPDFObjectTypeReal       PDFObjectType = 4
	KPDFObjectTypeStream     PDFObjectType = 9
	KPDFObjectTypeString     PDFObjectType = 6
)

// A PDF real value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdfreal?language=objc
type PDFReal float64

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpdftagtype?language=objc
type PDFTagType int32

const (
	PDFTagTypeAnnotation         PDFTagType = 507
	PDFTagTypeArt                PDFTagType = 102
	PDFTagTypeBibliography       PDFTagType = 504
	PDFTagTypeBlockQuote         PDFTagType = 105
	PDFTagTypeCaption            PDFTagType = 106
	PDFTagTypeCode               PDFTagType = 505
	PDFTagTypeDiv                PDFTagType = 104
	PDFTagTypeDocument           PDFTagType = 100
	PDFTagTypeFigure             PDFTagType = 700
	PDFTagTypeForm               PDFTagType = 702
	PDFTagTypeFormula            PDFTagType = 701
	PDFTagTypeHeader             PDFTagType = 201
	PDFTagTypeHeader1            PDFTagType = 202
	PDFTagTypeHeader2            PDFTagType = 203
	PDFTagTypeHeader3            PDFTagType = 204
	PDFTagTypeHeader4            PDFTagType = 205
	PDFTagTypeHeader5            PDFTagType = 206
	PDFTagTypeHeader6            PDFTagType = 207
	PDFTagTypeIndex              PDFTagType = 109
	PDFTagTypeLabel              PDFTagType = 302
	PDFTagTypeLink               PDFTagType = 506
	PDFTagTypeList               PDFTagType = 300
	PDFTagTypeListBody           PDFTagType = 303
	PDFTagTypeListItem           PDFTagType = 301
	PDFTagTypeNonStructure       PDFTagType = 110
	PDFTagTypeNote               PDFTagType = 502
	PDFTagTypeParagraph          PDFTagType = 200
	PDFTagTypePart               PDFTagType = 101
	PDFTagTypePrivate            PDFTagType = 111
	PDFTagTypeQuote              PDFTagType = 501
	PDFTagTypeReference          PDFTagType = 503
	PDFTagTypeRuby               PDFTagType = 600
	PDFTagTypeRubyAnnotationText PDFTagType = 602
	PDFTagTypeRubyBaseText       PDFTagType = 601
	PDFTagTypeRubyPunctuation    PDFTagType = 603
	PDFTagTypeSection            PDFTagType = 103
	PDFTagTypeSpan               PDFTagType = 500
	PDFTagTypeTOC                PDFTagType = 107
	PDFTagTypeTOCI               PDFTagType = 108
	PDFTagTypeTable              PDFTagType = 400
	PDFTagTypeTableBody          PDFTagType = 405
	PDFTagTypeTableDataCell      PDFTagType = 403
	PDFTagTypeTableFooter        PDFTagType = 406
	PDFTagTypeTableHeader        PDFTagType = 404
	PDFTagTypeTableHeaderCell    PDFTagType = 402
	PDFTagTypeTableRow           PDFTagType = 401
	PDFTagTypeWarichu            PDFTagType = 604
	PDFTagTypeWarichuPunctiation PDFTagType = 606
	PDFTagTypeWarichuText        PDFTagType = 605
)

// Options for rendering a path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathdrawingmode?language=objc
type PathDrawingMode int32

const (
	KPathEOFill       PathDrawingMode = 1
	KPathEOFillStroke PathDrawingMode = 4
	KPathFill         PathDrawingMode = 0
	KPathFillStroke   PathDrawingMode = 3
	KPathStroke       PathDrawingMode = 2
)

// The type of element found in a path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpathelementtype?language=objc
type PathElementType int32

const (
	KPathElementAddCurveToPoint     PathElementType = 3
	KPathElementAddLineToPoint      PathElementType = 1
	KPathElementAddQuadCurveToPoint PathElementType = 2
	KPathElementCloseSubpath        PathElementType = 4
	KPathElementMoveToPoint         PathElementType = 0
)

// Different methods for rendering a tiled pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgpatterntiling?language=objc
type PatternTiling int32

const (
	KPatternTilingConstantSpacing                  PatternTiling = 2
	KPatternTilingConstantSpacingMinimalDistortion PatternTiling = 1
	KPatternTilingNoDistortion                     PatternTiling = 0
)

// The size of an array of Quartz rectangles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrectcount?language=objc
type RectCount uint32

// Coordinates that establish the edges of a rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrectedge?language=objc
type RectEdge uint32

const (
	RectMaxXEdge RectEdge = 2
	RectMaxYEdge RectEdge = 3
	RectMinXEdge RectEdge = 0
	RectMinYEdge RectEdge = 1
)

// A display’s refresh rate in frames per second. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgrefreshrate?language=objc
type RefreshRate float64

// Types of screen-update operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgscreenupdateoperation?language=objc
type ScreenUpdateOperation uint32

const (
	KScreenUpdateOperationMove                       ScreenUpdateOperation = 1
	KScreenUpdateOperationReducedDirtyRectangleCount ScreenUpdateOperation = 2147483648
	KScreenUpdateOperationRefresh                    ScreenUpdateOperation = 0
)

// Constants that specify the unit of measurement for a scrolling event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgscrolleventunit?language=objc
type ScrollEventUnit uint32

const (
	KScrollEventUnitLine  ScrollEventUnit = 1
	KScrollEventUnitPixel ScrollEventUnit = 0
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgscrollphase?language=objc
type ScrollPhase uint32

const (
	KScrollPhaseBegan     ScrollPhase = 1
	KScrollPhaseCancelled ScrollPhase = 8
	KScrollPhaseChanged   ScrollPhase = 2
	KScrollPhaseEnded     ScrollPhase = 4
	KScrollPhaseMayBegin  ScrollPhase = 128
)

// Modes for rendering text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgtextdrawingmode?language=objc
type TextDrawingMode int32

const (
	KTextClip           TextDrawingMode = 7
	KTextFill           TextDrawingMode = 0
	KTextFillClip       TextDrawingMode = 4
	KTextFillStroke     TextDrawingMode = 2
	KTextFillStrokeClip TextDrawingMode = 6
	KTextInvisible      TextDrawingMode = 3
	KTextStroke         TextDrawingMode = 1
	KTextStrokeClip     TextDrawingMode = 5
)

// Text encodings for fonts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgtextencoding?language=objc
type TextEncoding int32

const (
	KEncodingFontSpecific TextEncoding = 0
	KEncodingMacRoman     TextEncoding = 1
)

// Represents the number of wheels being set in a scroll wheel event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgwheelcount?language=objc
type WheelCount uint32

// The data type used to specify the backing option for a given window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgwindowbackingtype?language=objc
type WindowBackingType uint32

const (
	KBackingStoreBuffered    WindowBackingType = 2
	KBackingStoreNonretained WindowBackingType = 1
	KBackingStoreRetained    WindowBackingType = 0
)

// The data type used to store window identifiers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgwindowid?language=objc
type WindowID uint32

// The data type to use to specify the type of image to be generated for a window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgwindowimageoption?language=objc
type WindowImageOption uint32

const (
	KWindowImageBestResolution      WindowImageOption = 8
	KWindowImageBoundsIgnoreFraming WindowImageOption = 1
	KWindowImageDefault             WindowImageOption = 0
	KWindowImageNominalResolution   WindowImageOption = 16
	KWindowImageOnlyShadows         WindowImageOption = 4
	KWindowImageShouldBeOpaque      WindowImageOption = 2
)

// A level assigned to a window by an application framework. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgwindowlevel?language=objc
type WindowLevel int32

// Keys that represent the standard window levels in macOS. Quartz includes these keys to support application frameworks like Cocoa. Applications do not need to use them directly. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgwindowlevelkey?language=objc
type WindowLevelKey int32

const (
	KAssistiveTechHighWindowLevelKey WindowLevelKey = 20
	KBackstopMenuLevelKey            WindowLevelKey = 3
	KBaseWindowLevelKey              WindowLevelKey = 0
	KCursorWindowLevelKey            WindowLevelKey = 19
	KDesktopIconWindowLevelKey       WindowLevelKey = 18
	KDesktopWindowLevelKey           WindowLevelKey = 2
	KDockWindowLevelKey              WindowLevelKey = 7
	KDraggingWindowLevelKey          WindowLevelKey = 12
	KFloatingWindowLevelKey          WindowLevelKey = 5
	KHelpWindowLevelKey              WindowLevelKey = 16
	KMainMenuWindowLevelKey          WindowLevelKey = 8
	KMaximumWindowLevelKey           WindowLevelKey = 14
	KMinimumWindowLevelKey           WindowLevelKey = 1
	KModalPanelWindowLevelKey        WindowLevelKey = 10
	KNormalWindowLevelKey            WindowLevelKey = 4
	KNumberOfWindowLevelKeys         WindowLevelKey = 21
	KOverlayWindowLevelKey           WindowLevelKey = 15
	KPopUpMenuWindowLevelKey         WindowLevelKey = 11
	KScreenSaverWindowLevelKey       WindowLevelKey = 13
	KStatusWindowLevelKey            WindowLevelKey = 9
	KTornOffMenuWindowLevelKey       WindowLevelKey = 6
	KUtilityWindowLevelKey           WindowLevelKey = 17
)

// The data type used to specify the options for gathering a list of windows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgwindowlistoption?language=objc
type WindowListOption uint32

const (
	KWindowListExcludeDesktopElements    WindowListOption = 16
	KWindowListOptionAll                 WindowListOption = 0
	KWindowListOptionIncludingWindow     WindowListOption = 8
	KWindowListOptionOnScreenAboveWindow WindowListOption = 2
	KWindowListOptionOnScreenBelowWindow WindowListOption = 4
	KWindowListOptionOnScreenOnly        WindowListOption = 1
)

// The data type used to specify the sharing mode used by a window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coregraphics/cgwindowsharingtype?language=objc
type WindowSharingType uint32

const (
	KWindowSharingNone      WindowSharingType = 0
	KWindowSharingReadOnly  WindowSharingType = 1
	KWindowSharingReadWrite WindowSharingType = 2
)
