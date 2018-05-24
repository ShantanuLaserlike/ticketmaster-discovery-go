package main

import (
	"log"
	"strconv"
	"time"
	"./code"
	//"github.com/patjackson52/ticketmaster-discovery-go/disco"
)

// We can fetch a maximum of 12000 events per minute 
// We can call this function 83 times a day
// THERE IS A CATCH the API can access only the first 1000 elements :(
func generate(s_id string, n int){
	for i := 1;  i <= n; i++ {								
		p :=strconv.Itoa(i)
		code.Maincall(s_id, p, "199")					//s_id = StateCode , p = page number
		log.Println(p)
		time.Sleep(1 * time.Second)
	}
}
func fetch_state_data(id string){
	val := code.Maincall(id, "0", "49")
	val = val - 49
	if val > 50{
		code.Maincall(id, "1","49")
		val = val - 49
		time.Sleep(1 * time.Second)						// 1 second rest between each request
	}
	if val > 100{
		code.Maincall(id, "1","99")
		val = val - 99
		time.Sleep(1 * time.Second)						// 1 second rest between each request
	}
	val = val/200
	if (val > 1 && val > 4){
		generate(id, 4)
	}else if (val > 1 && val < 4){
		generate(id, val)
	}
}
func main(){
	states := [50]string{"AL","AK","AZ","AR","CA","CO","CT","DE","FL","GA","HI","ID","IL","IN","IA","KS","KY","LA","ME","MD","MA","MI","MN","MS","MO","MT","NE","NV","NH","NJ","NM","NY","NC","ND","OH","OK","OR","PA","RI","SC","SD","TN","TX","UT","VT","VA","WA","WV","WI","WY"}
	for i := 0; i < 50; i++ {
		fetch_state_data(states[i])	
	}
}
