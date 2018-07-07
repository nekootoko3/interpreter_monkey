package aa

import (
	"math/rand"
	"time"
)

type Aa []string

const ACCELERATE_YOUR_SPIRIT = `
精神を加速させろ
　　　　　　∩＿∩
　　　　　／ ＼ ／＼
　　　　 ｜ (ﾟ)=(ﾟ)｜
　　　　 ｜　●_● ｜
　　　　_/　　　　　ヽ
　　　 /｜〃―――-ヾ|
　　　/　／~)_二＿＿ノ
　　 / ／　/　　　　 )
　　｜　 ／　　　/ ／
　　 ＼／　　　 /￣
　　　/ 　 ／⌒＼
　　 /　 ／　　 ｜
　　/　　　 /　 /
　 /　　 ／/　／
　｜　 ／ / ／
　｜　/ ／ /
　｜ /　L／
'／ /
(_／
`

const SAVANNAH = `
お前それサバンナでも
同じ事言えんの？

　　ノ从从从从ヽ
　(⌒／ﾞﾞﾞﾞﾞﾞ＼⌒)
　ノｲ ＿　　＿｜ヽ
　彡|ヽ･〉〈･ﾉ｜ミ
　彡|　　▼　 ｜ミ
　彡ヽ ＿人＿ / ミ
'／ヾヽ '⌒′/ ツ＼
｜　ヾ ﾞﾞﾞﾞﾞﾞ ツ　｜
｜　| ヾ从从ツ |　｜
｜　'――――――⌒)
(＼＿＿＿＿＿＿＿＿)
（⌒　　　　　　 ノ
　￣|￣￣￣￣￣Ｔ
`

const MAGAO = `
　　 ／￣￣￣＼
　　/::　　　::ヽ
　 f:ｨ●ｧ　ｨ●ｧ:|
　 |::　　　　::|
　 |::　c{ っ ::|
　 |::　＿＿　::|
　 ヽ::　ー　::ノ
　　 ＼::　::／
　　　｜::::｜
　　　｜::::｜
'／￣￣　　　￣￣＼
｜::　　　　　　::｜
｜::　　　　　　::｜
`

const NYAAN = `
　∧Ｍ∧Ｍ∧
＜ ニャーン ＞
　ＷＷＶＷＶ
　　　　　 ∩＿＿_∩
　　　　　 |ノ　　 ヽ
　　　　　/　●　 ●|
　 (ヽ＿／｜　(_●_)ミ
　／　　　彡　 |∪|ミ
'/　 ｜　　 ＼ ヽノ/
｜　 /　　　　￣￣｜
/ _／＞、(　<＿／ /
L/　/ ／￣＼_) ヽ｜
　 (_/　　　　　|_)

　/ |　/　 ￣/ /二/＿
／　ヽ/＿ヽ／＼　/
`
const KUMA = `
これはクマった…
　　　 ∩＿＿＿∩
　　　 | ノ　　 ヽ
　　 /⌒) ●　 ●|
　　/　/　 (_●_)ミ
　 (　 ヽ　|∪|　＼
　　＼　　 ヽノ/>　)
　　 /　　　　/(_／
　　/　 ／⌒ヽ⌒ヽ
　 (　　　　ノ　ノ
　(￣￣￣￣￣￣￣￣)
|￣|￣|￣|￣|￣|￣|￣|
`

const PON_DE_LION = `
そんな事言われてもウチ

ポン・デ・ライオンやし

//////////////////////
//////／￣＼／￣＼////
//// ｜　　｜　　｜///
//／￣／￣￣~＼／￣＼/
/｜　/●　　　｜　　｜
//＼｜ (人_) ●＼　／/
//／｜　Lノ　　／￣＼/
/｜　＼　　　 ｜　　｜
//＼_／￣＼／￣＼＿／/
////｜　　｜　　｜////
/////＼＿／＼＿／/////
////////　　 ｜///////
`

const ITTAIKAN = `
すごい一体感を感じる
今までにない
何か熱い一体感を

　　　∩＿∩
　　／ ＼ ／＼
　 ｜ (ﾟ)=(ﾟ)｜
　 ｜　●_● ｜
　 /　　　　　＼
　｜〃―――-ヾ｜
　 ＼＿_二＿＿／
`

func New() Aa {
	aa := make(Aa, 0, 7)
	aa = append(aa, ACCELERATE_YOUR_SPIRIT)
	aa = append(aa, SAVANNAH)
	aa = append(aa, MAGAO)
	aa = append(aa, NYAAN)
	aa = append(aa, KUMA)
	aa = append(aa, PON_DE_LION)
	aa = append(aa, ITTAIKAN)
	return aa
}

func (aa Aa) PrintAa() string {
	rand.Seed(time.Now().UnixNano())
	return aa[rand.Intn(7)]
}
