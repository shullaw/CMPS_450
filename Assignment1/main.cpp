//  main.cpp (C++ version)
//  sebestaScannerCpp

//#include <string>
//#include <cstdio>
//#include <cctype>
#include <iostream>  // I/O
#include <fstream>   // file I/O
#include <iomanip>   // format manipulation
#include <unistd.h>  // unix standard header
//#include <std_lib_facilities.h>

using namespace std;

int    charClass;
string lexeme;  // In contrast to C, C++ has a true string data type
char   nextChar;
int    token;
int    nextToken;
ifstream in_fp; // input file stream object which has some methods

void my_getChar();
int  lex();

// C++ has true constants, whereas C simulates
// comments via preprocessor define directives
const int LETTER  =  0;
const int DIGIT   =  1;
const int UNKNOWN = 99;

const int INT_LIT     = 10;
const int IDENT       = 11;
const int ADD_OP      = 21;
const int SUB_OP      = 22;
const int MULT_OP     = 23;
const int DIV_OP      = 24;
const int LEFT_PAREN  = 25;
const int RIGHT_PAREN = 26;

int main(int argc, const char * argv[]) // Same as C
{
    in_fp.open("front_cpp.in");
    if (in_fp.is_open()) {
//        cout << "front.in opened" << endl;
        my_getChar();
        do {
            lex();
        } while (nextToken!=EOF);
    }
    else {
        cout << "Cannot open front_cpp.in" << endl;
        char cwd[1024]; // Another string buffer w/ capacity = 1023 chars
        if (getcwd(cwd, sizeof(cwd)) != NULL) {
            cout << "Current working dir: " << cwd << endl;
        } // # C++ I/O is stream based and harder to understand
    }
    return 0;
}

void my_addChar() { // much improved in C++
    lexeme += nextChar;  // string concatenation
}

void my_getChar() {
    in_fp>>nextChar; // read the next character
    if (in_fp.eof())
        charClass = EOF;
    else {
        if (isalpha(nextChar))
            charClass = LETTER;
        else if (isdigit(nextChar))
            charClass = DIGIT;
        else charClass = UNKNOWN;
    }
//    cout << nextChar<<endl;  // for debugging
}

void getNonBlank() {
    while (isspace(nextChar))
        my_getChar();
}

int lookup(char ch) { // Same as C version
    switch (ch) {
        case '(':
            my_addChar();
            nextToken = LEFT_PAREN;
            break;
        case ')':
            my_addChar();
            nextToken = RIGHT_PAREN;
            break;
        case '+':
            my_addChar();
            nextToken = ADD_OP;
            break;
        case '-':
            my_addChar();
            nextToken = SUB_OP;
            break;
        case '*':
            my_addChar();
            nextToken = MULT_OP;
            break;
        case '/':
            my_addChar();
            nextToken = DIV_OP;
            break;
        default:
            my_addChar();
            nextToken = EOF;
            break;
    }
    return nextToken;
}

int lex() {
    lexeme = "";
    getNonBlank();
    switch (charClass) {
        case LETTER:
            my_addChar();
            my_getChar();
            while (charClass == LETTER || charClass == DIGIT) {
                my_addChar();
                my_getChar();
            }
            nextToken = IDENT;
            break;
        case DIGIT:
            my_addChar();
            my_getChar();
            while (charClass == DIGIT) {
                my_addChar();
                my_getChar();
            }
            nextToken = INT_LIT;
            break;
        case UNKNOWN:
            lookup(nextChar);
            my_getChar();
            break;
        case EOF:
            nextToken = EOF;
            lexeme = "EOF";
    }
    cout<<"Next token is: "<<nextToken<<", Next lexeme is: "<<lexeme<<endl;
    return nextToken;
}
