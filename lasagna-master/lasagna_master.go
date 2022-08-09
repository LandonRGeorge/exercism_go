package lasagna

func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}
	return len(layers) * timePerLayer
}

func Quantities(layers []string) (int, float64) {
	var noodleLayers int
	var sauceLayers int
	noodlesPerLayer := 50
	saucePerLayer := 0.2
	for _, layer := range layers {
		if layer == "noodles" {
			noodleLayers += 1
		} else if layer == "sauce" {
			sauceLayers += 1
		}
	}
	return noodleLayers * noodlesPerLayer, float64(sauceLayers) * saucePerLayer

}

func AddSecretIngredient(friendsList []string, myList []string) {
	friendsListMaxIndex := len(friendsList) - 1
	myListMaxIndex := len(myList) - 1
	secretIngredient := friendsList[friendsListMaxIndex]
	myList[myListMaxIndex] = secretIngredient
}

func ScaleRecipe(amounts []float64, portions int) []float64 {
	amountsScaled := make([]float64, len(amounts))
	for index, amount := range amounts {
		amountsScaled[index] = amount / 2 * float64(portions)
	}
	return amountsScaled
}
