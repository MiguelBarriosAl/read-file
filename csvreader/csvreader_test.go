package csvreader

import (
    "reflect"
    "testing"
)

func TestCountEmailDomains(t *testing.T) {
    rows := [][]string{
        {"first_name", "last_name", "email", "gender", "ip_address"},
        {"Mildred", "Hernandez", "mhernandez0@github.io", "Female", "38.194.51.128"},
        {"Bonnie", "Ortiz", "bortiz1@cyberchimps.com", "Female", "197.54.209.129"},
        {"Dennis", "Henry", "dhenry2@hubpages.com", "Male", "155.75.186.217"},
        {"Sarah", "Davis", "sdavis2@hubpages.com", "Female", "135.75.234.234"},
    }
    expected := map[string]int{
        "github.io": 1,
        "cyberchimps.com": 1,
        "hubpages.com": 2,
    }
    actual := CountEmailDomains(rows)
    if !reflect.DeepEqual(actual, expected) {
        t.Errorf("Expected %v, but got %v", expected, actual)
    }
}
