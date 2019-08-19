package utils

type ResponseData struct {
	message    string
	data       interface{}
	statusCode string
	total      int
}

func Response(data interface{}, message string, status string, total int) ResponseData {
	return ResponseData{
		message:    message,
		data:       data,
		statusCode: status,
		total:      total,
	}
}
