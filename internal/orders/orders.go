// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package orders

// Order is the interface for all types of orders.
type Order interface {
	Execute() error
}

type Assemble struct {
	Line        int
	Id          int // id of unit being ordered
	DepositId   int // optional id of deposit to assemble mines at
	FactGroup   int // optional id of factory group to add units to
	MiningGroup int // optional id of mining group to add units to
	Quantity    int // number of units to assemble
	Product     int // product to assemble
	Research    int // research to assemble
}

// Execute implements the Order interface.
func (o *Assemble) Execute() error { panic("!") }

type AssembleFactoryGroup struct {
	Line        int
	Id          int // id of unit being ordered
	Group       int // optional id of group to add units to
	Quantity    int // number of units to assemble
	Units       int // factory units to assemble
	Manufacture int // product unit to be manufactured
}

type AssembleMineGroup struct {
	Line      int
	Id        int // id of unit being ordered
	DepositId int // deposit to assemble mines at
	Quantity  int // number of units to assemble
	Units     int // mine units to assemble
}

type AssembleProduct struct {
	Line     int
	Id       int // id of unit being ordered
	Quantity int // number of units to assemble
	Unit     int // unit to assemble
}

type AssembleResearchGroup struct {
	Line     int
	Id       int // id of unit being ordered
	Group    int // optional id of group to add units to
	Quantity int // number of units to assemble
	Units    int // factory units to assemble
}

type Bombard struct {
	Line         int
	Id           int // id of unit being ordered
	TargetId     int // id of unit being attacked
	PctCommitted int
	Errors       []error
}

// Execute implements the Order interface.
func (o *Bombard) Execute() error { panic("!") }

type Coordinates struct { // location being set up
	X, Y, Z int
	Orbit   int
}

type Invade struct {
	Line         int
	Id           int // id of unit being ordered
	TargetId     int // id of unit being attacked
	PctCommitted int
	Errors       []error
}

// Execute implements the Order interface.
func (o *Invade) Execute() error { panic("!") }

type Raid struct {
	Line         int
	Id           int // id of unit being ordered
	TargetId     int // id of unit being attacked
	PctCommitted int
	Material     string // material to raid
	Errors       []error
}

// Execute implements the Order interface.
func (o *Raid) Execute() error { panic("!") }

type Setup struct {
	Line     int
	Id       int         // id of unit establishing ship or colony
	Location Coordinates // location being set up
	Kind     string      // must be 'colony' or 'ship'
	Action   string      // must be 'transfer'
	Items    []*TransferItem
	Errors   []error
}

// Execute implements the Order interface.
func (o *Setup) Execute() error { panic("!") }

type SupportAttack struct {
	Line         int
	Id           int // id of unit being ordered
	SupportId    int // id of unit being supported
	TargetId     int // id of unit being attacked
	PctCommitted int
	Errors       []error
}

// Execute implements the Order interface.
func (o *SupportAttack) Execute() error { panic("!") }

type SupportDefend struct {
	Line         int
	Id           int // id of unit being ordered
	SupportId    int // id of unit being supported
	PctCommitted int
	Errors       []error
}

// Execute implements the Order interface.
func (o *SupportDefend) Execute() error { panic("!") }

type TransferItem struct {
	Item      string // name
	Qty       int
	TechLevel int // optional tech level
}
