// Code generated by DarwinKit. DO NOT EDIT.

package vision

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TrajectoryObservation] class.
var TrajectoryObservationClass = _TrajectoryObservationClass{objc.GetClass("VNTrajectoryObservation")}

type _TrajectoryObservationClass struct {
	objc.Class
}

// An interface definition for the [TrajectoryObservation] class.
type ITrajectoryObservation interface {
	IObservation
	EquationCoefficients() objc.Object
	DetectedPoints() []Point
	MovingAverageRadius() float64
	ProjectedPoints() []Point
}

// An observation that describes a detected trajectory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrajectoryobservation?language=objc
type TrajectoryObservation struct {
	Observation
}

func TrajectoryObservationFrom(ptr unsafe.Pointer) TrajectoryObservation {
	return TrajectoryObservation{
		Observation: ObservationFrom(ptr),
	}
}

func (tc _TrajectoryObservationClass) Alloc() TrajectoryObservation {
	rv := objc.Call[TrajectoryObservation](tc, objc.Sel("alloc"))
	return rv
}

func (tc _TrajectoryObservationClass) New() TrajectoryObservation {
	rv := objc.Call[TrajectoryObservation](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTrajectoryObservation() TrajectoryObservation {
	return TrajectoryObservationClass.New()
}

func (t_ TrajectoryObservation) Init() TrajectoryObservation {
	rv := objc.Call[TrajectoryObservation](t_, objc.Sel("init"))
	return rv
}

// The coefficients of the parabolic equation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrajectoryobservation/3564825-equationcoefficients?language=objc
func (t_ TrajectoryObservation) EquationCoefficients() objc.Object {
	rv := objc.Call[objc.Object](t_, objc.Sel("equationCoefficients"))
	return rv
}

// The centroid points of the detected contour along the trajectory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrajectoryobservation/3564824-detectedpoints?language=objc
func (t_ TrajectoryObservation) DetectedPoints() []Point {
	rv := objc.Call[[]Point](t_, objc.Sel("detectedPoints"))
	return rv
}

// The moving average radius of the object the request is tracking. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrajectoryobservation/3751001-movingaverageradius?language=objc
func (t_ TrajectoryObservation) MovingAverageRadius() float64 {
	rv := objc.Call[float64](t_, objc.Sel("movingAverageRadius"))
	return rv
}

// The centroids of the calculated trajectory from the detected points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrajectoryobservation/3564826-projectedpoints?language=objc
func (t_ TrajectoryObservation) ProjectedPoints() []Point {
	rv := objc.Call[[]Point](t_, objc.Sel("projectedPoints"))
	return rv
}
