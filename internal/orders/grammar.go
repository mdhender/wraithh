// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package orders

// Grammar is EBNF representation of an order file's syntax.
const Grammar = `
	orders = {order | EOL} EOF .

	order = assemble | bombard | invade | raid | setup | support .

	assemble = "assemble" INTEGER INTEGER TEXT [ INTEGER | TEXT ] EOL .
	bombard  = "bombard"  INTEGER INTEGER PERCENTAGE EOL .
	invade   = "invade"   INTEGER INTEGER PERCENTAGE EOL .
	raid     = "raid"     INTEGER INTEGER PERCENTAGE TEXT EOL .
	setup    = "setup"    INTEGER TEXT ("ship" || "colony") "transfer" EOL
	           {xfer_detail}
	           "end" EOL . 
	support  = "support"  INTEGER INTEGER [INTEGER] PERCENTAGE EOL .

	xfer_detail = INTEGER TEXT EOL . 
`
