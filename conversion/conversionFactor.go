package conversion


type CurrencyType struct {
    Name          string
    ConversionFactor float64
}

const (
    INR float64 = 1.0
    USD float64 = 83.10
    EUR float64 = 89.04
    INV float64 = 0.0
)

var currencyTypes = map[string]CurrencyType{
    "INR": {Name: "INR", ConversionFactor: INR},
    "USD": {Name: "USD", ConversionFactor: USD},
    "EUR": {Name: "EUR", ConversionFactor: EUR},
    "INV": {Name: "INV", ConversionFactor: INV},
}

func GetCurrencyType(name string) (CurrencyType,bool) {
    currencyType, ok := currencyTypes[name]
    return currencyType, ok
}
