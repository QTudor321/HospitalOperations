package database
import (
	"fmt"
	"strings"
	"HospitalQOps/errorspacket"
)
var connectionString= "host=localhost port=5432 user=postgres password=Quantum132 dbname=hospital sslmode=disable"
func GetNursesAndDoctorsByDepartments() (string,error){
	db,err:=DatabaseConnection(connectionString)
	if err!=nil{
		fmt.Println(errorspacket.DatabaseConnectionError)
		return "", errorspacket.DatabaseConnectionError
	}
	defer db.Close()
	query:=`
		SELECT Nurses.first_name, Nurses.last_name, Doctors.first_name, Doctors.last_name, Departments.name
		FROM Nurses
		INNER JOIN Departments on Nurses.department_id = Departments.department_id
		INNER JOIN Doctors on Departments.department_id = Doctors.department_id`
	ListOfStaff,err:=db.Query(query)
	if err!=nil{
		fmt.Println(errorspacket.QueryError)
		return "",errorspacket.QueryError
	}
	defer ListOfStaff.Close()
	var ListOfStaffString strings.Builder
	for ListOfStaff.Next(){
		var nurseFirstN, nurseLastN, doctorFirstN, doctorLastN, departmentName string
		if err:=ListOfStaff.Scan(&nurseFirstN,&nurseLastN,&doctorFirstN,&doctorLastN,&departmentName); err!=nil{
			return "", errorspacket.QueryError
		}
		ListOfStaffString.WriteString(fmt.Sprintf("Department: %s\n", departmentName))
		ListOfStaffString.WriteString(fmt.Sprintf("  Nurse: %s %s\n", nurseFirstN, nurseLastN))
		ListOfStaffString.WriteString(fmt.Sprintf("  Doctor: %s %s\n\n", doctorFirstN, doctorLastN))
	}
	return ListOfStaffString.String(),nil
}
func GetDepartmentsID() (string,error){
	db,err:=DatabaseConnection(connectionString)
	if err!=nil{
		fmt.Println(errorspacket.DatabaseConnectionError)
		return "", errorspacket.DatabaseConnectionError
	}
	defer db.Close()
	query:=`SELECT department_id,name
	FROM Departments`
	rows,err:=db.Query(query)
	if err!=nil{
		fmt.Println(errorspacket.QueryError)
		return "",errorspacket.QueryError
	}
	defer rows.Close()
	var DepartmentList strings.Builder
	for rows.Next(){
		var departmentID int
		var depName string
		if err:=rows.Scan(&departmentID,&depName); err!=nil{
			return "", errorspacket.QueryError
		}
		DepartmentList.WriteString(fmt.Sprintf("Department ID: %s\n", departmentID))
		DepartmentList.WriteString(fmt.Sprintf("Name: %s\n\n", depName))
	}
	return DepartmentList.String(),nil
}