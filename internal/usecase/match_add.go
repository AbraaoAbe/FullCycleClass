package usecase

import (
	"context"
	"time"

	"github.com/AbraaoAbe/FullCycleClass/internal/domain/entity"
	"github.com/AbraaoAbe/FullCycleClass/internal/domain/repository"
	"github.com/AbraaoAbe/FullCycleClass/pkg/uow"
)

type MatchInput struct {
	ID      string
	Date    time.Time
	TeamAID string
	TeamBID string
}

type MatchUseCase struct {
	Uow uow.UowInterface
}

func (m *MatchUseCase) Execute(ctx context.Context, input MatchInput) error {
	return m.Uow.Do(ctx, func(uow *uow.Uow) error {
		//get match repository from unit of work
		matchRepository := m.getMatchRepository(ctx)

		//get team repository from unit of work
		teamRepository := m.getTeamRepository(ctx)

		teamA, err := teamRepository.FindByID(ctx, input.TeamAID)
		if err != nil {
			return err
		}
		teamB, err := teamRepository.FindByID(ctx, input.TeamBID)
		if err != nil {
			return err
		}

		//create a new match with input data and save it in database
		match := entity.NewMatch(input.ID, teamA, teamB, input.Date)
		err = matchRepository.Create(ctx, match)
		if err != nil {
			return err
		}
		return nil
	})
}

func (m *MatchUseCase) getMatchRepository(ctx context.Context) repository.MatchRepositoryInterface {
	matchRepository, err := m.Uow.GetRepository(ctx, "MatchRepository")
	if err != nil {
		panic(err)
	}
	return matchRepository.(repository.MatchRepositoryInterface)
}

func (m *MatchUseCase) getTeamRepository(ctx context.Context) repository.TeamRepositoryInterface {
	teamRepository, err := m.Uow.GetRepository(ctx, "TeamRepository")
	if err != nil {
		panic(err)
	}
	return teamRepository.(repository.TeamRepositoryInterface)
}
