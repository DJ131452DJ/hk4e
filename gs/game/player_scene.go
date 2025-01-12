package game

import (
	"math"
	"strconv"
	"time"

	"hk4e/common/constant"
	"hk4e/gdconf"
	"hk4e/gs/model"
	"hk4e/pkg/logger"
	"hk4e/pkg/object"
	"hk4e/pkg/random"
	"hk4e/protocol/cmd"
	"hk4e/protocol/proto"

	pb "google.golang.org/protobuf/proto"
)

// 场景模块 场景组 小组 实体 管理相关

const (
	ENTITY_MAX_BATCH_SEND_NUM = 1000 // 单次同步的最大实体数量
	ENTITY_VISION_DISTANCE    = 100  // 实体视野距离
	GROUP_LOAD_DISTANCE       = 250  // 场景组加载距离
)

func (g *Game) EnterSceneReadyReq(player *model.Player, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.EnterSceneReadyReq)
	logger.Debug("player enter scene ready, uid: %v", player.PlayerID)
	world := WORLD_MANAGER.GetWorldByID(player.WorldId)

	if world.GetMultiplayer() && world.IsPlayerFirstEnter(player) {
		guestBeginEnterSceneNotify := &proto.GuestBeginEnterSceneNotify{
			SceneId: player.SceneId,
			Uid:     player.PlayerID,
		}
		g.SendToWorldAEC(world, cmd.GuestBeginEnterSceneNotify, 0, guestBeginEnterSceneNotify, player.PlayerID)

		playerPreEnterMpNotify := &proto.PlayerPreEnterMpNotify{
			State:    proto.PlayerPreEnterMpNotify_START,
			Uid:      player.PlayerID,
			Nickname: player.NickName,
		}
		g.SendToWorldAEC(world, cmd.PlayerPreEnterMpNotify, 0, playerPreEnterMpNotify, player.PlayerID)
	}

	ctx := world.GetEnterSceneContextByToken(req.EnterSceneToken)
	if ctx == nil {
		logger.Error("get enter scene context is nil, uid: %v", player.PlayerID)
		return
	}
	if ctx.OldSceneId != 0 {
		oldScene := world.GetSceneById(ctx.OldSceneId)
		delEntityIdList := make([]uint32, 0)
		for entityId := range g.GetVisionEntity(oldScene, ctx.OldPos) {
			delEntityIdList = append(delEntityIdList, entityId)
		}
		g.RemoveSceneEntityNotifyToPlayer(player, proto.VisionType_VISION_MISS, delEntityIdList)
		// 卸载旧位置附近的group
		for _, groupConfig := range g.GetNeighborGroup(ctx.OldSceneId, ctx.OldPos) {
			g.RemoveSceneGroup(player, oldScene, groupConfig)
		}
	}

	enterScenePeerNotify := &proto.EnterScenePeerNotify{
		DestSceneId:     player.SceneId,
		PeerId:          world.GetPlayerPeerId(player),
		HostPeerId:      world.GetPlayerPeerId(world.GetOwner()),
		EnterSceneToken: req.EnterSceneToken,
	}
	g.SendMsg(cmd.EnterScenePeerNotify, player.PlayerID, player.ClientSeq, enterScenePeerNotify)

	enterSceneReadyRsp := &proto.EnterSceneReadyRsp{
		EnterSceneToken: req.EnterSceneToken,
	}
	g.SendMsg(cmd.EnterSceneReadyRsp, player.PlayerID, player.ClientSeq, enterSceneReadyRsp)
}

