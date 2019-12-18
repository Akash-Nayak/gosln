package main

// Token represents a lexical token.
type Token int

const (
	// Unknown caracters that doesn't fit in any other group
	Unknown Token = iota
	// EOF - end of file
	EOF
	// WS - whitespace
	WS

	// IDENT Characters
	IDENT

	// ASTERISK *
	ASTERISK
	// COMMA ,
	COMMA
	// OPENPAREN (
	OPENPAREN
	// CLOSEPAREN )
	CLOSEPAREN
	// QUOTE "
	QUOTE
	// EQUAL =
	EQUAL

	// PROJECT Project begin
	PROJECT
	// ENDPROJECT Project end
	ENDPROJECT
)