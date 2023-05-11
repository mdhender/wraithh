// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package orders

import (
	"fmt"
	"github.com/mdhender/wraithh/internal/parser"
	"github.com/mdhender/wraithh/internal/tokenizer"
	"strings"
)

func ptOrdersWalker(pt *parser.Tree) (orders []Order, err error) {
	if s, ok := pt.Symbol.(string); !ok || s != "orders" {
		return nil, fmt.Errorf("expected orders")
	}
	for _, st := range pt.Subtrees {
		if s, ok := st.Symbol.(string); ok && s == "order" {
			order, err := ptOrderWalker(st)
			if err != nil {
				return nil, err
			}
			orders = append(orders, order)
		} else {
			return nil, fmt.Errorf("expected order")
		}
	}
	return nil, nil
}

func ptOrderWalker(pt *parser.Tree) (order Order, err error) {
	for _, st := range pt.Subtrees {
		if s, ok := st.Symbol.(string); ok && s == "bombard" {
			order, err := ptBombardWalker(st)
			if err != nil {
				return nil, err
			}
			return order, nil
		} else if s, ok := st.Symbol.(string); ok && s == "invade" {
			order, err := ptInvadeWalker(st)
			if err != nil {
				return nil, err
			}
			return order, nil
		} else if s, ok := st.Symbol.(string); ok && s == "raid" {
			order, err := ptRaidWalker(st)
			if err != nil {
				return nil, err
			}
			return order, nil
		} else if s, ok := st.Symbol.(string); ok && s == "setup" {
			order, err := ptSetupWalker(st)
			if err != nil {
				return nil, err
			}
			return order, nil
		} else if s, ok := st.Symbol.(string); ok && s == "support" {
			order, err := ptSupportWalker(st)
			if err != nil {
				return nil, err
			}
			return order, nil
		} else {
			return nil, fmt.Errorf("expected order: got %q", s)
		}
	}
	return nil, nil
}

func ptAssembleWalker(pt *parser.Tree) (Order, error) {
	order := Assemble{}
	return &order, nil
}

