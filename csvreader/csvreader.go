package csvreader

import (
    "encoding/csv"
    "fmt"
    "os"
	"strings"
	"sort"
    "log"
)

func ReadCSV(fileName string) {
    file, err := os.Open(fileName)
    if err != nil {
        panic(err)
    }

    defer file.Close()
    reader := csv.NewReader(file)
	_, err = reader.Read()

    if err != nil {
        log.Println("Error reading the first line of the CSV file:", err)
        return
    }

    rows, err := reader.ReadAll()
    if err != nil {
        log.Println("Error reading line of the CSV file:", err)
        return
    }

	emailDomains := CountEmailDomains(rows)
    var keys []string
    for key := range emailDomains {
        keys = append(keys, key)
    }

    sort.Strings(keys)
    fmt.Println("Email domain count:")
    for _, key := range keys {
        fmt.Printf("%s: %d\n", key, emailDomains[key])
    }
}

func CountEmailDomains(rows [][]string) map[string]int {
    emailDomains := make(map[string]int)

    for _, row := range rows {
        email := row[2]
        domainParts := strings.Split(email, "@")
        if len(domainParts) != 2 {
            continue
        }
        domain := domainParts[1]
        emailDomains[domain]++
    }

    return emailDomains
}