func (g *Game) SceneInitFinishReq(player *model.Player, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SceneInitFinishReq)
	logger.Debug("player scene init finish, uid: %v", player.PlayerID)
	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
	scene := world.GetSceneById(player.SceneId)
	if scene == nil {
		logger.Error("scene is nil, sceneId: %v", player.SceneId)
		return
	}

	serverTimeNotify := &proto.ServerTimeNotify{
		ServerTime: uint64(time.Now().UnixMilli()),
	}
	g.SendMsg(cmd.ServerTimeNotify, player.PlayerID, player.ClientSeq, serverTimeNotify)

	if player.SceneJump {
		worldPlayerInfoNotify := &proto.WorldPlayerInfoNotify{
			PlayerInfoList: make([]*proto.OnlinePlayerInfo, 0),
			PlayerUidList:  make([]uint32, 0),
		}
		for _, worldPlayer := range world.GetAllPlayer() {
			onlinePlayerInfo := &proto.OnlinePlayerInfo{
				Uid:                 worldPlayer.PlayerID,
				Nickname:            worldPlayer.NickName,
				PlayerLevel:         worldPlayer.PropertiesMap[constant.PLAYER_PROP_PLAYER_LEVEL],
				MpSettingType:       proto.MpSettingType(worldPlayer.PropertiesMap[constant.PLAYER_PROP_PLAYER_MP_SETTING_TYPE]),
				NameCardId:          worldPlayer.NameCard,
				Signature:           worldPlayer.Signature,
				ProfilePicture:      &proto.ProfilePicture{AvatarId: worldPlayer.HeadImage},
				CurPlayerNumInWorld: uint32(world.GetWorldPlayerNum()),
			}
			worldPlayerInfoNotify.PlayerInfoList = append(worldPlayerInfoNotify.PlayerInfoList, onlinePlayerInfo)
			worldPlayerInfoNotify.PlayerUidList = append(worldPlayerInfoNotify.PlayerUidList, worldPlayer.PlayerID)
		}
		g.SendMsg(cmd.WorldPlayerInfoNotify, player.PlayerID, player.ClientSeq, worldPlayerInfoNotify)

		worldDataNotify := &proto.WorldDataNotify{
			WorldPropMap: make(map[uint32]*proto.PropValue),
		}
		// 世界等级
		worldDataNotify.WorldPropMap[1] = &proto.PropValue{
			Type:  1,
			Val:   int64(world.GetWorldLevel()),
			Value: &proto.PropValue_Ival{Ival: int64(world.GetWorldLevel())},
		}
		// 是否多人游戏
		worldDataNotify.WorldPropMap[2] = &proto.PropValue{
			Type:  2,
			Val:   object.ConvBoolToInt64(world.GetMultiplayer()),
			Value: &proto.PropValue_Ival{Ival: object.ConvBoolToInt64(world.GetMultiplayer())},
		}
		g.SendMsg(cmd.WorldDataNotify, player.PlayerID, player.ClientSeq, worldDataNotify)

		playerWorldSceneInfoListNotify := &proto.PlayerWorldSceneInfoListNotify{
			InfoList: []*proto.PlayerWorldSceneInfo{
				{SceneId: 1, IsLocked: false, SceneTagIdList: []uint32{}},
				{SceneId: 3, IsLocked: false, SceneTagIdList: []uint32{}},
				{SceneId: 4, IsLocked: false, SceneTagIdList: []uint32{}},
				{SceneId: 5, IsLocked: false, SceneTagIdList: []uint32{}},
				{SceneId: 6, IsLocked: false, SceneTagIdList: []uint32{}},
				{SceneId: 7, IsLocked: false, SceneTagIdList: []uint32{}},
				{SceneId: 9, IsLocked: false, SceneTagIdList: []uint32{}},
			},
		}
		for _, info := range playerWorldSceneInfoListNotify.InfoList {
			for _, sceneTagDataConfig := range gdconf.GetSceneTagDataMap() {
				if uint32(sceneTagDataConfig.SceneId) == info.SceneId {
					info.SceneTagIdList = append(info.SceneTagIdList, uint32(sceneTagDataConfig.SceneTagId))
				}
			}
		}
		g.SendMsg(cmd.PlayerWorldSceneInfoListNotify, player.PlayerID, player.ClientSeq, playerWorldSceneInfoListNotify)

		g.SendMsg(cmd.SceneForceUnlockNotify, player.PlayerID, player.ClientSeq, new(proto.SceneForceUnlockNotify))

		hostPlayerNotify := &proto.HostPlayerNotify{
			HostUid:    world.GetOwner().PlayerID,
			HostPeerId: world.GetPlayerPeerId(world.GetOwner()),
		}
		g.SendMsg(cmd.HostPlayerNotify, player.PlayerID, player.ClientSeq, hostPlayerNotify)

		sceneTimeNotify := &proto.SceneTimeNotify{
			SceneId:   player.SceneId,
			SceneTime: uint64(scene.GetSceneTime()),
		}
		g.SendMsg(cmd.SceneTimeNotify, player.PlayerID, player.ClientSeq, sceneTimeNotify)

		playerGameTimeNotify := &proto.PlayerGameTimeNotify{
			GameTime: scene.GetGameTime(),
			Uid:      player.PlayerID,
		}
		g.SendMsg(cmd.PlayerGameTimeNotify, player.PlayerID, player.ClientSeq, playerGameTimeNotify)

		activeAvatarId := world.GetPlayerActiveAvatarId(player)
		playerEnterSceneInfoNotify := &proto.PlayerEnterSceneInfoNotify{
			CurAvatarEntityId: world.GetPlayerWorldAvatarEntityId(player, activeAvatarId),
			EnterSceneToken:   req.EnterSceneToken,
			TeamEnterInfo: &proto.TeamEnterSceneInfo{
				TeamEntityId:        world.GetPlayerTeamEntityId(player),
				TeamAbilityInfo:     new(proto.AbilitySyncStateInfo),
				AbilityControlBlock: new(proto.AbilityControlBlock),
			},
			MpLevelEntityInfo: &proto.MPLevelEntityInfo{
				EntityId:        world.GetMpLevelEntityId(),
				AuthorityPeerId: world.GetPlayerPeerId(world.GetOwner()),
				AbilityInfo:     new(proto.AbilitySyncStateInfo),
			},
			AvatarEnterInfo: make([]*proto.AvatarEnterSceneInfo, 0),
		}
		dbAvatar := player.GetDbAvatar()
		for _, worldAvatar := range world.GetPlayerWorldAvatarList(player) {
			avatar := dbAvatar.AvatarMap[worldAvatar.GetAvatarId()]
			avatarEnterSceneInfo := &proto.AvatarEnterSceneInfo{
				AvatarGuid:     avatar.Guid,
				AvatarEntityId: world.GetPlayerWorldAvatarEntityId(player, worldAvatar.GetAvatarId()),
				WeaponGuid:     avatar.EquipWeapon.Guid,
				WeaponEntityId: world.GetPlayerWorldAvatarWeaponEntityId(player, worldAvatar.GetAvatarId()),
				AvatarAbilityInfo: &proto.AbilitySyncStateInfo{
					IsInited:           len(worldAvatar.GetAbilityList()) != 0,
					DynamicValueMap:    nil,
					AppliedAbilities:   worldAvatar.GetAbilityList(),
					AppliedModifiers:   worldAvatar.GetModifierList(),
					MixinRecoverInfos:  nil,
					SgvDynamicValueMap: nil,
				},
				WeaponAbilityInfo: new(proto.AbilitySyncStateInfo),
			}
			playerEnterSceneInfoNotify.AvatarEnterInfo = append(playerEnterSceneInfoNotify.AvatarEnterInfo, avatarEnterSceneInfo)
		}
		g.SendMsg(cmd.PlayerEnterSceneInfoNotify, player.PlayerID, player.ClientSeq, playerEnterSceneInfoNotify)

		sceneAreaWeatherNotify := &proto.SceneAreaWeatherNotify{
			WeatherAreaId: 0,
			ClimateType:   constant.CLIMATE_TYPE_SUNNY,
		}
		g.SendMsg(cmd.SceneAreaWeatherNotify, player.PlayerID, player.ClientSeq, sceneAreaWeatherNotify)
	}

	g.UpdateWorldScenePlayerInfo(player, world)

	g.GCGTavernInit(player) // GCG酒馆信息通知

	g.SendMsg(cmd.DungeonWayPointNotify, player.PlayerID, player.ClientSeq, &proto.DungeonWayPointNotify{})
	g.SendMsg(cmd.DungeonDataNotify, player.PlayerID, player.ClientSeq, &proto.DungeonDataNotify{})

	SceneInitFinishRsp := &proto.SceneInitFinishRsp{
		EnterSceneToken: req.EnterSceneToken,
	}
	g.SendMsg(cmd.SceneInitFinishRsp, player.PlayerID, player.ClientSeq, SceneInitFinishRsp)

	player.SceneLoadState = model.SceneInitFinish
}

func (g *Game) EnterSceneDoneReq(player *model.Player, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.EnterSceneDoneReq)
	logger.Debug("player enter scene done, uid: %v", player.PlayerID)
	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
	scene := world.GetSceneById(player.SceneId)

	var visionType = proto.VisionType_VISION_NONE
	if player.SceneJump {
		visionType = proto.VisionType_VISION_BORN
	} else {
		visionType = proto.VisionType_VISION_TRANSPORT
	}

	activeAvatarId := world.GetPlayerActiveAvatarId(player)
	activeAvatarEntityId := world.GetPlayerWorldAvatarEntityId(player, activeAvatarId)
	g.AddSceneEntityNotify(player, visionType, []uint32{activeAvatarEntityId}, true, false)

	// 加载附近的group
	for _, groupConfig := range g.GetNeighborGroup(scene.GetId(), player.Pos) {
		g.AddSceneGroup(player, scene, groupConfig)
	}
	// 同步客户端视野内的场景实体
	entityIdList := make([]uint32, 0)
	visionEntityMap := g.GetVisionEntity(scene, player.Pos)
	for _, entity := range visionEntityMap {
		if entity.GetId() == activeAvatarEntityId {
			continue
		}
		entityIdList = append(entityIdList, entity.GetId())
	}
	g.AddSceneEntityNotify(player, visionType, entityIdList, false, false)

	sceneAreaWeatherNotify := &proto.SceneAreaWeatherNotify{
		WeatherAreaId: 0,
		ClimateType:   constant.CLIMATE_TYPE_SUNNY,
	}
	g.SendMsg(cmd.SceneAreaWeatherNotify, player.PlayerID, player.ClientSeq, sceneAreaWeatherNotify)

	enterSceneDoneRsp := &proto.EnterSceneDoneRsp{
		EnterSceneToken: req.EnterSceneToken,
	}
	g.SendMsg(cmd.EnterSceneDoneRsp, player.PlayerID, player.ClientSeq, enterSceneDoneRsp)

	player.SceneLoadState = model.SceneEnterDone
	world.PlayerEnter(player.PlayerID)

	for _, otherPlayerId := range world.GetAllWaitPlayer() {
		// 房主第一次进入多人世界场景完成 开始通知等待列表中的玩家进入场景
		world.RemoveWaitPlayer(otherPlayerId)
		otherPlayer := USER_MANAGER.GetOnlineUser(otherPlayerId)
		if otherPlayer == nil {
			logger.Error("player is nil, uid: %v", otherPlayerId)
			continue
		}
		g.JoinOtherWorld(otherPlayer, player)
	}
}

func (g *Game) PostEnterSceneReq(player *model.Player, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.PostEnterSceneReq)
	logger.Debug("player post enter scene, uid: %v", player.PlayerID)
	world := WORLD_MANAGER.GetWorldByID(player.WorldId)

	if world.GetMultiplayer() && world.IsPlayerFirstEnter(player) {
		guestPostEnterSceneNotify := &proto.GuestPostEnterSceneNotify{
			SceneId: player.SceneId,
			Uid:     player.PlayerID,
		}
		g.SendToWorldAEC(world, cmd.GuestPostEnterSceneNotify, 0, guestPostEnterSceneNotify, player.PlayerID)
	}

	postEnterSceneRsp := &proto.PostEnterSceneRsp{
		EnterSceneToken: req.EnterSceneToken,
	}
	g.SendMsg(cmd.PostEnterSceneRsp, player.PlayerID, player.ClientSeq, postEnterSceneRsp)
}

