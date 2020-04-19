package linkedlist

func checkIfPalindrome(l *LinkedList) bool {
	if l.length <= 1 {
		return true
	}

	prevHalf := make([]interface{}, 0, l.length/2)
	cur := l.head
	for i := 1; i <= l.length; i++ {
		cur = cur.next
		if l.length%2 != 0 && i == l.length/2+1 {
			continue
		}
		if i <= l.length/2 {
			prevHalf = append(prevHalf, cur.GetValue())
		} else {
			if prevHalf[l.length-i] != cur.GetValue() {
				return false
			}
		}
	}
	return true
}
