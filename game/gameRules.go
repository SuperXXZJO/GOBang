package game

import (
	"strconv"
	"strings"
	"sync"
)

const (
	ROW = 5
	COLUMN = 5

)

var lock sync.Mutex

//棋盘
type checkerboard [ROW][COLUMN]string

//新建一个棋盘
func NewCheckboard() *checkerboard   {
	c := &checkerboard{}
	return c
}

//落子
func (c *checkerboard) LandPoint (pionts string,player string)  {

	lock.Lock()
	defer lock.Unlock()

	res :=strings.Split(pionts,",")
	x,_ := strconv.Atoi(res[0])
	y,_ := strconv.Atoi(res[1])

	c[x][y] = player

}


//落子



//判断输赢
