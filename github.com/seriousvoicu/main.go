package main

import (
	/*"encoding/json"
	"fmt"
	"io/ioutil"*/
	"core"
	"log"
	"objects"
	/*"os"
	"strconv"*/)

// indexHandler responds to requests with our greeting.
/*func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)

	//Todos los satelites especificados en el json

	var satellities Satellities
	json.Unmarshal(reqBody, &satellities)

	for i := 0; i < len(satellities.Satellities); i++ {
		fmt.Fprintf(w, "name: "+satellities.Satellities[i].Name)
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, "distance: "+strconv.Itoa(satellities.Satellities[i].Distance))
		fmt.Fprintf(w, "\n")
		fmt.Fprintf(w, "message: "+satellities.Satellities[i].Message[0])
		fmt.Fprintf(w, "\n")
	}
}
*/
func main() {

	/*	sateliteOne := &objects.Satellitie{Name: "jose", Distance: 0.66025, Message: []string{"", "", "mensaje", ""}, Pos: &core.Vector2{X: 0, Y: 8}}

		sateliteDos := &objects.Satellitie{Name: "Papa", Distance: 10, Message: []string{"", "", "secreto"}, Pos: &core.Vector2{X: 5, Y: 0}}

		sateliteTres := &objects.Satellitie{Name: "Ojota", Distance: 10, Message: []string{"mensaje", ""}, Pos: &core.Vector2{X: -5, Y: 0}}*/

	//sateliteCuatro := &Satellitie{Name: "FAFA", Distance: 100, Message: []string{"", "", ""}, Position: &Vector2{x: 100, y: -100}}

	/*sateliteOne := &Satellitie{Name: "jose", Distance: 100, Message: []string{"", "", "", "", "este", "", "", "mensaje", ""}, Position: &Vector2{x: -500, y: -200}}

	sateliteDos := &Satellitie{Name: "Papa", Distance: 100, Message: []string{"", "", "es", "un", "", "secreto"}, Position: &Vector2{x: 100, y: -100}}

	sateliteTres := &Satellitie{Name: "Ojota", Distance: 100, Message: []string{"", "", "", "es", "", "mensaje", ""}, Position: &Vector2{x: 100, y: -100}}

	sateliteCuatro := &Satellitie{Name: "FAFA", Distance: 100, Message: []string{"", "", "", "", "", "", ""}, Position: &Vector2{x: 100, y: -100}}*/

	//var satellities [2]*Satellitie
	/*var satArr [3]objects.Satellitie
	satArr[0] = sateliteOne
	satArr[1] = sateliteDos
	satArr[2] = sateliteTres*/

	var cluster *objects.SatellitieCluster

	cluster = &objects.SatellitieCluster{Satellities: []objects.Satellitie{
		objects.Satellitie{Name: "jose", Distance: 0.66025, Message: []string{"", "", "mensaje", ""}, Pos: &core.Vector2{X: 0, Y: 8}},
		objects.Satellitie{Name: "Papa", Distance: 10, Message: []string{"", "", "secreto"}, Pos: &core.Vector2{X: 5, Y: 0}},
		objects.Satellitie{Name: "Ojota", Distance: 10, Message: []string{"mensaje", ""}, Pos: &core.Vector2{X: -5, Y: 0}}}}

	log.Printf("Mensaje %s", cluster.GetLocation())

	log.Printf("Mensaje %s", cluster.GetMessage())

	//objects.SatellitieCluster := &objects.SatellitieCluster{Satellities: &[3]objects.Satellitie}
	//satellities[3] = sateliteCuatro

	/*log.Printf("Interseccion final %s", GetIntersection(satellities[:]))

	log.Printf("Mensaje %s", GetMessage(satellities[:]))*/

	/*http.HandleFunc("/", homepage)
	http.HandleFunc("/articles", createNewArticle)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
	*/
}

