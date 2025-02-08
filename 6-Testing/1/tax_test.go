/*
INFO: 6.1 - Iniciando com testes automatizados | 6.2-Testando em batch | 6.3-Cobertura de código
INFO: 6.4-Trabalhando com Benchmarking 6.5-Fuzzing
*/
package tax

/*
go test -coverprofile=coverage.out -> show code coverage on terminal
go tool cover -html=coverage.out -> generate html file with code coverage
*/
import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expected := 5.0
	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expected float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
	}

	for _, entry := range table {
		result := CalculateTax(entry.amount)
		if result != entry.expected {
			t.Errorf("Expected %f but got %f", entry.expected, result)
		}
	}
}

/*
go test -bench=.
go test -bench=. -run=^# -> run only benchmark functions
go test -bench=. -run=^# -benchmem -> show memory allocation on benchmak
*/
func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}
func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}

/*
go test -fuzz=. -run=^# -> run only fuzz functions
go test -fuzz=. -run=^# -> run only fuzz functions
go test -fuzz=. -fuzztime=5s -run=^# -> set for how long it will test
*/

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Received %f but expected 0", result)
		}
        

	})
}
