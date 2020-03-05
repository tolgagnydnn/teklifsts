<template>

  <section class="login">
    <div class="container">
        <div class="row">
          <div class="col-lg-12 col-md-12 col-sm-12 col-12">
              <p class="login__heading">TeklifSTS</p>
              <p class="login__subheading">Teklif Hazırlamanın Kolay Yolu</p>
          </div>
      </div>

      <p class="login__error">{{ message }}</p>

      <div class="row">
        <div class="col-lg-12 col-md-12 col-sm-12 col-12 d-flex justify-content-center align-items-center">
          <form  @submit.prevent="userLogin" method="POST" class="login__form">
            <img src="images/userlogo.png" />
            <h1 class="login__title"> Hoşgeldiniz </h1>
            <div class="form-group py-2">
              <div class="mb-0 input-group">
                <span class="login__prepend">
                  <div class="login__icon">
                    <i class="fas fa-envelope"> </i>
                  </div>
                </span>
                <input v-model.trim="user.email" class="login__input" type="email" placeholder="Email Giriniz" required>
              </div>
            </div>
            <div class="form-group py-2">
              <div class="mb-0 input-group">
                <span class="login__prepend">
                  <div class="login__icon">
                    <i class="fas fa-lock"> </i>
                  </div>
                </span>
                <input v-model="user.password" class="login__input" type="password"  placeholder="Şifre Giriniz" required>
              </div>
            </div>
            <div class="login__forgotpassword">
              <a href="#" class="login__password"> Şifremi Unuttum </a>
            </div>

            <div class="login__btn">
               <button> Giriş Yap </button>
            </div>

            <div class="login__register">
              <a href="#" @click="openregister"> Hemen Üye Olun </a>
            </div>

          </form>
        </div>
      </div>
    </div>
    <p class="login__version" >Sürüm: {{ appVersion }}</p>

  <transition name="fade-up" appear>
    <appregister v-if="showmodal"> </appregister>
  </transition>
</section>



</template>



<script>
import packageJson from '../../package.json';
import Register from './Register'
import {eventBus} from '../main'
//import customAxios from '../customaxios'

export default{
    components:{
      "appregister":Register
    },
    data() {
      return {
        message: '',
        showmodal:false,
        user: {
          email: 'demo@demo.com',
          password: 'demo'
        }
      }
    },
    methods:{
      userLogin() {
        this.$store.dispatch('login', this.user)
        .then(res => {
          if(res == 'success'){
            this.$router.push({name:"dashboard"})
            console.log("buradayım...")
          }
          else {
            this.message = res
            setTimeout(() => {
              this.message = ''
            }, 3000);
          }
        });
      },
      openregister(){
        this.showmodal = true
      },
    },
    created(){
      eventBus.$on("registerclose", () => {
        this.showmodal = false
      });
    },
    computed: {
      appVersion() {
        return packageJson.version;
      }
    }
}
</script>
