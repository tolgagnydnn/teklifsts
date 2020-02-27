<template>
        <div class="registermodal">
          <form @submit.prevent="successRegister" class="registermodal__form">
            <img class="registermodal__img" src="images/registerlogo.png" />
              <p  class="registermodal__error"> {{message}} </p>
            <div class="form-group py-2">
              <div class="mb-0 input-group">
                <span class="registermodal__prepend">
                  <div class="registermodal__icon">
                    <i class="fas fa-envelope"> </i>
                  </div>
                </span>
                <input  v-model="register.email" class="registermodal__input" type="email" placeholder="Email Giriniz" required/>
              </div>
            </div>
            <div class="form-group py-2">
              <div class="mb-0 input-group">
                <span class="registermodal__prepend">
                  <div class="registermodal__icon">
                    <i class="fas fa-user"> </i>
                  </div>
                </span>
                <input v-model="register.firstName" class="registermodal__input" type="text" placeholder="Adınızı Giriniz" required/>
              </div>
            </div>
            <div class="form-group py-2">
              <div class="mb-0 input-group">
                <span class="registermodal__prepend">
                  <div class="registermodal__icon">
                    <i class="fas fa-user"> </i>
                  </div>
                </span>
                <input  v-model="register.lastName" class="registermodal__input" type="text" placeholder="Soyadınızı Giriniz" required/>
              </div>
            </div>
            <div class="form-group py-2">
              <div class="mb-0 input-group">
                <span class="registermodal__prepend">
                  <div class="registermodal__icon">
                    <i class="fas fa-phone"> </i>
                  </div>
                </span>
                <input  v-model="register.phone" class="registermodal__input" type="text" autocomplete="off" placeholder="Telefon Numaranızı Giriniz" required/>
              </div>
            </div>
            <div class="form-group py-2">
              <div class="mb-0 input-group">
                <span class="registermodal__prepend">
                  <div class="registermodal__icon">
                    <i class="fas fa-lock"> </i>
                  </div>
                </span>
                <input  v-model="register.password" class="registermodal__input" type="password" autocomplete="off" placeholder="Şifrenizi Giriniz" required/>
              </div>
            </div>

            <div>
              <button class="registermodal__btn" type="submit"> Üye Ol </button>
            </div>
            <div>
              <button class="registermodal__close" @click="closeregister"> <i class="fas fa-times"> </i> </button>
            </div>
          </form>
        </div>

</template>


<script>
import {eventBus} from '../main';
import customAxios from '../customaxios';
export default {
  data(){
    return {
      message:'',
      register: {
        email:'',
        firstName:'',
        lastName:'',
        phone:'',
        password:'',
      }
    }
  },
  methods:{
    closeregister(){
      eventBus.$emit('registerclose');
    },
    successRegister(){
      customAxios.post("user/add", { ...this.register})
      .then(res => {
        if (res.data.status) {
            console.log(res);
            eventBus.$emit('registerclose');
        } else {
            this.message = res.data.error;
            setTimeout(() => {
            this.message = ''
          }, 3000)
        }
      })
      .catch(e => console.log(e));
    }
  }
}
</script>
