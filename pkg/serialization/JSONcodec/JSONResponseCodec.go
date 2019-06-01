package JSONcodec

import (
	"encoding/json"
	"github.com/akashdeep-singh/errol-go/pkg/models"
	"github.com/akashdeep-singh/errol-go/pkg/serialization"
)

type JSONResponseCodec struct {
	serialization.ResponseCodec
}

func (r *JSONResponseCodec) EncodeResponse(response models.Response) []byte {
	responseType := response.ResponseType()
	body := response.Body()

	serializedResponse, _ := json.Marshal(models.ResponseStruct{responseType, body})
	return serializedResponse
}

func (r *JSONResponseCodec) DecodeResponse(encoded []byte) (models.Response, error) {
	var deseralizedResponse models.ResponseStruct
	_ = json.Unmarshal(encoded, &deseralizedResponse)
	responseType := deseralizedResponse.Type
	responseBody := deseralizedResponse.Body

	var response models.Response
	switch responseType {
	case models.WhoAmIResponseType:
		response = models.NewWhoAmIResponse(uint64(responseBody.(float64)))
	case models.GetConnectedClientsResponseType:
		var clients []uint64
		for _, element := range responseBody.([]interface{}) {
			clients = append(clients, uint64(element.(float64)))
		}
		response = models.NewGetConnectedClientsResponse(clients)
	case models.RelayResponseType:
		response = models.NewRelayResponse(responseBody.(string))
	default:
		err := models.ResponseDecodingError
		return nil, err
	}

	return response, nil
}
