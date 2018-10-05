package sizeof

import (
	"testing"
)

func TestQuantileProcessing(t *testing.T) {
	type testCase struct {
		obj  interface{}
		size uint64
	}

	type otra struct { // just a struct to use in tests
		una    int
		dos    int8
		tres   int8
		cuatro int8
		cinco  int8
		seis   int8
		siete  string
	}

	type cosa struct { // just a struct to use in tests
		una    int
		dos    int8
		tres   int8
		cuatro int8
		cinco  int8
		seis   int8
		siete  string
		sub    otra
		SubPtr *otra
	}

	var testCases = []testCase{
		testCase{
			obj:  int8(120),
			size: 1,
		},
		testCase{
			obj:  uint16(5000),
			size: 2,
		},
		testCase{
			obj:  uint32(65535),
			size: 4,
		},
		testCase{
			obj:  float64(3.14),
			size: 8,
		},
		testCase{
			obj:  []int{1, 2, 3},
			size: 48,
		},
		testCase{
			obj:  [3]int{1, 2, 3},
			size: 24,
		},
		testCase{
			obj:  &[3]int{1, 2, 3},
			size: 32,
		},
		testCase{
			obj: struct {
				a int64
				b int8
				c int32
				d int8
				e int8
			}{
				1, 2, 3, 4, 5,
			},
			size: 24,
		},
		testCase{
			obj:  "",
			size: 16,
		},
		testCase{
			obj:  "Hello World!",
			size: 28,
		},
		testCase{
			obj: []*cosa{
				&cosa{},
				&cosa{},
			},
			size: 184,
		},
		testCase{
			obj: map[cosa]*cosa{
				cosa{una: 1}: &cosa{},
				cosa{una: 2}: nil,
			},
			size: 428,
		},
		testCase{
			obj: map[*cosa]*cosa{
				&cosa{una: 1}: &cosa{},
				&cosa{una: 2}: nil,
			},
			size: 444,
		},
		testCase{
			obj: map[[2]int]int{
				[2]int{1, 2}: 3,
			},
			size: 196,
		},
		testCase{
			obj: map[otra]*cosa{
				otra{}: &cosa{
					siete: "  ",
					SubPtr: &otra{
						siete: "   ",
					},
				},
				otra{siete: "7  "}: nil,
			},
			size: 388,
		},
	}

	for i, tc := range testCases {
		size := SizeOf(tc.obj)
		if size != tc.size {
			t.Errorf("Failed test %d, expected %d but got %d", i, tc.size, size)
		}
	}
}
