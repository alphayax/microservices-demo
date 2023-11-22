<script setup lang="ts">
import {ref} from "vue";
import {fetchArticles} from "../../api";

const articleCount = ref(0)

refreshArticleCount()

async function refreshArticleCount() {
  const articles = await fetchArticles()
  articleCount.value = articles.length
}
</script>

<template>
  <v-container>
    <v-toolbar density="compact" title="Home"/>

    <v-container>
      <v-row>
        <v-col cols="6">

          <v-card title="Articles">
            <v-list>
              <v-list-item>
                <v-list-item-title>Articles in database</v-list-item-title>
                <template v-slot:append>
                  <span class="text-amber">{{ articleCount }}</span>
                </template>
              </v-list-item>
            </v-list>
          </v-card>

        </v-col>
      </v-row>
    </v-container>
  </v-container>
</template>
