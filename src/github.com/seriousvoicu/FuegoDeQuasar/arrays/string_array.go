package arrays

import "github.com/seriousvoicu/FuegoDeQuasar/exestate"

//Genera un arreglo a partir de dos arreglos
//Solo verifica las coincidencias
func MergeStringArrays(messageOne []string, messageTwo []string) ([]string, *exestate.State) {
	var lastElement = 0
	var biggestSize = 0
	var oneBiggest = true

	//Agarro el elemento mas chico
	if len(messageOne) > len(messageTwo) {
		lastElement = len(messageTwo) - 1
		biggestSize = len(messageOne)
		oneBiggest = true
	} else {
		lastElement = len(messageOne) - 1
		biggestSize = len(messageTwo)
		oneBiggest = false
	}

	var newMessage []string = make([]string, lastElement+1)

	//Recorro de atras para delante todos los elementos
	for i := 1; i <= biggestSize; i++ {

		if i <= len(newMessage) {
			if messageOne[len(messageOne)-i] != messageTwo[len(messageTwo)-i] && messageOne[len(messageOne)-i] != "" && messageTwo[len(messageTwo)-i] != "" {
				return nil, exestate.ControlledError("No hay coincidencias (arrays.string_array)")
			}

			if messageOne[len(messageOne)-i] != "" || messageTwo[len(messageTwo)-i] != "" {
				if messageOne[len(messageOne)-i] == "" {
					newMessage[len(newMessage)-i] = messageTwo[len(messageTwo)-i]
				} else {
					newMessage[len(newMessage)-i] = messageOne[len(messageOne)-i]
				}
			}

			lastElement = i
		} else {
			//Valido que todos los elementos en la zona de "desfase" sean elementos vacíos
			if oneBiggest {
				if messageOne[len(messageOne)-i] != "" {
					return nil, exestate.ControlledError("No se pudo mergear (arrays.string_array)")
				}
			} else {
				if messageTwo[len(messageTwo)-i] != "" {
					return nil, exestate.ControlledError("No se pudo mergear (arrays.string_array)")
				}
			}
		}

	}

	return newMessage[:], exestate.Ok()
}

func StringArrayToString(array []string) (string, *exestate.State) {
	msg := ""

	for i := 0; i < len(array); i++ {
		msg += array[i]

		if i != len(array)-1 {
			msg += " "
		}
	}

	return msg, exestate.Ok()
}
