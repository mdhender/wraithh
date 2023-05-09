// Copyright (c) 2023 Michael D Henderson.
// SPDX-License-Identifier: AGPL-3.0-or-later

package orders

// Grammar is EBNF representation of an order file's syntax.
const Grammar = `
	orders = {order | EOL} EOF .

	order = assemble | bombard | invade | raid | retool | setup | support .

	bombard  = "bombard"  CSID CSID        PERCENTAGE       EOL .
	invade   = "invade"   CSID CSID        PERCENTAGE       EOL .
	raid     = "raid"     CSID CSID        PERCENTAGE cargo EOL .
	support  = "support"  CSID CSID [CSID] PERCENTAGE       EOL .

	setup    = "setup"    CSID coordinate ("ship" | "colony") "transfer" EOL
	           {xfer_detail}
	           "end" EOL .

	transfer = "transfer" CSID INTEGER cargo CSID EOL .

	assemble    = "assemble"    CSID [DEPOSITID | FACTGRP | MINEGRP] INTEGER material EOL .
	disassemble = "disassemble" CSID [            FACTGRP | MINEGRP] INTEGER material EOL .
	retool      = "retool"      CSID FACTGRP material EOL .

	buy  = "buy"  CSID (RESEARCH | (PRODUCT INTEGER)) number EOL .
	sell = "sell" CSID (RESEARCH | (PRODUCT INTEGER)) number EOL .

	survey = "survey" CSID EOL .
	probe  = "probe"  CSID (INTEGER | coordinates) EOL .

	spy = mission CSID INTEGER [CSID] EOL .

	news = "news" coordinate TEXT TEXT EOL .

	move = "move" CSID coordinate EOL .

	draft   = "draft"   CSID INTEGER POPULATION EOL.
	disband = "disband" CSID INTEGER POPULATION EOL.

	pay = "pay" [CSID] POPULATION NUMBER EOL.

	ration = "ration" [CSID] PERCENTAGE EOL.

	control = "control" CSID coordinate EOL .

	cargo       = POPULATION | PRODUCT | RESEARCH | RESOURCE .
	coordinate  = PARENOP INTEGER COMMA INTEGER COMMA INTEGER [COMMA INTEGER] PARENCL .
	material    = "research" | PRODUCT .
	mission     = "check-rebels" | "convert-rebels" | "counter-agents"
	            | "suppress-agents" | "incite-rebels" | "steal-reports" .
	xfer_detail = QUANTITY cargo TEXT EOL .
`
