<script setup lang="ts">
import ArticleCard from "@/components/ArticleCard.vue";
import {fetchArticles, addArticle} from "../../api";
import {ref} from "vue";
import AddArticleDialog from "@/components/AddArticleDialog.vue";

const isLoading = ref(false);
const articles = ref([]);

refreshArticles();

async function refreshArticles() {
  isLoading.value = true;
  const articleData = await fetchArticles();
  articles.value = articleData.articles
  isLoading.value = false;
}

async function saveArticle(article) {
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
