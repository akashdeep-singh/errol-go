package models

type WhoAmIResponse struct {
	Id     uint64
	Target uint64
}

func NewWhoAmIResponse(data uint64) WhoAmIResponse {
	return WhoAmIResponse{Id: data}
}

func (_ WhoAmIResponse) ResponseType() ResponseType {
	return WhoAmIResponseType
}

func (whoAmIRequest WhoAmIResponse) Body() interface{} {
	return whoAmIRequest.Id
}

func (whoAmIRequest WhoAmIResponse) Targets() []uint64 {
	var targets []uint64
	return append(targets, whoAmIRequest.Target)
}
