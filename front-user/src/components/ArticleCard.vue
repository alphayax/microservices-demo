<script setup lang="ts">
import type {Article} from "@/model";

interface Props {
  article: Article
}

const props = defineProps<Props>()
const emit = defineEmits(['onArticleAddedToCart'])

function addToCart() {
  emit('onArticleAddedToCart', props.article.id)
}
</script>

<template>
  <v-card :title="article.title" :subtitle="article.id">
    <v-card-text>{{ article.description }}</v-card-text>
    <v-card-actions>
      <v-spacer/>
      <v-btn prepend-icon="mdi-plus" text="Add to cart" @click="addToCart"/>
    </v-card-actions>
  </v-card>
</template>

<style scoped>

</style>
