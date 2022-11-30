-- 基础信息
local base_info = {
	group_id = 199001212
}

-- DEFS_MISCS
local shadowConfigIDList = {212003,212004,212005}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
	{ config_id = 212001, gadget_id = 70500033, pos = { x = 129.162, y = 121.608, z = -170.012 }, rot = { x = 0.000, y = 11.555, z = 0.000 }, level = 1, arguments = { 40 }, area_id = 401 },
	{ config_id = 212003, gadget_id = 70500047, pos = { x = 120.184, y = 120.203, z = -162.498 }, rot = { x = 0.000, y = 341.695, z = 0.000 }, level = 1, area_id = 401 },
	{ config_id = 212004, gadget_id = 70500051, pos = { x = 120.644, y = 120.382, z = -172.346 }, rot = { x = 0.000, y = 339.893, z = 0.000 }, level = 1, area_id = 401 },
	{ config_id = 212005, gadget_id = 70500041, pos = { x = 118.801, y = 120.444, z = -172.910 }, rot = { x = 0.000, y = 341.373, z = 0.000 }, level = 1, area_id = 401 }
}

-- 区域
regions = {
	{ config_id = 212002, shape = RegionShape.SPHERE, radius = 10, pos = { x = 125.796, y = 120.270, z = -167.710 }, area_id = 401 }
}

-- 触发器
triggers = {
}

-- 变量
variables = {
}

--================================================================
-- 
-- 初始化配置
-- 
--================================================================

-- 初始化时创建
init_config = {
	suite = 1,
	end_suite = 0,
	rand_suite = false
}

--================================================================
-- 
-- 小组配置
-- 
--================================================================

suites = {
	{
		-- suite_id = 1,
		-- description = ,
		monsters = { },
		gadgets = { 212001 },
		regions = { 212002 },
		triggers = { },
		rand_weight = 100
	}
}

--================================================================
-- 
-- 触发器
-- 
--================================================================

require "V2_8/EchoConch"