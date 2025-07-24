package database
import (
	"HospitalQOps/errorspacket"
	"fmt"
	"strings"
	"time"
	"database/sql"
)
func DoctorValidatesAppointment(doctorID int, appointmentID int, response string) error {
	db, err := DatabaseConnection(connectionString)
	if err != nil {
		fmt.Println(errorspacket.DatabaseConnectionError)
		return err
	}
	defer db.Close()

	var assignedDoctorID int
	var appointmentStatus string

	err = db.QueryRow(`SELECT doctor_id, status FROM Appointments WHERE appointment_id = $1`, appointmentID).Scan(&assignedDoctorID, &appointmentStatus)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("appointment ID %d does not exist", appointmentID)
		}
		return err
	}
	if assignedDoctorID != doctorID {
		return fmt.Errorf("doctor %d is not assigned to appointment %d", doctorID, appointmentID)
	}
	if response != "approved" && response != "denied" {
		return fmt.Errorf("invalid response '%s', must be 'approved' or 'denied'", response)
	}
	var newStatus string
	if response == "approved" {
		newStatus = "scheduled"
	} else {
		newStatus = "canceled"
	}
	_, err = db.Exec(`UPDATE Appointments SET status = $1 WHERE appointment_id = $2`, newStatus, appointmentID)
	if err != nil {
		return err
	}

	return nil
}

func GetAppointments() (string,error){
	db, err := DatabaseConnection(connectionString)
	if err != nil {
		fmt.Println(errorspacket.DatabaseConnectionError)
		return "",err
	}
	defer db.Close()

	query := `SELECT appointment_id,patient_id,doctor_id,room_id,schedule_date,status
	FROM Appointments`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(errorspacket.QueryError)
		return "", errorspacket.QueryError
	}
	defer rows.Close()

	var listAppointments strings.Builder

	for rows.Next() {
		var appID, patID, docID, roomID int
		var schDate time.Time
		var statApp string
		if err := rows.Scan(&appID, &patID, &docID, &roomID, &schDate, &statApp); err != nil {
			return "", errorspacket.QueryError
		}
		listAppointments.WriteString(fmt.Sprintf("Appointment, Patient, Doctor, Room ID's: %d %d %d %d\n", appID, patID, docID, roomID))
		listAppointments.WriteString(fmt.Sprintf("Appointment Date: [%s]\n", schDate.Format("2006-01-02 15:04:05")))
		listAppointments.WriteString(fmt.Sprintf("Appointment Status: %s\n\n", statApp))
	}

	if err = rows.Err(); err != nil {
		return "", err
	}

	return listAppointments.String(), nil
}
func GetAppointmentsByPatient(patID int) (string,error){
	db, err := DatabaseConnection(connectionString)
	if err != nil {
		fmt.Println(errorspacket.DatabaseConnectionError)
		return "",err
	}
	defer db.Close()

	query := `SELECT appointment_id,patient_id,doctor_id,room_id,schedule_date,status
	FROM Appointments
	WHERE patient_id=$1`
	rows, err := db.Query(query,patID)
	if err != nil {
		fmt.Println(errorspacket.QueryError)
		return "", errorspacket.QueryError
	}
	defer rows.Close()

	var listAppointmentsP strings.Builder

	for rows.Next() {
		var appID, patIDI, docID, roomID int
		var schDate time.Time
		var statApp string
		if err := rows.Scan(&appID, &patIDI, &docID, &roomID, &schDate, &statApp); err != nil {
			return "", errorspacket.QueryError
		}
		listAppointmentsP.WriteString(fmt.Sprintf("Appointment, Patient, Doctor, Room ID's: %d %d %d %d\n", appID, patIDI, docID, roomID))
		listAppointmentsP.WriteString(fmt.Sprintf("Appointment Date: [%s]\n", schDate.Format("2006-01-02 15:04:05")))
		listAppointmentsP.WriteString(fmt.Sprintf("Appointment Status: %s\n\n", statApp))
	}

	if err = rows.Err(); err != nil {
		return "", err
	}

	return listAppointmentsP.String(), nil
}
func InsertAppointment(patID int, docID int, roomID int, scheduleDate time.Time) error{
	db,err:=DatabaseConnection(connectionString)
	if err!=nil{
		fmt.Println(errorspacket.DatabaseConnectionError)
		return err
	}
	defer db.Close()
	instruction:=`INSERT INTO Appointments (patient_id, doctor_id, room_id, schedule_date, status)
	VALUES ($1, $2, $3, $4, 'waiting')`
	_,err=db.Exec(instruction,patID,docID, roomID, scheduleDate)
	if err!=nil{
		fmt.Println(errorspacket.ExecutionError)
		return errorspacket.ExecutionError
	}
	return nil
}