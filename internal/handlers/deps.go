package handlers

//go:generate moq -out converter_mock.go . Converter

type Converter interface {
	Convert(num string, base int, toBase int) (string, error)
}
