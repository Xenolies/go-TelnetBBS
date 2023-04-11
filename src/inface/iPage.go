package inface

// IPage 页面接口
type IPage interface {
	// Header 网站页面的头部区域，一般来讲，它包含网站的logo和一些其他元素。
	Header(IConnection)

	// Navbar 横向的导航栏，是最典型的网页元素。
	Navbar(IConnection)

	// Main 网站的主要区域，如果是博客的话它将包含的日志。
	Main(IConnection)

	// SidebarLeft 可以包含网站的次要内容，比如最近更新内容列表、关于网站的介绍或广告元素
	SidebarLeft(IConnection)

	// SidebarRight 可以包含网站的次要内容，比如最近更新内容列表、关于网站的介绍或广告元素
	SidebarRight(IConnection)

	// Footer 网站的一些附加信息，
	Footer(IConnection)
}
