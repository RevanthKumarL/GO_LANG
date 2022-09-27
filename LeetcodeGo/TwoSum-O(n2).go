package main

func twoSums(nums []int, target int) []int {

	// iterating over both slices
	// checking for the target

	//O(n^2)

	for i, left := range(nums){
		for j, right := range(nums){
			if left + right == target && i!=j {
				return []int{i,j}
			}
		}
	}



	// just return nothing since 
	// there will always be an answer

	return nil


}