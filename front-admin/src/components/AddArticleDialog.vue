<script setup lang="ts">
import {computed, ref} from "vue";

const emit = defineEmits(['onArticleAdded'])
const isDialogOpen = ref(false)
const articleName = ref('')
const articleDescription = ref('')
const isFormValid = computed(() => {
  return !!articleName.value && !!articleDescription.value
})

function saveArticle() {
  emit('onArticleAdded', {
    name: articleName.value,
    description: articleDescription.value
  })
  closeDialog()
}

function closeDialog() {
  resetForm()
  isDialogOpen.value = false
}

function resetForm() {
  articleName.value = ''
  articleDescription.value = ''
}
</script>

<template>
  <v-dialog width="500" v-model="isDialogOpen">
    <template v-slot:activator="{ props }">
      <v-btn v-bind="props" icon="mdi-plus" class="text-primary"/>
    </template>

    <v-card title="Add article">
      <v-card-text>
        <v-form>
          <v-text-field
            v-model="articleName"
            :rules="[v => !!v || 'Name cannot be empty']"
            :required="true"
            label="Article name"
          />
          <v-text-field
            v-model="articleDescription"
            :rules="[v => !!v || 'Description cannot be empty']"
            :required="true"
            label="Article description"
          />

        </v-form>
      </v-card-text>

      <v-card-actions>
        <v-btn text="save" color="primary" @click="saveArticle" :disabled="!isFormValid"/>
        <v-spacer/>
        <v-btn text="Close Dialog" @click="closeDialog"/>
      </v-card-actions>
    </v-card>
  </v-dialog>

</template>

<style scoped>

</style>
