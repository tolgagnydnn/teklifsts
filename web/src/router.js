import Vue from 'vue'
import Router from 'vue-router'

import Pano from './components/Pano.vue'
import TeklifListe from './components/Teklif/Liste.vue'
import MusteriListe from './components/Musteri/Liste.vue'
import Profil from './components/Profil.vue'
import Login from './components/Login.vue'

Vue.use(Router)

const routes = [
    {
      path:'/',
      name:'login',
      component:Login
    },
    {
        path: '/pano',
        name: 'pano',
        component: Pano,
    },
    {
        path: '/teklif/liste',
        name: "teklif",
        component: TeklifListe,
    },
    {
        path: '/musteri/liste',
        name: 'musteri',
        component: MusteriListe,
    },
    {
      path:'/profil',
      name:'profil',
      component:Profil
    },
    {
      path:'/login',
      name:'login',
      component:Login
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
