func groupAnagrams(strs []string) [][]string {
    // Initialize a map to store the anagrams
    strsMap := make(map[string][]string)

    // Iterate over the input strings
    for _, s := range strs {
        // Encode the string to a unique identifier
        code := encode(s)
        // Add the string to the corresponding group in the map
        strsMap[code] = append(strsMap[code], s)
    }

    // Initialize a slice to store the result
    var ret [][]string

    // Iterate over the groups in the map
    for _, list := range strsMap {
        // Add each group to the result
        ret = append(ret, list)
    }

    // Return the result
    return ret
}

func encode(s string) string {
    // Initialize a byte slice to store the count of each character
    count := make([]byte, 26)

    // Iterate over the characters in the string
    for i := 0; i < len(s); i++ {
        // Calculate the index in the count slice for the current character
        delta := s[i] - 'a'
        // Increment the count for the current character
        count[delta]++
    }

    // Convert the count slice to a string and return it
    return string(count)
}