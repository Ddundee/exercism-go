package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}

	return len(layers) * time
}

// TODO: define the 'Quantities()' function
func Quantities(slice []string) (noodles int, sauce float64) {

	for i := 0; i < len(slice); i++ {
		if slice[i] == "noodles" {
			noodles += 50
		}
		if slice[i] == "sauce" {
			sauce += .2
		}
	}

	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList) - 1] = friendsList[len(friendsList) - 1]

	// return myList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quan []float64, num int) []float64 {
	quantities := make([]float64, len(quan))

	for i := 0; i < len(quan); i++ {
		quantities[i] = float64(num) * quan[i] / 2
	}

	return quantities
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
