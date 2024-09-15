package httpclient

type HttpHeaderDTO struct {
	Key   string
	Value string
}

func ToHttpHeaderDto(key string, value string) HttpHeaderDTO {
	headerDTO := HttpHeaderDTO{}
	headerDTO.Key = key
	headerDTO.Value = value

	return headerDTO
}
