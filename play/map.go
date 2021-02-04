package play

// NumberToString 是将数字翻译为对应扑克牌的函数
var NumberToString map[int]string

func init() {
	Translation := map[int]string{
		0:  "梅花3",
		1:  "梅花4",
		2:  "梅花5",
		3:  "梅花6",
		4:  "梅花7",
		5:  "梅花8",
		6:  "梅花9",
		7:  "梅花10",
		8:  "梅花J",
		9:  "梅花Q",
		10: "梅花K",
		11: "梅花A",
		12: "梅花2",
		13: "方块3",
		14: "方块4",
		15: "方块5",
		16: "方块6",
		17: "方块7",
		18: "方块8",
		19: "方块9",
		20: "方块10",
		21: "方块J",
		22: "方块Q",
		23: "方块K",
		24: "方块A",
		25: "方块2",
		26: "红桃3",
		27: "红桃4",
		28: "红桃5",
		29: "红桃6",
		30: "红桃7",
		31: "红桃8",
		32: "红桃9",
		33: "红桃10",
		34: "红桃J",
		35: "红桃Q",
		36: "红桃K",
		37: "红桃A",
		38: "红桃2",
		39: "黑桃3",
		40: "黑桃4",
		41: "黑桃5",
		42: "黑桃6",
		43: "黑桃7",
		44: "黑桃8",
		45: "黑桃9",
		46: "黑桃10",
		47: "黑桃J",
		48: "黑桃Q",
		49: "黑桃K",
		50: "黑桃A",
		51: "黑桃2",
		52: "小王",
		53: "大王",
	}
	NumberToString = Translation
}
