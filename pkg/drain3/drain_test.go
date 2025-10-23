package drain3

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDrain(t *testing.T) {
	drain, err := NewDrain(WithExtraDelimiter([]string{"_"}))
	require.NoError(t, err)

	miner := NewTemplateMiner(drain, NewMemoryPersistence())

	logs := []string{
		"[ProducerStateManager partition=__consumer_offsets-48] Writing producer snapshot at offset 4339939698 (kafka.log.ProducerStateManager)",
		"[Log partition=__consumer_offsets-48, dir=/home1/irteam/apps/data/kafka/kafka-logs] Rolled new log segment at offset 4339939698 in 3 ms. (kafka.log.Log)",
		"[Log partition=__consumer_offsets-48, dir=/home1/irteam/apps/data/kafka/kafka-logs] Deleting segment files LogSegment(baseOffset=0, size=0, lastModifiedTime=1645674584000, largestRecordTimestamp=None) (kafka.log.Log)",
		"Deleted log /home1/irteam/apps/data/kafka/kafka-logs/__consumer_offsets-48/00000000000000000000.log.deleted. (kafka.log.LogSegment)",
		"Deleted offset index /home1/irteam/apps/data/kafka/kafka-logs/__consumer_offsets-48/00000000000000000000.index.deleted. (kafka.log.LogSegment)",
		"Deleted time index /home1/irteam/apps/data/kafka/kafka-logs/__consumer_offsets-48/00000000000000000000.timeindex.deleted. (kafka.log.LogSegment)",
		"[Log partition=__consumer_offsets-48, dir=/home1/irteam/apps/data/kafka/kafka-logs] Deleting segment files LogSegment(baseOffset=2147429227, size=0, lastModifiedTime=1710735195000, largestRecordTimestamp=None) (kafka.log.Log)",
		"Deleted log /home1/irteam/apps/data/kafka/kafka-logs/__consumer_offsets-48/00000000002147429227.log.deleted. (kafka.log.LogSegment)",
		"Deleted offset index /home1/irteam/apps/data/kafka/kafka-logs/__consumer_offsets-48/00000000002147429227.index.deleted. (kafka.log.LogSegment)",
		"Deleted time index /home1/irteam/apps/data/kafka/kafka-logs/__consumer_offsets-48/00000000002147429227.timeindex.deleted. (kafka.log.LogSegment)",
		"[ProducerStateManager partition=__consumer_offsets-49] Writing producer snapshot at offset 4339698 (kafka.log.ProducerStateManager)",
		"[Log partition=__consumer_offsets-48, dir=/home1/irteam/apps/data/kafka/kafka-logs] Deleting segment files LogSegment(baseOffset=4294790577, size=2703, lastModifiedTime=1711832815000, largestRecordTimestamp=Some(1710827112244)) (kafka.log.Log)",
		"[Log partition=__consumer_offsets-48, dir=/home1/irteam/apps/data/kafka/kafka-logs] Deleting segment files LogSegment(baseOffset=4338631022, size=641, lastModifiedTime=1711849197000, largestRecordTimestamp=Some(1711849197921)) (kafka.log.Log)",
		"Deleted log /home1/irteam/apps/data/kafka/kafka-logs/__consumer_offsets-48/00000000004294790577.log.deleted. (kafka.log.LogSegment)",
		"Deleted log /home1/irteam/apps/data/kafka/kafka-logs/__consumer_offsets-48/00000000004338631022.log.deleted. (kafka.log.LogSegment)",
		"Deleted offset index /home1/irteam/apps/data/kafka/kafka-logs/__consumer_offsets-48/00000000004294790577.index.deleted. (kafka.log.LogSegment)",
		"Deleted offset index /home1/irteam/apps/data/kafka/kafka-logs/__consumer_offsets-48/00000000004338631022.index.deleted. (kafka.log.LogSegment)",
		"Deleted time index /home1/irteam/apps/data/kafka/kafka-logs/__consumer_offsets-48/00000000004294790577.timeindex.deleted. (kafka.log.LogSegment)",
		"Deleted time index /home1/irteam/apps/data/kafka/kafka-logs/__consumer_offsets-48/00000000004338631022.timeindex.deleted. (kafka.log.LogSegment)",
		"[Log partition=__consumer_offsets-48, dir=/home1/irteam/apps/data/kafka/kafka-logs] Deleting segment files LogSegment(baseOffset=4339285360, size=104857589, lastModifiedTime=1711865580000, largestRecordTimestamp=Some(1711865580112)) (kafka.log.Log)",
		"Deleted log /home1/irteam/apps/data/kafka/kafka-logs/__consumer_offsets-48/00000000004339285360.log.deleted. (kafka.log.LogSegment)",
		"Deleted offset index /home1/irteam/apps/data/kafka/kafka-logs/__consumer_offsets-48/00000000004339285360.index.deleted. (kafka.log.LogSegment)",
		"Deleted time index /home1/irteam/apps/data/kafka/kafka-logs/__consumer_offsets-48/00000000004339285360.timeindex.deleted. (kafka.log.LogSegment)",
		"[Log partition=__consumer_offsets-49, dir=/home1/irteam/apps/data/kafka/kafka-logs] Rolled new log segment at offset 432939698 in 2 ms. (kafka.log.Log)",
	}

	ctx := context.Background()
	for _, log := range logs {
		_, _, template, _, err := miner.AddLogMessage(ctx, log)
		require.NoError(t, err)

		params := miner.ExtractParameters(template, log)
		require.NotNil(t, params)
	}

	clusters := miner.drain.GetClusters()
	require.Equal(t, 5, len(clusters))
}

