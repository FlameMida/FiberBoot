package source

import (
	"FiberBoot/global"
	"FiberBoot/model/system"
	"github.com/gookit/color"
)

var AuthorityMenu = new(authorityMenu)

type authorityMenu struct{}

// Init @author: Flame
//@description: authority_menu 视图数据初始化
func (a *authorityMenu) Init() error {
	if global.DB.Find(&[]system.Menu{}).RowsAffected > 0 {
		color.Danger.Println("\n[Mysql] --> authority_menu 视图已存在!")
		return nil
	}
	if err := global.DB.Exec("CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `authority_menu` AS select `base_menus`.`id` AS `id`,`base_menus`.`created_at` AS `created_at`, `base_menus`.`updated_at` AS `updated_at`, `base_menus`.`deleted_at` AS `deleted_at`, `base_menus`.`menu_level` AS `menu_level`,`base_menus`.`parent_id` AS `parent_id`,`base_menus`.`path` AS `path`,`base_menus`.`name` AS `name`,`base_menus`.`hidden` AS `hidden`,`base_menus`.`component` AS `component`, `base_menus`.`title`  AS `title`,`base_menus`.`icon` AS `icon`,`base_menus`.`sort` AS `sort`,`authority_menus`.`authority_authority_id` AS `authority_id`,`authority_menus`.`base_menu_id` AS `menu_id`,`base_menus`.`keep_alive` AS `keep_alive`,`base_menus`.`close_tab` AS `close_tab`,`base_menus`.`default_menu` AS `default_menu` from (`authority_menus` join `base_menus` on ((`authority_menus`.`base_menu_id` = `base_menus`.`id`)))").Error; err != nil {
		return err
	}
	color.Info.Println("\n[Mysql] --> authority_menu 视图创建成功!")
	return nil
}
