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

export function addArticle(article: Article) {
  return fetch('http://localhost:8080/', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(article)
  })
}