/*********************
Ensamblado de mensaje
*********************/
/*func GetMessage(messages ...[]string) (msg string) {
	return "asd"
}*/
/*
func GetMessage(satellities []*objects.Satellitie) []string {

	//Para evitar perder tiempo, doy por hecho que el desfasaje es solo al principio del mensaje
	//Solucion: Recorro de atras para adelante xD

	//El desfasaje implica que hay elementos de mas, por lo tanto el mas chicos de los mensajes tiene que representar el tamaño correcto
	//Si todos tienen desfasaje, va a dar error, porque va haber una hilera de elementos vacios, lo que nos deja un elemento vacio sin poder determinar que es
	//A su vez, ninguno de los elementos mas largos deben tener un valor valido en el area de "desfasaje", porque esto implica que hay perdida de informacion

	sort.Slice(satellities, func(i, j int) bool {
		return len(satellities[i].Message) < len(satellities[j].Message)
	})

	//Recorro elemento por elemento de cada uno de los arreglos, levantando el valor que le corresponde
	var message = MergeMessages(satellities[0].Message, satellities[1].Message)

	for i := 2; i < len(satellities); i++ {
		message = MergeMessages(message, satellities[i].Message)
	}

	return message
}

//Mergea dos mensajes y devuelve el mensaje mergeado
func MergeMessages(messageOne []string, messageTwo []string) []string {
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
			if (messageOne[len(messageOne)-i] != messageTwo[len(messageTwo)-i] && messageOne[len(messageOne)-i] != "" && messageTwo[len(messageTwo)-i] != "") || (messageOne[len(messageOne)-i] == "" && messageTwo[len(messageTwo)-i] == "") {
				return nil //No hay coincidencia
			}

			if messageOne[len(messageOne)-i] == "" {
				newMessage[len(newMessage)-i] = messageTwo[len(messageTwo)-i]
			} else {
				newMessage[len(newMessage)-i] = messageOne[len(messageOne)-i]
			}

			lastElement = i
		} else {
			//Valido que todos los elementos en la zona de "desfase" sean elementos vacíos
			if oneBiggest {
				if messageOne[len(messageOne)-i] != "" {
					return nil
				}
			} else {
				if messageTwo[len(messageTwo)-i] != "" {
					return nil
				}
			}
		}

	}

	return newMessage[:]
}
*/
/*********************
Satelite y calculo de inserseccion entre dos de ellos (hay que agregar los restantes)
*********************/
/*
// Due to double rounding precision the value passed into the Math.acos
// function may be outside its domain of [-1, +1] which would return
// the value NaN which we do not want.
func acossafe(x float64) float64 {
	if x >= +1.0 {
		return 0
	}
	if x <= -1.0 {
		return math.Pi
	}

	return math.Acos(x)
}

// Rotates a point about a fixed point at some angle 'a'
func rotatePoint(fp *objects.Vector2, pt *objects.Vector2, a float64) *objects.Vector2 {
	var x = pt.X - fp.X
	var y = pt.Y - fp.Y
	var xRot = x*math.Cos(a) + y*math.Sin(a)
	var yRot = y*math.Cos(a) - x*math.Sin(a)

	var point = new(objects.Vector2)
	point.X = fp.X + xRot
	point.Y = fp.Y + yRot

	return point
}

func GetLocation(satellities ...objects.Satellitie) (x, y float32) {

	return 2, 2
}

//Dados N satelites, evalua el punto donde todas las distancias coinciden
func GetIntersection(satellities []*objects.Satellitie) (point *objects.Vector2) {

	//Agarro los dos primeros satelites y calculo la interseccion
	//La idea es que dos satelites pueden coincidir en cero o 2 puntos
	intersectionOne, intersectionTwo := GetIntersections(satellities[0].Distance, satellities[0].Position, satellities[1].Distance, satellities[1].Position)

	var oneIntersects = true
	var twoIntersects = true

	var distance float64

	//Para los siguientes satelites es una cuestion de calcular las distancias entre cada uno y estos puntos
	if len(satellities) > 2 {
		for i := 2; i < len(satellities); i++ {
			if oneIntersects {
				distance = GetDistance(intersectionOne, satellities[i].Position)

				if math.Abs(distance-satellities[i].Distance) > 0.005 {
					oneIntersects = false
				}
			}

			if twoIntersects {
				distance = GetDistance(intersectionTwo, satellities[i].Position)
				if math.Abs(distance-satellities[i].Distance) > 0.005 {
					twoIntersects = false
				}
			}

			if oneIntersects == false && twoIntersects == false {
				return nil
			}
		}
	}

	if oneIntersects {
		return intersectionOne
	} else if twoIntersects {
		return intersectionTwo
	} else {
		return nil
	}
}

func GetIntersections(radOne float32, posOne *objects.Vector2, radTwo float32, posTwo *objects.Vector2) (intOne *objects.Vector2, intTwo *objects.Vector2) {

	var d = GetDistance(posOne, posTwo) // math.Hypot(posTwo.x-posOne.x, posTwo.y-posOne.y) //Distancia entre ambos puntos

	if d <= radOne+radTwo && d >= math.Abs(radTwo-radOne) {

		var ex = (posTwo.X - posOne.X) / d
		var ey = (posTwo.Y - posOne.Y) / d

		var x = (radOne*radOne - radTwo*radTwo + d*d) / (2 * d)
		var y = math.Sqrt(radOne*radOne - x*x)

		var P1, P2 *objects.Vector2
		P1 = &objects.Vector2{X: CapZero(posOne.X + x*ex - y*ey), Y: CapZero(posOne.Y + x*ey + y*ex)}
		P2 = &objects.Vector2{X: CapZero(posOne.X + x*ex + y*ey), Y: CapZero(posOne.Y + x*ey - y*ex)}

		return P1, P2
	}

	return nil, nil
}*/
