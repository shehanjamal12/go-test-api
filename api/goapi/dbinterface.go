package goapi

type Dbinter interface {
	Getitem() []Items
	Add(items Items)
	Deleteitem(name string)
}
