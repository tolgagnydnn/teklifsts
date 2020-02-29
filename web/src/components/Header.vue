<template>
  <!-- Navigation -->
  <nav class="navbar navbar-expand-lg navbar-dark bg-dark static-top mb-3">
    <div class="container">
      <router-link to="/dashboard" class="navbar-brand">TeklifSTS</router-link>
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarResponsive" aria-controls="navbarResponsive" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarResponsive">
        <ul class="navbar-nav ml-auto">
          <li class="nav-item">
            <router-link active-class="active" :to="{name: 'dashboard'}" class="nav-link" exact>Pano</router-link>
          </li>
          <li class="nav-item">
            <router-link active-class="active"  to="/offer/list" class="nav-link" exact>Teklif</router-link>
          </li>
          <li class="nav-item">
            <router-link active-class="active"  to="/customer/list" class="nav-link" exact>Müşteri</router-link>
          </li>
          <li class="nav-item">
              <router-link active-class="active"  to="/about" class="nav-link" exact>Hakkımızda</router-link>
          </li>
          <!-- Dropdown -->
          <li class="nav-item dropdown">
            <a class="nav-link dropdown-toggle" href="#" id="navbardrop" data-toggle="dropdown">
              {{`${getActiveUser.firstName} ${getActiveUser.lastName}`}}
            </a>
            <div class="dropdown-menu">
              <router-link to="/profile" class="dropdown-item">Profil</router-link>
              <router-link to="/" class="dropdown-item">Çıkış</router-link>
            </div>
          </li>

        </ul>
      </div>
    </div>
  </nav>
</template>


<script>
import {mapGetters} from 'vuex'
  export default {
    computed:{
      ...mapGetters(['getActiveUser'])
    }
  }

</script>
