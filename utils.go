package main

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func encodeBase62(num int64) string {
	if num == 0 {
		return "0"
	}

	var result []byte
	for num > 0 {
		remainder := num % 62
		result = append([]byte{base62Chars[remainder]}, result...)
		num /= 62
	}
	return string(result)
}