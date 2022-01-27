package main

import(
	"fmt"
)

var Character_Names[4] string
var Character_Positions[4] int
var BoatLeftSide bool


func main(){
Init()
	pws()
	PutIn("Mann")
	pws()
	CrossRiver()
	pws()
}

func pws(){
	fmt.Println(MakeWorldState())
}

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

func ReturnWorldState() string{
	Init()
	return MakeWorldState()
}

func MakeWorldState() string{
	rightSide := MakeRightSide()
	boat := MakeBoat()
	leftSide := MakeLeftSide()
	res := leftSide + boat + rightSide
	return res
}

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

func PutIn(item string) string{
	for i := 0; i < len(Character_Names); i++{
		if(Character_Names[i] == item){
			Character_Positions[i] = 1
		}
	}
	return MakeWorldState()
}

func MakeBoat() string{
	var endResult[4]string
	var res string
	for i := 0; i < len(Character_Names); i++{
		if(Character_Positions[i] == 1){
			endResult[i] = Character_Names[i]
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

func CrossRiver() string{
	BoatLeftSide = !BoatLeftSide
	
	return MakeWorldState()
}
