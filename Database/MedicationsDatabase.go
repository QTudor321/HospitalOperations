package database
import (
	"fmt"
	"strings"
	"HospitalQOps/errorspacket"
)
func GetMedications() (string,error){
	db,err:=DatabaseConnection(connectionString)
	if err!=nil{
		fmt.Println(errorspacket.DatabaseConnectionError)
		return "", errorspacket.DatabaseConnectionError
	}
	defer db.Close()
	query:=`
		SELECT *
		FROM Medications`
	rows,err:=db.Query(query)
	if err!=nil{
		fmt.Println(errorspacket.QueryError)
		return "",errorspacket.QueryError
	}
	defer rows.Close()
	var ListOfMedications strings.Builder
	for rows.Next(){
		var medicationID, depID, inventoryMed int
		var priceMed float64
		var nameMed, descrMed string
		if err:=rows.Scan(&medicationID,&nameMed,&descrMed,&priceMed,&depID,&inventoryMed); err!=nil{
			return "", errorspacket.QueryError
		}
		ListOfMedications.WriteString(fmt.Sprintf("Medication ID: %d\n", medicationID))
		ListOfMedications.WriteString(fmt.Sprintf("Name: %s\n", nameMed))
		ListOfMedications.WriteString(fmt.Sprintf("Description: %s\n", descrMed))
		ListOfMedications.WriteString(fmt.Sprintf("Price: %.2f\n", priceMed))
		ListOfMedications.WriteString(fmt.Sprintf("Department ID: %d\n", depID))
		ListOfMedications.WriteString(fmt.Sprintf("Inventory: %d\n\n", inventoryMed))
	}
	return ListOfMedications.String(),nil
}