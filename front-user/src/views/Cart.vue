<script setup lang="ts">
import {ref} from "vue";
import {saveCart} from "@/api";
import {useAppStore} from "@/store/app";

const appStore = useAppStore()
const cartItems = ref<string[]>([])

refresh()

async function refresh() {
  await appStore.loadCart()
  cartItems.value = appStore.cartItems
}

async function removeFromCart(articleId: string) {
  appStore.removeFromCart(articleId)
  await saveCart(appStore.cartId, appStore.cartItems)
  await refresh()
}
</script>

<template>
  <v-container>
    Items in cart:
    <v-list>
      <v-list-item v-for="item in cartItems" :key="item" :title="item">
        <template v-slot:append>
          <v-btn @click="removeFromCart(item)" icon="mdi-delete" variant="tonal" color="red" size="small"/>
        </template>
      </v-list-item>
    </v-list>
  </v-container>
</template>

<style scoped>

</style>
