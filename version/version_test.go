package version

import (
	"testing"
)

func TestParse(t *testing.T) {
	testData := []struct {
		input  string
		expect PMSVersion
	}{
		{"0.0.0.0", PMSVersion{[]uint{0, 0, 0, 0}}},
		{"100.032.10.0", PMSVersion{[]uint{100, 32, 10, 0}}},
	}

	for i, data := range testData {
		// TODO: test ok
		// TODO: rework when adding segments
		result, _ := parse(data.input)
		same := true
		for index, _ := range result.nums {
			if result.nums[index] != data.expect.nums[index] {
				same = false
			}
		}
		if same {
			t.Logf("[%d/%d] Passed: %v was equal to %v", i+1, len(testData), result, data.expect)
		} else {

			t.Errorf("[%d/%d] Falied: %v wasn't equal to %v", i+1, len(testData), result, data.expect)
		}
	}

}
func TestCompare(t *testing.T) {
	testData := []struct {
		inputA PMSVersion
		inputB PMSVersion
		expect int
	}{
		{PMSVersion{[]uint{1, 1, 1}}, PMSVersion{[]uint{1, 1, 1}}, 0},
		{PMSVersion{[]uint{1, 0, 1}}, PMSVersion{[]uint{1, 1, 1}}, -1},
		{PMSVersion{[]uint{1, 1, 1}}, PMSVersion{[]uint{1, 1}}, 1},
		{PMSVersion{[]uint{10, 1, 1}}, PMSVersion{[]uint{1, 1, 1}}, 1},
		{PMSVersion{[]uint{1, 1, 0}}, PMSVersion{[]uint{1, 10, 1}}, -1},
	}
	for i, data := range testData {
		result := compare(data.inputA, data.inputB)
		if result == data.expect {
			t.Logf("[%d/%d] Passed: %v was equal to %v", i+1, len(testData), result, data.expect)
		} else {
			t.Errorf("[%d/%d] Falied: %v wasn't equal to %v", i+1, len(testData), result, data.expect)
		}
	}

}

func TestParseInt(t *testing.T) {
	testData := []struct {
		input  string
		expect []uint
	}{
		{"0.0.0.0", []uint{0, 0, 0, 0}},
		{"100.032.10.0", []uint{100, 32, 10, 0}},
	}

	for i, data := range testData {
		result := parseNums(data.input)
		same := true
		for index, _ := range result {
			if result[index] != data.expect[index] {
				same = false
			}
		}
		if same {
			t.Logf("[%d/%d] Passed: %v was equal to %v", i+1, len(testData), result, data.expect)
		} else {

			t.Errorf("[%d/%d] Falied: %v wasn't equal to %v", i+1, len(testData), result, data.expect)
		}
	}
}

func TestCompareNums(t *testing.T) {
	testData := []struct {
		inputA []uint
		inputB []uint
		expect int
	}{
		{[]uint{1, 1, 1}, []uint{1, 1, 1}, 0},
		{[]uint{1, 0, 1}, []uint{1, 1, 1}, -1},
		{[]uint{1, 1, 1}, []uint{1, 1}, 1},
		{[]uint{10, 1, 1}, []uint{1, 1, 1}, 1},
		{[]uint{1, 1, 0}, []uint{1, 10, 1}, -1},
	}
	for i, data := range testData {
		result := compareNums(data.inputA, data.inputB)
		if result == data.expect {
			t.Logf("[%d/%d] Passed: %v was equal to %v", i+1, len(testData), result, data.expect)
		} else {
			t.Errorf("[%d/%d] Falied: %v wasn't equal to %v", i+1, len(testData), result, data.expect)
		}
	}
}
