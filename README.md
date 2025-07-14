## 📋 Table of Contents
- [Overview](#overview)
- [Requirements](#requirements)
- [Menu Items](#menu-items)
- [Business Rules](#business-rules)
- [Examples](#examples)
- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)


## 🎯 Overview
**Homework test for software engineer (Backend)**

**Time:** 60 - 120 minutes recommended.

**Submission:** recruit-tech@freshket.co

**Scoring Criteria:**
- ✅ Readability
- ✅ Maintainability  
- ✅ Extensibility
- ✅ Logic

## 📝 Requirements

### 1) Calculator
Write a Calculator class for food store (you can use any programming languages)

This food store only have 7 items in menu

## 🍽️ Menu Items

| Item | Price (THB) |
|------|-------------|
| Red set | 50 |
| Green set | 40 |
| Blue set | 30 |
| Yellow set | 50 |
| Pink set | 80 |
| Purple set | 90 |
| Orange set | 120 |

## 📐 Business Rules

### Core Features
- ✅ Customers can order multiple items
- ✅ Calculate and return the total price

### Discount Rules
1. **Member Discount:** 10% off total if customer has member card
2. **Pair Discount:** 5% off for every 2 items of Orange, Pink, or Green sets

### 🎯 Examples

#### Example 1: Basic Order
```
Order: Red set + Green set = 90 THB
With member card: 90 - (90 × 10%) = 81 THB
```

#### Example 2: Pair Discount
```
Order: 5x Orange sets
- 4 items get pair discount (2 pairs)
- 1 item at regular price
```

## 🚀 Installation

```bash
# Clone the repository
git clone https://github.com/GallowBell/freshket-test
cd freshket-test

# Run the application
go run .
```

## 💻 Usage

```go 
// main.go
// Create new order
order := InitOrderList()

// Add items to order using newMenu
// parameter Name of set, Quantity
order.newMenu("Red set", 1)
order.newMenu("Green set", 1)

// Calculate total (with member status)
total := order.Sum(true) // true = member, false = non-member
```

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test -run TestOrderCalculation
```

## 📊 Test Cases

- Basic price calculation
- Member discount application
- Pair discount for Orange/Pink/Green sets
- Other