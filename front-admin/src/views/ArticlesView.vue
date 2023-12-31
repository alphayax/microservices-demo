<script setup lang="ts">
import {ref} from "vue";
import {notify} from "@kyvg/vue3-notification";
import {fetchArticles, addArticle} from "../api";
import ArticleCard from "@/components/ArticleCard.vue";
import AddArticleDialog from "@/components/AddArticleDialog.vue";
import type {Article} from "@/model";

const isLoading = ref(false);
const articles = ref<Article[]>([]);

refreshArticles();

async function refreshArticles() {
  try {
    isLoading.value = true;
    articles.value = await fetchArticles();
    notify({
      title: 'Articles fetched',
      text: `Fetched ${articles.value.length} articles`,
      type: 'success'
    });
  } catch (e: any) {
    notify({
      title: 'Unable to fetch articles',
      text: e,
      type: 'error'
    });
  } finally {
    isLoading.value = false;
  }
}

async function saveArticle(article: Article) {
  await addArticle(article)
  await refreshArticles()
}
</script>

<template>
  <v-container>

    <v-toolbar density="compact" title="Articles">
      <v-spacer/>
      <AddArticleDialog @onArticleAdded="saveArticle"/>
      <v-btn icon="mdi-refresh" @click="refreshArticles" :loading="isLoading"/>
    </v-toolbar>

    <v-container>
      <v-row>
        <v-col cols="4" v-for="article in articles" :key="article.id">
          <ArticleCard :article="article"/>
        </v-col>
      </v-row>
    </v-container>
  </v-container>
</template>

<style scoped>

</style>
