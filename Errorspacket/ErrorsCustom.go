package errorspacket
import "errors"
var InvalidInputError=errors.New("Error: Invalid Input")
var DoctorCredentialsError=errors.New("Error: Wrong Doctor Credentials Input")
var PatientCredentialsError=errors.New("Error: Wrong Patient Credentials Input")
var DialNetworkIpAddressError=errors.New("Error: Dialing Server IP Address")
var SendingNetworkError=errors.New("Error: Sending TCP Packet")
var ReadingNetworkError=errors.New("Error: Reading TCP Packet")
var ListenerNetworkError=errors.New("Error: Listener Server Error")
var NetworkConnectionError=errors.New("Error: Connection to the Server Network")
var VaultError=errors.New("Error: Passcodes Vault Opening")
var DatabaseConnectionError=errors.New("ERROR: Database Connection Driver")
var QueryError=errors.New("Error: Querying Information")
var ExecutionError=errors.New("Error: Execution Database Statement")
