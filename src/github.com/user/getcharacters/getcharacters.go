package main

import (
        "gopkg.in/resty.v1"
        "github.com/fatih/color"
        "github.com/o1egl/govatar"
	"encoding/json"
	"strconv"
)

type Person struct {
    Name string
    Height string
    Mass string
    Hair_color string
    Skin_color string
    Eye_color string
    Birth_year string
    Gender string
    Homeworld string
    Films []string
    Species []string
    Vehicles []string
    Starships []string
    Created string
    Edited string
    Url string
}

type PaginatedPeopleResponse struct {
    Count int
    Next_url string
    Prev_url string
    Results []Person
}

func main() {

    people := getStarWarsCharacters(1)

    printPeople(people)

}

func getStarWarsCharacters(page int) []Person {

resty.SetHTTPMode()

        var paginatedPeopleResponse PaginatedPeopleResponse

        resp, err := resty.R().Get("https://swapi.co/api/people?page=" + strconv.Itoa(page))

        if (err != nil || resp.StatusCode() != 200) {

            color.Red("Error: %s", err)

        } else {

            if err := json.Unmarshal(resp.Body(), &paginatedPeopleResponse); err != nil {
                color.Red("Error: %s", err)
            }   

        } 

        return paginatedPeopleResponse.Results
}

func printPeople(people []Person) {
    for i := 0; i < len(people); i++ {
        person := people[i]
        name := person.Name
        govatar.GenerateFile(govatar.MALE, "avatars/" + person.Name + ".jpg")
        if (person.Gender == "female") {
            color.Red(name)
        } else {
            color.Blue(name)
        }   
    } 
}
