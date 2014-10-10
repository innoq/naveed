package htptest

import "strconv"
import "strings"

type lineLexer struct {
	line string
	lineCount int
	tokens chan string
	state stateFn
}

type stateFn func(lex *lineLexer) stateFn

func start(lex *lineLexer) stateFn {
	switch {
	default:
		return abort(lex)
	case lex.line == "": // skip
		return start
	case strings.HasPrefix(lex.line, "# "):
		lex.tokens <- "COMMENT"
		return start
	case strings.HasPrefix(lex.line, "> "):
		lex.tokens <- "REQUEST"
		return request
	}
}

func request(lex *lineLexer) stateFn {
	switch {
	default:
		return abort(lex)
	case lex.line == "":
		return response
	case strings.HasPrefix(lex.line, "> "):
		lex.tokens <- "REQUEST"
		return request
	case strings.HasPrefix(lex.line, "< "):
		lex.tokens <- "RESPONSE"
		return response
	}
}

func response(lex *lineLexer) stateFn {
	switch {
	default:
		return abort(lex)
	case lex.line == "":
		return start
	case strings.HasPrefix(lex.line, "< "):
		lex.tokens <- "RESPONSE"
		return response
	}
}

func abort(lex *lineLexer) stateFn {
	panic("invalid input at line #" + strconv.Itoa(lex.lineCount) + ": " + lex.line) // XXX: crude
}
