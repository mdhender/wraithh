// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package parser

type Kind int

const (
	EOF Kind = iota
	EOL
	ASSEMBLE
	BLANKLINE
	BOMBARD
	CARGO
	COMMA
	COMMAND
	ID
	INTEGER
	INVADE
	LINE
	LOCATION
	NAME
	PERCENTAGE
	PRODUCT
	PROFESSION
	QUANTITY
	RAID
	RESOURCE
	SUPPORT
)

type Product int

const (
	ANTIMISSILE Product = iota
	ASSAULTCRAFT
	ASSAULTWEAPONS
	AUTOMATION
	CONSUMERGOODS
	ENERGYSHIELD
	ENERGYWEAPON
	FACTORY
	FARM
	FOOD
	HYPERENGINE
	LIFESUPPORT
	LIGHTSTRUCTURAL
	MILITARYROBOT
	MILITARYSUPPLIES
	MINE
	MISSILE
	SENSOR
	SPACEDRIVE
	STRUCTURALUNIT
	SUPERLIGHTSTRUCTURAL
	TRANSPORT
)

type Profession int

const (
	CIVILIAN Profession = iota
	PROFESSIONAL
	SOLDIER
	UNSKILLEDWORKER
	CONSTRUCTIONCREW
	SPY
)

type Resource int

const (
	FUEL Resource = iota
	GOLD
	METALLICS
	NONMETALLICS
)