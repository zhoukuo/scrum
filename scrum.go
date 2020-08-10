package main

import (
	"fmt"
	"strings"
	"time"
)

type Requirement struct {
	// 需求编号
	id int
	// 标准答案
	answer string
	// 价值点数
	score int
	// 截至时间，这里记录的是回合编号
	period int
	// ROI，投入产出比
	roi string
}

type Team struct {
	// 临时记录当前需求编号
	reqId int
	// 临时记录当前需求的答案
	answer string
	// 临时记录当前所在回合
	round int
	// 记录错误次数
	wrong int
	// 记录完成需求个数
	count int
	// 记录已完成需求的分数
	scores [REQ_COUNT]int
	// 每回合需求数量
	round_reqs [ROUND_COUNT]int
	// 每回合剩余需求数量
	round_remaining [ROUND_COUNT]int
}

// 判断需求编号是否有效
func (t *Team) isInvalid() bool {
	if t.reqId > REQ_COUNT || t.reqId < 1 {
		return true
	}
	return false
}

// 判断需求是否已完成
func (t *Team) isDone() bool {
	if t.scores[t.reqId-1] != 0 {
		return true
	}
	return false
}

// 判断需求是否过期
func (t *Team) isExpired(r *Requirement) bool {
	t.round = int(time.Now().Sub(BEGIN).Minutes()/float64(ROUND_TIME)) + 1
	if t.round > r.period {
		return true
	}
	return false
}

// 计算总分，需要减去答错扣的分数
func (t *Team) getTotal() int {
	total := 0
	for i := 0; i < REQ_COUNT; i++ {
		total = total + t.scores[i]
	}
	return total - t.wrong*WRONG_SCORE
}

// 判断答案是否正确
func (t *Team) isRight(r *Requirement) bool {
	// 兼容小写字母
	t.answer = strings.ToUpper(t.answer)
	// 兼容-分隔符
	t.answer = strings.Replace(t.answer, "-", "", -1)
	r.answer = strings.Replace(r.answer, "-", "", -1)

	if t.answer != r.answer {
		t.wrong++
		return false
	}

	t.count++
	t.scores[t.reqId-1] = r.score

	// 计算当前回合剩余需求个数
	if t.round == r.period {
		t.round_remaining[t.round-1]--
	}
	return true
}

// 记录开始时间，用于判断需求是否过期
var BEGIN time.Time

// 回合数量
const ROUND_COUNT = 4

// 需求数量，默认23个
const REQ_COUNT = 23

// 每个回合15分钟，回顾5分钟
const ROUND_TIME = 20

// 答错一次扣10分
const WRONG_SCORE = 10

// 需求信息
var reqs = []Requirement{
	Requirement{
		1,
		"DZ-50-CH-1Z",
		85,
		3,
		"高",
	},
	Requirement{
		2,
		"X83Q1P-A2671B-OQBBAR-N2M5K6-VGHJ",
		200,
		3,
		"低",
	},
	Requirement{
		3,
		"M-3-A-J",
		65,
		2,
		"高",
	},
	Requirement{
		4,
		"1X-16-2L-1O",
		35,
		1,
		"中",
	},
	Requirement{
		5,
		"3V-AF-FJ-F3",
		50,
		1,
		"中",
	},
	Requirement{
		6,
		"2YB5-6FPZ-KY77-9C7E",
		170,
		4,
		"高",
	},
	Requirement{
		7,
		"NOP8QT-8CHMRW-GK87PT-6K5M2N-K2XP",
		180,
		2,
		"低",
	},
	Requirement{
		8,
		"87-A5-0J-8I",
		55,
		4,
		"中",
	},
	Requirement{
		9,
		"HUW2-7B49-YA9W-D5R8-YI4A",
		100,
		2,
		"低",
	},
	Requirement{
		10,
		"R4IS5B-KPUZ50-N5EU6D-ILE7TB-J30C",
		220,
		1,
		"低",
	},
	Requirement{
		11,
		"4329-EJOT-BO2Z-S37W",
		50,
		3,
		"低",
	},
	Requirement{
		12,
		"8A-NS-DS-4V",
		20,
		2,
		"中",
	},
	Requirement{
		13,
		"A-1-X-3",
		50,
		4,
		"高",
	},
	Requirement{
		14,
		"71-Y4-38-9U",
		55,
		3,
		"中",
	},
	Requirement{
		15,
		"6C-X3-M4-2X",
		95,
		2,
		"高",
	},
	Requirement{
		16,
		"6ACC-UKRL-9T5W-ODF5",
		130,
		4,
		"高",
	},
	Requirement{
		17,
		"70AC-27BG-Q69B-9G0Q-LDVB",
		30,
		1,
		"低",
	},
	Requirement{
		18,
		"FD9E-LQV1-075I-A4DC",
		140,
		3,
		"高",
	},
	Requirement{
		19,
		"TRKHIU-B832MC-D5408U-8HD4A0-IQDG",
		140,
		3,
		"中",
	},
	Requirement{
		20,
		"0-9-1-0",
		40,
		3,
		"高",
	},
	Requirement{
		21,
		"B5-8C-1C-6Y",
		30,
		2,
		"低",
	},
	Requirement{
		22,
		"G6-6A-3R-H8",
		120,
		1,
		"高",
	},
	Requirement{
		23,
		"BF-VH-L6-RA",
		55,
		4,
		"中",
	},
}

