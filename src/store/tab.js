import Cookie from 'js-cookie'
export default {
    state: {
        isCollapse: false,//控制菜单栏的收起和展开
        tabsList: [
            {
                path: "/home",
                name: "home",
                label: "首页",
                icon: "s-home",
                url: "Home/Home",
            }
        ],//面包屑数据
        menu: []
    },

    mutations: {
        //修改菜单收起展开的方法
        collapseMenu(state) {
            state.isCollapse = !state.isCollapse
        },
        //更新面包屑数据
        selectMenu(state, val) {
            if (val.name !== "home") {
                const index = state.tabsList.findIndex(item => item.name === val.name)
                if (index === -1) {
                    state.tabsList.push(val)
                }
            }
        },
        //删除面包屑数据
        closeTag(state, item) {
            const index = state.tabsList.findIndex(val => val.name === item.name)
            state.tabsList.splice(index, 1)
        },
        //设置menu的数据
        setMenu(state, val) {
            state.menu = val
            Cookie.set("menu", JSON.stringify(val))
        },
        //动态注册路由
        addMenu(state, router) {
            //判断缓存是否有数据
            if (!Cookie.get('menu')) return
            const menu = JSON.parse(Cookie.get('menu'))
            state.menu = menu
            //组装路由数据
            const menuArr = []
            menu.forEach(item => {
                if (item.children) {
                    item.children.map(item => {
                        item.component = () => import(`../views/${item.url}`)
                        return item
                    })
                    menuArr.push(...item.children)
                } else {
                    item.component = () => import(`../views/${item.url}`)
                    menuArr.push(item)

                }
            });
            //路由的动态添加
            menuArr.forEach(item => {
                router.addRoute("Main", item)
            })
        }
    }
}