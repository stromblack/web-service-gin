package statusmodel

type StatusModel int64
type strvalue string

const (
	InvalidParameter       StatusModel = 2002
	RequestHeaderEmpty     StatusModel = 2003
	RequestHeaderIncorrect StatusModel = 2004
	InvalidToken           StatusModel = 2005
	ErrorInDatabase        StatusModel = 2006
)

// more detail implement https://www.golangprograms.com/golang/interface-type/
// implement the String method for the StatusModel type
func (s StatusModel) String() string {
	switch s {
	case InvalidParameter:
		return "Invalid Parameter"
	case RequestHeaderEmpty:
		return "Request Header Empty"
	case RequestHeaderIncorrect:
		return "Request Header Incorrect"
	case InvalidToken:
		return "Invalid Token"
	case ErrorInDatabase:
		return "Error In Database"
	}
	return "not found"
}
