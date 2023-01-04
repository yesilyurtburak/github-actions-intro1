package main

import "testing"

func TestGreeting(t *testing.T) {
	inputs := []struct {
		name   string
		result string
	}{
		{name: "Burak", result: "Hello Burak"},
		{name: "Ali", result: "Hello Ali"},
		{name: "Elif", result: "Hello Elif"},
		{name: "Murat", result: "Hello Murat"},
	}

	for _, item := range inputs {
		result := greeting(item.name)
		if result != item.result {
			t.Errorf("\ngreeting('%s') failed,\n\texpected -> %v, \n\tgot -> %v", item.name, item.result, result)
		} else {
			t.Logf("\ngreeting('%s') succeded,\n\texpected -> %v, \n\tgot -> %v", item.name, item.result, result)
		}
	}
}
