package main

import (
	fmt "fmt"
	mat "math"
)

/********************************* package.go *********************************/

// PACKAGE CONSTANTS

// This package constant represents the value tau (τ).
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

// This abstract type defines the set of class constants, constructors and
// functions that must be supported by all angle-class-like types.
type AngleClassLike interface {
	// constants
	GetPi() AngleLike
	GetTau() AngleLike
	// constructors
	FromFloat(value float64) AngleLike
	FromString(value string) AngleLike
	// functions
	Difference(first, second AngleLike) AngleLike
}

// This abstract type defines the set of abstract interfaces that must be
// supported by all angle-like types.
type AngleLike interface {
	Angular
	Continuous
}

/********************************** angle.go **********************************/

// CLASS NAMESPACE

// Private Class Namespace Type

type angleClass_ struct {
	pi  AngleLike
	tau AngleLike
}

// Public Class Namespace Access

var angleClass = &angleClass_{
	pi:  angle_(mat.Pi), // Angle.Pi()
	tau: angle_(Tau),    // Angle.Tau()
}

func AngleClass() AngleClassLike {
	return angleClass
}

// Public Class Constants

func (c *angleClass_) GetPi() AngleLike {
	return c.pi
}

func (c *angleClass_) GetTau() AngleLike {
	return c.tau
}

// Public Class Constructors

func (c *angleClass_) FromFloat(float float64) AngleLike {
	var angle = angle_(float)
	return angle
}

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

// Public Class Functions

func (c *angleClass_) Difference(first, second AngleLike) AngleLike {
	var float = first.AsFloat() - second.AsFloat()
	var angle = c.FromFloat(float)
	return angle
}

// CLASS TYPE

// Private Class Type Definition

type angle_ float64

// Angular Interface

func (v angle_) AsDegrees() float64 {
	return float64(v * 360.0 / Tau)
}

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

// Continuous Interface

func (v angle_) AsFloat() float64 {
	return float64(v)
}

func (v angle_) IsZero() bool {
	return v == 0
}

/******************************************************************************/

// USAGE EXAMPLE

func main() {
	// Retrieve the angle class namespace.
	var Angle = AngleClass()

	// Retrieve a class constant.
	var pi = Angle.GetPi()

	// Call a class constructor.
	var angle = Angle.FromFloat(1.23)

	// Call a class function.
	var delta = Angle.Difference(angle, pi)

	// Call class methods.
	fmt.Printf("degrees: %v\n", delta.AsDegrees())
	fmt.Printf("radians: %v\n", delta.AsFloat())
	fmt.Printf("normalized: %v\n", delta.AsNormalized())
}
