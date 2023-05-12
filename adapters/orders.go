// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package adapters

import (
	"fmt"
	"github.com/mdhender/wraithh/ec"
	"github.com/mdhender/wraithh/parsers/orders"
)

func CoordToEngineCoord(in orders.Coordinates) ec.Coordinates {
	return ec.Coordinates{
		X:      in.X,
		Y:      in.Y,
		Z:      in.Z,
		System: in.System,
		Orbit:  in.Orbit,
	}
}
func ItemsToEngineItems(in []*orders.TransferDetail) (out []*ec.TransferDetail) {
	for _, item := range in {
		out = append(out, &ec.TransferDetail{
			Unit:     UnitToEngineUnit(item.Unit),
			Quantity: item.Quantity,
		})
	}
	return out
}
func UnitToEngineUnit(in orders.Unit) ec.Unit {
	return ec.Unit{
		Name:      in.Name,
		TechLevel: in.TechLevel,
	}
}
func OrdersToEngineOrders(in []any) (out []ec.Order) {
	for _, o := range in {
		switch order := o.(type) {
		case *orders.Abandon:
			out = append(out, &ec.Abandon{
				Line:     order.Line,
				Location: CoordToEngineCoord(order.Location),
			})
		case *orders.AssembleFactoryGroup:
			out = append(out, &ec.AssembleFactoryGroup{
				Line:        order.Line,
				Id:          order.Id,
				Quantity:    order.Quantity,
				Unit:        UnitToEngineUnit(order.Unit),
				Manufacture: UnitToEngineUnit(order.Manufacture),
			})
		case *orders.AssembleMineGroup:
			out = append(out, &ec.AssembleMineGroup{
				Line:      order.Line,
				Id:        order.Id,
				DepositId: order.DepositId,
				Quantity:  order.Quantity,
				Unit:      UnitToEngineUnit(order.Unit),
			})
		case *orders.AssembleUnit:
			out = append(out, &ec.AssembleUnit{
				Line:     order.Line,
				Id:       order.Id,
				Quantity: order.Quantity,
				Unit:     UnitToEngineUnit(order.Unit),
			})
		case *orders.Bombard:
			out = append(out, &ec.Bombard{
				Line:         order.Line,
				Id:           order.Id,
				PctCommitted: order.PctCommitted,
				TargetId:     order.TargetId,
			})
		case *orders.Buy:
			out = append(out, &ec.Buy{
				Line:     order.Line,
				Id:       order.Id,
				Quantity: order.Quantity,
				Unit:     UnitToEngineUnit(order.Unit),
				Bid:      order.Bid,
			})
		case *orders.CheckRebels:
			out = append(out, &ec.CheckRebels{
				Line:     order.Line,
				Id:       order.Id,
				Quantity: order.Quantity,
			})
		case *orders.Claim:
			out = append(out, &ec.Claim{
				Line:     order.Line,
				Id:       order.Id,
				Location: CoordToEngineCoord(order.Location),
			})
		case *orders.ConvertRebels:
			out = append(out, &ec.ConvertRebels{
				Line:     order.Line,
				Id:       order.Id,
				Quantity: order.Quantity,
			})
		case *orders.CounterAgents:
			out = append(out, &ec.CounterAgents{
				Line:     order.Line,
				Id:       order.Id,
				Quantity: order.Quantity,
			})
		case *orders.Discharge:
			out = append(out, &ec.Discharge{
				Line:       order.Line,
				Id:         order.Id,
				Quantity:   order.Quantity,
				Profession: order.Profession,
			})
		case *orders.Draft:
			out = append(out, &ec.Draft{
				Line:       order.Line,
				Id:         order.Id,
				Quantity:   order.Quantity,
				Profession: order.Profession,
			})
		case *orders.ExpandFactoryGroup:
			out = append(out, &ec.ExpandFactoryGroup{
				Line:         order.Line,
				Id:           order.Id,
				FactoryGroup: order.FactoryGroup,
				Quantity:     order.Quantity,
				Unit:         UnitToEngineUnit(order.Unit),
			})
		case *orders.ExpandMineGroup:
			out = append(out, &ec.ExpandMineGroup{
				Line:      order.Line,
				Id:        order.Id,
				MineGroup: order.MineGroup,
				Quantity:  order.Quantity,
				Unit:      UnitToEngineUnit(order.Unit),
			})
		case *orders.Grant:
			out = append(out, &ec.Grant{
				Line:     order.Line,
				Location: CoordToEngineCoord(order.Location),
				Kind:     order.Kind,
				TargetId: order.TargetId,
			})
		case *orders.InciteRebels:
			out = append(out, &ec.InciteRebels{
				Line:     order.Line,
				Id:       order.Id,
				Quantity: order.Quantity,
				TargetId: order.TargetId,
			})
		case *orders.Invade:
			out = append(out, &ec.Invade{
				Line:         order.Line,
				Id:           order.Id,
				PctCommitted: order.PctCommitted,
				TargetId:     order.TargetId,
			})
		case *orders.Jump:
			out = append(out, &ec.Jump{
				Line:     order.Line,
				Id:       order.Id,
				Location: CoordToEngineCoord(order.Location),
			})
		case *orders.Move:
			out = append(out, &ec.Move{
				Line:  order.Line,
				Id:    order.Id,
				Orbit: order.Orbit,
			})
		case *orders.Name:
			out = append(out, &ec.Name{
				Line:     order.Line,
				Location: CoordToEngineCoord(order.Location),
				Name:     order.Name,
			})
		case *orders.NameUnit:
			out = append(out, &ec.NameUnit{
				Line: order.Line,
				Id:   order.Id,
				Name: order.Name,
			})
		case *orders.News:
			out = append(out, &ec.News{
				Line:      order.Line,
				Location:  CoordToEngineCoord(order.Location),
				Article:   order.Article,
				Signature: order.Signature,
			})
		case *orders.PayAll:
			out = append(out, &ec.PayAll{
				Line:       order.Line,
				Profession: order.Profession,
				Rate:       order.Rate,
			})
		case *orders.PayLocal:
			out = append(out, &ec.PayLocal{
				Line:       order.Line,
				Id:         order.Id,
				Profession: order.Profession,
				Rate:       order.Rate,
			})
		case *orders.Probe:
			out = append(out, &ec.Probe{
				Line:  order.Line,
				Id:    order.Id,
				Orbit: order.Orbit,
			})
		case *orders.ProbeSystem:
			out = append(out, &ec.ProbeSystem{
				Line:     order.Line,
				Id:       order.Id,
				Location: CoordToEngineCoord(order.Location),
			})
		case *orders.Raid:
			out = append(out, &ec.Raid{
				Line:         order.Line,
				Id:           order.Id,
				PctCommitted: order.PctCommitted,
				TargetId:     order.TargetId,
				TargetUnit:   UnitToEngineUnit(order.TargetUnit),
			})
		case *orders.RationAll:
			out = append(out, &ec.RationAll{
				Line: order.Line,
				Rate: order.Rate,
			})
		case *orders.RationLocal:
			out = append(out, &ec.RationLocal{
				Line: order.Line,
				Id:   order.Id,
				Rate: order.Rate,
			})
		case *orders.RecycleFactoryGroup:
			out = append(out, &ec.RecycleFactoryGroup{
				Line:         order.Line,
				Id:           order.Id,
				FactoryGroup: order.FactoryGroup,
				Quantity:     order.Quantity,
				Unit:         UnitToEngineUnit(order.Unit),
			})
		case *orders.RecycleMineGroup:
			out = append(out, &ec.RecycleMineGroup{
				Line:      order.Line,
				Id:        order.Id,
				MineGroup: order.MineGroup,
				Quantity:  order.Quantity,
				Unit:      UnitToEngineUnit(order.Unit),
			})
		case *orders.RecycleUnit:
			out = append(out, &ec.RecycleUnit{
				Line:     order.Line,
				Id:       order.Id,
				Quantity: order.Quantity,
				Unit:     UnitToEngineUnit(order.Unit),
			})
		case *orders.RetoolFactoryGroup:
			out = append(out, &ec.RetoolFactoryGroup{
				Line:         order.Line,
				Id:           order.Id,
				FactoryGroup: order.FactoryGroup,
				Unit:         UnitToEngineUnit(order.Unit),
			})
		case *orders.Revoke:
			out = append(out, &ec.Revoke{
				Line:     order.Line,
				Location: CoordToEngineCoord(order.Location),
				Kind:     order.Kind,
				TargetId: order.TargetId,
			})
		case *orders.ScrapFactoryGroup:
			out = append(out, &ec.ScrapFactoryGroup{
				Line:         order.Line,
				Id:           order.Id,
				FactoryGroup: order.FactoryGroup,
				Quantity:     order.Quantity,
				Unit:         UnitToEngineUnit(order.Unit),
			})
		case *orders.ScrapMineGroup:
			out = append(out, &ec.ScrapMineGroup{
				Line:      order.Line,
				Id:        order.Id,
				MineGroup: order.MineGroup,
				Quantity:  order.Quantity,
				Unit:      UnitToEngineUnit(order.Unit),
			})
		case *orders.ScrapUnit:
			out = append(out, &ec.ScrapUnit{
				Line:     order.Line,
				Id:       order.Id,
				Quantity: order.Quantity,
				Unit:     UnitToEngineUnit(order.Unit),
			})
		case *orders.Secret:
			out = append(out, &ec.Secret{
				Line:   order.Line,
				Handle: order.Handle,
				Game:   order.Game,
				Turn:   order.Turn,
				Token:  order.Token,
			})
		case *orders.Sell:
			out = append(out, &ec.Sell{
				Line:     order.Line,
				Id:       order.Id,
				Quantity: order.Quantity,
				Unit:     UnitToEngineUnit(order.Unit),
				Ask:      order.Ask,
			})
		case *orders.Setup:
			out = append(out, &ec.Setup{
				Line:     order.Line,
				Id:       order.Id,
				Location: CoordToEngineCoord(order.Location),
				Kind:     order.Kind,
				Action:   order.Action,
				Items:    ItemsToEngineItems(order.Items),
			})
		case *orders.StealSecrets:
			out = append(out, &ec.StealSecrets{
				Line:     order.Line,
				Id:       order.Id,
				Quantity: order.Quantity,
				TargetId: order.TargetId,
			})
		case *orders.StoreFactoryGroup:
			out = append(out, &ec.StoreFactoryGroup{
				Line:         order.Line,
				Id:           order.Id,
				FactoryGroup: order.FactoryGroup,
				Quantity:     order.Quantity,
				Unit:         UnitToEngineUnit(order.Unit),
			})
		case *orders.StoreMineGroup:
			out = append(out, &ec.StoreMineGroup{
				Line:      order.Line,
				Id:        order.Id,
				MineGroup: order.MineGroup,
				Quantity:  order.Quantity,
				Unit:      UnitToEngineUnit(order.Unit),
			})
		case *orders.StoreUnit:
			out = append(out, &ec.StoreUnit{
				Line:     order.Line,
				Id:       order.Id,
				Quantity: order.Quantity,
				Unit:     UnitToEngineUnit(order.Unit),
			})
		case *orders.SupportAttack:
			out = append(out, &ec.SupportAttack{
				Line:         order.Line,
				Id:           order.Id,
				PctCommitted: order.PctCommitted,
				SupportId:    order.SupportId,
				TargetId:     order.TargetId,
			})
		case *orders.SupportDefend:
			out = append(out, &ec.SupportDefend{
				Line:         order.Line,
				Id:           order.Id,
				SupportId:    order.SupportId,
				PctCommitted: order.PctCommitted,
			})
		case *orders.SuppressAgents:
			out = append(out, &ec.SuppressAgents{
				Line:     order.Line,
				Id:       order.Id,
				Quantity: order.Quantity,
				TargetId: order.TargetId,
			})
		case *orders.Survey:
			out = append(out, &ec.Survey{
				Line:  order.Line,
				Id:    order.Id,
				Orbit: order.Orbit,
			})
		case *orders.SurveySystem:
			out = append(out, &ec.SurveySystem{
				Line:     order.Line,
				Id:       order.Id,
				Location: CoordToEngineCoord(order.Location),
			})
		case *orders.Transfer:
			out = append(out, &ec.Transfer{
				Line:     order.Line,
				Id:       order.Id,
				Quantity: order.Quantity,
				Unit:     UnitToEngineUnit(order.Unit),
				TargetId: order.TargetId,
			})
		case *orders.Unknown:
			// ignore unknown orders
		default:
			panic(fmt.Sprintf("unknown type %T", o))
		}
	}
	return out
}
