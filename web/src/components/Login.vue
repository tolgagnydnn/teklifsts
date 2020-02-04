<template>

  <section class="login">
  <div class="container">
      <div class="row">
        <div class="col-lg-12 col-md-12 col-sm-12 col-12">
            <p class="login__heading"> TeklifSTS </p>
            <p class="login__subheading"> Teklif Hazırlama Sistemi </p>
        </div>
    </div>

    <h2>**********{{ message }}***********</h2>

    <div class="row">
      <div class="col-lg-12 col-md-12 col-sm-12 col-12 d-flex justify-content-center align-items-center">
        <form class="login__form">
          <img src="images/userlogo.png" />
          <h1 class="login__title"> Hoşgeldiniz </h1>
          <div class="form-group py-2">
            <div class="mb-0 input-group">
              <span class="login__prepend">
                <div class="login__icon">
                  <i class="fas fa-envelope"> </i>
                </div>
              </span>
              <input v-model.trim="login.email" class="login__input" type="email" placeholder="Email Giriniz" />
            </div>
          </div>
          <div class="form-group py-2">
            <div class="mb-0 input-group">
              <span class="login__prepend">
                <div class="login__icon">
                  <i class="fas fa-lock"> </i>
                </div>
              </span>
              <input v-model="login.password" class="login__input" type="password" autocomplete="off" placeholder="Şifre Giriniz" />
            </div>
          </div>
          <div class="login__forgotpassword">
            <a href="#" class="login__password"> Şifremi Unuttum </a>
          </div>

          <div class="login__btn">
            <button  @click.prevent="successLogin"> Giriş Yap </button>
          </div>

          <div class="login__register">
            <a href="#" @click="register"> Hemen Üye Olun </a>
          </div>

        </form>
      </div>
    </div>
  </div>
  <p style="text-align:center;margin-top:10px">Sürüm: {{ appVersion }}</p>
</section>
</template>



<script>
import packageJson from '../../package.json';
export default{
    data() {
      return {
        message: '',
        login: {
          email: '',
          password: ''
        }
      }
    },
    methods:{
      successLogin(){
        console.log("--> Email: ", this.login.email);
        console.log("--> Password: ", this.login.password);

        fetch(process.env.VUE_APP_API + "user/login", {
          method: 'POST',
          headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
          },
          mode: 'cors',
          cache: 'default',
          body: JSON.stringify({"email": this.login.email, "password": this.login.password})
        })
        .then((res) => { return res.json() })
        .then((data) => {
            console.log(data);
            if (data.status == false) {
                this.message = data.error;
                console.log(data.error);
                //this.$router.push({name:"login"});
            }
        })

        //this.$router.push({name:"dashboard"});
      },
      register(){
        this.$router.push({name:"register"});
      }
    },
    computed: {
      appVersion() {
        return packageJson.version;
      }
    }
}
</script>
