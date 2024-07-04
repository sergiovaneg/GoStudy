package main

import "math"

const StartHP = 50
const StartMP = 500

const BossStartHP = 58

const BossDMG = 9
const ReducedBossDMG = max(1, BossDMG-7)

const HardMode = true

type GameStatus struct {
	hp              int
	mp              int
	bossHp          int
	shieldCounter   int
	poisonCounter   int
	rechargeCounter int
}

func (status *GameStatus) updateTimers() {
	if status.shieldCounter > 0 {
		status.shieldCounter--
	}
	if status.poisonCounter > 0 {
		status.bossHp -= 3
		status.poisonCounter--
	}
	if status.rechargeCounter > 0 {
		status.mp += 101
		status.rechargeCounter--
	}
}

func wizardTurn(
	manaSpent int,
	currentMinimum *int,
	status GameStatus) {
	// Hard Mode
	if HardMode {
		status.hp--
		if status.hp <= 0 {
			return
		}
	}

	// Run timers
	status.updateTimers()

	// Check death by poison
	if status.bossHp <= 0 {
		*currentMinimum = manaSpent
		return
	}

	// Try Magic Missile
	if status.mp >= 53 && manaSpent+53 < *currentMinimum {
		bossTurn(manaSpent+53, currentMinimum, GameStatus{
			hp:              status.hp,
			mp:              status.mp - 53,
			bossHp:          status.bossHp - 4,
			shieldCounter:   status.shieldCounter,
			poisonCounter:   status.poisonCounter,
			rechargeCounter: status.rechargeCounter,
		})
	}

	// Try Drain
	if status.mp >= 73 && manaSpent+73 < *currentMinimum {
		bossTurn(manaSpent+73, currentMinimum, GameStatus{
			hp:              status.hp + 2,
			mp:              status.mp - 73,
			bossHp:          status.bossHp - 2,
			shieldCounter:   status.shieldCounter,
			poisonCounter:   status.poisonCounter,
			rechargeCounter: status.rechargeCounter,
		})
	}

	// Try Shield
	if status.shieldCounter == 0 &&
		status.mp >= 113 &&
		manaSpent+113 < *currentMinimum {
		bossTurn(manaSpent+113, currentMinimum, GameStatus{
			hp:              status.hp,
			mp:              status.mp - 113,
			bossHp:          status.bossHp,
			shieldCounter:   6,
			poisonCounter:   status.poisonCounter,
			rechargeCounter: status.rechargeCounter,
		})
	}

	// Try Poison
	if status.poisonCounter == 0 &&
		status.mp >= 173 &&
		manaSpent+173 < *currentMinimum {
		bossTurn(manaSpent+173, currentMinimum, GameStatus{
			hp:              status.hp,
			mp:              status.mp - 173,
			bossHp:          status.bossHp,
			shieldCounter:   status.shieldCounter,
			poisonCounter:   6,
			rechargeCounter: status.rechargeCounter,
		})
	}

	// Try Recharge
	if status.rechargeCounter == 0 &&
		status.mp >= 229 &&
		manaSpent+229 < *currentMinimum {
		bossTurn(manaSpent+229, currentMinimum, GameStatus{
			hp:              status.hp,
			mp:              status.mp - 229,
			bossHp:          status.bossHp,
			shieldCounter:   status.shieldCounter,
			poisonCounter:   status.poisonCounter,
			rechargeCounter: 5,
		})
	}
}

func bossTurn(manaSpent int,
	currentMinimum *int,
	status GameStatus) {
	// Run timers
	status.updateTimers()

	// Check death by poison
	if status.bossHp <= 0 {
		*currentMinimum = manaSpent
		return
	}

	if status.shieldCounter > 0 {
		status.hp -= ReducedBossDMG
	} else {
		status.hp -= BossDMG
	}

	if status.hp <= 0 {
		return
	}

	wizardTurn(manaSpent, currentMinimum, status)
}

func main() {
	minimum := math.MaxInt
	wizardTurn(0, &minimum, GameStatus{
		hp:     StartHP,
		mp:     StartMP,
		bossHp: BossStartHP,
	})
	println(minimum)
}
