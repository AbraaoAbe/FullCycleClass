package service

import (
	"errors"

	"github.com/AbraaoAbe/FullCycleClass/internal/domain/entity"
)

var errNotEnoughMoney = errors.New("not enough money")

func ChoosePlayers(myTeam *entity.MyTeam, players []entity.Player) error {
	totalCost := 0.0
	totalEarned := 0.0

	//check if I have enough money to buy all players
	for _, player := range players {
		//when a player is in my team and not in players list, I sold the player
		if playerInMyTeam(player, myTeam) && !playerInPlayersList(player, &players) {
			totalEarned += player.Price
		}
		//when a player is not in my team and is in players list, I bought the player
		if !playerInMyTeam(player, myTeam) && playerInPlayersList(player, &players) {
			totalCost += player.Price
		}
	}

	if totalCost > myTeam.Score+totalEarned {
		return errNotEnoughMoney
	}
	//update my team score
	myTeam.Score += totalEarned - totalCost
	//clear my team players
	myTeam.Players = []string{}

	//add new players to my team
	for _, player := range players {
		myTeam.Players = append(myTeam.Players, player.ID)
	}
	return nil
}

func playerInMyTeam(player entity.Player, myTeam *entity.MyTeam) bool {
	for _, p := range myTeam.Players {
		if p == player.ID {
			return true
		}
	}
	return false
}

func playerInPlayersList(player entity.Player, players *[]entity.Player) bool {
	for _, p := range *players {
		if p == player {
			return true
		}
	}
	return false
}
