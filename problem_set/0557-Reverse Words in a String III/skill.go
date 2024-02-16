// Function reverseWords reverses each word in a given string s
func reverseWords(s string) string {
    // Convert the string to a byte array to allow for in-place reversal.
    b := []byte(s)

    // Initialize pointers i and j to 0.
    for i, j := 0, 0; i < len(b); i = j + 1 {
        // Find the beginning of each word.
        for i < len(b) && b[i] == ' ' {
            i++
        }

        // Find the end of each word.
        j = i + 1
        for j < len(b) && b[j] != ' ' {
            j++
        }

        // Reverse the current word in-place.
        for k := j - 1; i < k; i, k = i+1, k-1 {
            b[i], b[k] = b[k], b[i]
        }
    }

    // Convert the byte array back to a string and return.
    return string(b)
}