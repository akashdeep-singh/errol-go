package models

type WhoAmIRequest struct {
}

func NewWhoAmIRequest() WhoAmIRequest {
	return WhoAmIRequest{}
}

func (_ WhoAmIRequest) RequestType() RequestType {
	return WhoAmIRequestType
}

func (_ WhoAmIRequest) Body() interface{} {
	return nil
}
