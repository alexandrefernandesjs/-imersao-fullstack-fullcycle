package usecase

import (
	"errors"

	"github.com/sSchmidtT/imersao-fullstack-fullcycle/codepix/domain/model"
)

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (p *PixUseCase) RegisterKey(key, kind, accountId string) (*model.PixKey, error) {
	account, err := p.PixKeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixkey, err := model.NewPixKey(account, kind, key)
	if err != nil {
		return nil, err
	}

	p.PixKeyRepository.RegisterKey(pixkey)
	if pixkey.ID == "" {
		return nil, errors.New("unable to create new key at the moment")
	}

	return pixkey, nil

}

func (p *PixUseCase) FindKey(key, kind string) (*model.PixKey, error) {
	return p.PixKeyRepository.FindKeyByKind(key, kind)
}
