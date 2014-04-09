package number_test

import "testing"
import "reflect"
import . "github.com/roongr2k7/number"

func TestGetPrimeFactor(t *testing.T) {
  tests := []struct{
    input MyInt
    expected []int
  }{
    {2, []int{2}},
    {3, []int{3}},
    {4, []int{2, 2}},
    {5, []int{5}},
    {6, []int{2, 3}},
    {8, []int{2, 2, 2}},
    {9, []int{3, 3}},
    {16, []int{2, 2, 2, 2}},
  }
  for _, test := range tests {
    actual := test.input.GetPrimeFactor()
    if !reflect.DeepEqual(actual, test.expected) {
      t.Errorf("ERROR: prime factor of %v expect %v, actual %v", test.input, test.expected, actual)
    }
  }
}

func BenchmarkGetPrimeFactor(b *testing.B) {
  for n := 0; n <= b.N; n++ {
    MyInt(n).GetPrimeFactor()
  }
}

