package sub

import (
	"fmt"
)

type Result interface {
}
type Cloned struct {
	Result
	Name    string
	Message string
}
type Pulled struct {
	Result
	Name    string
	Message string
}
type LocalArchived struct {
	Result
	Name    string
	Message string
}
type Error struct {
	Result
	Name    string
	Message string
}

func Consolidate(input []Result) string {
	if len(input) == 0 {
		return "Nothing happened!"
	}
	var pulled = 0
	var cloned []Cloned
	var localArchived []LocalArchived
	var errors []Error
	for _, x := range input {
		switch x.(type) {
		case Pulled:
			pulled += 1
		case Cloned:
			cloned = append(cloned, x.(Cloned))
		case LocalArchived:
			localArchived = append(localArchived, x.(LocalArchived))
		case Error:
			errors = append(errors, x.(Error))
		default:
			fmt.Println("unknown")
		}
	}

	var resultStr = "\n"
	if len(cloned) != 0 {
		resultStr += fmt.Sprintf("Cloned %d new repositories:\n", len(cloned))
		for _, e := range cloned {
			resultStr += fmt.Sprintf("\t%s:\t%s\n", e.Name, e.Message)
		}
	}
	if len(localArchived) != 0 {
		resultStr += fmt.Sprintf("Local copies of archived repositories:\n")
		for _, e := range localArchived {
			resultStr += fmt.Sprintf("\t%s:\t%s\n", e.Name, e.Message)
		}
	}
	if len(errors) != 0 {
		resultStr += fmt.Sprintf("Errors happened in these repositories:\n")
		for _, e := range errors {
			resultStr += fmt.Sprintf("\t%s:\t%s\n", e.Name, e.Message)
		}
	}
	if pulled != 0 {
		resultStr += fmt.Sprintf("Pulled %d repositories.", pulled)
	}

	return resultStr
}
