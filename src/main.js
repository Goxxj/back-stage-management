import Vue from 'vue'
import App from './App.vue'
import {
  Container, Header, RadioGroup, RadioButton, DropdownItem, Row, Col, Card, Table, TableColumn, Tag, Dialog,
  Menu, Submenu, MenuItem, Aside, MenuItemGroup, Main, Button, DropdownMenu, Dropdown, Breadcrumb, BreadcrumbItem,
  Form, FormItem, Input, Select, Option, DatePicker, Message, MessageBox, Pagination
} from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import router from './router'
import store from './store'
import Cookie from "js-cookie";


Vue.config.productionTip = false


Vue.prototype.$confirm = MessageBox.confirm;
Vue.prototype.$message = Message;
Vue.prototype.$msgbox = MessageBox;
Vue.use(Header);
Vue.use(Pagination);
Vue.use(DatePicker);
Vue.use(Option);
Vue.use(Select);
Vue.use(Form);
Vue.use(FormItem);
Vue.use(Input);
Vue.use(Dialog);
Vue.use(Tag);
Vue.use(Breadcrumb);
Vue.use(BreadcrumbItem);
Vue.use(TableColumn);
Vue.use(Table);
Vue.use(Row);
Vue.use(Col);
Vue.use(Card);
Vue.use(Container);
Vue.use(RadioGroup);
Vue.use(RadioButton);
Vue.use(Menu);
Vue.use(DropdownItem);
Vue.use(Submenu);
Vue.use(MenuItem);
Vue.use(Aside);
Vue.use(MenuItemGroup);
Vue.use(Main);
Vue.use(Button);
Vue.use(DropdownMenu);
Vue.use(Dropdown);

router.beforeEach((to, from, next) => {
  const token = Cookie.get('token')
  if (!token) {
    if (to.path === '/login' || to.path === '/' || to.path === '/register') {
      next()
    } else {
      next({ path: '/login' })
    }
  } else {
    next()
  }
})


new Vue({
  router,
  store,
  render: h => h(App),
  created() {
    store.commit("addMenu", router)
  }
}).$mount('#app')
