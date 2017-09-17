//line parser.go.y:2
package vm

import __yyfmt__ "fmt"

//line parser.go.y:2
//line parser.go.y:5
type yySymType struct {
	yys   int
	expr  Expr
	exprs []Expr
	str   string
	lit   interface{}
}

const IDENT = 57346
const LIT = 57347
const FOR = 57348
const IN = 57349

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"LIT",
	"FOR",
	"IN",
	"','",
	"'('",
	"')'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'.'",
	"'['",
	"']'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:94

/* vim: set et sw=2: */

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 53

var yyAct = [...]int{

	26, 3, 31, 25, 37, 31, 14, 32, 29, 18,
	19, 20, 21, 15, 23, 16, 17, 27, 8, 9,
	10, 11, 12, 13, 30, 24, 8, 9, 10, 11,
	12, 13, 35, 34, 36, 8, 9, 10, 11, 12,
	13, 6, 4, 2, 6, 4, 5, 28, 33, 5,
	22, 7, 1,
}
var yyPact = [...]int{

	37, -1000, 47, 24, -1000, 40, 4, 8, 40, 40,
	40, 40, 46, 40, 15, 40, 40, 43, 24, 24,
	24, 24, -1, 7, -1000, -3, 24, 24, 41, 40,
	-1000, 40, -1000, 40, -6, 24, 24, -1000,
}
var yyPgo = [...]int{

	0, 52, 0, 3,
}
var yyR1 = [...]int{

	0, 1, 1, 1, 3, 3, 3, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2,
}
var yyR2 = [...]int{

	0, 4, 6, 1, 0, 1, 3, 1, 3, 3,
	3, 3, 3, 4, 6, 3, 4, 1,
}
var yyChk = [...]int{

	-1000, -1, 6, -2, 5, 9, 4, 4, 11, 12,
	13, 14, 15, 16, -2, 9, 7, 8, -2, -2,
	-2, -2, 4, -2, 10, -3, -2, -2, 4, 9,
	17, 8, 10, 7, -3, -2, -2, 10,
}
var yyDef = [...]int{

	0, -2, 0, 3, 7, 0, 17, 0, 0, 0,
	0, 0, 0, 0, 0, 4, 0, 0, 9, 10,
	11, 12, 15, 0, 8, 0, 5, 1, 0, 4,
	16, 0, 13, 0, 0, 6, 2, 14,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	9, 10, 13, 11, 8, 12, 15, 14, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 16, 3, 17,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:21
		{
			yylex.(*Lexer).e = &ForExpr{yyDollar[2].str, "", yyDollar[4].expr}
		}
	case 2:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:25
		{
			yylex.(*Lexer).e = &ForExpr{yyDollar[2].str, yyDollar[4].str, yyDollar[6].expr}
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:29
		{
			yylex.(*Lexer).e = yyDollar[1].expr
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:35
		{
			yyVAL.exprs = nil
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:39
		{
			yyVAL.exprs = []Expr{yyDollar[1].expr}
		}
	case 6:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:43
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].expr)
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:49
		{
			yyVAL.expr = &LitExpr{yyDollar[1].lit}
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:53
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:57
		{
			yyVAL.expr = &BinOpExpr{"+", yyDollar[1].expr, yyDollar[3].expr}
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:61
		{
			yyVAL.expr = &BinOpExpr{"-", yyDollar[1].expr, yyDollar[3].expr}
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:65
		{
			yyVAL.expr = &BinOpExpr{"*", yyDollar[1].expr, yyDollar[3].expr}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:69
		{
			yyVAL.expr = &BinOpExpr{"/", yyDollar[1].expr, yyDollar[3].expr}
		}
	case 13:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:73
		{
			yyVAL.expr = &CallExpr{yyDollar[1].str, yyDollar[3].exprs}
		}
	case 14:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:77
		{
			yyVAL.expr = &MethodCallExpr{Lhs: yyDollar[1].expr, Name: yyDollar[3].str, Exprs: yyDollar[5].exprs}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:81
		{
			yyVAL.expr = &MemberExpr{Lhs: yyDollar[1].expr, Name: yyDollar[3].str}
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:85
		{
			yyVAL.expr = &ItemExpr{Lhs: yyDollar[1].expr, Index: yyDollar[3].expr}
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:89
		{
			yyVAL.expr = &IdentExpr{yyDollar[1].str}
		}
	}
	goto yystack /* stack new state and value */
}
