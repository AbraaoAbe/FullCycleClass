package usecase

import (
	"context"

	"github.com/AbraaoAbe/FullCycleClass/internal/domain/entity"
	"github.com/AbraaoAbe/FullCycleClass/internal/domain/repository"
	"github.com/AbraaoAbe/FullCycleClass/pkg/uow"
)

type AddPlayerInput struct {
	ID           string
	Name         string
	InitialPrice float64
}

type AddPLayerUseCase struct {
	Uow uow.UowInterface
}

func (a *AddPLayerUseCase) Execute(ctx context.Context, input AddPlayerInput) error {
	playerRepository := a.getPlayerRepository(ctx)
	player := entity.NewPlayer(input.ID, input.Name, input.InitialPrice)
	err := playerRepository.Create(ctx, player)
	if err != nil {
		return err
	}
	return a.Uow.CommitOrRollback()
}

func (a *AddPLayerUseCase) getPlayerRepository(ctx context.Context) repository.PlayerRepositoryInterface {
	playerRepository, err := a.Uow.GetRepository(ctx, "PlayerRepository")
	if err != nil {
		panic(err)
	}
	return playerRepository.(repository.PlayerRepositoryInterface)
}
