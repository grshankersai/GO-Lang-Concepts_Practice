package generic

func Add[T int|string|float64](a,b T) T{
	return a+b
}