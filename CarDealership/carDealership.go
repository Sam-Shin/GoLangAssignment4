package CarDealership

import (
	"fmt"
	"time"
)

//initializing the objects into the complex map
var carMap map[owner]car = map[owner]car{obj1: obj1.ownerCar, obj2: obj2.ownerCar}


var obj1 owner = owner{firstName: "Sam", lastName: "Shin", ownersInsurance: insurance{insuranceCompany: "Geico",
	phoneNumber: "949-123-4567", priceInsurance: "235", address: "54 Kellogg Drive"},
	ownerCar: car{price: "70000", manufacturer: "Mercedes", model: "E350"}}

var obj2 owner = owner{firstName: "Sam", lastName: "Shin", ownersInsurance: insurance{insuranceCompany: "Geico",
	phoneNumber: "949-123-4567", priceInsurance: "235", address: "54 Kellogg Drive"},
	ownerCar: car{price: "70000", manufacturer: "Mercedes", model: "E350"}}

type car struct{
	price string
	manufacturer string
	model string
}

type owner struct{
	firstName string
	lastName string
	ownersInsurance insurance
	ownerCar car
}

type insurance struct{
	insuranceCompany string
	phoneNumber string
	priceInsurance string
	address string
}

func (i *car) getCarName() string{
	return i.manufacturer + " " + i.model
}

//Prints ALL data
func query(){
	for i,_ := range carMap{
		fmt.Println("\nOwner:",i.firstName, i.lastName, "\nInsurance Company:", i.ownersInsurance.insuranceCompany,
			        "\nPrice of insurance: $", i.ownersInsurance.priceInsurance,"\nPhone number:", i.ownersInsurance.phoneNumber,
			        "\n Address: ", i.ownersInsurance.address)
		fmt.Println("Car Information:\nManufacturer:", carMap[i].manufacturer, "\nModel:", carMap[i].model, "\nPrice:",
			         carMap[i].price)
	}
}

//Use list/array for the stock of cars only
func stock() {
	fmt.Println("\nStock of all available cars")
	for i := range carMap{
		fmt.Println(carMap[i].manufacturer, carMap[i].model, "$",carMap[i].price)
	}
}
//using the first and last name to find the key to delete from the carmap
func deleteCar(fname, lname string) {
	for i, _ := range carMap{
		if (fname == i.firstName) && (lname == i.lastName){
			delete(carMap, i)
		}
	}
}

func addNew(fName, lName, addy, phoneNum, insuranceName, manufactName, modelName, cPrice, iPrice string) {
	newObject := owner{firstName: fName, lastName: lName, ownersInsurance: insurance{insuranceCompany: insuranceName,
		phoneNumber: phoneNum, priceInsurance: iPrice, address: addy}, ownerCar: car{price: cPrice, manufacturer: manufactName, model: modelName}}
	carMap[newObject] = newObject.ownerCar
}

//get user input, then calls the deleteCar function at the end of this function
func sell() {
	var fname, lname string
	fmt.Print("\nEnter the name of the person you would like to delete: ")
	fmt.Scan(&fname, &lname)
	deleteCar(fname,lname)
	fmt.Print("Done.")
}
//get user input, then calls addNew function at the end of this function
func buy(){
	var fName, lName, addy, phoneNum, insuranceName, manufactName, modelName, cPrice, iPrice string
	fmt.Print("\nEnter the first name of the new person: ")
	fmt.Scanln(&fName)
	fmt.Print("Enter the last name of the new person: ")
	fmt.Scanln(&lName)
	fmt.Print("Enter ", fName, " ",lName, "'s insurance company: ")
	fmt.Scanln(&insuranceName)
	fmt.Print("Enter the price of the insurance: ")
	fmt.Scanln(&iPrice)
	fmt.Print("Enter ", fName, " ",lName, "'s address (NO SPACES): ")
	//for this the user would need to input "123FakeAddress" instead of "123 Fake Address")

	//i'm not sure how to get user input with a string that has spaces (i.e. "123 Fake Street"), the Scanln only takes each word
	//but not the entire string. I understand the spaces are delimiters. Next time I will try to implement "bufio" as I did some research
	//and found that bufio could be a solution to this problem
	fmt.Scanln(&addy)
	fmt.Print("Enter ", fName, " ",lName, "'s phone number: ")
	fmt.Scanln(&phoneNum)

	fmt.Print("Enter ", fName, " ",lName, "'s car manufacturer: ")
	fmt.Scanln(&manufactName)
	fmt.Print("Enter ", fName, " ",lName, "'s car model: ")
	fmt.Scanln(&modelName)
	fmt.Print("Enter the price of the car: ")
	fmt.Scanln(&cPrice)
	addNew(fName, lName, addy, phoneNum, insuranceName, manufactName, modelName, cPrice,iPrice)
}

//this is the ticker that prints "ping" every second.
func backgroundTask() {
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C{
		fmt.Println("Ping")
	}
}
//use breakpoint, when user quits program, it also breaks the ticker
func Main1(){
	//loop breakpoint, giving user to choose what happens
loop:
	for {
		go backgroundTask()                     //you can comment this out if you want to run the program normally. this calls the function above
		fmt.Println("\nEnter '1' to buy a car\nEnter '2' to sell a car\nEnter '3' to list the stock of cars available"+
			"\nEnter '4' to list the query\nEnter '5' to quit")
		var x int
		fmt.Scanf("%d", &x)
		switch x {
		case 1:
			buy()
			break
		case 2:
			sell()
			break
		case 3:
			stock()
			break
		case 4:
			query()
			break
		case 5:
			fmt.Println("Good bye")
			break loop
		default:
			fmt.Println("Invalid input, please try again")
		}
	}
}

/*
As stated in the buy() function, for the next project I will better implement getting user input that has spaces.
I realize that the fmt.Scan() function does not work well when getting strings with spaces.

you can comment out go backGroundTask() to run the program normally. I decided to have user input with the breakpoint loop
in this project and having the ticker function print "ping" every second was very interesting. Next time, I could plan better
on the design of this program somehow to make the ticker not cause as much of a distraction and interrupting the program.
Thanks Puneet! - Sam Shin
 */