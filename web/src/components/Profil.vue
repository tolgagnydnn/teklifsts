<template>

<section class="profile">
  <div class="container">
    <div class="title">
      <p> Kişisel Bilgilerim </p>
    </div>
    <div class="row py-5">
      <div class="col-lg-12 col-md-12 col-sm-12 col-12">
        <form>
          <div class="form-group row">
            <label class="col-lg-2 col-md-2 col-sm-2 col-6 col-form-label">E Posta Adresiniz :</label>
            <div class="col-lg-3 col-md-3 col-sm-3 col-6 mt-1">
              <span> tolgagunaydin@gmail.com </span>
            </div>
          </div>
          <div class="form-group row">
            <label class="col-lg-2 col-md-2 col-sm-2 col-6 col-form-label">Adınız Soyadınız :</label>
            <div class="col-lg-3 col-md-3 col-sm-3 col-6">
              <input v-model.trim="kullanici.adi" type="text" class="form-control" placeholder="Tolga Günaydın">
            </div>
          </div>
          <div class="form-group row">
            <label class="col-lg-2 col-md-2 col-sm-2 col-6 col-form-label">Telefon Numaranız :</label>
            <div class="col-lg-3 col-md-3 col-sm-3 col-6">
              <input type="text" v-model.trim.number="kullanici.telefon" class="form-control" placeholder="0999 999 99 99">
            </div>
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
        kullanici:{
          adi:"",
          telefon:""
        }
      }
    }
  }
</script>
