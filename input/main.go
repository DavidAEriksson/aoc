/*
AoC 2022 Utility functions for puzzle input
*/
package input

import (
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

func init() {

	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

/*
Get input of given day.
*/
func Get(day_count string) string {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://adventofcode.com/2022/day/"+day_count+"/input", nil)
	if err != nil {
		log.Fatal(err)

	}
	req.Header.Add("Cookie", "session="+os.Getenv("SESSION"))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}
	return string(body)
}
