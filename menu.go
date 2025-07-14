package main

import "fmt"

type MainMenu struct {
	name  string
	price int
}

var MainMenuList = []MainMenu{
	{"Red set", 50},
	{"Green set", 40},
	{"Blue set", 30},
	{"Yellow set", 50},
	{"Pink set", 80},
	{"Purple set", 90},
	{"Orange set", 120},
}

type MenuList struct {
	menu     MainMenu
	quantity int
}

type OrderList []MenuList

func InitOrderList() OrderList {
	return OrderList{}
}

func (ol *OrderList) newMenu(menuName string, quantity int) {
	for i, item := range MainMenuList {
		if item.name == menuName {
			Menu := MainMenuList[i]
			newItem := MenuList{menu: Menu, quantity: quantity}
			*ol = append(*ol, newItem)
			return
		}
	}
}

func (data OrderList) print() {
	for i, ol := range data {
		fmt.Println(i, ol.menu.name, ol.menu.price, ol.quantity)
	}
}

/*
	Calculate the total order amount with discounts applied
	** status: true for member, false for non-member
*/
func (ol *OrderList) Sum(status bool) int {
	var total int
	discountItems := []string{"Orange set", "Pink set", "Green set"}
	for _, od := range *ol {
		price := od.menu.price
		println("price", price)
		discountPromotion := 0
		isDiscountItem := false
		for _, item := range discountItems {
			if od.menu.name == item {
				isDiscountItem = true
				break
			}
		}
		if isDiscountItem && od.quantity >= 2 {
			discountPromotion = discountPrice(price, 5)
			println("discountPromotion", discountPromotion)
			isOdd := od.quantity % 2
			println("isOdd", isOdd)
			price = price - discountPromotion
			if isOdd == 0 {
				discountPromotion = 0
			}
			println("price after discount", price)
		}
		total += (price * od.quantity) + discountPromotion
	}
	if status {
		total = total - discountPrice(total, 10)
		println("Member discount applied: 10%")
	}
	return total
}

/* Calculate the discount price based on the total and discount percentage */
func discountPrice(total int, discount float64) int {
	return int(float64(total) * (discount / 100))
}
