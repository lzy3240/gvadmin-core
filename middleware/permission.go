package middleware

func AuthCheck() {
	//TODO
	//  从缓存中获取用户Menu,并判断是否具有权限
	//	-- 根据user_id查找所有权限id
	//	select a.user_id,b.role_id,d.menu_id from sys_user a
	//		LEFT JOIN sys_user_role b on a.user_id = b.user_id
	//		LEFT JOIN sys_role c on b.role_id = c.role_id
	//		LEFT JOIN sys_role_menu d on c.role_id = d.role_id
	//		where a.user_id = 100;
}

func DataCheck() {
	//TODO
	//	-- 根据user_id查询数据权限标识
	//	select a.user_id,b.role_id,c.data_scope,d.dept_id from sys_user a
	//		LEFT JOIN sys_user_role b on a.user_id = b.user_id
	//		LEFT JOIN sys_role c on b.role_id = c.role_id
	//		LEFT JOIN sys_role_dept d on b.role_id = d.role_id
	//		where a.user_id = 100;
}
