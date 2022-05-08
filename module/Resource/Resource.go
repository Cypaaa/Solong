package Resource

func OneLineIfElseStr(condition bool, vTrue string, vFalse string) string {
	if condition {
		return vTrue
	}
	return vFalse
}
