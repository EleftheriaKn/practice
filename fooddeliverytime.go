package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	foods := food{}
	if order == "burger" {
		foods.preptime = 15
	} else if order == "chips" {
		foods.preptime = 10
	} else if order == "nuggets" {
		foods.preptime = 12
	} else {
		return 404
	}
	return foods.preptime
}
