package network
import (
	"fmt"
	"net"
	"encoding/json"
	"HospitalQOps/errorspacket"
	//"HospitalQOps/logger"
	"HospitalQOps/model"
)
func HospitalServerConnect(ipAddress string) net.Conn{
	con,err:=net.Dial("tcp",ipAddress)
	if err!=nil{
		fmt.Println(errorspacket.DialNetworkIpAddressError)
		return nil
	}
	return con
}
func SendJSON(con net.Conn, msg model.Message) error{
	sending,_:=json.Marshal(msg)
	_,err:=con.Write(sending)
	//Future logger
	fmt.Printf("Doctor %s sent message to %s with content: %s\n",msg.Sender,msg.Receiver,msg.Content)
	if err!=nil{
		return errorspacket.SendingNetworkError
	}
	return nil
}
func ReadJSON(con net.Conn) (model.Message, error){
	var m model.Message
	buffer:=make([]byte, 1024)
	reading,err:=con.Read(buffer)
	if err!=nil{
		return model.Message{},errorspacket.ReadingNetworkError
	}
	err=json.Unmarshal(buffer[:reading],&m)
	//Future logger
	fmt.Printf("Doctor received message: %+v\n",m)
	return m,err
}