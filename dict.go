package randname

import (
	"math/rand"
	"time"
)

type (
	words []rune
)

var (
	names = []words{
		words{'王', '李', '张', '刘', '陈', '杨', '黄', '吴', '赵', '周', '徐', '孙', '马', '朱', '胡', '林', '郭', '何', '高', '罗', '郑', '梁', '谢', '宋', '唐', '许', '邓', '冯', '韩', '曹', '曾', '彭', '肖', '蔡', '潘', '田', '董', '袁', '于', '余', '蒋', '叶', '杜', '苏', '魏', '程', '吕', '丁', '沈', '任', '姚', '卢', '钟', '姜', '崔', '谭', '廖', '范', '汪', '陆', '金', '石', '戴', '贾', '韦', '夏', '邱', '方', '侯', '邹', '熊', '孟', '秦', '白', '毛', '江', '闫', '薛', '尹', '付', '段', '雷', '黎', '史', '龙', '钱', '贺', '陶', '顾', '龚', '郝', '邵', '万', '严', '洪', '赖', '武', '傅', '莫', '孔', '汤', '向', '常', '温', '康', '施', '文', '牛', '樊', '葛', '邢', '安', '齐', '易', '乔', '伍', '庞', '颜', '倪', '庄', '聂', '章', '鲁', '岳', '翟', '申', '殷', '詹', '欧', '耿', '关', '覃', '兰', '焦', '俞', '左', '柳', '甘', '祝', '包', '代', '宁', '符', '阮', '尚', '舒', '纪', '柯', '梅', '童', '毕', '凌', '单', '季', '成', '霍', '苗', '裴', '涂', '谷', '曲', '盛', '冉', '翁', '蓝', '骆', '路', '游', '靳', '辛'},
		words{'静', '敏', '丹', '丽', '燕', '婷', '娟', '艳', '芳', '娜', '玲', '洋', '洁', '佳', '琳', '雪', '阳', '颖', '莉', '慧', '萍', '倩', '宁', '莹', '霞', '云', '晶', '欣', '晨', '欢', '乐', '璐', '梅', '瑞', '琪', '琦', '岩', '海', '琼', '迪', '聪', '菲', '珊', '薇', '蓉', '凡', '茜', '悦', '雯', '扬', '睿', '瑜', '萌', '妍', '群', '娇', '玮', '虹', '凤', '珍', '梦', '越', '枫', '怡', '蕊', '畅', '芬', '靖', '惠', '硕', '柳', '晴', '莎', '曦', '昕', '艺', '烨', '婕', '朋', '涵', '瑾', '灿', '芸', '岚', '钰', '嘉', '曼', '捷', '彤', '卓', '珂', '彦', '真', '菁', '培', '容', '瑛', '馨', '灵', '雅', '鸿', '园', '凌', '翠', '琛', '煜', '锦', '勤', '霖', '秀', '蓓', '香', '寒', '言', '斐', '奕', '巧', '焱', '莲', '蔚', '姗', '姣', '雁', '舒', '娅', '韵', '琰', '桐', '昱', '盈', '婵', '娴', '竹', '臻', '媚', '霄', '仪', '榕', '贞', '依', '珠', '滢', '焕', '婉', '沁', '弘', '芝', '语', '淇', '诗', '茗', '宜', '彩', '宣', '素', '毓', '碧', '姿'},
		words{'小', '晓', '佳', '海', '嘉', '云', '雪', '新', '瑞', '亚', '丽', '梦', '雅', '秀', '欣', '锦', '梓', '静', '艳', '若', '慧', '彦', '艺', '书', '诗', '淑', '晨', '凤', '燕', '怡', '惠', '乐', '钰', '利', '敏', '奕', '悦', '芳', '鸿', '丹', '培', '彩', '颖', '馨', '琳', '舒', '雯', '睿', '宜', '婉', '香', '亦', '依', '韵', '如', '玲', '凡', '素', '巧', '莉', '洁', '育', '靖', '宁', '楚', '卓', '碧', '翠', '涵', '倩', '芝', '珍', '灵', '语', '凌', '芸', '妍', '贞', '笑', '玮', '煜', '毓', '婷', '琦', '琼', '晶', '仪', '茗', '昕', '昱', '真', '曼', '梅', '虹', '茜', '璐', '雁', '菲', '娅', '聪', '寒', '沛', '竹', '欢', '薇', '莲', '言', '柔', '瑾', '勤', '蓉', '琪', '娜', '泓', '莹', '萍', '瑜', '焕', '蕊', '霞', '娟', '弘', '晋', '柳', '枫', '京', '阳', '珊', '姿', '滢', '菁', '群', '岚', '芬', '映', '霄', '园', '沁', '朋', '婕', '蓓', '岩', '苑', '琰', '意', '影', '莎', '荷', '知', '泳', '妙', '珠', '迎', '迪', '忆', '婵', '念', '娴', '姣', '宣', '品', '盈', '烨', '荔', '芯', '姗', '越', '歆', '灿', '珂', '榕', '曦', '扬', '洋', '琛', '容', '轶', '焱', '晴', '蕴'},
	}

	pinyin = map[rune]string{
		20399: "hou", 32736: "cui", 26460: "du", 20613: "fu", 23092: "xian", 32599: "luo", 28504: "pan", 27784: "chen", 31069: "zhu", 38669: "huo", 38451: "yang", 32874: "cong", 23481: "rong", 29020: "yu", 37045: "shao", 30003: "shen", 38742: "jing", 38470: "lu", 37049: "zou", 24247: "kang", 32426: "ji", 23452: "yi", 23002: "yao", 37329: "jin", 20052: "qiao", 27915: "yang", 31243: "cheng", 40654: "li", 23068: "na", 29840: "lu", 24464: "xu", 38639: "wen", 24433: "ying", 26216: "chen", 33778: "fei", 21220: "qin", 33275: "zhen", 33509: "ruo", 26131: "yi", 38678: "lin", 24344: "hong", 20313: "yu", 24278: "liao", 36158: "jia", 32831: "geng", 22925: "yan", 26133: "xin", 26032: "xin", 20140: "jing", 40784: "qi", 26539: "feng", 20202: "yi", 32918: "xiao", 26757: "mei", 38597: "ya", 37041: "qiu", 29618: "ling", 24800: "hui", 26228: "qing", 26195: "xiao", 27462: "xin", 33673: "li", 29788: "yu", 29066: "xiong", 34203: "xue", 29141: "yan", 24935: "hui", 30021: "chang", 38047: "zhong", 20522: "ni", 29822: "jin", 34164: "yun", 35885: "tan", 36842: "di", 29787: "ying", 26446: "li", 33931: "jiang", 25140: "dai", 29579: "wang", 35895: "gu", 29664: "zhu", 26187: "jin", 37011: "deng", 24120: "chang", 23731: "yue", 33495: "miao", 39062: "ying", 26790: "meng", 24352: "zhang", 26041: "fang", 23380: "kong", 32735: "di", 29738: "qi", 29642: "shan", 21016: "liu", 31456: "zhang", 38745: "jing", 33452: "fen", 24039: "qiao", 20122: "ya", 26970: "chu", 19969: "ding", 23004: "jiang", 22799: "xia", 20961: "fan", 39321: "xiang", 23011: "jiao", 39038: "gu", 27605: "bi", 20113: "yun", 33678: "sha", 23706: "lan", 38901: "yun", 35799: "shi", 33620: "li", 23433: "an", 29614: "wei", 28977: "yan", 23113: "wan", 37073: "zheng", 26366: "ceng", 20005: "yan", 26045: "shi", 33459: "fang", 23506: "han", 40644: "huang", 20446: "yu", 28949: "huan", 27859: "hong", 21346: "lu", 28034: "tu", 28904: "ye", 30408: "ying", 21033: "li", 20134: "yi", 30693: "zhi", 27754: "wang", 21490: "shi", 32834: "nie", 27431: "ou", 33564: "qian", 20964: "feng", 29723: "chen", 33489: "yuan", 27891: "yong", 23425: "ning", 39558: "luo", 23047: "jiao", 25463: "jie", 28789: "ling", 36213: "zhao", 33891: "dong", 33883: "ge", 34013: "lan", 27905: "jie", 38686: "xia", 25196: "yang", 34003: "bei", 40857: "long", 38064: "yu", 21331: "zhuo", 26144: "ying", 24609: "yi", 26000: "fei", 24425: "cai", 21776: "tang", 37085: "hao", 30427: "sheng", 29747: "lin", 30805: "shuo", 23157: "chan", 28386: "ying", 27803: "pei", 26417: "zhu", 35874: "xie", 29976: "gan", 20873: "ran", 38182: "jin", 24429: "peng", 23828: "cui", 31206: "qin", 19975: "wan", 21521: "xiang", 23578: "shang", 33298: "shu", 25104: "cheng", 36763: "xin", 24742: "yue", 28085: "han", 33464: "yun", 34074: "wei", 21697: "pin", 26361: "cao", 38634: "xue", 34183: "wei", 38593: "yan", 23194: "mei", 26580: "rou", 39640: "gao", 27946: "hong", 26607: "ke", 38771: "jin", 33437: "zhi", 23391: "meng", 20184: "fu", 37026: "xing", 24196: "zhuang", 24038: "zuo", 32676: "qun", 34425: "hong", 31481: "zhu", 27777: "qin", 32993: "hu", 26753: "liang", 30000: "tian", 27494: "wu", 35203: "tan", 35060: "pei", 29756: "qiong", 26342: "xi", 26771: "zi", 28113: "shu", 36182: "lai", 23159: "ting", 33993: "rong", 23125: "jie", 29634: "ke", 31168: "xiu", 27603: "yu", 22914: "ru", 39532: "ma", 20309: "he", 28966: "jiao", 32705: "weng", 29645: "zhen", 24518: "yi", 33487: "su", 20851: "guan", 20521: "qian", 24422: "yan", 33559: "ming", 40858: "gong", 36335: "lu", 33805: "ping", 23435: "song", 33707: "mo", 21333: "dan", 25935: "min", 20029: "li", 29734: "qi", 22025: "jia", 33729: "jing", 35328: "yan", 29744: "yan", 20237: "wu", 33804: "meng", 22999: "shan", 23567: "xiao", 34945: "yuan", 31526: "fu", 26354: "qu", 33721: "ying", 38472: "chen", 21556: "wu", 20911: "feng", 38889: "han", 39759: "wei", 24222: "pang", 21253: "bao", 33395: "yan", 20339: "jia", 27426: "huan", 29790: "rui", 28023: "hai", 26364: "man", 33714: "lian", 27029: "rong", 26472: "yang", 35768: "xu", 20110: "yu", 33539: "fan", 26611: "liu", 26230: "jing", 36234: "yue", 20070: "shu", 36814: "ying", 21525: "lv", 36154: "he", 27748: "tang", 28216: "you", 22253: "yuan", 23045: "ya", 31505: "xiao", 21494: "ye", 23609: "yin", 20195: "dai", 36126: "zhen", 20381: "yi", 20940: "ling", 34081: "cai", 33402: "yi", 26379: "peng", 28799: "can", 30495: "zhen", 22521: "pei", 35821: "yu", 36726: "yi", 20219: "ren", 40065: "lu", 20848: "lan", 26704: "tong", 38660: "xiao", 30707: "shi", 30333: "bai", 27575: "yin", 32946: "yu", 33655: "he", 27743: "jiang", 23385: "sun", 27611: "mao", 27573: "duan", 29275: "niu", 23071: "juan", 30591: "rui", 39336: "xin", 23459: "xuan", 32032: "su", 26519: "lin", 31461: "tong", 24420: "tong", 24847: "yi", 21608: "zhou", 25991: "wen", 27146: "fan", 35449: "zhan", 38446: "ruan", 27427: "xin", 40511: "hong", 30887: "bi", 33455: "xin", 20025: "dan", 23721: "yan", 22869: "yi", 24565: "nian", 37101: "guo", 38886: "wei", 38647: "lei", 38518: "tao", 39068: "yan", 20048: "le", 34122: "rui", 26161: "yu", 38379: "yan", 28201: "wen", 28103: "qi", 38065: "qian", 23395: "ji", 23039: "zi", 22937: "miao",
	}

	randomBytes = []byte("abcdefghjkmnpqrstuvwxyz123456789")
)

func (w words) Get() rune {
	return w[rand.Intn(len(w)-1)]
}

func RandomBytes(min, max int, special bool) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63()))
	size := min + r.Intn(max-min)
	b := make([]byte, size)
	r.Read(b[:])
	for i, x := range b {
		b[i] = randomBytes[x&31]
	}
	if special {
		var lx = 0
		for {
			if lx > 2 {
				break
			}
			l := r.Intn(size - 1)
			if 'a' <= b[l] && b[l] <= 'z' {
				b[l] = randomBytes[r.Intn(len(randomBytes)-1)]
				lx++
			}
		}
		lx = 0
		for l, x := range b {
			if l > 2 {
				break
			}
			if 'a' <= x && x <= 'z' {
				b[l] -= 'a' - 'A'
				lx++
			}
		}
	}
	return string(b)
}
