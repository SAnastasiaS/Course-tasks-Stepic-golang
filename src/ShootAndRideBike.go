/*Задание было

"В рамках этого урока мы постарались представить себе уже привычные нам переменные и функции, как объекты из реальной жизни. Чтобы закрепить результат мы предлагаем вам небольшую творческую задачу.

Вам необходимо реализовать структуру со свойствами-полями On, Ammo и Power, с типами bool, int, int соответственно. У этой структуры должны быть методы: Shoot и RideBike, которые не принимают аргументов, но возвращают значение bool.

Если значение On == false, то оба метода вернут false.

Делать Shoot можно только при наличии Ammo (тогда Ammo уменьшается на единицу, а метод возвращает true), если его нет, то метод вернет false. Метод RideBike работает также, но только зависит от свойства Power.

Чтобы проверить, что вы все сделали правильно, вы должны создать указатель на экземпляр этой структуры с именем testStruct в функции main, в дальнейшем программа проверит результат.

Закрывающая фигурная скобка в конце main() вам не видна, но она есть.

Пакет main объявлять не нужно!

Удачи!"

Для запуска у себя добавила некоторую самопроверку в main*/

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
