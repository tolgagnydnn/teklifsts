<template>
<section class="profile">
  <div class="container">
    <div class="row">
      <div class="col-lg-12 col-md-12 col-sm-12 col-12 mt-5">
        <div>
          <span class="profile__title">Kişisel Bilgilerim</span>
          <span class="profile__edit">
            <button @click="duzenle()">Düzenle</button>
          </span>
          <hr>
        </div>
        <div class="profile__info">
          <span class="profile__itemtitle">E-Postanız:</span>
          <span class="profile__itemcontent">
            <input type="text" v-model="profil.eposta">
          </span>
        </div>
        <div class="profile__info">
          <span class="profile__itemtitle">Adınız:</span>
          <span class="profile__itemcontent">
            <input type="text" v-model="profil.adi">
          </span>
        </div>
        <div class="profile__info">
          <span class="profile__itemtitle">Soyadınız:</span>
          <span class="profile__itemcontent">
            <input type="text" v-model="profil.soyadi">
          </span>
        </div>
        <div class="profile__info">
          <span class="profile__itemtitle">Telefon Numaranız:</span>
          <span class="profile__itemcontent">
            <input type="text" v-model="profil.telefon">
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
export default {
  name: 'Profil',
  data() {
    return {
      profil: {
        id: null,
        eposta: null,
        adi: null,
        soyadi: null,
        telefon: null
      }
    }
  },
  methods: {
    duzenle: function() {
      console.log(this.profil);
      fetch("http://127.0.0.1:8090/v1/kullanici/ekle", {
        method: 'POST',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        mode: 'cors',
        cache: 'default',
        body: JSON.stringify({"id": this.profil.id})
      })
      .then((res) => { return res.json() })
      .then((data) => { console.log(data) })
    }
  },
  created: function() {
    fetch("http://127.0.0.1:8090/v1/kullanici/bilgi/1")
    .then((res) => { return res.json(); })
    .then((res) => {
      if (res.status) {
        this.profil = res.data;
      } else {
        alert(res.data.error);
      }
    })
  }
}
</script>
