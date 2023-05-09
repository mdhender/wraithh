// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package orders

// Grammar is EBNF representation of an order file's syntax.
const Grammar = `
	orders = {order | EOL} EOF .

	order = assemble | bombard | invade | raid | retool | setup | support .

	assemble    = "assemble"    CSID [DEPOSITID | FACTGRP | MINEGRP] QUANTITY material EOL .
	disassemble = "disassemble" CSID [            FACTGRP | MINEGRP] QUANTITY material EOL .
	retool      = "retool"      CSID FACTGRP material EOL .

	bombard  = "bombard"  CSID CSID        PERCENTAGE       EOL .
	invade   = "invade"   CSID CSID        PERCENTAGE       EOL .
	raid     = "raid"     CSID CSID        PERCENTAGE cargo EOL .
	support  = "support"  CSID CSID [CSID] PERCENTAGE       EOL .

	transfer = "transfer" CSID QUANTITY cargo CSID EOL .

	setup    = "setup"    CSID coordinate ("ship" | "colony") "transfer" EOL
	           {xfer_detail}
	           "end" EOL .

	cargo       = POPULATION | PRODUCT | RESEARCH | RESOURCE .
	coordinate  = PARENOP INTEGER COMMA INTEGER COMMA INTEGER [COMMA INTEGER] PARENCL .
	material    = "research" | PRODUCT .
	xfer_detail = QUANTITY cargo TEXT EOL .
`
