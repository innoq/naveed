// HTTP plain text testing DSL
//
// inspired by TiddlyWeb's YAML-based tests:
// https://github.com/tiddlyweb/tiddlyweb/blob/master/test/httptest.yaml

package htptest

import "os"
import "bufio"

func Process(testcases *os.File) {
	scanner := bufio.NewScanner(testcases)

	lex := new(lineLexer)
	lex.tokens = make(chan string)
	lex.lineCount = 0
	go parse(lex.tokens)

	lex.state = start
	for scanner.Scan() {
		lex.lineCount++
		lex.line = scanner.Text()
		lex.state = lex.state(lex)
	}
}
