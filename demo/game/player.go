package game

import (
	"github.com/xiaomingping/landlord/poker"
	"github.com/xiaomingping/landlord/pokergame"
)

type User struct {
	Id       int    `gorm:"column:id;primary_key;auto_increment;"`
	NickName string `gorm:"column:nickname;unique;type:varchar(100);index:name;not null"`
	Avatar   string `gorm:"column:avatar;type:varchar(255);not null"`
}

type IPlayer interface {
	//由game回调的方法
	PlayCardSuccess(cardIndexs []int) //出牌成功
	PlayCardError(err string)         //出牌错误
	GetReadyStatus() bool
	GetAutoStatus() bool
	GetPlayerUser() *User
	GetIndex() int
	SetIndex(index int)
	SetPokerCards(cards poker.PokerSet)
	StartCallScore()
	StartPlay()
	IsOutOfCards() bool //是否出完牌
	SendMsg(msg []byte)

	//响应客户端请求的方法
	JoinGame(gameType int, gameId int)    //加入游戏
	CreateGame(gameID int, baseScore int) //创建游戏
	LeaveGame()                           //离开游戏
	SayToOthers(msg []byte)               //给同一游戏的其他玩家发消息
	SayToAnother(id int, msg []byte)      //给同意游戏的特定索引玩家发消息
	ResolveMsg(msgB []byte) error         //处理客户端消息
	Ready()                               //玩家准备
	UnReady()                             //玩家取消准备
	CallScore(score int)                  //抢地主
	PlayCards(cards []int)                //出牌
	Pass()                                //过牌
	HintCards()                           //提示出牌
	GetPlayedCardIndexs() []int
	GetPlayerCards(indexs []int) poker.PokerSet

	SetPokerRecorder(recorder pokergame.IRecorder)
	SetPokerAnalyzer(analyzer pokergame.IAnalyzer)
}
