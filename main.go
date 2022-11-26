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

//Defined a type of ths date structure person

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

//Defined a type of ths date structure contact

type contact struct {
	name     string
	contacts []person
}

//Inicializar contact( a slice of  person)

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
	/*for i, ingrediente := range unaReceta.Ingredientes {
		fmt.Printf("%d - %s - comprar %v %s - Es indispensable? %v\n", i+1, ingrediente.Nombre, ingrediente.Cantidad, ingrediente.Unidad, ingrediente.Indispensable)
	}
	*/

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

/*func showPerson(aPerson person) {
	fmt.Println("Fullname: " + aPerson.fullname)
	fmt.Println("Country: " + aPerson.country)
	fmt.Println("LinkedIn Profile : " + aPerson.linkedinProfile)
	fmt.Println("email: " + aPerson.email)
	fmt.Println("Speciality: " + aPerson.speciality)
	fmt.Println("Techs: " + aPerson.techs)

}*/
