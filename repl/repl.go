package repl

import (
	"bufio"
	"fmt"
	"interpreterInGo/lexer"
	"interpreterInGo/token"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		_, _ = fmt.Fprintf(out, PROMPT)
		// 读取一行输入
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		// 读取token
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			// 打印token
			_, _ = fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
