package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/GiterLab/urllib"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type infoInical struct {

	url string
	concurrencia int
	solicitudes int
	rutaCarga string
	timeOut int


}

type strucData struct {
	Nombre string
	Departamento string
	Edad int
	Formadecontagio string
	Estado string

}

var peticionesEnviadas int =0
var flag bool = false
var timeCurrent = time.Now()

func main() {

	fmt.Println(". . . . iniciando")

	leerData()



}

func leerData(){

	var info infoInical
	var option int = -1

	for true {

		fmt.Println("------menu-----")
		fmt.Println("1. Ingresar url:")
		fmt.Println("2. Ingresar Concurrencia:")
		fmt.Println("3. Ingresar Solicitudes:")
		fmt.Println("4. Ingresar ruta del archivo que se desea cargar:")
		fmt.Println("5. Ingresar timeout:")
		fmt.Println("6. Aceptar")
		fmt.Println("7. Mostrar info ingresada:")
		fmt.Println("8. Cancelar")
		fmt.Println("9. Salir")

		fmt.Scanf("%d", &option)


		if option == 0{
			break
		}

		optionSwitch(option, &info)

	}

}

func optionSwitch(option int, info *infoInical) {
	switch option {
	case 1:
		readUrl(&info)
		break
	case 2:
		readConcurrence(&info)
		break
	case 3:
		readSolicitodes(&info)
		break
	case 4:
		readRuta(&info)
		break
	case 5:
		readTimeOut(&info)
		break
	case 6:
		execConcurrence(&info)
		break
	case 7:
		mostrarInfo(&info)
		break
	case 8:
		break

	}

}


func readUrl(info **infoInical) {
	fmt.Println("url...")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}

	(*info).url= scanner.Text()


}

func readConcurrence(info **infoInical){
	fmt.Println("no. concurrencia...")
	var concu int
	fmt.Scanf("%d", &concu)
	(*info).concurrencia = concu


}

func readSolicitodes(info **infoInical){

	fmt.Println("no. solicitudes...")
	var sol int
	fmt.Scanf("%d", &sol)

	(*info).solicitudes = sol



}

func readTimeOut(info **infoInical){

	fmt.Println("timeout...")
	var timeOut int
	fmt.Scanf("%d", &timeOut)

	(*info).timeOut = timeOut

}

func readRuta(info **infoInical){

	fmt.Println("ruta de carga...")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	(*info).rutaCarga = scanner.Text()

}

func mostrarInfo(info **infoInical)  {

	fmt.Println("url: "+ (*info).url )
	fmt.Println("concurrencia: ", (*info).concurrencia)
	fmt.Println("solicitudes: ", (*info).solicitudes)
	fmt.Println("ruta carga: "+ (*info).rutaCarga )
	fmt.Println("timeout: ", (*info).timeOut)

	option:=-1
	fmt.Scanf("%d", &option)

}

func validarNumero(value int, valorMenor int, valorMax int) int {


	if value < valorMenor || value > valorMax{
		return -1
	}

	return value
}

func execConcurrence(info **infoInical){

	var url string = (*info).url
	var urlParams string = (*info).rutaCarga
	var concurrence int = validarNumero((*info).concurrencia, 1, math.MaxInt64 )
	var solcitudes int = validarNumero((*info).solicitudes, 1, math.MaxInt64 )
	var timeWait int = validarNumero((*info).timeOut, 1 ,math.MaxInt64)


	if len(strings.TrimSpace(url)) == 0{

		println("La url no puede estar vacia")
		return
	}

	if len(strings.TrimSpace(urlParams)) == 0{

		println("La ruta de carga no puede estar vacia")
		return
	}

	if concurrence == -1{

		println("La concurrencia esta fuera de rango")
		return

	}

	if solcitudes == -1{

     	println("La solicitud esta fuera de rango")
		return

	}

	if timeWait == -1{

		println("La solicitud esta fuera de rango")
		return

	}

	var data[] strucData = getParams(urlParams)
	peticionExterna(concurrence, url, data, solcitudes, timeWait)





}

func peticionExterna(concurrence int, url string, data []strucData, solicitudes int, timeout int){

	peticionesEnviadas = 0
	flag = false


	for i:=0 ;i< concurrence; i++{

		go emular(i, url, data, solicitudes, timeout)

	}

	fmt.Println("fin de envios de datos")
	var input string
	fmt.Scanln(&input)

}

func emular(noPeticion int, url string, data [] strucData, total int, timesiu int)  {

	for peticionesEnviadas < total {

		var queryFinal string = url+"?"+getStringforRequest(data)
		//aca se enviara la petcion al servidro
		//str, err := urllib.Get("https://jsonplaceholder.typicode.com/users/1").String()

		println("No. peticion: ", noPeticion, " Request: ", queryFinal, " total peticiones: ", peticionesEnviadas)
		peticionesEnviadas+=1

	}
	time.Sleep((time.Second) )

}

func getStringforRequest(data []strucData) string  {
	var max int = len(data)

	if max == 0{
		println("Esta vacio, la deta de las peticiones")
		return "error"
	}

	var value int = rand.Intn(max- 0)+ 0
	var query string = "Nombre="+data[value].Nombre+"&Departamento="+data[value].Departamento+"&Edad="+strconv.Itoa(data[value].Edad) + "&Formadecontagio="+data[value].Formadecontagio+"&Estado="+data[value].Estado

	return query

}

func getParams(url string) []strucData {

	bytesLeidos, err := ioutil.ReadFile(url)

	if err != nil{
		fmt.Println(". . .error al leer archivo")

	}

	var packs []strucData
	//repleace porque si no no se va leer bien el JSON
	contenido := strings.ReplaceAll(string(bytesLeidos),"Forma de contagio","Formadecontagio")
	json.Unmarshal([]byte(contenido), &packs);
	return packs
}

func envio_data() {

	str, err := urllib.Get("https://jsonplaceholder.typicode.com/users/1").String()
	if err != nil {
		// error
	}
	fmt.Println(str)

}


