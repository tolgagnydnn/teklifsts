import Vue from 'vue'
import Router from 'vue-router'

import Pano from './components/Pano.vue'
import TeklifListe from './components/Teklif/Liste.vue'
import MusteriListe from './components/Musteri/Liste.vue'

Vue.use(Router)

const routes = [
    { path: '/', component: Pano },
    { path: '/teklif/liste', component: TeklifListe },
    { path: '/musteri/liste', component: MusteriListe }
]

export const router = new Router({
	mode: "history",
	routes: routes
});