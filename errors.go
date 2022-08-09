package main

type GithubError struct {
	message string
	code    int
	w_err   error
}

func (g *GithubError) Error() string {
	return g.message
}

func (g *GithubError) GetFullError() error {
	return g.w_err
}
