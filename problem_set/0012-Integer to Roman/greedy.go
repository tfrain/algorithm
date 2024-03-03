// Function to convert an integer to a Roman numeral
func intToRoman(num int) string {
    // Define arrays of Roman symbols and their corresponding integer values
    var (
        symbols = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
        values  = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    )

    // Initialize a string builder, which is more efficient for string concatenation
    var b strings.Builder
    
    // Loop over the arrays of Roman symbols and values
    // Start from the largest value (e.g., 1000 corresponds to "M")
    for i := 0; num > 0; i++ {
        // Loop until the given number is less than the current symbol's integer value
        for num >= values[i] {
            // Subtract the current symbol's integer value from the given number
            num -= values[i]
            // Write the current symbol to the string builder
            b.WriteString(symbols[i])
        }
    }
    // Output the final Roman numeral representation by converting the string builder to a string
    return b.String()
}