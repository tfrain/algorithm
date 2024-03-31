// Function to convert Roman numerals to integer using a different approach
func romanToInt(s string) int {
    // Function to get the integer value of a Roman numeral
    getInt := func(v byte) int {
        switch v {
        case 'I':
            return 1
        case 'V':
            return 5
        case 'X':
            return 10
        case 'L':
            return 50
        case 'C':
            return 100
        case 'D':
            return 500
        case 'M':
            return 1000
        default:
            return -1
        }
    }

    ret := 0
    var prev int

    // Iterate over the string
    for i := range s {
        // Get the integer value of the current numeral
        now := getInt(s[i])
        // Add the value to the result
        ret += now
        // If the current value is greater than the previous value, subtract twice the previous value from the result
        if now > prev {
            ret -= 2 * prev
        }
        // Update the previous value
        prev = now
    }

    // Return the result
    return ret
}