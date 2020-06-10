package parser

import(
	"fmt"
	"strings"
)

const (
	DEBUG = false
	IDENT = 2
)

var (
	level = 0
)

/*
PROCEDURE debug(msg: ARRAY OF CHAR; level: INTEGER);
	VAR sps: ARRAY 64 OF CHAR; i: INTEGER;
BEGIN i := 0;
	IF (DEBUG) THEN
		WHILE level > 0 DO sps [i] := " "; INC(i); DEC(level); END;
		sps[i] := 0X; (* END *);
		Out.String(sps); Out.String(msg); Out.Ln;
	END;
END debug;
*/
func debug(msg string, level *int) {
	if DEBUG {
		spaces := strings.Repeat(" ", *level)
		fmt.Printf("%s%s\n", spaces, msg)
	}
}

//
//PROCEDURE incL(VAR level: INTEGER); BEGIN level := level + 4; END incL;
//
func incLevel(level *int) {
	*level += IDENT
}

