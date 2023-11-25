<script setup lang="ts">
import {computed} from "vue";
import {storeToRefs} from "pinia";
import {useAppStore} from "@/store/app";
import type {Article} from "@/model";

interface Props {
  article: Article
}

const appStore = useAppStore()
const { isItemInCart } = storeToRefs(appStore)

const props = defineProps<Props>()
const emit = defineEmits(['onArticleAdded', 'onArticleRemoved'])
const isInCart = computed(() => {
  return isItemInCart.value(props.article.id)
})

function addToCart() {
  emit('onArticleAdded', props.article.id)
}

function removeFromCart() {
  emit('onArticleRemoved', props.article.id)
}
</script>

<template>
  <v-card :title="article.title" :subtitle="article.id">
    <v-card-text>{{ article.description }}</v-card-text>
    <v-card-actions>
      <v-spacer/>
      <v-btn
        v-if="!isInCart"
        @click="addToCart"
        prepend-icon="mdi-plus"
        text="Add to cart"
        color="green"
      />
      <v-btn
        v-else
        @click="removeFromCart"
        prepend-icon="mdi-minus"
        text="Remove from cart"
        color="red"
        />
    </v-card-actions>
  </v-card>
</template>

<style scoped>

</style>
