<template>

  <section class="müsteriliste">
    <div class="container">
      <div class="row">
          <div class="col-md-12">
               <h1> Müşteri Listesi </h1>
              <div class="mb-2"><a href="#" class="btn btn-success">Müşteri Ekle</a></div>
              <table class="table table-hover">
                  <thead>
                  <tr>
                      <th>Firstname</th>
                      <th>Lastname</th>
                      <th>Email</th>
                  </tr>
                  </thead>
                  <tbody>
                  <tr>
                      <td>John</td>
                      <td>Doe</td>
                      <td>john@example.com</td>
                  </tr>
                  <tr>
                      <td>Mary</td>
                      <td>Moe</td>
                      <td>mary@example.com</td>
                  </tr>
                  <tr>
                      <td>July</td>
                      <td>Dooley</td>
                      <td>july@example.com</td>
                  </tr>
                  </tbody>
              </table>
          </div>
      </div>
    </div>
  </section>
</template>

<script>
export default {
    name: 'Liste'
}
</script>
