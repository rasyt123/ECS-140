package term

import (
	"errors"
 // "strconv"
)

// ErrParser is the error value returned by the Parser if the string is not a
// valid term.
// See also https://golang.org/pkg/errors/#New
// and // https://golang.org/pkg/builtin/#error
var ErrParser = errors.New("parser error")

//
// <start>    ::= <term> | \epsilon
// <term>     ::= ATOM | NUM | VAR | <compound>
// <compound> ::= <functor> LPAR <args> RPAR
// <functor>  ::= ATOM
// <args>     ::= <term> | <term> COMMA <args>
//

// Parser is the interface for the term parser.
// Do not change the definition of this interface.
type Parser interface {
	Parse(string) (*Term, error)
}


type ParserImpl struct {
	lex     *lexer
	peekTok *token
}




// NewParser creates a struct of a type that satisfies the Parser interface.
func NewParser() Parser {
	return &ParserImpl{}
}


func (p *ParserImpl) nextToken() (*token, error) {
	if tok := p.peekTok; tok != nil {
		p.peekTok = nil
		return tok, nil
	}

	tok, err := p.lex.next()
	if err != nil {
		return nil, ErrParser
	}

	return tok, nil
}

// Helper function which puts a token back as the next token.
func (p *ParserImpl) backToken(tok *token) {
	p.peekTok = tok
}

// Helper function to peek the next token.
func (p *ParserImpl) peekToken() (*token, error) {
	tok, err := p.nextToken()
	if err != nil {
		return nil, ErrParser
	}

	p.backToken(tok)

	return tok, nil
}


//Example declarations 
//&Term{Typ: TermAtom, Literal: "foo"}


//f := &Term{Typ: TermAtom, Literal: "f"}
//	X := &Term{Typ: TermVariable, Literal: "X"}
//	return "f(X)", &Term{Typ: TermCompound, Functor: f, Args: []*Term{X}}
//}

func (p *ParserImpl) Parse(input string) (*Term, error) {


	p.lex = newLexer(input)

	//<start> := <term> | epsilon 
	//peektoken 
	









}

//start -> term | EPSILON
//term ->  NUM | VAR | ATOM ender
//ender -> EPSILON | ( args )
//args -> term argsprime
//argsprime -> EPSILON  | , args


//foo vs foo(X, 5 )
func (p *ParserImpl) termNT() (*Term, error) {
	//Praisnt term <term> := | NUM | VAR | ATOM <ender>
	tok, err := p.nextToken()
	if err != nil {
		return nil, ErrParser
	}

	switch tok.typ {
		case tokenAtom:
			//Parse Ender 
			//term -> atom <ender>
			ender, err := p.enderNT()
			if err != nil {
				return nil, ErrParser
			}
			return ender, nill
			  
		case tokenNumber:
			//term -> <num>
			return &Term{Typ: termNum, Literal: tok.literal}
		case tokenVariable:
			//<term> -> <variable>
			return &Term{Typ: termVariable, Literal: tok.literal}
		default: 
			return nil, ErrParser
			

	}

}
//foo(4)
//p.peektoken  parse
//p.nextoken   term 
//p.peektoken  ender
//p.nexttoken  args
//current token is ()


func (p *ParserImpl) EnderNT(enderstr string) (*Term, error) {
	//ender -> EPSILON | ( args )
	//two choices, we peek at the next token

	//get the next token and if its epsilon or empty 
	//we that means we parsed an ATOM only 
	tok, err := p.peekToken()



	switch tok.typ {
			case tokenLpar:
					//ender -> ( args )
					//
					//crrate a term with compound probably? 
					//add result of ArgsNt to this compound 
			

			case tokenEOF:
				return &Term{Typ: TermAtom, Literal: enderstr}
			case tokenRpar:
				//<ender> -> epsilon
				return &Term{Typ: TermAtom, Literal: enderstr}
			case tokenComma:
				//<ender> -> epsilon 
				return &Term{Typ: TermAtom, Literal: enderstr}
			default:
				return nil, ErrParser
	}

	
}

//peektok has atom 
/*
func termWithoutSharingTest8() (string, *Term) {
	foo := &Term{Typ: TermAtom, Literal: "foo"}
	a := &Term{Typ: TermAtom, Literal: "a"}
	X := &Term{Typ: TermVariable, Literal: "X"}
	return "foo  ( a , X )", &Term{Typ: TermCompound, Functor: foo, Args: []*Term{a, X}}
}

*/

//foo(4)
//p.peektoken  parse
//p.nextoken   term 
//p.peektoken  ender
//p.nexttoken  args
//current token is ()

//foo( 4 )
func (p *ParserImpl) ArgsNT() (*Term, error) {
	//<args> -> <term> <argsprime>
	//we parse both term and args prime here 
	//
	tok, err := p.nextToken()
	if err != nil {
		return nil, ErrParser
	}
	switch tok.typ {
		case tokenAtom, tokenNumber, tokenVariable:
			//parse <term>
			term, err := p.termNT()
			if err != nil {
				return nil, ErrParser
			}

			//parse <argsprime>
			expr, err := p.ArgsPrimeNT()
			if err != nil {
				return nil, ErrParser
			}

			//args prime should be added to term, add this later 
		default:
			return nil, ErrParser	
			
	}
}



func (p *ParserImpl) ArgsPrimeNT() (*Term, error) {
	//argsprime -> epsiolon | , args 
	//two choices, peek token 

	tok, err := p.peekToken()
	if err != nil {
		return nil, ErrParser
	}

	switch tok.typ {
			case tokenRpar:
				//argsprime -> epsilon
				return &Term{Typ: nil, Literal: nil}
			case tokenComma:
				//argsprime -> , args		
				args, err := p.ArgsNT()
				if err != nil {
					return nil, ErrParser
				}
				return args
			default:
				return nil, ErrParser	
	}
}



//parse start implicit 
 //parseTermFunction
 //Parse Compound 
 //Parse args
 //parse argsprime 

