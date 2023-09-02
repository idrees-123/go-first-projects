package main

import "fmt"



func main (){

	var val1, val2 uint;

	
	val1 = getUserInput();

	val2 = getUserInput();



	opertor := getOpertor();
	
	if(opertor == "+"){
		fmt.Println(val2+val1)
	} else if(opertor == "-"){
		fmt.Println(val2-val1)
	} else if(opertor == "*"){
		fmt.Println(val2*val1)
	} else if(opertor == "/"){
		fmt.Println(val2/val1)
	}

}

func plus ( val1, val2 uint)(uint){
	return val1 + val2;
}

func getUserInput () (uint){
	var val uint
 	fmt.Println("Please Enter Value :");
	_, err :=fmt.Scanln(&val);
	if  ! (val > 0 ) && err != nil {
		return 0;
	}

	return val;
 }


 func getOpertor() (string){
	var op string;

	fmt.Println("Please Enter Opertor");
	_, err := fmt.Scanln(&op);

	if err != nil {
		return op;
	}
	return op;
 }
