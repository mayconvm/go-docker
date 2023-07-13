package person

import "encoding/json"

type StudentEntity struct {
	ID            string `json:"id,omitempty" example:"28a"`
	Name          string `json:"name,omitempty" query:"name"`
	Phone         string `json:"phone,omitempty" query:"phone"`
	Birthday      string `json:"birthday,omitempty" query:"birthday"`
	FatherID      string `json:"father_id,omitempty" query:"father_id"`
	MotherID      string `json:"mother_id,omitempty" query:"mother_id"`
	CreatedAt     string `json:"created_at,omitempty" query:"created_at" example:"2023-02-21"`
	LastUpdatedAt string `json:"last_updated_at,omitempty" query:"last_updated_at" example:"2023-02-21"`
}

func (entity *StudentEntity) ToJson() []byte {
	result, _ := json.Marshal(entity)

	return result
}

func (entity *StudentEntity) ToMap() map[string]interface{} {
	var result map[string]interface{}
	dataJson, _ := json.Marshal(entity)

	json.Unmarshal(dataJson, &result)

	return result
}

func NewStudentEntity(data interface{}) *StudentEntity {
	result := new(StudentEntity)
	dataJson, _ := json.Marshal(data)

	json.Unmarshal(dataJson, result)

	return result
}
