package test

import (
	"fmt"
	"testing"

	"github.com/dailymotion/allure-go"
)

func TestStep(t *testing.T) {
	allure.Test(t, allure.Action(func() {
		fmt.Println("This block of code is a test")
	}))
}

func TestParameter(t *testing.T) {
	allure.Test(t,
		allure.Description("Test that has parameters"),
		allure.Parameter("testParameter", "test"),
		allure.Action(func() {
			allure.Step(allure.Description("Step with parameters"),
				allure.Action(func() {}),
				allure.Parameter("stepParameter", "step"))
		}))
}