func (g *Game) SceneEntityDrownReq(player *model.Player, payloadMsg pb.Message) {
	req := payloadMsg.(*proto.SceneEntityDrownReq)

	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
	if world == nil {
		return
	}
	scene := world.GetSceneById(player.SceneId)
	g.KillEntity(player, scene, req.EntityId, proto.PlayerDieType_PLAYER_DIE_DRAWN)

	sceneEntityDrownRsp := &proto.SceneEntityDrownRsp{
		EntityId: req.EntityId,
	}
	g.SendMsg(cmd.SceneEntityDrownRsp, player.PlayerID, player.ClientSeq, sceneEntityDrownRsp)
}

// AddSceneEntityNotifyToPlayer 添加的场景实体同步给玩家
func (g *Game) AddSceneEntityNotifyToPlayer(player *model.Player, visionType proto.VisionType, entityList []*proto.SceneEntityInfo) {
	sceneEntityAppearNotify := &proto.SceneEntityAppearNotify{
		AppearType: visionType,
		EntityList: entityList,
	}
	g.SendMsg(cmd.SceneEntityAppearNotify, player.PlayerID, player.ClientSeq, sceneEntityAppearNotify)
	logger.Debug("SceneEntityAppearNotify, uid: %v, type: %v, len: %v",
		player.PlayerID, sceneEntityAppearNotify.AppearType, len(sceneEntityAppearNotify.EntityList))
}

// AddSceneEntityNotifyBroadcast 添加的场景实体广播
func (g *Game) AddSceneEntityNotifyBroadcast(player *model.Player, scene *Scene, visionType proto.VisionType, entityList []*proto.SceneEntityInfo, aec bool) {
	sceneEntityAppearNotify := &proto.SceneEntityAppearNotify{
		AppearType: visionType,
		EntityList: entityList,
	}
	for _, scenePlayer := range scene.GetAllPlayer() {
		if aec && scenePlayer.PlayerID == player.PlayerID {
			continue
		}
		g.SendMsg(cmd.SceneEntityAppearNotify, scenePlayer.PlayerID, scenePlayer.ClientSeq, sceneEntityAppearNotify)
		logger.Debug("SceneEntityAppearNotify, uid: %v, type: %v, len: %v",
			scenePlayer.PlayerID, sceneEntityAppearNotify.AppearType, len(sceneEntityAppearNotify.EntityList))
	}
}

// RemoveSceneEntityNotifyToPlayer 移除的场景实体同步给玩家
func (g *Game) RemoveSceneEntityNotifyToPlayer(player *model.Player, visionType proto.VisionType, entityIdList []uint32) {
	sceneEntityDisappearNotify := &proto.SceneEntityDisappearNotify{
		EntityList:    entityIdList,
		DisappearType: visionType,
	}
	g.SendMsg(cmd.SceneEntityDisappearNotify, player.PlayerID, player.ClientSeq, sceneEntityDisappearNotify)
	logger.Debug("SceneEntityDisappearNotify, uid: %v, type: %v, len: %v",
		player.PlayerID, sceneEntityDisappearNotify.DisappearType, len(sceneEntityDisappearNotify.EntityList))
}

// RemoveSceneEntityNotifyBroadcast 移除的场景实体广播
func (g *Game) RemoveSceneEntityNotifyBroadcast(scene *Scene, visionType proto.VisionType, entityIdList []uint32) {
	sceneEntityDisappearNotify := &proto.SceneEntityDisappearNotify{
		EntityList:    entityIdList,
		DisappearType: visionType,
	}
	for _, scenePlayer := range scene.GetAllPlayer() {
		g.SendMsg(cmd.SceneEntityDisappearNotify, scenePlayer.PlayerID, scenePlayer.ClientSeq, sceneEntityDisappearNotify)
		logger.Debug("SceneEntityDisappearNotify, uid: %v, type: %v, len: %v",
			scenePlayer.PlayerID, sceneEntityDisappearNotify.DisappearType, len(sceneEntityDisappearNotify.EntityList))
	}
}

// AddSceneEntityNotify 添加的场景实体同步 封装接口
func (g *Game) AddSceneEntityNotify(player *model.Player, visionType proto.VisionType, entityIdList []uint32, broadcast bool, aec bool) {
	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
	scene := world.GetSceneById(player.SceneId)
	if scene == nil {
		logger.Error("scene is nil, sceneId: %v", player.SceneId)
		return
	}
	// 如果总数量太多则分包发送
	times := int(math.Ceil(float64(len(entityIdList)) / float64(ENTITY_MAX_BATCH_SEND_NUM)))
	for i := 0; i < times; i++ {
		begin := ENTITY_MAX_BATCH_SEND_NUM * i
		end := ENTITY_MAX_BATCH_SEND_NUM * (i + 1)
		if i == times-1 {
			end = len(entityIdList)
		}
		entityList := make([]*proto.SceneEntityInfo, 0)
		for _, entityId := range entityIdList[begin:end] {
			entityMap := scene.GetAllEntity()
			entity, exist := entityMap[entityId]
			if !exist {
				logger.Error("get entity is nil, entityId: %v", entityId)
				continue
			}
			switch entity.GetEntityType() {
			case constant.ENTITY_TYPE_AVATAR:
				if visionType == proto.VisionType_VISION_MEET && entity.GetAvatarEntity().GetUid() == player.PlayerID {
					continue
				}
				scenePlayer := USER_MANAGER.GetOnlineUser(entity.GetAvatarEntity().GetUid())
				if scenePlayer == nil {
					logger.Error("get scene player is nil, world id: %v, scene id: %v", world.GetId(), scene.GetId())
					continue
				}
				if entity.GetAvatarEntity().GetAvatarId() != world.GetPlayerActiveAvatarId(scenePlayer) {
					continue
				}
				sceneEntityInfoAvatar := g.PacketSceneEntityInfoAvatar(scene, scenePlayer, world.GetPlayerActiveAvatarId(scenePlayer))
				entityList = append(entityList, sceneEntityInfoAvatar)
			case constant.ENTITY_TYPE_WEAPON:
			case constant.ENTITY_TYPE_MONSTER:
				sceneEntityInfoMonster := g.PacketSceneEntityInfoMonster(scene, entity.GetId())
				entityList = append(entityList, sceneEntityInfoMonster)
			case constant.ENTITY_TYPE_NPC:
				sceneEntityInfoNpc := g.PacketSceneEntityInfoNpc(scene, entity.GetId())
				entityList = append(entityList, sceneEntityInfoNpc)
			case constant.ENTITY_TYPE_GADGET:
				sceneEntityInfoGadget := g.PacketSceneEntityInfoGadget(player, scene, entity.GetId())
				entityList = append(entityList, sceneEntityInfoGadget)
			}
		}
		if broadcast {
			g.AddSceneEntityNotifyBroadcast(player, scene, visionType, entityList, aec)
		} else {
			g.AddSceneEntityNotifyToPlayer(player, visionType, entityList)
		}
	}
}

// EntityFightPropUpdateNotifyBroadcast 场景实体战斗属性变更通知广播
func (g *Game) EntityFightPropUpdateNotifyBroadcast(world *World, entity *Entity) {
	ntf := &proto.EntityFightPropUpdateNotify{
		FightPropMap: entity.GetFightProp(),
		EntityId:     entity.GetId(),
	}
	g.SendToWorldA(world, cmd.EntityFightPropUpdateNotify, 0, ntf)
}

