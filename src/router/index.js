import Vue from 'vue'
import VueRouter from 'vue-router'
import Main from '../views/Main'
// import Home from '../views/Home'
// import User from '../views/User'
// import Mall from '../views/Mall'
// import PageOne from '../views/PageOne'
// import PageTwo from '../views/PageTwo'
import Login from '../views/Login'
import Register from '../views/Register'

Vue.use(VueRouter)

const routes = [
    {
        path: '/',
        component: Main,
        name: "Main",
        redirect: '/home',
        children: [
            // { path: 'home', name: "home", component: Home },//主页
            // { path: 'user', name: "user", component: User },//用户管理
            // { path: 'mall', name: "mall", component: Mall },//商品管理
            // { path: 'page1', name: "page1", component: PageOne },//分页1
            // { path: 'page2', name: "page2", component: PageTwo },//分页2
        ]
    },
    {
        path: "/login",
        name: "login",
        component: Login,
    },
    {
        path: "/register",
        name: "register",
        component: Register,
    },

]

const router = new VueRouter({
    routes // (缩写) 相当于 routes: routes
})

export default router