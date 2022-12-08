package usecase

import (
	"context"

	"github.com/AbraaoAbe/FullCycleClass/internal/domain/entity"
	"github.com/AbraaoAbe/FullCycleClass/internal/domain/repository"
	"github.com/AbraaoAbe/FullCycleClass/pkg/uow"
)

type AddMyTeamInput struct {
	ID    string
	Name  string
	Score int
}

type AddMyTeamUseCase struct {
	Uow uow.UowInterface
}

func (a *AddMyTeamUseCase) Execute(ctx context.Context, input AddMyTeamInput) error {
	//get repository from unit of work
	myTeamRepository := a.getMyTeamRepository(ctx)
	//create new entity with input data
	myTeam := entity.NewMyTeam(input.ID, input.Name)
	//save entity in database using repository
	err := myTeamRepository.Create(ctx, myTeam)
	if err != nil {
		return err
	}
	return a.Uow.CommitOrRollback()
}

func (a *AddMyTeamUseCase) getMyTeamRepository(ctx context.Context) repository.MyTeamRepositoryInterface {
	myTeamRepository, err := a.Uow.GetRepository(ctx, "MyTeamRepository")
	if err != nil {
		panic(err)
	}
	return myTeamRepository.(repository.MyTeamRepositoryInterface)
}
