// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package adapters

import "fmt"

type Abandon struct {
	Line     int
	Location Coordinates // location to be abandoned
	Errors   []error
}

type AssembleFactoryGroup struct {
	Line        int
	Id          int  // id of unit being ordered
	Quantity    int  // number of units to assemble
	Unit        Unit // factory units to assemble
	Manufacture Unit // product unit to be manufactured
	Errors      []error
}

type AssembleMineGroup struct {
	Line      int
	Id        int    // id of unit being ordered
	DepositId string // deposit to assemble mines at
	Quantity  int    // number of units to assemble
	Unit      Unit   // mine units to assemble
	Errors    []error
}

type AssembleUnit struct {
	Line     int
	Id       int  // id of unit being ordered
	Quantity int  // number of units to assemble
	Unit     Unit // unit to assemble
	Errors   []error
}

type Bombard struct {
	Line         int
	Id           int // id of unit being ordered
	PctCommitted int
	TargetId     int // id of unit being attacked
	Errors       []error
}

type Buy struct {
	Line     int
	Id       int     // id of unit being ordered
	Quantity int     // number of units to purchase
	Unit     Unit    // unit to sell
	Bid      float64 // bid per unit
	Errors   []error
}

type CheckRebels struct {
	Line     int
	Id       int // id of unit being ordered
	Quantity int // number of units to use
	Errors   []error
}

type Claim struct {
	Line     int
	Id       int         // id of unit being ordered
	Location Coordinates // location to be claimed
	Errors   []error
}

type ConvertRebels struct {
	Line     int
	Id       int // id of unit being ordered
	Quantity int // number of units to use
	Errors   []error
}

type Coordinates struct { // location being set up
	X, Y, Z int
	Orbit   int
}

type CounterAgents struct {
	Line     int
	Id       int // id of unit being ordered
	Quantity int // number of units to use
	Errors   []error
}

type Discharge struct {
	Line       int
	Id         int    // id of unit being ordered
	Quantity   int    // number of units to use
	Profession string // profession to discharge from
	Errors     []error
}

type Draft struct {
	Line       int
	Id         int    // id of unit being ordered
	Quantity   int    // number of units to use
	Profession string // profession to draft into
	Errors     []error
}

type ExpandFactoryGroup struct {
	Line         int
	Id           int    // id of unit being ordered
	FactoryGroup string // factory group to expand
	Quantity     int    // number of units to assemble
	Unit         Unit   // mine units to assemble
	Errors       []error
}

type ExpandMineGroup struct {
	Line      int
	Id        int    // id of unit being ordered
	MineGroup string // mine group to expand
	Quantity  int    // number of units to assemble
	Unit      Unit   // mine units to assemble
	Errors    []error
}

type Grant struct {
	Line     int
	Location Coordinates // coordinates of system and orbit
	Kind     string      // kind of grant
	TargetId int         // nation to grant
	Errors   []error
}

type InciteRebels struct {
	Line     int
	Id       int // id of unit being ordered
	Quantity int // number of units to use
	TargetId int // id of nation to target
	Errors   []error
}

type Invade struct {
	Line         int
	Id           int // id of unit being ordered
	PctCommitted int
	TargetId     int // id of unit being attacked
	Errors       []error
}

type Move struct {
	Line     int
	Id       int         // id of unit being ordered
	Location Coordinates // coordinates to move to
	Errors   []error
}

type NameOrbit struct {
	Line     int
	Location Coordinates // coordinates of system and orbit to name
	Name     string      // new name for unit
	Errors   []error
}

type NameSystem struct {
	Line     int
	Location Coordinates // coordinates of system to name
	Name     string      // new name for unit
	Errors   []error
}

type NameUnit struct {
	Line   int
	Id     int    // id of unit being ordered
	Name   string // new name for unit
	Errors []error
}

type News struct {
	Line      int
	Location  Coordinates // location to send news to
	Article   string
	Signature string
	Errors    []error
}

type PayAll struct {
	Line       int
	Profession string  // profession to change pay for
	Rate       float64 // new pay rate
	Errors     []error
}

type PayLocal struct {
	Line       int
	Id         int     // id of unit being ordered
	Profession string  // profession to change pay for
	Rate       float64 // new pay rate
	Errors     []error
}

type Probe struct {
	Line   int
	Id     int // id of unit being ordered
	Orbit  int // orbit to probe
	Errors []error
}

type ProbeSystem struct {
	Line     int
	Id       int         // id of unit being ordered
	Location Coordinates // location to probe
	Errors   []error
}

type Raid struct {
	Line         int
	Id           int // id of unit being ordered
	PctCommitted int
	TargetId     int      // id of unit being raided
	TargetUnit   struct { // material to raid
		Unit      string
		TechLevel int
	}
	Errors []error
}

type RationAll struct {
	Line   int
	Rate   int // new ration percentage
	Errors []error
}

