package thefarm

import (
    "errors"
    "fmt"
)

// See types.go for the types defined for this exercise.

type SillyNephewError struct {
    cows int
}

func (e *SillyNephewError) Error() string {
    return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
    amount, err := weightFodder.FodderAmount()
    if err == ErrScaleMalfunction && amount > 0 {
        return (amount * 2) / float64(cows), nil
    }
	if amount < 0 && (err == ErrScaleMalfunction || err == nil) {
        return 0, errors.New("negative fodder")
    }
	if cows == 0 {
        return 0, errors.New("division by zero")
    }
	if cows < 0 {
        return 0, &SillyNephewError{cows}
    }
    if err != nil {
        return 0, err
    }
    return amount / float64(cows), nil
}
