package main

// I will import the package fmt and in the future a json

import "fmt"

func main() {
	contactList := generateContacts()
	showContacts(contactList)
}

//I tried to use iota to create a estructure type such as a enum, but I changed my mind for now

/*
type SeniorityType int

 const(

	entryLevel= iota +1
	softwareEngineerI
	softwareEngineerII
	softwareEngineerIII
	technicalLead
	projectLead
	manager
	vp
	director
	cto

 )
*/

//Defined a type of this date structure person

type person struct {
	fullname        string
	country         string
	linkedinProfile string
	email           string
	speciality      string
	//Type		SeniorityType
	techs        string
	observations string
}

//Defined a type of this date structure contact

type contact struct {
	name     string
	contacts []person
}

//initialize contact( a slice of  person)

func generateContacts() contact {
	//initialize a person

	person1 := person{
		fullname:        "Laura Eroles",
		country:         "Argentina",
		linkedinProfile: "https://www.linkedin.com/in/mar%C3%ADa-laura-eroles-b1030173/",
		email:           "laura.eroles@test.com",
		speciality:      "Software Engineer",
		techs:           "Golag, Java",
		observations:    "trainee/junior profile",
	}
	person2 := person{
		fullname:        "Maria Babot",
		country:         "Argentina",
		linkedinProfile: "https://www.linkedin.com/in/mar%C3%ADa-laura-eroles-b1030173/",
		email:           "mariab@test.com",
		speciality:      "Software Engineer",
		techs:           "Go, C#",
		observations:    "Ssr profile",
	}
	person3 := person{
		fullname:        "Diego Lopez",
		country:         "Argentino",
		linkedinProfile: "",
		email:           "diegol@test.com",
		speciality:      "Software Engineer",
		techs:           "Golag, C#",
		observations:    "Senior profile",
	}

	contactList := contact{
		name:     "Agenda",
		contacts: []person{person1, person2, person3},
	}
	return contactList
}

func showContacts(oneContact contact) {
	fmt.Println(" The agenda have the following contacts: ")

	for i, agenda := range oneContact.contacts {
		fmt.Println("=======================================================")
		//Printf ---> format %d define where is going to be the placeholder for the next parameter
		fmt.Printf("Contacto N°: %d\n", i+1)
		fmt.Println("Fullname, " + agenda.fullname)
		fmt.Println("Country, " + agenda.country)
		fmt.Println("LinkedIn, " + agenda.linkedinProfile)
		fmt.Println("email, " + agenda.email)
		fmt.Println("speciality, " + agenda.speciality)
		fmt.Println(" tech, " + agenda.techs)
		fmt.Println("Observations, " + agenda.observations)
	}

}
