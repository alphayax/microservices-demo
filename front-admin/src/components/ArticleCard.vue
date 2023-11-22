<script setup lang="ts">

interface Props {
  article: {
    id: string,
    title: string,
    description: string,
  }
}

defineProps<Props>()

</script>

<template>
  <v-card :title="article.title" :subtitle="article.id">
    <v-card-text>{{ article.description }}</v-card-text>
  </v-card>
</template>

<style scoped>

</style>
