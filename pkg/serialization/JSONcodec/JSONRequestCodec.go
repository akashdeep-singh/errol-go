package JSONcodec

import (
	"encoding/json"
	"github.com/akashdeep-singh/errol/pkg/models"
	"github.com/akashdeep-singh/errol/pkg/serialization"
)

type JSONRequestCodec struct {
	serialization.RequestCodec
}

func (r *JSONRequestCodec) EncodeRequest(request models.Request) []byte {
	requestType := request.RequestType()
	body := request.Body()
	var serializedRequest []byte
	if body != nil {
		serializedRequest, _ = json.Marshal(models.RequestStruct{requestType, body.(models.RelayRequest)})
	} else {
		serializedRequest, _ = json.Marshal(models.RequestStruct{requestType, models.RelayRequest{}})
	}
	return serializedRequest
}

func (r *JSONRequestCodec) DecodeRequest(encoded []byte) (models.Request, error) {
	var deserializedRequest models.RequestStruct
	_ = json.Unmarshal(encoded, &deserializedRequest)
	requestType := deserializedRequest.Type
	requestBody := deserializedRequest.Body

	var request models.Request

	switch requestType {
	case models.WhoAmIRequestType:
		request = models.NewWhoAmIRequest()
	case models.GetConnectedClientsRequestType:
		request = models.NewGetConnectedClientsRequest()
	case models.RelayRequestType:
		request = models.NewRelayRequest(requestBody.Receivers, requestBody.Message)
	default:
		err := models.RequestDecodingError
		return nil, err
	}

	return request, nil
}
