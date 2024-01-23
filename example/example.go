package example

type IDThing struct {
	ID string
	TB string
}

type Record struct {
	ID          IDThing `surreal:"id"`
	RequestedAt string  `surreal:"requested_at"`
	RespondedAt string  `surreal:"responded_at"`
}
