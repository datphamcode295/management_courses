package consts

type SalaryRatio struct {
	Ratio       int
	AfterBlock  int
	BeforeBlock int
}

var SalaryRatios []SalaryRatio = []SalaryRatio{
	{Ratio: 100, AfterBlock: 0, BeforeBlock: 15 * 28800},
	{Ratio: 80, AfterBlock: 15 * 28800, BeforeBlock: 30 * 28800},
	{Ratio: 5, AfterBlock: 30 * 28800, BeforeBlock: 60 * 28800},
	{Ratio: 0, AfterBlock: 60 * 28800, BeforeBlock: 1095 * 28800},
}
