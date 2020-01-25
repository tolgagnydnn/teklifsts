<template>
  <!-- Navigation -->
  <nav class="navbar navbar-expand-lg navbar-dark bg-dark static-top mb-3">
    <div class="container">
      <router-link to="/pano" class="navbar-brand">TeklifSTS</router-link>
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarResponsive" aria-controls="navbarResponsive" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarResponsive">
        <ul class="navbar-nav ml-auto">
          <li class="nav-item">
            <router-link active-class="active"  :to="{name: 'pano'}" class="nav-link" exact>Pano</router-link>
          </li>
          <li class="nav-item">
            <router-link active-class="active"  to="/teklif/liste" class="nav-link" exact>Teklif</router-link>
          </li>
          <li class="nav-item">
            <router-link active-class="active"  to="/musteri/liste" class="nav-link" exact>Müşteri</router-link>
          </li>
          <li class="nav-item">
              <router-link active-class="active"  to="/hakkimizda" class="nav-link" exact>Hakkımızda</router-link>
          </li>
          <!-- Dropdown -->
          <li class="nav-item dropdown">
            <a class="nav-link dropdown-toggle" href="#" id="navbardrop" data-toggle="dropdown">
              Hesabım
            </a>
            <div class="dropdown-menu">
              <router-link to="/profil" class="dropdown-item">Profil</router-link>
              <router-link to="/" class="dropdown-item">Çıkış</router-link>
            </div>
          </li>

        </ul>
      </div>
    </div>
  </nav>
</template>


<script>


</script>
