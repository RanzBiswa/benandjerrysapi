package loggers

import (
	"errors"
	"testing"
)

func TestTypeDesc(t *testing.T) {

	ts := testStruct{
		value: "Test",
	}

	cs := childStruct{
		value1: "Child",
		value2: 100,
	}

	ts.childStruct = cs

	LogError("TESTEVT", "TestMethod", ts, "", nil)
	LogWarning("TESTEVT", "TestMethod", ts, "", nil)
	LogData("TESTEVT", "TestMethod", ts, "", nil)
	LogTrace("TESTEVT", "TestMethod", ts, "", nil)

	tsn := testStructNillableChild{
		value: "Test1",
	}

	csn := &childStruct{
		value1: "Child1",
		value2: 200,
	}

	tsn.childStruct = csn

	LogError("TESTEVT", "TestMethod", tsn, "", nil)
	LogWarning("TESTEVT", "TestMethod", tsn, "", nil)
	LogData("TESTEVT", "TestMethod", tsn, "", nil)
	LogTrace("TESTEVT", "TestMethod", tsn, "", nil)

	e := errors.New("Some Test Error")

	LogError("TESTERR", "TestMethod", e, "", nil)
	LogError("TESTERR", "TestMethod", e.Error(), "", nil)

	LogWarning("TESTWARNING", "TestMethod", "This is testing", "", nil)

	tsChlds := testStructWithChilds{
		value: "Test2",
	}

	cs1 := childStruct{
		value1: "Child1",
		value2: 200,
	}

	cs2 := childStruct{
		value1: "Child2",
		value2: 600,
	}

	chlds := []childStruct{}

	chlds = append(chlds, cs1)
	chlds = append(chlds, cs2)

	tsChlds.childStruct = &chlds

	LogError("TESTEVT", "TestMethod", tsChlds, "", nil)
	LogWarning("TESTEVT", "TestMethod", tsChlds, "", nil)

	tsChlds.childStruct = nil

	LogError("TESTEVT", "TestMethod", tsChlds, "", nil)
	LogWarning("TESTEVT", "TestMethod", tsChlds, "", nil)
}

type testStruct struct {
	value       string
	childStruct childStruct
}

type testStructNillableChild struct {
	value       string
	childStruct *childStruct
}

type testStructWithChilds struct {
	value       string
	childStruct *[]childStruct
}

type childStruct struct {
	value1 string
	value2 int
}
