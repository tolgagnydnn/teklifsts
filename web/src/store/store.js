import Vue from 'vue'
import Vuex from 'vuex'
import customAxios from '../customaxios';


Vue.use(Vuex);

const store = new Vuex.Store({
  state:{
    activeUser: null
  },
  getters: {
    getActiveUser(state) {
      return state.activeUser
    }
  },
  mutations:{
    setActiveUser(state, user) {
      state.activeUser = user
    }
  },
  actions:{
    login(context, user) {
      return customAxios.post('user/login?email='+ user.email +'&password='+ user.password)
        .then(res => {

          if(res.data.status) {
              console.log(res.data)
            context.commit('setActiveUser', res.data.data);
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
