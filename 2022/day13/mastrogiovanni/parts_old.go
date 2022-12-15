package main

import "fmt"

type TokenType int

const (
	Unknown = iota
	Open
	Close
	Comma
	Value
	Eof
)

type Token struct {
	Start int
	End   int
	Type  TokenType
	Value int
}

func (t Token) String() string {
	if t.Type == Value {
		return fmt.Sprintf("%d", t.Value)
	} else {
		return DecodeType(t.Type)
	}
}

func DecodeType(ttype TokenType) string {
	switch ttype {
	case Open:
		return "Open"
	case Close:
		return "Close"
	case Comma:
		return "Comma"
	case Value:
		return "Value"
	case Eof:
		return "Eof"
	case Unknown:
		return "Unknown"
	}
	return ""
}

func DumpTokens(line string) {
	for i := 0; ; {
		token := NextToken(line, i)
		if token.Type == Eof {
			break
		}
		if token.Type == Value {
			fmt.Printf("%s (start: %d, end: %d) => %d\n", line[token.Start:token.End], token.Start, token.End, token.Value)
		} else {
			fmt.Printf("%s (start: %d, end: %d)\n", DecodeType(token.Type), token.Start, token.End)
		}
		i = token.End
	}
}

func NextToken(line string, index int) Token {
	tot := 0
	ttype := Unknown
	for i := index; i < len(line); i++ {
		if line[i] >= '0' && line[i] <= '9' {
			tot = tot*10 + int(line[i]-'0')
			ttype = Value
		} else if line[i] == '[' {
			if ttype == Value {
				return Token{
					Start: index,
					End:   i,
					Type:  Value,
					Value: tot,
				}
			} else {
				return Token{
					Start: index,
					End:   i + 1,
					Type:  Open,
					Value: -1,
				}
			}
		} else if line[i] == ']' {
			if ttype == Value {
				return Token{
					Start: index,
					End:   i,
					Type:  Value,
					Value: tot,
				}
			} else {
				return Token{
					Start: index,
					End:   i + 1,
					Type:  Close,
					Value: -1,
				}
			}
		} else if line[i] == ',' {
			if ttype == Value {
				return Token{
					Start: index,
					End:   i,
					Type:  Value,
					Value: tot,
				}
			} else {
				return Token{
					Start: index,
					End:   i + 1,
					Type:  Comma,
					Value: -1,
				}
			}
		}
	}
	if ttype == Unknown {
		return Token{
			Start: -1,
			End:   -1,
			Type:  Eof,
			Value: -1,
		}
	} else {
		return Token{
			Start: index,
			End:   len(line),
			Type:  Value,
			Value: tot,
		}
	}
}

func Reverse(tokens []Token) []Token {
	for i, j := 0, len(tokens)-1; i < j; i, j = i+1, j-1 {
		tokens[i], tokens[j] = tokens[j], tokens[i]
	}
	return tokens
}

func GetTokens(line string) []Token {
	tokens := make([]Token, 0)
	for i := 0; ; {
		token := NextToken(line, i)
		if token.Type == Eof {
			break
		}
		i = token.End
		tokens = append(tokens, token)
	}
	Reverse(tokens)
	return tokens
}

func IsArray(tokens []Token) bool {
	if len(tokens) == 0 {
		return true
	}
	return tokens[0].Type == Open
}

func Parse(tokens []Token) [][]Token {
	level := 0
	res := make([][]Token, 0)
	var current []Token
	for i := 0; i < len(tokens); i++ {
		if tokens[i].Type == Open {
			level++
			// fmt.Println(DecodeType(tokens[i].Type), level, current)
			if level == 1 {
				current = make([]Token, 0)
			} else {
				current = append(current, tokens[i])
			}
		} else if tokens[i].Type == Close {
			// fmt.Println(DecodeType(tokens[i].Type), level, current)
			if level > 1 {
				current = append(current, tokens[i])
			} else {
				res = append(res, current)
				return res
			}
			level--
		} else if tokens[i].Type == Comma {
			// fmt.Println(DecodeType(tokens[i].Type), level, current)
			if level == 1 {
				res = append(res, current)
				current = make([]Token, 0)
			} else {
				current = append(current, tokens[i])
			}
		} else {
			// fmt.Println(DecodeType(tokens[i].Type), level, current)
			current = append(current, tokens[i])
		}
	}
	return res
}

func CompareTokens(tokensLeft, tokensRight []Token, level int) bool {

	if len(tokensLeft) == 0 {
		fmt.Println("Left ended")
		return true
	}
	if len(tokensRight) == 0 {
		fmt.Println("Right ended")
		return false
	}

	lastLeft := tokensLeft[len(tokensLeft)-1]
	lastRight := tokensRight[len(tokensRight)-1]
	restLeft := tokensLeft[0 : len(tokensLeft)-1]
	restRight := tokensRight[0 : len(tokensRight)-1]

	fmt.Printf("%s = %s\n", lastLeft, lastRight)

	switch {
	case lastLeft.Type == Comma && lastRight.Type == Comma:
		// fmt.Println("Case 1")
		return CompareTokens(restLeft, restRight, level)
	case lastLeft.Type == Open && lastRight.Type == Open:
		// fmt.Println("Case 2")
		return CompareTokens(restLeft, restRight, level+1)
	case lastLeft.Type == Close && lastRight.Type == Close:
		// fmt.Println("Case 3")
		return CompareTokens(restLeft, restRight, level-1)
	case lastLeft.Type == Value && lastRight.Type == Value:
		if lastLeft.Value > lastRight.Value {
			// fmt.Println("Case 4.1")
			return false
		} else if lastLeft.Value < lastRight.Value {
			// fmt.Println("Case 4.2")
			return true
		} else {
			// fmt.Println("Case 4.3")
			return CompareTokens(restLeft, restRight, level)
		}
	case lastLeft.Type == Value && lastRight.Type == Open:
		fmt.Println("Case 5")
		res := append(tokensLeft[0:len(tokensLeft)-1], Token{
			Start: -1,
			End:   -1,
			Type:  Close,
			Value: -1,
		}, lastLeft, Token{
			Start: -1,
			End:   -1,
			Type:  Open,
			Value: -1,
		})
		return CompareTokens(res, tokensRight, level)
	case lastLeft.Type == Open && lastRight.Type == Value:
		fmt.Println("Case 6")
		res := append(tokensRight[0:len(tokensRight)-1], Token{
			Start: -1,
			End:   -1,
			Type:  Close,
			Value: -1,
		}, lastRight, Token{
			Start: -1,
			End:   -1,
			Type:  Open,
			Value: -1,
		})
		return CompareTokens(tokensLeft, res, level)
	case lastLeft.Type == Close && lastRight.Type != Close:
		fmt.Println("Case 7")
		level := 1
		for i := 0; i < len(tokensRight); i++ {
			if tokensRight[len(tokensRight)-1-i].Type == Open {
				level++
				fmt.Println("Case 7.1:", level)
			} else if tokensRight[len(tokensRight)-1-i].Type == Close {
				level--
				fmt.Println("Case 7.2:", level)
				if level == 0 {
					return CompareTokens(tokensLeft, tokensRight[0:len(tokensRight)-i], level)
				}
			}
		}
		fmt.Println("Should never be here 2")
	case lastLeft.Type != Close && lastRight.Type == Close:
		fmt.Println("Case 8")
		return false
	}
	fmt.Println("Should never be here")
	return false
}

func Compare1(left string, right string) bool {
	tokensLeft := GetTokens(left)
	tokensRight := GetTokens(right)
	return CompareTokens(tokensLeft, tokensRight, 0)
}
