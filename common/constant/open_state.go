package constant

const (
	OPEN_STATE_NONE                                    uint16 = 0
	OPEN_STATE_PAIMON                                  uint16 = 1
	OPEN_STATE_PAIMON_NAVIGATION                       uint16 = 2
	OPEN_STATE_AVATAR_PROMOTE                          uint16 = 3
	OPEN_STATE_AVATAR_TALENT                           uint16 = 4
	OPEN_STATE_WEAPON_PROMOTE                          uint16 = 5
	OPEN_STATE_WEAPON_AWAKEN                           uint16 = 6
	OPEN_STATE_QUEST_REMIND                            uint16 = 7
	OPEN_STATE_GAME_GUIDE                              uint16 = 8
	OPEN_STATE_COOK                                    uint16 = 9
	OPEN_STATE_WEAPON_UPGRADE                          uint16 = 10
	OPEN_STATE_RELIQUARY_UPGRADE                       uint16 = 11
	OPEN_STATE_RELIQUARY_PROMOTE                       uint16 = 12
	OPEN_STATE_WEAPON_PROMOTE_GUIDE                    uint16 = 13
	OPEN_STATE_WEAPON_CHANGE_GUIDE                     uint16 = 14
	OPEN_STATE_PLAYER_LVUP_GUIDE                       uint16 = 15
	OPEN_STATE_FRESHMAN_GUIDE                          uint16 = 16
	OPEN_STATE_SKIP_FRESHMAN_GUIDE                     uint16 = 17
	OPEN_STATE_GUIDE_MOVE_CAMERA                       uint16 = 18
	OPEN_STATE_GUIDE_SCALE_CAMERA                      uint16 = 19
	OPEN_STATE_GUIDE_KEYBOARD                          uint16 = 20
	OPEN_STATE_GUIDE_MOVE                              uint16 = 21
	OPEN_STATE_GUIDE_JUMP                              uint16 = 22
	OPEN_STATE_GUIDE_SPRINT                            uint16 = 23
	OPEN_STATE_GUIDE_MAP                               uint16 = 24
	OPEN_STATE_GUIDE_ATTACK                            uint16 = 25
	OPEN_STATE_GUIDE_FLY                               uint16 = 26
	OPEN_STATE_GUIDE_TALENT                            uint16 = 27
	OPEN_STATE_GUIDE_RELIC                             uint16 = 28
	OPEN_STATE_GUIDE_RELIC_PROM                        uint16 = 29
	OPEN_STATE_COMBINE                                 uint16 = 30
	OPEN_STATE_GACHA                                   uint16 = 31
	OPEN_STATE_GUIDE_GACHA                             uint16 = 32
	OPEN_STATE_GUIDE_TEAM                              uint16 = 33
	OPEN_STATE_GUIDE_PROUD                             uint16 = 34
	OPEN_STATE_GUIDE_AVATAR_PROMOTE                    uint16 = 35
	OPEN_STATE_GUIDE_ADVENTURE_CARD                    uint16 = 36
	OPEN_STATE_FORGE                                   uint16 = 37
	OPEN_STATE_GUIDE_BAG                               uint16 = 38
	OPEN_STATE_EXPEDITION                              uint16 = 39
	OPEN_STATE_GUIDE_ADVENTURE_DAILYTASK               uint16 = 40
	OPEN_STATE_GUIDE_ADVENTURE_DUNGEON                 uint16 = 41
	OPEN_STATE_TOWER                                   uint16 = 42
	OPEN_STATE_WORLD_STAMINA                           uint16 = 43
	OPEN_STATE_TOWER_FIRST_ENTER                       uint16 = 44
	OPEN_STATE_RESIN                                   uint16 = 45
	OPEN_STATE_LIMIT_REGION_FRESHMEAT                  uint16 = 47
	OPEN_STATE_LIMIT_REGION_GLOBAL                     uint16 = 48
	OPEN_STATE_MULTIPLAYER                             uint16 = 49
	OPEN_STATE_GUIDE_MOUSEPC                           uint16 = 50
	OPEN_STATE_GUIDE_MULTIPLAYER                       uint16 = 51
	OPEN_STATE_GUIDE_DUNGEONREWARD                     uint16 = 52
	OPEN_STATE_GUIDE_BLOSSOM                           uint16 = 53
	OPEN_STATE_AVATAR_FASHION                          uint16 = 54
	OPEN_STATE_PHOTOGRAPH                              uint16 = 55
	OPEN_STATE_GUIDE_KSLQUEST                          uint16 = 56
	OPEN_STATE_PERSONAL_LINE                           uint16 = 57
	OPEN_STATE_GUIDE_PERSONAL_LINE                     uint16 = 58
	OPEN_STATE_GUIDE_APPEARANCE                        uint16 = 59
	OPEN_STATE_GUIDE_PROCESS                           uint16 = 60
	OPEN_STATE_GUIDE_PERSONAL_LINE_KEY                 uint16 = 61
	OPEN_STATE_GUIDE_WIDGET                            uint16 = 62
	OPEN_STATE_GUIDE_ACTIVITY_SKILL_ASTER              uint16 = 63
	OPEN_STATE_GUIDE_COLDCLIMATE                       uint16 = 64
	OPEN_STATE_DERIVATIVE_MALL                         uint16 = 65
	OPEN_STATE_GUIDE_EXITMULTIPLAYER                   uint16 = 66
	OPEN_STATE_GUIDE_THEATREMACHANICUS_BUILD           uint16 = 67
	OPEN_STATE_GUIDE_THEATREMACHANICUS_REBUILD         uint16 = 68
	OPEN_STATE_GUIDE_THEATREMACHANICUS_CARD            uint16 = 69
	OPEN_STATE_GUIDE_THEATREMACHANICUS_MONSTER         uint16 = 70
	OPEN_STATE_GUIDE_THEATREMACHANICUS_MISSION_CHECK   uint16 = 71
	OPEN_STATE_GUIDE_THEATREMACHANICUS_BUILD_SELECT    uint16 = 72
	OPEN_STATE_GUIDE_THEATREMACHANICUS_CHALLENGE_START uint16 = 73
	OPEN_STATE_GUIDE_CONVERT                           uint16 = 74
	OPEN_STATE_GUIDE_THEATREMACHANICUS_MULTIPLAYER     uint16 = 75
	OPEN_STATE_GUIDE_COOP_TASK                         uint16 = 76
	OPEN_STATE_GUIDE_HOMEWORLD_ADEPTIABODE             uint16 = 77
	OPEN_STATE_GUIDE_HOMEWORLD_DEPLOY                  uint16 = 78
	OPEN_STATE_GUIDE_CHANNELLERSLAB_EQUIP              uint16 = 79
	OPEN_STATE_GUIDE_CHANNELLERSLAB_MP_SOLUTION        uint16 = 80
	OPEN_STATE_GUIDE_CHANNELLERSLAB_POWER              uint16 = 81
	OPEN_STATE_GUIDE_HIDEANDSEEK_SKILL                 uint16 = 82
	OPEN_STATE_GUIDE_HOMEWORLD_MAPLIST                 uint16 = 83
	OPEN_STATE_GUIDE_RELICRESOLVE                      uint16 = 84
	OPEN_STATE_GUIDE_GGUIDE                            uint16 = 85
	OPEN_STATE_GUIDE_GGUIDE_HINT                       uint16 = 86
	OPEN_STATE_CITY_REPUATION_MENGDE                   uint16 = 800
	OPEN_STATE_CITY_REPUATION_LIYUE                    uint16 = 801
	OPEN_STATE_CITY_REPUATION_UI_HINT                  uint16 = 802
	OPEN_STATE_CITY_REPUATION_INAZUMA                  uint16 = 803
	OPEN_STATE_SHOP_TYPE_MALL                          uint16 = 900
	OPEN_STATE_SHOP_TYPE_RECOMMANDED                   uint16 = 901
	OPEN_STATE_SHOP_TYPE_GENESISCRYSTAL                uint16 = 902
	OPEN_STATE_SHOP_TYPE_GIFTPACKAGE                   uint16 = 903
	OPEN_STATE_SHOP_TYPE_PAIMON                        uint16 = 1001
	OPEN_STATE_SHOP_TYPE_CITY                          uint16 = 1002
	OPEN_STATE_SHOP_TYPE_BLACKSMITH                    uint16 = 1003
	OPEN_STATE_SHOP_TYPE_GROCERY                       uint16 = 1004
	OPEN_STATE_SHOP_TYPE_FOOD                          uint16 = 1005
	OPEN_STATE_SHOP_TYPE_SEA_LAMP                      uint16 = 1006
	OPEN_STATE_SHOP_TYPE_VIRTUAL_SHOP                  uint16 = 1007
	OPEN_STATE_SHOP_TYPE_LIYUE_GROCERY                 uint16 = 1008
	OPEN_STATE_SHOP_TYPE_LIYUE_SOUVENIR                uint16 = 1009
	OPEN_STATE_SHOP_TYPE_LIYUE_RESTAURANT              uint16 = 1010
	OPEN_STATE_SHOP_TYPE_INAZUMA_SOUVENIR              uint16 = 1011
	OPEN_STATE_SHOP_TYPE_NPC_TOMOKI                    uint16 = 1012
	OPEN_ADVENTURE_MANUAL                              uint16 = 1100
	OPEN_ADVENTURE_MANUAL_CITY_MENGDE                  uint16 = 1101
	OPEN_ADVENTURE_MANUAL_CITY_LIYUE                   uint16 = 1102
	OPEN_ADVENTURE_MANUAL_MONSTER                      uint16 = 1103
	OPEN_ADVENTURE_MANUAL_BOSS_DUNGEON                 uint16 = 1104
	OPEN_STATE_ACTIVITY_SEALAMP                        uint16 = 1200
	OPEN_STATE_ACTIVITY_SEALAMP_TAB2                   uint16 = 1201
	OPEN_STATE_ACTIVITY_SEALAMP_TAB3                   uint16 = 1202
	OPEN_STATE_BATTLE_PASS                             uint16 = 1300
	OPEN_STATE_BATTLE_PASS_ENTRY                       uint16 = 1301
	OPEN_STATE_ACTIVITY_CRUCIBLE                       uint16 = 1400
	OPEN_STATE_ACTIVITY_NEWBEEBOUNS_OPEN               uint16 = 1401
	OPEN_STATE_ACTIVITY_NEWBEEBOUNS_CLOSE              uint16 = 1402
	OPEN_STATE_ACTIVITY_ENTRY_OPEN                     uint16 = 1403
	OPEN_STATE_MENGDE_INFUSEDCRYSTAL                   uint16 = 1404
	OPEN_STATE_LIYUE_INFUSEDCRYSTAL                    uint16 = 1405
	OPEN_STATE_SNOW_MOUNTAIN_ELDER_TREE                uint16 = 1406
	OPEN_STATE_MIRACLE_RING                            uint16 = 1407
	OPEN_STATE_COOP_LINE                               uint16 = 1408
	OPEN_STATE_INAZUMA_INFUSEDCRYSTAL                  uint16 = 1409
	OPEN_STATE_FISH                                    uint16 = 1410
	OPEN_STATE_GUIDE_SUMO_TEAM_SKILL                   uint16 = 1411
	OPEN_STATE_GUIDE_FISH_RECIPE                       uint16 = 1412
	OPEN_STATE_HOME                                    uint16 = 1500
	OPEN_STATE_ACTIVITY_HOMEWORLD                      uint16 = 1501
	OPEN_STATE_ADEPTIABODE                             uint16 = 1502
	OPEN_STATE_HOME_AVATAR                             uint16 = 1503
	OPEN_STATE_HOME_EDIT                               uint16 = 1504
	OPEN_STATE_HOME_EDIT_TIPS                          uint16 = 1505
	OPEN_STATE_RELIQUARY_DECOMPOSE                     uint16 = 1600
	OPEN_STATE_ACTIVITY_H5                             uint16 = 1700
	OPEN_STATE_ORAIONOKAMI                             uint16 = 2000
	OPEN_STATE_GUIDE_CHESS_MISSION_CHECK               uint16 = 2001
	OPEN_STATE_GUIDE_CHESS_BUILD                       uint16 = 2002
	OPEN_STATE_GUIDE_CHESS_WIND_TOWER_CIRCLE           uint16 = 2003
	OPEN_STATE_GUIDE_CHESS_CARD_SELECT                 uint16 = 2004
	OPEN_STATE_INAZUMA_MAINQUEST_FINISHED              uint16 = 2005
	OPEN_STATE_PAIMON_LVINFO                           uint16 = 2100
	OPEN_STATE_TELEPORT_HUD                            uint16 = 2101
	OPEN_STATE_GUIDE_MAP_UNLOCK                        uint16 = 2102
	OPEN_STATE_GUIDE_PAIMON_LVINFO                     uint16 = 2103
	OPEN_STATE_GUIDE_AMBORTRANSPORT                    uint16 = 2104
	OPEN_STATE_GUIDE_FLY_SECOND                        uint16 = 2105
	OPEN_STATE_GUIDE_KAEYA_CLUE                        uint16 = 2106
	OPEN_STATE_CAPTURE_CODEX                           uint16 = 2107
	OPEN_STATE_ACTIVITY_FISH_OPEN                      uint16 = 2200
	OPEN_STATE_ACTIVITY_FISH_CLOSE                     uint16 = 2201
	OPEN_STATE_GUIDE_ROGUE_MAP                         uint16 = 2205
	OPEN_STATE_GUIDE_ROGUE_RUNE                        uint16 = 2206
	OPEN_STATE_GUIDE_BARTENDER_FORMULA                 uint16 = 2210
	OPEN_STATE_GUIDE_BARTENDER_MIX                     uint16 = 2211
	OPEN_STATE_GUIDE_BARTENDER_CUP                     uint16 = 2212
	OPEN_STATE_GUIDE_MAIL_FAVORITES                    uint16 = 2400
	OPEN_STATE_GUIDE_POTION_CONFIGURE                  uint16 = 2401
	OPEN_STATE_GUIDE_LANV2_FIREWORK                    uint16 = 2402
	OPEN_STATE_LOADINGTIPS_ENKANOMIYA                  uint16 = 2403
	OPEN_STATE_MICHIAE_CASKET                          uint16 = 2500
	OPEN_STATE_MAIL_COLLECT_UNLOCK_RED_POINT           uint16 = 2501
	OPEN_STATE_LUMEN_STONE                             uint16 = 2600
	OPEN_STATE_GUIDE_CRYSTALLINK_BUFF                  uint16 = 2601
)

