package number

type MyInt int

func (m MyInt) GetPrimeFactor() []int {
  factors := []int{}
  number := int(m)
  for factor := 2; factor <= number; factor++ {
    for (number % factor == 0) {
      factors = append(factors, factor)
      number /= factor
    }
  }
  return factors
}

