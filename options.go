package dbquery

import "strings"

// SearchOption - a tool for constructing a query from parameters and conditions
type SearchOption struct {
	SQL    string
	Values []interface{}
}

// build sql string & values array
func (o *SearchOption) build() (string, []interface{}) {
	// TODO
	return "", nil
}

func optionVal(key, option string, value interface{}) SearchOption {
	return SearchOption{
		SQL:    "(" + key + option + "{R})",
		Values: []interface{}{value},
	}
}

// Eq - values are equal
func Eq(key string, value interface{}) SearchOption {
	return optionVal(key, "=", value)
}

// EqMore - value are equal or more than
func EqMore(key string, value interface{}) SearchOption {
	return optionVal(key, ">=", value)
}

// EqLess - value are equal or less than
func EqLess(key string, value interface{}) SearchOption {
	return optionVal(key, "<=", value)
}

// Less - than value
func Less(key string, value interface{}) SearchOption {
	return optionVal(key, "<", value)
}

// More - than value
func More(key string, value interface{}) SearchOption {
	return optionVal(key, ">", value)
}

func joinOptions(glue string, options ...SearchOption) SearchOption {
	resultValues := []interface{}{}
	optionsSQL := []string{}

	for _, option := range options {
		optionsSQL = append(optionsSQL, option.SQL)
		resultValues = append(resultValues, option.Values...)
	}

	resultSQL := "( " + strings.Join(optionsSQL, glue) + " )"
	return SearchOption{
		SQL:    resultSQL,
		Values: resultValues,
	}
}

// And - multiple values should be observed together
func And(options ...SearchOption) SearchOption {
	return joinOptions("AND")
}

// Or - one of the values must be true
func Or(options ...SearchOption) SearchOption {
	return joinOptions("OR")
}
