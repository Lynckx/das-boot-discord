package rules

type Rules struct {
	Name  string
	rules map[Rule]interface{}
}
type Rule struct {
	Name string
}