type RationLocal struct {
	Line   int
	Id     int // id of unit being ordered
	Rate   int // new ration percentage
	Errors []error
}

type RecycleFactoryGroup struct {
	Line         int
	Id           int    // id of unit being ordered
	FactoryGroup string // factory group to recycle units from
	Quantity     int    // number of units to recycle
	Unit         Unit   // unit to recycle
	Errors       []error
}

type RecycleMineGroup struct {
	Line      int
	Id        int    // id of unit being ordered
	MineGroup string // mine group to recycle units from
	Quantity  int    // number of units to recycle
	Unit      Unit   // unit to recycle
	Errors    []error
}

type RecycleUnit struct {
	Line     int
	Id       int  // id of unit being ordered
	Quantity int  // number of units to recycle
	Unit     Unit // unit to recycle
	Errors   []error
}

type RetoolFactoryGroup struct {
	Line         int
	Id           int    // id of unit being ordered
	FactoryGroup string // factory group to retool
	Unit         Unit   // new unit to manufacture
	Errors       []error
}

type Revoke struct {
	Line     int
	Location Coordinates // coordinates of system and orbit
	Kind     string      // kind of grant
	TargetId int         // nation to grant
	Errors   []error
}

type ScrapFactoryGroup struct {
	Line         int
	Id           int    // id of unit being ordered
	FactoryGroup string // factory group to scrap units from
	Quantity     int    // number of units to scrap
	Unit         Unit   // unit to scrap
	Errors       []error
}

type ScrapMineGroup struct {
	Line      int
	Id        int    // id of unit being ordered
	MineGroup string // mine group to scrap units from
	Quantity  int    // number of units to scrap
	Unit      Unit   // unit to scrap
	Errors    []error
}

type ScrapUnit struct {
	Line     int
	Id       int  // id of unit being ordered
	Quantity int  // number of units to scrap
	Unit     Unit // unit to scrap
	Errors   []error
}

type Sell struct {
	Line     int
	Id       int     // id of unit being ordered
	Quantity int     // number of units to sell
	Unit     Unit    // unit to sell
	Ask      float64 // ask per unit
	Errors   []error
}

type Setup struct {
	Line     int
	Id       int         // id of unit establishing ship or colony
	Location Coordinates // location being set up
	Kind     string      // must be 'colony' or 'ship'
	Action   string      // must be 'transfer'
	Items    []*TransferDetail
	Errors   []error
}

type StealSecrets struct {
	Line     int
	Id       int // id of unit being ordered
	Quantity int // number of units to use
	TargetId int // id of nation to target
	Errors   []error
}

type StoreFactoryGroup struct {
	Line         int
	Id           int    // id of unit being ordered
	FactoryGroup string // factory group to store units from
	Quantity     int    // number of units to store
	Unit         Unit   // unit to store
	Errors       []error
}

type StoreMineGroup struct {
	Line      int
	Id        int    // id of unit being ordered
	MineGroup string // mine group to store units from
	Quantity  int    // number of units to store
	Unit      Unit   // unit to store
	Errors    []error
}

type StoreUnit struct {
	Line     int
	Id       int  // id of unit being ordered
	Quantity int  // number of units to store
	Unit     Unit // unit to store
	Errors   []error
}

type SupportAttack struct {
	Line         int
	Id           int // id of unit being ordered
	PctCommitted int
	SupportId    int // id of unit being supported
	TargetId     int // id of unit being attacked
	Errors       []error
}

type SupportDefend struct {
	Line         int
	Id           int // id of unit being ordered
	SupportId    int // id of unit being supported
	PctCommitted int
	Errors       []error
}

type SuppressAgents struct {
	Line     int
	Id       int // id of unit being ordered
	Quantity int // number of units to use
	TargetId int // id of nation to target
	Errors   []error
}

type Survey struct {
	Line   int
	Id     int // id of unit being ordered
	Orbit  int // orbit to survey
	Errors []error
}

type SurveySystem struct {
	Line     int
	Id       int         // id of unit being ordered
	Location Coordinates // location to survey
	Errors   []error
}

type Transfer struct {
	Line     int
	Id       int  // id of unit being ordered
	Quantity int  // number of units to transfer
	Unit     Unit // unit to transfer
	TargetId int  // id of unit receiving units
	Errors   []error
}

type TransferDetail struct {
	Unit     Unit
	Quantity int
}

func (td *TransferDetail) String() string {
	return fmt.Sprintf("{%d %s}", td.Quantity, td.Unit)
}

type Unit struct {
	Name      string // name
	TechLevel int    // optional tech level
}

func (u Unit) String() string {
	if u.TechLevel == 0 {
		return u.Name
	}
	return fmt.Sprintf("%s-%d", u.Name, u.TechLevel)
}

type Unknown struct {
	Line    int
	Command string
	Errors  []error
}
