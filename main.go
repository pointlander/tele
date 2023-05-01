// Copyright 2023 The Tele Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"math"
	"math/cmplx"

	"github.com/itsubaki/q"
	"github.com/itsubaki/q/math/matrix"
)

// SR send receive
func SR() {
	qsim := q.New()
	q0 := qsim.Zero()
	q1 := qsim.Zero()

	// Define the parameters
	k := 1.0
	h := 1.5

	// Prepare the ground state
	alpha := -math.Asin((1 / math.Sqrt(2)) * (math.Sqrt(1 + h/math.Sqrt(h*h+k*k))))

	qsim.RY(2*alpha, q0)
	qsim.CNOT(q0, q1)
	fmt.Println("2*alpha=", 2*alpha)

	qsim.H(q0)

	sin := func(k, h float64) float64 {
		a := h*h + 2*k*k
		b := h * k
		return b / math.Sqrt(a*a+b*b)
	}
	phi := 0.5 * math.Asin(sin(k, h))

	rotate := func(v float64) matrix.Matrix {
		return matrix.Matrix{
			[]complex128{cmplx.Cos(complex(v, 0)), -1 * cmplx.Sin(complex(v, 0))},
			[]complex128{cmplx.Sin(complex(v, 0)), cmplx.Cos(complex(v, 0))},
		}
	}

	qsim.C(rotate(-2*phi), q0, q1)
	qsim.I(q0)
	qsim.C(rotate(2*phi), q0, q1)
	fmt.Println("2*phi=", 2*phi)

	qsim.I(q0)
	qsim.H(q1)

	for _, s := range qsim.State() {
		fmt.Println(s)
	}
}

// R receive only
func R() {
	qsim := q.New()
	q0 := qsim.Zero()

	// Define the parameters
	k := 1.0
	h := 1.5

	// Prepare the ground state
	alpha := -math.Asin((1 / math.Sqrt(2)) * (math.Sqrt(1 + h/math.Sqrt(h*h+k*k))))

	qsim.RY(2*alpha, q0)
	qsim.I(q0)
	fmt.Println("2*alpha=", 2*alpha)

	sin := func(k, h float64) float64 {
		a := h*h + 2*k*k
		b := h * k
		return b / math.Sqrt(a*a+b*b)
	}
	phi := 0.5 * math.Asin(sin(k, h))

	qsim.RY(2*phi, q0)
	fmt.Println("2*phi=", 2*phi)

	qsim.H(q0)
	//qsim.Measure(q0)

	for _, s := range qsim.State() {
		fmt.Println(s)
	}
}

// Split split mode
func Split() {
	qsim := q.New()
	q0 := qsim.Zero()
	q1 := qsim.Zero()

	// Define the parameters
	k := 1.0
	h := 1.5

	// Prepare the ground state
	alpha := -math.Asin((1 / math.Sqrt(2)) * (math.Sqrt(1 + h/math.Sqrt(h*h+k*k))))

	qsim.RY(2*alpha, q0)
	qsim.I(q0)
	fmt.Println("2*alpha=", 2*alpha)

	sin := func(k, h float64) float64 {
		a := h*h + 2*k*k
		b := h * k
		return b / math.Sqrt(a*a+b*b)
	}
	phi := 0.5 * math.Asin(sin(k, h))

	qsim.RY(2*phi, q0)
	fmt.Println("2*phi=", 2*phi)

	qsim.H(q0)

	qsim.RY(2*alpha, q1)
	qsim.I(q1)
	qsim.RY(-2*phi, q1)
	qsim.H(q1)

	//qsim.Measure(q0, q1)

	for _, s := range qsim.State() {
		fmt.Println(s)
	}
}

func main() {
	fmt.Println("Split:")
	Split()
	fmt.Println("\nR:")
	R()
	fmt.Println("\nSR:")
	SR()
}
