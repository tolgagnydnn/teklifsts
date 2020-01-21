<template>

  <section class="login">
  <div class="container">
    <div class="row">
      <div class="col-lg-12 col-md-12 col-sm-12 col-12 d-flex justify-content-center align-items-center">
        <form>
          <img src="images/userlogo.png" />
          <h1> Hoşgeldiniz </h1>
          <div class="form-group py-2">
            <div class="mb-0 input-group">
              <span class="input-group-prepend">
                <div class="input-group-text">
                  <i class="fas fa-envelope"> </i>
                </div>
              </span>
              <input v-model.trim="kullanicigiris.eposta" class="form-control" type="email" placeholder="Email Giriniz" />
            </div>
          </div>
          <div class="form-group py-2">
            <div class="mb-0 input-group">
              <span class="input-group-prepend">
                <div class="input-group-text">
                  <i class="fas fa-lock"> </i>
                </div>
              </span>
              <input v-model="kullanicigiris.sifre" class="form-control" type="password" autocomplete="off" placeholder="Şifre Giriniz" />
            </div>
          </div>
          <div class="float-left">
            <a href="#" class="sifre"> Şifremi Unuttum </a>
          </div>

          <div class="loginBtn float-right">
            <button type="submit"> Giriş Yap </button>
          </div>
        </form>
      </div>
    </div>
  </div>
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
    }
  }
</script>
