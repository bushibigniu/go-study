package engine


//定义内容
type Request struct {
	Url string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items []interface{}
}

func NilParse([]byte)  ParseResult{
	return ParseResult{}
}
