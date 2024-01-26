func isValid(s string) bool {
    brackets := map[byte]byte{'}': '{', ']': '[', ')': '('}
    stack := make([]byte, 0)
    for i := range s {
        switch s[i] {
        case '(', '{', '[':
            stack = append(stack, s[i])
        default:
            if len(stack) > 0 && stack[len(stack)-1] == brackets[s[i]] {
                stack = stack[:len(stack)-1]
            } else {
                return false
            }
        }
    }
    return len(stack) == 0
}

func isValid(s string) bool {
    brackets := map[byte]byte{')': '(', '}': '{', ']': '['}
    stack := make([]byte, len(s))
    var i int
    for j := 0; j < len(s); j++ {
        switch s[j] {
        case '(', '{', '[':
            stack[i] = s[j]     // push opening bracket onto stack
            i++
        default:
            if i == 0 || stack[i-1] != brackets[s[j]] {   // check top of stack for corresponding opening bracket
                return false
            }
            i--   // pop off opening bracket from stack
        }
    }
    return i == 0
}