var ALL_OPEN_STATE []uint16

func init() {
	ALL_OPEN_STATE = []uint16{
		OPEN_STATE_NONE,
		OPEN_STATE_PAIMON,
		OPEN_STATE_PAIMON_NAVIGATION,
		OPEN_STATE_AVATAR_PROMOTE,
		OPEN_STATE_AVATAR_TALENT,
		OPEN_STATE_WEAPON_PROMOTE,
		OPEN_STATE_WEAPON_AWAKEN,
		OPEN_STATE_QUEST_REMIND,
		OPEN_STATE_GAME_GUIDE,
		OPEN_STATE_COOK,
		OPEN_STATE_WEAPON_UPGRADE,
		OPEN_STATE_RELIQUARY_UPGRADE,
		OPEN_STATE_RELIQUARY_PROMOTE,
		OPEN_STATE_WEAPON_PROMOTE_GUIDE,
		OPEN_STATE_WEAPON_CHANGE_GUIDE,
		OPEN_STATE_PLAYER_LVUP_GUIDE,
		OPEN_STATE_FRESHMAN_GUIDE,
		OPEN_STATE_SKIP_FRESHMAN_GUIDE,
		OPEN_STATE_GUIDE_MOVE_CAMERA,
		OPEN_STATE_GUIDE_SCALE_CAMERA,
		OPEN_STATE_GUIDE_KEYBOARD,
		OPEN_STATE_GUIDE_MOVE,
		OPEN_STATE_GUIDE_JUMP,
		OPEN_STATE_GUIDE_SPRINT,
		OPEN_STATE_GUIDE_MAP,
		OPEN_STATE_GUIDE_ATTACK,
		OPEN_STATE_GUIDE_FLY,
		OPEN_STATE_GUIDE_TALENT,
		OPEN_STATE_GUIDE_RELIC,
		OPEN_STATE_GUIDE_RELIC_PROM,
		OPEN_STATE_COMBINE,
		OPEN_STATE_GACHA,
		OPEN_STATE_GUIDE_GACHA,
		OPEN_STATE_GUIDE_TEAM,
		OPEN_STATE_GUIDE_PROUD,
		OPEN_STATE_GUIDE_AVATAR_PROMOTE,
		OPEN_STATE_GUIDE_ADVENTURE_CARD,
		OPEN_STATE_FORGE,
		OPEN_STATE_GUIDE_BAG,
		OPEN_STATE_EXPEDITION,
		OPEN_STATE_GUIDE_ADVENTURE_DAILYTASK,
		OPEN_STATE_GUIDE_ADVENTURE_DUNGEON,
		OPEN_STATE_TOWER,
		OPEN_STATE_WORLD_STAMINA,
		OPEN_STATE_TOWER_FIRST_ENTER,
		OPEN_STATE_RESIN,
		OPEN_STATE_LIMIT_REGION_FRESHMEAT,
		OPEN_STATE_LIMIT_REGION_GLOBAL,
		OPEN_STATE_MULTIPLAYER,
		OPEN_STATE_GUIDE_MOUSEPC,
		OPEN_STATE_GUIDE_MULTIPLAYER,
		OPEN_STATE_GUIDE_DUNGEONREWARD,
		OPEN_STATE_GUIDE_BLOSSOM,
		OPEN_STATE_AVATAR_FASHION,
		OPEN_STATE_PHOTOGRAPH,
		OPEN_STATE_GUIDE_KSLQUEST,
		OPEN_STATE_PERSONAL_LINE,
		OPEN_STATE_GUIDE_PERSONAL_LINE,
		OPEN_STATE_GUIDE_APPEARANCE,
		OPEN_STATE_GUIDE_PROCESS,
		OPEN_STATE_GUIDE_PERSONAL_LINE_KEY,
		OPEN_STATE_GUIDE_WIDGET,
		OPEN_STATE_GUIDE_ACTIVITY_SKILL_ASTER,
		OPEN_STATE_GUIDE_COLDCLIMATE,
		OPEN_STATE_DERIVATIVE_MALL,
		OPEN_STATE_GUIDE_EXITMULTIPLAYER,
		OPEN_STATE_GUIDE_THEATREMACHANICUS_BUILD,
		OPEN_STATE_GUIDE_THEATREMACHANICUS_REBUILD,
		OPEN_STATE_GUIDE_THEATREMACHANICUS_CARD,
		OPEN_STATE_GUIDE_THEATREMACHANICUS_MONSTER,
		OPEN_STATE_GUIDE_THEATREMACHANICUS_MISSION_CHECK,
		OPEN_STATE_GUIDE_THEATREMACHANICUS_BUILD_SELECT,
		OPEN_STATE_GUIDE_THEATREMACHANICUS_CHALLENGE_START,
		OPEN_STATE_GUIDE_CONVERT,
		OPEN_STATE_GUIDE_THEATREMACHANICUS_MULTIPLAYER,
		OPEN_STATE_GUIDE_COOP_TASK,
		OPEN_STATE_GUIDE_HOMEWORLD_ADEPTIABODE,
		OPEN_STATE_GUIDE_HOMEWORLD_DEPLOY,
		OPEN_STATE_GUIDE_CHANNELLERSLAB_EQUIP,
		OPEN_STATE_GUIDE_CHANNELLERSLAB_MP_SOLUTION,
		OPEN_STATE_GUIDE_CHANNELLERSLAB_POWER,
		OPEN_STATE_GUIDE_HIDEANDSEEK_SKILL,
		OPEN_STATE_GUIDE_HOMEWORLD_MAPLIST,
		OPEN_STATE_GUIDE_RELICRESOLVE,
		OPEN_STATE_GUIDE_GGUIDE,
		OPEN_STATE_GUIDE_GGUIDE_HINT,
		OPEN_STATE_CITY_REPUATION_MENGDE,
		OPEN_STATE_CITY_REPUATION_LIYUE,
		OPEN_STATE_CITY_REPUATION_UI_HINT,
		OPEN_STATE_CITY_REPUATION_INAZUMA,
		OPEN_STATE_SHOP_TYPE_MALL,
		OPEN_STATE_SHOP_TYPE_RECOMMANDED,
		OPEN_STATE_SHOP_TYPE_GENESISCRYSTAL,
		OPEN_STATE_SHOP_TYPE_GIFTPACKAGE,
		OPEN_STATE_SHOP_TYPE_PAIMON,
		OPEN_STATE_SHOP_TYPE_CITY,
		OPEN_STATE_SHOP_TYPE_BLACKSMITH,
		OPEN_STATE_SHOP_TYPE_GROCERY,
		OPEN_STATE_SHOP_TYPE_FOOD,
		OPEN_STATE_SHOP_TYPE_SEA_LAMP,
		OPEN_STATE_SHOP_TYPE_VIRTUAL_SHOP,
		OPEN_STATE_SHOP_TYPE_LIYUE_GROCERY,
		OPEN_STATE_SHOP_TYPE_LIYUE_SOUVENIR,
		OPEN_STATE_SHOP_TYPE_LIYUE_RESTAURANT,
		OPEN_STATE_SHOP_TYPE_INAZUMA_SOUVENIR,
		OPEN_STATE_SHOP_TYPE_NPC_TOMOKI,
		OPEN_ADVENTURE_MANUAL,
		OPEN_ADVENTURE_MANUAL_CITY_MENGDE,
		OPEN_ADVENTURE_MANUAL_CITY_LIYUE,
		OPEN_ADVENTURE_MANUAL_MONSTER,
		OPEN_ADVENTURE_MANUAL_BOSS_DUNGEON,
		OPEN_STATE_ACTIVITY_SEALAMP,
		OPEN_STATE_ACTIVITY_SEALAMP_TAB2,
		OPEN_STATE_ACTIVITY_SEALAMP_TAB3,
		OPEN_STATE_BATTLE_PASS,
		OPEN_STATE_BATTLE_PASS_ENTRY,
		OPEN_STATE_ACTIVITY_CRUCIBLE,
		OPEN_STATE_ACTIVITY_NEWBEEBOUNS_OPEN,
		OPEN_STATE_ACTIVITY_NEWBEEBOUNS_CLOSE,
		OPEN_STATE_ACTIVITY_ENTRY_OPEN,
		OPEN_STATE_MENGDE_INFUSEDCRYSTAL,
		OPEN_STATE_LIYUE_INFUSEDCRYSTAL,
		OPEN_STATE_SNOW_MOUNTAIN_ELDER_TREE,
		OPEN_STATE_MIRACLE_RING,
		OPEN_STATE_COOP_LINE,
		OPEN_STATE_INAZUMA_INFUSEDCRYSTAL,
		OPEN_STATE_FISH,
		OPEN_STATE_GUIDE_SUMO_TEAM_SKILL,
		OPEN_STATE_GUIDE_FISH_RECIPE,
		OPEN_STATE_HOME,
		OPEN_STATE_ACTIVITY_HOMEWORLD,
		OPEN_STATE_ADEPTIABODE,
		OPEN_STATE_HOME_AVATAR,
		OPEN_STATE_HOME_EDIT,
		OPEN_STATE_HOME_EDIT_TIPS,
		OPEN_STATE_RELIQUARY_DECOMPOSE,
		OPEN_STATE_ACTIVITY_H5,
		OPEN_STATE_ORAIONOKAMI,
		OPEN_STATE_GUIDE_CHESS_MISSION_CHECK,
		OPEN_STATE_GUIDE_CHESS_BUILD,
		OPEN_STATE_GUIDE_CHESS_WIND_TOWER_CIRCLE,
		OPEN_STATE_GUIDE_CHESS_CARD_SELECT,
		OPEN_STATE_INAZUMA_MAINQUEST_FINISHED,
		OPEN_STATE_PAIMON_LVINFO,
		OPEN_STATE_TELEPORT_HUD,
		OPEN_STATE_GUIDE_MAP_UNLOCK,
		OPEN_STATE_GUIDE_PAIMON_LVINFO,
		OPEN_STATE_GUIDE_AMBORTRANSPORT,
		OPEN_STATE_GUIDE_FLY_SECOND,
		OPEN_STATE_GUIDE_KAEYA_CLUE,
		OPEN_STATE_CAPTURE_CODEX,
		OPEN_STATE_ACTIVITY_FISH_OPEN,
		OPEN_STATE_ACTIVITY_FISH_CLOSE,
		OPEN_STATE_GUIDE_ROGUE_MAP,
		OPEN_STATE_GUIDE_ROGUE_RUNE,
		OPEN_STATE_GUIDE_BARTENDER_FORMULA,
		OPEN_STATE_GUIDE_BARTENDER_MIX,
		OPEN_STATE_GUIDE_BARTENDER_CUP,
		OPEN_STATE_GUIDE_MAIL_FAVORITES,
		OPEN_STATE_GUIDE_POTION_CONFIGURE,
		OPEN_STATE_GUIDE_LANV2_FIREWORK,
		OPEN_STATE_LOADINGTIPS_ENKANOMIYA,
		OPEN_STATE_MICHIAE_CASKET,
		OPEN_STATE_MAIL_COLLECT_UNLOCK_RED_POINT,
		OPEN_STATE_LUMEN_STONE,
		OPEN_STATE_GUIDE_CRYSTALLINK_BUFF,
	}
}
