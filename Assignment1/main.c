//  main.c
//  sebestaScanner (p. 172 from Concepts of PLs, Robert Sebesta, 10th ed)
// Additional comments added 1/30/19. A. Maida
// Written in a subset of C known as clean C.
// This means it should be acceptable to a C++ compiler.

#include <stdio.h>
#include <ctype.h>
#include <unistd.h>
//  stdio:  printf, fopen
//  ctype:  isalpha, isdigit, isspace
//  unistd: getcwd

/* global variables */
int  charClass;

// A C-string is a character array terminated by the null char
// That means lexeme has a capacity of 99 chars.
char lexeme[100]; // Buffer for character array. Assume initialized to garbage.
char nextChar;

int  lexLen;
int  token;
int  nextToken;
FILE *in_fp, *fopen();

/* Function declarations */
void addChar();
void getChar();
void getNonBlank();
int  lex();

/* Character classes */
/* Use preprocessor directives 
to emulate symbolic constants */
#define LETTER 0
#define DIGIT 1
#define UNKNOWN 99

/* Token codes */
#define INT_LIT 10
#define IDENT 11
#define ASSIGN_OP 20
#define ADD_OP 21
#define SUB_OP 22
#define MULT_OP 23
#define DIV_OP 24
#define LEFT_PAREN 25
#define RIGHT_PAREN 26

/* main driver */
int main(int argc, const char * argv[]) // # of args, arg vector
{
    char cwd[1024];
    if ((in_fp = fopen("front_c.in","r")) == NULL) {
        printf("ERROR: cannot open front.in \n"); // print formated
        if (getcwd(cwd, sizeof(cwd)) != NULL)
            printf("Current working dir: %s\n", cwd);
        else
            perror("getcwd() error");
        return -1;
    }
    else {
        getChar();
        do {
            lex();
        } while (nextToken != EOF);
    }

    return 0; // successful termination
}

// Adds nextChar to the lexeme string.
void addChar() {
    if (lexLen <= 98) { // can't hold more than 99 chars
        lexeme[lexLen++] = nextChar;
        lexeme[lexLen] = 0;  // null-terminated string
    }
    else
        printf("Error: lexeme is too long.");
}

// Reads the next char and saves it in nextChar
void getChar() {
    if ((nextChar = getc(in_fp)) != EOF) {
        if (isalpha(nextChar))
            charClass = LETTER;
        else if (isdigit(nextChar))
            charClass = DIGIT;
        else charClass = UNKNOWN;
    }
    else
        charClass = EOF;
}

// Skip over blanks
void getNonBlank() {
    while (isspace(nextChar))
        getChar();
}

int lookup(char ch) {
    switch (ch) {
        case '(':
            addChar();
            nextToken = LEFT_PAREN;
            break;
        case ')':
            addChar();
            nextToken = RIGHT_PAREN;
            break;
        case '+':
            addChar();
            nextToken = ADD_OP;
            break;
        case '-':
            addChar();
            nextToken = SUB_OP;
            break;
        case '*':
            addChar();
            nextToken = MULT_OP;
            break;
        case '/':
            addChar();
            nextToken = DIV_OP;
            break;
        default:
            addChar();
            nextToken = EOF;
            break;
    }
    return nextToken;
}

/* Corresponds to FSA in chapter 4 of Sebesta */
int lex() {
    lexLen = 0;
    getNonBlank();
    switch (charClass) {
        case LETTER:
            addChar();
            getChar();
            while (charClass == LETTER || charClass == DIGIT) {
                addChar();
                getChar();
            }
            nextToken = IDENT;
            break;
        case DIGIT:
            addChar();
            getChar();
            while (charClass == DIGIT) {
                addChar();
                getChar();
            }
            nextToken = INT_LIT;
            break;
        case UNKNOWN:
            lookup(nextChar);
            getChar();
            break;
        case EOF:
            nextToken = EOF;
            lexeme[0] = 'E';
            lexeme[1] = 'O';
            lexeme[2] = 'F';
            lexeme[3] = 0;
    }
    printf("Next token is: %d, Next lexeme is %s\n", nextToken, lexeme);
    return nextToken;
}
