package reqerrs

type Error interface {
	StatusCode() int
	ErrorCode() ErrCode
	Message() string
	Error() string
	Fragments() map[string]any
}
