package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTime int) int{
	if (prepTime == 0){
		return len(layers) * 2
	} else {
		return len(layers) * prepTime
	}
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64){
	sauceLayers := 0
	noodleLayers := 0
	for i := 0; i < len(layers); i++ {
		switch layers[i] {
		case  "noodles":
			noodleLayers += 1
		case  "sauce":
			sauceLayers += 1
		}
	}
	totalSauce := float64(sauceLayers) * 0.2
	totalNoodles := noodleLayers * 50
	return totalNoodles, totalSauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string){
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64{
	var newQuantities []float64
	for i := 0; i < len(quantities); i++ {
		newQuantities = append(newQuantities, quantities[i] * (float64(portions) / 2))
	}
	return newQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
