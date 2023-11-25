// Utilities
import {defineStore} from 'pinia'
import {computed, Ref, ref} from "vue";
import {fetchCart} from "@/api";

export const useAppStore = defineStore('app', () => {
  const cartItems: Ref<string[]> = ref<string[]>([])
  const cartId: Ref<string> = ref<string>("777") // Hardcoded for demo purposes

  const cartItemsCount = computed(() => {
    return cartItems.value.length
  })

  const isItemInCart = computed(() => {
    return (itemId: string): boolean => {
      return cartItems.value.find(cartItemId => (itemId === cartItemId)) !== undefined
    }
  })

  async function loadCart() {
    cartItems.value = await fetchCart(cartId.value)
  }

  function addToCart(itemId: string) {
    cartItems.value.push(itemId)
  }

  function removeFromCart(itemId: string) {
    cartItems.value = cartItems.value.filter((cartItemId) => cartItemId !== itemId)
  }

  return {
    // state
    cartItems,
    cartId,
    // actions
    loadCart,
    addToCart,
    removeFromCart,
    // getters
    cartItemsCount,
    isItemInCart,
  }
})
