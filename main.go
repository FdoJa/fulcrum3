package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	pb "github.com/FdoJa/ServidoresFulcrum/proto"
	"google.golang.org/grpc"
)

// Variable asociada al identificador del servidor
var fulcrumId = 2 // Dado que un array empieza desde 0 se tiene el mapeo -> Fulcrum1 = 0 ; Fulcrum2 = 1 ; Fulcrum3 = 2
var timeVector = [3]int{0, 0, 0}

type fulcrumServer struct {
	pb.UnimplementedInformantesServer
}

type consistenciaServer struct {
	pb.UnimplementedConsistenciaServer
}

func (s *fulcrumServer) AgregarBase(ctx context.Context, base *pb.Base) (*pb.Vector, error) {
	log.Printf("- Comando recibido: AgregarBase %s %s %s \n", base.Sector, base.Base, base.Soldados)

	agregarDato(base.Sector, base.Base, base.Soldados)

	timeVector[fulcrumId]++

	return &pb.Vector{
		S1: int32(timeVector[0]),
		S2: int32(timeVector[1]),
		S3: int32(timeVector[2]),
	}, nil
}

func agregarDato(sector string, base string, soldados string) {
	filePath := "/app/" + sector + ".txt"

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Error al leer archivo: %v", err)
	}
	defer file.Close()

	data := sector + " " + base + " " + soldados + " \n"
	_, err = file.WriteString(data)
	if err != nil {
		log.Fatalf("Error al escribir en el archivo: %v", err)
	}

	// Luego de realizar la correcta modificación, pasar a colocarla en los logs
	logLine := "AgregarBase " + sector + " " + base + " " + soldados + " \n"
	agregarLog(logLine)
}

func (s *fulcrumServer) RenombrarBase(ctx context.Context, nuevaBase *pb.BaseModificada) (*pb.Vector, error) {
	log.Printf("- Comando recibido: RenombrarBase %s %s %s \n", nuevaBase.Sector, nuevaBase.Base, nuevaBase.ActualizacionBase)

	modificarBase(nuevaBase.Sector, nuevaBase.Base, nuevaBase.ActualizacionBase)

	timeVector[fulcrumId]++

	return &pb.Vector{
		S1: int32(timeVector[0]),
		S2: int32(timeVector[1]),
		S3: int32(timeVector[2]),
	}, nil
}

func modificarBase(sector string, base string, nuevaBase string) {
	filePath := "/app/" + sector + ".txt"

	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		// En caso de que no existe el archivo, se crea
		file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Fatalf("Error al abrir o crear el archivo: %v", err)
		}
		defer file.Close()

		data := sector + " " + nuevaBase + " 0 \n"
		_, err = file.WriteString(data)
		if err != nil {
			log.Fatalf("Error al escribir en el archivo: %v", err)
		}

		logLine := "RenombrarBase " + sector + " " + base + " " + nuevaBase + " \n"
		agregarLog(logLine)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var text []string

	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Fields(line)

		baseTxt := data[1]

		if baseTxt == base {
			data[1] = nuevaBase
		}

		newLine := strings.Join(data, " ")
		text = append(text, newLine)
	}

	file.Seek(0, 0)
	file.Truncate(0)

	for _, linea := range text {
		_, err := file.WriteString(linea + " \n")
		if err != nil {
			log.Fatalf("Error al modificar archivo: %v", err)
		}
	}

	// Luego de realizar la correcta modificación, pasar a colocarla en los logs
	logLine := "RenombrarBase " + sector + " " + base + " " + nuevaBase + " \n"
	agregarLog(logLine)
}

func (s *fulcrumServer) ActualizarValor(ctx context.Context, actualizar *pb.ActualizarSoldados) (*pb.Vector, error) {
	log.Printf("Comando recibido: ActualizarValor  %s %s %s \n", actualizar.Sector, actualizar.Base, actualizar.ActualizacionSoldados)

	actualizarCantidad(actualizar.Sector, actualizar.Base, actualizar.ActualizacionSoldados)

	timeVector[fulcrumId]++

	return &pb.Vector{
		S1: int32(timeVector[0]),
		S2: int32(timeVector[1]),
		S3: int32(timeVector[2]),
	}, nil
}

func actualizarCantidad(sector string, base string, soldados string) {
	filePath := "/app/" + sector + ".txt"

	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		// En caso de que no existe el archivo, se crea
		file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Fatalf("Error al abrir o crear el archivo: %v", err)
		}
		defer file.Close()

		data := sector + " " + base + " " + soldados + " \n"
		_, err = file.WriteString(data)
		if err != nil {
			log.Fatalf("Error al escribir en el archivo: %v", err)
		}

		logLine := "ActualizarValor " + sector + " " + base + " " + soldados + " \n"
		agregarLog(logLine)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var text []string

	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Fields(line)

		baseTxt := data[1]
		if baseTxt == base {
			data[2] = soldados
		}

		newLine := strings.Join(data, " ")
		text = append(text, newLine)
	}

	file.Seek(0, 0)
	file.Truncate(0)

	for _, linea := range text {
		_, err := file.WriteString(linea + " \n")
		if err != nil {
			log.Fatalf("Error al modificar archivo: %v", err)
		}
	}

	// Luego de realizar la correcta modificación, pasar a colocarla en los logs
	logLine := "ActualizarValor " + sector + " " + base + " " + soldados + " \n"
	agregarLog(logLine)
}