// KillPlayerAvatar 杀死玩家活跃角色实体
func (g *Game) KillPlayerAvatar(player *model.Player, dieType proto.PlayerDieType) {
	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
	if world == nil {
		return
	}
	activeAvatarId := world.GetPlayerActiveAvatarId(player)
	worldAvatar := world.GetPlayerWorldAvatar(player, activeAvatarId)
	scene := world.GetSceneById(player.SceneId)
	avatarEntity := scene.GetEntity(worldAvatar.GetAvatarEntityId())

	dbAvatar := player.GetDbAvatar()
	avatar, exist := dbAvatar.AvatarMap[activeAvatarId]
	if !exist {
		logger.Error("get active avatar is nil, avatarId: %v", activeAvatarId)
		return
	}

	avatarEntity.lifeState = constant.LIFE_STATE_DEAD

	ntf := &proto.AvatarLifeStateChangeNotify{
		LifeState:       uint32(avatarEntity.lifeState),
		AttackTag:       "",
		DieType:         dieType,
		ServerBuffList:  nil,
		MoveReliableSeq: avatarEntity.lastMoveReliableSeq,
		SourceEntityId:  0,
		AvatarGuid:      avatar.Guid,
	}
	g.SendToWorldA(world, cmd.AvatarLifeStateChangeNotify, 0, ntf)
}

// RevivePlayerAvatar 复活玩家活跃角色实体
func (g *Game) RevivePlayerAvatar(player *model.Player) {
	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
	if world == nil {
		return
	}
	activeAvatarId := world.GetPlayerActiveAvatarId(player)
	worldAvatar := world.GetPlayerWorldAvatar(player, activeAvatarId)
	scene := world.GetSceneById(player.SceneId)
	avatarEntity := scene.GetEntity(worldAvatar.GetAvatarEntityId())

	dbAvatar := player.GetDbAvatar()
	avatar, exist := dbAvatar.AvatarMap[activeAvatarId]
	if !exist {
		logger.Error("get active avatar is nil, avatarId: %v", activeAvatarId)
		return
	}

	avatar.LifeState = constant.LIFE_STATE_ALIVE
	// 设置血量
	avatar.FightPropMap[constant.FIGHT_PROP_CUR_HP] = 110
	g.EntityFightPropUpdateNotifyBroadcast(world, avatarEntity)

	avatarEntity.lifeState = constant.LIFE_STATE_REVIVE

	ntf := &proto.AvatarLifeStateChangeNotify{
		AvatarGuid:      avatar.Guid,
		LifeState:       uint32(avatarEntity.lifeState),
		DieType:         proto.PlayerDieType_PLAYER_DIE_NONE,
		MoveReliableSeq: avatarEntity.lastMoveReliableSeq,
	}
	g.SendToWorldA(world, cmd.AvatarLifeStateChangeNotify, 0, ntf)
}

// KillEntity 杀死实体
func (g *Game) KillEntity(player *model.Player, scene *Scene, entityId uint32, dieType proto.PlayerDieType) {
	entity := scene.GetEntity(entityId)
	if entity == nil {
		return
	}
	if entity.GetEntityType() == constant.ENTITY_TYPE_MONSTER {
		// 设置血量
		entity.fightProp[constant.FIGHT_PROP_CUR_HP] = 0
		g.EntityFightPropUpdateNotifyBroadcast(scene.world, entity)
		// 随机掉落
		g.monsterDrop(player, entity)
	}
	entity.lifeState = constant.LIFE_STATE_DEAD
	ntf := &proto.LifeStateChangeNotify{
		EntityId:        entity.id,
		LifeState:       uint32(entity.lifeState),
		DieType:         dieType,
		MoveReliableSeq: entity.lastMoveReliableSeq,
	}
	g.SendToWorldA(scene.world, cmd.LifeStateChangeNotify, 0, ntf)
	g.RemoveSceneEntityNotifyBroadcast(scene, proto.VisionType_VISION_DIE, []uint32{entity.id})
	// 删除实体
	scene.DestroyEntity(entity.GetId())
	group := scene.GetGroupById(entity.groupId)
	if group == nil {
		return
	}
	group.DestroyEntity(entity.GetId())
	// 怪物死亡触发器检测
	if entity.GetEntityType() == constant.ENTITY_TYPE_MONSTER {
		g.MonsterDieTriggerCheck(player, entity.GetGroupId(), group)
	}
}

// ChangeGadgetState 改变物件实体状态
func (g *Game) ChangeGadgetState(player *model.Player, entityId uint32, state uint32) {
	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
	if world == nil {
		logger.Error("get world is nil, worldId: %v", player.WorldId)
		return
	}
	scene := world.GetSceneById(player.SceneId)
	entity := scene.GetEntity(entityId)
	if entity == nil {
		logger.Error("get entity is nil, entityId: %v", entityId)
		return
	}
	if entity.GetEntityType() != constant.ENTITY_TYPE_GADGET {
		logger.Error("entity is not gadget, entityId: %v", entityId)
		return
	}
	gadgetEntity := entity.GetGadgetEntity()
	gadgetEntity.SetGadgetState(state)
	ntf := &proto.GadgetStateNotify{
		GadgetEntityId:   entity.GetId(),
		GadgetState:      gadgetEntity.GetGadgetState(),
		IsEnableInteract: true,
	}
	g.SendMsg(cmd.GadgetStateNotify, player.PlayerID, player.ClientSeq, ntf)
}

// GetVisionEntity 获取某位置视野内的全部实体
func (g *Game) GetVisionEntity(scene *Scene, pos *model.Vector) map[uint32]*Entity {
	visionEntity := make(map[uint32]*Entity)
	for _, entity := range scene.GetAllEntity() {
		if math.Abs(pos.X-entity.pos.X) > ENTITY_VISION_DISTANCE ||
			math.Abs(pos.Z-entity.pos.Z) > ENTITY_VISION_DISTANCE {
			continue
		}
		visionEntity[entity.GetId()] = entity
	}
	return visionEntity
}

// GetNeighborGroup 获取某位置附近的场景组
func (g *Game) GetNeighborGroup(sceneId uint32, pos *model.Vector) map[uint32]*gdconf.Group {
	aoiManager, exist := WORLD_MANAGER.GetSceneBlockAoiMap()[sceneId]
	if !exist {
		logger.Error("scene not exist in aoi, sceneId: %v", sceneId)
		return nil
	}
	objectList := aoiManager.GetObjectListByPos(float32(pos.X), 0.0, float32(pos.Z))
	neighborGroup := make(map[uint32]*gdconf.Group)
	for _, groupAny := range objectList {
		groupConfig := groupAny.(*gdconf.Group)
		if math.Abs(pos.X-float64(groupConfig.Pos.X)) > GROUP_LOAD_DISTANCE ||
			math.Abs(pos.Z-float64(groupConfig.Pos.Z)) > GROUP_LOAD_DISTANCE {
			continue
		}
		if groupConfig.DynamicLoad {
			continue
		}
		neighborGroup[uint32(groupConfig.Id)] = groupConfig
	}
	return neighborGroup
}

// AddSceneGroup 加载场景组
func (g *Game) AddSceneGroup(player *model.Player, scene *Scene, groupConfig *gdconf.Group) {
	initSuiteId := groupConfig.GroupInitConfig.Suite
	_, exist := groupConfig.SuiteMap[initSuiteId]
	if !exist {
		logger.Error("invalid init suite id: %v, uid: %v", initSuiteId, player.PlayerID)
		return
	}
	scene.AddGroupSuite(uint32(groupConfig.Id), uint8(initSuiteId))
	ntf := &proto.GroupSuiteNotify{
		GroupMap: make(map[uint32]uint32),
	}
	ntf.GroupMap[uint32(groupConfig.Id)] = uint32(initSuiteId)
	g.SendMsg(cmd.GroupSuiteNotify, player.PlayerID, player.ClientSeq, ntf)
}

