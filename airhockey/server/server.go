package server

import (
	"airhockey-multiplayer-server/airhockey/game"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Start() {
	gin.SetMode(gin.TestMode)
	game.Init()
	go game.Loop()
	setupRouter()
}

func setupRouter() {
	r := gin.New()
	r.GET("game", gameUpdateHandler)
	r.GET("/", func(context *gin.Context) {
		context.String(200, "OK")
	})

	r.Run()
}

func gameUpdateHandler(ctx *gin.Context) {
	playerIdString := ctx.Query("player")
	xString := ctx.Query("x")
	yString := ctx.Query("y")
	uString := ctx.Query("u")

	playerId, _ := strconv.Atoi(playerIdString)
	x, _ := strconv.ParseFloat(xString, 64)
	y, _ := strconv.ParseFloat(yString, 64)
	u, _ := strconv.ParseBool(uString)

	if u {
		game.UpdatePlayer(playerId == 1, game.Vec2d{X:x, Y:y})
	}

	state := game.GetGameState()
	ctx.JSON(200, state)
}


