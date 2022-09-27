package main


func twoSums(nums []int, target int) []int {

	// Hash map
	// Create a hash map
	
	hm := make(map[int]int)

	// Iterate the numbers
	// Add compliments to the hash map

	for i, num := range(nums){
		_, ok := hm[num]
		if ok {
			return []int{i, hm[num]}
			
		}

		// compliment to the target

		hm[target - num] = i
	}

	return nil

	// O(n)

}