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

/*
PROCEDURE error(errNo: INTEGER);
	PROCEDURE msg(s: ARRAY OF CHAR); 
	BEGIN Out.String("error("); Out.Int(errNo,1); Out.String(") - "); Out.String(s); Out.Ln;
	END msg;
BEGIN 
	CASE errNo OF
		1: msg("Left paren is missing");  | (* could never happen *)
		2: msg("Right paren is missing"); 
	END;
END error;
*/
func printError(errNo int) {

	msg := func(s string) {
		fmt.Printf("error(%v) - %s\n", errNo, s)
	}
	
	txt, found := errors[errNo]
	if found {
		msg(txt)
	} else {
		msg("Unknown error no:" + string(errNo))
	}

}

