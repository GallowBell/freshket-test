package main

import "testing"

func TestInitOrderList(t *testing.T) {
	ordser_list := InitOrderList()
	if len(ordser_list) != 0 {
		t.Errorf("Expected order list to be initialized as empty, got length: %d", len(ordser_list))
	}
}

func TestAddNewMenu(t *testing.T) {
	ordser_list := InitOrderList()
	ordser_list.newMenu("Red set", 3)
	if len(ordser_list) != 1 {
		t.Errorf("Expected order list to have 1 item, got: %d", len(ordser_list))
	}
	if ordser_list[0].menu.name != "Red set" || ordser_list[0].quantity != 3 {
		t.Errorf("Expected menu item 'Red set' with quantity 3, got: %s with quantity %d", ordser_list[0].menu.name, ordser_list[0].quantity)
	}
}

func TestAddMultipleNewMenu(t *testing.T) {
	ordser_list := InitOrderList()
	ordser_list.newMenu("Red set", 3)
	ordser_list.newMenu("Purple set", 5)
	if len(ordser_list) != 2 {
		t.Errorf("Expected order list to have 2 item, got: %d", len(ordser_list))
	}
	if ordser_list[0].menu.name != "Red set" || ordser_list[0].quantity != 3 {
		t.Errorf("Expected menu item 'Red set' with quantity 3, got: %s with quantity %d", ordser_list[0].menu.name, ordser_list[0].quantity)
	}
	if ordser_list[1].menu.name != "Purple set" || ordser_list[1].quantity != 5 {
		t.Errorf("Expected menu item 'Purple set' with quantity 5, got: %s with quantity %d", ordser_list[0].menu.name, ordser_list[0].quantity)
	}
}

func TestCalculateDiscount(t *testing.T) {
	total := discountPrice(100, 5)
	expectedTotal := 100 * 5 / 100
	if total != int(expectedTotal) {
		t.Errorf("Expected total with discount to be %d, got: %d", expectedTotal, total)
	}
}

func TestSumOrderList(t *testing.T) {
	ordser_list := InitOrderList()
	ordser_list.newMenu("Red set", 3)
	ordser_list.newMenu("Yellow set", 2)
	total := ordser_list.Sum(false)
	expectedTotal := (50 * 3) + (50 * 2)
	if total != expectedTotal {
		t.Errorf("Expected total to be %d, got: %d", expectedTotal, total)
	}
}

func TestSumOrderListWithDiscountPair(t *testing.T) {
	ordser_list := InitOrderList()
	ordser_list.newMenu("Orange set", 2)
	ordser_list.newMenu("Pink set", 3)
	ordser_list.newMenu("Green set", 3)
	total := ordser_list.Sum(false)
	orangeDiscount := discountPrice(120, 5) * 2
	pinkDiscount := discountPrice(80, 5) * 2
	greenDiscount := discountPrice(40, 5) * 2
	expectedTotal := (120*2 + 80*3 + 40*3) - (orangeDiscount + pinkDiscount + greenDiscount)
	if total != expectedTotal {
		t.Errorf("Expected total with discount to be %d, got: %d", expectedTotal, total)
	}
}

func TestSumOrderListWithMemberDiscount(t *testing.T) {
	ordser_list := InitOrderList()
	ordser_list.newMenu("Red set", 3)
	ordser_list.newMenu("Yellow set", 2)
	total := ordser_list.Sum(true)
	expectedTotal := ((50 * 3) + (50 * 2)) - discountPrice((50*3)+(50*2), 10)
	if total != expectedTotal {
		t.Errorf("Expected total with member discount to be %d, got: %d", expectedTotal, total)
	}
}

func TestFreshketOne(t *testing.T) {
	ordser_list := InitOrderList()
	ordser_list.newMenu("Red set", 1)
	ordser_list.newMenu("Green set", 1)
	total := ordser_list.Sum(false)
	expectedTotal := 50 + 40
	if total != expectedTotal {
		t.Errorf("Expected total for Freshket One to be %d, got: %d", expectedTotal, total)
	}
}

func TestFreshketTwo(t *testing.T) {
	ordser_list := InitOrderList()
	ordser_list.newMenu("Red set", 1)
	ordser_list.newMenu("Green set", 1)
	total := ordser_list.Sum(true)
	expectedTotal := (50 * 1) + (40 * 1) - discountPrice((50*1)+(40*1), 10)
	if total != expectedTotal {
		t.Errorf("Expected total for Freshket Two to be %d, got: %d", expectedTotal, total)
	}
}

func TestFreshketThree(t *testing.T) {
	ordser_list := InitOrderList()
	ordser_list.newMenu("Orange set", 5)
	total := ordser_list.Sum(false)
	expectedTotal := (120 * 5) - (discountPrice(120, 5) * 4)
	if total != expectedTotal {
		t.Errorf("Expected total for Freshket Three to be %d, got: %d", expectedTotal, total)
	}
}
