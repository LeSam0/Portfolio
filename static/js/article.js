    
var json = []

async function LoadPage() {
    const queryString = window.location.search;
    const urlParams = new URLSearchParams(queryString);
    const product = urlParams.get('numero')
    const response = await fetch('/blog/article/data?numero='+product, { method: 'GET', })
    json = await response.json();
    console.log(json)
    LoadArticle()
}

function LoadArticle() {
    document.querySelector('#Image').src = "/static/image/"+json.Image;
    document.querySelector('#Date').innerHTML = json.Date;
    document.querySelector('#Title').innerHTML = json.Name;
    document.querySelector('#Context').innerHTML = json.Context;
    document.querySelector('#Text').innerHTML = json.Text;
}