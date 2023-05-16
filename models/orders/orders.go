// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package orders

import (
	"fmt"
	"github.com/mdhender/wraithh/models/units"
)

type Order interface {
	Execute() error
}

type TransferDetail struct {
	Unit     units.Unit
	Quantity int
}

func (td *TransferDetail) String() string {
	return fmt.Sprintf("{%d %s}", td.Quantity, td.Unit)
}
