package codeTest

func MonkeyTrouble(MonkeyA, MonkeyB bool) bool {
	if MonkeyA && MonkeyB {
		return true
	}
	if !MonkeyA && !MonkeyB {
		return true
	}
	return false
}
