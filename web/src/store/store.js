import Vue from 'vue'
import Vuex from 'vuex'
import customAxios from '../customaxios';


Vue.use(Vuex);

const store = new Vuex.Store({
  state:{
    activeUser:[]
  },
  getters: {
    getActiveUser(state) {
      return state.activeUser
    }
  },
  mutations:{
    setActiveUser(state, authData ){
      state.activeUser = authData
    }
  },
  actions:{
    login(vuexcontext,authData) {
      return customAxios.post('user/login?email='+ authData.email +'&password='+ authData.password)
        .then(res => {
          if(res.data.status) {
            vuexcontext.commit('setActiveUser', res.data.data);
            return 'success'
          }
          else {
             return res.data.error
           }
        })
    }

  }

});


export default store
