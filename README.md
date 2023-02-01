# Email Domain Counter

This project provides a tool to count the number of customers with email addresses for each domain in a CSV file.
## Requirements

The following Go packages are required:

    encoding/csv
    os
    sort
    strings
    fmt

## Functionality

The project consists of two main functions: ReadCSV and CountEmailDomains. The ReadCSV function reads a CSV file and passes the rows to the CountEmailDomains function. The CountEmailDomains function counts the number of customers with email addresses for each domain and returns a map. The keys in the map represent the email domains and the values represent the number of customers with email addresses for each domain.

The ReadCSV function opens the CSV file, reads all the rows from the file, and passes the rows to the CountEmailDomains function. The function then sorts the keys in the map returned by CountEmailDomains and prints the results to the console.

Any errors that occur during the execution of the ReadCSV function are logged.
## Usage

To use the Email Domain Counter, run the ReadCSV function and pass the path to the CSV file as an argument. The function will read the file, count the email domains, and print the results to the console.

Example:
    
    go run main.go

Output:

    123-reg.co.uk: 8
    163.com: 6
    1688.com: 3
    1und1.de: 5
    360.cn: 6
    4shared.com: 5
    51.la: 4
    a8.net: 6
    abc.net.au: 7

The expected output will be a list of email domains and the number of customers with email addresses for each domain, sorted in alphabetical order.
## Testing

The project includes unit tests for the CountEmailDomains function. To run the tests, use the following command:

    go test
