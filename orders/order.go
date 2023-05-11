// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

// Package orders implements order parsing.
package orders

// Orders is a slice of Order.
type Orders []Executer

// Executer is the interface for executing all types of orders.
type Executer interface {
	Execute() error
}