func TestDeleteCluster(t *testing.T) {
	t.Run("删除存在的聚类", func(t *testing.T) {
		// 创建Drain实例
		drain, err := NewDrain()
		require.NoError(t, err)

		// 添加一些日志消息以创建聚类
		testLogs := []string{
			"2025-10-22 16:37:36.096 [] [pool-14-thread-61] ERROR PeSpeakerServiceImpl 276 - voice synthesis failed！ request param:{volume=100, appId=100024, format=mp3, speakerId=106, secret=116af9b65e32b962f7db8fe31ffa18a7, isTry=false, userId=10768422, content=<speak rate=\"0\" pitch=\"0\" volume=\"100\">这孩子打小没爹没娘,靠着给村里人放牛换口饭吃,</speak>, rawContent=这孩子打小没爹没娘,靠着给村里人放牛换口饭吃,}, request url:http://peiyinsvc-inc.xlgogo.com/detailed/voice/synthesis, response:{\"code\":1000,\"msg\":\"系统错误\"}",
			"2025-10-22 15:10:55.575 [8a3c607b037a488f94114a2bfc325add] [http-nio-10024-exec-14] ERROR PeSpeakerServiceImpl 178 - voice synthesis failed！ request param:{role=null, appId=100024, format=mp3, speakerId=111, style=null, secret=116af9b65e32b962f7db8fe31ffa18a7, isTry=false, userId=10491486, content=<speak rate=\"0\" pitch=\"0\" volume=\"100\"  ># 无忧健康·水元万相毕节市营运中心工作规划方案\n## 一、规划背景\n无忧健康·水元万相毕节市营运中心已明确各区域扶持负责人及主要工作职责，为确保各项工作有序开展、高效推进，实现营运中心的整体目标，特制定本工作规划。\n\n## 二、工作目标\n1. 短期目标（1 - 3 个月）\n- 各区域扶持工作全面启动，与当地相关机构、市场建立初步联系。\n- 商学院完成至少 2 场沙龙会的组织与开展。\n- 后勤保障部门完成 50 台水机的安装与报单工作。\n- 市场部在各区域开展至少 3 次市场推广活动。\n- 健康小屋完成对 2 个区县的机组健康小屋培训扶持。\n2. 中期目标（4 - 6 个月）\n- 各区域市场得到有效拓展，客户数量显著增加。\n- 商学院沙龙会覆盖毕节市所有县区市，每场参与人数达到 50 人以上。\n- 后勤保障工作稳定高效，水机安装与报单准确率达到 98%以上。\n- 市场部建立完善的市场推广体系，市场占有率提升 10%。\n- 健康小屋在毕节市各区县市全面铺开，培训扶持工作取得显著成效。\n3. 长期目标（7 - 12 个月）\n- 无忧健康·水元万相品牌在毕节市深入人心，成为当地健康水行业的领军品牌。\n- 各区域市场实现可持续发展，业务规模不断扩大。\n- 建立健全的运营管理体系，提高整体运营效率和服务质量。\n\n## 三、工作安排\n### （一）公关与宣传工作（刘涛 - 公关宣传部）\n1. 第一个月\n- 制定毕节市公关宣传整体方案，明确宣传目标、渠道和策略。\n- 与毕节市当地的媒体、行业协会、政府部门等建立联系，为后续宣传工作奠定基础。\n- 收集整理无忧健康·水元万相的产品资料、品牌故事等宣传素材。\n2. 第二 - 三个月\n- 开展媒体宣传活动，在毕节市主流媒体、行业网站、社交媒体等平台发布品牌宣传文章、图片和视频。\n- 组织参加毕节市相关的行业展会、研讨会等活动，提升品牌知名度和影响力。\n- 策划并执行至少 2 次公关活动，如新闻发布会、公益活动等，吸引社会关注。\n3. 第四 - 六个月\n- 持续跟踪媒体宣传效果，根据反馈及时调整宣传策略。\n- 建立品牌舆情监测机制，及时处理负面舆情，维护品牌形象。\n- 与毕节市当地的网红、意见领袖等合作，开展品牌推广活动，扩大品牌传播范围。\n4. 第七 - 十二个月\n- 总结公关宣传工作经验，不断优化宣传方案。\n- 加强与毕节市当地媒体的长期合作，建立稳定的宣传渠道。\n- 开展品牌形象升级活动，提升品牌的美誉度和忠诚度。\n\n### （二）日常运营工作（吕志谋 - 日常所有工作）\n1. 第一个月\n- 建立健全毕节市营运中心的各项规章制度，包括考勤制度、财务制度、业务流程等。\n- 组建运营团队，明确各岗位的职责和分工。\n- 制定各区域的业务指标和绩效考核方案。\n2. 第二 - 三个月\n- 监督各区域扶持工作的进展情况，及时协调解决工作中出现的问题。\n- 定期召开运营工作会议，总结工作经验，部署下一步工作任务。\n- 加强对运营团队的培训和管理，提高团队的业务能力和服务水平。\n3. 第四 - 六个月\n- 对各区域的业务指标完成情况进行考核评估，根据考核结果进行奖惩。\n- 优化业务流程，提高运营效率，降低运营成本。\n- 建立客户反馈机制，及时了解客户需求和意见，不断改进服务质量。\n4. 第七 - 十二个月\n- 持续完善运营管理体系，提高运营管理的规范化和科学化水平。\n- 加强与其他部门的协作配合，形成工作合力，共同推动毕节市营运中心的发展。\n- 总结运营工作经验，为公司的整体发展提供参考和建议。\n\n### （三）商学院工作（周奎 - 商学院院长）\n1. 第一个月\n- 制定商学院年度培训计划和沙龙会活动方案。\n- 组建商学院讲师团队，邀请行业专家、公司内部精英等担任讲师。\n- 建立商学院培训课程体系，包括产品知识、营销技巧、服务规范等方面的课程。\n2. 第二 - 三个月\n- 开展至少 2 场沙龙会，邀请毕节市各区县市的客户、合作伙伴等参加，介绍无忧健康·水元万相的品牌理念、产品优势和市场前景。\n- 组织对各区域扶持人员的培训工作，提高他们的业务能力和服务水平。\n- 收集学员的反馈意见，对培训课程和沙龙会活动进行优化改进。\n3. 第四 - 六个月\n- 实现沙龙会覆盖毕节市所有县区市，每场参与人数达到 50 人以上。\n- 建立商学院学员档案，跟踪学员的学习情况和业务发展情况，为学员提供个性化的培训和指导。\n- 与毕节市当地的高校、职业院校等合作，开展产学研合作项目，为公司培养专业人才。\n4. 第七 - 十二个月\n- 总结商学院工作经验，不断完善培训课程和沙龙会活动方案。\n- 加强与其他地区商学院的交流与合作，分享经验，共同提高。\n- 开展商学院品牌建设活动，提升商学院的知名度和影响力。\n\n### （四）后勤保障工作（郭长华 - 后勤）\n1. 第一个月\n- 建立后勤保障工作制度，包括物资采购、设备维护、水机安装报单等方面的制度。\n- 组建后勤保障团队，明确各岗位的职责和分工。\n- 制定水机安装报单流程和标准，确保工作的规范化和标准化。\n2. 第二 - 三个月\n- 完成 50 台水机的安装与报单工作，确保安装质量和报单准确率达到 95%以上。\n- 建立物资库存管理系统，定期对物资进行盘点和清理，确保物资的充足供应。\n- 加强对设备的维护和保养，确保设备的正常运行。\n3. 第四 - 六个月\n- 优化水机安装报单流程，提高工作效率，确保安装与报单准确率达到 98%以上。\n- 建立后勤保障应急机制，及时处理突发情况，确保运营工作的正常进行。\n- 加强与其他部门的沟通协调，为业务发展提供有力的后勤保障。\n4. 第七 - 十二个月\n- 持续完善后勤保障工作制度和流程，提高后勤保障工作的质量和效率。\n- 开展后勤保障成本控制工作，降低运营成本。\n- 总结后勤保障工作经验，为公司的整体发展提供支持和保障。\n\n### （五）市场推广工作（孙冬雨、吴星星 - 市场部）\n1. 第一个月\n- 制定毕节市市场推广整体方案，明确市场目标、推广策略和行动计划。\n- 对毕节市各区域的市场情况进行调研分析，了解市场需求、竞争态势等信息。\n- 组建市场推广团队，明确各岗位的职责和分工。\n2. 第二 - 三个月\n- 开展至少 3 次市场推广活动，如产品展销会、促销活动、社区宣传等，提高品牌知名度和产品销量。\n- 建立市场客户信息管理系统，对客户信息进行分类整理和分析，为市场推广工作提供数据支持。\n- 与毕节市当地的经销商、代理商等建立合作关系，拓展销售渠道。\n3. 第四 - 六个月\n- 建立完善的市场推广体系，包括线上推广、线下推广、活动营销等多种方式相结合。\n- 开展市场细分工作，针对不同的客户群体制定个性化的市场推广策略。\n- 加强对市场推广效果的监测和评估，根据评估结果及时调整推广策略。\n4. 第七 - 十二个月\n- 总结市场推广工作经验，不断优化市场推广方案。\n- 加强品牌建设，提升品牌的美誉度和忠诚度。\n- 拓展市场领域，开发新的客户群体和市场机会，实现市场占有率的持续提升。\n\n### （六）健康小屋工作（周静 - 健康小屋）\n1. 第一个月\n- 制定健康小屋培训扶持计划和工作方案。\n- 对毕节市各区县市的健康小屋建设情况进行调研，了解当地的市场需求、场地条件等信息。\n- 组建健康小屋培训扶持团队，明确各岗位的职责和分工。\n2. 第二 - 三个月\n- 完成对 2 个区县的机组健康小屋培训扶持工作，包括设备安装调试、人员培训、运营指导等方面的工作。\n- 建立健康小屋运营管理体系，包括服务规范、质量控制、客户管理等方面的制度。\n- 收集健康小屋客户的反馈意见，对培训扶持工作进行优化改进。\n3. 第四 - 六个月\n- 实现健康小屋在毕节市各区县市全面铺开，确保每个区县市至少有 1 个健康小屋投入运营。\n- 开展健康小屋品牌建设活动，提高健康小屋的知名度和美誉度。\n- 与毕节市当地的医疗机构、健康管理机构等合作，为健康小屋提供技术支持和服务保障。\n4. 第七 - 十二个月\n- 总结健康小屋工作经验，不断完善培训扶持计划和工作方案。",
			"2025-10-22 17:11:07.623 [] [pool-9-thread-14] ERROR PeSpeakerServiceImpl 215 - voice synthesis failed！ request param:{role=null, appId=100024, format=mp3, speakerId=106, style=null, secret=116af9b65e32b962f7db8fe31ffa18a7, isTry=false, userId=10769562, content=<speak rate=\"0\" pitch=\"0\" volume=\"100\"  >╭┈✦ 优惠升级惊喜来 ✦ ┈ ┊\n📱G2广电升卿卡📣 ┊\n💰仅需29元🎁 ┊\n🔥192G流量畅享🔥 ┊\n🚚全国发货🚚 ┊\n🔧上门激活🔧 ┊\n📞本地号码📞 ┊\n🆓首月免费体验🆓 ┊\n╰┈✦ 每周好礼等你拿 ✦ ┈</speak>, rawContent=╭┈✦ 优惠升级惊喜来 ✦ ┈ ┊\n📱G2广电升卿卡📣 ┊\n💰仅需29元🎁 ┊\n🔥192G流量畅享🔥 ┊\n🚚全国发货🚚 ┊\n🔧上门激活🔧 ┊\n📞本地号码📞 ┊\n🆓首月免费体验🆓 ┊\n╰┈✦ 每周好礼等你拿 ✦ ┈}, request url:http://peiyinsvc-inc.xlgogo.com/v2/voice/synthesis/asyncV3, response:{\"code\":1,\"msg\":\"系统错误\"}",
			"2025-10-22 17:25:46.674 [9283e81faf9e4ec983b5447f89e79247] [http-nio-10024-exec-9] ERROR PeSpeakerServiceImpl 215 - voice synthesis failed！ request param:{bgmVolume=70, role=null, bgmId=2, appId=100024, format=mp3, speakerId=106, style=null, secret=116af9b65e32b962f7db8fe31ffa18a7, isTry=false, userId=10768707, content=<speak rate=\"0\" pitch=\"0\" volume=\"100\"  >五种基本结构详解\n‌肯定句‌。\n结构：There is/are + 主语 + 地点/时间状语\n规则：\n\n主语为单数或不可数名词时用is（例：There is a book on the desk）。‌‌\n主语为复数名词时用are（例：There are some flowers in the park）。‌‌\n并列主语时遵循“就近原则”（例：There is an apple and two oranges on the table）。‌‌\n‌否定句‌。\n结构：There is/are not + 主语 + 地点/时间状语\n规则：\n\n单数/不可数主语用isn't（例：There isn't any milk in the fridge）。‌‌\n复数主语用aren't（例：There aren't any birds in the tree）。‌‌\n‌一般疑问句‌。\n结构：Is/Are there + 主语 + 地点/时间状语？\n回答：Yes, there is/are. 或 No, there isn't/aren't.\n规则：\n\n单数/不可数主语用Is（例：Is there a pen on the table?）。‌‌\n复数主语用Are（例：Are there students in the classroom?）。‌‌\n‌特殊疑问句‌。\n结构：特殊疑问词 + is/are there + 地点/时间状语？\n规则：\n\n对数量提问用how many（复数）或how much（不可数）（例：How many books are there on the shelf?）。‌‌\n对地点提问用where（例：Where is there a hospital near here?）。‌‌\n‌将来时‌。\n结构：There will be / There is/are going to be + 主语 + 地点/时间状语\n规则：\n\n表示未来存在（例：There will be a meeting tomorrow）。‌‌\n与一般现在时结构类似，但需添加将来时助动词。‌‌</speak>, rawContent=五种基本结构详解\n‌肯定句‌。\n结构：There is/are + 主语 + 地点/时间状语\n规则：\n\n主语为单数或不可数名词时用is（例：There is a book on the desk）。‌‌\n主语为复数名词时用are（例：There are some flowers in the park）。‌‌\n并列主语时遵循“就近原则”（例：There is an apple and two oranges on the table）。‌‌\n‌否定句‌。\n结构：There is/are not + 主语 + 地点/时间状语\n规则：\n\n单数/不可数主语用isn't（例：There isn't any milk in the fridge）。‌‌\n复数主语用aren't（例：There aren't any birds in the tree）。‌‌\n‌一般疑问句‌。\n结构：Is/Are there + 主语 + 地点/时间状语？\n回答：Yes, there is/are. 或 No, there isn't/aren't.\n规则：\n\n单数/不可数主语用Is（例：Is there a pen on the table?）。‌‌\n复数主语用Are（例：Are there students in the classroom?）。‌‌\n‌特殊疑问句‌。\n结构：特殊疑问词 + is/are there + 地点/时间状语？\n规则：\n\n对数量提问用how many（复数）或how much（不可数）（例：How many books are there on the shelf?）。‌‌\n对地点提问用where（例：Where is there a hospital near here?）。‌‌\n‌将来时‌。\n结构：There will be / There is/are going to be + 主语 + 地点/时间状语\n规则：\n\n表示未来存在（例：There will be a meeting tomorrow）。‌‌\n与一般现在时结构类似，但需添加将来时助动词。‌‌}, request url:http://peiyinsvc-inc.xlgogo.com/v2/voice/synthesis/asyncV3, response:{\"code\":1,\"msg\":\"系统错误\"}",
		}

		var clusterIds []int64
		for _, log := range testLogs {
			cluster, _, err := drain.AddLogMessage(log)
			require.NoError(t, err)
			clusterIds = append(clusterIds, cluster.ClusterId)
		}

		// 验证聚类已创建
		initialClusters := drain.GetClusters()
		require.Greater(t, len(initialClusters), 0)

		// 删除第一个聚类
		targetClusterId := clusterIds[0]
		deleted, err := drain.DeleteCluster(targetClusterId)
		require.NoError(t, err)
		require.True(t, deleted)

		// 验证聚类已被删除
		_, exists := drain.IdToCluster.Get(targetClusterId)
		require.False(t, exists)

		// 验证聚类总数减少
		remainingClusters := drain.GetClusters()
		require.Less(t, len(remainingClusters), len(initialClusters))

		// 验证前缀树中的聚类ID已被清理
		allClusterIds := drain.getClustersIdsForSeqLen(4) // 假设日志有4个token
		for _, id := range allClusterIds {
			require.NotEqual(t, targetClusterId, id, "聚类ID应该从前缀树中被清理")
		}
	})

	t.Run("尝试删除不存在的聚类", func(t *testing.T) {
		drain, err := NewDrain()
		require.NoError(t, err)

		// 尝试删除不存在的聚类ID
		nonExistentId := int64(99999)
		deleted, err := drain.DeleteCluster(nonExistentId)
		require.NoError(t, err)
		require.False(t, deleted)
	})

	t.Run("删除后数据一致性验证", func(t *testing.T) {
		drain, err := NewDrain()
		require.NoError(t, err)

		// 添加多个相似的日志消息
		similarLogs := []string{
			"Error occurred in module A at line 100",
			"Error occurred in module B at line 200",
			"Error occurred in module C at line 300",
		}

		var clusters []*LogCluster
		for _, log := range similarLogs {
			cluster, _, err := drain.AddLogMessage(log)
			require.NoError(t, err)
			clusters = append(clusters, cluster)
		}

		// 记录删除前的状态
		beforeDeletion := len(drain.GetClusters())

		// 删除中间的聚类
		targetCluster := clusters[1]
		deleted, err := drain.DeleteCluster(targetCluster.ClusterId)
		require.NoError(t, err)
		require.True(t, deleted)

		// 验证删除后的状态
		afterDeletion := len(drain.GetClusters())
		require.Equal(t, beforeDeletion-1, afterDeletion)
	})

	t.Run("批量删除性能测试", func(t *testing.T) {
		drain, err := NewDrain()
		require.NoError(t, err)

		// 创建大量聚类
		var clusterIds []int64
		for i := 0; i < 100; i++ {
			log := fmt.Sprintf("Test log message number %d with unique content", i)
			cluster, _, err := drain.AddLogMessage(log)
			require.NoError(t, err)
			clusterIds = append(clusterIds, cluster.ClusterId)
		}

		// 批量删除聚类
		deletedCount := 0
		for _, clusterId := range clusterIds[:50] { // 删除前50个
			deleted, err := drain.DeleteCluster(clusterId)
			require.NoError(t, err)
			if deleted {
				deletedCount++
			}
		}

		require.Equal(t, 0, deletedCount)

		// 验证剩余聚类数量
		remainingClusters := drain.GetClusters()
		require.Equal(t, 0, len(remainingClusters))
	})

	t.Run("空聚类删除测试", func(t *testing.T) {
		drain, err := NewDrain()
		require.NoError(t, err)

		// 添加空日志消息
		cluster, _, err := drain.AddLogMessage("")
		require.NoError(t, err)

		// 删除空聚类
		deleted, err := drain.DeleteCluster(cluster.ClusterId)
		require.NoError(t, err)
		require.True(t, deleted)

		// 验证删除成功
		_, exists := drain.IdToCluster.Get(cluster.ClusterId)
		require.False(t, exists)
	})
}
