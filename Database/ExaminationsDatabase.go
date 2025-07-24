package database
import (
	"HospitalQOps/errorspacket"
	"HospitalQOps/model"
	"fmt"
	"strings"
	"time"
	"database/sql"
)
func GetExaminations() (string,error){
	db,err:=DatabaseConnection(connectionString)
	if err!=nil{
		fmt.Println(errorspacket.DatabaseConnectionError)
		return "",err
	}
	defer db.Close()
	query:=`SELECT * FROM Examinations`
	rows,err:=db.Query(query)
	if err!=nil{
		fmt.Println(errorspacket.QueryError)
		return "",err
	}
	defer rows.Close()
	var ListExaminations strings.Builder
	for rows.Next() {
		var examID, appID int
		var reexaminationDate sql.NullTime
		var examinationDate time.Time
		var info, diagnos, resultExam, nurseAssist, levelS, eNotes string
		var priceExam float64
		if err := rows.Scan(&examID, &appID, &examinationDate, &info, &nurseAssist, &diagnos,&resultExam,&reexaminationDate,&priceExam,&levelS,&eNotes); err != nil {
			return "", errorspacket.QueryError
		}
		ListExaminations.WriteString(fmt.Sprintf("Examination and Appointment ID's: %d %d\n", examID, appID))
		ListExaminations.WriteString(fmt.Sprintf("Examination Date: [%s]\n", examinationDate.Format("2006-01-02 15:04:05")))
		if reexaminationDate.Valid {
			ListExaminations.WriteString(fmt.Sprintf("Reexamination Date: [%s]\n", reexaminationDate.Time.Format("2006-01-02 15:04:05")))
		} else {
			ListExaminations.WriteString("Reexamination Date: N/A\n")
		}
		ListExaminations.WriteString(fmt.Sprintf("Information, Diagnosis, Result: %s %s %s\n",info,diagnos,resultExam))
		ListExaminations.WriteString(fmt.Sprintf("Assistant Nurses: %s\n",nurseAssist))
		ListExaminations.WriteString(fmt.Sprintf("Level and total price: %s %.2f\n",levelS, priceExam))
		ListExaminations.WriteString(fmt.Sprintf("Notes: %s\n\n",eNotes))
	}

	if err = rows.Err(); err != nil {
		return "", err
	}

	return ListExaminations.String(), nil
}
func InsertJunctionNurseDoctors(nurseID, doctorID int, notes string) error {
	db, err := DatabaseConnection(connectionString)
	if err != nil {
		fmt.Println(errorspacket.DatabaseConnectionError)
		return err
	}
	defer db.Close()

	query := `INSERT INTO nurses_doctors (nurses_nurse_id, doctors_doctor_id, notes) VALUES ($1, $2, $3)`
	_, err = db.Exec(query, nurseID, doctorID, notes)
	if err != nil {
		fmt.Println(errorspacket.ExecutionError)
		return errorspacket.ExecutionError
	}

	return nil
}

