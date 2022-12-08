package usecase

import (
	"context"

	"github.com/AbraaoAbe/FullCycleClass/internal/domain/entity"
	"github.com/AbraaoAbe/FullCycleClass/internal/domain/repository"
	"github.com/AbraaoAbe/FullCycleClass/pkg/uow"
)

type ActionAddInput struct {
	MacthID  string
	TeamID   string
	PlayerID string
	Minute   int
	Action   string
}

type ActionAddUseCase struct {
	Uow         uow.UowInterface
	ActionTable entity.ActionTableInterface
}

func (a *ActionAddUseCase) Execute(ctx context.Context, input ActionAddInput) error {
	return a.Uow.Do(ctx, func(uow *uow.Uow) error {
		//get repositories from unit of work
		matchRepository := a.getMatchRepository(ctx)
		myTeamRepository := a.getMyTeamRepository(ctx)
		playerRepository := a.getPlayerRepository(ctx)
		//actionRepository := a.getActionRepository(ctx)

		//get entities from repositories
		match, err := matchRepository.FindByID(ctx, input.MacthID)
		if err != nil {
			return err
		}

		//get action score
		score, err := a.ActionTable.GetScore(input.Action)
		if err != nil {
			return err
		}
		//create new action entity
		action := entity.NewGameAction(input.PlayerID, input.Minute, input.Action, score)
		//add action to match
		match.Actions = append(match.Actions, *action)
		//save match in database
		err = matchRepository.SaveActions(ctx, match, float64(score))

		player, err := playerRepository.FindByID(ctx, input.PlayerID)
		if err != nil {
			return err
		}

		player.Price += float64(score)

		myTeam, err := myTeamRepository.FindByID(ctx, input.TeamID)
		if err != nil {
			return err
		}

		err = myTeamRepository.AddScore(ctx, myTeam, float64(score))
		if err != nil {
			return err
		}

		return nil
	})
}

func (a *ActionAddUseCase) getMatchRepository(ctx context.Context) repository.MatchRepositoryInterface {
	matchRepository, err := a.Uow.GetRepository(ctx, "MatchRepository")
	if err != nil {
		panic(err)
	}
	return matchRepository.(repository.MatchRepositoryInterface)
}

func (a *ActionAddUseCase) getMyTeamRepository(ctx context.Context) repository.MyTeamRepositoryInterface {
	myTeamRepository, err := a.Uow.GetRepository(ctx, "MyTeamRepository")
	if err != nil {
		panic(err)
	}
	return myTeamRepository.(repository.MyTeamRepositoryInterface)
}

func (a *ActionAddUseCase) getPlayerRepository(ctx context.Context) repository.PlayerRepositoryInterface {
	playerRepository, err := a.Uow.GetRepository(ctx, "PlayerRepository")
	if err != nil {
		panic(err)
	}
	return playerRepository.(repository.PlayerRepositoryInterface)
}