// 初始化
func init() {
	fmt.Printf("\n%s ", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("演练开始，每个回合%d分钟，开始计时！\n\n", ROUND_TIME-5)
	fmt.Printf("----------------------------------------------------\n\n")
	// 比赛计时器
	BEGIN = time.Now()
}

// 主要逻辑
func main() {

	var t Team
	t.round_reqs = [ROUND_COUNT]int{5, 6, 7, 5}
	t.round_remaining = t.round_reqs
	for {

		fmt.Printf("%s ", time.Now().Format("2006-01-02 15:04:05"))
		fmt.Printf("请输入需求编号：\t\t#")
		fmt.Scan(&t.reqId)

		if t.isInvalid() {
			fmt.Printf("%s ", time.Now().Format("2006-01-02 15:04:05"))
			fmt.Printf("需求编号无效，请重新输入！\n\n")
			fmt.Printf("----------------------------------------------------\n\n")
			continue
		}

		if t.isDone() {
			fmt.Printf("%s ", time.Now().Format("2006-01-02 15:04:05"))
			fmt.Printf("需求已完成，请重新输入!\n\n")
			fmt.Printf("----------------------------------------------------\n\n")
			continue
		}

		fmt.Printf("%s ", time.Now().Format("2006-01-02 15:04:05"))
		fmt.Printf("请输入识别码：")
		fmt.Scan(&t.answer)

		if t.isExpired(&(reqs[t.reqId-1])) {
			fmt.Printf("%s ", time.Now().Format("2006-01-02 15:04:05"))
			fmt.Printf("当前所在回合：\t%9d/4\n", t.round)
			fmt.Printf("%s ", time.Now().Format("2006-01-02 15:04:05"))
			fmt.Printf("此需求截至时间: \t%9d/4\n", reqs[t.reqId-1].period)
			fmt.Printf("%s ", time.Now().Format("2006-01-02 15:04:05"))
			fmt.Printf("需求已过期，请重新输入!\n\n")
			fmt.Printf("----------------------------------------------------\n\n")
			continue
		}

		if t.isRight(&(reqs[t.reqId-1])) {
			fmt.Printf("%s ", time.Now().Format("2006-01-02 15:04:05"))
			fmt.Printf("截至时间: \t\t%9d/4\n", reqs[t.reqId-1].period)
			fmt.Printf("%s ", time.Now().Format("2006-01-02 15:04:05"))
			fmt.Printf("投入产出比(ROI): \t%10s\n", reqs[t.reqId-1].roi)
			fmt.Printf("%s ", time.Now().Format("2006-01-02 15:04:05"))
			// fmt.Printf("%c[1;31;32m回答正确！\t\t      +%4d%c[0m\n\n", 0x1B, reqs[t.reqId-1].score, 0x1B)
			fmt.Printf("回答正确！\t\t      +%4d\n\n", reqs[t.reqId-1].score)

			fmt.Printf("%s ", time.Now().Format("2006-01-02 15:04:05"))
			fmt.Printf("已完成需求数量: \t%11d\n", t.count)
			fmt.Printf("%s ", time.Now().Format("2006-01-02 15:04:05"))
			fmt.Printf("当前错误次数：\t%11d\n", t.wrong)
			fmt.Printf("%s ", time.Now().Format("2006-01-02 15:04:05"))
			fmt.Printf("当前所在回合：\t%9d/4\n", t.round)
			fmt.Printf("%s ", time.Now().Format("2006-01-02 15:04:05"))
			fmt.Printf("当前回合剩余时间：\t%7d分钟\n", ROUND_TIME-int(time.Now().Sub(BEGIN).Minutes())%(ROUND_TIME)-5)
			fmt.Printf("%s ", time.Now().Format("2006-01-02 15:04:05"))
			fmt.Printf("当前回合剩余需求：\t%11d\n", t.round_remaining[t.round-1])
			fmt.Printf("%s ", time.Now().Format("2006-01-02 15:04:05"))
			fmt.Printf("总分： \t\t%11d\n\n", t.getTotal())

			fmt.Printf("----------------------------------------------------\n\n")

		} else {
			fmt.Printf("%s ", time.Now().Format("2006-01-02 15:04:05"))
			fmt.Printf("截至时间: \t\t%9d/4\n", reqs[t.reqId-1].period)
			fmt.Printf("%s ", time.Now().Format("2006-01-02 15:04:05"))
			fmt.Printf("价值点数：\t\t%11d\n", reqs[t.reqId-1].score)
			fmt.Printf("%s ", time.Now().Format("2006-01-02 15:04:05"))
			// fmt.Printf("%c[1;31;31m回答错误！\t\t\t-10%c[0m\n\n", 0x1B, 0x1B)
			fmt.Printf("回答错误！\t\t\t-10\n\n")

			fmt.Printf("%s ", time.Now().Format("2006-01-02 15:04:05"))
			fmt.Printf("已完成需求数量： \t%11d\n", t.count)
			fmt.Printf("%s ", time.Now().Format("2006-01-02 15:04:05"))
			fmt.Printf("当前错误次数：\t%11d\n", t.wrong)
			fmt.Printf("%s ", time.Now().Format("2006-01-02 15:04:05"))
			fmt.Printf("当前所在回合：\t%9d/4\n", t.round)
			fmt.Printf("%s ", time.Now().Format("2006-01-02 15:04:05"))
			fmt.Printf("当前回合剩余时间：\t%7d分钟\n", ROUND_TIME-int(time.Now().Sub(BEGIN).Minutes())%(ROUND_TIME)-5)
			fmt.Printf("%s ", time.Now().Format("2006-01-02 15:04:05"))
			fmt.Printf("当前回合剩余需求：\t%11d\n", t.round_remaining[t.round-1])
			fmt.Printf("%s ", time.Now().Format("2006-01-02 15:04:05"))
			fmt.Printf("总分： \t\t%11d\n\n", t.getTotal())

			fmt.Printf("----------------------------------------------------\n\n")

		}
	}
}
