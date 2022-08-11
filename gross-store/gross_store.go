package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
        "quarter_of_a_dozen": 3,
        "half_of_a_dozen": 6,
        "dozen": 12,
        "small_gross": 120,
        "gross": 144,
        "great_gross": 1728,
    }
}

// NewBill creates a new bill.
func NewBill() map[string]int {
    bill := make(map[string]int)
    return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	quantity, ok := units[unit]
    if ok {
        bill[item] += quantity
    }
    return ok	
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	billQuantity, billQuantityOk := bill[item]
	unitQuantity, unitQuantityOk := units[unit]
	newQuantity := billQuantity - unitQuantity
	if billQuantityOk && unitQuantityOk && newQuantity >= 0 {
        if newQuantity == 0 {
            delete(bill, item)
        } else {
            bill[item] = newQuantity
        }
    	return true
    } else {
    	return false
    }
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	billQuantity, ok := bill[item]
    return billQuantity, ok
}
