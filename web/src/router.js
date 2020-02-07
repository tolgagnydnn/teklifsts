import Vue from 'vue'
import Router from 'vue-router'

import Dashboard from './components/Dashboard.vue'
import OfferList from './components/offer/List.vue'
import CustomerList from './components/customer/List.vue'
import Profile from './components/Profile.vue'
import Login from './components/Login.vue'
import Header from './components/Header.vue'
import Register from './components/Register.vue'
import About from './components/About.vue'

Vue.use(Router)

const routes = [
    {
      path:'/',
      name:'login',
      component:Login
    },
    {
      path:'/register',
      name:'register',
      component:Register
    },
    {
        path: '/dashboard',
        name: 'dashboard',
        components:{
          default:Dashboard,
          "headertop":Header
        }
    },
    {
        path: '/offer/list',
        name: "offer",
        components:{
          default:OfferList,
          "headertop":Header
        }
    },
    {
        path: '/customer/list',
        name: 'customer',
        components:{
          default:CustomerList,
          "headertop":Header
        }
    },
    {
      path:'/profile',
      name:'profile',
      component:Profile,
      components:{
        default:Profile,
        "headertop":Header
      }
    },
    {
        path: '/about',
        name: 'about',
        component: About,
        components: {
            default: About,
            "headertop": Header
        }
    }
]


export const router = new Router({
	mode: "history",
	routes: routes
});

router.beforeResolve((to, from, next) => {
    if (to.name) {
        //NProgress.start()
    }
    next()
})
