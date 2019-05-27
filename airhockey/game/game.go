package game

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type Vec2d struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func scale(v Vec2d, s float64) Vec2d {
	return Vec2d{
		X: v.X * s,
		Y: v.Y * s,
	}
}

func add(v1, v2 Vec2d) Vec2d  {
	return Vec2d{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
}

type State struct {
	mallet1Pos, mallet2Pos, puckVec, puckPos Vec2d
}

const malletRadius = 0.08
const puckRadius = 0.06

var state *State

func Init()  {
	state = &State{
		mallet1Pos:Vec2d{X:0, Y:0.4},
		mallet2Pos:Vec2d{X:0, Y:-0.4},
		puckVec:Vec2d{X:0, Y:0},
		puckPos:Vec2d{X:0, Y:0},
	}
}

func fts(f float64) string  {
	return fmt.Sprintf("%f", f)
}

func GetGameState() gin.H  {
	return gin.H{
		"mallet1Pos.x": fts(state.mallet1Pos.X),
		"mallet1Pos.y": fts(state.mallet1Pos.Y),
		"mallet2Pos.x": fts(state.mallet2Pos.X),
		"mallet2Pos.y": fts(state.mallet2Pos.Y),
		"puckPos.x": fts(state.puckPos.X),
		"puckPos.y": fts(state.puckPos.Y),
	}
}

const leftBound float64 = -0.5
const rightBound float64 = 0.5
const farBound float64 = -0.8
const nearBound float64 = 0.8

func cycle()  {
	state.puckVec = scale(state.puckVec, 0.99)
	state.puckPos = add(state.puckPos, state.puckVec)

	if state.puckPos.X < leftBound + puckRadius || state.puckPos.X > rightBound - puckRadius {
		println("Hit wall lr")
		state.puckVec = Vec2d{X: -state.puckVec.X, Y: state.puckVec.Y}
		state.puckVec = scale(state.puckVec, 0.9)
	}

	if state.puckPos.Y < farBound+ puckRadius || state.puckPos.Y > nearBound - puckRadius {
		println("Hit wall tb")
		state.puckVec = Vec2d{X: state.puckVec.X, Y: -state.puckVec.Y}
		state.puckVec = scale(state.puckVec, 0.9)
	}


	state.puckPos = Vec2d{
		X: clamp(state.puckPos.X, leftBound + puckRadius, rightBound - puckRadius),
		Y: clamp(state.puckPos.Y, farBound + puckRadius, nearBound - puckRadius),
	}
}

const fps = 60
const interval = time.Second / fps
func Loop()  {
	for {
		cycle()
		time.Sleep(interval)
	}
}


func UpdatePlayer(playerOne bool, pos Vec2d)  {
	var normalizedMalletPos Vec2d
	var currentMalletPos *Vec2d

	if playerOne {
		normalizedMalletPos = normalizeMallet1Pos(pos)
		currentMalletPos = &state.mallet1Pos
	} else {
		normalizedMalletPos = normalizeMallet2Pos(pos)
		currentMalletPos = &state.mallet2Pos
	}

	dis := distance(normalizedMalletPos, state.puckPos)

	if dis < (puckRadius + malletRadius) {
		state.puckVec = vecBetween(normalizedMalletPos, state.puckPos)
	}

	*currentMalletPos = normalizedMalletPos
}