<template>
        <div class="registermodal">
          <form class="registermodal__form">
            <img class="registermodal__img" src="images/registerlogo.png" />
            <div class="form-group py-2">
              <div class="mb-0 input-group">
                <span class="registermodal__prepend">
                  <div class="registermodal__icon">
                    <i class="fas fa-envelope"> </i>
                  </div>
                </span>
                <input  class="registermodal__input" type="email" placeholder="Email Giriniz" />
              </div>
            </div>
            <div class="form-group py-2">
              <div class="mb-0 input-group">
                <span class="registermodal__prepend">
                  <div class="registermodal__icon">
                    <i class="fas fa-user"> </i>
                  </div>
                </span>
                <input  class="registermodal__input" type="email" placeholder="Adınızı Giriniz" />
              </div>
            </div>
            <div class="form-group py-2">
              <div class="mb-0 input-group">
                <span class="registermodal__prepend">
                  <div class="registermodal__icon">
                    <i class="fas fa-user"> </i>
                  </div>
                </span>
                <input  class="registermodal__input" type="email" placeholder="Soyadınızı Giriniz" />
              </div>
            </div>
            <div class="form-group py-2">
              <div class="mb-0 input-group">
                <span class="registermodal__prepend">
                  <div class="registermodal__icon">
                    <i class="fas fa-lock"> </i>
                  </div>
                </span>
                <input  class="registermodal__input" type="password" autocomplete="off" placeholder="Şifrenizi Giriniz" />
              </div>
            </div>
            <div class="form-group py-2">
              <div class="mb-0 input-group">
                <span class="registermodal__prepend">
                  <div class="registermodal__icon">
                    <i class="fas fa-lock"> </i>
                  </div>
                </span>
                <input  class="registermodal__input" type="password" autocomplete="off" placeholder="Şifrenizi Tekrar Giriniz" />
              </div>
            </div>

            <div>
              <button class="registermodal__btn" type="submit"> Üye Ol </button>
            </div>
            <div>
              <button class="registermodal__close" @click="closeregister"> <i class="fas fa-times"> </i> </button>
            </div>
          </form>
        </div>

</template>


<script>
import {eventBus} from '../main'
export default {
  methods:{
    closeregister(){
      eventBus.$emit('registerclose');
    }
  }
}
</script>
