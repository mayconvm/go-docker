package person

import (
	Entity "github.com/<TEMPLATE>/internal/entity/person"
	Rep "github.com/<TEMPLATE>/internal/repository"
	"github.com/<TEMPLATE>/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type StudentUserCase struct {
	Mongo *mongo.Client
	Log   *logger.Logger
}

func (duc *StudentUserCase) SaveDocument(entity *Entity.StudentEntity) *Entity.StudentEntity {
	repDoc := Rep.NewStudentRepository(duc.Mongo, duc.Log)

	return repDoc.SaveNewEntity(entity)
}

func (duc *StudentUserCase) SearchDocuments(fields Entity.StudentEntity) []*Entity.StudentEntity {
	repDoc := Rep.NewStudentRepository(duc.Mongo, duc.Log)

	return repDoc.Search(fields)
}
