package main

import (
	fmt "fmt"
	mat "math"
)

/********************************* package.go *********************************/

// PACKAGE CONSTANTS

// Public Constants

// This constant represents the value tau (τ).
// See "The Tau Manifesto" at https://tauday.com/tau-manifesto
const Tau = 2.0 * mat.Pi

// Private Constants

// These private constants implement the singleton pattern to provide a single
// reference to each class type structure.  They also initialize any class
// constants.
var (
	angleClassSingleton = &angleClass_{
		angle_(mat.Pi), // Angle.Pi()
		angle_(Tau),    // Angle.Tau()
	}
)

// PACKAGE ABSTRACTIONS

// Abstract Interfaces

// This abstract interface defines the set of method signatures that must be
// supported by all continuous numeric types.
type Continuous interface {
	AsFloat() float64
	IsZero() bool
}

// This abstract interface defines the set of method signatures that must be
// supported by all angular numeric types.
type Angular interface {
	AsDegrees() float64
	AsNormalized() float64
}

// Abstract Types

// This abstract type defines the set of abstract interfaces that must be
// supported by all angle-like types.
type AngleLike interface {
	Continuous
	Angular
}

// PACKAGE FUNCTIONS

// Public Functions

// This public function returns the circumference of a circle with the specified
// radius.
func Circumference(radius float64) float64 {
	var circumference = Tau * radius
	return circumference
}

// Private Functions

// This private function returns the magnitude (absolute value) of the specified
// value.
func magnitude(value float64) float64 {
	if value < 0 {
		return -value
	}
	return value
}

// PACKAGE CLASSES

// This function returns a reference to the angle class type singleton.
func Angle() *angleClass_ {
	return angleClassSingleton
}

/********************************** angle.go **********************************/

// CLASS TYPE

// This private class type defines the structure associated with the constants,
// constructors and functions for the angle class.
type angleClass_ struct {
	pi  AngleLike
	tau AngleLike
}

// CLASS CONSTANTS

// This class constant represents an angle with the value pi.
func (c *angleClass_) Pi() AngleLike {
	return c.pi
}

// This class constant represents an angle with the value tau.
func (c *angleClass_) Tau() AngleLike {
	return c.tau
}

// CLASS CONSTRUCTORS

// This class constructor creates a new angle from the specified float value.
func (c *angleClass_) FromFloat(float float64) AngleLike {
	var angle = angle_(float)
	return angle
}

// This class constructor creates a new angle from the specified string value.
func (c *angleClass_) FromString(string_ string) AngleLike {
	var angle AngleLike
	switch string_ {
	case "pi", "π":
		angle = c.pi
	case "tau", "τ":
		angle = c.tau
	default:
		var message = fmt.Sprintf("Attempted to construct an angle from an invalid string: %v", string_)
		panic(message)
	}
	return angle
}

// CLASS FUNCTIONS

// This class function returns the difference between the specified angles.
func (c *angleClass_) Difference(first, second AngleLike) AngleLike {
	return c.FromFloat(first.AsFloat() - second.AsFloat())
}

// CLASS METHODS

// This private concrete type extends the primitive Go float64 data type and
// defines the methods that implement the angle-like abstract type.  It
// represents a radian based angle.
type angle_ float64

// Continuous Interface

// This class method returns the floating point value for this angle.
func (v angle_) AsFloat() float64 {
	return float64(v)
}

// This class method whether or not the value for this angle is zero.
func (v angle_) IsZero() bool {
	return v == 0
}

// Angular Interface

// This class method returns the floating point value for this angle in degrees.
func (v angle_) AsDegrees() float64 {
	return float64(v * 360.0 / Tau)
}

// This class method returns this angle normalized to the range [0..τ).
func (v angle_) AsNormalized() float64 {
	var angle = v.AsFloat()
	if magnitude(angle) >= Tau {
		// Normalize the angle to the range (-τ..τ).
		angle = mat.Remainder(angle, Tau)
	}
	if angle < 0.0 {
		// Normalize the angle to the range [0..τ).
		angle = angle + Tau
	}
	if angle == -0 {
		// This is annoying so fix it.
		angle = 0
	}
	return angle
}

/******************************************************************************/

// USAGE EXAMPLE

func main() {
	// Retrieve the angle class type.
	var Angle = Angle()

	// Retrieve a class constant.
	var pi = Angle.Pi()

	// Call a class constructor.
	var angle = Angle.FromFloat(1.23)

	// Call a class function.
	var delta = Angle.Difference(angle, pi)

	// Call class methods.
	fmt.Printf("radians: %v\n", delta.AsFloat())
	fmt.Printf("degrees: %v\n", delta.AsDegrees())
	fmt.Printf("normalized: %v\n", delta.AsNormalized())
}

