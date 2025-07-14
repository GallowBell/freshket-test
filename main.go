package main

func main() {

	isMember := false
	ordser_list := InitOrderList()
	ordser_list.newMenu("Red set", 2)
	ordser_list.newMenu("Green set", 5)

	ordser_list.print()
	total := ordser_list.Sum(isMember)
	println("Total order amount:", total)
}
