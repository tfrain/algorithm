// canCompleteCircuit determines if you can travel around a gas station circuit starting from one station.
func canCompleteCircuit(gas []int, cost []int) int {
    // Calculate the total gas and cost difference to check if the circuit can be completed.
    n := len(gas)
    sum := 0
    for i := range gas {
        sum += gas[i] - cost[i]
    }
    // If total gas is less than total cost, the circuit cannot be completed.
    if sum < 0 {
        return -1
    }

    // Initialize variables to track the gas in the tank and the starting station.
    tank := 0
    start := 0
    for i := 0; i < n; i++ {
        // Update the tank gas after visiting each station.
        tank += gas[i] - cost[i]
        // If the tank is negative, you can't reach station i from the start.
        if tank < 0 {
            // Set the next station as the new starting point and reset the tank.
            tank = 0
            start = i + 1
        }
    }
    // If start equals the number of stations, it means starting from the first station is possible.
    if start == n {
        return 0
    }
    // Return the index of the starting station.
    return start
}
