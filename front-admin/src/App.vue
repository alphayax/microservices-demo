<script setup lang="ts">
import {RouterView} from 'vue-router'
</script>

<template>
  <v-app>

    <v-navigation-drawer>
      <v-list density="comfortable">
        <v-list-item to="/" prepend-icon="mdi-home" title="Home"/>
        <v-list-item to="/articles" prepend-icon="mdi-home" title="Articles"/>
        <v-list-item to="/about" prepend-icon="mdi-home" title="About"/>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar title="Article Admin"/>

    <v-main>
      <RouterView/>
    </v-main>

    <v-footer :app="true">
      <span class="text-grey-darken-1 font-italic">WeScale Training</span>
    </v-footer>

    <notifications />
  </v-app>
</template>

<style scoped>

</style>
