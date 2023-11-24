<template>
  <v-app>

    <v-navigation-drawer>
      <v-list density="comfortable">
        <v-list-item to="/" prepend-icon="mdi-home" title="Home"/>
        <v-list-item to="/shop" prepend-icon="mdi-store" title="Shop"/>
        <v-list-item to="/cart" prepend-icon="mdi-cart" title="Cart"/>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar title="The awesome alph&alpha; shop"/>

    <v-main>
      <RouterView/>
    </v-main>

    <v-footer :app="true">
      <span class="text-grey-darken-1 font-italic">WeScale Training</span>
    </v-footer>

  </v-app>
</template>

<script lang="ts" setup>
  //
</script>
