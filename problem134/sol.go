package problem134

/*
	LeetCode Problem 134: Gas Station
	Level: Medium
	Description: There are n gas stations along a circular route, where the amount of gas at the ith station is gas[i].
	You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from the ith station to its next (i + 1)th station. You begin the journey with an empty tank at one of the gas stations.
	Given two integer arrays gas and cost, return the starting gas station's index if you can travel around the circuit once in the clockwise direction, otherwise return -1. If there exists a solution, it is guaranteed to be unique
	Link: https://leetcode.com/problems/gas-station/
*/

// CanCompleteCircuit return the starting gas station's index if it is able to travel around the circuit
// once in the clockwise direction, otherwise return -1
func CanCompleteCircuit(gas []int, cost []int) int {
	gases, costs := SumInt(gas), SumInt(cost)
	if gases < costs {
		return -1
	}
	ans := 0
	net := 0
	for i := 0; i < len(gas); i++ {
		net += gas[i] - cost[i]
		if net < 0 {
			ans = i + 1
			net = 0
		}
	}
	return ans
}

// SumInt return total sum of ints
func SumInt(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// MaxInt return the larger one
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
