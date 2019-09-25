package engine

type Request struct {
	URL       string
	ParseFunc func([]byte) ParseResult
}

type ParseResult struct {
	Request []Request
	Items   []interface{}
}

func NilParser([] byte) ParseResult {
	return ParseResult{}
}
