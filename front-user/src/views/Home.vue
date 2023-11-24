<template>
  <v-container class="fill-height">
    <v-responsive class="align-center text-center fill-height">
      <v-row class="d-flex align-center justify-center">
        <div class="text-body-2 font-weight-light mb-n1">Welcome to</div>
      </v-row>
      <v-row class="d-flex align-center justify-center">
        <h1 class="text-h1">The awesome alph<span class="text-blue">&alpha;</span> shop !</h1>
      </v-row>
      <v-row class="d-flex align-center justify-center">
        <v-btn-group>
          <v-btn text="Go to the shop" color="primary" variant="outlined" prepend-icon="mdi-store" to="/shop" />
          <v-btn text="Go to the cart" color="" variant="outlined" prepend-icon="mdi-cart" to="/cart" />
        </v-btn-group>
      </v-row>
    </v-responsive>
  </v-container>
</template>

<script lang="ts" setup>

</script>