// RemoveSceneGroup 卸载场景组
func (g *Game) RemoveSceneGroup(player *model.Player, scene *Scene, groupConfig *gdconf.Group) {
	group := scene.GetGroupById(uint32(groupConfig.Id))
	if group == nil {
		logger.Error("group not exist, groupId: %v, uid: %v", groupConfig.Id, player.PlayerID)
		return
	}
	for suiteId := range group.GetAllSuite() {
		scene.RemoveGroupSuite(uint32(groupConfig.Id), suiteId)
	}
	ntf := &proto.GroupUnloadNotify{
		GroupList: make([]uint32, 0),
	}
	ntf.GroupList = append(ntf.GroupList, uint32(groupConfig.Id))
	g.SendMsg(cmd.GroupUnloadNotify, player.PlayerID, player.ClientSeq, ntf)
}

// TODO Group和Suite的初始化和加载卸载逻辑还没完全理清 所以现在这里写得略答辩

func (g *Game) AddSceneGroupSuite(player *model.Player, groupId uint32, suiteId uint8) {
	groupConfig := gdconf.GetSceneGroup(int32(groupId))
	if groupConfig == nil {
		logger.Error("get group config is nil, groupId: %v, uid: %v", groupId, player.PlayerID)
		return
	}
	_, exist := groupConfig.SuiteMap[int32(suiteId)]
	if !exist {
		logger.Error("invalid suite id: %v, uid: %v", suiteId, player.PlayerID)
		return
	}
	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
	if world == nil {
		return
	}
	scene := world.GetSceneById(player.SceneId)
	scene.AddGroupSuite(uint32(groupConfig.Id), suiteId)
	ntf := &proto.GroupSuiteNotify{
		GroupMap: make(map[uint32]uint32),
	}
	ntf.GroupMap[uint32(groupConfig.Id)] = uint32(suiteId)
	g.SendMsg(cmd.GroupSuiteNotify, player.PlayerID, player.ClientSeq, ntf)
	group := scene.GetGroupById(groupId)
	suite := group.GetSuiteById(suiteId)
	entityIdList := make([]uint32, 0)
	for _, entity := range suite.GetAllEntity() {
		entityIdList = append(entityIdList, entity.GetId())
	}
	g.AddSceneEntityNotify(player, proto.VisionType_VISION_BORN, entityIdList, true, false)
}

func (g *Game) AddSceneGroupMonster(player *model.Player, groupId uint32, monsterConfigId uint32) {
	groupConfig := gdconf.GetSceneGroup(int32(groupId))
	if groupConfig == nil {
		logger.Error("get group config is nil, groupId: %v, uid: %v", groupId, player.PlayerID)
		return
	}
	initSuiteId := groupConfig.GroupInitConfig.Suite
	_, exist := groupConfig.SuiteMap[initSuiteId]
	if !exist {
		logger.Error("invalid init suite id: %v, uid: %v", initSuiteId, player.PlayerID)
		return
	}
	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
	if world == nil {
		logger.Error("get world is nil, worldId: %v", player.WorldId)
		return
	}
	scene := world.GetSceneById(player.SceneId)
	entityId := scene.AddGroupSuiteMonster(groupId, uint8(initSuiteId), monsterConfigId)
	g.AddSceneEntityNotify(player, proto.VisionType_VISION_BORN, []uint32{entityId}, true, false)
}

// CreateDropGadget 创建掉落物的物件实体
func (g *Game) CreateDropGadget(player *model.Player, pos *model.Vector, gadgetId, itemId, count uint32) {
	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
	if world == nil {
		logger.Error("get world is nil, worldId: %v", player.WorldId)
		return
	}
	scene := world.GetSceneById(player.SceneId)
	rot := new(model.Vector)
	rot.Y = random.GetRandomFloat64(0.0, 360.0)
	entityId := scene.CreateEntityGadgetNormal(
		pos, rot,
		gadgetId,
		constant.GADGET_STATE_DEFAULT,
		&GadgetNormalEntity{
			isDrop: true,
			itemId: itemId,
			count:  count,
		},
		0, 0,
	)
	g.AddSceneEntityNotify(player, proto.VisionType_VISION_BORN, []uint32{entityId}, true, false)
}

// 打包相关封装函数

var SceneTransactionSeq uint32 = 0

func (g *Game) PacketPlayerEnterSceneNotifyLogin(player *model.Player, enterType proto.EnterType) *proto.PlayerEnterSceneNotify {
	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
	if world == nil {
		logger.Error("get world is nil, worldId: %v", player.WorldId)
		return new(proto.PlayerEnterSceneNotify)
	}
	scene := world.GetSceneById(player.SceneId)
	enterSceneToken := world.AddEnterSceneContext(&EnterSceneContext{
		OldSceneId: 0,
		Uid:        player.PlayerID,
	})
	playerEnterSceneNotify := &proto.PlayerEnterSceneNotify{
		SceneId:                player.SceneId,
		Pos:                    &proto.Vector{X: float32(player.Pos.X), Y: float32(player.Pos.Y), Z: float32(player.Pos.Z)},
		SceneBeginTime:         uint64(scene.GetSceneCreateTime()),
		Type:                   enterType,
		TargetUid:              player.PlayerID,
		EnterSceneToken:        enterSceneToken,
		WorldLevel:             player.PropertiesMap[constant.PLAYER_PROP_PLAYER_WORLD_LEVEL],
		EnterReason:            uint32(proto.EnterReason_ENTER_REASON_LOGIN),
		IsFirstLoginEnterScene: true,
		WorldType:              1,
		SceneTagIdList:         make([]uint32, 0),
	}
	SceneTransactionSeq++
	playerEnterSceneNotify.SceneTransaction = strconv.Itoa(int(player.SceneId)) + "-" +
		strconv.Itoa(int(player.PlayerID)) + "-" +
		strconv.Itoa(int(time.Now().Unix())) + "-" +
		strconv.Itoa(int(SceneTransactionSeq))
	for _, sceneTagDataConfig := range gdconf.GetSceneTagDataMap() {
		if uint32(sceneTagDataConfig.SceneId) == player.SceneId {
			playerEnterSceneNotify.SceneTagIdList = append(playerEnterSceneNotify.SceneTagIdList, uint32(sceneTagDataConfig.SceneTagId))
		}
	}
	return playerEnterSceneNotify
}

func (g *Game) PacketPlayerEnterSceneNotifyTp(
	player *model.Player,
	enterType proto.EnterType,
	enterReason proto.EnterReason,
	prevSceneId uint32,
	prevPos *model.Vector,
	dungeonId uint32,
	dungeonPointId uint32,
) *proto.PlayerEnterSceneNotify {
	return g.PacketPlayerEnterSceneNotifyCore(player, player, enterType, enterReason, prevSceneId, prevPos, dungeonId, dungeonPointId)
}

func (g *Game) PacketPlayerEnterSceneNotifyMp(
	player *model.Player,
	targetPlayer *model.Player,
	enterType proto.EnterType,
	enterReason proto.EnterReason,
	prevSceneId uint32,
	prevPos *model.Vector,
) *proto.PlayerEnterSceneNotify {
	return g.PacketPlayerEnterSceneNotifyCore(player, targetPlayer, enterType, enterReason, prevSceneId, prevPos, 0, 0)
}

