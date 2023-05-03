// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package parser

type Nodes []Node

type Node interface {
	Kind() Kind
}

type AssembleNode struct {
	line    int
	id      int           // id of unit being ordered
	qty     int           // number of units to assemble
	what    *ProductNode  // type of units to assemble
	mine    *LocationNode // for mines, which deposit to work
	factory *ProductNode  // for factories, which product to manufacture
	err     error
}

func (n *AssembleNode) Kind() Kind {
	return ASSEMBLE
}

type BlankNode struct {
	line int
}

func (n *BlankNode) Kind() Kind {
	return BLANKLINE
}

type BombardNode struct {
	line         int
	id           int
	targetId     int
	pctCommitted int
	err          error
}

func (n *BombardNode) Kind() Kind {
	return BOMBARD
}

type CommaNode struct {
	line int
}

func (n *CommaNode) Kind() Kind {
	return COMMA
}

type CommandNode struct {
	line  int
	value string
}

func (n *CommandNode) Kind() Kind {
	return COMMAND
}

type EOFNode struct {
	line int
}

func (n *EOFNode) Kind() Kind {
	return EOF
}

type EOLNode struct {
	line int
}

func (n *EOLNode) Kind() Kind {
	return EOL
}

type IDNode struct {
	line int
	id   int
}

func (n *IDNode) Kind() Kind {
	return ID
}

type IntegerNode struct {
	line  int
	value int
}

func (n *IntegerNode) Kind() Kind {
	return INTEGER
}

type InvadeNode struct {
	line         int
	id           int
	targetId     int
	pctCommitted int
	err          error
}

func (n *InvadeNode) Kind() Kind {
	return INVADE
}

type LineNode struct {
	line  int
	value string
}

func (n *LineNode) Kind() Kind {
	return LINE
}

type LocationNode struct {
	line int
}

func (n *LocationNode) Kind() Kind {
	return LOCATION
}

type NameNode struct {
	line  int
	value string
}

func (n *NameNode) Kind() Kind {
	return NAME
}

type PercentageNode struct {
	line  int
	value int
}

func (n *PercentageNode) Kind() Kind {
	return PERCENTAGE
}

type ProductNode struct {
	line      int
	_type     Product
	techLevel int
}

func (n *ProductNode) Kind() Kind {
	return PRODUCT
}

type ProfessionNode struct {
	line  int
	_type Profession
}

func (n *ProfessionNode) Kind() Kind {
	return PROFESSION
}

func (n *ProfessionNode) Type() Profession {
	return n._type
}

type QuantityNode struct {
	line   int
	amount int
}

func (n *QuantityNode) Kind() Kind {
	return QUANTITY
}

type RaidNode struct {
	line         int
	id           int
	targetId     int
	pctCommitted int
	product      *ProductNode
	profession   *ProfessionNode
	resource     *ResourceNode
	err          error
}

func (n *RaidNode) Kind() Kind {
	return RAID
}

type ResourceNode struct {
	line  int
	_type Resource
}

func (n *ResourceNode) Kind() Kind {
	return RESOURCE
}
func (n *ResourceNode) Type() Resource {
	return n._type
}

type SupportNode struct {
	line         int
	id           int // id of unit being ordered
	supportId    int // id of unit being supported
	targetId     int // id of unit being attacked
	pctCommitted int
	err          error
}

func (n *SupportNode) Kind() Kind {
	return SUPPORT
}
