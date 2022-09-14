package Calc

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestCalc(t *testing.T) {
	testCases := []struct {
		name   string
		symbol string
		a      float64
		b      float64
		output float64
	}{
		{
			name:   "+ positive",
			symbol: "+",
			a:      3,
			b:      4,
			output: 7,
		},
		{
			name:   "+ negative",
			symbol: "+",
			a:      -5.3,
			b:      3,
			output: -2.3,
		},
		{
			name:   "- positive",
			symbol: "-",
			a:      3,
			b:      4,
			output: -1,
		},
		{
			name:   "- negative",
			symbol: "-",
			a:      -5.3,
			b:      3,
			output: -8.3,
		},
		{
			name:   "* positive",
			symbol: "*",
			a:      3,
			b:      4,
			output: 12,
		},
		{
			name:   "- negative",
			symbol: "*",
			a:      -5.3,
			b:      3,
			output: -15.9,
		},
		{
			name:   "/ positive",
			symbol: "/",
			a:      3,
			b:      4,
			output: 0.75,
		},
		{
			name:   "/ negative",
			symbol: "/",
			a:      -5.3,
			b:      3,
			output: -1.77,
		},
		{
			name:   "n! positive",
			symbol: "n!",
			a:      3,
			b:      4,
			output: 6,
		},
		{
			name:   "n! negative",
			symbol: "n!",
			a:      -5.3,
			b:      3,
			output: -0.1,
		},
		{
			name:   "^ positive",
			symbol: "^",
			a:      3,
			b:      4,
			output: 81,
		},
		{
			name:   "^ negative",
			symbol: "^",
			a:      -5.3,
			b:      3,
			output: -148.88,
		},
		{
			name:   "V positive",
			symbol: "V",
			a:      3,
			b:      4,
			output: 1.32,
		},
		{
			name:   "V negative b",
			symbol: "V",
			a:      5.3,
			b:      -3,
			output: 0.57,
		},
		{
			name:   "V negative a",
			symbol: "V",
			a:      -5.3,
			b:      3,
			output: -1.74,
		},
		{
			name:   "Wrong input",
			symbol: "T",
			a:      5.3,
			b:      -3,
			output: 0,
		},
	}

	for _, cse := range testCases {
		cse := cse
		
		// Testify
		assert.Equal(t, cse.output, (math.Round(Run(cse.a, cse.b, cse.symbol)*100) / 100), "error in: "+cse.name)

		//Go Tests
		t.Run(cse.name, func(t *testing.T) {
			result := Run(cse.a, cse.b, cse.symbol)

			// Большое спасибо - данный тест явно показывает, что само домашнее задание было выполнено с плохим качеством кода - я не могу тест написать полноценно на всю Функцию. Мне приходится закомментировать часть кода (12ю строку - // a, b, symbol = inputAndValidation()), что бы тест выполнился. Необходимо инпут вынести в отдельную функцию и таким образом все что вне него вызывать.
			if (math.Round(result*100) / 100) != cse.output {
				t.Errorf("function returns wrong answer: expected %f, got: %f", cse.output, result)
			}
		})
	}
}
