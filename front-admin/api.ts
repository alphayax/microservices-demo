
export async function fetchArticles() {
  const articleResponse = await fetch('http://localhost:8080/')
  if (!articleResponse.ok) {
    return Promise.reject(articleResponse.statusText)
  }
  const articleData = await articleResponse.json()
  console.log(articleResponse, articleData)
  return articleData.articles || []
}

export function addArticle(article) {
  return fetch('http://localhost:8080/', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(article)
  })
}
