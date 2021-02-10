package misc

type ID int

var IdCounter ID = 0

func (counter *ID) Next() ID {
	val := *counter
	(*counter)++
	return val
}
