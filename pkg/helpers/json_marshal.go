package helpers

import "encoding/json"

// SafeJsonMarshal safely marshals an interface{} to a []byte using json.Marshal.
// If there is an error during the marshal, it will re-marshal the original data
// and return it.
// It returns a []byte and an error. If there is an error during the marshal,
// the error will be returned. Otherwise, nil will be returned.
func SafeJsonMarshal(data interface{}) ([]byte, error) {
	resp, err := json.Marshal(data)
	if err != nil {
		return resp, err
	}
	var tempData interface{}
	err = json.Unmarshal(resp, &tempData)
	if err != nil {
		return json.Marshal(data)
	}
	return resp, nil
}
