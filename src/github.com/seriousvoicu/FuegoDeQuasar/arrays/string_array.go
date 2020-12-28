package arrays

import "github.com/seriousvoicu/FuegoDeQuasar/exestate"

/*
//Indice de la primera aparicion de 'value' desde 'startIndex' con barrido a la derecha
func GetFirstRight(array []string, startIndex int, value string) (int, *exestate.State) {

	for ii := startIndex; ii <= len(array); ii++ {
		if array[ii] == value {
			return ii, exestate.Ok()
		}
	}

	return 0, exestate.ControlledError("No se encontro elemento (arrays.string_array.GetFirstRight)")
}

//Indice de la primera aparicion de 'value' desde 'startIndex' con barrido a la izquierda
func GetFirstLeft(array []string, startIndex int, value string) (int, *exestate.State) {

	for ii := startIndex; ii >= 0; ii-- {
		if array[ii] == value {
			return ii, exestate.Ok()
		}
	}

	return 0, exestate.ControlledError("No se encontro elemento (arrays.string_array.GetFirstLeft)")
}

//Devuelve los indices de 'arrayA' y 'arrayB' donde los elementos coinciden con barrido a la derecha
func GetMatchIndexesRight(arrayA []string, arrayB []string) (int, int, *exestate.State) {

	for ii := 0; ii < len(arrayA); ii++ {
		index, state := GetFirstRight(arrayB, ii, arrayA[ii])

		if state.IsOk() {
			return ii, index, exestate.Ok()
		}
	}

	return 0, 0, exestate.ControlledError("No se pudo match (arrays.string_array.GetMatchIndexRight)")
}

//Devuelve los indices de 'arrayA' y 'arrayB' donde los elementos coinciden con barrido a la izquierda
func GetMatchIndexesLeft(arrayA []string, arrayB []string) (int, int, *exestate.State) {

	for ii := len(arrayA); ii > 0; ii-- {
		index, state := GetFirstLeft(arrayB, ii, arrayA[ii])

		if state.IsOk() {
			return ii, index, exestate.Ok()
		}
	}

	return 0, 0, exestate.ControlledError("No se pudo match (arrays.string_array.GetMatchIndexLeft)")
}

//Mergea dos arreglos 'arrayA' y 'arrayB' bajo las siguientes premisas:
//En los mismos indices en cada arreglo debe darse la combinacion: Iguales valores, valor y vacio o dos vacios
//Si se encuentran dos valores distintos (descartando espacios y vacio) retorna error
func FullMatchMerge(arrayA []string, arrayB []string) (*string, *exestate.State) {
	return nil, exestate.Ok()
}

func DirtyMatcheMerge(arrayA []string, arrayB []string) (*[]string, *exestate.State) {
	var smaller *[]string
	var other *[]string

	if len(arrayA) < len(arrayB) {
		smaller = &arrayA
		other = &arrayB
	} else {
		smaller = &arrayB
		other = &arrayA
	}

	var leftOffset bool
	var rightOffset bool

	smallIndex, otherIndex, state = GetMatchIndexesRight(arrayA, arrayB)

	smallIndex, otherIndex, state = GetMatchIndexesLeft(arrayA, arrayB)

	//Verifico si existe un match en el primer o ultimo indice
	if (*smaller)[0] == (*other)[0] {
		rightOffset = true
	}
	if (*smaller)[len(*smaller)-1] == (*other)[len(*other)-1] {
		leftOffset = true
	}

	var smallIndex int
	var otherIndex int
	var state *exestate.State

	//Busco un punto de anclaje, excluyendo los extremos
	//Los puntos de anclaje determinan elementos que no pueden ser desplazados
	if rightOffset || (!rightOffset && !leftOffset) {

	} else if leftOffset {

	}



	return nil, exestate.Ok()
}
*/
//Genera un arreglo a partir de dos arreglos
//Solo verifica las coincidencias
//Los espacios en blanco coincidentes son eliminados
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
