// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

// Package orders implements order parsing.
package orders

// Orders is a slice of Order.
type Orders []Order

// Order is the interface for all types of orders.
type Order interface {
	Id() int                             // id of unit being ordered
	AddError(format string, args ...any) // add a new error to the order
	Errors() []error                     // any errors parsing
	Execute() error                      // any error executing
	Line() int                           // line number in orders file
	String() string                      // implement the Stringer interface
}
