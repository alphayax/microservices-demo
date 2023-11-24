import type {Article} from "@/model";

interface EndpointsConfig {
  apiArticlesEndpoint: string
}

let endpoints: EndpointsConfig = {
  apiArticlesEndpoint: ""
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

export async function addArticle(article: Article) {
  const endpoints = await getEndpoints()
  return fetch(endpoints.apiArticlesEndpoint + "/", {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(article)
  })
}
