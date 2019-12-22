package configmodel


type target struct {
	Key   string
	Type  string
	Value string
}

type operation struct {
	Key   string
	Type  string
	Value string
}

type option struct {
	Cookie string
	Proxy  string
}
type replace struct {
	Target string
	From   string
	To     string
}

type transform struct {
	Target string
	From   string
	To     string
}

type after struct {
	Transform transform
	Replace   replace
}

type before struct {
	Option  option
	Replace replace
}

type action struct {
	Target    target
	Operation operation
	Before    before
	After     after
	Return    string
}

type Actions struct {
	Action []action
}
