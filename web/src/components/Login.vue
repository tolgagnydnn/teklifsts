<template>

  <section class="login">
  <div class="container">
      <div class="row">
        <div class="col-lg-12 col-md-12 col-sm-12 col-12 d-flex justify-content-center align-items-center">
            <h2>TeklifSTS<br><h4>Kolay Teklif Hazırlama Sistemi</h4></h2>
        </div>
    </div>

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
              <input v-model.trim="kullanicigiris.eposta" class="login__input" type="email" placeholder="Email Giriniz" />
            </div>
          </div>
          <div class="form-group py-2">
            <div class="mb-0 input-group">
              <span class="login__prepend">
                <div class="login__icon">
                  <i class="fas fa-lock"> </i>
                </div>
              </span>
              <input v-model="kullanicigiris.sifre" class="login__input" type="password" autocomplete="off" placeholder="Şifre Giriniz" />
            </div>
          </div>
          <div class="login__forgotpassword">
            <a href="#" class="login__password"> Şifremi Unuttum </a>
          </div>

          <div class="login__btn">
            <button type="submit" @click="successLogin"> Giriş Yap </button>
          </div>

          <div class="login__register">
            <a href="#"> Hemen Üye Olun </a>
          </div>

        </form>
      </div>
    </div>
  </div>
  <p style="text-align:center;margin-top:10px">Sürüm: 0.3.1</p>
</section>
</template>



<script>
  export default{
    data(){
      return {
        kullanicigiris:{
          eposta:"",
          sifre:""
        }
      }
    },
    methods:{
      successLogin(){
        this.$router.push({name:"pano"});
      }
    }
  }
</script>