func (s *fulcrumServer) BorrarBase(ctx context.Context, borrar *pb.Base) (*pb.Vector, error) {
	log.Printf("Comando recibido: BorrarBase  %s %s \n", borrar.Base, borrar.Sector)

	borrarDato(borrar.Sector, borrar.Base)

	timeVector[fulcrumId]++

	return &pb.Vector{
		S1: int32(timeVector[0]),
		S2: int32(timeVector[1]),
		S3: int32(timeVector[2]),
	}, nil
}

func borrarDato(sector string, base string) {
	filePathData := "/app/" + sector + ".txt"

	fileData, err := os.OpenFile(filePathData, os.O_RDWR, 0644)
	if err != nil {
		logLine := "BorrarBase " + sector + " " + base + " \n"
		agregarLog(logLine)
		return
		//log.Fatalf("Error al leer archivo: %v", err)
	}
	defer fileData.Close()

	scanner := bufio.NewScanner(fileData)
	var text []string

	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Fields(line)

		baseTxt := data[1]
		if baseTxt == base {
			continue
		}

		newLine := strings.Join(data, " ")
		text = append(text, newLine)
	}

	fileData.Seek(0, 0)
	fileData.Truncate(0)

	for _, linea := range text {
		_, err := fileData.WriteString(linea + " \n")
		if err != nil {
			log.Fatalf("Error al modificar archivo: %v", err)
		}
	}

	// Luego de realizar la correcta modificación, pasar a colocarla en los logs
	logLine := "BorrarBase " + sector + " " + base + " \n"
	agregarLog(logLine)
}

func agregarLog(datos string) {
	filePathLog := "/app/log.txt"

	fileLog, err := os.OpenFile(filePathLog, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Error en manejo de logs: %v", err)
	}
	defer fileLog.Close()

	currentTime := time.Now()

	logLine := currentTime.Format("15:04:05.000000") + " " + datos
	_, err = fileLog.WriteString(logLine)
	if err != nil {
		log.Fatalf("Error al escribir en el archivo de logs: %v", err)
	}
}

func (s *consistenciaServer) ConseguirLogs(ctx context.Context, ok *pb.Recepcion) (*pb.LogList, error) {
	filePathLog := "/app/log.txt"

	fileLog, err := os.OpenFile(filePathLog, os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("Error en manejo de logs: %v", err)
	}
	defer fileLog.Close()

	scanner := bufio.NewScanner(fileLog)

	var listaDatos []*pb.Log

	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Fields(line)

		tiempo := data[0]
		accion := data[1]
		sectorAfectado := data[2]
		baseAfectada := data[3]
		var nuevoValor string

		if accion == "BorrarBase" {
			nuevoValor = "0"
		} else {
			nuevoValor = data[4]
		}

		dato := &pb.Log{
			Tiempo:         tiempo,
			Accion:         accion,
			SectorAfectado: sectorAfectado,
			BaseAfectada:   baseAfectada,
			NuevoValor:     nuevoValor,
		}

		listaDatos = append(listaDatos, dato)
	}

	var vec []int32
	for _, v := range timeVector {
		vec = append(vec, int32(v))
	}

	response := &pb.LogList{
		ListaLogs: listaDatos,
		Vector:    vec,
	}
	return response, nil
}

func (s *consistenciaServer) EnviarDatosActualizados(ctx context.Context, datos *pb.Datos) (*pb.Recepcion, error) {
	log.Printf("- Llegaron datos del merge")

	sectoresRevisados := make(map[string]bool)

	for _, dato := range datos.ListaBases {
		filePath := "/app/" + dato.Sector + ".txt"

		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Fatalf("Error al leer archivo: %v", err)
		}
		defer file.Close()

		if !sectoresRevisados[dato.Sector] {
			sectoresRevisados[dato.Sector] = true

			file.Seek(0, 0)
			file.Truncate(0)
		}

		_, err = file.WriteString(dato.Sector + " " + dato.Base + " " + dato.Soldados + " \n")
		if err != nil {
			log.Fatalf("Error al escribir en el archivo: %v", err)
		}
	}

	timeVector[0] = int(datos.Vector[0])
	timeVector[1] = int(datos.Vector[1])
	timeVector[2] = int(datos.Vector[2])

	fileLogPath := "/app/log.txt"
	err := os.Remove(fileLogPath)
	if err != nil {
		log.Fatalf("Error borrar logs: %v", err)
	}

	return &pb.Recepcion{
		Ok: "Ok",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("Fallo en escuchar: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterConsistenciaServer(s, &consistenciaServer{})
	pb.RegisterInformantesServer(s, &fulcrumServer{})

	fmt.Println("Servidor Fulcrum 3 escuchando")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Fallo en serve: %v", err)
	}
}
