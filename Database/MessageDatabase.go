package database
import (
	"fmt"
	"HospitalQOps/model"
	"HospitalQOps/errorspacket"
)
func InsertMessage(msg model.Message) error {
	db,err:=DatabaseConnection(connectionString)
	if err!=nil{
		fmt.Println(errorspacket.DatabaseConnectionError)
		return errorspacket.DatabaseConnectionError
	}
	defer db.Close()
	query := `INSERT INTO messages (sender_email, receiver_email, title, content) VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(query, msg.Sender, msg.Receiver, msg.Title, msg.Content)
	if err!=nil{
		fmt.Println(errorspacket.ExecutionError)
		return errorspacket.ExecutionError
	}
	return err
}

func GetDoctorInbox(doc model.Doctor) ([]model.Message, error) {
	db,err:=DatabaseConnection(connectionString)
	if err!=nil{
		fmt.Println(errorspacket.DatabaseConnectionError)
		return nil, errorspacket.DatabaseConnectionError
	}
	defer db.Close()
	query := `SELECT sender_email, receiver_email, title, content FROM messages WHERE receiver_email = $1 ORDER BY id DESC`
	rows, err := db.Query(query, doc.Email)
	if err != nil {
		fmt.Println(errorspacket.QueryError)
		return nil, errorspacket.QueryError
	}
	defer rows.Close()

	var messages []model.Message
	for rows.Next() {
		var msg model.Message
		if err := rows.Scan(&msg.Sender, &msg.Receiver, &msg.Title, &msg.Content); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}

func GetConversationEmail(doc model.Doctor, email2 string) ([]model.Message, error) {
	db,err:=DatabaseConnection(connectionString)
	if err!=nil{
		fmt.Println(errorspacket.DatabaseConnectionError)
		return nil, errorspacket.DatabaseConnectionError
	}
	defer db.Close()
	query := `
	SELECT sender_email, receiver_email, title, content 
	FROM messages 
	WHERE (sender_email = $1 AND receiver_email = $2) 
	   OR (sender_email = $2 AND receiver_email = $1)
	ORDER BY id`
	rows, err := db.Query(query, doc.Email, email2)
	if err != nil {
		fmt.Println(errorspacket.QueryError)
		return nil, errorspacket.QueryError
	}
	defer rows.Close()

	var messages []model.Message
	for rows.Next() {
		var msg model.Message
		if err := rows.Scan(&msg.Sender, &msg.Receiver, &msg.Title, &msg.Content); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}
func GetPatientInbox(pat model.Patient) ([]model.Message, error) {
	db,err:=DatabaseConnection(connectionString)
	if err!=nil{
		fmt.Println(errorspacket.DatabaseConnectionError)
		return nil, errorspacket.DatabaseConnectionError
	}
	defer db.Close()
	query := `SELECT sender_email, receiver_email, title, content FROM messages WHERE receiver_email = $1 ORDER BY id DESC`
	rows, err := db.Query(query, pat.Email)
	if err != nil {
		fmt.Println(errorspacket.QueryError)
		return nil, errorspacket.QueryError
	}
	defer rows.Close()

	var messages []model.Message
	for rows.Next() {
		var msg model.Message
		if err := rows.Scan(&msg.Sender, &msg.Receiver, &msg.Title, &msg.Content); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}

func GetConversationEmailPatient(pat model.Patient, email2 string) ([]model.Message, error) {
	db,err:=DatabaseConnection(connectionString)
	if err!=nil{
		fmt.Println(errorspacket.DatabaseConnectionError)
		return nil, errorspacket.DatabaseConnectionError
	}
	defer db.Close()
	query := `
	SELECT sender_email, receiver_email, title, content 
	FROM messages 
	WHERE (sender_email = $1 AND receiver_email = $2) 
	   OR (sender_email = $2 AND receiver_email = $1)
	ORDER BY id`
	rows, err := db.Query(query, pat.Email, email2)
	if err != nil {
		fmt.Println(errorspacket.QueryError)
		return nil, errorspacket.QueryError
	}
	defer rows.Close()

	var messages []model.Message
	for rows.Next() {
		var msg model.Message
		if err := rows.Scan(&msg.Sender, &msg.Receiver, &msg.Title, &msg.Content); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}