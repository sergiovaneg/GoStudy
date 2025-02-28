package fractals

type Solver interface {
	String() string
	Solve(string, int, []string) uint
}
