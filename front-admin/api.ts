
export function fetchArticles() {
  return fetch('http://localhost:8080/')
    .then(response => response.json())
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
