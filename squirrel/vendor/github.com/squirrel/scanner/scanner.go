package scanner

import (
	"bytes"
	"fmt"
	"io"
)

const (
	Lparen 			= 0		// 	(
	Symbol 			= 10	// 	'foo
	String 			= 11 	// 	"bar"
	Number 			= 12 	// 	-1234.567e-10 (DEC64 - Douglas Crockford)
	Quote 			= 20 	// 	'
	Backquote 		= 21	// 	`
	Unquote			= 22    //  ,
	UnquoteSplicing = 23	//  ,@
	Rparen 			= 100  	//	)
	Other 			= 255	// 	EOF 
)
	
var (
	Ch 		rune
	Ch2     rune
	Size	int
	Err		error
	Sym 	int
	Id		[IdLen]rune
	R		*bytes.Reader
	//error	bool
	errpos	int64
	Debug	bool
)

func IdStr() string {
	return asStr(Id)
}

// NextCh reads next rune
// At the end Err is set to io.EOF
func NextCh() {
	Ch, Size, Err = R.ReadRune()
}

// PeekCh peek next rune
func PeekCh() {
	Ch2, Size, Err = R.ReadRune()	// TODO: CHECK ????
	err := R.UnreadRune()
	if err != nil {
		panic(err)
	}
}

func printSym() {
	fmt.Printf("sym: %3d  ", Sym)
	switch {
		case Sym == Symbol:
			fmt.Printf("symbol: %v", Id)
		case Sym == String:
			fmt.Printf("string: %v", Id)
		case Sym == Lparen:
			fmt.Printf("lparen: (")
		case Sym == Rparen:
			fmt.Printf("rparen: )")
		case Sym == Quote:
			fmt.Printf("quote: '")
		case Sym == Backquote:
			fmt.Printf("backquote: `")
		case Sym == Unquote:
			fmt.Printf("unquote: ,")
		case Sym == UnquoteSplicing:
			fmt.Printf("unquote-splicing: ,@")
		case Sym == Other:
			fmt.Printf("other")
	}
	fmt.Println("")
}

func Mark(msg string) {
	pos := R.Size() - int64(R.Len()) // TODO: Check
	if pos > errpos {
		fmt.Printf(" pos %v - %v", pos, msg)
		fmt.Println()
	}
}

func GetSym() {

	inc := func(i *int) { 
		*i++
	}
	
	// readSymbol reads symbols like 'foo
	//
	readSymbol := func() {
		i := 0
		for ;; {
			Id[i] = Ch
			inc(&i)
			NextCh()
			if isNotLetter(Ch) {
				break
			}
		}
		Id[i] = '\x00' // = 0 see https://golang.org/ref/spec#Rune_literals
	}	

	// readString reads string symbols
	//
	readString := func() {
		NextCh()
		i := 0
		for ;Ch != '"' && Ch > ' '; {
			Id[i] = Ch
			inc(&i)
			NextCh()
		}
		if Ch <= ' ' {
			Mark("No control chars allowed in string")
		}
		Id[i] = '\x00'
		NextCh()
	}
	
	// readNumber reads a number symbol
	//
	readNumber := func() {
	
		i := 0
		for ;('0' <=Ch && Ch <= '9') || Ch == '.'; {
			Id[i] = Ch
			inc(&i)
			NextCh()
		}		
		Id[i] = '\x00'
	}
		
	// Skips blanks like \n, \r, SPACE, TAB
	//
	for ; Err != io.EOF && Ch <= ' ' ; {
		NextCh()
	}
		
	// Read next symbol
	//
	switch {
		case 'a' <= Ch && Ch <= 'z' || 
		     'A' <= Ch && Ch <= 'Z':
			Sym = Symbol
			readSymbol()
			
		// TODO: -1234.45e-12
		case '0' <= Ch && Ch <='9':		
			Sym = Number
			readNumber()
		
		case '\'' == Ch:
			Sym = Quote
			NextCh()
			
		case '`' == Ch:
			Sym = Backquote
			NextCh()
						
		case ',' == Ch:
			PeekCh()
			if '@' == Ch2 {
				Sym = UnquoteSplicing
				NextCh()
			} else {
				Sym = Unquote
			}
			NextCh()
			
		// TODO: UnquoteSplicing
		// ,@
		
		case '"' == Ch:
			Sym = String
			readString()
		
		case '(' == Ch:
			Sym = Lparen
			NextCh()
			
		case ')' == Ch:
			Sym = Rparen
			NextCh()
		
		default:
			Sym = Other // EOT
			NextCh()
	}

}

// Init inits the scanner with given string
func Init(b []byte) {	
	R = bytes.NewReader(b)
	NextCh()
}

