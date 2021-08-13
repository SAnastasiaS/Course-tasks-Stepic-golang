package main
import "fmt"

type GamePerson struct{
	On bool
	Ammo, Power int
}
func (person *GamePerson) IsAnyAmmo() bool{
	if person.Ammo>0{
		return true
	}
	return false
}

func (person *GamePerson) IsAnyPower() bool{
	if person.Power>0{
		return true
	}
	return false
}

func (person *GamePerson) Shoot() bool{
	if person.On && person.IsAnyAmmo(){
		person.Ammo--
		return true
	}
	return false
}

func (person *GamePerson) RideBike() bool{
	if person.On && person.IsAnyPower(){
		person.Power--
		return true
	}
	return false
}

func main() {
	testStruct := &GamePerson{true,2,3}
	if testStruct.On == true{		
		for testStruct.Ammo!=0 || testStruct.Power!=0{
			fmt.Println("How many ammo do we have:", testStruct.Ammo, ". How many power do we have:", testStruct.Power)
			if testStruct.Shoot(){
				fmt.Println("Shoot")
			}else{
				fmt.Println("Ammo is ended")
			}
			if testStruct.RideBike(){
				fmt.Println("Ride a bike")
			}else{
				fmt.Println("Power is ended")
			}
			continue 
		}
		fmt.Println("Ammo and power are ended :(")  
	}else{
	fmt.Println("We are off :(")
	}     	
}
