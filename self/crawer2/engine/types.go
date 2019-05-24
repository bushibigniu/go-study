package engine


//定义内容
type Request struct {
	Url string
	ParserFunc func([]byte) ParseResult
}
/*
<a href="http://www.zhenai.com/zhenghun/xuancheng" data-v-5e16505f>宣城</a>
	<a href="http://www.zhenai.com/zhenghun/xuchang" data-v-5e16505f>许昌</a>
		<a href="http://www.zhenai.com/zhenghun/xuhui" data-v-5e16505f>徐汇</a>
*/
type ParseResult struct {
	Requests []Request
	Items []interface{}
}

func NilParse([]byte)  ParseResult{
	return ParseResult{}
}
