package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
)

const (
	/* Character classes */
	/* Use preprocessor directives
	   to emulate symbolic constants */
	LETTER  = 0
	DIGIT   = 1
	UNKNOWN = 99
	SPACE   = -2

	/* Token codes */
	INT_LIT     = 10
	IDENT       = 11
	ASSIGN_OP   = 20
	ADD_OP      = 21
	SUB_OP      = 22
	MULT_OP     = 23
	DIV_OP      = 24
	LEFT_PAREN  = 25
	RIGHT_PAREN = 26
	EOF         = -1
)

func letter(r byte) bool {
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
		// fmt.Printf("LETTER: %s\n", string(r))
		return true
	}

	return false
}

func digit(r byte) bool {

	if r >= '0' && r <= '9' {
		// fmt.Printf("DIGIT: %s\n", string(r))
		return true
	}

	return false
}

type lexeme struct {
	PATH          string
	FILE          os.File
	LEXEMES       []byte // CAPITALIZE TO EXPORT
	CurrentChar_b byte
	CurrentChar_s string
	Token         int
	CHARCLASS     int
	Err           error
	MAX           int
	POSITION      int
}

// func (lex *lexeme) ADDCHAR() {
// 	if len(lex.LEXEMES) <= 98 {
// 		lex.LEXEMES = append(lex.LEXEMES, lex.CurrentChar_b)
// 	} else {
// 		fmt.Println("Error: lexeme is too long.")
// 		fmt.Println(lex.LEXEMES)
// 		os.Exit(0)
// 	}
// }
func (lex *lexeme) GETCHAR() {

	if letter(lex.CurrentChar_b) {
		lex.CHARCLASS = LETTER
		fmt.Printf("GETCHAR: LETTER %s\n", string(lex.CurrentChar_b))
	} else if digit(lex.CurrentChar_b) {
		lex.CHARCLASS = DIGIT
		fmt.Printf("GETCHAR: DIGIT %s\n", string(lex.CurrentChar_b))
	} else {
		lex.CHARCLASS = UNKNOWN
		fmt.Printf("GETCHAR: UNKNOWN %s\n", string(lex.CurrentChar_b))
	}
}
func (lex *lexeme) READCHAR() {
	// for i, c := range lex.LEXEMES {
	// 	fmt.Println(i, ":", c)
	// }
	fmt.Println("READCHAR: ")
	if lex.POSITION < lex.MAX {
		lex.CurrentChar_b = lex.LEXEMES[lex.POSITION]
	} else {
		lex.Token = EOF
	}
}

func (lex *lexeme) LOOKUP() {
	// fmt.Printf("LOOKUP: '%s'\n", string(lex.CurrentChar_b))
	switch lex.CurrentChar_b {
	case '(':
		// lex.READCHAR()
		lex.Token = LEFT_PAREN
		// fmt.Println("LEFT_PAREN")
	case ')':
		// lex.READCHAR()
		lex.Token = RIGHT_PAREN
		// fmt.Println("RIGHT_PAREN")
	case '+':
		// lex.READCHAR()
		lex.Token = ADD_OP
	case '-':
		// lex.READCHAR()
		lex.Token = SUB_OP
	case '*':
		// lex.READCHAR()
		lex.Token = MULT_OP
	case '/':
		// lex.READCHAR()
		lex.Token = DIV_OP
	default:
		lex.Token = EOF
	}
}

func (lex *lexeme) LEX() {

	for lex.POSITION < lex.MAX {
		if lex.POSITION == lex.MAX {
			os.Exit(0)
		}
		lex.POSITION++
		lex.READCHAR()
		lex.GETCHAR()
		var sb strings.Builder
		switch lex.CHARCLASS {
		case LETTER:
			for lex.CHARCLASS == LETTER || lex.CHARCLASS == DIGIT {
				sb.WriteString(string(lex.CurrentChar_b))
				if lex.POSITION < lex.MAX {
					lex.POSITION++
					lex.READCHAR()
					lex.GETCHAR()
				}
			}
			lex.Token = IDENT
			lex.POSITION--
			fmt.Printf("LETTER: Next token is: %d. Next lexeme is %s\n", lex.Token, sb.String())
			// sb.Reset()
			// lex.READCHAR()
		case DIGIT:
			for lex.CHARCLASS == DIGIT {
				sb.WriteString(string(lex.CurrentChar_b))
				if lex.POSITION < lex.MAX {
					lex.POSITION++
					lex.READCHAR()
					lex.GETCHAR()
				}
			}
			lex.Token = INT_LIT
			// sb.WriteString(string(lex.CurrentChar_b))
			lex.POSITION--
			fmt.Printf("DIGIT: Next token is: %d. Next lexeme is %s\n", lex.Token, sb.String())
			// lex.RESET()
			// lex.READCHAR()
		case UNKNOWN:
			lex.LOOKUP()
			sb.WriteString(string(lex.CurrentChar_b))
			fmt.Printf("UKNOWN: Next token is: %d. Next lexeme is %s\n", lex.Token, sb.String())
			// sb.Reset()
			// lex.RESET()
		default:
			// fmt.Printf("Next token is: %d. Next lexeme is %s\n", lex.Token, strconv.Itoa(int(lex.CurrentChar_b)))
			os.Exit(0)
		}
	}
}
func (lex *lexeme) RESET() {
	lex.CHARCLASS = -42
}
func main() {

	path, err := os.Getwd() // get current working dir
	if err != nil {
		log.Println(err)
		os.Exit(-42)
	}
	fmt.Println(path)

	path_file := path + "\\front.in" // add file name to path (file should be in cwd)
	f, err := os.Open(path_file)
	// f, err := ioutil.ReadFile(path_file)

	if err != nil { // if opening file results in error stop
		log.Println(err)
		os.Exit(-42)
	} else {
		fmt.Printf("%s\\front.in\n", path)
	}
	MAX := 99
	lex := lexeme{
		FILE:          *f,
		LEXEMES:       make([]byte, 0, MAX),
		CurrentChar_b: 0,
		CurrentChar_s: "",
		Token:         0,
		CHARCLASS:     0,
		Err:           err,
		MAX:           MAX,
		POSITION:      -1,
	}
	text, err := ioutil.ReadFile(lex.FILE.Name())
	for _, c := range text {
		if unicode.IsSpace(rune(c)) {
		} else {
			lex.LEXEMES = append(lex.LEXEMES, c)
		}
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lex.MAX = len(lex.LEXEMES)

	lex.LEX()
}
