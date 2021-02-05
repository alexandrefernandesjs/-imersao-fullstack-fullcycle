package factory

import (
	"github.com/jinzhu/gorm"
	"github.com/sSchmidtT/imersao-fullstack-fullcycle/codepix/application/usecase"
	"github.com/sSchmidtT/imersao-fullstack-fullcycle/codepix/infrastructure/repository"
)

func TransactionUseCaseFactory(database *gorm.DB) usecase.TransactionUseCase {
	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	transactionUseCaseFactory := repository.TransactionRepositoryDb{Db: database}

	transactionUseCase := usecase.TransactionUseCase{
		TransactionRepository: &transactionUseCaseFactory,
		PixRepository:         pixRepository,
	}
	return transactionUseCase
}