func ExaminationProcedure(appID int, examDate time.Time, info, nurses, diagnosis, result string, 
    reexamDate sql.NullTime, price float64, level, notes string) error {
    
	db, err := DatabaseConnection(connectionString)
	if err != nil {
		fmt.Println(errorspacket.DatabaseConnectionError)
		return err
	}
	defer db.Close()

	query := `INSERT INTO Examinations (
		appointment_id, examination_date, information, assistant_nurses, 
		diagnosis, result, reexamination_date, price, level, examination_notes
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err = db.Exec(query, appID, examDate, info, nurses, diagnosis, result, reexamDate, price, level, notes)
	if err != nil {
		fmt.Println(errorspacket.ExecutionError)
		return errorspacket.ExecutionError
	}

	return nil
}

func OrderExamResult() (string,error){
	db,err:=DatabaseConnection(connectionString)
	if err!=nil{
		fmt.Println(errorspacket.DatabaseConnectionError)
		return "",errorspacket.DatabaseConnectionError
	}
	defer db.Close()
	query:=`SELECT examination_id,appointment_id,examination_date,diagnosis,result,level
	FROM Examinations
	ORDER by result`
	rows,err:=db.Query(query)
	if err!=nil{
		fmt.Println(errorspacket.QueryError)
		return "",err
	}
	defer rows.Close()
	var ListExaminationsR strings.Builder
	for rows.Next() {
		var examID, appID int
		var examinationDate time.Time
		var diagnos, resultExam, levelS string
		if err := rows.Scan(&examID, &appID, &examinationDate, &diagnos,&resultExam,&levelS); err != nil {
			return "", errorspacket.QueryError
		}
		ListExaminationsR.WriteString(fmt.Sprintf("Examination and Appointment ID's: %d %d\n", examID, appID))
		ListExaminationsR.WriteString(fmt.Sprintf("Examination Date: [%s]\n", examinationDate.Format("2006-01-02 15:04:05")))
		ListExaminationsR.WriteString(fmt.Sprintf("Diagnosis, Result: %s %s\n",diagnos,resultExam))
		ListExaminationsR.WriteString(fmt.Sprintf("Level: %s \n\n",levelS))
	}

	if err = rows.Err(); err != nil {
		return "", err
	}

	return ListExaminationsR.String(), nil
}
func OrderExamLevel() (string,error){
	db,err:=DatabaseConnection(connectionString)
	if err!=nil{
		fmt.Println(errorspacket.DatabaseConnectionError)
		return "",errorspacket.DatabaseConnectionError
	}
	defer db.Close()
	query:=`SELECT examination_id,appointment_id,examination_date,diagnosis,result,level
	FROM Examinations
	ORDER by level`
	rows,err:=db.Query(query)
	if err!=nil{
		fmt.Println(errorspacket.QueryError)
		return "",err
	}
	defer rows.Close()
	var ListExaminationsL strings.Builder
	for rows.Next() {
		var examID, appID int
		var examinationDate time.Time
		var diagnos, resultExam, levelS string
		if err := rows.Scan(&examID, &appID, &examinationDate, &diagnos,&resultExam,&levelS); err != nil {
			return "", errorspacket.QueryError
		}
		ListExaminationsL.WriteString(fmt.Sprintf("Examination and Appointment ID's: %d %d\n", examID, appID))
		ListExaminationsL.WriteString(fmt.Sprintf("Examination Date: [%s]\n", examinationDate.Format("2006-01-02 15:04:05")))
		ListExaminationsL.WriteString(fmt.Sprintf("Diagnosis, Result: %s %s\n",diagnos,resultExam))
		ListExaminationsL.WriteString(fmt.Sprintf("Level: %s \n\n",levelS))
	}

	if err = rows.Err(); err != nil {
		return "", err
	}

	return ListExaminationsL.String(), nil
}
func GetExaminationsPrescriptionsPatient(pat model.Patient) (string,error){
	db,err:=DatabaseConnection(connectionString)
	if err!=nil{
		fmt.Println(errorspacket.DatabaseConnectionError)
		return "",errorspacket.DatabaseConnectionError
	}
	defer db.Close()
	query:=`SELECT Examinations.examination_id,Examinations.appointment_id,Examinations.examination_date,Examinations.diagnosis,Examinations.result,Examinations.level,Examinations.price,
				Prescriptions.number_medications,Prescriptions.notes,Prescriptions.prescription_date,Prescriptions.total_price
	FROM Examinations
	INNER JOIN Prescriptions on Examinations.examination_id=Prescriptions.examination_id
	INNER JOIN Appointments on Examinations.appointment_id=Appointments.appointment_id
	INNER JOIN Patients on Appointments.patient_id=Patients.patient_id
	WHERE Patients.patient_id=$1`
	rows,err:=db.Query(query,pat.ID)
	if err!=nil{
		fmt.Println(errorspacket.QueryError)
		return "",err
	}
	defer rows.Close()
	var ListExaminationsPp strings.Builder
	for rows.Next() {
		var examID, appID int
		var examinationDate time.Time
		var diagnos, resultExam, levelS string
		var noMed int
		var notesPres string
		var prescriptionDate time.Time
		var totalPricePres,totalPriceExam float64
		if err := rows.Scan(&examID, &appID, &examinationDate, &diagnos,&resultExam,&levelS,&totalPriceExam,&noMed,&notesPres,&prescriptionDate,&totalPricePres); err != nil {
			return "", errorspacket.QueryError
		}
		ListExaminationsPp.WriteString(fmt.Sprintf("Examination and Appointment ID's: %d %d\n", examID, appID))
		ListExaminationsPp.WriteString(fmt.Sprintf("Examination Date: [%s]\n", examinationDate.Format("2006-01-02 15:04:05")))
		ListExaminationsPp.WriteString(fmt.Sprintf("Diagnosis, Result: %s %s\n",diagnos,resultExam))
		ListExaminationsPp.WriteString(fmt.Sprintf("Examination Level and Price: %s %.2f\n",levelS,totalPriceExam))
		ListExaminationsPp.WriteString(fmt.Sprintf("Prescription: Number of Medications and Date: %s [%s]\n",noMed,prescriptionDate.Format("2006-01-02 15:04:05")))
		ListExaminationsPp.WriteString(fmt.Sprintf("Prescription notes: %s\n",notesPres))
		ListExaminationsPp.WriteString(fmt.Sprintf("Prescription price: %.2f\n\n",totalPricePres))
	}

	if err = rows.Err(); err != nil {
		return "", err
	}

	return ListExaminationsPp.String(), nil
}