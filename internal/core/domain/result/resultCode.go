package result

const (
	CodeOk      = 200
	CodeCreated = 201
	CodeUpdated = 202
	CodeDeleted = 203
	//Errors Code
	CodeInternalError = 500
)

func CodeText(code int) string {
	switch code {
	case CodeOk:
		return "Ok"
	case CodeCreated:
		return "Created"
	case CodeUpdated:
		return "Updated"
	case CodeDeleted:
		return "Deleted"
	case CodeInternalError:
		return "Internal Error"
	default:
		return ""
	}
}