func ptBombardWalker(pt *parser.Tree) (Order, error) {
	order := Bombard{}
	st := pt.Subtrees
	if len(st) == 0 {
		return nil, fmt.Errorf("expected bombard order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.TEXT || tok.Text != "bombard" {
		return nil, fmt.Errorf("expected bombard order")
	} else {
		order.Line = tok.Line
	}
	// csid
	if st = st[1:]; len(st) == 0 {
		return nil, fmt.Errorf("expected bombard order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.INTEGER {
		return nil, fmt.Errorf("expected bombard order")
	} else {
		order.Id = tok.Integer
	}
	// target id
	if st = st[1:]; len(st) == 0 {
		return nil, fmt.Errorf("expected bombard order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.INTEGER {
		return nil, fmt.Errorf("expected bombard order")
	} else {
		order.TargetId = tok.Integer
	}
	// pct committed
	if st = st[1:]; len(st) == 0 {
		return nil, fmt.Errorf("expected bombard order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.PERCENTAGE {
		return nil, fmt.Errorf("expected bombard order")
	} else {
		order.PctCommitted = tok.Integer
	}
	// eol
	if st = st[1:]; len(st) == 0 {
		return nil, fmt.Errorf("expected bombard order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.EOL {
		return nil, fmt.Errorf("expected bombard order")
	}

	if st = st[1:]; len(st) != 0 {
		return nil, fmt.Errorf("expected bombard order")
	}
	return &order, nil
}

func ptCargoWalker(pt *parser.Tree) (string, error) {
	if s, ok := pt.Symbol.(string); !ok || s != "cargo" {
		return "", fmt.Errorf("expected cargo")
	}
	st := pt.Subtrees
	if len(st) == 0 {
		return "", fmt.Errorf("expected cargo")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok {
		return "", fmt.Errorf("expected cargo")
	} else {
		// population or product or research or resource
		switch tok.Kind {
		case tokenizer.POPULATION:
			return tok.Text, nil
		case tokenizer.PRODUCT:
			return tok.Text, nil
		case tokenizer.RESEARCH:
			return tok.Text, nil
		case tokenizer.RESOURCE:
			return tok.Text, nil
		}
	}
	return "", fmt.Errorf("expected cargo")
}

func ptCoordinateWalker(pt *parser.Tree) (int, int, int, int, error) {
	if s, ok := pt.Symbol.(string); !ok || s != "coordinate" {
		return 0, 0, 0, 0, fmt.Errorf("expected coordinate")
	}
	st := pt.Subtrees
	if len(st) == 0 {
		return 0, 0, 0, 0, fmt.Errorf("expected coordinate")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.PARENOP {
		return 0, 0, 0, 0, fmt.Errorf("expected coordinate")
	}
	var x int
	if st = st[1:]; len(st) == 0 {
		return 0, 0, 0, 0, fmt.Errorf("expected coordinate")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.INTEGER {
		return 0, 0, 0, 0, fmt.Errorf("expected coordinate")
	} else {
		x = tok.Integer
	}
	// comma
	if st = st[1:]; len(st) == 0 {
		return 0, 0, 0, 0, fmt.Errorf("expected coordinate")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.COMMA {
		return 0, 0, 0, 0, fmt.Errorf("expected coordinate")
	}
	var y int
	if st = st[1:]; len(st) == 0 {
		return 0, 0, 0, 0, fmt.Errorf("expected coordinate")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.INTEGER {
		return 0, 0, 0, 0, fmt.Errorf("expected coordinate")
	} else {
		y = tok.Integer
	}
	// comma
	if st = st[1:]; len(st) == 0 {
		return 0, 0, 0, 0, fmt.Errorf("expected coordinate")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.COMMA {
		return 0, 0, 0, 0, fmt.Errorf("expected coordinate")
	}
	var z int
	if st = st[1:]; len(st) == 0 {
		return 0, 0, 0, 0, fmt.Errorf("expected coordinate")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.INTEGER {
		return 0, 0, 0, 0, fmt.Errorf("expected coordinate")
	} else {
		z = tok.Integer
	}
	// optional orbit
	var orbit int
	if tok, ok := st[0].Symbol.(*tokenizer.Token); ok && tok.Kind == tokenizer.COMMA {
		if st = st[1:]; len(st) == 0 {
			return 0, 0, 0, 0, fmt.Errorf("expected coordinate")
		} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.INTEGER {
			return 0, 0, 0, 0, fmt.Errorf("expected coordinate")
		} else {
			orbit = tok.Integer
		}
	}
	if st = st[1:]; len(st) == 0 {
		return 0, 0, 0, 0, fmt.Errorf("expected coordinate")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.PARENCL {
		return 0, 0, 0, 0, fmt.Errorf("expected coordinate")
	}
	return x, y, z, orbit, nil
}

func ptInvadeWalker(pt *parser.Tree) (Order, error) {
	order := Invade{}
	st := pt.Subtrees
	if len(st) == 0 {
		return nil, fmt.Errorf("expected invade order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.TEXT || tok.Text != "invade" {
		return nil, fmt.Errorf("expected invade order")
	} else {
		order.Line = tok.Line
	}
	// csid
	if st = st[1:]; len(st) == 0 {
		return nil, fmt.Errorf("expected invade order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.INTEGER {
		return nil, fmt.Errorf("expected invade order")
	} else {
		order.Id = tok.Integer
	}
	// target id
	if st = st[1:]; len(st) == 0 {
		return nil, fmt.Errorf("expected invade order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.INTEGER {
		return nil, fmt.Errorf("expected invade order")
	} else {
		order.TargetId = tok.Integer
	}
	// pct committed
	if st = st[1:]; len(st) == 0 {
		return nil, fmt.Errorf("expected invade order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.PERCENTAGE {
		return nil, fmt.Errorf("expected invade order")
	} else {
		order.PctCommitted = tok.Integer
	}
	// eol
	if st = st[1:]; len(st) == 0 {
		return nil, fmt.Errorf("expected invade order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.EOL {
		return nil, fmt.Errorf("expected invade order")
	}

	if st = st[1:]; len(st) != 0 {
		return nil, fmt.Errorf("expected invade order")
	}
	return &order, nil
}

func ptRaidWalker(pt *parser.Tree) (Order, error) {
	order := Raid{}
	st := pt.Subtrees
	if len(st) == 0 {
		return nil, fmt.Errorf("expected raid order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.TEXT || tok.Text != "raid" {
		return nil, fmt.Errorf("expected raid order")
	} else {
		order.Line = tok.Line
	}
	// csid
	if st = st[1:]; len(st) == 0 {
		return nil, fmt.Errorf("expected raid order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.INTEGER {
		return nil, fmt.Errorf("expected raid order")
	} else {
		order.Id = tok.Integer
	}
	// target id
	if st = st[1:]; len(st) == 0 {
		return nil, fmt.Errorf("expected raid order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.INTEGER {
		return nil, fmt.Errorf("expected raid order")
	} else {
		order.TargetId = tok.Integer
	}
	// pct committed
	if st = st[1:]; len(st) == 0 {
		return nil, fmt.Errorf("expected raid order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.PERCENTAGE {
		return nil, fmt.Errorf("expected raid order")
	} else {
		order.PctCommitted = tok.Integer
	}
	// material
	if st = st[1:]; len(st) == 0 {
		return nil, fmt.Errorf("expected raid order")
	} else if material, err := ptCargoWalker(st[0]); err != nil {
		return nil, fmt.Errorf("expected raid order")
	} else {
		order.Material = material
	}
	// eol
	if st = st[1:]; len(st) == 0 {
		return nil, fmt.Errorf("expected raid order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.EOL {
		return nil, fmt.Errorf("expected raid order")
	}
	if st = st[1:]; len(st) != 0 {
		return nil, fmt.Errorf("expected raid order")
	}
	return &order, nil
}

func ptSetupWalker(pt *parser.Tree) (Order, error) {
	order := Setup{}
	st := pt.Subtrees
	if len(st) == 0 {
		return nil, fmt.Errorf("expected setup order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.TEXT || strings.ToLower(tok.Text) != "setup" {
		return nil, fmt.Errorf("expected setup order")
	} else {
		order.Line = tok.Line
	}
	// csid
	if st = st[1:]; len(st) == 0 {
		return nil, fmt.Errorf("expected setup order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.INTEGER {
		return nil, fmt.Errorf("expected setup order")
	} else {
		order.Id = tok.Integer
	}
	// coordinates
	if st = st[1:]; len(st) == 0 {
		return nil, fmt.Errorf("expected setup order")
	} else if x, y, z, orbit, err := ptCoordinateWalker(st[0]); err != nil {
		return nil, fmt.Errorf("expected raid order")
	} else {
		order.Location = Coordinates{x, y, z, orbit}
	}
	// colony or ship
	if st = st[1:]; len(st) == 0 {
		return nil, fmt.Errorf("expected setup order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.TEXT {
		return nil, fmt.Errorf("expected setup order")
	} else {
		switch strings.ToLower(tok.Text) {
		case "colony":
			order.Kind = "COLONY"
		case "ship":
			order.Kind = "SHIP"
		default:
			return nil, fmt.Errorf("expected setup order")
		}
	}
	// transfer
	if st = st[1:]; len(st) == 0 {
		return nil, fmt.Errorf("expected setup order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.TEXT {
		return nil, fmt.Errorf("expected setup order")
	} else {
		switch strings.ToLower(tok.Text) {
		case "transfer":
			order.Action = "TRANSFER"
		default:
			return nil, fmt.Errorf("expected setup order")
		}
	}
	// eol
	if st = st[1:]; len(st) == 0 {
		return nil, fmt.Errorf("expected setup order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.EOL {
		return nil, fmt.Errorf("expected setup order")
	}
	// items
	for st = st[1:]; len(st) != 0; {
		if tok, ok := st[0].Symbol.(*tokenizer.Token); ok && tok.Kind == tokenizer.TEXT && strings.ToLower(tok.Text) == "end" {
			break
		}
		// cargo
		if item, err := ptXferDetailWalker(st[0]); err != nil {
			return nil, fmt.Errorf("expected setup order")
		} else {
			order.Items = append(order.Items, item)
		}
		st = st[1:]
	}
	// end
	if len(st) == 0 {
		return nil, fmt.Errorf("expected setup order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.TEXT || strings.ToLower(tok.Text) != "end" {
		return nil, fmt.Errorf("expected setup order")
	}
	// eol
	if st = st[1:]; len(st) == 0 {
		return nil, fmt.Errorf("expected setup order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.EOL {
		return nil, fmt.Errorf("expected setup order")
	}
	if st = st[1:]; len(st) != 0 {
		return nil, fmt.Errorf("expected setup order")
	}
	return &order, nil
}

func ptSupportWalker(pt *parser.Tree) (Order, error) {
	var sa *SupportAttack
	sd := SupportDefend{}
	st := pt.Subtrees
	if len(st) == 0 {
		return nil, fmt.Errorf("expected support order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.TEXT || tok.Text != "support" {
		return nil, fmt.Errorf("expected support order")
	} else {
		sd.Line = tok.Line
	}
	// csid
	if st = st[1:]; len(st) == 0 {
		return nil, fmt.Errorf("expected support order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.INTEGER {
		return nil, fmt.Errorf("expected support order")
	} else {
		sd.Id = tok.Integer
	}
	// support id
	if st = st[1:]; len(st) == 0 {
		return nil, fmt.Errorf("expected support order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.INTEGER {
		return nil, fmt.Errorf("expected support order")
	} else {
		sd.SupportId = tok.Integer
	}
	// pct committed
	if st = st[1:]; len(st) == 0 {
		return nil, fmt.Errorf("expected support order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.PERCENTAGE {
		return nil, fmt.Errorf("expected support order")
	} else {
		sd.PctCommitted = tok.Integer
	}
	// optional target id
	saved := st
	if st = st[1:]; len(st) == 0 {
		return nil, fmt.Errorf("expected support order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); ok && tok.Kind == tokenizer.INTEGER {
		sa = &SupportAttack{
			Line:      sd.Line,
			Id:        sd.Id,
			SupportId: sd.SupportId,
			TargetId:  tok.Integer,
		}
	} else {
		st = saved
	}
	// eol
	if st = st[1:]; len(st) == 0 {
		return nil, fmt.Errorf("expected support order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.EOL {
		return nil, fmt.Errorf("expected support order")
	}
	if st = st[1:]; len(st) != 0 {
		return nil, fmt.Errorf("expected support order")
	}
	if sa != nil {
		return sa, nil
	}
	return &sd, nil
}

func ptXferDetailWalker(pt *parser.Tree) (*TransferItem, error) {
	if s, ok := pt.Symbol.(string); !ok || s != "xfer_detail" {
		return nil, fmt.Errorf("expected xfer_detail")
	}
	item := TransferItem{}
	st := pt.Subtrees
	// quantity
	if len(st) == 0 {
		return nil, fmt.Errorf("expected xfer_detail")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.INTEGER {
		return nil, fmt.Errorf("expected xfer_detail")
	} else {
		item.Qty = tok.Integer
	}
	// cargo
	if st = st[1:]; len(st) == 0 {
		return nil, fmt.Errorf("expected xfer_detail")
	} else if cargo, err := ptCargoWalker(st[0]); err != nil {
		return nil, fmt.Errorf("expected xfer_detail")
	} else {
		item.Item = cargo
	}
	// eol
	if st = st[1:]; len(st) == 0 {
		return nil, fmt.Errorf("expected support order")
	} else if tok, ok := st[0].Symbol.(*tokenizer.Token); !ok || tok.Kind != tokenizer.EOL {
		return nil, fmt.Errorf("expected support order")
	}
	if st = st[1:]; len(st) != 0 {
		return nil, fmt.Errorf("expected support order")
	}
	return &item, nil
}
