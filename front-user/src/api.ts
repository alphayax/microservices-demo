import type {Article} from "@/model";

interface EndpointsConfig {
  apiArticlesEndpoint: string
  apiCartEndpoint: string
}

let endpoints: EndpointsConfig = {
  apiArticlesEndpoint: "",
  apiCartEndpoint: ""
}

async function getEndpoints() : Promise<EndpointsConfig> {
  if (endpoints.apiArticlesEndpoint === "") {
    endpoints = await fetch("config/endpoints.json").then((response) => response.json())
  }
  return endpoints
}

export async function fetchArticles(): Promise<Article[]> {
  const endpoints = await getEndpoints()
  const articleResponse = await fetch(endpoints.apiArticlesEndpoint + "/")
  if (!articleResponse.ok) {
    return Promise.reject(articleResponse.statusText)
  }
  const articleData = await articleResponse.json()
  return articleData.articles || []
}

export async function fetchCart(cartId: string): Promise<string[]> {
  const endpoints = await getEndpoints()
  const cartResponse = await fetch(endpoints.apiCartEndpoint + "/" + cartId + "/")
  if (!cartResponse.ok) {
    return Promise.reject(cartResponse.statusText)
  }
  const cartData = await cartResponse.json()
  return cartData.cart.items || []
}

export async function saveCart(cartId: string, items: string[]) {
  const endpoints = await getEndpoints()
  return fetch(endpoints.apiCartEndpoint + "/" + cartId + "/", {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      items: items
    })
  })
}

export async function deleteCart(cartId: string) {
  const endpoints = await getEndpoints()
  return fetch(endpoints.apiCartEndpoint + "/" + cartId + "/", {
    method: 'DELETE'
  })
}
