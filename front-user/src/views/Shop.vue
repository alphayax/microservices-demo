<script setup lang="ts">

import {Article} from "@/model";
import {ref} from "vue";
import {fetchArticles, saveCart} from "@/api";
import {useAppStore} from "@/store/app";

const appStore = useAppStore()
const articles = ref<Article[]>([])

refresh()

async function refresh() {
  articles.value = await fetchArticles()
}

async function addToCart(articleId: string) {
  console.log(`Add ${articleId} to cart`)
  appStore.addToCart(articleId)
  await saveCart(appStore.cartId, appStore.cartItems)
}

</script>

<template>
  <v-container>
    <v-row>
      <v-col v-for="article in articles" cols="4">
        <v-card :title="article.title" :subtitle="article.id">
          <v-card-text>{{ article.description }}</v-card-text>
          <v-card-actions>
            <v-spacer/>
            <v-btn prepend-icon="mdi-plus" text="Add to cart" @click="addToCart(article.id)"/>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<style scoped>

</style>
