package csvreader

import (
    "encoding/csv"
    "fmt"
    "os"
)

func ReadCSV(fileName string) {
    // Abrir el archivo CSV
    file, err := os.Open(fileName)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // Crear un nuevo lector CSV
    reader := csv.NewReader(file)

    // Leer todas las filas del archivo
    rows, err := reader.ReadAll()
    if err != nil {
        panic(err)
    }

    // Imprimir cada fila
    for _, row := range rows {
        fmt.Println(row)
    }
}