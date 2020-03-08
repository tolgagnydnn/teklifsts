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
          <span class="profile__itemtitle">E-Postanız</span>
          <span class="profile__itemcontent">
            <p> {{getActiveUser.email}} </p>
          </span>
        </div>
        <div class="profile__info">
          <span class="profile__itemtitle">Adınız</span>
          <span class="profile__itemcontent">
            <input type="text" class="profile__input"  v-model="getActiveUser.firstName">
          </span>
        </div>
        <div class="profile__info">
          <span class="profile__itemtitle">Soyadınız</span>
          <span class="profile__itemcontent">
            <input type="text" class="profile__input" v-model="getActiveUser.lastName">
          </span>
        </div>
        <div class="profile__info">
          <span class="profile__itemtitle">Telefon Numaranız</span>
          <span class="profile__itemcontent">
            <input type="text" class="profile__input"  v-model="getActiveUser.phone">
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
          <span class="profile__itemtitle"> Mevcut Şifreniz </span>
          <span class="profile__itemcontent"> ****** </span>
        </div>
     </div>

   </div>
  </div>
</section>

</template>


<script>

import {mapGetters} from 'vuex'

export default {
  name: 'Profil',
  data() {
    return {
      disabled:true,
      profile: {
        id: null,
        email: null,
        firstName: null,
        lastName: null,
        phone: null
      }
    }
  },
  computed:{
      ...mapGetters(['getActiveUser'])
  }
}
</script>
