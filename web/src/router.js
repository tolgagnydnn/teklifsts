import Vue from 'vue'
import Router from 'vue-router'
import store from './store/store'

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
        },
        beforeEnter(to, from, next) {
          if (store.getters.getActiveUser.userID !== undefined) {
            next()
          } else {
            next('/')
          }
        }
    },
    {
        path: '/offer/list',
        name: "offer",
        components:{
          default:OfferList,
          "headertop":Header
        },
        beforeEnter(to, from, next) {
          if (store.getters.getActiveUser.userID !== undefined) {
            next()
          } else {
            next('/')
          }
        }
    },
    {
        path: '/customer/list',
        name: 'customer',
        components:{
          default:CustomerList,
          "headertop":Header
        },
        beforeEnter(to, from, next) {
          if (store.getters.getActiveUser.userID !== undefined) {
            next()
          } else {
            next('/')
          }
        }
    },
    {
      path:'/profile',
      name:'profile',
      component:Profile,
      components:{
        default:Profile,
        "headertop":Header
      },
      beforeEnter(to, from, next) {
        if (store.getters.getActiveUser.userID !== undefined) {
          next()
        } else {
          next('/')
        }
      }
    },
    {
        path: '/about',
        name: 'about',
        component: About,
        components: {
            default: About,
            "headertop": Header
        },
        beforeEnter(to, from, next) {
          if (store.getters.getActiveUser.userID !== undefined) {
            next()
          } else {
            next('/')
          }
        }
    },

]


export const router = new Router({
	mode: "history",
	routes: routes
});
