package gockle

import (
	"fmt"

	"github.com/maraino/go-mock"
)

var mySession = &SessionMock{}

func Example_dump() {
	var rows, _ = mySession.QuerySliceMap("select * from users")

	for _, row := range rows {
		fmt.Println(row)
	}
}

func Example_insert() {
	mySession.QueryExec("insert into users (id, name) values (123, 'me')")
}

func Example_print() {
	var i = mySession.QueryIterator("select * from users")

	for done := false; !done; {
		var m = map[string]interface{}{}

		done = i.ScanMap(m)

		fmt.Println(m)
	}
}

func init() {
	var i = &IteratorMock{}

	i.When("ScanMap", mock.Any).Call(func(m map[string]interface{}) bool {
		m["id"] = 123
		m["name"] = "me"

		return false
	})

	i.When("Close").Return(nil)

	mySession.When("QueryExec", mock.Any).Return(nil)
	mySession.When("QueryIterator", mock.Any).Return(i)
	mySession.When("QueryScanMap", mock.Any).Return(map[string]interface{}{"id": 1, "name": "me"}, nil)
}
