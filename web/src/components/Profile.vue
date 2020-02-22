<template>
<section class="profile">
  <div class="container">
    <div class="row">
      <div class="col-lg-12 col-md-12 col-sm-12 col-12 mt-5">
        <div>
          <span class="profile__title">Kişisel Bilgilerim</span>
          <span class="profile__edit"> Düzenle </span>

          <hr>
        </div>
        <div class="profile__info">
          <span class="profile__itemtitle">E-Postanız:</span>
          <span class="profile__itemcontent">
            <input type="text" v-model="profile.email">
          </span>
        </div>
        <div class="profile__info">
          <span class="profile__itemtitle">Adınız:</span>
          <span class="profile__itemcontent">
            <input type="text" v-model="profile.firstName">
          </span>
        </div>
        <div class="profile__info">
          <span class="profile__itemtitle">Soyadınız:</span>
          <span class="profile__itemcontent">
            <input type="text" v-model="profile.lastName">
          </span>
        </div>
        <div class="profile__info">
          <span class="profile__itemtitle">Telefon Numaranız:</span>
          <span class="profile__itemcontent">
            <input type="text" v-model="profile.phone">
          </span>
        </div>
      </div>

      <div class="col-lg-12 col-md-12 col-sm-12 col-12 mt-5">
        <div>
          <span class="profile__title"> Şifre Bilgilerim </span>
          <span class="profile__edit">Düzenle</span>
          <hr>
        </div>
        <div class="profile__password">
          <span class="profile__itemtitle"> Mevcut Şifreniz: </span>
          <span class="profile__itemcontent"> *******  </span>
        </div>
     </div>

   </div>
  </div>
</section>

</template>


<script>
import customAxios from '../customaxios'

export default {
  name: 'Profil',
  data() {
    return {
      profile: {
        id: null,
        email: null,
        firstName: null,
        lastName: null,
        phone: null
      }
    }
  },
  methods: {
    duzenle: function() {
      console.log(this.profile);
      fetch(process.env.VUE_APP_API + "user/ekle", {
        method: 'POST',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        mode: 'cors',
        cache: 'default',
        body: JSON.stringify({"id": this.profile.id})
      })
      .then((res) => { return res.json() })
      .then((data) => { console.log(data) })
    }
  },
  created: function() {
    console.log("calisti.....")
    customAxios.get("user/3")
    .then(res => {
      console.log(res)
      this.profile = res.data.data
    })
    /*
    fetch(process.env.VUE_APP_API + "user/bilgi/1")
    .then((res) => { return res.json(); })
    .then((res) => {
      if (res.status) {
        this.profile = res.data;
      } else {
        alert(res.data.error);
      }
    })
    */
  }
}
</script>
