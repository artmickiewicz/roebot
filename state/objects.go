package state

type Template struct {
	ID               int
	TargetChannel    string
	SourceMessagePtr MessagePtr
	TargetMessagePtr MessagePtr
	Text             string
}

type MessagePtr struct {
	ChatID    int64
	MessageID int
}

func (tpl Template) IsPosted() bool {
	return tpl.TargetMessagePtr.MessageID > 0
}

type State int

const Null State = 0
const Clean State = 1
const Added State = 2
const Updated State = 3
const Deleted State = 4
