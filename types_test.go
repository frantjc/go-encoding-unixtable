package unixtable_test

type Unixtable struct {
	One, Two string
}

type UnixtableTagged struct {
	One string `unixtable:"one"`
	Two string `unixtable:"two"`
}

type UnixtableOmitempty struct {
	One string `unixtable:"one,omitempty"`
	Two string `unixtable:",omitempty"`
}
