package usercase

import (
	"github.com/<TEMPLATE>/config"
	InUserCase "github.com/<TEMPLATE>/internal/usercase/person"
	"github.com/<TEMPLATE>/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserCase struct {
	Cfg   *config.Config
	Mongo *mongo.Client
	Log   *logger.Logger
}

func New(uc UserCase) UserCase {
	return uc
}

func (uc *UserCase) Students() *InUserCase.StudentUserCase {
	return &InUserCase.StudentUserCase{
		Mongo: uc.Mongo,
		Log:   uc.Log,
	}
}
