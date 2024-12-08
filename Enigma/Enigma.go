package main

import (
	"fmt"
	"strings"
)

var rotorI = "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
var rotorII = "AJDKSIRUXBLHWTMCQGZNPYFVOE"
var rotorIII = "BDFHJLCPRTXVZNYEIWGAKMUSQO"
var reflectorB = "YRUHQSLDPXNGOKMIEBFZCWVJAT"
var notchI, notchII, notchIII = 'Q', 'E', 'V'

func forwardThroughRotor(rotor string, pos int, char rune) rune {
	offset := (int(char-'A') + pos) % 26
	/* fmt.Println("Alphapetical Index of char", int(char-'A'), "Current rotor Postion", pos,
	"Finding the mapped charcter", offset) */
	return rune(rotor[offset])
}

func backwardThroughRotor(rotor string, pos int, char rune) rune {
	index := strings.IndexRune(rotor, char)
	return rune('A' + ((index - pos + 26) % 26))
}

func reflect(reflector string, char rune) rune {
	index := strings.IndexRune("ABCDEFGHIJKLMNOPQRSTUVWXYZ", char)
	return rune(reflector[index])
}

func stepRotors(posI, posII, posIII *int) {
	*posI = (*posI + 1) % 26
	if *posI == int(notchI-'A') {
		*posII = (*posII + 1) % 26
		if *posII == int(notchII-'A') {
			*posIII = (*posIII + 1) % 26
		}
	}
}

func enigmaEncrypt(message string, posI, posII, posIII int) string {
	var encryptedMessage string
	for _, char := range message {
		if char < 'A' || char > 'Z' {
			continue
		}

		stepRotors(&posI, &posII, &posIII)

		c1 := forwardThroughRotor(rotorI, posI, char)
		c2 := forwardThroughRotor(rotorII, posII, c1)
		c3 := forwardThroughRotor(rotorIII, posIII, c2)
		reflected := reflect(reflectorB, c3)
		c4 := backwardThroughRotor(rotorIII, posIII, reflected)
		c5 := backwardThroughRotor(rotorII, posII, c4)
		c6 := backwardThroughRotor(rotorI, posI, c5)

		encryptedMessage += string(c6)
	}

	return encryptedMessage
}

func main() {
	initialPosI, initialPosII, initialPosIII := 0, 0, 0

	message := "MARAM"

	encrypted := enigmaEncrypt(message, initialPosI, initialPosII, initialPosIII)
	fmt.Println("Encrypted message:", encrypted)

	decrypted := enigmaEncrypt(encrypted, initialPosI, initialPosII, initialPosIII)
	fmt.Println("Decrypted message:", decrypted)
}
