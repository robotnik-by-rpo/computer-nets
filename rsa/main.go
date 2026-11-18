package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

// RSA — structure
type RSA struct {
	P int64
	Q int64
	E int64
	N int64
	D int64
}

func NewRSA(p, q int64) (*RSA, error) {
	r := &RSA{P: p, Q: q}
	if err := r.GenerateKeys(); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *RSA) isPrime(n int64) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := int64(3); i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func (r *RSA) gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func (r *RSA) modInverse(a, m int64) int64 {
	m0 := m
	y := int64(0)
	x := int64(1)
	if m == 1 {
		return 0
	}
	for a > 1 {
		q := a / m
		t := m
		m = a % m
		a = t
		t = y
		y = x - q*y
		x = t
	}
	if x < 0 {
		x += m0
	}
	return x
}

func (r *RSA) modPow(base, exp, modulus int64) int64 {
	result := int64(1)
	base = base % modulus
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % modulus
		}
		exp = exp >> 1
		base = (base * base) % modulus
	}
	return result
}

func (r *RSA) GenerateKeys() error {
	if !r.isPrime(r.P) || !r.isPrime(r.Q) {
		return fmt.Errorf("P and Q must be prime numbers")
	}
	if r.P == r.Q {
		return fmt.Errorf("P and Q mustn't be equal")
	}

	r.N = r.P * r.Q
	phi := (r.P - 1) * (r.Q - 1)

	r.E = 65537
	if r.E >= phi {
		r.E = 3
		for r.E < phi {
			if r.gcd(r.E, phi) == 1 {
				break
			}
			r.E += 2
		}
	} else if r.gcd(r.E, phi) != 1 {
		r.E = 3
		for r.gcd(r.E, phi) != 1 {
			r.E += 2
		}
	}

	r.D = r.modInverse(r.E, phi)
	if r.D == 0 {
		return fmt.Errorf("could not be calculated d")
	}
	return nil
}

func (r *RSA) PublicKey() (int64, int64) {
	return r.E, r.N
}

func (r *RSA) PrivateKey() (int64, int64) {
	return r.D, r.N
}

func (r *RSA) EncryptNumber(m int64) (int64, error) {
	if m < 0 || m >= r.N {
		return 0, fmt.Errorf("m must be in a range [0, n-1]")
	}
	return r.modPow(m, r.E, r.N), nil
}

func (r *RSA) DecryptNumber(c int64) int64 {
	return r.modPow(c, r.D, r.N)
}

func (r *RSA) EncryptString(message string) (string, error) {
	var ciphertexts []string
	for _, ch := range message {
		m := int64(ch)
		if m >= r.N {
			return "", fmt.Errorf("character code '%c' (%d) is >= n (%d), increase P and Q", ch, m, r.N)
		}
		c, err := r.EncryptNumber(m)
		if err != nil {
			return "", err
		}
		encoded := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", c)))
		ciphertexts = append(ciphertexts, encoded)
	}
	return strings.Join(ciphertexts, "."), nil
}

func (r *RSA) DecryptString(encodedCipher string) (string, error) {
	parts := strings.Split(encodedCipher, ".")
	var result strings.Builder
	for _, part := range parts {
		decodedBytes, err := base64.StdEncoding.DecodeString(part)
		if err != nil {
			return "", fmt.Errorf("error decoding Base64: %v", err)
		}
		var c int64
		_, err = fmt.Sscanf(string(decodedBytes), "%d", &c)
		if err != nil {
			return "", fmt.Errorf("error parsing number: %v", err)
		}
		m := r.DecryptNumber(c)
		result.WriteRune(rune(m))
	}
	return result.String(), nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	readLine := func() string {
		line, _ := reader.ReadString('\n')
		return strings.TrimRight(line, "\r\n")
	}

	fmt.Println("=== Task 1 ===")

	var p, q int64
	fmt.Print("Enter prime number P: ")
	fmt.Scan(&p)
	fmt.Print("Enter prime number Q: ")
	fmt.Scan(&q)

	rsa, err := NewRSA(p, q)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	e, n := rsa.PublicKey()
	d, _ := rsa.PrivateKey()
	fmt.Printf("\nOpen key: (e=%d, n=%d)\n", e, n)
	fmt.Printf("Close key: (d=%d, n=%d)\n", d, n)

	var m int64
	fmt.Print("\nEnter number m (0 <= m < n): ")
	fmt.Scan(&m)

	c, err := rsa.EncryptNumber(m)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Encrypted number: %d\n", c)

	decrypted := rsa.DecryptNumber(c)
	fmt.Printf("Decrypted number: %d\n", decrypted)

	if decrypted == m {
		fmt.Println("Check passed")
	} else {
		fmt.Println("Error: The transcript does not match.")
	}

	fmt.Println("\n=== Task 2 ===")

	readLine()

	fmt.Print("Enter the message to encrypt: ")
	message := readLine()
	if message == "" {
		fmt.Println("Empty message, using default: \"Hello, world!\"")
		message = "Hello, world!"
	}

	maxCode := int64(0)
	for _, ch := range message {
		if int64(ch) > maxCode {
			maxCode = int64(ch)
		}
	}
	if n <= maxCode {
		fmt.Println("Using P=1009, Q=1013.")
		rsa, err = NewRSA(1009, 1013)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		e, n = rsa.PublicKey()
		d, _ = rsa.PrivateKey()
		fmt.Printf("New key: e=%d, n=%d, d=%d\n", e, n, d)
	}

	fmt.Printf("\nSource message: %s\n", message)

	encryptedMsg, err := rsa.EncryptString(message)
	if err != nil {
		fmt.Println("Error encryption:", err)
		return
	}
	fmt.Printf("Encrypted message (Base64): %s\n", encryptedMsg)

	decryptedMsg, err := rsa.DecryptString(encryptedMsg)
	if err != nil {
		fmt.Println("Error decryption:", err)
		os.Exit(1)
	}
	fmt.Printf("Decrypted message: %s\n", decryptedMsg)

	if decryptedMsg == message {
		fmt.Println("Verification passed: the text message was correctly decrypted.")
	} else {
		fmt.Println("Error: The decrypted text does not match")
	}
}