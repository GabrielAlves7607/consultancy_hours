package constControllers

// ScheduleUnavailable message + Error ScheduleUnavailable message
const (
	ErrorScheduleUnavailable    = "horario indisponiveis"
	MsgErrorScheduleUnavailable = "This time is already booked"
)

// TimeId + Erro message
const (
	MsgErrorTimeId = "id_horario invalid"
	ErrorTimeId    = "The id_horario sent is not a valid slot"
)

// Error Database
const (
	MsgErrorDatabase = "Failed to schedule in database"
)

// HTTP Headers
const (
	HeaderContentType     = "Content-Type"
	HeaderContentTypeJSON = "application/json"
)

// Keys JSON
const (
	JSONKeyError      = "erro"
	JSONKeyMessage    = "message"
	JSONKeyInsertedID = "id_inserido"
)

// Message of users (Mistakes and Successes)
const (
	MsgErrorFailedDBQuery = "Failed to query database"
	MsgErrorInvalidJSON   = "JSON invalid"
	MsgErrorMissingFields = "id_horario and nome_cliente are mandatory"
	MsgScheduleSuccess    = "Schedule completed successfully!" // Corrigi o "Shedule"
)

// Log Formats
const (
	LogRequestReceived = "Request for schedule received %s of the customer %s\n"
	LogInsertSuccess   = "Inserted into DB with ID: %v\n"
)