func (g *Game) PacketPlayerEnterSceneNotifyCore(
	player *model.Player,
	targetPlayer *model.Player,
	enterType proto.EnterType,
	enterReason proto.EnterReason,
	prevSceneId uint32,
	prevPos *model.Vector,
	dungeonId uint32,
	dungeonPointId uint32,
) *proto.PlayerEnterSceneNotify {
	world := WORLD_MANAGER.GetWorldByID(targetPlayer.WorldId)
	scene := world.GetSceneById(targetPlayer.SceneId)
	enterSceneToken := world.AddEnterSceneContext(&EnterSceneContext{
		OldSceneId: prevSceneId,
		OldPos: &model.Vector{
			X: prevPos.X,
			Y: prevPos.Y,
			Z: prevPos.Z,
		},
		OldDungeonPointId: dungeonPointId,
		Uid:               player.PlayerID,
	})
	playerEnterSceneNotify := &proto.PlayerEnterSceneNotify{
		PrevSceneId:     prevSceneId,
		PrevPos:         &proto.Vector{X: float32(prevPos.X), Y: float32(prevPos.Y), Z: float32(prevPos.Z)},
		SceneId:         targetPlayer.SceneId,
		Pos:             &proto.Vector{X: float32(targetPlayer.Pos.X), Y: float32(targetPlayer.Pos.Y), Z: float32(targetPlayer.Pos.Z)},
		SceneBeginTime:  uint64(scene.GetSceneCreateTime()),
		Type:            enterType,
		TargetUid:       targetPlayer.PlayerID,
		EnterSceneToken: enterSceneToken,
		WorldLevel:      targetPlayer.PropertiesMap[constant.PLAYER_PROP_PLAYER_WORLD_LEVEL],
		EnterReason:     uint32(enterReason),
		WorldType:       1,
		DungeonId:       dungeonId,
		SceneTagIdList:  make([]uint32, 0),
	}
	SceneTransactionSeq++
	playerEnterSceneNotify.SceneTransaction = strconv.Itoa(int(targetPlayer.SceneId)) + "-" +
		strconv.Itoa(int(player.PlayerID)) + "-" +
		strconv.Itoa(int(time.Now().Unix())) + "-" +
		strconv.Itoa(int(SceneTransactionSeq))
	for _, sceneTagDataConfig := range gdconf.GetSceneTagDataMap() {
		if uint32(sceneTagDataConfig.SceneId) == targetPlayer.SceneId {
			playerEnterSceneNotify.SceneTagIdList = append(playerEnterSceneNotify.SceneTagIdList, uint32(sceneTagDataConfig.SceneTagId))
		}
	}
	return playerEnterSceneNotify
}

func (g *Game) PacketFightPropMapToPbFightPropList(fightPropMap map[uint32]float32) []*proto.FightPropPair {
	fightPropList := []*proto.FightPropPair{
		{PropType: constant.FIGHT_PROP_BASE_HP, PropValue: fightPropMap[constant.FIGHT_PROP_BASE_HP]},
		{PropType: constant.FIGHT_PROP_BASE_ATTACK, PropValue: fightPropMap[constant.FIGHT_PROP_BASE_ATTACK]},
		{PropType: constant.FIGHT_PROP_BASE_DEFENSE, PropValue: fightPropMap[constant.FIGHT_PROP_BASE_DEFENSE]},
		{PropType: constant.FIGHT_PROP_CRITICAL, PropValue: fightPropMap[constant.FIGHT_PROP_CRITICAL]},
		{PropType: constant.FIGHT_PROP_CRITICAL_HURT, PropValue: fightPropMap[constant.FIGHT_PROP_CRITICAL_HURT]},
		{PropType: constant.FIGHT_PROP_CHARGE_EFFICIENCY, PropValue: fightPropMap[constant.FIGHT_PROP_CHARGE_EFFICIENCY]},
		{PropType: constant.FIGHT_PROP_CUR_HP, PropValue: fightPropMap[constant.FIGHT_PROP_CUR_HP]},
		{PropType: constant.FIGHT_PROP_MAX_HP, PropValue: fightPropMap[constant.FIGHT_PROP_MAX_HP]},
		{PropType: constant.FIGHT_PROP_CUR_ATTACK, PropValue: fightPropMap[constant.FIGHT_PROP_CUR_ATTACK]},
		{PropType: constant.FIGHT_PROP_CUR_DEFENSE, PropValue: fightPropMap[constant.FIGHT_PROP_CUR_DEFENSE]},
	}
	return fightPropList
}

func (g *Game) PacketSceneEntityInfoAvatar(scene *Scene, player *model.Player, avatarId uint32) *proto.SceneEntityInfo {
	entity := scene.GetEntity(scene.GetWorld().GetPlayerWorldAvatarEntityId(player, avatarId))
	if entity == nil {
		return new(proto.SceneEntityInfo)
	}
	pos := &proto.Vector{
		X: float32(entity.GetPos().X),
		Y: float32(entity.GetPos().Y),
		Z: float32(entity.GetPos().Z),
	}
	worldAvatar := scene.GetWorld().GetWorldAvatarByEntityId(entity.GetId())
	dbAvatar := player.GetDbAvatar()
	avatar, ok := dbAvatar.AvatarMap[worldAvatar.GetAvatarId()]
	if !ok {
		logger.Error("avatar error, avatarId: %v", worldAvatar.GetAvatarId())
		return new(proto.SceneEntityInfo)
	}
	sceneEntityInfo := &proto.SceneEntityInfo{
		EntityType: proto.ProtEntityType_PROT_ENTITY_AVATAR,
		EntityId:   entity.GetId(),
		MotionInfo: &proto.MotionInfo{
			Pos: pos,
			Rot: &proto.Vector{
				X: float32(entity.GetRot().X),
				Y: float32(entity.GetRot().Y),
				Z: float32(entity.GetRot().Z),
			},
			Speed: &proto.Vector{},
			State: proto.MotionState(entity.GetMoveState()),
		},
		PropList: []*proto.PropPair{
			{
				Type: uint32(constant.PLAYER_PROP_LEVEL),
				PropValue: &proto.PropValue{
					Type:  uint32(constant.PLAYER_PROP_LEVEL),
					Value: &proto.PropValue_Ival{Ival: int64(avatar.Level)},
					Val:   int64(avatar.Level)},
			},
			{
				Type: uint32(constant.PLAYER_PROP_EXP),
				PropValue: &proto.PropValue{
					Type:  uint32(constant.PLAYER_PROP_EXP),
					Value: &proto.PropValue_Ival{Ival: int64(avatar.Exp)},
					Val:   int64(avatar.Exp)},
			},
			{
				Type: uint32(constant.PLAYER_PROP_BREAK_LEVEL),
				PropValue: &proto.PropValue{
					Type:  uint32(constant.PLAYER_PROP_BREAK_LEVEL),
					Value: &proto.PropValue_Ival{Ival: int64(avatar.Promote)},
					Val:   int64(avatar.Promote)},
			},
			{
				Type: uint32(constant.PLAYER_PROP_SATIATION_VAL),
				PropValue: &proto.PropValue{
					Type:  uint32(constant.PLAYER_PROP_SATIATION_VAL),
					Value: &proto.PropValue_Ival{Ival: int64(avatar.Satiation)},
					Val:   int64(avatar.Satiation)},
			},
			{
				Type: uint32(constant.PLAYER_PROP_SATIATION_PENALTY_TIME),
				PropValue: &proto.PropValue{
					Type:  uint32(constant.PLAYER_PROP_SATIATION_PENALTY_TIME),
					Value: &proto.PropValue_Ival{Ival: int64(avatar.SatiationPenalty)},
					Val:   int64(avatar.SatiationPenalty)},
			},
		},
		FightPropList:    g.PacketFightPropMapToPbFightPropList(avatar.FightPropMap),
		LifeState:        uint32(avatar.LifeState),
		AnimatorParaList: make([]*proto.AnimatorParameterValueInfoPair, 0),
		Entity: &proto.SceneEntityInfo_Avatar{
			Avatar: g.PacketSceneAvatarInfo(scene, player, avatarId),
		},
		EntityClientData: new(proto.EntityClientData),
		EntityAuthorityInfo: &proto.EntityAuthorityInfo{
			AbilityInfo: &proto.AbilitySyncStateInfo{
				IsInited:           len(worldAvatar.GetAbilityList()) != 0,
				DynamicValueMap:    nil,
				AppliedAbilities:   worldAvatar.GetAbilityList(),
				AppliedModifiers:   worldAvatar.GetModifierList(),
				MixinRecoverInfos:  nil,
				SgvDynamicValueMap: nil,
			},
			RendererChangedInfo: new(proto.EntityRendererChangedInfo),
			AiInfo: &proto.SceneEntityAiInfo{
				IsAiOpen: true,
				BornPos:  pos,
			},
			BornPos: pos,
		},
		LastMoveSceneTimeMs: entity.GetLastMoveSceneTimeMs(),
		LastMoveReliableSeq: entity.GetLastMoveReliableSeq(),
	}
	return sceneEntityInfo
}

