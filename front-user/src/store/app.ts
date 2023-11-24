// Utilities
import { defineStore } from 'pinia'

export const useAppStore = defineStore('app', {
  state: () => ({
    cartItems: [],
    cartId: "777" // Hardcoded for demo purposes
  }),
  getters: {
    cartItemsCount() {
      return this.cartItems.length
    },
  },
  actions: {
    addToCart(item) {
      this.cartItems.push(item)
    },
    removeFromCart(item) {
      this.cartItems = this.cartItems.filter((i) => i !== item)
    }
  },
})
