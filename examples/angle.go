package main

import (
	fmt "fmt"
	mat "math"
)

/********************************* package.go *********************************/

// PACKAGE CONSTANTS

// This constant represents the value tau (τ).
// See "The Tau Manifesto" at https://tauday.com/tau-manifesto
const Tau = 2.0 * mat.Pi

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

// This function returns the circumference of a circle with the specified
// radius.
func Circumference(radius float64) float64 {
	var circumference = Tau * radius
	return circumference
}

/********************************** angle.go **********************************/

// CLASS NAMESPACE

// This private type defines the namepace structure associated with the
// constants, constructors and functions for the angle class namespace.
type angleClass_ struct {
	pi  AngleLike
	tau AngleLike
}

// This private constant defines the singleton reference to the angle
// class namespace.  It also initializes any class constants as needed.
var angleClassSingleton = &angleClass_{
	angle_(mat.Pi), // Angle.Pi()
	angle_(Tau),    // Angle.Tau()
}

// This public function returns the singleton reference to the angle
// class namespace.
func Angle() *angleClass_ {
	return angleClassSingleton
}

// CLASS CONSTANTS

// This public class constant represents an angle with the value pi.
func (c *angleClass_) Pi() AngleLike {
	return c.pi
}

// This public class constant represents an angle with the value tau.
func (c *angleClass_) Tau() AngleLike {
	return c.tau
}

// CLASS CONSTRUCTORS

// This public class constructor creates a new angle from the specified float
// value.
func (c *angleClass_) FromFloat(float float64) AngleLike {
	var angle = angle_(float)
	return angle
}

// This public class constructor creates a new angle from the specified string
// value.
func (c *angleClass_) FromString(string_ string) AngleLike {
	var angle AngleLike
	switch string_ {
	case "pi", "π":
		angle = c.pi
	case "tau", "τ":
		angle = c.tau
	default:
		var message = fmt.Sprintf(
			"Attempted to construct an angle from an invalid string: %v",
			string_,
		)
		panic(message)
	}
	return angle
}

// CLASS FUNCTIONS

// This public class function returns the difference between the specified
// angles.
func (c *angleClass_) Difference(first, second AngleLike) AngleLike {
	var float = first.AsFloat() - second.AsFloat()
	var angle = c.FromFloat(float)
	return angle
}

// CLASS TYPE

// This private class type extends the primitive Go float64 data type
// and defines the methods that implement the angle-like abstract type.
// It represents a radian based angle.
type angle_ float64

// Continuous Interface

// This public class method returns the floating point value for this angle.
func (v angle_) AsFloat() float64 {
	return float64(v)
}

// This public class method whether or not the value for this angle is zero.
func (v angle_) IsZero() bool {
	return v == 0
}

// Angular Interface

// This public class method returns the floating point value for this angle in
// degrees.
func (v angle_) AsDegrees() float64 {
	return float64(v * 360.0 / Tau)
}

// This public class method returns this angle normalized to the range [0..τ).
func (v angle_) AsNormalized() float64 {
	var angle = v.AsFloat()
	if angle <= -Tau || angle >= Tau {
		// Normalize the angle to the range (-τ..τ).
		angle = mat.Remainder(angle, Tau)
	}
	if angle == -0 {
		// This is annoying so fix it.
		angle = 0
	}
	if angle < 0.0 {
		// Normalize the angle to the range [0..τ).
		angle = angle + Tau
	}
	return angle
}

/******************************************************************************/

// USAGE EXAMPLE

func main() {
	// Retrieve the angle class namespace.
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