func (g *Game) PacketSceneEntityInfoMonster(scene *Scene, entityId uint32) *proto.SceneEntityInfo {
	entity := scene.GetEntity(entityId)
	if entity == nil {
		return new(proto.SceneEntityInfo)
	}
	pos := &proto.Vector{
		X: float32(entity.GetPos().X),
		Y: float32(entity.GetPos().Y),
		Z: float32(entity.GetPos().Z),
	}
	sceneEntityInfo := &proto.SceneEntityInfo{
		EntityType: proto.ProtEntityType_PROT_ENTITY_MONSTER,
		EntityId:   entity.GetId(),
		MotionInfo: &proto.MotionInfo{
			Pos: pos,
			Rot: &proto.Vector{
				X: float32(entity.GetRot().X),
				Y: float32(entity.GetRot().Y),
				Z: float32(entity.GetRot().Z),
			},
			Speed: &proto.Vector{},
			State: proto.MotionState(entity.GetMoveState()),
		},
		PropList: []*proto.PropPair{{Type: uint32(constant.PLAYER_PROP_LEVEL), PropValue: &proto.PropValue{
			Type:  uint32(constant.PLAYER_PROP_LEVEL),
			Value: &proto.PropValue_Ival{Ival: int64(entity.GetLevel())},
			Val:   int64(entity.GetLevel()),
		}}},
		FightPropList:    g.PacketFightPropMapToPbFightPropList(entity.GetFightProp()),
		LifeState:        uint32(entity.GetLifeState()),
		AnimatorParaList: make([]*proto.AnimatorParameterValueInfoPair, 0),
		Entity: &proto.SceneEntityInfo_Monster{
			Monster: g.PacketSceneMonsterInfo(entity),
		},
		EntityClientData: new(proto.EntityClientData),
		EntityAuthorityInfo: &proto.EntityAuthorityInfo{
			AbilityInfo:         new(proto.AbilitySyncStateInfo),
			RendererChangedInfo: new(proto.EntityRendererChangedInfo),
			AiInfo: &proto.SceneEntityAiInfo{
				IsAiOpen: true,
				BornPos:  pos,
			},
			BornPos: pos,
		},
	}
	return sceneEntityInfo
}

func (g *Game) PacketSceneEntityInfoNpc(scene *Scene, entityId uint32) *proto.SceneEntityInfo {
	entity := scene.GetEntity(entityId)
	if entity == nil {
		return new(proto.SceneEntityInfo)
	}
	pos := &proto.Vector{
		X: float32(entity.GetPos().X),
		Y: float32(entity.GetPos().Y),
		Z: float32(entity.GetPos().Z),
	}
	sceneEntityInfo := &proto.SceneEntityInfo{
		EntityType: proto.ProtEntityType_PROT_ENTITY_NPC,
		EntityId:   entity.GetId(),
		MotionInfo: &proto.MotionInfo{
			Pos: pos,
			Rot: &proto.Vector{
				X: float32(entity.GetRot().X),
				Y: float32(entity.GetRot().Y),
				Z: float32(entity.GetRot().Z),
			},
			Speed: &proto.Vector{},
			State: proto.MotionState(entity.GetMoveState()),
		},
		PropList: []*proto.PropPair{{Type: uint32(constant.PLAYER_PROP_LEVEL), PropValue: &proto.PropValue{
			Type:  uint32(constant.PLAYER_PROP_LEVEL),
			Value: &proto.PropValue_Ival{Ival: int64(entity.GetLevel())},
			Val:   int64(entity.GetLevel()),
		}}},
		FightPropList:    g.PacketFightPropMapToPbFightPropList(entity.GetFightProp()),
		LifeState:        uint32(entity.GetLifeState()),
		AnimatorParaList: make([]*proto.AnimatorParameterValueInfoPair, 0),
		Entity: &proto.SceneEntityInfo_Npc{
			Npc: g.PacketSceneNpcInfo(entity.GetNpcEntity()),
		},
		EntityClientData: new(proto.EntityClientData),
		EntityAuthorityInfo: &proto.EntityAuthorityInfo{
			AbilityInfo:         new(proto.AbilitySyncStateInfo),
			RendererChangedInfo: new(proto.EntityRendererChangedInfo),
			AiInfo: &proto.SceneEntityAiInfo{
				IsAiOpen: true,
				BornPos:  pos,
			},
			BornPos: pos,
		},
	}
	return sceneEntityInfo
}

func (g *Game) PacketSceneEntityInfoGadget(player *model.Player, scene *Scene, entityId uint32) *proto.SceneEntityInfo {
	entity := scene.GetEntity(entityId)
	if entity == nil {
		return new(proto.SceneEntityInfo)
	}
	pos := &proto.Vector{
		X: float32(entity.GetPos().X),
		Y: float32(entity.GetPos().Y),
		Z: float32(entity.GetPos().Z),
	}
	sceneEntityInfo := &proto.SceneEntityInfo{
		EntityType: proto.ProtEntityType_PROT_ENTITY_GADGET,
		EntityId:   entity.GetId(),
		MotionInfo: &proto.MotionInfo{
			Pos: pos,
			Rot: &proto.Vector{
				X: float32(entity.GetRot().X),
				Y: float32(entity.GetRot().Y),
				Z: float32(entity.GetRot().Z),
			},
			Speed: &proto.Vector{},
			State: proto.MotionState(entity.GetMoveState()),
		},
		PropList: []*proto.PropPair{{Type: uint32(constant.PLAYER_PROP_LEVEL), PropValue: &proto.PropValue{
			Type:  uint32(constant.PLAYER_PROP_LEVEL),
			Value: &proto.PropValue_Ival{Ival: int64(1)},
			Val:   int64(1),
		}}},
		FightPropList:    g.PacketFightPropMapToPbFightPropList(entity.GetFightProp()),
		LifeState:        uint32(entity.GetLifeState()),
		AnimatorParaList: make([]*proto.AnimatorParameterValueInfoPair, 0),
		EntityClientData: new(proto.EntityClientData),
		EntityAuthorityInfo: &proto.EntityAuthorityInfo{
			AbilityInfo:         new(proto.AbilitySyncStateInfo),
			RendererChangedInfo: new(proto.EntityRendererChangedInfo),
			AiInfo: &proto.SceneEntityAiInfo{
				IsAiOpen: true,
				BornPos:  pos,
			},
			BornPos: pos,
		},
	}
	gadgetEntity := entity.GetGadgetEntity()
	switch gadgetEntity.GetGadgetType() {
	case GADGET_TYPE_NORMAL:
		sceneEntityInfo.Entity = &proto.SceneEntityInfo_Gadget{
			Gadget: g.PacketSceneGadgetInfoNormal(player, entity),
		}
	case GADGET_TYPE_CLIENT:
		sceneEntityInfo.Entity = &proto.SceneEntityInfo_Gadget{
			Gadget: g.PacketSceneGadgetInfoClient(gadgetEntity.GetGadgetClientEntity()),
		}
	case GADGET_TYPE_VEHICLE:
		sceneEntityInfo.Entity = &proto.SceneEntityInfo_Gadget{
			Gadget: g.PacketSceneGadgetInfoVehicle(gadgetEntity.GetGadgetVehicleEntity()),
		}
	}
	return sceneEntityInfo
}

