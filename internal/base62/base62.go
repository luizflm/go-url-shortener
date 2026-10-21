package base62

const base62chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func ConvertToBase62(decimal int64) string {
	if decimal == 0 {
		return "0"
	}

	b62 := ""

	numerator := decimal
	for numerator > 0 {
		remainder := numerator % 62
		b62 = string(base62chars[remainder]) + b62
		numerator = numerator / 62
	}

	return b62
}
