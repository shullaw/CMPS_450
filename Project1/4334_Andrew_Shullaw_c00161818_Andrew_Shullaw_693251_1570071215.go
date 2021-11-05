// Author :: Andrew Shullaw c00161818
package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

//---CHARACTER CLASSES AND TOKENS--//
const (
	LETTER  = 0
	DIGIT   = 1
	UNKNOWN = 99
	SPACE   = -2

	INT_LIT   = 10
	IDENT     = 11
	ASSIGN_OP = 20
	ADD_OP    = 21
	SUB_OP    = 22
	MULT_OP   = 23
	DIV_OP    = 24
	L_PAREN   = 25
	R_PAREN   = 26
	EOF       = -1

	MAX_SIZE = 99
)

//---DECLARE VARS---//
var (
	path       string
	err        error
	char_class int
	lexeme     string
	next_char  rune
	next_token int
)

func main() {

	e := func(error) {
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			fmt.Println(err)
			os.Exit(42)
		}
	}

	file_name := "/front.in"
	open_file := func(string) *bufio.Reader {

		path, err = os.Getwd() // current working directory
		e(err)
		fmt.Printf("cwd: %s\n", path)

		path_to_file := path + file_name // file to read
		file, err := os.Open(path_to_file)
		e(err)
		reader := bufio.NewReader(file)

		return reader
	}

	reader := open_file(file_name)

	add_char := func() {
		lexeme += string(next_char)
	}

	get_char := func() {
		r, s, _ := reader.ReadRune()
		next_char = r
		if s == 0 {
			char_class = EOF
		} else {
			if unicode.IsLetter(next_char) {
				char_class = LETTER
			} else if unicode.IsDigit(next_char) {
				char_class = DIGIT
			} else {
				char_class = UNKNOWN
			}
		}

	}

	skip_space := func() {
		for unicode.IsSpace(next_char) {
			get_char()
		}
	}

	look_up := func() {
		switch next_char {
		case '(':
			add_char()
			next_token = L_PAREN
		case ')':
			add_char()
			next_token = R_PAREN
		case '+':
			add_char()
			next_token = ADD_OP
		case '-':
			add_char()
			next_token = SUB_OP
		case '*':
			add_char()
			next_token = MULT_OP
		case '/':
			add_char()
			next_token = DIV_OP
		default:
			add_char()
			next_token = EOF
		}
	}

	lex := func() {
		lexeme = " "
		skip_space()
		switch char_class {
		case LETTER:
			add_char()
			get_char()
			for char_class == LETTER || char_class == DIGIT {
				add_char()
				get_char()
			}
			next_token = IDENT
		case DIGIT:
			add_char()
			get_char()
			for {
				add_char()
				get_char()
				if char_class != DIGIT {
					break
				}
			}
			next_token = INT_LIT
		case UNKNOWN:
			look_up()
			get_char()
		case EOF:
			next_token = EOF
			lexeme = "EOF"
		}
		fmt.Printf("Next token is: %d, Next lexeme is %s\n", next_token, lexeme)
	}

	//-----------------------------------------------------------------------//
	get_char()
	for next_token != EOF {
		lex()
	}

}