func (g *Game) PacketSceneAvatarInfo(scene *Scene, player *model.Player, avatarId uint32) *proto.SceneAvatarInfo {
	dbAvatar := player.GetDbAvatar()
	avatar, ok := dbAvatar.AvatarMap[avatarId]
	if !ok {
		logger.Error("avatar error, avatarId: %v", avatarId)
		return new(proto.SceneAvatarInfo)
	}
	equipIdList := make([]uint32, len(avatar.EquipReliquaryMap)+1)
	for _, reliquary := range avatar.EquipReliquaryMap {
		equipIdList = append(equipIdList, reliquary.ItemId)
	}
	equipIdList = append(equipIdList, avatar.EquipWeapon.ItemId)
	reliquaryList := make([]*proto.SceneReliquaryInfo, 0, len(avatar.EquipReliquaryMap))
	for _, reliquary := range avatar.EquipReliquaryMap {
		reliquaryList = append(reliquaryList, &proto.SceneReliquaryInfo{
			ItemId:       reliquary.ItemId,
			Guid:         reliquary.Guid,
			Level:        uint32(reliquary.Level),
			PromoteLevel: uint32(reliquary.Promote),
		})
	}
	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
	sceneAvatarInfo := &proto.SceneAvatarInfo{
		Uid:          player.PlayerID,
		AvatarId:     avatarId,
		Guid:         avatar.Guid,
		PeerId:       world.GetPlayerPeerId(player),
		EquipIdList:  equipIdList,
		SkillDepotId: avatar.SkillDepotId,
		Weapon: &proto.SceneWeaponInfo{
			EntityId:    scene.GetWorld().GetPlayerWorldAvatarWeaponEntityId(player, avatarId),
			GadgetId:    uint32(gdconf.GetItemDataById(int32(avatar.EquipWeapon.ItemId)).GadgetId),
			ItemId:      avatar.EquipWeapon.ItemId,
			Guid:        avatar.EquipWeapon.Guid,
			Level:       uint32(avatar.EquipWeapon.Level),
			AbilityInfo: new(proto.AbilitySyncStateInfo),
		},
		ReliquaryList:     reliquaryList,
		SkillLevelMap:     avatar.SkillLevelMap,
		WearingFlycloakId: avatar.FlyCloak,
		CostumeId:         avatar.Costume,
		BornTime:          uint32(avatar.BornTime),
		TeamResonanceList: make([]uint32, 0),
	}
	// for id := range player.TeamConfig.TeamResonances {
	//	sceneAvatarInfo.TeamResonanceList = append(sceneAvatarInfo.TeamResonanceList, uint32(id))
	// }
	return sceneAvatarInfo
}

func (g *Game) PacketSceneMonsterInfo(entity *Entity) *proto.SceneMonsterInfo {
	sceneMonsterInfo := &proto.SceneMonsterInfo{
		MonsterId:       entity.GetMonsterEntity().GetMonsterId(),
		AuthorityPeerId: 1,
		BornType:        proto.MonsterBornType_MONSTER_BORN_DEFAULT,
		// BlockId:         3001,
		// TitleId:         3001,
		// SpecialNameId:   40,
	}
	return sceneMonsterInfo
}

func (g *Game) PacketSceneNpcInfo(entity *NpcEntity) *proto.SceneNpcInfo {
	sceneNpcInfo := &proto.SceneNpcInfo{
		NpcId:         entity.NpcId,
		RoomId:        entity.RoomId,
		ParentQuestId: entity.ParentQuestId,
		BlockId:       entity.BlockId,
	}
	return sceneNpcInfo
}

func (g *Game) PacketSceneGadgetInfoNormal(player *model.Player, entity *Entity) *proto.SceneGadgetInfo {
	gadgetEntity := entity.GetGadgetEntity()
	gadgetDataConfig := gdconf.GetGadgetDataById(int32(gadgetEntity.GetGadgetId()))
	if gadgetDataConfig == nil {
		logger.Error("get gadget data config is nil, gadgetId: %v", gadgetEntity.GetGadgetId())
		return new(proto.SceneGadgetInfo)
	}
	sceneGadgetInfo := &proto.SceneGadgetInfo{
		GadgetId:         gadgetEntity.GetGadgetId(),
		GroupId:          entity.GetGroupId(),
		ConfigId:         entity.GetConfigId(),
		GadgetState:      gadgetEntity.GetGadgetState(),
		IsEnableInteract: true,
		AuthorityPeerId:  1,
	}
	gadgetNormalEntity := gadgetEntity.GetGadgetNormalEntity()
	if gadgetNormalEntity.GetIsDrop() {
		dbItem := player.GetDbItem()
		sceneGadgetInfo.Content = &proto.SceneGadgetInfo_TrifleItem{
			TrifleItem: &proto.Item{
				ItemId: gadgetNormalEntity.GetItemId(),
				Guid:   dbItem.GetItemGuid(gadgetNormalEntity.GetItemId()),
				Detail: &proto.Item_Material{
					Material: &proto.Material{
						Count: gadgetNormalEntity.GetCount(),
					},
				},
			},
		}
	} else if gadgetDataConfig.Type == constant.GADGET_TYPE_GATHER_OBJECT {
		sceneGadgetInfo.Content = &proto.SceneGadgetInfo_GatherGadget{
			GatherGadget: &proto.GatherGadgetInfo{
				ItemId:        gadgetNormalEntity.GetItemId(),
				IsForbidGuest: false,
			},
		}
	}
	return sceneGadgetInfo
}

func (g *Game) PacketSceneGadgetInfoClient(gadgetClientEntity *GadgetClientEntity) *proto.SceneGadgetInfo {
	sceneGadgetInfo := &proto.SceneGadgetInfo{
		GadgetId:         gadgetClientEntity.GetConfigId(),
		OwnerEntityId:    gadgetClientEntity.GetOwnerEntityId(),
		AuthorityPeerId:  1,
		IsEnableInteract: true,
		Content: &proto.SceneGadgetInfo_ClientGadget{
			ClientGadget: &proto.ClientGadgetInfo{
				CampId:         gadgetClientEntity.GetCampId(),
				CampType:       gadgetClientEntity.GetCampType(),
				OwnerEntityId:  gadgetClientEntity.GetOwnerEntityId(),
				TargetEntityId: gadgetClientEntity.GetTargetEntityId(),
			},
		},
		PropOwnerEntityId: gadgetClientEntity.GetPropOwnerEntityId(),
	}
	return sceneGadgetInfo
}

func (g *Game) PacketSceneGadgetInfoVehicle(gadgetVehicleEntity *GadgetVehicleEntity) *proto.SceneGadgetInfo {
	sceneGadgetInfo := &proto.SceneGadgetInfo{
		GadgetId:         gadgetVehicleEntity.GetVehicleId(),
		AuthorityPeerId:  WORLD_MANAGER.GetWorldByID(gadgetVehicleEntity.GetOwner().WorldId).GetPlayerPeerId(gadgetVehicleEntity.GetOwner()),
		IsEnableInteract: true,
		Content: &proto.SceneGadgetInfo_VehicleInfo{
			VehicleInfo: &proto.VehicleInfo{
				MemberList: make([]*proto.VehicleMember, 0, len(gadgetVehicleEntity.GetMemberMap())),
				OwnerUid:   gadgetVehicleEntity.GetOwner().PlayerID,
				CurStamina: gadgetVehicleEntity.GetCurStamina(),
			},
		},
	}
	return sceneGadgetInfo
}

func (g *Game) PacketDelTeamEntityNotify(scene *Scene, player *model.Player) *proto.DelTeamEntityNotify {
	delTeamEntityNotify := &proto.DelTeamEntityNotify{
		SceneId:         player.SceneId,
		DelEntityIdList: []uint32{scene.GetWorld().GetPlayerTeamEntityId(player)},
	}
	return delTeamEntityNotify
}
