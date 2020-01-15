<template>
  <div id="app">
    <!-- Navigation -->
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark static-top mb-3">
      <div class="container">
        <a class="navbar-brand" href="#">Start Bootstrap</a>
        <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarResponsive" aria-controls="navbarResponsive" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarResponsive">
          <ul class="navbar-nav ml-auto">
            <li class="nav-item">
              <router-link :to="{name: 'pano'}" class="nav-link" exact>Pano</router-link>
            </li>
            <li class="nav-item">
              <router-link to="/teklif/liste" class="nav-link" exact>Teklif</router-link>
            </li>
            <li class="nav-item">
              <router-link to="/musteri/liste" class="nav-link" exact>Müşteri</router-link>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="#">İletişim</a>
            </li>
          </ul>
        </div>
      </div>
    </nav>

    <!-- Page Content -->
    <div class="container">
        <router-view></router-view>
    </div>
    
  </div>
</template>

<script>
export default {
  name: 'app'
}
</script>