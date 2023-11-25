<script setup lang="ts">
import {Article} from "@/model";
import {ref} from "vue";
import {fetchArticles, saveCart} from "@/api";
import {useAppStore} from "@/store/app";
import ArticleCard from "@/components/ArticleCard.vue";

const appStore = useAppStore()
const articles = ref<Article[]>([])

refresh()

async function refresh() {
  await appStore.loadCart()
  articles.value = await fetchArticles()
}

async function addToCart(articleId: string) {
  appStore.addToCart(articleId)
  await saveCart(appStore.cartId, appStore.cartItems)
}

async function removeFromCart(articleId: string) {
  console.log('removeFromCart', articleId)
  appStore.removeFromCart(articleId)
  await saveCart(appStore.cartId, appStore.cartItems)
}
</script>

<template>
  <v-container>
    <v-row>
      <v-col v-for="article in articles" cols="4">
        <ArticleCard
          :article="article"
          @onArticleAdded="addToCart"
          @onArticleRemoved="removeFromCart"
        />
      </v-col>
    </v-row>
  </v-container>
</template>

<style scoped>

</style>
