package gosln

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"
)

type Solution struct {
	Projects []Project
}

type Project struct {
	ID           string
	Name         string
	ProjectFile  string
	TypeGUID     string
	IsDependency bool
}

// Parser represents a parser.
type Parser struct {
	s   *Scanner
	buf struct {
		tok Token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

func (p *Parser) ParseString() (string, error) {
	tok, lit := p.scanIgnoreWhitespace()
	if tok != QUOTE {
		return lit, nil
	} else {
		var s string
		for {
			tok, lit := p.scan()
			if tok != QUOTE {
				s = s + lit
			} else {
				break
			}
		}
		return s, nil
	}
}

func (p *Parser) ParseProject() (Project, error) {
	var proj Project
	if ok, err := p.expect(OPENPAREN); !ok {
		return proj, err
	}
	proj.TypeGUID, _ = p.ParseString()
	if ok, err := p.expect(CLOSEPAREN, EQUAL); !ok {
		return proj, err
	}
	proj.Name, _ = p.ParseString()
	if ok, err := p.expect(COMMA); !ok {
		return proj, err
	}
	s, _ := p.ParseString()
	proj.ProjectFile = strings.Replace(s, `\`, string(filepath.Separator), -1)
	if ok, err := p.expect(COMMA); !ok {
		return proj, err
	}

	if ok, err := p.expect(OPENPAREN); !ok {
		return proj, err
	}
	ProjectDependencies, _ := p.ParseString()
	if ProjectDependencies == "ProjectDependencies" {
		proj.IsDependency = true
	}
	proj.ID, _ = p.ParseString()
	if ok, err := p.expect(ENDPROJECT); !ok {
		return proj, err
	}
	return proj, nil
}

func (p *Parser) expect(expected ...Token) (bool, error) {
	for _, exp := range expected {
		if tok, lit := p.scanIgnoreWhitespace(); tok != exp {
			return false, fmt.Errorf("unexpected token %q", lit)
		}
	}
	return true, nil
}

// Parse parses a SQL SELECT statement.
func (p *Parser) Parse() (Solution, error) {
	var sln Solution
	for {
		tok, _ := p.scanIgnoreWhitespace()
		switch tok {
		case EOF:
			p.unscan()
		case PROJECT:
			proj, _ := p.ParseProject()
			sln.Projects = append(sln.Projects, proj)
		}
		if tok == EOF {
			break
		}
	}
	return sln, nil
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.n = 1 }
