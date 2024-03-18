// findRadius is the primary function that calculates the minimum radius the heaters must have to heat all the houses in a neighborhood.
func findRadius(houses []int, heaters []int) int {
    // Sorting the heater positions
    sort.Ints(heaters)
    var radius int
    // Iterating over each house
    for _, house := range houses {
        // Finding the index of heater that is closest to and less than or equal to the current house's position
        i := findIndex(heaters, house)
        min := math.MaxInt32
        // If some heater was found that is less than or equal to house
        if i >= 0 {
            // The potential radius could be the distance between the house and this heater
            min = house - heaters[i]
        }
        // Checking the heater to the right of the current house's position, if it exists and is closer than the former found heater
        if i+1 < len(heaters) && heaters[i+1]-house < min {
            // The potential radius could be the distance between the house and this heater
            min = heaters[i+1] - house
        }
        // Updating the radius if a larger potential radius is found
        if min > radius {
            radius = min
        }
    }
    // Return the found minimum necessary radius of heaters
    return radius
}

// findIndex is a helper function to search the heater array using binary search for the given house's position, and find out the closest heater that is less than or equal to the house's position.
func findIndex(heaters []int, num int) int {
    // Initiating pointers for the boundaries of the search space within the heaters' positions array
    lo, hi := -1, len(heaters)-1
    // Until the lower boundary is less than the upper boundary
    for lo < hi {
        // Calculating the middle point within the search space
        mid := hi - (hi-lo)/2
        // Adjusting the search space depending on the heater's position at the mid-point
        if heaters[mid] < num { 
            lo = mid
        } else if heaters[mid] > num {
            hi = mid - 1
        } else {
            // If the heater's position equals the house's position, return the index immediately
            return mid
        }
    }
    // Return the index of the closest heater that is less than the house's position
    return lo
}