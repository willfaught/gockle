package gockle

import (
	"fmt"

	"github.com/maraino/go-mock"
)

var mySession = &SessionMock{}

func ExampleBatch() {
	var b = mySession.Batch(BatchLogged)

	b.Add("insert into users (id, name) values (123, 'me')")
	b.Add("insert into users (id, name) values (456, 'you')")

	b.Exec()
}

func ExampleIterator() {
	var i = mySession.Iterate("select * from users")

	for done := false; !done; {
		var m = map[string]interface{}{}

		done = i.ScanMap(m)

		fmt.Println(m)
	}
}

func ExampleSession() {
	var rows, _ = mySession.ScanMapSlice("select * from users")

	for _, row := range rows {
		fmt.Println(row)
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

	mySession.When("Exec", mock.Any).Return(nil)
	mySession.When("Iterate", mock.Any).Return(i)
	mySession.When("ScanMap", mock.Any).Return(map[string]interface{}{"id": 1, "name": "me"}, nil)
}
