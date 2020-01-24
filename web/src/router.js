import Vue from 'vue'
import Router from 'vue-router'

import Pano from './components/Pano.vue'
import TeklifListe from './components/Teklif/Liste.vue'
import MusteriListe from './components/Musteri/Liste.vue'
import Profil from './components/Profil.vue'
import Login from './components/Login.vue'
import Header from './components/Header.vue'

Vue.use(Router)

const routes = [
    {
      path:'/',
      name:'login',
      components:{
        default:Login
      }
    },
    {
        path: '/pano',
        name: 'pano',
        components:{
          default:Pano,
          "headertop":Header
        }
    },
    {
        path: '/teklif/liste',
        name: "teklif",
        components:{
          default:TeklifListe,
          "headertop":Header
        }
    },
    {
        path: '/musteri/liste',
        name: 'musteri',
        components:{
          default:MusteriListe,
          "headertop":Header
        }
    },
    {
      path:'/profil',
      name:'profil',
      component:Profil,
      components:{
        default:Profil,
        "headertop":Header
      }
    },
    {
      path:'/login',
      name:'login',
      components:{
        default:Login,
        "headertop":Header
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
