package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rotors = []string{
	"EKMFLGDQVZNTOWYHXUSPAIBRCJ", // Rotor I
	"AJDKSIRUXBLHWTMCQGZNPYFVOE", // Rotor II
	"BDFHJLCPRTXVZNYEIWGAKMUSQO", // Rotor III
	"ESOVPZJAYQUIRHXLNFTGKDCMWB", // Rotor IV
	"VZBRGITYUPSDNHLXAWMJQOFECK", // Rotor V
	"YRUHQSLDPXNGOKMIEBFZCWVJAT", // Reflector A
	"RDOBJNTKVEHMLFCWZAXGYIPSUQ", // Reflector B
}

var rotorsInv = []string{
	"UWYGADFPVZBECKMTHXSLRINQOJ", // Rotor I
	"AJPCZWRLFBDKOTYUQGENHXMIVS", // Rotor II
	"TAGBPCSDQEUFVNZHYIXJWLRKOM", // Rotor III
	"HZWVARTNLGUPXQCEJMBSKDYOIF", // Rotor IV
	"QCYLXWENFTZOSMVJUDKGIARPHB", // Rotor V
}

/**
 * cleanMessage function :
 * Use to clean encrypted message
 */
func cleanMessage(_message string) string {
	return strings.ToUpper(strings.ReplaceAll(_message, " ", ""))
}

/**
 * coincidenceIndex function :
 * Use to compute the coincidence index of a string
 */
func coincidenceIndex(_message string) float64 {
	var count [26]float64
	var coincidenceIndex float64
	for _, char := range _message {
		count[int(char)-65]++
	}
	for _, c := range count {
		coincidenceIndex += (c * (c - 1.0)) / (float64(len(_message)) * (float64(len(_message)) - 1.0))
	}
	return coincidenceIndex
}

/**
 * rotors increment function :
 * Use to increment rotors key values
 */
func rotorsIncr(_key [3]int, _rotors [3]int) [3]int {
	var notch = [5]int{16, 4, 21, 9, 25}
	if _key[1] == notch[_rotors[1]] {
		_key[0] = (_key[0] + 1) % 26
		_key[1] = (_key[1] + 1) % 26
	}
	if _key[2] == notch[_rotors[2]] {
		_key[1] = (_key[1] + 1) % 26
	}
	_key[2] = (_key[2] + 1) % 26
	return _key
}

/**
 * decrypt function :
 * Use to decrypt a message using a key
 */
func decrypt(_message string, _rotors [3]int, _ref int, _key [3]int) string {
	var builder strings.Builder

	for _, char := range _message {
		_key = rotorsIncr(_key, _rotors)
		var rd = (byte(rotors[_rotors[2]][(byte(char)-65+byte(_key[2])+26)%26]) - 65 + 26 - byte(_key[2])) % 26
		var rm = (byte(rotors[_rotors[1]][(rd+byte(_key[1])+26)%26]) - 65 + 26 - byte(_key[1])) % 26
		var rg = (byte(rotors[_rotors[0]][(rm+byte(_key[0])+26)%26]) - 65 + 26 - byte(_key[0])) % 26
		var r = byte(rotors[_ref][rg] - 65)

		var rg2 = (byte(rotorsInv[_rotors[0]][(r+byte(_key[0])+26)%26]) - 65 + 26 - byte(_key[0])) % 26
		var rm2 = (byte(rotorsInv[_rotors[1]][(rg2+byte(_key[1])+26)%26]) - 65 + 26 - byte(_key[1])) % 26
		var rd2 = (byte(rotorsInv[_rotors[2]][(rm2+byte(_key[2])+26)%26]) - 65 + 26 - byte(_key[2])) % 26
		builder.WriteRune(rune(rd2 + 65))
	}

	return builder.String()
}

/**
 * Function cryptanalys :
 * Use to try all possible key and show possible matches
 */
func cryptanalys(_message string) {
	var count = 0
	for rl := 0; rl < 5; rl++ {
		for rm := 0; rm < 5; rm++ {
			if rm != rl {
				for rr := 0; rr < 5; rr++ {
					if rr != rm && rr != rl {
						for ref := 5; ref < 7; ref++ {
							for kl := 0; kl < 26; kl++ {
								for km := 0; km < 26; km++ {
									for kr := 0; kr < 26; kr++ {
										count++
										var txt = decrypt(_message, [3]int{rl, rm, rr}, ref, [3]int{kl, km, kr})
										var ic = coincidenceIndex(txt)
										if ic > 0.0745 {
											fmt.Print("\n", txt, "\n IC -> ", ic, " | Rotors -> ", rl+1, rm+1, rr+1, " | Reflector -> ", ref-4, " | Key -> ", kl, km, kr, "\n")
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

/**
 * main function
 */
func main() {
	// Configuration variables
	rotorLeft := []int{1, 2, 3, 4, 5}
	var rotors = [3]int{-1, -1, -1}
	var ref = -1
	var key = [3]int{-1, -1, -1}

	// Get input from user
	scanner := bufio.NewScanner(os.Stdin)

	// Get the rotors
	for i := 0; i < 3; i++ {
		for rotors[i] == -1 {
			fmt.Println("Séléctionnez le rotor n°", i+1, " parmis ", rotorLeft)
			scanner.Scan()
			value, err := strconv.Atoi(scanner.Text())
			if err == nil {
				var pos = find(rotorLeft, value)
				if pos != -1 {
					rotorLeft = append(rotorLeft[:pos], rotorLeft[pos+1:]...)
					rotors[i] = value - 1
				} else {
					fmt.Print("Valeur incorrecte !\n\n")
				}
			} else {
				fmt.Print("Valeur incorrecte !\n\n")
			}
		}
	}

	// Get the reflector
	for ref == -1 {
		fmt.Println("Séléctionnez le réflecteur (1 ou 2)")
		scanner.Scan()
		value, err := strconv.Atoi(scanner.Text())
		if err == nil {
			if value > 0 && value < 3 {
				ref = value + 4
			} else {
				fmt.Print("Valeur incorrecte !\n\n")
			}
		} else {
			fmt.Print("Valeur incorrecte !\n\n")
		}
	}

	// Get the key
	for i := 0; i < 3; i++ {
		for key[i] == -1 {
			fmt.Println("Séléctionnez la clé n°", i+1, " (nombre entre 0 et 25)")
			scanner.Scan()
			value, err := strconv.Atoi(scanner.Text())
			if err == nil {
				if value >= 0 && value <= 25 {
					key[i] = value
				} else {
					fmt.Print("Valeur incorrecte !\n\n")
				}
			} else {
				fmt.Print("Valeur incorrecte !\n\n")
			}
		}
	}

	// Break the code
	fmt.Println("Entrer le texte")
	scanner.Scan()
	line := scanner.Text()
	var encrypt = decrypt(cleanMessage(line), rotors, ref, key)
	fmt.Println(encrypt)
	cryptanalys(encrypt)
}

/**
 * Function find :
 * Use to find an element inside a slice
 */
func find(slice []int, value int) int {
	for i, item := range slice {
		if item == value {
			return i
		}
	}
	return -1
}
