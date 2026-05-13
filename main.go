package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"google.golang.org/genai"
)

func main() {
	fmt.Println("Hello World!")

	ctx := context.Background()
	// Client gets API key automatically from environment variable (GEMINI_API_KEY)
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	input := fmt.Sprint(
		`I need you to translate this chinese webnovel chapter into english. Here are 2
		earlier chapters both raw and then translated:
		I'd like another chapter translated, here are the preceding 2 chapters both raw and then translated:

Chapter 1045: Storm I

乌迪尔站在纯白色的巨大起源树上,远远遥望着星空中那颗正在急速崩塌的四季之树。

那庞大甚至能比拟星系的巨树,此时正犹如在烈火中,正在收缩的海绵。

他能够看到翠绿的树枝树叶在枯萎焦黑,大量的花朵和果实纷纷化为白色灰烬。

“这是第四颗被虚无毁灭的四季之树。”

一旁的猎魔人低声带着哀戚道。

乌迪尔拄着手中的囚笼之杖转过身。

这根他用虚灵大妖兽的头颅打造而成的手杖,顶端有着半透明的无数触须正在四处挥舞。也是他无数年来和虚无抗争中,得力的左膀右臂。

“接下来虚无的动向如何?”他低沉问。

“虚灵界已经同时向七千多个不同世界宇宙进行同步侵蚀。我们的世界只是其中之一。根据调查反馈的情报来看,其中超能级三个,高能级十七个。数量还在缓慢增加中。”

一旁皮肤洁白的女性猎魔人迅速回答道。

“犄角防线呢?”乌迪尔沉声问。

“犄角防线的其中一端。天魔界的星灵联盟崩溃了其余两个点恐怕支撑不了多久军团长,我们必须尽快做好后撤准备。”一个猎魔人提醒道。

乌迪尔环视周围,最初跟随他离开族群的猎魔人,还有上千,而现在,只剩下五十一个

这些猎魔人,每一个都在长期的厮杀中,身体被虚无之力侵染,成了半虚无半实体的半透明怪物。

他们的右臂变得和那些虚无守卫者相似,全是半透明的无色触须。

现在,其实已经很难界定他们还算不算起源族人了

但不管身体如何改变,乌迪尔依旧能从这些兄弟姐妹眼中,看到不灭的希望。

“撤吧。不过,我们还会回来。”乌迪尔嘶哑着声音道。

他的嗓子早在常年操纵零质火的过程中被烧伤。很难想像,一个本身便是从起源之火出生的特殊火焰族人,会有一天被火焰烧伤。

但这就是代价。

为了对抗虚无,他用自己一半的灵魂为代价,换来了最强大的灵魂火焰零质火,的掌控权。

但虚无的力量实在太庞大了,那无穷无尽的虚无守卫者仿佛根本杀不光。

它们本身就代表着毁灭,就算被杀,也能够源源不断的重生出现。

族群的防线节节败退,从最初开始,到现在已经抗争了无数年。但得来的结果,却依旧是毫无希望。

“撤吧”乌迪尔最后看了眼那庞大的正在毁灭的四季之树,手杖轻轻一顿,整个人渐渐淡化,消失在原地。

其余猎魔人也纷纷淡化消失,只剩下孤零零的纯白起源树,依旧留在这片宇宙里,静静等待着最终的毁灭。

***

***

“废话少说,说吧,到底谁让你们绑我过来的?”

“我明明刚刚在做饭啊?怎么?”

“我手机呢?我电脑呢??”

路胜无语的看着这群说着乱七八糟语言的地球人。

是的,虽然不知道他们是不是来自于他当初那个地球,但这些人说的话,有汉语,有韩语,还有英语法语。

他基本都听出来了。虽然他不懂汉语之外的语种,但那种口音,大概是能听出来的。

很明显,这群人自己都没搞清楚是怎么回事,就被传送到这里来了。

路胜目光扫视一圈,很快便在这群人身上,发现了一些特异之处。

这一百人闹哄哄的挤在一起,又相互之间保持了一定程度的安全距离。

而其中的一小撮人,大约十几人的数量,一直保持着冷静。

和其他闹哄哄的人不同,这些人看起来似乎并不惊讶自己忽然来到这里。

路胜感觉到,这群人身上或多或少的都带着一些细微特异处。

“交给你了,按照外围成员一样训练。”路胜对一边的白郡城低声道。

“是。”白郡城点头。

“现在,都给我听好!”他直接大声呵斥起来。“你们在这里,将会受到最严苛的高强度训练。不管你以前是干什么的?不管你以前是什么身份,来到这里,那就是我们九命堂的外围训练生”

他训练这些新人简直不要太容易。

最近的九命堂新人,其实都是他和魏韩冬一起训练的,所以这一百多人他也是游刃有余。

比较奇妙的是,路胜看出,这些人明明来自不同国家地区,但却都能统一的听懂白郡城所说的话。

“先观察下,有什么异常随时给我禀报。”路胜低声对白郡城道。

“明白。”

路胜交代好后,便不去管他们,在总堂停留了一个多小时,他分别个给核心班的弟子指点一二技艺,然后打算出去吃饭。

刚从大门走出去,没几步。

早已等候在外的九命堂弟子恭敬的给他打开车门。

“堂主,能不能借一步说话。”右侧早已等候多时的几人,一个穿暗红色短袖,牛仔长裤的漂亮女子,起身试图靠近路胜。但马上便被九命堂的人拦住。

路胜扫视了眼她,这女子看起来在十八九岁的样子,但眼神里透出的年纪,绝不只是十八九岁。

她容貌很漂亮,几乎看不到半点瑕疵,有种白玉的味道。五官也很标致,胸部高耸,腰肢纤细,臀腿圆润修长,几乎是最标准的美女模板。

但就是给路胜一种不自然的感觉。

“有什么事?”路胜也对这群人很好奇。

记得他曾经在地球时,还看过一些书,上边的情况和眼下他遇到的这些人,似乎很类似。

女子柔媚的笑了笑,挺了挺胸走得近了些。

拦住她的九命堂弟子也主动让开,让其走近。

“堂主,不知道您知不知道,用不了多久,这里,这颗星球,就将迎来最终的毁灭?”

女子一开口便是危言耸听之语。

“哦?”

路胜饶有兴趣的看着这人。

如果是普通人,或许已经把对方当成是神经病了。但他不同。他一开始就知道,这群人似乎不简单。

“您或许不信,但如果说,我能预测您接下来会发生的诸多事件,不知道您愿不愿意相信我。”女子认真低声道。

奇妙的是,她的话语对于近在咫尺的两个九命堂弟子,似乎毫无影响。他们压根就似乎没听到她说话。

“你确定?你们和蓝色星光是什么关系?”路胜直白问道。

“只是雇佣关系。”女子微微松了口气,看样子这位九命堂堂主似乎感兴趣了。他们就怕对方毫无兴趣,让他们连靠近的机会也没。

感兴趣就好,他们有太多的办法能让对方感兴趣。

她悄悄看了眼路胜脸色,继续道。

“我们其实一共分成了两批人,一批跟着蓝色星光,一批则是在另外的地方。我们这一批,人手虽然多,但各个实力都不怎么样。

而另一批就不同了”

女子滔滔不绝的开始给路胜普及起来他们内部的一些信息。

很快路胜便知道了,她叫苏芩,她身后的四个人,都是和她一个团队的。团队名字叫歌舞升平。

路胜针对他们的来历,随意询问了几句,但得到的不是不能回答,就是说到一半,苏芩忽然脸色一白,彻底忘了他刚才在问什么。

重复几次后,路胜便明白,似乎有着某种机制,死死的将他们的思维阻挡在某些信息之外。不让他们自由说出来。

“那么,你们找我,有什么目的?”路胜直接进入正题。

苏芩郑重道。

“我们,希望能借助您的力量,保证自身安全。作为代价,我们愿意付出您感兴趣的任何东西。”

路胜大概已经了解了她们的模式。他们身上有着某种能力,似乎能距离极远都实时传递信息。

“你们有什么东西能让我感兴趣?”路胜反问。

“大难将至,我们能帮你避开最麻烦的几次大灾。”苏芩自信道。

对于这个电影里实力不错,但运气实在不怎么样的王木,他们这次是势在必得,一定要争取对方的好感。

虽然在母星战争这部电影里,这个倍加星只不过是星际战争中的一颗普通星球。

也就是主角在偶然流落时,在这颗星球上度过了一阵岁月时光。才让这儿能在荧幕上有点存在感。

和其他人不同,苏芩之前就很熟悉这部电影,对于电影中,在倍加星只出场过很短时间的九命堂王木,有着很深的印象。

这个穷其一生,都在追求武道的强大人类格斗大师,曾经短暂的抵挡住过念能联盟和蓝色星光的巨大压力。

在念能联盟和蓝色星光都被虚无侵蚀渗透的大局下,九命堂虽然只支撑了很短一段时间便覆灭。

但在主角还没遭遇大难前,这里应该算是少有的电影里最安全的一段时间。

而九命堂,虽然只是个土著星球的一方豪强,但若是利用得当,未必不能得到巨大的回报和利益。

而且,苏芩心头还有个隐晦的念头,若是能够把九命堂堂主,直接拉进自己队伍

虽然比不上那些念能联盟和蓝色星光顶尖的操纵使,但王木也是走在人类极限的强者。只不过比不过那些顶尖黑手罢了。

“找个地方坐下聊吧。”路胜明显看出了这个苏芩没有说谎。

对于心理引导术上千级的他而言,判断一个人类说没说谎,是再简单不过的事。

“多谢堂主信任!”苏芩大喜,赶紧抱拳。

Udir stood atop the massive, pure white Origin Tree, gazing from afar at the Four Seasons Tree rapidly collapsing in the starry void.

The colossal tree, so vast it rivaled a galaxy, now seemed like a sponge shrinking in a raging fire.

He watched its emerald branches and leaves wither and blacken, its countless flowers and fruits turning to white ash.

"This is the fourth Four Seasons Tree destroyed by the Nothingness," the Demon Hunter beside him murmured, his voice filled with sorrow.

Udir turned, leaning on the Cage Staff in his hand.

This staff, crafted from the skull of a Void Spirit Great Demon Beast, had semi-transparent tendrils writhing at its tip. It had been his trusted left arm in his countless years of battling the Nothingness.

"What's the Nothingness's next move?" he asked in a deep voice.

"The Void Spirit Realm is simultaneously invading over seven thousand different worlds and universes. Our world is just one of them. Intelligence reports indicate three super-tier and seventeen high-tier incursions. The numbers are still slowly increasing."

A fair-skinned female Demon Hunter beside him quickly answered.

"What about the Horned Defense Line?" Udir asked in a deep voice.

"One end of the Horned Defense Line. The Star Spirit Alliance in the Demonic Realm has collapsed. The other two points won't hold much longer, Legion Commander. We must prepare to retreat immediately," a Demon Hunter warned.

Udir surveyed his surroundings. Of the thousand Demon Hunters who had initially followed him out of the tribe, only fifty-one remained.

Each of these Demon Hunters, after prolonged combat, had been infected by the Power of Emptiness, transforming into semi-transparent monsters that were half-ethereal and half-physical.

Their right arms now resembled those of the Void Guardians, covered in semi-transparent, colorless tendrils.

It was now difficult to determine if they still qualified as Origin Clan members.

Yet despite their physical changes, Udir could still see an unyielding hope in the eyes of these brothers and sisters.

"Withdraw. But we will return," Udir said hoarsely.

His throat had been scorched long ago from years of manipulating the Zero-Matter Fire. It was hard to imagine that an Origin Clan member, born from the very heart of the primordial flame, could ever be burned by fire.

But this was the price.

To fight the Nothingness, he had sacrificed half his soul to gain control of the most powerful soul flame: the Zero-Matter Fire.

Yet the power of the Nothingness was simply too vast. The endless ranks of Void Guardians seemed impossible to eradicate.

They embodied destruction itself, and even when slain, they would rise again in an endless cycle.

The clan's defenses crumbled, and after countless years of resistance, the result remained as hopeless as ever.

"Withdraw," Udir said, casting one last glance at the colossal, decaying Four Seasons Tree. He tapped his staff lightly, and his form gradually faded until he vanished from the spot.

The remaining Demon Hunters followed suit, leaving only the solitary, pure white Origin Tree to remain in this corner of the Cosmos, silently awaiting its final destruction.

***

"Enough talk. Tell me, who ordered you to bring me here?"

"I was just cooking! What's going on?"

"Where's my phone? My laptop?"

Lu Sheng stared wordlessly at the group of Earthlings babbling in a jumble of languages.

Yes, though he couldn't be sure if they were from his original Earth, their speech included Chinese, Korean, English, and French.

He could understand most of it. While he only knew Chinese, the accents were distinct enough to identify the languages.

It was clear this group hadn't figured out what was happening either, having been teleported here without warning.

Lu Sheng scanned the crowd and quickly noticed some peculiar traits.

The hundred people huddled together in a noisy cluster, yet maintained a certain safety distance between themselves.

A small group, about a dozen individuals, remained calm.

Unlike the clamoring crowd, these people didn't seem surprised by their sudden arrival.

Lu Sheng sensed subtle anomalies in each of them.

"Leave them to me. Train them like the outer members," he whispered to Bai Juncheng beside him.

"Understood," Bai Juncheng nodded.

"Listen up, everyone!" he barked, his voice echoing through the hall. "Here, you will undergo the most rigorous and high-intensity training. Regardless of your past profession or status, once you set foot here, you are nothing more than peripheral trainees of the Nine Lives Hall."

Training these newcomers was child's play for him.

Most recent recruits had been trained by him and Wei Handong, so managing this group of over a hundred was effortless.

What struck Lu Sheng as peculiar was that despite their diverse origins, all these individuals understood Bai Juncheng's words perfectly.

"Observe them for now. Report any anomalies immediately," Lu Sheng murmured to Bai Juncheng.

"Understood."

After giving his instructions, Lu Sheng paid them no further attention. He spent over an hour at the Main Hall, offering brief technical guidance to the core disciples before preparing to leave for a meal.

He had barely stepped beyond the main gate when a Nine Lives Hall disciple waiting outside respectfully opened his car door.

"Hall Master, may I have a moment of your time?" A beautiful young woman, dressed in a dark red short-sleeved shirt and jeans, rose from the group waiting to his right and attempted to approach Lu Sheng, but was immediately restrained by the Hall members.

Lu Sheng glanced at her. She appeared to be around nineteen, yet her eyes betrayed a maturity far beyond her years.

Her beauty was almost flawless, like polished white jade. Her features were perfectly proportioned: a high-set chest, a slender waist, and rounded, shapely hips and legs—the very embodiment of the ideal female form.

Yet, something about her felt unnatural to Lu Sheng.

"What is it?" Lu Sheng was equally curious about this group.

He recalled reading books on Earth that described situations strikingly similar to the one he was now facing.

The woman smiled coquettishly, puffing out her chest as she moved closer.

The Nine Lives Hall disciples who had blocked her path stepped aside, allowing her to approach.

"Hall Master, do you know that this planet—and everything on it—is about to face its final destruction?"

The woman's words were alarmist.

"Oh?"

Lu Sheng watched her with keen interest.

An ordinary person might have dismissed her as a lunatic. But he was different. He had sensed from the beginning that this group was no ordinary faction.

"You may not believe me, but if I told you I could predict the many events that will unfold before you, would you be willing to trust me?" she whispered earnestly.

Strangely, her words seemed to have no effect on the two Nine Lives Hall disciples standing close by. They acted as if they hadn't heard her speak at all.

"Are you sure? What's your relationship with Blue Starlight?" Lu Sheng asked directly.

"Just a contractual relationship," the woman replied, visibly relieved. It seemed the Hall Master of Nine Lives Hall was interested. They had feared he would be indifferent, leaving them no chance to even approach him.

*Good, he's interested.* They had plenty of ways to make him even more interested.

She subtly glanced at Lu Sheng's expression before continuing, "We're actually divided into two groups. One group follows Blue Starlight, while the other is elsewhere. Our group has many members, but none are particularly strong."

"The other group is different..."

The woman began to eagerly share internal information with Lu Sheng.

Soon, Lu Sheng learned her name was Su Qin, and the four people behind her were all part of her team, named Peaceful Prosperity.

Lu Sheng casually inquired about their origins, but his questions were either met with refusal or left half-answered. Su Qin's face would suddenly pale, and she would completely forget what he had asked.

After repeating this several times, Lu Sheng realized there must be some mechanism rigidly blocking their minds from revealing certain information, preventing them from speaking freely.

"So, what's your purpose in seeking me out?" Lu Sheng cut to the chase.

Su Qin said solemnly, "We hope to borrow your power to ensure our safety. In return, we're willing to offer anything that interests you."

Lu Sheng had already figured out their method. They possessed some kind of ability that allowed them to transmit information in real-time over vast distances.

"What do you have that would interest me?" Lu Sheng countered.

"A great disaster is approaching. We can help you avoid the most devastating catastrophes," Su Qin said confidently.

They were determined to win over Wang Mu this time, a character from the movie who was skilled but notoriously unlucky.

Although Beijia Star was merely an ordinary planet in the grand scheme of the *Mother Planet War* film, its brief appearance on screen was due to the protagonist's accidental stay there.

Unlike the others, Su Qin was deeply familiar with the movie. She had a particularly strong impression of Wang Mu from Nine Lives Hall, who only had a fleeting role on Beijia Star.

This human martial artist, who had dedicated his entire life to pursuing ultimate power in the Martial Dao, had once briefly withstood the immense pressure from the Psychic Alliance and Blue Starlight.

Even though the Nine Lives Hall was ultimately annihilated after a short resistance, under the overwhelming tide of the Nothingness that had infiltrated both the Psychic Alliance and Blue Starlight, this period remained one of the safest in the film before the protagonist faced their greatest trials.

Though the Nine Lives Hall was merely a local power on an indigenous planet, if properly utilized, it could potentially yield immense rewards and benefits.

Moreover, Su Qin harbored a hidden thought: what if she could directly recruit the Hall Master into her own ranks?

While not as powerful as the top Manipulators of the Psychic Alliance and Blue Starlight, Wang Mu was still a formidable warrior pushing the boundaries of human potential, albeit not quite reaching the level of the elite masterminds.

"Let's find a place to sit down and talk," Lu Sheng said, clearly sensing Su Qin wasn't lying.

For someone with over a thousand levels in the Psychological Guidance Technique, discerning whether a human was lying was as simple as breathing.

"Thank you for your trust, Hall Master!" Su Qin exclaimed, her face lighting up as she quickly clasped her hands in a formal greeting.

---

Chapter 1046: Storm II

片刻后。

几人坐在一间附近的茶楼静室里。

侍者上了一壶茶后,便缓缓安静退下。

静室门关拢。

苏芩和路胜面对面坐下,其余四个形态各异的年轻男女,则分散在她身后坐下。

和她不同,其余四人显得颇为紧张,似乎对他颇有畏惧。

路胜一身黑色西装外套,盘着腿,一手捏起茶杯喝了口。

“你们,似乎对我有过了解?说说看?”

他感觉从这几人口中,或许能得到一些自己完全没有预料到的情报。

苏芩清了清嗓子。

“还请您摒退周围,检查下是否有监听装置。

另外,接下来我要说的内容,可能对您有很大冲击,还请您千万要冷静!”

“芩姐!”她身后一个男子有些担心的叫了声,“还是我来吧?”

“别担心,我知道怎么处理。”苏芩扬手郑重道。

路胜不以为意。

拿出手机发了条语音通信,很快,外面守卫的九命堂高手都纷纷离开。

然后他又看到苏芩从怀里拿出一个小巧的化妆镜一样的东西,拿着对着周围照射了一圈。

“现在好了,没有任何额外装置。”苏芩松口气。

“说吧,什么事这么郑重。”路胜放下茶杯淡淡道。

苏芩咬住嘴唇,深吸一口气。

“堂主,如果我没记错的话,你应该是在两年前,才真正开始崛起的吧?”

“不错。九命堂也是那时候成立。这种事到处都能查到,毫无意义。”路胜不在意道。

“而您崛起的关键,应该就在那个闻达图书馆吧?”苏芩再度道。

“在图书馆里,您无意中找到了一本名为千锁定龙功的强大功决,从而一举突破限制,走上了超越人体极限的道路。”

苏芩缓了口气,继续道。

“您一开始以办理补习班的方式,缓慢积累。之后无意中遇到您的左膀右臂郑欢,在正式击败他后,将其收服,带领门下诸多弟子,开始九命堂真正的扩张之路。”

“如果您不信,那本千锁定龙功,当时应该是藏在闻达图书馆的第二层,古籍区第二个书架最下方。

功决的表面是用天机地理四个大字作为掩护,表面记录的是天文地理方面的知识,但如果用油浸泡之后,就能显露出真实内容”

路胜一开始还只是当成故事在听,但在苏芩甚至给出了极其详细精确的方位和描述后,他顿时有些迟疑起来。

“现在,您该相信我们了吧?”苏芩一口气说完自己掌握的信息,面色依旧肃然。

“如果说这些,您可以说我们是情报能力很强,掌握得当,那么接下来,我可以向您证明,我们能帮你避开最大的麻烦。”

“什么麻烦?”路胜不知不觉间,面色也微微肃然起来。

“倍加星,会在一年之后,毁于星际战争中的奇数死光射击。

念能联盟和蓝色星光全面开战,双方的操纵使每一位都是灭星级存在。而操纵使之上,还有更强存在的暗能使”苏芩沉声道。

“如果您不相信,我们可以提供一点小小的私人情报。”

“什么情报?”路胜不知不觉已经一开始重视这个苏芩所说的内容了。

“您门下的一个弟子,一个名叫姚崇的男子,会在后天和人因为酒醉斗殴结仇,他会被当场打断右手,回来求援。您到时候就能判断,我们所提供的情报是否真实了。”苏芩自信道。

路胜沉默了下。

“我明白了。你的话,我会仔细考虑。那么”

“那么,一切,等后天见分晓吧。”苏芩站起身,认真的朝路胜鞠了一躬。“告辞。”

路胜目送他们五人鱼贯离开静室。心头有种莫名的怪异。

“如果按照这个苏芩所言,那岂不是就算我不降临,王木也会在闻达图书馆,无意间找到那本千锁定龙功,然后走上苦修格斗的道路?

还有,他们到底是怎么知道这么清楚的?或者说,难不成就像那个里提到的,我其实是一直生活在一部别人观看过的电影或者故事里?”

路胜心头一时间各种念头纷沓而至。

离开静室,天色已经有些黯淡了。

他没有回去,而是开着车,直奔图书馆。

闻达图书馆已经在他的管理下,重新调整,稍微翻修了下外边,看起来崭新多了。

路胜按照苏芩所说的,直接走上二楼,进了古籍存放的藏书室。

幽暗的藏书室内,一排排书架上稀稀落落的摆放着一本本价值不大的旧书。

所谓的古籍,其实最长也不过是四五十年前的书册。

路胜径直走向苏芩所言的书架,蹲下来翻了翻。很快便找到了她说的那本天机地理旧书。

“还真有!?”路胜微微一惊,看着手里的书册。

这是本标准的介绍地理和天文方面知识的旧书。路胜反过来看了下,出版社是倍加星综合新闻出版社。内容也就是些过时了的老版知识。

“放进油里浸泡?”

路胜取出手机,发了条指令。

然后他拿着书走出藏书室,来到一楼大厅。

大厅里已经有九命堂的人等着了,手里还提着一桶才买的食用油。

路胜接过油,进了图书馆的员工专用室。

在水槽里,他迅速将油往书册上一倒。

嘶

一阵细微的嘶嘶声响起,书册遇到油水仿佛像是化学反应一样,整个封面,书页,都完全大变样。

一开始淡黄色的书页,在浸泡了油水后,变得惨白惨白。

纸张的厚度也变薄了不少。

路胜拿起翻看了下,满是油水上的书页已经显露出了密密麻麻的大量字迹和图形。

字是几十年前还带有点点过度简化迹象的简化字。

魔图帝国原本使用的文字,因为太过简化,导致出现诸多重叠使用现象。于是帝国文化部重新制定了国家字典,将部分简化字再度恢复了繁体。

所以现在已经几乎看不到这种极度简化的字迹了。

路胜一眼就认出了这上边的字样。

“千锁定龙功?还真有?”

他心头一动,翻开大概扫视了一遍。

这功法,和他的螺旋九命法没得比,要差了好几个档次。但在开发自身潜能上,也算是首屈一指了。起码是远远超过这个星球所谓的魔图格斗术之类技艺。

和负荷格斗术差不多一个级别。

而且最重要的是,这门功法进展速度极快。

负荷格斗术从开始修习,到大成,安沙红刺利用电磁能和核能加速修行,也足足用了一年多。这还是在他的帮助下,不断进行自限性最大锻炼。

而这个千锁定龙功,按照上边记录的,只要能上手入门,就算最慢也只要三个月,就能大成。

“三个月??这是在搞笑?”路胜迅速翻开后续内容。

很快也明白了这功法的本质。

这是一本不折不扣的魔道功法。

如果说一开始还算循规蹈矩,正正经经,那么后续,就是各种以透支身体潜能为代价,不择手段的提升自身格斗术。

如果真的按照这上边的法子修炼,服药。功法是能练成,但寿命起码减寿四十年。

“这更像是某种创造出来实验性质的残本。”路胜合上功法,心头有了数。

“那个苏芩,能掌握这么多我不知道的情报。身上又有着我也没见识过的神秘力量气息。如果后天真的会发生她所言的那件事,或许”路胜心中若有所思。

蓝色星光交到路胜手里的一百人,素质参差不齐。

其中强的,直刚测试就达标了,根本不需要强化训练。而弱的,压根就是普通人,没有经过任何训练,甚至还有人以为自己是在拍电视剧电影。

路胜让白郡城下去一顿乱打后,这些乱叫的家伙立马安静了。

训练分成两部分,路胜想了想,没有教导他们负荷格斗术,而是直接改良了下,将千锁定龙功的弊端剪除部分,再简化后,丢给这群人开始习练。

时间缓缓流逝,转眼,两天过去了。

在苏芩预言的当天晚上,路胜下令严加监控的九命堂那人,居然真的和另外两人在ktv发生矛盾。

而且和苏芩所言的一模一样,九命堂弟子姚崇,这个在堂内也算好手的家伙,居然真的在冲突中被人打断了右手,然后回到分部求援。

如今的九命堂,雄霸整个阿诺丝省,甚至周边行省也无可奈何,被其影响力渐渐扩散。

这么一个庞然大物的好手,居然当场被打断右手。

当场分部的十几人,马上便起身打算去活动活动筋骨。

没想到打断姚崇手掌的那两人,居然主动找上门了。

“情况就是这样。对方使用的是念能师手段,攻击很凌厉干脆,显然不是第一次这么动手。”魏真鱼在手机里低声道。

路胜坐在沙发上,桌上放了一脸盆的蛋挞,他随手端起来喝一口。

十几个蛋挞像豆子一样哗啦全滚进他嘴里。

然后稍微咀嚼下,就全进了肚子。

“让白安义去看看。调查清楚是什么原因。”路胜放下脸盆,一旁的女弟子迅速上来用丝绢给他擦干净嘴。

“还有,把姚崇带来,还有那个打人的家伙,资料图像都传过来。”他吐了口气,拿起一边的矿泉水,丢进嘴里轻轻一咬,像吃果冻一样,一口气吃了十几瓶,算是解渴了。

如果苏芩所言的是真的,那么这个打人的家伙,很有可能是个关键人物。


A moment later, the group was seated in a private room at a nearby tea house.

After serving a pot of tea, the waiter quietly withdrew.

The door to the room closed.

Su Qin and Lu Sheng sat facing each other, while the other four young men and women, each with a distinct appearance, took seats behind her.

Unlike Su Qin, the other four appeared noticeably nervous, seemingly intimidated by Lu Sheng.

Lu Sheng, dressed in a black suit jacket, crossed his legs and took a sip from his teacup.

"You seem to have researched me. Tell me what you know."

He sensed that these individuals might possess information he hadn't anticipated.

Su Qin cleared her throat.

"Please dismiss our surroundings and check for any listening devices.

Moreover, what I'm about to reveal could be quite shocking. Please remain calm."

"Sister Qin!" a man behind her called out worriedly. "Let me handle this instead?"

"Don't worry, I know how to proceed," Su Qin said gravely, raising her hand.

Lu Sheng seemed unconcerned.

He pulled out his phone and sent a voice message. Within moments, the Nine Lives Hall experts guarding the entrance outside had all departed.

Then he watched as Su Qin produced a small, mirror-like device from her pocket and scanned the room with it.

"All clear. No additional devices," Su Qin said, letting out a sigh of relief.

"Speak. What's so serious?" Lu Sheng said calmly, setting down his teacup.

Su Qin bit her lip and took a deep breath.

"Hall Master, if I remember correctly, you only truly rose to prominence two years ago?"

"That's right. The Nine Lives Hall was also established then. You can find this information anywhere; it's hardly a secret," Lu Sheng replied dismissively.

"And the key to your rise was the Wenda Library, wasn't it?" Su Qin continued.

"In the library, you accidentally discovered the powerful Cultivation Manual, the Thousand Locks Dragon Art. This breakthrough allowed you to transcend human limits and embark on a path beyond the body's natural boundaries."

Su Qin paused for a moment before continuing.

"Initially, you slowly accumulated power by running a tutoring business. Later, you encountered Zheng Huan, your right-hand woman. After defeating her in a formal duel, you brought her into the fold. Together with your disciples, you began the true expansion of the Nine Lives Hall."

"If you don't believe me, the Thousand Locks Dragon Art was hidden on the second floor of the Wenda Library, on the bottom shelf of the second bookshelf in the ancient texts section.

The cover of the manual is disguised with the four characters 'Heavenly Secrets of Geography,' and it appears to contain only knowledge of astronomy and geography. But if you soak it in oil, its true contents will be revealed."

Lu Sheng had initially dismissed her words as mere storytelling. But when Su Qin provided extremely precise and detailed coordinates and descriptions, he hesitated.

"Now, do you believe us?" Su Qin finished her explanation in one breath, her expression remaining solemn.

"If that's all, you might say we're just well-informed. But next, I can prove that we can help you avoid the greatest catastrophe."

"What catastrophe?" Lu Sheng found his own expression growing solemn.

"Beijia Star will be destroyed in one year by an Odd Death Ray during the Interstellar War. The Psychic Alliance and Blue Starlight are fully at war, and their Manipulators are each Star-Destroying entities. Above the Manipulators, there are even stronger Dark Energy Users," Su Qin said gravely.

"If you still don't believe us, we can provide a small piece of private intelligence."

"What intelligence?" Lu Sheng found himself unconsciously taking Su Qin's words seriously.

"One of your disciples, a man named Yao Chong, will get into a drunken brawl the day after tomorrow and make an enemy. His right arm will be broken on the spot, and he'll return to you for help. At that time, you'll be able to judge whether our intelligence is true," Su Qin said confidently.

Lu Sheng fell silent for a moment.

"I understand. I'll carefully consider your words. Then..."

"Then, we'll see how everything unfolds two days from now," Su Qin said, standing up and bowing seriously to Lu Sheng. "Farewell."

Lu Sheng watched the five of them file out of the quiet room one by one. A strange feeling settled in his heart.

*If what Su Qin said is true, wouldn't Wang Mu have stumbled upon the Thousand Locks Dragon Art at Wenda Library even if I hadn't intervened? And wouldn't he have embarked on a path of arduous martial arts training anyway?*

*Also, how could they know all this so clearly? Or... could it be like that story mentioned—that I've actually been living in a movie or story that someone else has watched all along?*

A flurry of thoughts raced through Lu Sheng's mind.

When he left the quiet room, the sky had already darkened.

Instead of returning home, he drove straight to the library.

Under his management, Wenda Library had been renovated and its exterior spruced up, making it look much newer.

Following Su Qin's instructions, Lu Sheng went straight to the second floor and entered the rare book storage room.

In the dimly lit room, rows of shelves held scattered volumes of low-value old books.

The so-called rare books were actually no older than forty or fifty years.

Lu Sheng walked straight to the bookshelf Su Qin had mentioned and crouched down to search through it. He quickly found the old book she had described, **[** Heavenly Secrets of Geography **]**.

"It's actually here!?" Lu Sheng was slightly surprised, looking at the book in his hands.

It was a standard old book introducing geographical and astronomical knowledge. Lu Sheng turned it over to check the publisher: Beijia Star Comprehensive News Press. The content was just outdated, old-fashioned knowledge.

"Soak it in oil?"

Lu Sheng took out his phone and sent a command.

He took the book out of the library and went to the first-floor lobby.

Nine Lives Hall members were already waiting there, carrying a newly purchased bucket of cooking oil.

Lu Sheng took the oil and entered the library's staff-only room.

In the sink, he quickly poured the oil over the book.

*Hiss*

A faint hissing sound arose. The book reacted with the oil as if undergoing a chemical reaction, its entire cover and pages transforming completely.

The initially pale yellow pages turned deathly white after being soaked in the oil.

The paper also became significantly thinner.

Lu Sheng picked up the book and flipped through it. The oil-soaked pages now revealed dense clusters of characters and diagrams.

The characters were simplified characters, still showing faint signs of over-simplification from decades past.

The original script of the Demon Map Empire had become so simplified that it led to numerous overlapping usages. As a result, the Empire's Ministry of Culture revised the national dictionary, restoring some of the simplified characters to their traditional forms.

Consequently, such extremely simplified characters are rarely seen today.

Lu Sheng recognized the script at a glance.

"Thousand Locks Dragon Art? It actually exists?"

His heart stirred as he flipped through the pages, scanning them briefly.

This cultivation technique couldn't compare to his Spiral Nine Lives Method—it was several tiers below. Yet, in developing one's potential, it was still considered top-tier, far surpassing the planet's so-called Magic Diagram Martial Arts and similar techniques.

It was on par with the Load-Bearing Combat Art.

Most importantly, this technique progressed at an extremely rapid pace.

Even with Ansha Hongci using electromagnetic and nuclear energy to accelerate his training, it took over a year to master the Load-Bearing Combat Art from start to finish. This was even with Lu Sheng's assistance, constantly pushing him through self-imposed maximum exertion.

According to the records, the Thousand Locks Dragon Art only requires three months to master, even for the slowest learners.

"Three months? Is this a joke?" Lu Sheng quickly flipped through the subsequent sections.

He soon understood the true nature of this Cultivation Technique.

It was an out-and-out demonic art.

While the initial stages seemed relatively conventional and proper, the later stages involved various methods of ruthlessly enhancing one's combat skills by draining the body's potential.

If one truly followed the methods described in this manual and took the prescribed medicines, they could indeed master the technique, but it would shave at least forty years off their lifespan.

"This feels more like some experimental fragment than a complete manual," Lu Sheng thought as he closed the manual, his mind made up.

"That Su Qin knows so much that I don't. She also carries a mysterious aura that I've never encountered before. If what she said about tomorrow actually happens, then perhaps..." Lu Sheng mused, his mind deep in thought.

The hundred individuals Blue Starlight had entrusted to Lu Sheng were of wildly varying quality.

The stronger candidates passed the initial assessment without needing any further training. The weaker ones were essentially ordinary people with no training at all; some even thought they were filming a TV show or movie.

After Lu Sheng ordered Bai Juncheng to beat them up thoroughly, the noisy rabble quickly fell silent.

The training was divided into two parts. After some thought, Lu Sheng decided not to teach them the Load-Bearing Combat Art. Instead, he modified the Thousand Locks Dragon Art, removing some of its drawbacks and simplifying it before teaching it to the group.

Time passed slowly. In the blink of an eye, two days had gone by.

On the night of Su Qin's predicted event, the Nine Lives Hall member Lu Sheng had ordered to be closely monitored actually got into an altercation with two others at a karaoke bar.

And just as Su Qin had foretold, Yao Chong, a skilled disciple of the Nine Lives Hall, had his right arm broken during the conflict and returned to the branch for help.

The Nine Lives Hall now dominated the entire Anos Province, and its influence was gradually spreading to neighboring provinces, which were powerless to stop it.

For a top fighter from such a colossal power to have his right arm broken on the spot was truly shocking.

The dozen or so members of the branch immediately rose, ready to stretch their limbs.

Unexpectedly, the two who had interrupted Yao Chong's hand came to them directly.

"That's the situation. The attackers used Nen Master techniques, their strikes were sharp and decisive. They've clearly done this before," Wei Zhenyu whispered into her phone.

Lu Sheng sat on the sofa, a basin full of egg tarts on the table. He casually picked it up and took a sip.

A dozen egg tarts spilled into his mouth like beans.

After a quick chew, they were all swallowed.

"Send Bai Anyi to investigate the cause," Lu Sheng said, setting the basin down. A female disciple quickly stepped forward to wipe his mouth with a silk handkerchief.

"Also, bring Yao Chong and the assailant here. Send me their files and images." He sighed, picked up a bottle of mineral water, and bit it gently. He downed a dozen bottles in one go, like eating jelly, finally quenching his thirst.

If what Su Qin said was true, that assailant might be a key figure.


Here's the next chapter I'd like to be translated in the same style then turned into an html page:

    一道道或飞行，或疾驰的人影，不断从混乱的城市里穿行，朝着路胜所在方向赶去。

    整个安明市已经有三分之一，毁于混乱中的厮杀缠斗，轮回者之间的死斗，和九命堂干部们的清场，都不断加大整个城市的破坏度。

    但不管如何，处于人道主义，和这里是自己总部地盘的意识。九命堂的干部和弟子们，在肃清丧尸和轮回者的同时，也在不断的尽量营救还活着的人。

    他们的营救，与其说是营救，不如说只是帮助解除威胁。

    从九命堂总部开始，诸多干部带领成员，逐步朝着城市各个角落步步推进，所有安全区域，由后续赶到的警部残存人员和军方留下的少部分部队一起协同防守。

    一些被救出的富豪们惊魂未定后，挥动着支票，要求雇佣九命堂为他们救出亲人妻女。

    但九命堂一概不理。

    情况爆发得太突然，以至于很多上层人士根本没有做好任何准备。

    病毒的扩散只维持了一小时，基本上整个安明市完全陷入了感染。不过根据统计，这种病毒似乎能够免疫的人不少，很多体质强健的正常人，都能安全度过这次危机。但最可怜的还是小孩。

    他们体质衰弱，被病毒感染扩散得最快。

    这次整个安明市的孩童，十岁以下的几乎无一幸免。

    各种建筑物小半出现不同程度损毁，幸存者们大多在九命堂的引领下，经过一个个的身体筛查后，确认没有伤口，才被允许进入防线后方。

    很快，两个小时后，城市一半的面积被彻底肃清完毕。只剩下最后的重灾区。

    丧尸最集中和轮回者最多的的中心商贸区。

    ***************

    ***************

    路胜静静站在一栋三层小别墅的屋顶，身后足足五米多宽的肉翼，缓缓扇动着，不时带起阵阵劲风。

    下方已经有一道道认识的，不认识的人影出现在废墟上，遥遥朝着他单膝跪地。

    灭情种不需要路胜主动控制，这完全是被动激发，只要是被他击败而不死的，都会被强制破开心灵防线，在意识最深处种下灭情种。

    理论上，只有意志极其坚定的类型，才能抵抗灭情种的植入。但迄今为止，路胜还没见过意志力能坚韧到那个程度的人。

    人数慢慢增多，逐步从个位数，增长到了两位数。

    很快，两个关键人物出现了。

    之前还在和人缠斗，制造了大屠杀恶魔之门的布莱特，此时正颤抖着身体，用一种惊恐的眼神盯着路胜，双脚却不自觉的朝着路胜缓缓走近。

    他一边恐惧一边又不得不听从路胜的命令。

    另一人，则是之前从路胜手下逃离了的白颂碟。

    她双臂鲜血淋漓，似乎才经过某场厮杀，此时脸色泛白，楚楚可怜，就像是受伤的小兽，急需别人的安慰和怀抱。

    路胜并不知道病毒就是白颂碟下的，此时只是眼神一扫，便淡淡划过去。

    白颂碟原本惊恐无比的身体，这才缓缓平息下来。她最害怕的就是路胜知道是她下的病毒。

    灭情种的弊端此时也显露出来了。

    灭情种产生的恐惧，会让人不自觉的敬畏和听从路胜的指令，但这种恐惧，也可能会导致别人为了避免更恐惧，更害怕的结果，从而隐瞒一些关键。

    又等了一会儿，周围陆陆续续居然又走出来好几人，全都是轮回者。

    再等便没有人抵达了。

    路胜这才缓缓开口。

    “不论你们来自何处。不论你们身负什么任务。现在，你们都是我的仆人。”

    他顿了顿，目光在布莱特身上停顿。

    嘭！！！

    布莱特还没来得及说话。便感觉额头一阵剧痛，随即整个人眼前一黑，然后便什么也不知道了。

    其余人只看到布莱特整个人瞬间爆开，化为一团血雾，转眼便被风吹散，消失在天地间。

    只有几个看起来十分精致的黑盒子掉落下来。

    “其余人，去肃清现在的残留城区，记住，换上九命堂的特殊标志。”路胜平淡道。

    下方众人纷纷行礼，然后各自沉默不语的朝着远处狂奔离开。

    白颂碟也是一样，心头一阵庆幸，正要转身飞离。

    “你留下。”路胜的声音忽然在她耳边响起。

    白颂碟浑身一颤，几乎就要夺路而逃，但灭情种产生的对路胜的惊恐和敬畏，又硬生生的让她停在原地。

    “你比其他人的心跳，每分钟快了十二下。”路胜身形带出一串残影，转瞬便到了她身后。

    “告诉我，你们的人，这次降临的最强者，是谁？刚才那个老头？”路胜试图从白颂碟口中弄清楚一些轮回者的底细。

    “我不”白颂碟哭丧着脸，缓缓转过身。

    “病毒是谁放的？”路胜忽然问。

    “我我不是我！！啊！！！”白颂碟猛地惊恐万分，连连后退，情绪大变。

    路胜看到这一幕，已经知道答案了。

    “还有么？拿出来。”他平静道。

    白颂碟神色惊惶，后退数步，脸上闪过迟疑，犹豫，煎熬等各种表情。最后，她还是伸出手，掌心一闪之下，多出了两支密封的，装着银色旋转液体的试管。

    路胜伸手接过试管，抬头对着光查看了下。

    “走吧，把你知道的关键人物，都说出来。”他一把抓住白颂碟长发，拖着她纵身朝远处飞去。

    商贸区，数座商场最集中的街区。

    黑压压的丧尸群不断在这附近游走，一个接一个，宛如生前一样在街道上闲逛。

    他们的皮肤已经溃烂，身上不少地方能看到腐烂的血肉。很多器官组织，除开消化系统外其余都已经彻底报废。

    如果从高空往下看，就能看到，整个这附近数百米范围，全是密密麻麻无以计数的丧尸群。

    铛。

    忽然一处五金店门前，传出清脆刺耳的巨大响声。

    大量丧尸纷纷嘶吼着，抬头朝着声音传出的方向缓缓走去。

    铛铛铛铛！！

    忽然一阵更加激烈的金属敲击声接连响起。

    这下，更多的丧尸彻底确定了方位，加快速度朝着声音传出方向赶去。

    一群群丧尸纷纷汇集一起，组成更大的群体，朝着声音来源处走去。

    安沙手里不断盘旋着短刀，一下下的敲击在五金店门前地上的几根金属管上。

    他注视着前面慢慢逼近的黑压压尸潮，神情冷漠而平淡。

    “所以说，这种毫无意义的垃圾，再来多少，也只是垃圾啊”

    很快，丧尸聚集得够多了。

    他停下敲击，竖起短刀，刀背对准自身，屈指一弹。

    嘣！！！

    刹那间整个短刀彻底崩散。化为无数金属碎片。

    碎片再度炸碎成更细小的银色砂砾，彷如一团银雾长龙，狠狠吹进丧尸潮。

    呼

    细微的噗噗声连续不断响起。

    顷刻间整个尸潮金属僵直不动。

    安沙转过身，看也不看身后，朝着远处离去。

    身后是大片不断倒地碎成烂肉的无数丧尸。

    另一处大商场边缘。整个三角形商城，一半的楼体，正缓缓往下垮塌，溅起大量灰黄色烟尘。

    嘭！！！

    完好的一半处，大块的楼层陡然倒塌炸碎。

    烟尘中，红刺高达四米多的庞大身躯，缓缓显露出来。轻轻一跃落地，站在大片废墟中。

    她整个人皮肤上浮现出一根根血管一样的绿色脉络，面部的绿色管道高速输送着海量绿色液体。

    一身恐怖的肌肉几乎到了畸形的地步，像是无数肿瘤堆砌到一起，又像是巨树地下胡乱生长的树根树瘤。一些肌肉块甚至比她整个脑袋还要大。

    “滚出来！一群贱人！”红刺走在一片废墟的大楼砂石上。声音咆哮得地面都在震动。

    啊！

    几个躲在断墙缝隙里的男女被惊动，狼狈的惊恐钻出来，朝着远处逃去。

    红刺看也不看，只是环顾四周，寻找刚才和她交手的几个轮回者。

    轮回者，师尊王木就是这么告诉他们的。她不知道是什么含义，不过不管什么含义，反正都是些死人，知不知道也无所谓。

    “区区np！！”一个阴测测的声音忽然响起。

    同一时间，一道灰色人影速度奇快的朝着红刺飞扑过去。

    另外两道灰色人影各自远远站立，从烟尘中显出身躯，他们手中持有类似弩箭一样的复杂机器，此时已经对准红刺。

    “雷龙！”

    轰！！

    两道弩箭离弦而出，在半空中居然瞬间变成两条蓝色电光长龙。

    “找死！”

    红刺身体弓起，掌心急速震动起来，汇聚出大片暗绿色光点。

    “冥星高翔！！”

    她猛地往地面一掌打出，大片墨绿光点咆哮着化为绿光飞出，钻入大地。

    顷刻间静寂了一瞬。

    轰！！

    地面大片水泥猛然炸开。无数石块炮弹般朝四面八方爆射。瞬间便将雷龙和靠近的那道灰色人影打得吐血倒飞出去。

    魏韩冬一身黑色皮西装，露出的强壮胸膛戴着一个妹妹送给他的黑玉吊坠，形状是个奇怪的倒月牙。

    他手里拖着一个浑身是血的蒙面女子，拖的地方是女子的腿，虽然女子看起来有几分姿色，但魏韩冬除了自己妹妹外，对其他女人完全没兴趣。

    所以在这女人对他展开美色魅惑的瞬间，他上去就是一记黑虎掏心，胜负就彻底确定了。

    “小莲！！啊啊啊啊！我杀了你！！”

    一道黑色人影从后方一个闪现跳出，手里的重型战斧从上往下带着白光，狠狠斩在魏韩冬脑门上。

    铛！！！！

    一阵敲钟一般的巨响猛地炸开。

    魏韩冬脚下一片龟裂，但他只是无动于衷的看着面前这人。

    因为这样的情景已经不是第一次了。

    这人掌握着瞬移一样的技巧，他抓不到他，但这人也砍不动他。

    索性他也懒得动手了，让他砍。
`)

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-3-flash-preview",
		genai.Text(input),
		nil,
	)

	if err != nil {
		tries := 1

		fmt.Println("Retry:", tries)
		for tries != 5 {
			result, err = client.Models.GenerateContent(
				ctx,
				"gemini-3-flash-preview",
				genai.Text(input),
				nil,
			)

			if err == nil {
				break
			}
			tries += 1
		}

		log.Fatal(err)
	}
	response := result.Text()
	//fmt.Println(response)

	/*html_start_idx := strings.Index(response, "```html")

	if html_start_idx == -1 {
		log.Fatal("No HTML block found in the response")
	}*/

	_, html_start, found := strings.Cut(response, "```html")
	if !found {
		log.Fatal("No HTML block found in the response")
	}

	html_content, _, found := strings.Cut(html_start, "```")
	if !found {
		log.Fatal("End of HTML block not found in the response. Possibly malformed.")
	}

	fmt.Println(html_content)
}
