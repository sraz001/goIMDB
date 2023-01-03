package goIMDB

type ErrRedirect struct {
	msg string
}

func (e ErrRedirect) Error() string {
	return e.msg
}

type ErrTitleNotFound struct {
	msg string
}

func (e ErrTitleNotFound) Error() string {
	return e.msg
}

// ErrInvalidImdbCode - bad imdb code (must be tt....)
type ErrInvalidImdbCode struct {
	msg string
}

func (e ErrInvalidImdbCode) Error() string {
	return e.msg
}

type ErrInvalidQuery struct {
	msg string
}

func (e ErrInvalidQuery) Error() string {
	return e.msg
}

type ErrEpisodeTitle struct {
	msg string
}

func (e ErrEpisodeTitle) Error() string {
	return e.msg
}

type ErrBadDateFormat struct {
	msg string
}

func (e ErrBadDateFormat) Error() string {
	return e.msg
}
