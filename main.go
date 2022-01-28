package main

import(
	"fmt"
)
//Setter opp character name og character posisjons array
var Character_Names[4] string
var Character_Positions[4] int

//Hvilken side er båten på?
var BoatLeftSide bool


func main(){
	Init()
	pws()
	PutIn("Rev")
	pws()
	CrossRiver()
	pws()
}

//Shorthand funksjon for å printe world state fordi det er langt å skrive
func pws(){
	fmt.Println(MakeWorldState())
}
//Initialiserer alle arrayene så karakternavn matcher med samme index i karakter posisjon
func Init(){
	BoatLeftSide = true
	Character_Names[0] = "Mann"
	Character_Names[1] = "Korn"
	Character_Names[2] = "Kylling"
	Character_Names[3] = "Rev"
	//Mann
	Character_Positions[0] = 0
	//Korn
	Character_Positions[1] = 0
	//Kylling
	Character_Positions[2] = 0
	//Rev
	Character_Positions[3] = 0
}
//For testing
func ReturnWorldState() string{
	Init()
	return MakeWorldState()
}

//Samler alle statene til en streng
func MakeWorldState() string{
	rightSide := MakeRightSide()
	boat := MakeBoat()
	leftSide := MakeLeftSide()
	res := leftSide + boat + rightSide
	return res
}
//Lager høyre side av elven ved å loope gjennom og sjekke hvilke som har posisjon 0 (Høyre) også legge de til i et array før den ferdige stringen returneres
func MakeRightSide() string{
	var endResult[4]string
	for i := 0; i < len(Character_Names); i++{
		if(Character_Positions[i] == 2){
			endResult[i] = Character_Names[i]
		}else{
			endResult[i] = " "
		}
	}
	res := fmt.Sprintf("%s-%s-%s-%s-", endResult[0], endResult[1], endResult[2], endResult[3])
	return res
}

//Samme som for makerightside men for venstre siden
func MakeLeftSide() string{
	var endResult[4]string
	for i := 0; i < len(Character_Names); i++{
		if(Character_Positions[i] == 0){
			endResult[i] = Character_Names[i]
		}else{
			endResult[i] = " "
		}
	}
	res := fmt.Sprintf("%s-%s-%s-%s-", endResult[0], endResult[1], endResult[2], endResult[3])
	return res
}

//Oppdaterer posisjonen til karkateren i posisjonsarrayet og returnerer statusen til verden, return er for testing
func PutIn(item string) string{
	for i := 0; i < len(Character_Names); i++{
		if(Character_Names[i] == item){
			Character_Positions[i] = 1
		}
	}
	return MakeWorldState()
}

//Fungerer på samme måte som makeright og makeleft men har en ekstra count variabel som fikser problemet med at karakterer som var 
//Etter index 0 i arrayet ikke kom med i båten
func MakeBoat() string{
	var endResult[4]string
	var res string
	count := 0
	for i := 0; i < len(Character_Names); i++{
		if(Character_Positions[i] == 1){
			endResult[count] = Character_Names[i]
			count++
		}else{
			endResult[i] = " "
		}
	}
	if(BoatLeftSide){
		res = fmt.Sprintf(" \\_%s_%s_/________", endResult[0], endResult[1])
	}else{
		res = fmt.Sprintf("________\\_%s_%s_/  ", endResult[0], endResult[1])
	}
	return res
}
//Flipper boatleftside og returnerer verdensstatusen
func CrossRiver() string{
	BoatLeftSide = !BoatLeftSide
	return MakeWorldState()
}
